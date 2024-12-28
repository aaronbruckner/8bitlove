package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func newMainPage() *tview.TextView {
	var page = tview.NewTextView().SetText("Hello world Main Menu").SetTextAlign(tview.AlignCenter)

	page.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		page.SetText("Main menu changed!")
		return nil
	})

	return page
}
