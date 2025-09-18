package pages

import (
	"fn-installer/components"
	"fn-installer/forms"
	"fn-installer/state"

	"github.com/gdamore/tcell/v2"
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
		app.SetFocus(forms.GHContainer)
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
	cloudPlatformDropDown.AddOption("None (local)", func() { state.State.CloudPlatform = "Local" })
	cloudPlatformDropDown.AddOption("AWS", func() { state.State.CloudPlatform = "AWS" })
	cloudPlatformDropDown.AddOption("Azure", func() { state.State.CloudPlatform = "Azure" })
	cloudPlatformDropDown.SetSelectedFunc(forms.HideFields)
	cloudPlatformDropDown.SetCurrentOption(0)
	components.ErrorBoard.SetTextColor(tcell.ColorRed).SetBorder(true).SetBorderColor(tcell.ColorRed)

	NamespaceDeployment.SetLabel("Deployment Namespace")
	NamespaceDeployment.SetText("default")

	introView := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(NamespaceDeployment, 0, 2, true).
		AddItem(cloudPlatformDropDown, 0, 1, true)

	page.AddPage("Select platform", introView, true, true)
	page.AddPage("Secrets", forms.SecretContainer, true, false)
	page.AddPage("General", forms.GeneralSettingsForm, true, false)
	page.AddPage("Storage", forms.StorageSettingsForm, true, false)
	page.AddPage("Outbound", forms.GHContainer, true, false)
	page.AddPage("Nginx", forms.NginxSettingsForm, true, false)
	page.AddPage("CertManager", forms.CertSettingsForm, true, false)
	page.AddPage("Namespaces", forms.NamespacesForm, true, false)

	return page, mainSideMenu
}
