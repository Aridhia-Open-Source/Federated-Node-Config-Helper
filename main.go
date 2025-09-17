/*
Copyright © 2025 Riccardo Casula <riccardocasula@aridhia.net>
*/
package main

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"

	"fn-installer/forms"
	"fn-installer/helpers"
	"fn-installer/pages"
)

func main() {
	app := tview.NewApplication()

	// Main Buttons
	saveButton := tview.NewButton("Save")
	saveButton.SetSelectedFunc(func() {
		getValuesAndSaveYaml()
	}).SetBorder(true)
	quitButton := tview.NewButton("Quit")
	quitButton.SetSelectedFunc(func() { app.Stop() }).SetBorder(true)

	footer := tview.NewTextView()
	footer.SetBorder(true)
	footer.SetText(fmt.Sprintf("Deploy command:\n\nhelm install federatednode -n %s -f values.yaml", helpers.State.Namespace))
	pages.NamespaceDeployment.SetChangedFunc(func(text string) {
		if text == "" {
			text = "default"
		}
		helpers.State.Namespace = text
		footer.SetText(fmt.Sprintf("Deploy command:\n\nhelm install federatednode -n %s -f values.yaml", text))
	})

	page, mainSideMenu := pages.CreateMainPage(app)

	// Main app handlers
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(mainSideMenu, 0, 8, true).
			AddItem(saveButton, 0, 1, false).
			AddItem(quitButton, 0, 1, false), 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(page, 0, 8, false).
			AddItem(footer, 0, 2, false), 0, 1, false)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func getValuesAndSaveYaml() {
	conf := helpers.Config{}
	var err error = nil

	conf.LocalDevelopment = forms.GeneralSettingsForm.GetFormItemByLabel("Is development deployment").(*tview.Checkbox).IsChecked()
	conf.TaskReview = forms.GeneralSettingsForm.GetFormItemByLabel("Use Task Result Review").(*tview.Checkbox).IsChecked()
	conf.Global.TaskReview = forms.GeneralSettingsForm.GetFormItemByLabel("Use Task Result Review").(*tview.Checkbox).IsChecked()
	conf.Smoketests = forms.GeneralSettingsForm.GetFormItemByLabel("Enable Smoketests").(*tview.Checkbox).IsChecked()
	conf.Database.Hostname = forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).GetText()
	conf.Database.User = forms.GeneralSettingsForm.GetFormItemByLabel("Database User").(*tview.InputField).GetText()
	conf.Database.Name = forms.GeneralSettingsForm.GetFormItemByLabel("Database Name").(*tview.InputField).GetText()
	conf.Database.Port, err = strconv.Atoi(forms.GeneralSettingsForm.GetFormItemByLabel("Database Port").(*tview.InputField).GetText())
	if err != nil {
		panic(err)
	}
	conf.Keycloak.Replicas, err = strconv.Atoi(forms.GeneralSettingsForm.GetFormItemByLabel("Keycloak Replicas").(*tview.InputField).GetText())
	if err != nil {
		panic(err)
	}
	conf.Database.Secret.Name = forms.GeneralSettingsForm.GetFormItemByLabel("Database Secret Name").(*tview.InputField).GetText()
	conf.Database.Secret.Key = forms.GeneralSettingsForm.GetFormItemByLabel("Database Secret Key").(*tview.InputField).GetText()

	conf.NginxIngress.Enabled = forms.NginxSettingsForm.GetFormItemByLabel("Use nginx").(*tview.Checkbox).IsChecked()
	conf.Host = forms.NginxSettingsForm.GetFormItemByLabel("Host URL").(*tview.InputField).GetText()
	conf.Global.Host = forms.NginxSettingsForm.GetFormItemByLabel("Host URL").(*tview.InputField).GetText()
	conf.NginxIngress.Controller.AllowSnippetAnnotations = forms.NginxSettingsForm.GetFormItemByLabel("Allow Snippet Annotations").(*tview.Checkbox).IsChecked()
	conf.NginxIngress.Controller.ExtraArgs.DefaultSslCertificate = forms.NginxSettingsForm.GetFormItemByLabel("Default SSL cert").(*tview.InputField).GetText()
	conf.NginxIngress.Controller.IngressClass = forms.NginxSettingsForm.GetFormItemByLabel("Ingress Class").(*tview.InputField).GetText()
	conf.NginxIngress.Controller.IngressClassResource.Name = forms.NginxSettingsForm.GetFormItemByLabel("Ingress Class").(*tview.InputField).GetText()
	conf.NginxIngress.NamespaceOverride = forms.NginxSettingsForm.GetFormItemByLabel("Namespace").(*tview.InputField).GetText()

	conf.CertManager.Enabled = forms.CertSettingsForm.GetFormItemByLabel("Use cert manager").(*tview.Checkbox).IsChecked()
	conf.CertManager.Namespace = forms.CertSettingsForm.GetFormItemByLabel("Namespace").(*tview.InputField).GetText()
	conf.CertManager.InstallCRD = forms.CertSettingsForm.GetFormItemByLabel("Install CRDs").(*tview.Checkbox).IsChecked()

	conf.OutboundMode = forms.OutboundSettingsForm.GetFormItemByLabel("Outbound mode").(*tview.Checkbox).IsChecked()
	deliveryIndex, _ := forms.OutboundSettingsForm.GetFormItemByLabel("Deliver to").(*tview.DropDown).GetCurrentOption()
	if deliveryIndex == 1 {
		conf.ControllerConfig.Delivery.Github = &helpers.DeliveryGH{Repository: "repo"}
	} else {
		conf.ControllerConfig.Delivery.Other = &helpers.DeliveryOther{Url: "test", AuthType: "Bearer"}
	}

	conf.ControllerConfig.Idp.Github.SecretName = forms.OutboundSettingsForm.GetFormItemByLabel("IDP Secret Name").(*tview.InputField).GetText()
	conf.ControllerConfig.Idp.Github.SecretKey = forms.OutboundSettingsForm.GetFormItemByLabel("IDP Secret Key").(*tview.InputField).GetText()
	conf.ControllerConfig.Idp.Github.ClientIDKey = forms.OutboundSettingsForm.GetFormItemByLabel("IDP Client ID Key").(*tview.InputField).GetText()
	_, choice := forms.StorageSettingsForm.GetFormItemByLabel("Storage type").(*tview.DropDown).GetCurrentOption()

	switch choice {
	case "aws":
		conf.OnEks = true
		awsStorage := &helpers.AwsStorage{
			FileSystemId:  forms.StorageSettingsForm.GetFormItemByLabel("AWS File System ID").(*tview.InputField).GetText(),
			AccessPointId: forms.StorageSettingsForm.GetFormItemByLabel("AWS Access Point ID").(*tview.InputField).GetText(),
		}
		conf.Storage.Aws = awsStorage
		conf.ControllerConfig.Storage.Aws = awsStorage
	case "azure":
		conf.OnAks = true
		azureStorage := &helpers.AzureStorage{
			SecretName:         forms.StorageSettingsForm.GetFormItemByLabel("Azure Storage Secret Name").(*tview.InputField).GetText(),
			ShareName:          forms.StorageSettingsForm.GetFormItemByLabel("Azure File Share").(*tview.InputField).GetText(),
			StorageAccountKey:  forms.StorageSettingsForm.GetFormItemByLabel("Azure Storage Account Secret Key").(*tview.InputField).GetText(),
			StorageAccountName: forms.StorageSettingsForm.GetFormItemByLabel("Azure Storage Account Name").(*tview.InputField).GetText(),
		}
		conf.Storage.Azure = azureStorage
		conf.ControllerConfig.Storage.Azure = azureStorage
	default:
		localStorage := &helpers.LocalStorage{
			Path:   forms.StorageSettingsForm.GetFormItemByLabel("Local Path").(*tview.InputField).GetText(),
			Dbpath: forms.StorageSettingsForm.GetFormItemByLabel("Local DB Path").(*tview.InputField).GetText(),
		}
		conf.Storage.Local = localStorage
		conf.ControllerConfig.Storage.Local = localStorage
	}

	conf.Storage.Capacity = forms.StorageSettingsForm.GetFormItemByLabel("Capacity").(*tview.InputField).GetText()
	conf.ControllerConfig.Storage.Capacity = forms.StorageSettingsForm.GetFormItemByLabel("Capacity").(*tview.InputField).GetText()

	namespaces := helpers.Namespaces{
		Keycloak:   forms.NamespacesForm.GetFormItemByLabel("Keycloak").(*tview.InputField).GetText(),
		Tasks:      forms.NamespacesForm.GetFormItemByLabel("Tasks").(*tview.InputField).GetText(),
		Controller: forms.NamespacesForm.GetFormItemByLabel("Controller").(*tview.InputField).GetText(),
	}
	conf.Global.Namespaces = namespaces
	conf.Namespaces = namespaces

	helpers.CreateYaml(&conf)
}
