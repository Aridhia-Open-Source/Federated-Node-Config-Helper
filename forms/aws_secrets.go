package forms

import (
	"fn-config-helper/helpers"

	"github.com/rivo/tview"
)

var awsSSLSecretName = tview.NewInputField().
	SetLabel("AWS SSL Secret Name")
var awsSSLEmail = tview.NewInputField().
	SetLabel("AWS SSL Email").
	SetChangedFunc(helpers.EmailValidator)
var awsSSLRegion = tview.NewInputField().
	SetLabel("AWS SSL Account ID")
var awsSSLAccountId = tview.NewInputField().
	SetLabel("AWS SSL Account ID")
var awsSSLRoleName = tview.NewInputField().
	SetLabel("AWS SSL Role Name")

var AwsSecretsForm = tview.NewForm().AddFormItem(awsSSLSecretName).
	AddFormItem(awsSSLEmail).
	AddFormItem(awsSSLRegion).
	AddFormItem(awsSSLAccountId).
	AddFormItem(awsSSLRoleName)
