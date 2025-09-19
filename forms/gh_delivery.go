package forms

import "github.com/rivo/tview"

var gitHubOrganization = tview.NewInputField().SetLabel("Github Delivery Organization")
var gitHubRepo = tview.NewInputField().SetLabel("Github Delivery Repository")
var gitHubPemFilePath = tview.NewInputField().SetLabel("Github Delivery Pem File")
var gitHubdeliveryClientId = tview.NewInputField().SetLabel("Github Delivery ClientID")

var GhDeliveryForm = tview.NewForm().
	AddFormItem(gitHubOrganization).
	AddFormItem(gitHubRepo).
	AddFormItem(gitHubPemFilePath).
	AddFormItem(gitHubdeliveryClientId)

func init() {
	GhDeliveryForm.
		SetTitle("Github Delivery Config").
		SetBorder(true)
}
