package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var CreateSecretButton = tview.NewButton("Create Secrets").
	SetDisabledStyle(tcell.StyleDefault.Background(tcell.ColorGrey))
