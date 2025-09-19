package forms

import (
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
	if StorageContainer.GetItem(1).(*tview.Form).GetFormItemByLabel("Local Path") == nil {
		t.Fatalf("Azure storage form not found")
	}
}
func TestStorageDropdownAzure(t *testing.T) {
	/*
		When setting the dropdown to azure, the correct
		form is rendered
	*/
	CloudPlatformDropDown.SetCurrentOption(dropdownChoices["azure"])
	if StorageContainer.GetItem(1).(*tview.Form).GetFormItemByLabel("Azure File Share") == nil {
		t.Fatalf("Azure storage form not found")
	}
}

func TestStorageDropdownAws(t *testing.T) {
	/*
		When setting the dropdown to AWS, the correct
		form is rendered
	*/
	CloudPlatformDropDown.SetCurrentOption(dropdownChoices["aws"])
	if StorageContainer.GetItem(1).(*tview.Form).GetFormItemByLabel("AWS File System ID") == nil {
		t.Fatalf("AWS storage form not found")
	}
}

func TestCertManagerDropdownDefault(t *testing.T) {
	CloudPlatformDropDown.SetCurrentOption(dropdownChoices["local"])
	if SecretContainer.GetItemCount() > 2 {
		t.Fatalf("Expected 2 elements. Got %d", SecretContainer.GetItemCount())
	}
}
func TestCertManagerDropdownAzure(t *testing.T) {
	CloudPlatformDropDown.SetCurrentOption(dropdownChoices["azure"])
	if SecretContainer.GetItemCount() != 3 {
		t.Fatalf("Azure Form not rendered")
	}
}
func TestCertManagerDropdownAws(t *testing.T) {
	CloudPlatformDropDown.SetCurrentOption(dropdownChoices["aws"])
	if SecretContainer.GetItemCount() != 3 {
		t.Fatalf("AWS Form not rendered")
	}
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
	if OutboundContainer.GetItem(2).(*tview.Form).GetFormItemByLabel("Github Delivery Repository") == nil {
		t.Fatalf("GitHub form not found")
	}
}
func TestDeliveryOther(t *testing.T) {
	OutboundSettingsForm.GetFormItemByLabel("Deliver to").(*tview.DropDown).SetCurrentOption(2)
	if OutboundContainer.GetItem(2).(*tview.Form).GetFormItemByLabel("Other Delivery Url") == nil {
		t.Fatalf("Other delivery form not found")
	}
}
