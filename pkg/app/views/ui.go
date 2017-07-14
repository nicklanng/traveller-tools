package views

import (
	tui "github.com/marcusolsson/tui-go"
)

var (
	ui     tui.UI
	status *tui.StatusBar
)

func InitializeUI() {
	status = tui.NewStatusBar("Ready.")

	root := tui.NewVBox(tui.NewHBox(), status)
	ui = tui.New(root)
	ui.SetKeybinding("Ctrl+c", func() { ui.Quit() })

	theme := tui.NewTheme()
	theme.SetStyle("list.item", tui.Style{Bg: tui.ColorDefault, Fg: tui.ColorWhite})
	theme.SetStyle("list.item.selected", tui.Style{Bg: tui.ColorBlue, Fg: tui.ColorWhite})
	theme.SetStyle("label.logo", tui.Style{Bg: tui.ColorDefault, Fg: tui.ColorRed})
	theme.SetStyle("label.value", tui.Style{Bg: tui.ColorDefault, Fg: tui.ColorWhite})
	ui.SetTheme(theme)

	LoadScreen(NewMainMenu())

	if err := ui.Run(); err != nil {
		panic(err)
	}
}

func LoadScreen(w tui.Widget) {
	root := tui.NewVBox(w, status)
	ui.SetWidget(root)
}
