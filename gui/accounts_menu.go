package gui

import (
	"sync"

	"github.com/twstrike/coyim/config"
	"github.com/twstrike/coyim/i18n"
	"github.com/twstrike/gotk3adapter/glibi"
	"github.com/twstrike/gotk3adapter/gtki"
)

var (
	// TODO: shouldn't this be specific to the account ID in question?
	accountChangedSignal glibi.Signal
)

var accountsLock sync.Mutex

func (u *gtkUI) buildStaticAccountsMenu(submenu gtki.Menu) {
	connectAutomaticallyItem, _ := g.gtk.CheckMenuItemNewWithMnemonic(i18n.Local("Connect On _Startup"))
	u.config.WhenLoaded(func(a *config.ApplicationConfig) {
		connectAutomaticallyItem.SetActive(a.ConnectAutomatically)
	})

	connectAutomaticallyItem.Connect("activate", func() {
		u.setConnectAllAutomatically(connectAutomaticallyItem.GetActive())
	})
	submenu.Append(connectAutomaticallyItem)

	connectAllMenu, _ := g.gtk.MenuItemNewWithMnemonic(i18n.Local("_Connect All"))
	connectAllMenu.Connect("activate", func() { u.connectAllAutomatics(true) })
	submenu.Append(connectAllMenu)

	disconnectAllMenu, _ := g.gtk.MenuItemNewWithMnemonic(i18n.Local("_Disconnect All"))
	disconnectAllMenu.Connect("activate", u.disconnectAll)
	submenu.Append(disconnectAllMenu)

	sep2, _ := g.gtk.SeparatorMenuItemNew()
	submenu.Append(sep2)

	registerAccMenu, _ := g.gtk.MenuItemNewWithMnemonic(i18n.Local("_New Account"))
	registerAccMenu.Connect("activate", u.showServerSelectionWindow)
	submenu.Append(registerAccMenu)

	addAccMenu, _ := g.gtk.MenuItemNewWithMnemonic(i18n.Local("_Add Account"))
	addAccMenu.Connect("activate", u.showAddAccountWindow)
	submenu.Append(addAccMenu)

	importMenu, _ := g.gtk.MenuItemNewWithMnemonic(i18n.Local("_Import Account"))
	importMenu.Connect("activate", u.runImporter)
	submenu.Append(importMenu)

}

func (u *gtkUI) buildAccountsMenu() {
	accountsLock.Lock()
	defer accountsLock.Unlock()

	submenu, _ := g.gtk.MenuNew()

	for _, account := range u.accounts {
		account.appendMenuTo(submenu)
	}

	if len(u.accounts) > 0 {
		sep, _ := g.gtk.SeparatorMenuItemNew()
		submenu.Append(sep)
	}

	u.buildStaticAccountsMenu(submenu)

	submenu.ShowAll()

	u.accountsMenu.SetSubmenu(submenu)
}
