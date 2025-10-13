package forms

import (
	"github.com/rivo/tview"
)

var azureSecretName = tview.NewInputField().
	SetLabel("Azure Storage Secret Name")
var azureStorageAccountName = tview.NewInputField().
	SetLabel("Azure Storage Account Name")
var azureStorageAccountKey = tview.NewInputField().
	SetLabel("Azure Storage Account Key")

var AzureStorageSecretsForm = tview.NewForm().
	AddFormItem(azureSecretName).
	AddFormItem(azureStorageAccountKey).
	AddFormItem(azureStorageAccountName)

func init() {
	AzureStorageSecretsForm.SetBorder(true).SetTitle("Storage")
}
