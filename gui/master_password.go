package gui

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/coyim/config"
	"github.com/twstrike/coyim/i18n"
)

func (u *gtkUI) captureInitialMasterPassword(k func()) {
	dialogID := "CaptureInitialMasterPassword"
	builder := builderForDefinition(dialogID)
	dialogOb, _ := builder.GetObject(dialogID)
	pwdDialog := dialogOb.(*gtk.Dialog)

	passObj, _ := builder.GetObject("password")
	password := passObj.(*gtk.Entry)

	pass2Obj, _ := builder.GetObject("password2")
	password2 := pass2Obj.(*gtk.Entry)

	msgObj, _ := builder.GetObject("passMessage")
	messageObj := msgObj.(*gtk.Label)

	builder.ConnectSignals(map[string]interface{}{
		"on_save_signal": func() {
			passText1, _ := password.GetText()
			passText2, _ := password2.GetText()
			if len(passText1) == 0 {
				messageObj.SetLabel(i18n.Local("Password can not be empty - please try again"))
				password.GrabFocus()
			} else if passText1 != passText2 {
				messageObj.SetLabel(i18n.Local("Passwords have to be the same - please try again"))
				password.GrabFocus()
			} else {
				u.keySupplier = &onetimeSavedPassword{
					savedPassword: passText1,
					realF:         u.keySupplier,
				}
				pwdDialog.Destroy()
				k()
			}
		},
		"on_cancel_signal": func() {
			pwdDialog.Destroy()
		},
	})

	glib.IdleAdd(func() {
		pwdDialog.SetTransientFor(u.window)
		pwdDialog.ShowAll()
	})
}

func (u *gtkUI) wouldYouLikeToEncryptYourFile(k func(bool)) {
	dialogID := "AskToEncrypt"
	builder := builderForDefinition(dialogID)

	dialogOb, _ := builder.GetObject(dialogID)
	encryptDialog := dialogOb.(*gtk.MessageDialog)
	encryptDialog.SetDefaultResponse(gtk.RESPONSE_YES)
	encryptDialog.SetTransientFor(u.window)

	responseType := gtk.ResponseType(encryptDialog.Run())
	result := responseType == gtk.RESPONSE_YES
	encryptDialog.Destroy()
	k(result)
}

type onetimeSavedPassword struct {
	savedPassword string
	realF         config.KeySupplier
}

func (o *onetimeSavedPassword) Invalidate() {
	o.realF.Invalidate()
}

func (o *onetimeSavedPassword) GenerateKey(params config.EncryptionParameters) ([]byte, []byte, bool) {
	if o.savedPassword != "" {
		ourPwd := o.savedPassword
		o.savedPassword = ""

		l, r := config.GenerateKeys(ourPwd, params)
		return l, r, true
	}
	return o.realF.GenerateKey(params)
}

func (u *gtkUI) getMasterPassword(params config.EncryptionParameters) ([]byte, []byte, bool) {
	dialogID := "MasterPassword"
	builder := builderForDefinition(dialogID)
	dialogOb, _ := builder.GetObject(dialogID)
	dialog := dialogOb.(*gtk.Dialog)

	passObj, _ := builder.GetObject("password")
	password := passObj.(*gtk.Entry)
	pwdResultChan := make(chan string)

	msgObj, _ := builder.GetObject("passMessage")
	messageObj := msgObj.(*gtk.Label)

	builder.ConnectSignals(map[string]interface{}{
		"on_save_signal": func() {
			passText, _ := password.GetText()
			if len(passText) > 0 {
				messageObj.SetLabel(i18n.Local("Checking password..."))
				pwdResultChan <- passText
				close(pwdResultChan)
			}
		},
		"on_cancel_signal": func() {
			close(pwdResultChan)
			u.quit()
		},
	})

	glib.IdleAdd(func() {
		dialog.SetTransientFor(u.window)
		dialog.ShowAll()
	})

	pwd, ok := <-pwdResultChan

	if !ok {
		glib.IdleAdd(func() {
			dialog.Destroy()
		})
		return nil, nil, false
	}

	l, r := config.GenerateKeys(pwd, params)
	glib.IdleAdd(func() {
		dialog.Destroy()
	})
	return l, r, true
}
