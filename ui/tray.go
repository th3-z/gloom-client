package ui

import (
	"github.com/therecipe/qt/widgets"
)

type TrayApp struct {
	widgets.QDialog

	systemTrayIcon *widgets.QSystemTrayIcon
	systemTrayMenu *widgets.QMenu
}

func InitTrayApp() *TrayApp {
	trayApp := NewTrayApp(nil, 0)

	systemTrayMenu := widgets.NewQMenu(trayApp)
	systemTrayMenu.AddAction("alertAction")
	trayApp.systemTrayMenu = systemTrayMenu

	systemTrayIcon := widgets.NewQSystemTrayIcon(trayApp)
	systemTrayIcon.SetContextMenu(systemTrayMenu)
	trayApp.systemTrayIcon = systemTrayIcon

	systemTrayIcon.Show()
	return trayApp
}

func (trayApp *TrayApp) alertAction(something bool) {
	widgets.QMessageBox_Information(nil, "", "", widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}
