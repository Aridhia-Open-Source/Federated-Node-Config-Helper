package components

import (
	"fmt"
	"fn-config-helper/state"

	"github.com/rivo/tview"
)

var Footer = tview.NewTextView()

func init() {
	Footer.SetBorder(true)
	Footer.SetText(fmt.Sprintf("Deploy command:\n\nhelm install federatednode -n %s -f values.yaml", state.State.Namespace))
}
