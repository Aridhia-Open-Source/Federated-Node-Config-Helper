package cmd

import (
	"fn-config-helper/components"
	"fn-config-helper/forms"
	"fn-config-helper/helpers"
	"os"
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func TestArgoEnabled(t *testing.T) {
	forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).SetText("host")
	forms.IntroForm.GetFormItemByLabel("Use ArgoCD to deploy?").(*tview.Checkbox).SetChecked(true)
	getValuesAndSaveYaml()

	_, err := os.Stat("values.yaml")
	assert.NotNil(t, err)
	assert.Nil(t, err, components.ErrorBoard.GetText(true))
	_, vals := helpers.ReadYAML("argo-app-deployment.yaml")
	assert.Equal(t, vals.Spec.SyncPolicy.SyncOptions, []string{"RespectIgnoreDifferences=true"})
	os.Remove("argo-app-deployment.yaml")
	forms.IntroForm.GetFormItemByLabel("Use ArgoCD to deploy?").(*tview.Checkbox).SetChecked(false)
}
