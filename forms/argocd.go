package forms

import "github.com/rivo/tview"

var AutomaticSync = tview.NewCheckbox()
var BranchOrTagName = tview.NewInputField()
var AppName = tview.NewInputField()

var ArgoForm = tview.NewForm().
	AddFormItem(AppName).
	AddFormItem(AutomaticSync).
	AddFormItem(BranchOrTagName)

func init() {
	AppName.SetLabel("App name")
	AutomaticSync.SetLabel("Automatically sync changes?")
	BranchOrTagName.SetLabel("Branch or tag to monitor")
	ArgoForm.SetBorder(true).SetTitle("ArgoCD settings")
}
