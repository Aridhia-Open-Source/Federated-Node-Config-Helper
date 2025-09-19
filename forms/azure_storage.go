package forms

import "github.com/rivo/tview"

var aksShareName = tview.NewInputField().SetLabel("Azure File Share")

var AzureStorageForm = tview.NewForm().
	AddFormItem(aksShareName)
