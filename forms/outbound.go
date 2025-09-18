package forms

import (
	"fmt"
	"fn-installer/components"
	"fn-installer/helpers"
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
		secretName := gitHubSecretName.GetText()
		if secretName != "" {
			helpers.CreateSecret(
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
	GHContainer.
		AddItem(OutboundSettingsForm, 0, 4, true).
		AddItem(GhIdpForm, 0, 4, true)
}

var GHContainer = tview.NewFlex().SetDirection(tview.FlexRow)

func HandleDeliveryOpts(option string, optionIndex int) {
	switch strings.ToLower(option) {
	case "github":
		GHContainer.RemoveItem(OtherDeliveryForm)
		GHContainer.RemoveItem(createGHSecretButton)
		GHContainer.AddItem(GhDeliveryForm, 0, 5, false)
		GHContainer.AddItem(createGHSecretButton, 0, 1, false)
	case "other":
		GHContainer.RemoveItem(GhDeliveryForm)
		GHContainer.RemoveItem(createGHSecretButton)
		GHContainer.AddItem(OtherDeliveryForm, 0, 5, false)
		GHContainer.AddItem(createGHSecretButton, 0, 1, false)
	default:
		GHContainer.RemoveItem(GhDeliveryForm)
		GHContainer.RemoveItem(OtherDeliveryForm)
	}
}
