package components

import (
	"github.com/rivo/tview"
)

var FileChoice = tview.NewModal()
var warningText = tview.NewTextView().SetText("Warning: Importing a yaml will not populate or look into existing secrets/configmaps. Once imported those fields will be left empty")

func init() {
	FilepathInput.SetLabel("File path")
	CancelModalButton.SetBorder(true)
	ConfirmButton.SetBorder(true)
}

var ConfirmButton = tview.NewButton("Load")
var CancelModalButton = tview.NewButton("Cancel")

var FilepathInput = tview.NewInputField()

var modal = tview.NewFlex().SetDirection(tview.FlexRow).
	AddItem(warningText, 0, 2, true).
	AddItem(FilepathInput, 0, 6, true).
	AddItem(ConfirmButton, 0, 1, false).
	AddItem(CancelModalButton, 0, 1, false)

var ModalPage = tview.NewPages().
	AddPage("modal", modal, true, true)
