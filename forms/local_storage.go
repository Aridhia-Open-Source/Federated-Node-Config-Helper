package forms

import "github.com/rivo/tview"

var localPath = tview.NewInputField().SetLabel("Local Path")
var localDbPath = tview.NewInputField().SetLabel("Local DB Path")

var LocalStorageForm = tview.NewForm().
	AddFormItem(localPath).
	AddFormItem(localDbPath)
