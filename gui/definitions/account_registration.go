package definitions

func init() {
	add(`AccountRegistration`, &defAccountRegistration{})
}

type defAccountRegistration struct{}

func (*defAccountRegistration) String() string {
	return `<interface>
  <object class="GtkAssistant" id="assistant">
    <signal name="close" handler="on_cancel_signal" />
    <signal name="cancel" handler="on_cancel_signal" />
    <signal name="prepare" handler="on_prepare" />
    <child>
      <object class="GtkBox" id="dialog">
        <property name="margin">10</property>
        <property name="spacing">10</property>
	<child>
          <object class="GtkGrid">
            <property name="margin-bottom">10</property>
            <property name="row-spacing">12</property>
            <property name="column-spacing">6</property>
	    <child>
              <object class="GtkLabel" id="server-label">
                <property name="label" translatable="yes">Recommended servers: </property>
                <property name="justify">GTK_JUSTIFY_RIGHT</property>
              </object>
              <packing>
                <property name="left-attach">0</property>
                <property name="top-attach">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkComboBoxText" id="server">
                <property name="has-entry">True</property>
              </object>
              <packing>
                <property name="left-attach">1</property>
                <property name="top-attach">1</property>
              </packing>
            </child>
	  </object>
        </child>
      </object>
      <packing>
        <property name="title" translatable="yes">1. Choose a server</property>
        <property name="complete">true</property>
      </packing>
    </child>
    <child>
      <object class="GtkBox" id="form">
        <child>
          <object class="GtkImage" id="formImage">
            <property name="margin_right">15</property>
          </object>
        </child>
	<child>
	  <object class="GtkLabel" id="formMessage">
            <property name="margin-top">20</property>
	  </object>
        </child>
        <child>
          <object class="GtkGrid" id="formGrid">
            <property name="margin-bottom">10</property>
            <property name="row-spacing">12</property>
            <property name="column-spacing">6</property>
          </object>
        </child>
      </object>
      <packing>
        <property name="title" translatable="yes">2. Create account</property>
        <property name="page-type">GTK_ASSISTANT_PAGE_CONFIRM</property>
      </packing>
    </child>
    <child>
      <object class="GtkBox">
        <child>
          <object class="GtkImage" id="doneImage">
            <property name="margin_right">15</property>
          </object>
	</child>
	<child>
          <object class="GtkLabel" id="doneMessage" />
        </child>
      </object>
      <packing>
         <property name="page-type">GTK_ASSISTANT_PAGE_SUMMARY</property>
      </packing>
    </child>
  </object>
</interface>
`
}
