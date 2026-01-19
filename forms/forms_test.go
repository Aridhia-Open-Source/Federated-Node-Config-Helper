package forms

import (
	"fmt"
	"fn-config-helper/components"
	"os"
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	CloudPlatformDropDown.SetCurrentOption(dropdownChoices["local"])
	exitcode := m.Run()
	os.Exit(exitcode)
}

var dropdownChoices = map[string]int{
	"local": 0,
	"aws":   1,
	"azure": 2,
}

func TestStorageDropdownBase(t *testing.T) {
	/*
		Default choice is local
	*/
	CloudPlatformDropDown.SetCurrentOption(dropdownChoices["local"])
	assert.NotNil(t,
		StorageContainer.GetItem(1).(*tview.Form).GetFormItemByLabel("Local Path"),
		"Local storage form not found",
	)
}
func TestStorageDropdownAzure(t *testing.T) {
	/*
		When setting the dropdown to azure, the correct
		form is rendered
	*/
	CloudPlatformDropDown.SetCurrentOption(dropdownChoices["azure"])
	assert.NotNil(t,
		StorageContainer.GetItem(1).(*tview.Form).GetFormItemByLabel("Azure File Share"),
		"Azure storage form not found",
	)
}

func TestStorageDropdownAws(t *testing.T) {
	/*
		When setting the dropdown to AWS, the correct
		form is rendered
	*/
	CloudPlatformDropDown.SetCurrentOption(dropdownChoices["aws"])
	assert.NotNil(t,
		StorageContainer.GetItem(1).(*tview.Form).GetFormItemByLabel("AWS File System ID"),
		"AWS storage form not found",
	)
}

func TestCertManagerDropdownDefault(t *testing.T) {
	CloudPlatformDropDown.SetCurrentOption(dropdownChoices["local"])
	assert.Equal(t, 1, subSideMenu.GetItemCount(), fmt.Sprintf("Expected 2 elements. Got %d", subSideMenu.GetItemCount()))
}
func TestCertManagerDropdownAzure(t *testing.T) {
	CloudPlatformDropDown.SetCurrentOption(dropdownChoices["azure"])
	assert.Equal(t, 2, subSideMenu.GetItemCount(), "Azure Form not rendered")
}
func TestCertManagerDropdownAws(t *testing.T) {
	CloudPlatformDropDown.SetCurrentOption(dropdownChoices["aws"])
	assert.Equal(t, 2, subSideMenu.GetItemCount(), "AWS Form not rendered")
}

func TestCertManagerEmailValidator(t *testing.T) {
	CloudPlatformDropDown.SetCurrentOption(dropdownChoices["aws"])
	emailTests := []string{
		"not email address",
		"user@email",
		"@email.com",
		"user@.com",
	}
	for _, em := range emailTests {
		awsSSLEmail.SetText(em)
		assert.Equal(t, components.ErrorBoard.GetText(false), "Email is not in a valid format")
		assert.True(t, components.CreateSecretButton.IsDisabled(), "Button is clickable")
	}
}

func TestDBPortIsInt(t *testing.T) {
	dbPort := GeneralSettingsForm.GetFormItemByLabel("Database Port").(*tview.InputField)
	dbPort.SetText("asdasd")
	assert.Equal(t, components.ErrorBoard.GetText(false), "Port should be an integer")
	assert.True(t, components.CreateSecretButton.IsDisabled(), "Button is clickable")
}

func TestDeliveryGithub(t *testing.T) {
	OutboundSettingsForm.GetFormItemByLabel("Deliver to").(*tview.DropDown).SetCurrentOption(1)
	assert.NotNil(t,
		GhSecretsContainer.GetItem(1).(*tview.Form).GetFormItemByLabel("Github Delivery Repository"),
		"GitHub form not found",
	)
}
func TestDeliveryOther(t *testing.T) {
	OutboundSettingsForm.GetFormItemByLabel("Deliver to").(*tview.DropDown).SetCurrentOption(2)
	assert.NotNil(t,
		GhSecretsContainer.GetItem(1).(*tview.Form).GetFormItemByLabel("Other Delivery Url"),
		"Other delivery form not found",
	)
}

func TestArgoForm(t *testing.T) {
	IntroForm.GetFormItemByLabel("Use ArgoCD to deploy?").(*tview.Checkbox).SetChecked(true)
	assert.NotNil(t,
		IntroContainer.GetItem(1).(*tview.Form).GetFormItemByLabel("App name"),
		"Argo form not found",
	)
}
