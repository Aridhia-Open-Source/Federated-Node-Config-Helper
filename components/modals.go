package components

import (
	"github.com/rivo/tview"
)

var FailureModal = tview.NewModal()
var SuccessModal = tview.NewModal()
var ModalContainer = tview.NewFlex().
	AddItem(FailureModal, 0, 1, true)

func init() {
	FailureModal.
		AddButtons([]string{"OK"}).
		SetTitle("Error")
	SuccessModal.
		AddButtons([]string{"OK"}).
		SetTitle("Success")
}
