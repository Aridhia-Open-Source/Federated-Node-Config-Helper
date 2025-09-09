package forms

import "github.com/rivo/tview"

var OutboundSettingsForm = tview.NewForm().
	AddCheckbox("Outbound mode", true, nil).
	AddDropDown("Deliver to", []string{"github", "other"}, 0, nil).
	AddInputField("IDP Secret Name", "", 20, nil, nil).
	AddInputField("IDP Secret Key", "", 20, nil, nil).
	AddInputField("IDP Client ID Key", "", 20, nil, nil)
