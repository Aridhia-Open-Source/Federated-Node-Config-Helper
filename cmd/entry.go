/*
Copyright © 2025 Riccardo Casula <riccardocasula@aridhia.net>
*/
package cmd

import (
	"strconv"
	"strings"

	"github.com/rivo/tview"
	"go.yaml.in/yaml/v2"

	"fn-config-helper/components"
	"fn-config-helper/forms"
	"fn-config-helper/helpers"
	"fn-config-helper/pages"
	"fn-config-helper/state"
)

var app = tview.NewApplication()

func Execute() {

	// Main Buttons
	saveButton := components.NewCreateConfigButton("Create Configuration File")
	saveButton.
		SetSelectedFunc(func() {
			getValuesAndSaveYaml()
		}).
		SetBorder(true)
	quitButton := components.NewCreateQuitButton("Quit")
	quitButton.
		SetSelectedFunc(func() { app.Stop() }).
		SetBorder(true)

	page, mainSideMenu := pages.CreateMainPage(app)

	// Main app handlers
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(components.Footer, 0, 1, false).
			AddItem(components.ErrorBoard, 0, 1, false), 0, 2, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(mainSideMenu, 0, 2, true).
			AddItem(page, 0, 8, false), 0, 7, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(saveButton, 0, 1, false).
			AddItem(quitButton, 0, 1, false), 0, 1, false)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func getValuesAndSaveYaml() {
	conf := helpers.Config{}
	var err error = nil

	// Checkboxes
	conf.LocalDevelopment = forms.GeneralSettingsForm.GetFormItemByLabel("Is development deployment").(*tview.Checkbox).IsChecked()
	conf.TaskReview = forms.GeneralSettingsForm.GetFormItemByLabel("Use Task Result Review").(*tview.Checkbox).IsChecked()
	conf.Global.TaskReview = forms.GeneralSettingsForm.GetFormItemByLabel("Use Task Result Review").(*tview.Checkbox).IsChecked()
	conf.Smoketests = forms.GeneralSettingsForm.GetFormItemByLabel("Enable Smoketests").(*tview.Checkbox).IsChecked()

	// Database
	conf.Database.Hostname = forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).GetText()
	if conf.Database.Hostname == "" {
		components.ErrorBoard.SetText("Database Hostname should not be empty")
		return
	}
	conf.Database.User = forms.GeneralSettingsForm.GetFormItemByLabel("Database User").(*tview.InputField).GetText()
	conf.Database.Name = forms.GeneralSettingsForm.GetFormItemByLabel("Database Name").(*tview.InputField).GetText()
	conf.Database.Port, err = strconv.Atoi(forms.GeneralSettingsForm.GetFormItemByLabel("Database Port").(*tview.InputField).GetText())
	if err != nil {
		components.ErrorBoard.SetText(err.Error())
		return
	}
	conf.Database.Secret.Name = forms.SecretsForm.GetFormItemByLabel("Database Secret Name").(*tview.InputField).GetText()
	conf.Database.Secret.Key = "password"

	// Keycloak
	conf.Keycloak.Replicas, err = strconv.Atoi(forms.GeneralSettingsForm.GetFormItemByLabel("Keycloak Replicas").(*tview.InputField).GetText())
	if err != nil {
		components.ErrorBoard.SetText(err.Error())
		return
	}

	// Nginx
	conf.NginxIngress.Enabled = forms.NginxSettingsForm.GetFormItemByLabel("Use nginx").(*tview.Checkbox).IsChecked()
	conf.Host = forms.NginxSettingsForm.GetFormItemByLabel("Host URL").(*tview.InputField).GetText()
	conf.Global.Host = forms.NginxSettingsForm.GetFormItemByLabel("Host URL").(*tview.InputField).GetText()
	if conf.Global.Host == "" && conf.NginxIngress.Enabled {
		components.ErrorBoard.SetText("Nginx Host URL should not be empty")
		return
	}
	conf.NginxIngress.Controller.AllowSnippetAnnotations = forms.NginxSettingsForm.GetFormItemByLabel("Allow Snippet Annotations").(*tview.Checkbox).IsChecked()
	conf.NginxIngress.Controller.ExtraArgs.DefaultSslCertificate = forms.NginxSettingsForm.GetFormItemByLabel("Default SSL cert").(*tview.InputField).GetText()
	conf.NginxIngress.Controller.IngressClass = forms.NginxSettingsForm.GetFormItemByLabel("Ingress Class").(*tview.InputField).GetText()
	conf.NginxIngress.Controller.IngressClassResource.Name = forms.NginxSettingsForm.GetFormItemByLabel("Ingress Class").(*tview.InputField).GetText()
	conf.NginxIngress.NamespaceOverride = forms.NginxSettingsForm.GetFormItemByLabel("Namespace").(*tview.InputField).GetText()

	// Cert manager
	conf.CertManager.Enabled = forms.CertSettingsForm.GetFormItemByLabel("Use cert manager").(*tview.Checkbox).IsChecked()
	conf.CertManager.Namespace = forms.CertSettingsForm.GetFormItemByLabel("Namespace").(*tview.InputField).GetText()
	conf.CertManager.InstallCRD = forms.CertSettingsForm.GetFormItemByLabel("Install CRDs").(*tview.Checkbox).IsChecked()
	conf.Certs.RotationPolicy = forms.CertSettingsForm.GetFormItemByLabel("Rotation Policy").(*tview.InputField).GetText()

	// Outbound mode/controller
	conf.OutboundMode = forms.OutboundSettingsForm.GetFormItemByLabel("Outbound mode").(*tview.Checkbox).IsChecked()
	if conf.OutboundMode {
		_, deliveryOption := forms.OutboundSettingsForm.GetFormItemByLabel("Deliver to").(*tview.DropDown).GetCurrentOption()
		if deliveryOption == "github" {
			conf.ControllerConfig.Delivery.Github = &helpers.DeliveryGH{Repository: forms.GhDeliveryForm.GetFormItemByLabel("Github Delivery Repository").(*tview.InputField).GetText()}
		} else {
			_, authType := forms.OtherDeliveryForm.GetFormItemByLabel("Authentication Type").(*tview.DropDown).GetCurrentOption()
			conf.ControllerConfig.Delivery.Other = &helpers.DeliveryOther{
				Url:      forms.OtherDeliveryForm.GetFormItemByLabel("Other Delivery Url").(*tview.InputField).GetText(),
				AuthType: authType,
			}
		}

		conf.ControllerConfig.Idp.Github.SecretName = forms.GhIdpForm.GetFormItemByLabel("Github App Secret Name").(*tview.InputField).GetText()
		conf.ControllerConfig.Idp.Github.SecretKey = "GH_SECRET"
		conf.ControllerConfig.Idp.Github.ClientIDKey = "GH_CLIENT_ID"
	}

	// Platforms
	choice := state.State.CloudPlatform

	switch strings.ToLower(choice) {
	case "aws":
		conf.OnEks = true
		awsStorage := &helpers.AwsStorage{
			FileSystemId:  forms.AwsStorageForm.GetFormItemByLabel("AWS File System ID").(*tview.InputField).GetText(),
			AccessPointId: forms.AwsStorageForm.GetFormItemByLabel("AWS Access Point ID").(*tview.InputField).GetText(),
		}
		conf.Storage.Aws = awsStorage
		conf.ControllerConfig.Storage.Aws = awsStorage
		conf.Certs.Azure.Configmap = forms.AzureStorageForm.GetFormItemByLabel("Azure SSL ConfigMap Name").(*tview.InputField).GetText()
		conf.Certs.Azure.SecretName = forms.AzureStorageForm.GetFormItemByLabel("Azure SSL SP Secret").(*tview.InputField).GetText()
	case "azure":
		conf.OnAks = true

		azureFileShare := forms.AzureStorageForm.GetFormItemByLabel("Azure File Share").(*tview.InputField).GetText()
		if azureFileShare == "" {
			components.ErrorBoard.SetText("Azure File Share should not be empty")
			return
		}
		azureStorage := &helpers.AzureStorage{
			SecretName:         forms.AzureSSLSecretsForm.GetFormItemByLabel("Azure Storage Secret Name").(*tview.InputField).GetText(),
			ShareName:          azureFileShare,
			StorageAccountKey:  "azurestorageaccountkey",
			StorageAccountName: "azurestorageaccountname",
		}
		conf.Storage.Azure = azureStorage
		conf.ControllerConfig.Storage.Azure = azureStorage
		conf.Certs.AWS = forms.AwsSecretsForm.GetFormItemByLabel("AWS SSL Secret Name").(*tview.InputField).GetText()
	default:
		localStorage := &helpers.LocalStorage{
			Path:   forms.LocalStorageForm.GetFormItemByLabel("Local Path").(*tview.InputField).GetText(),
			Dbpath: forms.LocalStorageForm.GetFormItemByLabel("Local DB Path").(*tview.InputField).GetText(),
		}
		conf.Storage.Local = localStorage
		conf.ControllerConfig.Storage.Local = localStorage
	}

	capacityField := forms.StorageSettingsForm.GetFormItemByLabel("Capacity").(*components.CapacityInput)
	if !capacityField.IsValid() {
		components.ErrorBoard.SetText(capacityField.ValidationErrorMessage)
		return
	}
	conf.Storage.Capacity = capacityField.GetText()
	conf.ControllerConfig.Storage.Capacity = capacityField.GetText()

	// Namespaces
	namespaces := helpers.Namespaces{
		Keycloak:   forms.NamespacesForm.GetFormItemByLabel("Keycloak").(*tview.InputField).GetText(),
		Tasks:      forms.NamespacesForm.GetFormItemByLabel("Tasks").(*tview.InputField).GetText(),
		Controller: forms.NamespacesForm.GetFormItemByLabel("Controller").(*tview.InputField).GetText(),
	}
	conf.Global.Namespaces = namespaces
	conf.Namespaces = namespaces

	// First User
	if forms.FirstUserFrom.GetFormItemByLabel("User password").(*tview.InputField).GetText() != "" {
		conf.FirstUserSecret = &helpers.FirstUser{
			Name:      forms.FirstUserFrom.GetFormItemByLabel("Secret Name").(*tview.InputField).GetText(),
			PassKey:   forms.FirstUserFrom.GetFormItemByLabel("Password Key").(*tview.InputField).GetText(),
			FirstName: forms.FirstUserFrom.GetFormItemByLabel("First Name").(*tview.InputField).GetText(),
			LastName:  forms.FirstUserFrom.GetFormItemByLabel("Last Name").(*tview.InputField).GetText(),
			Email:     forms.FirstUserFrom.GetFormItemByLabel("Email").(*tview.InputField).GetText(),
		}
	}

	// ArgoCD
	if forms.IntroForm.GetFormItemByLabel("Use ArgoCD to deploy?").(*tview.Checkbox).IsChecked() {
		argo := helpers.InitArgoStruct()
		argo.Metadata.Name = forms.AppName.GetText()
		argo.Spec.Project = "default"
		argo.Metadata.Namespace = "argocd"
		argo.Spec.Destination.Namespace = forms.IntroForm.GetFormItemByLabel("Deployment Namespace").(*tview.InputField).GetText()
		argo.Spec.Source.TargetRevision = forms.BranchOrTagName.GetText()

		stringConf, err := yaml.Marshal(&conf)
		if err != nil {
			panic(err)
		}

		argo.Spec.Source.Helm.Values = string(stringConf)
		argo.Spec.SyncPolicy.SyncOptions = []string{"RespectIgnoreDifferences=true"}
		if forms.AutomaticSync.IsChecked() {
			argo.Spec.SyncPolicy.Automated = &helpers.Automated{}
		}

		argo.CreateYaml("argo-app-deployment.yaml")
	} else {
		conf.CreateYaml("values.yaml")
	}
}
