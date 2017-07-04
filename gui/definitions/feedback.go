package definitions

func init() {
	add(`Feedback`, &defFeedback{})
}

type defFeedback struct{}

func (*defFeedback) String() string {
	return `<interface>
  <object class="GtkMessageDialog" id="dialog">
    <property name="title" translatable="yes">We would like to receive your feedback</property>
    <property name="text" translatable="yes">https://coy.im</property>
    <property name="secondary-text" translatable="yes">Please, let us know what you think of CoyIM.&#xA;This is the only way we can create a better tool that will keep your conversations private.</property>
    <property name="window-position">GTK_WIN_POS_CENTER</property>
    <property name="modal">true</property>
    <property name="message-type">GTK_MESSAGE_INFO</property>
    <property name="buttons">GTK_BUTTONS_CLOSE</property>
  </object>
</interface>
`
}
