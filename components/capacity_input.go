package components

import (
	"regexp"

	"github.com/rivo/tview"
)

type CapacityInput struct {
	*tview.InputField
	ValidationErrorMessage string
}

func (ti *CapacityInput) IsValid() bool {
	matched, _ := regexp.MatchString("\\d+(T|G|M|K)i", ti.GetText())
	return matched
}

func NewCapacityInput() *CapacityInput {
	newFieldInput := tview.NewInputField()

	custom := &CapacityInput{newFieldInput, "Capacity is not in a correct format. Try a number followed by: Ti, Gi, Mi, Ki"}
	custom.SetLabel("Capacity")
	custom.SetChangedFunc(func(text string) {
		if matched, _ := regexp.MatchString("[1-9]\\d*(T|G|M|K)i", text); matched == false {
			ErrorBoard.SetText(custom.ValidationErrorMessage)
		} else {
			ErrorBoard.SetText("")
		}
	})
	custom.SetText("1Gi")
	return custom
}
