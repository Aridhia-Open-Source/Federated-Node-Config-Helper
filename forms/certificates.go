package forms

import "github.com/rivo/tview"

var CertSettingsForm = tview.NewForm().
	AddCheckbox("Use cert manager", false, nil).
	AddInputField("Namespace", "fn-certmanager", 20, nil, nil).
	AddCheckbox("Install CRDs", true, nil)
