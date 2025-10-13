package forms

import (
	"fn-config-helper/helpers"

	"github.com/rivo/tview"
)

var azureSslSecretName = tview.NewInputField().
	SetLabel("Azure SSL Secret Name")
var azureSslSPSecret = tview.NewInputField().
	SetLabel("Azure SSL SP Secret").
	SetMaskCharacter('*')
var azureSslConfigMapName = tview.NewInputField().
	SetLabel("Azure SSL ConfigMap Name")
var azureSslEmail = tview.NewInputField().
	SetLabel("Azure SSL Email").
	SetChangedFunc(helpers.EmailValidator)
var azureSslHostedZone = tview.NewInputField().
	SetLabel("Azure SSL Hosted Zone")
var azureSslRGName = tview.NewInputField().
	SetLabel("Azure SSL RG Name")
var azureSslSPId = tview.NewInputField().
	SetLabel("Azure SSL SP ID")
var azureSslSubscriptionId = tview.NewInputField().
	SetLabel("Azure SSL Subscription ID")
var azureSslTenantId = tview.NewInputField().
	SetLabel("Azure SSL Tenant ID")

var AzureSSLSecretsForm = tview.NewForm().
	AddFormItem(azureSslSecretName).
	AddFormItem(azureSslSPSecret).
	AddFormItem(azureSslConfigMapName).
	AddFormItem(azureSslEmail).
	AddFormItem(azureSslHostedZone).
	AddFormItem(azureSslRGName).
	AddFormItem(azureSslSPId).
	AddFormItem(azureSslSubscriptionId).
	AddFormItem(azureSslTenantId)

func init() {
	AzureSSLSecretsForm.SetBorder(true).SetTitle("SSL")
}
