package helpers

import (
	"fn-config-helper/components"
	"regexp"
)

func EmailValidator(email string) {
	matched, _ := regexp.MatchString(".+@.+\\..+", email)
	if !matched {
		components.ErrorBoard.SetText("Email is not in a valid format")
		components.CreateSecretButton.SetDisabled(true)
	} else {
		components.ErrorBoard.SetText("")
		components.CreateSecretButton.SetDisabled(false)
	}
}

func PortValidator(str string) {
	intValidator(str, "Port should be an integer")
}

func ReplicasValidator(str string) {
	intValidator(str, "Replicas should be an integer")
}

func CleanupDaysValidator(str string) {
	intValidator(str, "Cleanup Time should be an integer")
}

func intValidator(str string, errorMessage string) {
	matched, _ := regexp.MatchString("\\d+", str)
	if !matched {
		components.ErrorBoard.SetText(errorMessage)
		components.CreateSecretButton.SetDisabled(true)
	} else {
		components.ErrorBoard.SetText("")
		components.CreateSecretButton.SetDisabled(false)
	}
}
