package forms

import (
	"fn-config-helper/helpers"

	"github.com/rivo/tview"
)

var GeneralSettingsForm = tview.NewForm().
	AddCheckbox("Is development deployment", false, nil).
	AddCheckbox("Use Task Result Review", false, nil).
	AddCheckbox("Enable Smoketests", false, nil).
	AddInputField("Database Host", "", 20, nil, nil).
	AddInputField("Database User", "admin", 20, nil, nil).
	AddInputField("Database Name", "fndb", 20, nil, nil).
	AddInputField("Database Port", "5432", 20, nil, helpers.PortValidator).
	AddInputField("Keycloak Replicas", "2", 20, nil, helpers.ReplicasValidator).
	AddInputField("Cleanup Time", "3", 20, nil, helpers.CleanupDaysValidator)
