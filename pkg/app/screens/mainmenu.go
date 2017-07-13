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

func NewMainMenu() tui.Widget {
	options := createOptions()
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

	content := tui.NewHBox(
		tui.NewSpacer(),
		wrapper,
		tui.NewSpacer(),
	)

	return content
}

type menuOption struct {
	text   string
	action func()
}

func createOptions() (options *tui.List) {
	mainMenuOptions := []*menuOption{
		&menuOption{
			text: "Create a Character",
			action: func() {
				ui.SetWidget(NewCreateCharacter())
			},
		},
		&menuOption{
			text: "List Characters (Not implemented)",
		},
		&menuOption{
			text: "Trade Calculator (Not implemented)",
		},
		&menuOption{
			text: "",
		},
		&menuOption{
			text: "Exit",
			action: func() {
				ui.Quit()
			},
		},
	}

	options = tui.NewList()
	for _, o := range mainMenuOptions {
		options.AddItems(o.text)
	}
	options.OnSelectionChanged(func(t *tui.List) {
		status.SetText(mainMenuOptions[t.Selected()].text)
	})
	options.OnItemActivated(func(t *tui.List) {
		action := mainMenuOptions[t.Selected()].action
		if action != nil {
			action()
		}
	})
	options.Select(0)

	return
}
