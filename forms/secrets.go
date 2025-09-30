package forms

import (
	"fn-config-helper/components"
	"fn-config-helper/helpers"
	"strings"

	"github.com/rivo/tview"
)

var SecretsForm = tview.NewForm().
	AddInputField("Database Secret Name", "", 20, nil, nil).
	AddPasswordField("Database Password", "", 20, rune('*'), nil)

func init() {
	components.CreateSecretButton.SetSelectedFunc(func() {
		client, _ := helpers.NewRealKubeClient()
		dbSecretName := SecretsForm.GetFormItemByLabel("Database Secret Name").(*tview.InputField).GetText()
		dbPass := SecretsForm.GetFormItemByLabel("Database Password").(*tview.InputField).GetText()
		if dbSecretName != "" {
			if dbPass == "" {
				components.ErrorBoard.SetText("The Database Password field cannot be empty")
				return
			} else {
				components.ErrorBoard.SetText("")
			}
			helpers.CreateSecret(
				client,
				dbSecretName,
				map[string]string{
					"password": dbPass,
				},
			)
		}
		secretName := azureSecretName.GetText()
		if secretName != "" {
			if azureStorageAccountKey.GetText() == "" || azureStorageAccountName.GetText() == "" {
				components.ErrorBoard.SetText("Both storage account key and name should not be empty")
				return
			} else {
				components.ErrorBoard.SetText("")
			}
			helpers.CreateSecret(
				client,
				secretName,
				map[string]string{
					"azurestorageaccountkey":  azureStorageAccountKey.GetText(),
					"azurestorageaccountname": azureStorageAccountName.GetText(),
				},
			)
		}
		secretName = azureSslSecretName.GetText()
		if secretName != "" {
			if azureSslSPSecret.GetText() == "" {
				components.ErrorBoard.SetText("The SSL secret field cannot be empty")
				return
			} else {
				components.ErrorBoard.SetText("")
			}
			helpers.CreateSecret(
				client,
				secretName,
				map[string]string{
					"SP_SECRET": azureSslSPSecret.GetText(),
				},
			)
		}
		cmName := azureSslConfigMapName.GetText()
		if cmName != "" {
			helpers.CreateConfigMap(
				client,
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
				client,
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
}

var SecretContainer = tview.NewFlex().SetDirection(tview.FlexRow).
	AddItem(SecretsForm, 0, 3, true).
	AddItem(components.CreateSecretButton, 0, 1, false)

func SecretsHide(option string) {
	switch strings.ToLower(option) {
	case "azure":
		SecretContainer.
			RemoveItem(AwsSecretsForm).
			RemoveItem(components.CreateSecretButton)
		SecretContainer.
			AddItem(AzureSecretsForm, 0, 3, true).
			AddItem(components.CreateSecretButton, 0, 1, false)
	case "aws":
		SecretContainer.
			RemoveItem(AzureSecretsForm).
			RemoveItem(components.CreateSecretButton)
		SecretContainer.
			AddItem(AwsSecretsForm, 0, 3, true).
			AddItem(components.CreateSecretButton, 0, 1, false)
	default:
		SecretContainer.RemoveItem(AzureSecretsForm)
		SecretContainer.RemoveItem(AwsSecretsForm)
	}
}
