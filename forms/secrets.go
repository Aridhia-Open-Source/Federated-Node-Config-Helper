package forms

import (
	"fn-installer/helpers"
	"regexp"
	"strings"

	"github.com/rivo/tview"
)

var azureSecretName = tview.NewInputField().SetLabel("Azure Secret Name")
var azureStorageAccountName = tview.NewInputField().SetLabel("Azure Storage Account Name")
var azureStorageAccountKey = tview.NewInputField().SetLabel("Azure Storage Account Key")
var azureSslSecretName = tview.NewInputField().SetLabel("Azure SSL Secret Name")
var azureSslSPSecret = tview.NewInputField().SetLabel("Azure SSL SP Secret").SetMaskCharacter('*')
var azureSslConfigMapName = tview.NewInputField().SetLabel("Azure SSL ConfigMap Name")
var azureSslEmail = tview.NewInputField().SetLabel("Azure SSL Email")
var azureSslHostedZone = tview.NewInputField().SetLabel("Azure SSL Hosted Zone")
var azureSslRGName = tview.NewInputField().SetLabel("Azure SSL RG Name")
var azureSslSPId = tview.NewInputField().SetLabel("Azure SSL SP ID")
var azureSslSubscriptionId = tview.NewInputField().SetLabel("Azure SSL Subscription ID")
var azureSslTenantId = tview.NewInputField().SetLabel("Azure SSL Tenant ID")
var awsSSLSecretName = tview.NewInputField().SetLabel("AWS SSL Secret Name")
var awsSSLEmail = tview.NewInputField().SetLabel("AWS SSL Email")
var awsSSLRegion = tview.NewInputField().SetLabel("AWS SSL Account ID")
var awsSSLAccountId = tview.NewInputField().SetLabel("AWS SSL Account ID")
var awsSSLRoleName = tview.NewInputField().SetLabel("AWS SSL Role Name")

var SecretsForm = tview.NewForm().
	AddInputField("Database Secret Name", "", 20, nil, nil).
	AddInputField("Database Username", "", 20, nil, nil).
	AddPasswordField("Database Password", "", 20, rune(1), nil).
	AddFormItem(azureSecretName).
	AddFormItem(azureStorageAccountKey).
	AddFormItem(azureStorageAccountName).
	AddFormItem(azureSslSecretName).
	AddFormItem(azureSslSPSecret).
	AddFormItem(azureSslConfigMapName).
	AddFormItem(azureSslEmail).
	AddFormItem(azureSslHostedZone).
	AddFormItem(azureSslRGName).
	AddFormItem(azureSslSPId).
	AddFormItem(azureSslSubscriptionId).
	AddFormItem(azureSslTenantId).
	AddFormItem(awsSSLSecretName).
	AddFormItem(awsSSLEmail).
	AddFormItem(awsSSLRegion).
	AddFormItem(awsSSLAccountId).
	AddFormItem(awsSSLRoleName)

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

func SecretsHide(option string) {
	switch strings.ToLower(option) {
	case "azure":
		for i := 0; i < SecretsForm.GetFormItemCount(); {
			matched, _ := regexp.MatchString("^(Database|Azure).*", SecretsForm.GetFormItem(i).GetLabel())
			if !matched {
				SecretsForm.RemoveFormItem(i)
			} else {
				i++
			}
		}
		if SecretsForm.GetFormItemCount() == 3 {
			SecretsForm.AddFormItem(azureSecretName).
				AddFormItem(azureStorageAccountKey).
				AddFormItem(azureStorageAccountName).
				AddFormItem(azureSslSecretName).
				AddFormItem(azureSslSPSecret).
				AddFormItem(azureSslConfigMapName).
				AddFormItem(azureSslEmail).
				AddFormItem(azureSslHostedZone).
				AddFormItem(azureSslRGName).
				AddFormItem(azureSslSPId).
				AddFormItem(azureSslSubscriptionId).
				AddFormItem(azureSslTenantId)
		}
	case "aws":
		for i := 0; i < SecretsForm.GetFormItemCount(); {
			currentLabel := SecretsForm.GetFormItem(i).GetLabel()
			matched, _ := regexp.MatchString("^(Database|AWS).*", currentLabel)
			if !matched {
				SecretsForm.RemoveFormItem(i)
			} else {
				i++
			}
		}
		if SecretsForm.GetFormItemCount() == 3 {
			SecretsForm.AddFormItem(awsSSLSecretName).
				AddFormItem(awsSSLEmail).
				AddFormItem(awsSSLRegion).
				AddFormItem(awsSSLAccountId).
				AddFormItem(awsSSLRoleName)
		}
	default:
		for i := 0; i < SecretsForm.GetFormItemCount(); i++ {
			SecretsForm.RemoveFormItem(i)
		}
	}
}
