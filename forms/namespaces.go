package forms

import "github.com/rivo/tview"

var NamespacesForm = tview.NewForm().
	AddInputField("Keycloak", "keycloak", 20, nil, nil).
	AddInputField("Tasks", "tasks", 20, nil, nil).
	AddInputField("Controller", "fn-controller", 20, nil, nil)
