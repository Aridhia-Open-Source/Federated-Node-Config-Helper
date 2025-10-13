package forms

import (
	"fn-config-helper/helpers"

	"github.com/rivo/tview"
)

var FirstUserFrom = tview.NewForm().
	AddInputField("Secret Name", "first-user", 0, nil, nil).
	AddInputField("Password Key", "passw", 0, nil, nil).
	AddInputField("User password", "", 0, nil, nil).
	AddInputField("First Name", "", 0, nil, nil).
	AddInputField("Last Name", "", 0, nil, nil).
	AddInputField("Email", "", 0, nil, helpers.EmailValidator)

func init() {
	FirstUserFrom.SetBorder(true).SetTitle("First User Setup (optional)")
}
