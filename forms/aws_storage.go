package forms

import "github.com/rivo/tview"

var awsFileId = tview.NewInputField().SetLabel("AWS File System ID")
var awsAccessPoint = tview.NewInputField().SetLabel("AWS Access Point ID")

var AwsStorageForm = tview.NewForm().
	AddFormItem(awsFileId).
	AddFormItem(awsAccessPoint)
