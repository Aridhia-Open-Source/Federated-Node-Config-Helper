package helpers

import (
	"fn-installer/components"
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
