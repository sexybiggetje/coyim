package definitions

func init() {
	add(`FirstAccountDialog`, &defFirstAccountDialog{})
}

type defFirstAccountDialog struct{}

func (*defFirstAccountDialog) String() string {
	return `<interface>
  <object class="GtkDialog" id="dialog">
    <property name="window-position">GTK_WIN_POS_CENTER</property>
    <property name="title" translatable="yes">Setup your first account</property>
    <signal name="delete-event" handler="on_cancel_signal" />
    <child internal-child="vbox">
      <object class="GtkBox" id="Vbox">
        <property name="homogeneous">True</property>
	<property name="margin">10</property>
        <child>
          <object class="GtkLabel" id="message">
            <property name="justify">center</property>
            <property name="label" translatable="yes">Welcome to CoyIM!

Since you do not have a configured account,
we need to set up one.
            </property>
            <property name="wrap">true</property>
          </object>
          <packing>
            <property name="expand">false</property>
            <property name="fill">true</property>
            <property name="position">0</property>
          </packing>
        </child>
        <child internal-child="action_area">
          <object class="GtkButtonBox" id="button_box">
            <property name="orientation">GTK_ORIENTATION_VERTICAL</property>
            <child>
              <object class="GtkButton" id="button_register">
                <property name="label" translatable="yes">Create a new account</property>
                <signal name="clicked" handler="on_register_signal" />
              </object>
            </child>
            <child>
              <object class="GtkButton" id="button_existing">
                <property name="label" translatable="yes">Add an existing account</property>
                <signal name="clicked" handler="on_existing_signal" />
              </object>
            </child>
            <child>
              <object class="GtkButton" id="button_import">
                <property name="label" translatable="yes">Import account from your computer</property>
                <signal name="clicked" handler="on_import_signal" />
              </object>
            </child>
	    <child>
              <object class="GtkButton" id="button_cancel">
                <property name="label">Cancel</property>
                <signal name="clicked" handler="on_cancel_signal" />
              </object>
            </child>
          </object>
        </child>
      </object>
    </child>
  </object>
</interface>
`
}
