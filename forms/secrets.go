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
	AddPasswordField("Database Password", "", 20, rune('*'), nil)

var createSecretButton = tview.NewButton("Create Secrets").SetSelectedFunc(func() {
	dbSecretName := SecretsForm.GetFormItemByLabel("Database Secret Name").(*tview.InputField).GetText()
	if dbSecretName != "" {
		helpers.CreateSecret(
			dbSecretName,
			map[string]string{
				"password": SecretsForm.GetFormItemByLabel("Database Password").(*tview.InputField).GetText(),
			},
		)
	}
	secretName := azureSecretName.GetText()
	if secretName != "" {
		helpers.CreateSecret(
			secretName,
			map[string]string{
				"azurestorageaccountkey":  azureStorageAccountKey.GetText(),
				"azurestorageaccountname": azureStorageAccountName.GetText(),
			},
		)
	}
	secretName = azureSslSecretName.GetText()
	if secretName != "" {
		helpers.CreateSecret(
			secretName,
			map[string]string{
				"SP_SECRET": azureSslSPSecret.GetText(),
			},
		)
	}
	cmName := azureSslConfigMapName.GetText()
	if cmName != "" {
		helpers.CreateConfigMap(
			cmName,
			map[string]string{
				"EMAIL_CERT":      azureSslEmail.GetText(),
				"HOSTED_ZONE":     azureSslHostedZone.GetText(),
				"RG_NAME":         azureSslRGName.GetText(),
				"SP_ID":           azureSslSPId.GetText(),
				"SUBSCRIPTION_ID": azureSslSubscriptionId.GetText(),
				"TENANT_ID":       azureSslTenantId.GetText(),
			},
		)
	}
	secretName = awsSSLSecretName.GetText()
	if secretName != "" {
		helpers.CreateSecret(
			secretName,
			map[string]string{
				"EMAIL_CERT": awsSSLEmail.GetText(),
				"REGION":     awsSSLRegion.GetText(),
				"ACCOUNT_ID": awsSSLAccountId.GetText(),
				"ROLE_NAME":  awsSSLRoleName.GetText(),
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
			matched, _ := regexp.MatchString("^Database.*", SecretsForm.GetFormItem(i).GetLabel())
			if !matched {
				SecretsForm.RemoveFormItem(i)
			} else {
				i++
			}
		}
	}
}
