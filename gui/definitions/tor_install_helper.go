package definitions

func init() {
	add(`TorInstallHelper`, &defTorInstallHelper{})
}

type defTorInstallHelper struct{}

func (*defTorInstallHelper) String() string {
	return `<interface>
  <object class="GtkDialog" id="dialog">
    <property name="window-position">GTK_WIN_POS_CENTER</property>
    <property name="border_width">7</property>
    <child internal-child="vbox">
      <object class="GtkBox" id="box">
        <property name="border-width">10</property>
        <property name="orientation">GTK_ORIENTATION_VERTICAL</property>
	<child>
          <object class="GtkLabel">
            <property name="can_focus">False</property>
            <property name="margin-top">5</property>
	    <property name="label" translatable="yes">
Install Tor:
            </property>
            <attributes>
              <attribute name="font-desc" value="&lt;Enter Value&gt; 14"/>
	      <attribute name="weight" value="semibold"/>
            </attributes>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">True</property>
            <property name="position">0</property>
          </packing>
        </child>
	<child>
          <object class="GtkLabel" id="message">
            <property name="can_focus">False</property>
	    <property name="label" translatable="yes">
    This should be easy:
            </property>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">True</property>
            <property name="position">1</property>
          </packing>
        </child>
        <child>
          <object class="GtkLabel">
            <property name="can_focus">False</property>
            <property name="margin-bottom">10</property>
	    <property name="label" translatable="yes">
    1. Go to https://www.torproject.org/

    2. Download and install Tor

    3. Start Tor or open Tor Browser

    4. Close CoyIM

    5. Reopen CoyIM

    6. Enjoy!
            </property>
            <property name="selectable">True</property>
	  </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">True</property>
            <property name="position">2</property>
          </packing>
        </child>
      </object>
    </child>
    <child internal-child="action_area">
      <object class="GtkButtonBox" id="button_box">
        <property name="orientation">GTK_ORIENTATION_HORIZONTAL</property>
        <child>
          <object class="GtkButton" id="close">
            <property name="label" translatable="yes">Close</property>
            <signal name="clicked" handler="on_close_signal" />
          </object>
        </child>
      </object>
    </child>
  </object>
</interface>
`
}
