package sgfw

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/godbus/dbus"
	"github.com/subgraph/fw-daemon/proc-coroner"
)

var gPrompter *prompter = nil

func newPrompter(dbus *dbusServer) *prompter {
	p := new(prompter)
	p.dbus = dbus
	p.cond = sync.NewCond(&p.lock)
	p.dbusObj = p.dbus.conn.Object("com.subgraph.FirewallPrompt", "/com/subgraph/FirewallPrompt")
	p.policyMap = make(map[string]*Policy)

	if gPrompter != nil {
		fmt.Println("Unexpected: global prompter variable was already set!")
	}

	gPrompter = p
	go p.promptLoop()
	return p
}

type prompter struct {
	dbus        *dbusServer
	dbusObj     dbus.BusObject
	lock        sync.Mutex
	cond        *sync.Cond
	policyMap   map[string]*Policy
	policyQueue []*Policy
}

func (p *prompter) prompt(policy *Policy) {
	p.lock.Lock()
	defer p.lock.Unlock()
	_, ok := p.policyMap[policy.sandbox+"|"+policy.path]
	if ok {
		p.cond.Signal()
		return
	}
	p.policyMap[policy.sandbox+"|"+policy.path] = policy
	log.Debugf("Saving policy key:" + policy.sandbox + "|" + policy.path)
	p.policyQueue = append(p.policyQueue, policy)
	p.cond.Signal()
}

func (p *prompter) promptLoop() {
	//	p.lock.Lock()
	for {
		p.processNextPacket()
	}
}

func (p *prompter) processNextPacket() bool {
	//fmt.Println("processNextPacket()")
	var pc pendingConnection = nil
	empty := true

	for {
		p.lock.Lock()
		pc, empty = p.nextConnection()
		p.lock.Unlock()
		if pc != nil {
			fmt.Println("Got next pending connection...")
		}
		//fmt.Println("XXX: processNextPacket() loop; empty = ", empty, " / pc = ", pc)
		if pc == nil && empty {
			time.Sleep(100 * time.Millisecond)
			return false
		} else if pc == nil {
			time.Sleep(100 * time.Millisecond)
			continue
		} else if pc != nil {
			break
		}
	}

	if pc.getPrompting() {
		log.Debugf("Skipping over already prompted connection")
		return false
	}

	pc.setPrompting(true)
	go p.processConnection(pc)
	return true
}

type PC2FDMapping struct {
	guid     string
	inode    uint64
	fd       int
	fdpath   string
	prompter *prompter
}

func dumpPendingQueues() string {
	result := ""
	pctotal := 0

	if gPrompter == nil {
		return "Cannot query pending connections; no prompts have been issued yet!"
	}

	all_policies := make([]*Policy, 0)
	gPrompter.lock.Lock()
	defer gPrompter.lock.Unlock()

	for _, policy := range gPrompter.policyMap {
		all_policies = append(all_policies, policy)
	}

	result += fmt.Sprintf("Total policies: %d\n", len(all_policies))

	for pind, policy := range all_policies {
		policy.lock.Lock()

		if len(policy.pendingQueue) > 0 {
			result += fmt.Sprintf("  Policy %d of %d (%s): #pc = %d\n", pind+1,
				len(all_policies), policy.application, len(policy.pendingQueue))

			for pcind, pc := range policy.pendingQueue {
				result += fmt.Sprintf("    %d: %s\n", pcind+1, pc.print())
				pctotal++
			}

		}

		policy.lock.Unlock()
	}

	result += "-----------------------------------\n"
	result = fmt.Sprintf("Pending Queues / total pending connections = %d\n", pctotal) + result
	return result
}

var PC2FDMap = map[string]PC2FDMapping{}
var PC2FDMapLock = &sync.Mutex{}

func monitorPromptFDs(pc pendingConnection) {
	guid := pc.getGUID()
	pid := pc.procInfo().Pid
	//leaderpid := pc.procInfo().LeaderPid
	inode := pc.procInfo().Inode
	fd := pc.procInfo().FD
	prompter := pc.getPrompter()

	//fmt.Printf("ADD TO MONITOR: %v | %v / %v / %v\n", pc.policy().application, guid, pid, fd)

	if pid == -1 || fd == -1 || prompter == nil {
		log.Warningf("Unexpected error condition occurred while adding socket fd to monitor: %d %d %v",pid, fd, prompter)
		return
	} else {
		log.Warning("No unexpected errors");
	}

	PC2FDMapLock.Lock()
	defer PC2FDMapLock.Unlock()
	var fdpath string
//	log.Warning("leaderpid:",pc.procInfo().LeaderPid) 
	//if pc.procInfo().LeaderPid != "" {
//		fdpath = fmt.Sprintf("/proc/%s/root/%d/fd/%d", leaderpid, pid, fd)
//	} else {
		fdpath = fmt.Sprintf("/proc/%d/fd/%d", pid, fd)
//	}
	PC2FDMap[guid] = PC2FDMapping{guid: guid, inode: inode, fd: fd, fdpath: fdpath, prompter: prompter}
	return
}

func dumpMonitoredFDs() string {
	PC2FDMapLock.Lock()
	defer PC2FDMapLock.Unlock()

	cnt := 1
	result := fmt.Sprintf("Monitored FDs: %v total\n", len(PC2FDMap))
	for guid, fdmon := range PC2FDMap {
		result += fmt.Sprintf("%d: %s -> [inode=%v, fd=%d, fdpath=%s]\n", cnt, guid, fdmon.inode, fdmon.fd, fdmon.fdpath)
		cnt++
	}

	result += "-----------------------------------\n"
	return result
}

func monitorPromptFDLoop() {

	for true {
		delete_guids := []string{}
		PC2FDMapLock.Lock()

		for guid, fdmon := range PC2FDMap {

			lsb, err := os.Stat(fdmon.fdpath)
			if err != nil {
				log.Warningf("Error looking up socket \"%s\": %v\n", fdmon.fdpath, err)
				delete_guids = append(delete_guids, guid)
				continue
			}

			sb, ok := lsb.Sys().(*syscall.Stat_t)
			if !ok {
				log.Warning("Not a syscall.Stat_t")
				delete_guids = append(delete_guids, guid)
				continue
			}

			inode := sb.Ino
			//			fmt.Println("+++ INODE = ", inode)

			if inode != fdmon.inode {
				fmt.Printf("inode mismatch: %v vs %v\n", inode, fdmon.inode)
				delete_guids = append(delete_guids, guid)
			}

		}

		if len(delete_guids) > 0 {
			fmt.Println("guids to delete: ", delete_guids)
		}

		saved_mappings := []PC2FDMapping{}
		for _, guid := range delete_guids {
			saved_mappings = append(saved_mappings, PC2FDMap[guid])
			delete(PC2FDMap, guid)
		}

		PC2FDMapLock.Unlock()

		for _, mapping := range saved_mappings {
			_ = mapping.prompter.dbusObj.Call("com.subgraph.FirewallPrompt.RemovePrompt", 0, mapping.guid)
			// fmt.Println("DISPOSING CALL = ", call)
			prompter := mapping.prompter
			found := false

			prompter.lock.Lock()

			for _, policy := range prompter.policyQueue {
				policy.lock.Lock()
				pcind := 0

				for pcind < len(policy.pendingQueue) {

					if policy.pendingQueue[pcind].getGUID() == mapping.guid {
						// fmt.Println("-------------- found guid to remove")
						policy.pendingQueue = append(policy.pendingQueue[:pcind], policy.pendingQueue[pcind+1:]...)
						found = true
					} else {
						pcind++
					}

				}

				policy.lock.Unlock()
			}

			if !found {
				fmt.Println("Warning: FD monitor could not find pending connection to map to removed GUID: %s", mapping.guid)
			}

			prompter.lock.Unlock()
		}

		time.Sleep(5 * time.Second)
	}

}

func InitPrompt() {
	go monitorPromptFDLoop()
}

func (p *prompter) processConnection(pc pendingConnection) {
	var scope int32
	var dres bool
	var rule string

	if pc.getPrompter() == nil {
		pc.setPrompter(p)
	}

	addr := pc.hostname()
	if addr == "" {
		addr = pc.dst().String()
	}
	policy := pc.policy()

	dststr := ""

	if pc.dst() != nil {
		dststr = pc.dst().String()
	} else {
		dststr = addr + " (via proxy resolver)"
	}

	monitorPromptFDs(pc)
	call := p.dbusObj.Call("com.subgraph.FirewallPrompt.RequestPromptAsync", 0,
		pc.getGUID(),
		policy.application,
		policy.icon,
		policy.path,
		addr,
		int32(pc.dstPort()),
		dststr,
		pc.src().String(),
		pc.proto(),
		int32(pc.procInfo().UID),
		int32(pc.procInfo().GID),
		uidToUser(pc.procInfo().UID),
		gidToGroup(pc.procInfo().GID),
		int32(pc.procInfo().Pid),
		pc.sandbox(),
		pc.socks(),
		pc.getTimestamp(),
		pc.getOptString(),
		FirewallConfig.PromptExpanded,
		FirewallConfig.PromptExpert,
		int32(FirewallConfig.DefaultActionID))

	err := call.Store(&dres)
	if err != nil {
		log.Warningf("Error sending dbus async RequestPrompt message: %v", err)
		policy.removePending(pc)
		pc.drop()
		return
	}

	if !dres {
		fmt.Println("Unexpected: fw-prompt async RequestPrompt message returned:", dres)
	}

	return

	// the prompt sends:
	// ALLOW|dest or DENY|dest
	//
	// rule string needs to be:
	// VERB|dst|class|uid:gid|sandbox|[src]

	// sometimes there's a src
	// this needs to be re-visited

	toks := strings.Split(rule, "|")
	//verb := toks[0]
	//target := toks[1]
	sandbox := ""

	if len(toks) > 2 {
		sandbox = toks[2]
	}

	tempRule := fmt.Sprintf("%s|%s", toks[0], toks[1])
	tempRule += "||-1:-1|" + sandbox + "|"

	if pc.src() != nil && !pc.src().IsLoopback() && sandbox != "" {

		//if !strings.HasSuffix(rule, "SYSTEM") && !strings.HasSuffix(rule, "||") {
		//rule += "||"
		//}
		//ule += "|||" + pc.src().String()

		//		tempRule += "||-1:-1|" + sandbox + "|" + pc.src().String()
		tempRule += pc.src().String()
	} else {
		//		tempRule += "||-1:-1|" + sandbox + "|"
	}
	r, err := policy.parseRule(tempRule, false)
	if err != nil {
		log.Warningf("Error parsing rule string returned from dbus RequestPrompt: %v", err)
		policy.removePending(pc)
		pc.drop()
		return
	}
	fscope := FilterScope(scope)
	if fscope == APPLY_SESSION {
		r.mode = RULE_MODE_SESSION
	} else if fscope == APPLY_PROCESS {
		r.mode = RULE_MODE_PROCESS
		r.pid = pc.procInfo().Pid
		pcoroner.MonitorProcess(r.pid)
	}
	if !policy.processNewRule(r, fscope) {
		p.lock.Lock()
		defer p.lock.Unlock()
		p.removePolicy(pc.policy())
	}
	if fscope == APPLY_FOREVER {
		r.mode = RULE_MODE_PERMANENT
		policy.fw.saveRules()
	}
	//log.Warningf("Prompt returning rule: %v", tempRule)
	p.dbus.emitRefresh("rules")
}

func (p *prompter) nextConnection() (pendingConnection, bool) {
	pind := 0

	if len(p.policyQueue) == 0 {
		return nil, true
	}
	//fmt.Println("policy queue len = ", len(p.policyQueue))

	for pind < len(p.policyQueue) {
		//fmt.Printf("policy loop %d of %d\n", pind, len(p.policyQueue))
		//fmt.Printf("XXX: pind = %v of %v\n", pind, len(p.policyQueue))
		policy := p.policyQueue[pind]
		pc, qempty := policy.nextPending()

		if pc == nil && qempty {
			p.removePolicy(policy)
			continue
		} else {
			pind++

			pendingOnce := make([]PendingRule, 0)
			pendingOther := make([]PendingRule, 0)

			for _, r := range policy.rulesPending {
				if r.scope == int(APPLY_ONCE) {
					pendingOnce = append(pendingOnce, r)
				} else {
					pendingOther = append(pendingOther, r)
				}
			}

			if len(pendingOnce) > 0 || len(pendingOther) > 0 {
				fmt.Printf("# pending once = %d, other = %d, pc = %p / policy = %p\n", len(pendingOnce), len(pendingOther), pc, policy)
			}

			policy.rulesPending = pendingOther

			// One time filters are all applied right here, at once.
			for _, pr := range pendingOnce {
				toks := strings.Split(pr.rule, "|")
				sandbox := ""

				if len(toks) > 2 {
					sandbox = toks[2]
				}

				tempRule := fmt.Sprintf("%s|%s", toks[0], toks[1])
				tempRule += "||-1:-1|" + sandbox + "|"

				r, err := policy.parseRule(tempRule, false)
				if err != nil {
					log.Warningf("Error parsing rule string returned from dbus RequestPrompt: %v", err)
					continue
				}

				r.mode = RuleMode(pr.scope)
				fmt.Println("+++++++ processing one time rule: ", pr.rule)
				policy.processNewRuleOnce(r, pr.guid)
			}

			//			if pc == nil && !qempty {
			if len(policy.rulesPending) > 0 {
				fmt.Println("non/once policy rules pending = ", len(policy.rulesPending))
				prule := policy.rulesPending[0]
				policy.rulesPending = append(policy.rulesPending[:0], policy.rulesPending[1:]...)
				toks := strings.Split(prule.rule, "|")
				sandbox := ""

				if len(toks) > 2 {
					sandbox = toks[2]
				}

				tempRule := fmt.Sprintf("%s|%s", toks[0], toks[1])
				tempRule += "||-1:-1|" + sandbox + "|"

				/*if pc.src() != nil && !pc.src().IsLoopback() && sandbox != "" {
					tempRule += "||-1:-1|" + sandbox + "|" + pc.src().String()
				} else {
					tempRule += "||-1:-1|" + sandbox + "|"
				}*/

				r, err := policy.parseRule(tempRule, false)
				if err != nil {
					log.Warningf("Error parsing rule string returned from dbus RequestPrompt: %v", err)
					continue
					//						policy.removePending(pc)
					//						pc.drop()
					//						return
				} else {
					fscope := FilterScope(prule.scope)
					if fscope == APPLY_SESSION {
						r.mode = RULE_MODE_SESSION
					} else if fscope == APPLY_PROCESS {
						r.mode = RULE_MODE_PROCESS
						/*r.pid = pc.procInfo().Pid
						pcoroner.MonitorProcess(r.pid)*/
					}
					if !policy.processNewRule(r, fscope) {
						//							p.lock.Lock()
						//							defer p.lock.Unlock()
						//							p.removePolicy(pc.policy())
					}
					if fscope == APPLY_FOREVER {
						r.mode = RULE_MODE_PERMANENT
						policy.fw.saveRules()
					}
					//log.Warningf("Prompt returning rule: %v", tempRule)
					p.dbus.emitRefresh("rules")
				}

			}

			if pc == nil && !qempty {
				//				log.Errorf("FIX ME: I NEED TO SLEEP ON A WAKEABLE CONDITION PROPERLY!!")
				time.Sleep(time.Millisecond * 300)
				continue
			}

			if pc != nil && pc.getPrompting() {
				fmt.Println("SKIPPING PROMPTED")
				continue
			}

			return pc, qempty
		}
	}

	return nil, true
}

func (p *prompter) removePolicy(policy *Policy) {
	var newQueue []*Policy = nil

	//	if DoMultiPrompt {
	if len(p.policyQueue) == 0 {
		log.Debugf("Skipping over zero length policy queue")
		newQueue = make([]*Policy, 0, 0)
	}
	//	}

	//	if !DoMultiPrompt || newQueue == nil {
	if newQueue == nil {
		newQueue = make([]*Policy, 0, len(p.policyQueue)-1)
	}
	for _, pol := range p.policyQueue {
		if pol != policy {
			newQueue = append(newQueue, pol)
		}
	}
	p.policyQueue = newQueue
	delete(p.policyMap, policy.sandbox+"|"+policy.path)
}

var userMap = make(map[int]string)
var groupMap = make(map[int]string)
var userMapLock = &sync.Mutex{}
var groupMapLock = &sync.Mutex{}

func lookupUser(uid int) string {
	if uid == -1 {
		return "[unknown]"
	}

	userMapLock.Lock()
	defer userMapLock.Unlock()

	u, err := user.LookupId(strconv.Itoa(uid))
	if err != nil {
		return fmt.Sprintf("%d", uid)
	}
	return u.Username
}

func lookupGroup(gid int) string {
	if gid == -1 {
		return "[unknown]"
	}

	groupMapLock.Lock()
	defer groupMapLock.Unlock()

	g, err := user.LookupGroupId(strconv.Itoa(gid))
	if err != nil {
		return fmt.Sprintf("%d", gid)
	}
	return g.Name
}

func uidToUser(uid int) string {
	uname, ok := userMap[uid]
	if ok {
		return uname
	}
	uname = lookupUser(uid)
	userMap[uid] = uname
	return uname
}

func gidToGroup(gid int) string {
	gname, ok := groupMap[gid]
	if ok {
		return gname
	}
	gname = lookupGroup(gid)
	groupMap[gid] = gname
	return gname
}
