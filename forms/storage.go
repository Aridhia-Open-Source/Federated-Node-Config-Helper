package forms

import (
	"fn-installer/helpers"
	"strings"

	"github.com/rivo/tview"
)

var storageChoices = []string{"azure", "aws", "local"}

var aksSecretName = tview.NewInputField().SetLabel("Azure Storage Secret Name")
var aksStorageAccountKey = tview.NewInputField().SetLabel("Azure Storage Account Secret Key")
var awsStorageAccountName = tview.NewInputField().SetLabel("Azure Storage Account Name")
var aksShareName = tview.NewInputField().SetLabel("Azure File Share")
var awsFileId = tview.NewInputField().SetLabel("AWS File System ID")
var awsAccessPoint = tview.NewInputField().SetLabel("AWS Access Point ID")
var localPath = tview.NewInputField().SetLabel("Local Path")
var localDbPath = tview.NewInputField().SetLabel("Local DB Path")

func getChoiceFromState() int {
	for i, st := range storageChoices {
		if st == strings.ToLower(helpers.State.CloudPlatform) {
			return i
		}
	}
	return 0
}

var StorageSettingsForm = tview.NewForm().
	AddInputField("Capacity", "1Gi", 20, nil, nil).
	AddDropDown("Storage type", storageChoices, getChoiceFromState(), hideFields).
	AddFormItem(aksSecretName).
	AddFormItem(aksStorageAccountKey).
	AddFormItem(awsStorageAccountName).
	AddFormItem(aksShareName).
	AddFormItem(awsFileId).
	AddFormItem(awsAccessPoint).
	AddFormItem(localPath).
	AddFormItem(localDbPath)

func hideFields(option string, optionIndex int) {
	switch option {
	case "azure":
		aksSecretName.SetDisabled(false)
		aksStorageAccountKey.SetDisabled(false)
		awsStorageAccountName.SetDisabled(false)
		aksShareName.SetDisabled(false)
		awsFileId.SetDisabled(true)
		awsAccessPoint.SetDisabled(true)
		localPath.SetDisabled(true)
		localDbPath.SetDisabled(true)
	case "aws":
		aksSecretName.SetDisabled(true)
		aksStorageAccountKey.SetDisabled(true)
		awsStorageAccountName.SetDisabled(true)
		aksShareName.SetDisabled(true)
		awsFileId.SetDisabled(false)
		awsAccessPoint.SetDisabled(false)
		localPath.SetDisabled(true)
		localDbPath.SetDisabled(true)
	case "local":
		aksSecretName.SetDisabled(true)
		aksStorageAccountKey.SetDisabled(true)
		awsStorageAccountName.SetDisabled(true)
		aksShareName.SetDisabled(true)
		awsFileId.SetDisabled(true)
		awsAccessPoint.SetDisabled(true)
		localPath.SetDisabled(false)
		localDbPath.SetDisabled(false)
	}
}
