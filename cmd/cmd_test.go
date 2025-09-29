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
}
func TestSaveYaml(t *testing.T) {
	/*
		Test to make sure the file is created successful with
		a minimal config provided
	*/
	defer tearDown()
	st := helpers.Config{
		LocalDevelopment: false,
	}
	st.CreateYaml("values.yaml")
	_, err := os.Stat("values.yaml")
	assert.Nil(t, err)
}

func TestProperFieldConverted(t *testing.T) {
	/*
		Simple test with all defaults to make sure
		all is converted smoothly
	*/
	defer tearDown()
	getValuesAndSaveYaml()
	_, err := os.Stat("values.yaml")
	assert.Nil(t, err)
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
