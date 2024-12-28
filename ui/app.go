package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	PAGE_WELCOME   string = "welcomePage"
	PAGE_MAIN_MENU string = "mainMenuPage"
)

func StartApp() {

	var pages = tview.NewPages()
	pages.AddPage(PAGE_WELCOME, newSplashScreenPage(), true, false)
	pages.AddPage(PAGE_MAIN_MENU, newMainPage(), true, false)
	pages.SwitchToPage(PAGE_WELCOME)
	pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		pages.SwitchToPage(PAGE_MAIN_MENU)
		return event
	})
	tview.NewApplication().SetRoot(pages, true).Run()

}
