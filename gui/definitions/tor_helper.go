package definitions

func init() {
	add(`TorHelper`, &defTorHelper{})
}

type defTorHelper struct{}

func (*defTorHelper) String() string {
	return `<interface>
  <object class="GtkMessageDialog" id="TorHelper">
    <property name="window-position">GTK_WIN_POS_CENTER_ALWAYS</property>
    <property name="modal">true</property>
    <property name="border_width">7</property>
    <property name="text" translatable="yes">Welcome!</property>
    <property name="secondary_text" translatable="yes">
    It is highly encouraged to use it with CoyIM. Tor is a tool for anonymity.

    Would you like to install it?

    If you already have it installed, please run it.
   </property>
    <property name="buttons">GTK_BUTTONS_YES_NO</property>
  </object>
</interface>
`
}
