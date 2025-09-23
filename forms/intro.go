package forms

import (
	"fmt"
	"fn-config-helper/components"
	"fn-config-helper/helpers"
	"fn-config-helper/state"

	"github.com/rivo/tview"
)

var createNamespaceButton = tview.NewButton("Create Namespace").
	SetSelectedFunc(func() {
		client, _ := helpers.NewRealKubeClient()
		helpers.CreateNamespace(client, namespaceText.GetText())
	})

var useArgoCheckbox = tview.NewCheckbox()

var namespaceText = tview.NewInputField().
	SetLabel("Deployment Namespace").
	SetText("default").
	SetChangedFunc(func(text string) {
		if !useArgoCheckbox.IsChecked() {
			if text == "" {
				text = "default"
			}
			state.State.Namespace = text
			components.Footer.SetText(fmt.Sprintf("Deploy command:\n\nhelm install federatednode -n %s -f values.yaml", text))
		}
	})

var CloudPlatformDropDown = tview.NewDropDown().
	SetLabel("Choose the cloud provider (if any)").
	AddOption("None (local)", func() { state.State.CloudPlatform = "Local" }).
	AddOption("AWS", func() { state.State.CloudPlatform = "AWS" }).
	AddOption("Azure", func() { state.State.CloudPlatform = "Azure" }).
	SetSelectedFunc(HideFields).
	SetCurrentOption(0)

var IntroForm = tview.NewForm().
	AddFormItem(namespaceText).
	AddFormItem(CloudPlatformDropDown).
	AddFormItem(useArgoCheckbox)

var IntroContainer = tview.NewFlex().SetDirection(tview.FlexRow)

func init() {
	createNamespaceButton.SetBorder(true)
	IntroContainer.AddItem(IntroForm, 0, 4, true).
		AddItem(createNamespaceButton, 0, 1, false)
	useArgoCheckbox.
		SetLabel("Deploying via ArgoCD?").
		SetChangedFunc(
			func(checked bool) {
				if checked {
					IntroContainer.RemoveItem(createNamespaceButton)
					IntroContainer.AddItem(ArgoForm, 0, 4, false)
					IntroContainer.AddItem(createNamespaceButton, 0, 1, false)
					components.Footer.SetText("Deploy command:\n\nkubectl apply -f argo-app-deployment.yaml")

				} else {
					IntroContainer.RemoveItem(ArgoForm)
					components.Footer.SetText(fmt.Sprintf("Deploy command:\n\nhelm install federatednode -n %s -f values.yaml", namespaceText.GetText()))

				}
			},
		)
}
