package forms

import "github.com/rivo/tview"

var gitHubSecretName = tview.NewInputField().SetLabel("Github App Secret Name")
var gitHubSecret = tview.NewInputField().SetLabel("Github App Secret")
var gitHubClientId = tview.NewInputField().SetLabel("Github App ClientID")

var GhIdpForm = tview.NewForm().
	AddFormItem(gitHubSecretName).
	AddFormItem(gitHubSecret).
	AddFormItem(gitHubClientId)

func init() {
	GhIdpForm.
		SetTitle("IdP Config").
		SetBorder(true)
}
