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

func tearDown() {
	os.Remove("values.yaml")
	forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).SetText("")
}

func TestSaveYaml(t *testing.T) {
	/*
		Test to make sure the file is created successful with
		a minimal config provided
	*/
	defer tearDown()
	forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).SetText("host")
	st := helpers.Config{
		LocalDevelopment: false,
	}
	st.CreateYaml("values.yaml")
	_, err := os.Stat("values.yaml")
	assert.Nil(t, err)
}

func TestMissingRequiredField(t *testing.T) {
	/*
		Simple test with missing required field does print an error
	*/
	defer tearDown()
	getValuesAndSaveYaml()
	_, err := os.Stat("values.yaml")
	assert.NotNil(t, err)
	assert.Equal(t, "Database Hostname should not be empty", components.ErrorBoard.GetText(true))
}

func TestProperFieldConverted(t *testing.T) {
	/*
		Simple test with all defaults to make sure
		all is converted smoothly
	*/
	defer tearDown()
	forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).SetText("host")
	getValuesAndSaveYaml()
	vals, _ := helpers.ReadYAML("values.yaml")
	assert.Equal(t, vals.Database.Hostname, "host")
}

func TestWrongValues(t *testing.T) {
	/*
		Simple test to make sure a file is not created
		if an error occurs
	*/
	forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).SetText("host")
	forms.GeneralSettingsForm.GetFormItemByLabel("Database Port").(*tview.InputField).SetText("asdasdas")
	getValuesAndSaveYaml()
	assert.Equal(t, components.ErrorBoard.GetText(false), "strconv.Atoi: parsing \"asdasdas\": invalid syntax")
	_, err := os.Stat("values.yaml")
	assert.NotNil(t, err)
	forms.GeneralSettingsForm.GetFormItemByLabel("Database Port").(*tview.InputField).SetText("5432")
}

func TestFirstUserIgnored(t *testing.T) {
	defer tearDown()
	forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).SetText("host")
	getValuesAndSaveYaml()

	vals, _ := helpers.ReadYAML("values.yaml")
	assert.Nil(t, vals.FirstUserSecret)
}

func TestFirstUserFilled(t *testing.T) {
	defer tearDown()
	forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).SetText("host")
	forms.FirstUserFrom.GetFormItemByLabel("User password").(*tview.InputField).SetText("asdasdas")
	getValuesAndSaveYaml()

	vals, _ := helpers.ReadYAML("values.yaml")
	assert.NotNil(t, vals.FirstUserSecret)
}
