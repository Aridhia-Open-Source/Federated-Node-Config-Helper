package forms

import "github.com/rivo/tview"

var otherSecretName = tview.NewInputField().SetLabel("Other Delivery Secret Name")
var otherUrl = tview.NewInputField().SetLabel("Other Delivery Url")
var otherAuth = tview.NewInputField().SetLabel("Other Delivery Auth")

var OtherDeliveryForm = tview.NewForm().
	AddFormItem(otherSecretName).
	AddFormItem(otherUrl).
	AddFormItem(otherAuth).
	AddDropDown("Authentication Type", []string{"Basic", "Bearer", "AzCopy"}, 1, nil)

func init() {
	OtherDeliveryForm.
		SetTitle("Other Delivery Config").
		SetBorder(true)
}
