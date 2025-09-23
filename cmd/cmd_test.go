package cmd

import (
	"fmt"
	"fn-config-helper/components"
	"fn-config-helper/forms"
	"fn-config-helper/helpers"
	"os"
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	exitcode := m.Run()
	err := os.Remove("values.yaml")
	if err != nil {
		fmt.Println("No files to delete. Ignoring")
	}
	os.Exit(exitcode)
}

func TestSaveYaml(t *testing.T) {
	/*
		Test to make sure the file is created successful with
		a minimal config provided
	*/
	st := helpers.Config{
		LocalDevelopment: false,
	}
	st.CreateYaml()
	_, err := os.Stat("values.yaml")
	assert.Nil(t, err)
}

func TestProperFieldConverted(t *testing.T) {
	/*
		Simple test with all defaults to make sure
		all is converted smoothly
	*/
	getValuesAndSaveYaml()
	_, err := os.Stat("values.yaml")
	assert.Nil(t, err)
	os.Remove("values.yaml")
}

func TestWrongValues(t *testing.T) {
	/*
		Simple test to make sure a file is not created
		if an error occurs
	*/
	forms.GeneralSettingsForm.GetFormItemByLabel("Database Port").(*tview.InputField).SetText("asdasdas")
	getValuesAndSaveYaml()
	assert.Equal(t, components.ErrorBoard.GetText(false), "strconv.Atoi: parsing \"asdasdas\": invalid syntax")
	_, err := os.Stat("values.yaml")
	assert.NotNil(t, err)
}

func TestArgoEnabled(t *testing.T) {
	forms.IntroForm.GetFormItemByLabel("Use ArgoCD to deploy?").(*tview.Checkbox).SetChecked(true)
	getValuesAndSaveYaml()
	_, err := os.Stat("values.yaml")
	assert.NotNil(t, err)
	_, err = os.Stat("argo-app-deployment.yaml")
	assert.Nil(t, err)
}
