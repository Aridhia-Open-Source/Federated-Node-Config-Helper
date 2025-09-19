package forms

import (
	"fmt"
	"fn-config-helper/components"
	"fn-config-helper/helpers"
	"os"
	"strings"

	"github.com/rivo/tview"
)

var OutboundSettingsForm = tview.NewForm().
	AddCheckbox("Outbound mode", true, nil).
	AddDropDown("Deliver to", []string{"none", "github", "other"}, 0, HandleDeliveryOpts)

var createGHSecretButton = tview.NewButton("Create Secrets")

func init() {
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
		AddItem(OutboundSettingsForm, 0, 4, true).
		AddItem(GhIdpForm, 0, 4, true)
}

var OutboundContainer = tview.NewFlex().SetDirection(tview.FlexRow)

func HandleDeliveryOpts(option string, optionIndex int) {
	switch strings.ToLower(option) {
	case "github":
		OutboundContainer.RemoveItem(OtherDeliveryForm)
		OutboundContainer.RemoveItem(createGHSecretButton)
		OutboundContainer.AddItem(GhDeliveryForm, 0, 5, false)
		OutboundContainer.AddItem(createGHSecretButton, 0, 1, false)
	case "other":
		OutboundContainer.RemoveItem(GhDeliveryForm)
		OutboundContainer.RemoveItem(createGHSecretButton)
		OutboundContainer.AddItem(OtherDeliveryForm, 0, 5, false)
		OutboundContainer.AddItem(createGHSecretButton, 0, 1, false)
	default:
		OutboundContainer.RemoveItem(GhDeliveryForm)
		OutboundContainer.RemoveItem(OtherDeliveryForm)
	}
}
