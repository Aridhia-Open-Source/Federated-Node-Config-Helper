package forms

import "github.com/rivo/tview"

var NginxSettingsForm = tview.NewForm().
	AddCheckbox("Use nginx", false, nil).
	AddInputField("Host URL", "", 20, nil, nil).
	AddInputField("Namespace", "ingress-nginx", 20, nil, nil).
	AddInputField("Ingress Class", "fn-nginx", 20, nil, nil).
	AddInputField("Default SSL cert", "default/tls", 20, nil, nil).
	AddCheckbox("Allow Snippet Annotations", false, nil)
