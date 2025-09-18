package forms

import (
	"regexp"
	"strings"

	"github.com/rivo/tview"
)

var aksShareName = tview.NewInputField().SetLabel("Azure File Share")
var awsFileId = tview.NewInputField().SetLabel("AWS File System ID")
var awsAccessPoint = tview.NewInputField().SetLabel("AWS Access Point ID")
var localPath = tview.NewInputField().SetLabel("Local Path")
var localDbPath = tview.NewInputField().SetLabel("Local DB Path")

var StorageSettingsForm = tview.NewForm().
	AddInputField("Capacity", "1Gi", 20, nil, nil).
	AddFormItem(aksShareName).
	AddFormItem(awsFileId).
	AddFormItem(awsAccessPoint).
	AddFormItem(localPath).
	AddFormItem(localDbPath)

func HideFields(option string, optionIndex int) {
	SecretsHide(option)
	switch strings.ToLower(option) {
	case "azure":
		for i := 0; i < StorageSettingsForm.GetFormItemCount(); {
			currentLabel := StorageSettingsForm.GetFormItem(i).GetLabel()
			matched, _ := regexp.MatchString("^(Capacity|Azure).*", currentLabel)
			if !matched {
				StorageSettingsForm.RemoveFormItem(i)
			} else {
				i++
			}
		}
		if StorageSettingsForm.GetFormItemCount() == 0 {
			StorageSettingsForm.AddFormItem(aksShareName)
		}
	case "aws":
		for i := 0; i < StorageSettingsForm.GetFormItemCount(); {
			currentLabel := StorageSettingsForm.GetFormItem(i).GetLabel()
			matched, _ := regexp.MatchString("^(Capacity|AWS).*", currentLabel)
			if !matched {
				StorageSettingsForm.RemoveFormItem(i)
			} else {
				i++
			}
		}
		if StorageSettingsForm.GetFormItemCount() == 0 {
			StorageSettingsForm.AddFormItem(awsFileId).
				AddFormItem(awsAccessPoint)
		}
	default:
		for i := 0; i < StorageSettingsForm.GetFormItemCount(); {
			currentLabel := StorageSettingsForm.GetFormItem(i).GetLabel()
			matched, _ := regexp.MatchString("^(Capacity|Local).*", currentLabel)
			if !matched {
				StorageSettingsForm.RemoveFormItem(i)
			} else {
				i++
			}
		}
		if StorageSettingsForm.GetFormItemCount() == 0 {
			StorageSettingsForm.AddFormItem(localPath).
				AddFormItem(localDbPath)
		}
	}
}
