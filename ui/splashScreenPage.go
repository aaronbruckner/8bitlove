package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const TITLE = `
 ░▒▓██████▓▒░       ░▒▓███████▓▒░░▒▓█▓▒░▒▓████████▓▒░      ░▒▓█▓▒░      ░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓████████▓▒░ 
░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░  ░▒▓█▓▒░          ░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░        
░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░  ░▒▓█▓▒░          ░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒▒▓█▓▒░░▒▓█▓▒░        
 ░▒▓██████▓▒░       ░▒▓███████▓▒░░▒▓█▓▒░  ░▒▓█▓▒░          ░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒▒▓█▓▒░░▒▓██████▓▒░   
░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░  ░▒▓█▓▒░          ░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▓█▓▒░ ░▒▓█▓▒░        
░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░  ░▒▓█▓▒░          ░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░ ░▒▓█▓▓█▓▒░ ░▒▓█▓▒░        
 ░▒▓██████▓▒░       ░▒▓███████▓▒░░▒▓█▓▒░  ░▒▓█▓▒░          ░▒▓████████▓▒░▒▓██████▓▒░   ░▒▓██▓▒░  ░▒▓████████▓▒░
`

func newSplashScreenPage(onContinue func()) *tview.Flex {
	var titleTextView = tview.NewTextView().SetText(TITLE).SetTextAlign(tview.AlignCenter)
	var continueTextView = tview.NewTextView().SetText("Press any key to continue...").SetTextAlign(tview.AlignCenter)

	var page = tview.NewFlex()
	page.SetDirection(tview.FlexRow)
	page.AddItem(titleTextView, 0, 1, true)
	page.AddItem(continueTextView, 0, 1, true)

	page.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		onContinue()
		return nil
	})

	return page
}
