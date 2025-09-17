package pages

import (
	"fn-installer/forms"
	"fn-installer/helpers"

	"github.com/rivo/tview"
)

var page = tview.NewPages()
var NamespaceDeployment = tview.NewInputField()

func CreateMainPage(app *tview.Application) (*tview.Pages, *tview.List) {
	mainSideMenu := tview.NewList().ShowSecondaryText(false)

	// side menu entries
	mainSideMenu.AddItem("Select platform", "", '0', func() {
		page.SwitchToPage("Select platform")
	})
	mainSideMenu.AddItem("Secrets", "", '1', func() {
		page.SwitchToPage("Secrets")
		app.SetFocus(forms.SecretContainer)
	})
	mainSideMenu.AddItem("General", "", '2', func() {
		page.SwitchToPage("General")
		app.SetFocus(forms.GeneralSettingsForm)
	})
	mainSideMenu.AddItem("Storage", "", '3', func() {
		page.SwitchToPage("Storage")
		app.SetFocus(forms.StorageSettingsForm)
	})
	mainSideMenu.AddItem("Outbound mode", "", '4', func() {
		page.SwitchToPage("Outbound")
		app.SetFocus(forms.OutboundSettingsForm)
	})
	mainSideMenu.AddItem("Certificate Manager", "", '5', func() {
		page.SwitchToPage("CertManager")
		app.SetFocus(forms.CertSettingsForm)
	})
	mainSideMenu.AddItem("Nginx", "", '6', func() {
		page.SwitchToPage("Nginx")
		app.SetFocus(forms.NginxSettingsForm)
	})
	mainSideMenu.AddItem("Namespaces", "", '7', func() {
		page.SwitchToPage("Namespaces")
		app.SetFocus(forms.NamespacesForm)
	})
	mainSideMenu.SetBorder(true).SetTitle("Categories")

	cloudPlatformDropDown := tview.NewDropDown()
	cloudPlatformDropDown.SetLabel("Choose the cloud provider (if any)")
	cloudPlatformDropDown.AddOption("None (local)", func() { helpers.State.CloudPlatform = "Local" })
	cloudPlatformDropDown.AddOption("AWS", func() { helpers.State.CloudPlatform = "AWS" })
	cloudPlatformDropDown.AddOption("Azure", func() { helpers.State.CloudPlatform = "Azure" })
	cloudPlatformDropDown.SetSelectedFunc(forms.HideFields)

	NamespaceDeployment.SetLabel("Deployment Namespace")

	introView := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(NamespaceDeployment, 0, 1, true).
		AddItem(cloudPlatformDropDown, 0, 1, true)

	page.AddPage("Select platform", introView, true, true)
	page.AddPage("Secrets", forms.SecretContainer, true, false)
	page.AddPage("General", forms.GeneralSettingsForm, true, false)
	page.AddPage("Storage", forms.StorageSettingsForm, true, false)
	page.AddPage("Outbound", forms.OutboundSettingsForm, true, false)
	page.AddPage("Nginx", forms.NginxSettingsForm, true, false)
	page.AddPage("CertManager", forms.CertSettingsForm, true, false)
	page.AddPage("Namespaces", forms.NamespacesForm, true, false)

	return page, mainSideMenu
}
