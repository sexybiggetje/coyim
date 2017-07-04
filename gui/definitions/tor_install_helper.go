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

    1. Go to https://www.torproject.org/

    2. Close CoyIM

    3. Download and install

    4. Open/run Tor Browser

    5. Open CoyIM

    </property>
    <property name="buttons">GTK_BUTTONS_CLOSE</property>
  </object>
</interface>
`
}
