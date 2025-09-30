package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type CreateConfigButton struct {
	*tview.Button
}

func NewCreateConfigButton(label string) *CreateConfigButton {
	baseButton := tview.NewButton(label)

	custom := &CreateConfigButton{baseButton}
	style := tcell.Style{}
	style = style.Background(tcell.ColorDarkGreen)
	style = style.Foreground(tcell.ColorDarkGreen)

	custom.SetStyle(style)

	custom.SetBackgroundColorActivated(tcell.ColorDarkGreen)
	custom.SetLabelColorActivated(tcell.ColorWhite)
	custom.SetLabelColor(tcell.ColorWhite)
	return custom
}

func NewCreateQuitButton(label string) *CreateConfigButton {
	baseButton := tview.NewButton(label)

	custom := &CreateConfigButton{baseButton}
	style := tcell.Style{}
	style = style.Background(tcell.ColorDarkRed)
	style = style.Foreground(tcell.ColorDarkRed)

	custom.SetStyle(style)

	custom.SetBackgroundColorActivated(tcell.ColorDarkRed)
	custom.SetLabelColorActivated(tcell.ColorWhite)
	custom.SetLabelColor(tcell.ColorWhite)
	return custom
}
