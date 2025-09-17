package forms

import (
	"fn-installer/helpers"

	"github.com/rivo/tview"
)

var SecretsForm = tview.NewForm().
	AddInputField("Database Secret Name", "", 20, nil, nil).
	AddInputField("Database Username", "", 20, nil, nil).
	AddPasswordField("Database Password", "", 20, rune(1), nil).
	AddInputField("Azure Secret Name", "", 20, nil, nil).
	AddInputField("Azure Storage Account Key", "", 20, nil, nil).
	AddPasswordField("Azure Storage Account Name", "", 20, rune(1), nil).
	AddPasswordField("Azure SSL Secret Name", "ssl-sp-secret", 20, rune(1), nil).
	AddPasswordField("Azure SSL SP Secret", "", 20, rune(1), nil).
	AddPasswordField("Azure SSL ConfigMap Name", "ssl-cm", 20, rune(1), nil).
	AddPasswordField("Azure SSL Email", "", 20, rune(1), nil).
	AddPasswordField("Azure SSL Hosted Zone", "", 20, rune(1), nil).
	AddPasswordField("Azure SSL RG Name", "", 20, rune(1), nil).
	AddPasswordField("Azure SSL SP ID", "", 20, rune(1), nil).
	AddPasswordField("Azure SSL Subscription ID", "", 20, rune(1), nil).
	AddPasswordField("Azure SSL Tebabt ID", "", 20, rune(1), nil).
	AddPasswordField("AWS SSL Secret Name", "", 20, rune(1), nil).
	AddPasswordField("AWS SSL Email", "", 20, rune(1), nil).
	AddPasswordField("AWS SSL Region", "", 20, rune(1), nil).
	AddPasswordField("AWS SSL Account ID", "", 20, rune(1), nil).
	AddPasswordField("AWS SSL Role Name", "", 20, rune(1), nil)

var createSecretButton = tview.NewButton("Create Secrets").SetSelectedFunc(func() {
	secretName := SecretsForm.GetFormItemByLabel("Azure Secret Name").(*tview.InputField).GetText()
	if secretName != "" {
		helpers.CreateSecret(
			SecretsForm.GetFormItemByLabel("Database Secret Name").(*tview.InputField).GetText(),
			map[string]string{
				"username": SecretsForm.GetFormItemByLabel("Database Username").(*tview.InputField).GetText(),
				"password": SecretsForm.GetFormItemByLabel("Database Password").(*tview.InputField).GetText(),
			},
		)
	}
	secretName = SecretsForm.GetFormItemByLabel("Azure Secret Name").(*tview.InputField).GetText()
	if secretName != "" {
		helpers.CreateSecret(
			secretName,
			map[string]string{
				"azurestorageaccountkey":  SecretsForm.GetFormItemByLabel("Azure Storage Account Key").(*tview.InputField).GetText(),
				"azurestorageaccountname": SecretsForm.GetFormItemByLabel("Azure Storage Account Name").(*tview.InputField).GetText(),
			},
		)
	}
	secretName = SecretsForm.GetFormItemByLabel("Azure SSL Secret Name").(*tview.InputField).GetText()
	if secretName != "" {
		helpers.CreateSecret(
			secretName,
			map[string]string{
				"SP_SECRET": SecretsForm.GetFormItemByLabel("Azure SSL SP Secret").(*tview.InputField).GetText(),
			},
		)
	}
	cmName := SecretsForm.GetFormItemByLabel("Azure SSL Secret Name").(*tview.InputField).GetText()
	if cmName != "" {
		helpers.CreateConfigMap(
			cmName,
			map[string]string{
				"EMAIL_CERT":      SecretsForm.GetFormItemByLabel("Azure SSL SP Secret").(*tview.InputField).GetText(),
				"HOSTED_ZONE":     SecretsForm.GetFormItemByLabel("Azure SSL SP Secret").(*tview.InputField).GetText(),
				"RG_NAME":         SecretsForm.GetFormItemByLabel("Azure SSL SP Secret").(*tview.InputField).GetText(),
				"SP_ID":           SecretsForm.GetFormItemByLabel("Azure SSL SP Secret").(*tview.InputField).GetText(),
				"SUBSCRIPTION_ID": SecretsForm.GetFormItemByLabel("Azure SSL SP Secret").(*tview.InputField).GetText(),
				"TENANT_ID":       SecretsForm.GetFormItemByLabel("Azure SSL SP Secret").(*tview.InputField).GetText(),
			},
		)
	}
	secretName = SecretsForm.GetFormItemByLabel("AWS SSL Secret Name").(*tview.InputField).GetText()
	if secretName != "" {
		helpers.CreateSecret(
			secretName,
			map[string]string{
				"EMAIL_CERT": SecretsForm.GetFormItemByLabel("AWS SSL Email").(*tview.InputField).GetText(),
				"REGION":     SecretsForm.GetFormItemByLabel("AWS SSL Region").(*tview.InputField).GetText(),
				"ACCOUNT_ID": SecretsForm.GetFormItemByLabel("AWS SSL Account ID").(*tview.InputField).GetText(),
				"ROLE_NAME":  SecretsForm.GetFormItemByLabel("AWS SSL Role Name").(*tview.InputField).GetText(),
			},
		)
	}
})

var SecretContainer = tview.NewFlex().SetDirection(tview.FlexRow).
	AddItem(SecretsForm, 0, 9, true).
	AddItem(createSecretButton, 0, 1, false)
