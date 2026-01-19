package forms

import "github.com/rivo/tview"

var gitHubTriggerOrganization = tview.NewInputField().SetLabel("Github Delivery Organization")
var gitHubTriggerRepo = tview.NewInputField().SetLabel("Github Delivery Repository")
var gitHubTriggerPemFilePath = tview.NewInputField().SetLabel("Github Delivery Pem File")
var gitHubTriggerClientId = tview.NewInputField().SetLabel("Github Delivery ClientID")
var SameAsTrigger = tview.NewCheckbox().SetLabel("Same delivery configuration as trigger?")

var GhTriggerForm = tview.NewForm().
	AddFormItem(gitHubTriggerOrganization).
	AddFormItem(gitHubTriggerRepo).
	AddFormItem(gitHubTriggerPemFilePath).
	AddFormItem(gitHubTriggerClientId).
	AddFormItem(SameAsTrigger)

func init() {
	SameAsTrigger.SetChecked(false).SetDisabled(true)
	GhTriggerForm.
		SetTitle("Github Trigger Config").
		SetBorder(true)
}
