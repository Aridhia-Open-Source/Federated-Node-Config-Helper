package forms

import "github.com/rivo/tview"

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

var AzureSecretsForm = tview.NewForm().AddFormItem(azureSecretName).
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
