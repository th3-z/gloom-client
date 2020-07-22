package main

import (
	"os"

	"github.com/th3-z/gloom-client/ui"

	"github.com/therecipe/qt/widgets"
)

func main() {

	// needs to be called once before you can start using the QWidgets
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// create a window
	// with a minimum size of 250*200
	// and sets the title to "Hello Widgets Example"
	window := ui.InitLoginForm()
	var availableGeometry = widgets.QApplication_Desktop().AvailableGeometry2(window)
	//mw.Resize2(availableGeometry.Width()/2, (availableGeometry.Height()*2)/3)
	window.Resize2(600, 450)
	window.Move2((availableGeometry.Width()-window.Width())/2,
		(availableGeometry.Height()-window.Height())/2)

	ui.InitTrayApp()
	//window.SetWindowTitle("Hello Widgets Example")

	// create a regular widget
	// give it a QVBoxLayout
	// and make it the central widget of the window
	//widget := widgets.NewQWidget(nil, 0)
	//widget.SetLayout(widgets.NewQVBoxLayout())
	//window.SetCentralWidget(widget)

	// create a line edit
	// with a custom placeholder text
	// and add it to the central widgets layout
	//input := widgets.NewQLineEdit(nil)
	//input.SetPlaceholderText("Write something ...")
	//widget.Layout().AddWidget(input)

	// create a button
	// connect the clicked signal
	// and add it to the central widgets layout
	//button := widgets.NewQPushButton2("and click me!", nil)
	//button.ConnectClicked(func(bool) {
	//	widgets.QMessageBox_Information(nil, "OK", input.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	//})
	//widget.Layout().AddWidget(button)

	// make the window visible
	window.Show()
	app.Exec()
}
