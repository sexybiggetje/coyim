package gtk_mock

import (
	"github.com/twstrike/gotk3adapter/glib_mock"
	"github.com/twstrike/gotk3adapter/gtki"
)

type MockApplication struct {
	glib_mock.MockApplication
}

func (*MockApplication) GetActiveWindow() gtki.Window {
	return nil
}
