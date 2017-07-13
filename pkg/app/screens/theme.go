package screens

import tui "github.com/marcusolsson/tui-go"

func createTheme() (t *tui.Theme) {
	t = tui.NewTheme()
	t.SetStyle("list.item", tui.Style{Bg: tui.ColorDefault, Fg: tui.ColorWhite})
	t.SetStyle("list.item.selected", tui.Style{Bg: tui.ColorBlue, Fg: tui.ColorWhite})

	t.SetStyle("label.logo", tui.Style{Bg: tui.ColorDefault, Fg: tui.ColorRed})
	return
}
