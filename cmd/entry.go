/*
Copyright © 2025 Riccardo Casula <riccardocasula@aridhia.net>
*/
package cmd

import (
	"os"
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
var flex = tview.NewFlex()

func Execute() {
	// Main Buttons
	loadButton := components.NewLoadConfigButton("Load yaml file")
	loadButton.SetSelectedFunc(func() {
		app.SetRoot(components.ModalPage, true)
	})
	loadButton.SetBorder(true)

	components.ConfirmButton.SetSelectedFunc(importHandler)

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
	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(components.Footer, 0, 1, false).
			AddItem(components.ErrorBoard, 0, 1, false), 0, 2, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(mainSideMenu, 0, 2, true).
			AddItem(page, 0, 8, false), 0, 7, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(saveButton, 0, 1, false).
			AddItem(loadButton, 0, 1, false).
			AddItem(quitButton, 0, 1, false), 0, 1, false)

	// components
	components.CancelModalButton.SetSelectedFunc(func() {
		app.SetRoot(flex, true)
	})

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func importHandler() {
	_, err := os.Stat(components.FilepathInput.GetText())
	if err != nil {
		components.ErrorBoard.SetText(err.Error())
	}
	mainConfig, argoConfig := helpers.ReadYAML(components.FilepathInput.GetText())
	if argoConfig != nil {
		err := yaml.Unmarshal([]byte(argoConfig.Spec.Source.Helm.Values), &mainConfig)
		if err != nil {
			panic(err)
		}
		// Set form items specifically to ArgoCD
		forms.IntroForm.GetFormItemByLabel("Deployment Namespace").(*tview.InputField).SetText(argoConfig.Spec.Destination.Namespace)
		forms.BranchOrTagName.SetText(argoConfig.Spec.Source.TargetRevision)
		forms.AppName.SetText(argoConfig.Metadata.Name)
		forms.IntroForm.GetFormItemByLabel("Use ArgoCD to deploy?").(*tview.Checkbox).SetChecked(true)
	}
	setFormsFromStucts(mainConfig)
	app.SetRoot(flex, true)
}

func getValuesAndSaveYaml() {
	conf := helpers.Config{}
	var err error = nil

	// Checkboxes
	localDevelopment := forms.GeneralSettingsForm.GetFormItemByLabel("Is development deployment").(*tview.Checkbox).IsChecked()
	conf.LocalDevelopment = &localDevelopment
	taskReview := forms.GeneralSettingsForm.GetFormItemByLabel("Use Task Result Review").(*tview.Checkbox).IsChecked()
	conf.TaskReview = &taskReview
	conf.Global.TaskReview = &taskReview
	smoketests := forms.GeneralSettingsForm.GetFormItemByLabel("Enable Smoketests").(*tview.Checkbox).IsChecked()
	conf.Smoketests = &smoketests

	// Auto cleanup
	cleanupEnabled := forms.GeneralSettingsForm.GetFormItemByLabel("Automatic Cleanup enabled").(*tview.Checkbox).IsChecked()
	conf.CleanupResults.Enabled = &cleanupEnabled
	conf.CleanupResults.CleanupTime, err = strconv.Atoi(forms.GeneralSettingsForm.GetFormItemByLabel("Cleanup Time").(*tview.InputField).GetText())
	if err != nil {
		components.ErrorBoard.SetText(err.Error())
		return
	}

	// Federated Node
	regSync := forms.GeneralSettingsForm.GetFormItemByLabel("Enable Registry Sync").(*tview.Checkbox).IsChecked()
	conf.FederatedNode.EnableRegistrySync = &regSync

	// Database
	conf.Database.Host = forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).GetText()
	enforceSSL := forms.GeneralSettingsForm.GetFormItemByLabel("Enforce DB SSL").(*tview.Checkbox).IsChecked()
	conf.Database.EnforceSSL = &enforceSSL
	if conf.Database.Host == "" {
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

	// Traefik
	traefikEnabled := forms.TraefikSettingsForm.GetFormItemByLabel("Use Traefik").(*tview.Checkbox).IsChecked()
	conf.Traefik.Enabled = &traefikEnabled
	conf.Host = forms.TraefikSettingsForm.GetFormItemByLabel("Host URL").(*tview.InputField).GetText()
	conf.Traefik.GatewayClass.Name = forms.TraefikSettingsForm.GetFormItemByLabel("Gateway Class").(*tview.InputField).GetText()
	conf.Traefik.Gateway.Name = forms.TraefikSettingsForm.GetFormItemByLabel("Gateway Name").(*tview.InputField).GetText()

	// Cert manager
	installCRD := forms.CertSettingsForm.GetFormItemByLabel("Install CRDs").(*tview.Checkbox).IsChecked()
	certMgrEnaled := forms.CertSettingsForm.GetFormItemByLabel("Use cert manager").(*tview.Checkbox).IsChecked()
	conf.CertManager.Enabled = &certMgrEnaled
	conf.CertManager.Namespace = forms.CertSettingsForm.GetFormItemByLabel("Namespace").(*tview.InputField).GetText()
	conf.CertManager.InstallCRD = &installCRD
	conf.Certs.RotationPolicy = forms.CertSettingsForm.GetFormItemByLabel("Rotation Policy").(*tview.InputField).GetText()

	// Outbound mode/controller
	outbound := forms.OutboundSettingsForm.GetFormItemByLabel("Outbound mode").(*tview.Checkbox).IsChecked()
	conf.OutboundMode = &outbound
	if *conf.OutboundMode {
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
		conf.Certs.AWS = forms.AwsSecretsForm.GetFormItemByLabel("AWS SSL Secret Name").(*tview.InputField).GetText()
	case "azure":
		conf.OnAks = true
		if traefikEnabled {
			conf.Traefik.Service.Spec.ExternalTrafficPolicy = "Local"
		}
		azureFileShare := forms.AzureStorageForm.GetFormItemByLabel("Azure File Share").(*tview.InputField).GetText()
		if azureFileShare == "" {
			components.ErrorBoard.SetText("Azure File Share should not be empty")
			return
		}
		azureStorage := &helpers.AzureStorage{
			SecretName:         forms.AzureStorageSecretsForm.GetFormItemByLabel("Azure Storage Secret Name").(*tview.InputField).GetText(),
			ShareName:          azureFileShare,
			StorageAccountKey:  "azurestorageaccountkey",
			StorageAccountName: "azurestorageaccountname",
		}
		conf.Storage.Azure = azureStorage
		conf.ControllerConfig.Storage.Azure = azureStorage
		azureSSL := &helpers.AzureCerts{
			Configmap: forms.AzureSSLSecretsForm.GetFormItemByLabel("Azure SSL ConfigMap Name").(*tview.InputField).GetText(),
			Secret:    forms.AzureSSLSecretsForm.GetFormItemByLabel("Azure SSL SP Secret").(*tview.InputField).GetText(),
		}
		conf.Certs.Azure = azureSSL
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

func setFormsFromStucts(conf *helpers.Config) {
	// Checkboxes
	devDeploy := forms.GeneralSettingsForm.GetFormItemByLabel("Is development deployment").(*tview.Checkbox)
	if conf.LocalDevelopment != nil {
		devDeploy.SetChecked(*conf.LocalDevelopment)
	}
	taskReview := forms.GeneralSettingsForm.GetFormItemByLabel("Use Task Result Review").(*tview.Checkbox)
	if conf.TaskReview != nil {

		taskReview.SetChecked(*conf.TaskReview)
	}
	smoketests := forms.GeneralSettingsForm.GetFormItemByLabel("Enable Smoketests").(*tview.Checkbox)
	if conf.Smoketests != nil {
		smoketests.SetChecked(*conf.Smoketests)
	}
	// Auto cleanup
	cleanupEnabled := forms.GeneralSettingsForm.GetFormItemByLabel("Automatic Cleanup enabled").(*tview.Checkbox)
	if conf.CleanupResults.Enabled != nil {
		cleanupEnabled.SetChecked(*conf.CleanupResults.Enabled)
	}
	forms.GeneralSettingsForm.GetFormItemByLabel("Cleanup Time").(*tview.InputField).SetText(strconv.Itoa(conf.CleanupResults.CleanupTime))
	// Federated Node
	regSync := forms.GeneralSettingsForm.GetFormItemByLabel("Enable Registry Sync").(*tview.Checkbox)
	if conf.FederatedNode.EnableRegistrySync != nil {
		regSync.SetChecked(*conf.FederatedNode.EnableRegistrySync)
	}
	// Database
	forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).SetText(conf.Database.Host)
	enforceSSL := forms.GeneralSettingsForm.GetFormItemByLabel("Enforce DB SSL").(*tview.Checkbox)
	if conf.Database.EnforceSSL != nil {
		enforceSSL.SetChecked(*conf.Database.EnforceSSL)
	}
	if conf.Database.User != "" {
		forms.GeneralSettingsForm.GetFormItemByLabel("Database User").(*tview.InputField).SetText(conf.Database.User)
	}
	if conf.Database.Name != "" {
		forms.GeneralSettingsForm.GetFormItemByLabel("Database Name").(*tview.InputField).SetText(conf.Database.Name)
	}
	if conf.Database.Port != 0 {
		forms.GeneralSettingsForm.GetFormItemByLabel("Database Port").(*tview.InputField).SetText(strconv.Itoa(conf.Database.Port))
	}
	if conf.Database.Secret.Name != "" {
		forms.SecretsForm.GetFormItemByLabel("Database Secret Name").(*tview.InputField).SetText(conf.Database.Secret.Name)
	}
	// Keycloak
	if conf.Keycloak.Replicas != 0 {
		forms.GeneralSettingsForm.GetFormItemByLabel("Keycloak Replicas").(*tview.InputField).SetText(strconv.Itoa(conf.Keycloak.Replicas))
	}
	// Storage
	if conf.Storage.Capacity != "" {
		forms.StorageSettingsForm.GetFormItemByLabel("Capacity").(*components.CapacityInput).SetText(conf.Storage.Capacity)
	}
	// Namespaces
	if conf.Namespaces.Controller != "" {
		forms.NamespacesForm.GetFormItemByLabel("Controller").(*tview.InputField).SetText(conf.Namespaces.Controller)
	}
	if conf.Namespaces.Tasks != "" {
		forms.NamespacesForm.GetFormItemByLabel("Tasks").(*tview.InputField).SetText(conf.Namespaces.Tasks)
	}
	if conf.Namespaces.Keycloak != "" {
		forms.NamespacesForm.GetFormItemByLabel("Keycloak").(*tview.InputField).SetText(conf.Namespaces.Keycloak)
	}
	// AWS
	if conf.OnEks {
		forms.AwsStorageForm.GetFormItemByLabel("AWS File System ID").(*tview.InputField).SetText(conf.Storage.Aws.FileSystemId)
		forms.AwsStorageForm.GetFormItemByLabel("AWS Access Point ID").(*tview.InputField).SetText(conf.Storage.Aws.AccessPointId)
		forms.AwsSecretsForm.GetFormItemByLabel("AWS SSL Secret Name").(*tview.InputField).SetText(conf.Certs.AWS)
	} else
	// Azure
	if conf.OnAks {
		forms.AzureStorageSecretsForm.GetFormItemByLabel("Azure Storage Secret Name").(*tview.InputField).SetText(conf.Storage.Azure.SecretName)
		forms.AzureStorageForm.GetFormItemByLabel("Azure File Share").(*tview.InputField).SetText(conf.Storage.Azure.ShareName)
		forms.AzureSSLSecretsForm.GetFormItemByLabel("Azure SSL ConfigMap Name").(*tview.InputField).SetText(conf.Certs.Azure.Configmap)
		forms.AzureSSLSecretsForm.GetFormItemByLabel("Azure SSL SP Secret").(*tview.InputField).SetText(conf.Certs.Azure.Secret)
	} else if conf.Storage.Local != nil {
		if conf.Storage.Local.Path != "" {
			forms.LocalStorageForm.GetFormItemByLabel("Local Path").(*tview.InputField).SetText(conf.Storage.Local.Path)
		}
		if conf.Storage.Local.Dbpath != "" {
			forms.LocalStorageForm.GetFormItemByLabel("Local DB Path").(*tview.InputField).SetText(conf.Storage.Local.Dbpath)
		}
	}
	// first user
	if conf.FirstUserSecret != nil {
		if conf.FirstUserSecret.Name != "" {
			forms.FirstUserFrom.GetFormItemByLabel("Secret Name").(*tview.InputField).SetText(conf.FirstUserSecret.Name)
		}
		if conf.FirstUserSecret.PassKey != "" {
			forms.FirstUserFrom.GetFormItemByLabel("Password Key").(*tview.InputField).SetText(conf.FirstUserSecret.PassKey)
		}
		forms.FirstUserFrom.GetFormItemByLabel("First Name").(*tview.InputField).SetText(conf.FirstUserSecret.FirstName)
		forms.FirstUserFrom.GetFormItemByLabel("Last Name").(*tview.InputField).SetText(conf.FirstUserSecret.LastName)
		forms.FirstUserFrom.GetFormItemByLabel("Email").(*tview.InputField).SetText(conf.FirstUserSecret.Email)
	}
	// Traefik
	if conf.Traefik.Enabled != nil {
		if *conf.Traefik.Enabled {
			useTraefikCheck := forms.TraefikSettingsForm.GetFormItemByLabel("Use Traefik").(*tview.Checkbox)
			if conf.Traefik.Enabled != nil {
				useTraefikCheck.SetChecked(*conf.Traefik.Enabled)
			}
			forms.TraefikSettingsForm.GetFormItemByLabel("Host URL").(*tview.InputField).SetText(conf.Host)

			if conf.Traefik.GatewayClass.Name != "" {
				forms.TraefikSettingsForm.GetFormItemByLabel("Gateway Class").(*tview.InputField).SetText(conf.Traefik.GatewayClass.Name)
			}
			if conf.Traefik.Gateway.Name != "" {
				forms.TraefikSettingsForm.GetFormItemByLabel("Gateway Name").(*tview.InputField).SetText(conf.Traefik.Gateway.Name)
			}
		}
	}
	// Cert Manager
	if conf.CertManager.Enabled != nil {
		if *conf.CertManager.Enabled {
			useCertMgrCheck := forms.CertSettingsForm.GetFormItemByLabel("Use cert manager").(*tview.Checkbox)
			if conf.CertManager.Enabled != nil {
				useCertMgrCheck.SetChecked(*conf.CertManager.Enabled)
			}

			installCRDs := forms.CertSettingsForm.GetFormItemByLabel("Install CRDs").(*tview.Checkbox)
			if nil != conf.CertManager.InstallCRD {
				installCRDs.SetChecked(*conf.CertManager.InstallCRD)
			}

			if conf.CertManager.Namespace != "" {
				forms.CertSettingsForm.GetFormItemByLabel("Namespace").(*tview.InputField).SetText(conf.CertManager.Namespace)
			}
			if conf.Certs.RotationPolicy != "" {
				forms.CertSettingsForm.GetFormItemByLabel("Rotation Policy").(*tview.InputField).SetText(conf.Certs.RotationPolicy)
			}
		}
	}
	// Outbound mode/controller
	if conf.OutboundMode != nil {
		if *conf.OutboundMode {
			if conf.ControllerConfig.Delivery.Github != nil {
				forms.OutboundSettingsForm.GetFormItemByLabel("Deliver to").(*tview.DropDown).SetCurrentOption(helpers.FindInList("github", forms.DeliveryOptions))
				forms.GhDeliveryForm.GetFormItemByLabel("Github Delivery Repository").(*tview.InputField).SetText(conf.ControllerConfig.Delivery.Github.Repository)
			}
			if conf.ControllerConfig.Delivery.Other != nil {
				forms.OutboundSettingsForm.GetFormItemByLabel("Deliver to").(*tview.DropDown).SetCurrentOption(helpers.FindInList("other", forms.DeliveryOptions))
				forms.OtherDeliveryForm.GetFormItemByLabel("Authentication Type").(*tview.DropDown).SetCurrentOption(helpers.FindInList(conf.ControllerConfig.Delivery.Other.AuthType, forms.AuthOptions))
				forms.OtherDeliveryForm.GetFormItemByLabel("Other Delivery Url").(*tview.InputField).SetText(conf.ControllerConfig.Delivery.Other.Url)
			}
			forms.GhIdpForm.GetFormItemByLabel("Github App Secret Name").(*tview.InputField).SetText(conf.ControllerConfig.Idp.Github.SecretName)
		}
	}
}
