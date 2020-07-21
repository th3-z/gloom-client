package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type LoginForm struct {
	widgets.QMainWindow

	actionLogin         *widgets.QAction
	actionResetPassword *widgets.QAction

	fieldServer   *widgets.QLineEdit
	fieldUsername *widgets.QLineEdit
	fieldPassword *widgets.QLineEdit
}

func InitLoginForm() *LoginForm {
	loginForm := NewLoginForm(nil, 0)

	loginForm.SetWindowTitle("Login - " + core.QCoreApplication_ApplicationName())

	iconPath := "/home/the_z/prog/go/src/github.com/th3-z/gloom-client/ui/res/gloom.png"
	iconPixmap := gui.NewQPixmap3(iconPath, "png", 0)
	iconPixmap = iconPixmap.Scaled2(128, 128, 0, 1)
	iconWidget := widgets.NewQLabel(nil, 0)
	iconWidget.SetPixmap(iconPixmap)
	iconWidget.SetFixedHeight(128)
	iconWidget.SetFixedWidth(128)

	iconCLayout := widgets.NewQHBoxLayout()
	iconCLayout.AddWidget(iconWidget, 0, 0)
	iconCWidget := widgets.NewQWidget(nil, 0)
	iconCWidget.SetLayout(iconCLayout)

	var (
		serverWidget   = widgets.NewQWidget(nil, 0)
		serverLayout   = widgets.NewQHBoxLayout()
		serverLabel    = widgets.NewQLabel2("Server:", loginForm, 0)
		serverLineEdit = widgets.NewQLineEdit(loginForm)
	)
	serverLineEdit.SetPlaceholderText("Server address")
	serverLayout.AddWidget(serverLabel, 0, 0)
	serverLayout.AddWidget(serverLineEdit, 0, 0)
	serverWidget.SetLayout(serverLayout)
	loginForm.fieldServer = serverLineEdit

	var (
		usernameWidget   = widgets.NewQWidget(nil, 0)
		usernameLayout   = widgets.NewQHBoxLayout()
		usernameLabel    = widgets.NewQLabel2("Username:", loginForm, 0)
		usernameLineEdit = widgets.NewQLineEdit(loginForm)
	)
	usernameLineEdit.SetPlaceholderText("Username")
	usernameLayout.AddWidget(usernameLabel, 0, 0)
	usernameLayout.AddWidget(usernameLineEdit, 0, 0)
	usernameWidget.SetLayout(usernameLayout)
	loginForm.fieldUsername = usernameLineEdit

	var (
		passwordWidget   = widgets.NewQWidget(nil, 0)
		passwordLayout   = widgets.NewQHBoxLayout()
		passwordLabel    = widgets.NewQLabel2("Password:", loginForm, 0)
		passwordLineEdit = widgets.NewQLineEdit(loginForm)
	)
	passwordLayout.AddWidget(passwordLabel, 0, 0)
	passwordLayout.AddWidget(passwordLineEdit, 0, 0)
	passwordWidget.SetLayout(passwordLayout)
	loginForm.fieldPassword = passwordLineEdit

	loginButton := widgets.NewQPushButton2("Login", loginForm)
	loginButton.ConnectClicked(loginForm.loginAction)

	var layout = widgets.NewQVBoxLayout()
	layout.AddWidget(iconCWidget, 0, 0)
	layout.AddWidget(serverWidget, 0, 0)
	layout.AddWidget(usernameWidget, 0, 0)
	layout.AddWidget(passwordWidget, 0, 0)
	layout.AddWidget(loginButton, 0, 0)

	var centralWidget = widgets.NewQWidget(loginForm, 0)
	centralWidget.SetLayout(layout)
	loginForm.SetCentralWidget(centralWidget)

	return loginForm
}

func (loginForm *LoginForm) loginAction(something bool) {
	widgets.QMessageBox_Information(nil, "", loginForm.fieldUsername.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func (loginForm *LoginForm) resetPasswordAction(something bool) {
	widgets.QMessageBox_Information(nil, "Reset action", loginForm.fieldPassword.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}
