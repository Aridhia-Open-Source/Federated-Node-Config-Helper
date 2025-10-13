package forms

import (
	"fn-config-helper/components"
	"fn-config-helper/helpers"
	"strings"

	"github.com/rivo/tview"
)

var SecretsForm = tview.NewForm().
	AddInputField("Database Secret Name", "internal-db", 0, nil, nil).
	AddPasswordField("Database Password", "", 0, rune('*'), nil)

func init() {
	subSideMenu.SetBorder(true).SetTitle("Platforms")
	SecretsForm.SetBorder(true).SetTitle("Internal Database")
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

		userPass := FirstUserFrom.GetFormItemByLabel("User password").(*tview.InputField).GetText()
		if userPass != "" {
			helpers.CreateSecret(
				client,
				FirstUserFrom.GetFormItemByLabel("Secret Name").(*tview.InputField).GetText(),
				map[string]string{
					FirstUserFrom.GetFormItemByLabel("Password Key").(*tview.InputField).GetText(): userPass,
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

var GlobalSecretContainer = tview.NewFlex().SetDirection(tview.FlexColumn).
	AddItem(SecretsForm, 0, 1, true).
	AddItem(FirstUserFrom, 0, 1, true)

var AzureSecretContainer = tview.NewFlex().SetDirection(tview.FlexColumn).
	AddItem(AzureStorageSecretsForm, 0, 2, true).
	AddItem(AzureSSLSecretsForm, 0, 2, true)

var page = tview.NewPages().
	AddPage("Global", GlobalSecretContainer, true, false).
	AddPage("Azure", AzureSecretContainer, true, false).
	AddPage("AWS", AwsSecretsForm, true, false)

var subSideMenu = tview.NewList().ShowSecondaryText(false).
	AddItem("Global", "", '0', func() {
		page.SwitchToPage("Global")
	})

var SecretFormContainer = tview.NewFlex().SetDirection(tview.FlexRow).
	AddItem(page, 0, 8, true).
	AddItem(components.CreateSecretButton, 0, 2, false)

var SecretContainer = tview.NewFlex().SetDirection(tview.FlexColumn).
	AddItem(subSideMenu, 0, 1, true).
	AddItem(SecretFormContainer, 0, 4, true)

func SecretsHide(option string) {
	switch strings.ToLower(option) {
	case "azure":
		if hasAwsItem() {
			subSideMenu.RemoveItem(subSideMenu.FindItems("AWS", "", false, false)[0])
		}
		subSideMenu.AddItem("Azure", "", '1', func() {
			page.SwitchToPage("Azure")
		})
	case "aws":
		if hasAzureItem() {
			subSideMenu.RemoveItem(subSideMenu.FindItems("Azure", "", false, false)[0])
		}
		subSideMenu.AddItem("AWS", "", '1', func() {
			page.SwitchToPage("AWS")
		})
	default:
		if hasAzureItem() {
			subSideMenu.RemoveItem(subSideMenu.FindItems("Azure", "", false, false)[0])
		}
		if hasAwsItem() {
			subSideMenu.RemoveItem(subSideMenu.FindItems("AWS", "", false, false)[0])
		}
		page.ShowPage("Global")
	}
}

func hasAzureItem() bool {
	return len(subSideMenu.FindItems("Azure", "", false, false)) > 0
}
func hasAwsItem() bool {
	return len(subSideMenu.FindItems("AWS", "", false, false)) > 0
}
