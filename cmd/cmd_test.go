package cmd

import (
	"fn-config-helper/components"
	"fn-config-helper/forms"
	"fn-config-helper/helpers"
	"fn-config-helper/state"
	"os"
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func tearDown() {
	os.Remove("values.yaml")
	components.FilepathInput.SetText("")
	// Reset all changed fields
	forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).SetText("")
	forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).SetText("")
	forms.AzureStorageForm.GetFormItemByLabel("Azure File Share").(*tview.InputField).SetText("")
	forms.AzureStorageSecretsForm.GetFormItemByLabel("Azure Storage Secret Name").(*tview.InputField).SetText("")
	forms.AzureSSLSecretsForm.GetFormItemByLabel("Azure SSL ConfigMap Name").(*tview.InputField).SetText("")
	forms.AzureSSLSecretsForm.GetFormItemByLabel("Azure SSL SP Secret").(*tview.InputField).SetText("")
	forms.AwsStorageForm.GetFormItemByLabel("AWS File System ID").(*tview.InputField).SetText("")
	forms.AwsStorageForm.GetFormItemByLabel("AWS Access Point ID").(*tview.InputField).SetText("")
	forms.AwsSecretsForm.GetFormItemByLabel("AWS SSL Secret Name").(*tview.InputField).SetText("")
}

func TestSaveYaml(t *testing.T) {
	/*
		Test to make sure the file is created successful with
		a minimal config provided
	*/
	defer tearDown()
	forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).SetText("host")
	lDev := false
	st := helpers.Config{
		LocalDevelopment: &lDev,
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
	state.State.CloudPlatform = "Local"
	forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).SetText("host")
	getValuesAndSaveYaml()
	vals, _ := helpers.ReadYAML("values.yaml")
	assert.Equal(t, vals.Database.Host, "host")
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

func TestAWSConfig(t *testing.T) {
	defer tearDown()
	state.State.CloudPlatform = "AWS"
	forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).SetText("host")
	forms.AwsStorageForm.GetFormItemByLabel("AWS File System ID").(*tview.InputField).SetText("something")
	forms.AwsStorageForm.GetFormItemByLabel("AWS Access Point ID").(*tview.InputField).SetText("ap-123123123")
	forms.AwsSecretsForm.GetFormItemByLabel("AWS SSL Secret Name").(*tview.InputField).SetText("secret")

	getValuesAndSaveYaml()

	vals, _ := helpers.ReadYAML("values.yaml")
	assert.True(t, vals.OnEks)
	assert.False(t, vals.OnAks)
	assert.Equal(t, "ap-123123123", vals.Storage.Aws.AccessPointId)
	assert.Equal(t, "something", vals.Storage.Aws.FileSystemId)
	assert.Equal(t, "secret", vals.Certs.AWS)
}

func TestAzureConfig(t *testing.T) {
	defer tearDown()
	state.State.CloudPlatform = "Azure"

	forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).SetText("host")
	forms.AzureStorageForm.GetFormItemByLabel("Azure File Share").(*tview.InputField).SetText("files")
	forms.AzureStorageSecretsForm.GetFormItemByLabel("Azure Storage Secret Name").(*tview.InputField).SetText("storage-sec")
	forms.AzureSSLSecretsForm.GetFormItemByLabel("Azure SSL ConfigMap Name").(*tview.InputField).SetText("ssl-cm")
	forms.AzureSSLSecretsForm.GetFormItemByLabel("Azure SSL SP Secret").(*tview.InputField).SetText("ssl-sec")

	getValuesAndSaveYaml()

	vals, _ := helpers.ReadYAML("values.yaml")
	assert.True(t, vals.OnAks)
	assert.False(t, vals.OnEks)
	assert.Equal(t, "storage-sec", vals.Storage.Azure.SecretName)
	assert.Equal(t, "files", vals.Storage.Azure.ShareName)
	assert.Equal(t, "ssl-sec", vals.Certs.Azure.Secret)
	assert.Equal(t, "ssl-cm", vals.Certs.Azure.Configmap)
}

func TestAzureMissingRequiredConfig(t *testing.T) {
	defer tearDown()
	state.State.CloudPlatform = "Azure"

	forms.GeneralSettingsForm.GetFormItemByLabel("Database Host").(*tview.InputField).SetText("host")
	forms.AzureStorageSecretsForm.GetFormItemByLabel("Azure Storage Secret Name").(*tview.InputField).SetText("storage-sec")
	forms.AzureSSLSecretsForm.GetFormItemByLabel("Azure SSL ConfigMap Name").(*tview.InputField).SetText("ssl-cm")
	forms.AzureSSLSecretsForm.GetFormItemByLabel("Azure SSL SP Secret").(*tview.InputField).SetText("ssl-sec")

	getValuesAndSaveYaml()
	assert.Equal(t, components.ErrorBoard.GetText(false), "Azure File Share should not be empty")
	_, err := os.Stat("values.yaml")
	assert.NotNil(t, err)
}
