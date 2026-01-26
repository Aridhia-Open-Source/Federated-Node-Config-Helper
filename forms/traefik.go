package forms

import "github.com/rivo/tview"

var TraefikSettingsForm = tview.NewForm().
	AddCheckbox("Use Traefik", false, nil).
	AddInputField("Host URL", "", 20, nil, nil).
	AddInputField("Gateway Name", "traefik-fn-gateway", 20, nil, nil).
	AddInputField("Gateway Class", "traefik-fn-gateway-class", 20, nil, nil)
