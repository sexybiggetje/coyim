package definitions

func init() {
	add(`TorInstallHelper`, &defTorInstallHelper{})
}

type defTorInstallHelper struct{}

func (*defTorInstallHelper) String() string {
	return `<interface>
  <object class="GtkMessageDialog" id="dialog">
    <property name="window-position">GTK_WIN_POS_CENTER_ALWAYS</property>
    <property name="modal">true</property>
    <property name="border_width">7</property>
    <property name="text" translatable="yes">Install Tor</property>
    <property name="secondary_text" translatable="yes">

    This should be easy:

    1. Close CoyIM

    2. Go to https://www.torproject.org/

    3. Download and install Tor

    4. Run Tor or open Tor Browser

    5. Reopen CoyIM

    6. Enjoy!

    </property>
    <property name="buttons">GTK_BUTTONS_CLOSE</property>
  </object>
</interface>
`
}
