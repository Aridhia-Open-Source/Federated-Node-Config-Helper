package forms

import (
	"fn-config-helper/components"
	"strings"

	"github.com/rivo/tview"
)

var storageCapacity = components.NewCapacityInput()

var StorageSettingsForm = tview.NewForm().
	AddFormItem(storageCapacity)

func HideFields(option string, optionIndex int) {
	SecretsHide(option)
	switch strings.ToLower(option) {
	case "azure":
		StorageContainer.RemoveItem(AwsStorageForm)
		StorageContainer.RemoveItem(LocalStorageForm)
		StorageContainer.AddItem(AzureStorageForm, 0, 4, false)
	case "aws":
		StorageContainer.RemoveItem(AzureStorageForm)
		StorageContainer.RemoveItem(LocalStorageForm)
		StorageContainer.AddItem(AwsStorageForm, 0, 4, false)
	default:
		StorageContainer.RemoveItem(AzureStorageForm)
		StorageContainer.RemoveItem(AwsStorageForm)
		StorageContainer.AddItem(LocalStorageForm, 0, 4, false)
	}
}

var StorageContainer = tview.NewFlex().
	SetDirection(tview.FlexRow).
	AddItem(StorageSettingsForm, 0, 4, true)
