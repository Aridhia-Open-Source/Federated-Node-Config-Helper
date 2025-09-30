package cmd

import (
	"fmt"
	"fn-config-helper/components"
	"fn-config-helper/forms"
	"log"
	"os"
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func TestArgoEnabled(t *testing.T) {
	forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).SetText("host")
	forms.IntroForm.GetFormItemByLabel("Use ArgoCD to deploy?").(*tview.Checkbox).SetChecked(true)
	getValuesAndSaveYaml()
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		fmt.Println(e.Name())
	}
	_, err = os.Stat("values.yaml")
	assert.NotNil(t, err)
	_, err = os.Stat("argo-app-deployment.yaml")
	assert.Nil(t, err, components.ErrorBoard.GetText(true))
	os.Remove("argo-app-deployment.yaml")
	forms.IntroForm.GetFormItemByLabel("Use ArgoCD to deploy?").(*tview.Checkbox).SetChecked(false)
}
