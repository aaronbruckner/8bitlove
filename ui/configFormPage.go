package ui

import (
	"github.com/aaronbruckner/8bitlove/config"
	"github.com/rivo/tview"
)

func newConfigFormPage(onSubmit func(authConfig *config.AuthConfig)) *tview.Form {
	var page = tview.NewForm()
	var config = config.AuthConfig{Aws_access_key: "", Aws_secret_key: ""}
	page.AddTextView("Enter credentials provided by Aaron. Values are saved locally.", "", 0, 1, false, false)
	page.AddPasswordField("AWS Access Key", "", 20, '*', func(value string) { config.Aws_access_key = value })
	page.AddPasswordField("AWS Secret Key", "", 20, '*', func(value string) { config.Aws_secret_key = value })
	page.AddButton("Save & Continue", func() {
		onSubmit(&config)
	})

	page.SetTitle("Auth Config")
	page.SetBorder(true)

	return page
}
