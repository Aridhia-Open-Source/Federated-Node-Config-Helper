package forms

import (
	"fn-installer/helpers"

	"github.com/rivo/tview"
)

var gitHubSecretName = tview.NewInputField().SetLabel("Github App Secret Name")
var gitHubSecret = tview.NewInputField().SetLabel("Github App Secret")
var gitHubClientId = tview.NewInputField().SetLabel("Github App ClientID")

var OutboundSettingsForm = tview.NewForm().
	AddCheckbox("Outbound mode", true, nil).
	AddDropDown("Deliver to", []string{"github", "other"}, 0, nil).
	AddFormItem(gitHubSecretName).
	AddFormItem(gitHubSecret).
	AddFormItem(gitHubClientId)

var createGHSecretButton = tview.NewButton("Create Secrets").SetSelectedFunc(func() {
	secretName := OutboundSettingsForm.GetFormItemByLabel("Github App Secret Name").(*tview.InputField).GetText()
	if secretName != "" {
		helpers.CreateSecret(
			secretName,
			map[string]string{
				"GH_SECRET":    OutboundSettingsForm.GetFormItemByLabel("Github App Secret").(*tview.InputField).GetText(),
				"GH_CLIENT_ID": OutboundSettingsForm.GetFormItemByLabel("Github App ClientID").(*tview.InputField).GetText(),
			},
		)
	}

})

var GHContainer = tview.NewFlex().SetDirection(tview.FlexRow).
	AddItem(OutboundSettingsForm, 0, 9, true).
	AddItem(createGHSecretButton, 0, 1, false)
