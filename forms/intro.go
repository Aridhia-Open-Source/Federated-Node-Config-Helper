package forms

import (
	"fn-installer/helpers"
	"fn-installer/state"

	"github.com/rivo/tview"
)

var createNamespaceButton = tview.NewButton("Create Namespace").
	SetSelectedFunc(func() {
		client, _ := helpers.NewRealKubeClient()
		helpers.CreateNamespace(client, namespaceText.GetText())
	})

var namespaceText = tview.NewInputField().
	SetLabel("Deployment Namespace").
	SetText("default")

var CloudPlatformDropDown = tview.NewDropDown().
	SetLabel("Choose the cloud provider (if any)").
	AddOption("None (local)", func() { state.State.CloudPlatform = "Local" }).
	AddOption("AWS", func() { state.State.CloudPlatform = "AWS" }).
	AddOption("Azure", func() { state.State.CloudPlatform = "Azure" }).
	SetSelectedFunc(HideFields).
	SetCurrentOption(0)

var IntroForm = tview.NewForm().
	AddFormItem(namespaceText).
	AddFormItem(CloudPlatformDropDown)

var IntroContainer = tview.NewFlex().SetDirection(tview.FlexRow).
	AddItem(IntroForm, 0, 3, true).
	AddItem(createNamespaceButton, 0, 1, false)
