package definitions

func init() {
	add(`PeerRequestsSMP`, &defPeerRequestsSMP{})
}

type defPeerRequestsSMP struct{}

func (*defPeerRequestsSMP) String() string {
	return `<interface>
  <object class="GtkInfoBar" id="peer_requests_smp">
    <property name="message-type">GTK_MESSAGE_WARNING</property>
    <child internal-child="content_area">
      <object class="GtkBox" id="box">
        <property name="homogeneous">false</property>
        <child>
          <object class="GtkGrid">
            <child>
              <object class="GtkLabel" id="message">
                <property name="wrap">true</property>
                <property name="margin-bottom">10</property>
              </object>
              <packing>
                <property name="left-attach">0</property>
                <property name="top-attach">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkButtonBox" id="button_box">
                <child>
                  <object class="GtkButton" id="cancel_button">
                    <property name="label" translatable="yes">Cancel</property>
                  </object>
                </child>
                <child>
                  <object class="GtkButton" id="verification_button">
                    <property name="label" translatable="yes">Finish verification</property>
                  </object>
                </child>
              </object>
              <packing>
                <property name="left-attach">0</property>
                <property name="top-attach">1</property>
              </packing>
            </child>
          </object>
        </child>
      </object>
    </child>
  </object>
</interface>
`
}
