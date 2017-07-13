package screens

import (
	tui "github.com/marcusolsson/tui-go"
)

var logo = `  _____                    _ _
 |_   _| __ __ ___   _____| | | ___ _ __
   | || '__/ _' \ \ / / _ \ | |/ _ \ '__|
   | || | | (_| |\ V /  __/ | |  __/ |
   |_||_|  \__,_| \_/ \___|_|_|\___|_|
                                         `

type MainMenu struct {
	baseScreen
}

func NewMainMenu() *MainMenu {
	return &MainMenu{newBaseScreen()}
}

func (s *MainMenu) Render() {
	options := s.createOptions()
	options.SetFocused(true)

	logo := tui.NewLabel(logo)
	logo.SetStyleName("logo")

	window := tui.NewVBox(
		tui.NewPadder(10, 1, logo),
		tui.NewPadder(17, 0, tui.NewLabel("Welcome to Traveller Tools!")),
		tui.NewPadder(1, 1, options),
	)
	window.SetBorder(true)

	wrapper := tui.NewVBox(
		tui.NewSpacer(),
		window,
		tui.NewSpacer(),
	)
	s.content = tui.NewHBox(tui.NewSpacer(), wrapper, tui.NewSpacer())

	s.render()
}

func (s *MainMenu) createOptions() (options *tui.List) {
	mainMenuOptions := []string{
		"Character Creation (In progress)",
		"Character List (Not implemented)",
		"Trade Calculator (Not implemented)",
		"",
		"Exit",
	}

	options = tui.NewList()
	for _, o := range mainMenuOptions {
		options.AddItems(o)
	}
	options.Select(0)
	options.OnItemActivated(func(t *tui.List) {
		s.status.SetText(mainMenuOptions[t.Selected()])
	})

	return
}
