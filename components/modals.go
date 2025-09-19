package components

import (
	"github.com/rivo/tview"
)

var FailureModal = tview.NewModal().AddButtons([]string{"OK"})
var SuccessModal = tview.NewModal().AddButtons([]string{"OK"})

func init() {
	FailureModal.SetTitle("Error")
	SuccessModal.SetTitle("Success")
}
