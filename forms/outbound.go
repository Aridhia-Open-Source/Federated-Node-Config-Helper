package forms

import (
	"fmt"
	"fn-config-helper/components"
	"fn-config-helper/helpers"
	"os"
	"strings"

	"github.com/rivo/tview"
)

var DeliveryOptions = []string{"none", "github", "other"}
var dropDownDeliveryType = tview.NewDropDown()
var OutboundSettingsForm = tview.NewForm().
	AddCheckbox("Outbound mode", true, nil).
	AddFormItem(dropDownDeliveryType)

var createGHSecretButton = tview.NewButton("Create Secrets")

func init() {
	dropDownDeliveryType.SetLabel("Deliver to").
		SetOptions(DeliveryOptions, nil).
		SetCurrentOption(0).
		SetSelectedFunc(HandleDeliveryOpts)
	SameAsTrigger.SetChangedFunc(func(checked bool) {
		_, deliveryIS := dropDownDeliveryType.GetCurrentOption()
		if !checked && deliveryIS == DeliveryOptions[1] {
			GhSecretsContainer.AddItem(GhDeliveryForm, 0, 1, false)
		} else {
			GhSecretsContainer.RemoveItem(GhDeliveryForm)
		}
	})

	createGHSecretButton.SetSelectedFunc(func() {
		client, _ := helpers.NewRealKubeClient()
		secretName := gitHubSecretName.GetText()
		if secretName != "" {
			helpers.CreateSecret(
				client,
				secretName,
				map[string]string{
					"GH_SECRET":    gitHubSecret.GetText(),
					"GH_CLIENT_ID": gitHubClientId.GetText(),
				},
			)
		}

		// Trigger secret
		secretName = fmt.Sprintf("%s-%s", strings.ToLower(gitHubTriggerOrganization.GetText()), strings.ToLower(gitHubTriggerRepo.GetText()))
		if secretName != "-" {
			// Read the pem file
			pemKey, err := os.ReadFile(gitHubTriggerPemFilePath.GetText())
			if err != nil {
				components.ErrorBoard.SetText(err.Error())
			} else {
				components.ErrorBoard.SetText("")
				helpers.CreateSecret(
					client,
					secretName,
					map[string]string{
						"key.pem":      string(pemKey),
						"GH_CLIENT_ID": gitHubTriggerClientId.GetText(),
					},
				)
			}
		}
		// If different, create the delivery one as well
		_, deliveryIS := dropDownDeliveryType.GetCurrentOption()
		if !SameAsTrigger.IsChecked() && deliveryIS == DeliveryOptions[1] {
			secretName = fmt.Sprintf("%s-%s", strings.ToLower(gitHubOrganization.GetText()), strings.ToLower(gitHubRepo.GetText()))
			if secretName != "-" {
				// Read the pem file
				pemKey, err := os.ReadFile(gitHubPemFilePath.GetText())
				if err != nil {
					components.ErrorBoard.SetText(err.Error())
				} else {
					components.ErrorBoard.SetText("")
					helpers.CreateSecret(
						client,
						secretName,
						map[string]string{
							"key.pem":      string(pemKey),
							"GH_CLIENT_ID": gitHubdeliveryClientId.GetText(),
						},
					)
				}
			}
		}
		// other delivery
		secretName = otherSecretName.GetText()
		if secretName != "" {
			helpers.CreateSecret(
				client,
				secretName,
				map[string]string{
					"auth": otherAuth.GetText(),
				},
				map[string]string{
					"url": otherUrl.GetText(),
				},
			)
		}
	})
	OutboundContainer.
		AddItem(OutboundSettingsForm, 0, 3, true).
		AddItem(GhIdpForm, 0, 6, false).
		AddItem(GhSecretsContainer, 0, 6, false).
		AddItem(createGHSecretButton, 0, 2, false)
	GhSecretsContainer.
		AddItem(GhTriggerForm, 0, 1, true)
}

var OutboundContainer = tview.NewFlex().SetDirection(tview.FlexRow)
var GhSecretsContainer = tview.NewFlex().SetDirection(tview.FlexColumn)

func HandleDeliveryOpts(option string, optionIndex int) {
	switch strings.ToLower(option) {
	case "github":
		SameAsTrigger.SetDisabled(false)
		if !SameAsTrigger.IsChecked() {
			GhSecretsContainer.AddItem(GhDeliveryForm, 0, 1, false)
		}
		GhSecretsContainer.RemoveItem(OtherDeliveryForm)
	case "other":
		GhSecretsContainer.RemoveItem(GhDeliveryForm)
		GhSecretsContainer.AddItem(OtherDeliveryForm, 0, 1, false)
		SameAsTrigger.SetDisabled(true)
	default:
		SameAsTrigger.SetDisabled(true)
		GhSecretsContainer.RemoveItem(GhDeliveryForm)
		GhSecretsContainer.RemoveItem(OtherDeliveryForm)
	}
}
