package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ConfigButton struct {
	*tview.Button
}

func NewConfigButton(label string, color tcell.Color) *ConfigButton {
	baseButton := tview.NewButton(label)

	custom := &ConfigButton{baseButton}
	style := tcell.Style{}
	style = style.Background(color)
	style = style.Foreground(color)

	custom.SetStyle(style)

	custom.SetBackgroundColorActivated(color)
	custom.SetLabelColorActivated(tcell.ColorWhite)
	custom.SetLabelColor(tcell.ColorWhite)
	return custom
}

func NewCreateConfigButton(label string) *ConfigButton {
	return NewConfigButton(label, tcell.ColorDarkGreen)
}

func NewCreateQuitButton(label string) *ConfigButton {
	return NewConfigButton(label, tcell.ColorDarkRed)
}
