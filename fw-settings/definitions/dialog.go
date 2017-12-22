package definitions

func init() {
	add(`Dialog`, &defDialog{})
}

type defDialog struct{}

func (*defDialog) String() string {
	return `
<?xml version="1.0" encoding="UTF-8"?>
<!-- Generated with glade 3.20.0 -->
<interface>
  <requires lib="gtk+" version="3.20"/>
  <object class="GtkImage" id="img_btn_add">
    <property name="visible">True</property>
    <property name="can_focus">False</property>
    <property name="icon_name">document-new-symbolic</property>
  </object>
  <object class="GtkImage" id="img_btn_search">
    <property name="visible">True</property>
    <property name="can_focus">False</property>
    <property name="icon_name">edit-find-symbolic</property>
  </object>
  <object class="GtkApplicationWindow" id="window">
    <property name="can_focus">False</property>
    <property name="title">Subgraph Firewall</property>
    <property name="window_position">center</property>
    <property name="default_width">700</property>
    <property name="default_height">600</property>
    <child>
      <object class="GtkBox" id="box1">
        <property name="can_focus">False</property>
        <property name="hexpand">True</property>
        <property name="vexpand">True</property>
        <property name="orientation">vertical</property>
        <child>
          <object class="GtkStack" id="toplevel_stack">
            <property name="can_focus">False</property>
            <property name="margin_bottom">5</property>
            <property name="transition_duration">250</property>
            <property name="transition_type">slide-left-right</property>
            <child>
              <object class="GtkScrolledWindow" id="prompt_scrollwindow">
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="shadow_type">in</property>
                <child>
                  <placeholder/>
                </child>
              </object>
              <packing>
                <property name="name">prompt</property>
                <property name="title" translatable="yes">Prompt</property>
              </packing>
            </child>
            <child>
              <object class="GtkBox">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="orientation">vertical</property>
                <child>
                  <object class="GtkRevealer" id="search_revealer">
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <property name="transition_type">slide-up</property>
                    <child>
                      <object class="GtkSearchEntry" id="search_entry">
                        <property name="visible">True</property>
                        <property name="can_focus">True</property>
                        <property name="tooltip_text" translatable="yes">Rule path or sandbox</property>
                        <property name="halign">center</property>
                        <property name="margin_top">5</property>
                        <property name="margin_bottom">5</property>
                        <property name="width_chars">48</property>
                        <property name="truncate_multiline">True</property>
                        <property name="caps_lock_warning">False</property>
                        <property name="primary_icon_name">edit-find-symbolic</property>
                        <property name="primary_icon_activatable">False</property>
                        <property name="primary_icon_sensitive">False</property>
                        <property name="placeholder_text" translatable="yes">Rule path or sandbox</property>
                      </object>
                    </child>
                  </object>
                  <packing>
                    <property name="expand">False</property>
                    <property name="fill">True</property>
                    <property name="position">0</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkNotebook" id="rulesnotebook">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="hexpand">True</property>
                    <property name="vexpand">True</property>
                    <child>
                      <object class="GtkScrolledWindow" id="swRulesPermanent">
                        <property name="visible">True</property>
                        <property name="can_focus">True</property>
                        <property name="hexpand">True</property>
                        <property name="vexpand">True</property>
                        <property name="hscrollbar_policy">never</property>
                        <property name="shadow_type">in</property>
                      </object>
                      <packing>
                        <property name="tab_expand">True</property>
                      </packing>
                    </child>
                    <child type="tab">
                      <object class="GtkLabel">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="label" translatable="yes">Permanent</property>
                      </object>
                      <packing>
                        <property name="tab_expand">True</property>
                        <property name="tab_fill">False</property>
                      </packing>
                    </child>
                    <child>
                      <object class="GtkScrolledWindow" id="swRulesSession">
                        <property name="visible">True</property>
                        <property name="can_focus">True</property>
                        <property name="hexpand">True</property>
                        <property name="vexpand">True</property>
                        <property name="hscrollbar_policy">never</property>
                        <property name="shadow_type">in</property>
                      </object>
                      <packing>
                        <property name="position">1</property>
                        <property name="tab_expand">True</property>
                      </packing>
                    </child>
                    <child type="tab">
                      <object class="GtkLabel">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="label" translatable="yes">Session</property>
                      </object>
                      <packing>
                        <property name="position">1</property>
                        <property name="tab_expand">True</property>
                        <property name="tab_fill">False</property>
                      </packing>
                    </child>
                    <child>
                      <object class="GtkScrolledWindow" id="swRulesProcess">
                        <property name="visible">True</property>
                        <property name="can_focus">True</property>
                        <property name="hexpand">True</property>
                        <property name="vexpand">True</property>
                        <property name="hscrollbar_policy">never</property>
                        <property name="shadow_type">in</property>
                      </object>
                      <packing>
                        <property name="position">2</property>
                        <property name="tab_expand">True</property>
                      </packing>
                    </child>
                    <child type="tab">
                      <object class="GtkLabel">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="label" translatable="yes">Process</property>
                      </object>
                      <packing>
                        <property name="position">2</property>
                        <property name="tab_expand">True</property>
                        <property name="tab_fill">False</property>
                      </packing>
                    </child>
                    <child>
                      <object class="GtkScrolledWindow" id="swRulesSystem">
                        <property name="visible">True</property>
                        <property name="can_focus">True</property>
                        <property name="hexpand">True</property>
                        <property name="vexpand">True</property>
                        <property name="hscrollbar_policy">never</property>
                        <property name="shadow_type">in</property>
                      </object>
                      <packing>
                        <property name="position">3</property>
                        <property name="tab_expand">True</property>
                      </packing>
                    </child>
                    <child type="tab">
                      <object class="GtkLabel">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="label" translatable="yes">System</property>
                      </object>
                      <packing>
                        <property name="position">3</property>
                        <property name="tab_expand">True</property>
                        <property name="tab_fill">False</property>
                      </packing>
                    </child>
                    <child type="action-start">
                      <object class="GtkBox">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="margin_left">5</property>
                        <property name="margin_right">5</property>
                        <property name="margin_top">2</property>
                        <property name="margin_bottom">2</property>
                        <child>
                          <object class="GtkToggleButton" id="btn_search">
                            <property name="visible">True</property>
                            <property name="can_focus">True</property>
                            <property name="receives_default">True</property>
                            <property name="image">img_btn_search</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">0</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkSeparator">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="margin_left">5</property>
                          </object>
                          <packing>
                            <property name="expand">False</property>
                            <property name="fill">True</property>
                            <property name="position">1</property>
                          </packing>
                        </child>
                      </object>
                      <packing>
                        <property name="tab_fill">False</property>
                      </packing>
                    </child>
                  </object>
                  <packing>
                    <property name="expand">False</property>
                    <property name="fill">True</property>
                    <property name="position">1</property>
                  </packing>
                </child>
              </object>
              <packing>
                <property name="name">rules</property>
                <property name="title" translatable="yes">Rules</property>
                <property name="position">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkGrid" id="config_grid">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_left">10</property>
                <property name="margin_right">10</property>
                <property name="margin_top">10</property>
                <property name="margin_bottom">10</property>
                <property name="hexpand">True</property>
                <property name="vexpand">True</property>
                <property name="row_spacing">5</property>
                <property name="column_homogeneous">True</property>
                <child>
                  <object class="GtkLabel" id="label1">
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <property name="halign">start</property>
                    <property name="margin_top">10</property>
                    <property name="margin_bottom">5</property>
                    <property name="label" translatable="yes">Prompt</property>
                    <property name="ellipsize">start</property>
                    <attributes>
                      <attribute name="weight" value="bold"/>
                    </attributes>
                  </object>
                  <packing>
                    <property name="left_attach">0</property>
                    <property name="top_attach">3</property>
                    <property name="width">2</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkLabel" id="label5">
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <property name="halign">start</property>
                    <property name="margin_bottom">5</property>
                    <property name="label" translatable="yes">Logging</property>
                    <property name="ellipsize">start</property>
                    <attributes>
                      <attribute name="weight" value="bold"/>
                    </attributes>
                  </object>
                  <packing>
                    <property name="left_attach">0</property>
                    <property name="top_attach">0</property>
                    <property name="width">2</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkLabel" id="label3">
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <property name="halign">start</property>
                    <property name="margin_left">10</property>
                    <property name="label" translatable="yes">Daemon Log Level:</property>
                  </object>
                  <packing>
                    <property name="left_attach">0</property>
                    <property name="top_attach">1</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkComboBoxText" id="level_combo">
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <property name="hexpand">True</property>
                    <property name="active">0</property>
                    <items>
                      <item id="error" translatable="yes">Error</item>
                      <item id="warning" translatable="yes">Warning</item>
                      <item id="notice" translatable="yes">Notice</item>
                      <item id="info" translatable="yes">Info</item>
                      <item id="debug" translatable="yes">Debug</item>
                    </items>
                  </object>
                  <packing>
                    <property name="left_attach">1</property>
                    <property name="top_attach">1</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkCheckButton" id="redact_checkbox">
                    <property name="label" translatable="yes">Remove host names and addresses from logs</property>
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="receives_default">False</property>
                    <property name="halign">start</property>
                    <property name="margin_left">10</property>
                    <property name="draw_indicator">True</property>
                  </object>
                  <packing>
                    <property name="left_attach">0</property>
                    <property name="top_attach">2</property>
                    <property name="width">2</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkCheckButton" id="expanded_checkbox">
                    <property name="label" translatable="yes">Always expand event prompt</property>
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="receives_default">False</property>
                    <property name="halign">start</property>
                    <property name="margin_left">10</property>
                    <property name="draw_indicator">True</property>
                  </object>
                  <packing>
                    <property name="left_attach">0</property>
                    <property name="top_attach">4</property>
                    <property name="width">2</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkCheckButton" id="expert_checkbox">
                    <property name="label" translatable="yes">Show expert options in event prompt</property>
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="receives_default">False</property>
                    <property name="halign">start</property>
                    <property name="margin_left">10</property>
                    <property name="draw_indicator">True</property>
                  </object>
                  <packing>
                    <property name="left_attach">0</property>
                    <property name="top_attach">5</property>
                    <property name="width">2</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkLabel" id="label2">
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <property name="halign">start</property>
                    <property name="margin_left">10</property>
                    <property name="label" translatable="yes">Default Scope:</property>
                  </object>
                  <packing>
                    <property name="left_attach">0</property>
                    <property name="top_attach">7</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkComboBoxText" id="action_combo">
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <property name="hexpand">True</property>
                    <property name="active">0</property>
                    <items>
                      <item id="FOREVER" translatable="yes">Forever</item>
                      <item id="SESSION" translatable="yes">Session</item>
                      <item id="PROCESS" translatable="yes">Process</item>
                      <item id="ONCE" translatable="yes">Once</item>
                    </items>
                  </object>
                  <packing>
                    <property name="left_attach">1</property>
                    <property name="top_attach">7</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkGrid" id="grib_gtkprompt">
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <property name="row_spacing">5</property>
                    <property name="column_homogeneous">True</property>
                    <child>
                      <object class="GtkCheckButton" id="toplevel_checkbox">
                        <property name="label" translatable="yes">Bring window front most and make sticky on new prompt events</property>
                        <property name="visible">True</property>
                        <property name="can_focus">True</property>
                        <property name="receives_default">False</property>
                        <property name="halign">start</property>
                        <property name="margin_left">10</property>
                        <property name="draw_indicator">True</property>
                      </object>
                      <packing>
                        <property name="left_attach">0</property>
                        <property name="top_attach">0</property>
                        <property name="width">2</property>
                      </packing>
                    </child>
                  </object>
                  <packing>
                    <property name="left_attach">0</property>
                    <property name="top_attach">6</property>
                    <property name="width">2</property>
                  </packing>
                </child>
              </object>
              <packing>
                <property name="name">config</property>
                <property name="title" translatable="yes">Options</property>
                <property name="position">2</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">True</property>
            <property name="position">0</property>
          </packing>
        </child>
      </object>
    </child>
    <child type="titlebar">
      <object class="GtkHeaderBar" id="headerbar">
        <property name="can_focus">False</property>
        <property name="show_close_button">True</property>
        <property name="decoration_layout">:minimize,maximize,close</property>
        <child>
          <object class="GtkBox" id="box_app_menu">
            <property name="can_focus">False</property>
            <property name="no_show_all">True</property>
            <property name="valign">center</property>
            <child>
              <object class="GtkMenuButton" id="btn_app_menu">
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
                <child>
                  <object class="GtkImage" id="gear_image">
                    <property name="can_focus">False</property>
                    <property name="icon_name">open-menu-symbolic</property>
                    <property name="icon_size">1</property>
                  </object>
                </child>
                <style>
                  <class name="image-button"/>
                  <class name="appmenu"/>
                </style>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkSeparator">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_left">5</property>
                <property name="orientation">vertical</property>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">1</property>
              </packing>
            </child>
          </object>
        </child>
        <child>
          <object class="GtkButton" id="btn_new_rule">
            <property name="visible">True</property>
            <property name="app_paintable">True</property>
            <property name="can_focus">True</property>
            <property name="can_default">True</property>
            <property name="has_default">True</property>
            <property name="receives_default">False</property>
            <property name="tooltip_text" translatable="yes">Add New Rule</property>
            <property name="valign">center</property>
            <property name="hexpand">False</property>
            <property name="vexpand">False</property>
            <property name="image">img_btn_add</property>
            <signal name="clicked" handler="on_btn_new_rule_clicked" swapped="no"/>
            <style>
              <class name="image-button"/>
              <class name="suggested-action"/>
              <class name="default"/>
            </style>
          </object>
          <packing>
            <property name="position">1</property>
          </packing>
        </child>
        <child type="title">
          <object class="GtkStackSwitcher" id="stack_switcher">
            <property name="can_focus">False</property>
            <property name="valign">center</property>
            <property name="icon_size">2</property>
            <property name="stack">toplevel_stack</property>
          </object>
          <packing>
            <property name="position">2</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
</interface>

`
}
