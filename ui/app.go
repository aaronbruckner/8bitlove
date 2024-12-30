package ui

import (
	"github.com/aaronbruckner/8bitlove/config"
	"github.com/rivo/tview"
)

const (
	PAGE_WELCOME          string = "welcomePage"
	PAGE_MAIN_MENU        string = "mainMenuPage"
	PAGE_CONFIG_FORM_PAGE string = "configFormPage"
)

func StartApp() {

	var pages = tview.NewPages()
	var authConfig, authConfigErr = config.LoadLocalAuthConfigIfPresent()

	var onSplashScreenContinue = func() {
		if authConfigErr != nil {
			pages.SwitchToPage(PAGE_CONFIG_FORM_PAGE)
		} else {
			pages.SwitchToPage(PAGE_MAIN_MENU)
		}
	}

	var onAuthConfigSubmit = func(c config.AuthConfig) {
		authConfig = c
		config.SaveLocalAuthConfig(c)
		pages.SwitchToPage(PAGE_MAIN_MENU)
		if authConfig.Aws_access_key != "" {

		}
	}

	pages.AddPage(PAGE_WELCOME, newSplashScreenPage(onSplashScreenContinue), true, false)
	pages.AddPage(PAGE_MAIN_MENU, newMainPage(), true, false)
	pages.AddPage(PAGE_CONFIG_FORM_PAGE, newConfigFormPage(onAuthConfigSubmit), true, false)
	pages.SwitchToPage(PAGE_WELCOME)
	err := tview.NewApplication().SetRoot(pages, true).EnableMouse(true).Run()
	if err != nil {
		panic(err)
	}

}
