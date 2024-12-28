package ui

import (
	"github.com/rivo/tview"
)

func newMainPage() *tview.TextView {
	return tview.NewTextView().SetText("Hello world Main Menu").SetTextAlign(tview.AlignCenter)
}
