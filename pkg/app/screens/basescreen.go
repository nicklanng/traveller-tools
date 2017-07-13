package screens

import tui "github.com/marcusolsson/tui-go"

type baseScreen struct {
	content *tui.Box
	status  *tui.StatusBar
}

func newBaseScreen() baseScreen {
	s := baseScreen{}
	s.status = tui.NewStatusBar("Ready.")
	return s
}

func (s *baseScreen) render() {
	theme := createTheme()

	root := tui.NewVBox(
		s.content,
		s.status,
	)

	ui := tui.New(root)
	ui.SetTheme(theme)
	ui.SetKeybinding("Ctrl+c", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
