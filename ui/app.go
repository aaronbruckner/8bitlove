package ui

import (
	"github.com/rivo/tview"
)

const (
	PAGE_WELCOME   string = "welcomePage"
	PAGE_MAIN_MENU string = "mainMenuPage"
)

func StartApp() {

	var pages = tview.NewPages()

	var onSplashScreenContinue = func() {
		pages.SwitchToPage(PAGE_MAIN_MENU)
	}
	pages.AddPage(PAGE_WELCOME, newSplashScreenPage(onSplashScreenContinue), true, false)
	pages.AddPage(PAGE_MAIN_MENU, newMainPage(), true, false)
	pages.SwitchToPage(PAGE_WELCOME)
	tview.NewApplication().SetRoot(pages, true).Run()

}
