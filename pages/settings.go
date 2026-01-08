package pages

import (
	"fn-config-helper/components"
	"fn-config-helper/forms"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var page = tview.NewPages()

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
		app.SetFocus(forms.StorageContainer)
	})
	mainSideMenu.AddItem("Outbound mode", "", '4', func() {
		page.SwitchToPage("Outbound")
		app.SetFocus(forms.OutboundContainer)
	})
	mainSideMenu.AddItem("Certificate Manager", "", '5', func() {
		page.SwitchToPage("CertManager")
		app.SetFocus(forms.CertSettingsForm)
	})
	mainSideMenu.AddItem("Traefik", "", '6', func() {
		page.SwitchToPage("Traefik")
		app.SetFocus(forms.TraefikSettingsForm)
	})
	mainSideMenu.AddItem("Namespaces", "", '7', func() {
		page.SwitchToPage("Namespaces")
		app.SetFocus(forms.NamespacesForm)
	})
	mainSideMenu.SetBorder(true).SetTitle("Categories")

	components.ErrorBoard.SetTextColor(tcell.ColorRed).SetBorder(true).SetBorderColor(tcell.ColorRed)

	introView := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(forms.IntroContainer, 0, 2, true)

	page.AddPage("Select platform", introView, true, true)
	page.AddPage("Secrets", forms.SecretContainer, true, false)
	page.AddPage("General", forms.GeneralSettingsForm, true, false)
	page.AddPage("Storage", forms.StorageContainer, true, false)
	page.AddPage("Outbound", forms.OutboundContainer, true, false)
	page.AddPage("Traefik", forms.TraefikSettingsForm, true, false)
	page.AddPage("CertManager", forms.CertSettingsForm, true, false)
	page.AddPage("Namespaces", forms.NamespacesForm, true, false)

	page.SetBorder(true)

	return page, mainSideMenu
}
