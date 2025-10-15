package cmd

import (
	"fn-config-helper/components"
	"fn-config-helper/forms"
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

var yamlTestFilePath = "../tests/import.values.yaml"
var yamlTestMissingFilePath = "../tests/missing.values.yaml"
var yamlTestArgoFilePath = "../tests/argo.yaml"
var yamlTestAwsFilePath = "../tests/aws.values.yaml"
var yamlTestAzureFilePath = "../tests/azure.values.yaml"

func TestBaseImport(t *testing.T) {
	defer tearDown()
	//set the path input text to the test yaml
	components.FilepathInput.SetText(yamlTestFilePath)
	importHandler()
	//clear text input
	components.FilepathInput.SetText("")
	assert.Equal(t, "db.default.svc", forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).GetText())
	// Check one of the default values are not overriden
	assert.Equal(t, "5432", forms.GeneralSettingsForm.GetFormItemByLabel("Database Port").(*tview.InputField).GetText())
}

func TestAWSImport(t *testing.T) {
	defer tearDown()
	//set the path input text to the test yaml
	components.FilepathInput.SetText(yamlTestAwsFilePath)
	importHandler()
	//clear text input
	components.FilepathInput.SetText("")
	assert.Equal(t, "fs-0cfbff2429fbe1c3a", forms.AwsStorageForm.GetFormItemByLabel("AWS File System ID").(*tview.InputField).GetText())
	assert.Equal(t, "fsap-0413cbd15dbc26d2a", forms.AwsStorageForm.GetFormItemByLabel("AWS Access Point ID").(*tview.InputField).GetText())
	assert.Equal(t, "config-aws", forms.AwsSecretsForm.GetFormItemByLabel("AWS SSL Secret Name").(*tview.InputField).GetText())
}

func TestAzureImport(t *testing.T) {
	defer tearDown()
	//set the path input text to the test yaml
	components.FilepathInput.SetText(yamlTestAzureFilePath)
	importHandler()
	//clear text input
	components.FilepathInput.SetText("")
	assert.Equal(t, "files", forms.AzureStorageForm.GetFormItemByLabel("Azure File Share").(*tview.InputField).GetText())
	assert.Equal(t, "azstorage", forms.AzureStorageSecretsForm.GetFormItemByLabel("Azure Storage Secret Name").(*tview.InputField).GetText())
	assert.Equal(t, "test-azure-qc", forms.AzureSSLSecretsForm.GetFormItemByLabel("Azure SSL ConfigMap Name").(*tview.InputField).GetText())
	assert.Equal(t, "test-azure-qc", forms.AzureSSLSecretsForm.GetFormItemByLabel("Azure SSL SP Secret").(*tview.InputField).GetText())
}

func TestMissingMandatoryFieldImport(t *testing.T) {
	/*
		Shouldn't cause panic, just leave the default empty field
		it will be then flagged on creating config
	*/
	//set the path input text to the test yaml
	components.FilepathInput.SetText(yamlTestMissingFilePath)
	importHandler()
	//clear text input
	components.FilepathInput.SetText("")
	assert.Equal(t, "", forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).GetText())
}

func TestArgoImport(t *testing.T) {
	//set the path input text to the test yaml
	components.FilepathInput.SetText(yamlTestArgoFilePath)
	importHandler()
	//clear text input
	components.FilepathInput.SetText("")
	assert.Equal(t, "db.default.svc", forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).GetText())
	assert.Equal(t, "5432", forms.GeneralSettingsForm.GetFormItemByLabel("Database Port").(*tview.InputField).GetText())

}
