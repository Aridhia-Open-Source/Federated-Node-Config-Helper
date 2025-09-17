package pages

import (
	"fn-installer/forms"

	"github.com/rivo/tview"
)

func CreateConfigPages(app *tview.Application) *tview.Pages {
	config_pages := tview.NewPages()

	// Config pages
	config_pages.AddPage("General", forms.GeneralSettingsForm, true, true)
	config_pages.AddPage("Storage", forms.StorageSettingsForm, true, false)
	config_pages.AddPage("Outbound", forms.OutboundSettingsForm, true, false)
	config_pages.AddPage("Nginx", forms.NginxSettingsForm, true, false)
	config_pages.AddPage("CertManager", forms.CertSettingsForm, true, false)
	config_pages.AddPage("Namespaces", forms.NamespacesForm, true, false)
	config_pages.SetBorder(true).SetTitle("Federated Node configuration helper")

	return config_pages
}
