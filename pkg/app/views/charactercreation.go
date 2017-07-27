package views

import (
	"github.com/leekchan/accounting"
	tui "github.com/marcusolsson/tui-go"
)

func NewCreateCharacter() tui.Widget {
	// left panel
	name := tui.NewLabel("Name: " + "Rinse Concord")
	name.SetStyleName("value")
	race := tui.NewLabel("Race: " + "Human")
	race.SetStyleName("value")
	age := tui.NewLabel("Age:  " + "18")
	age.SetStyleName("value")
	upp := tui.NewLabel("UPP:  " + "A5623B")
	upp.SetStyleName("value")
	profileBox := tui.NewVBox(
		tui.NewPadder(1, 0, name),
		tui.NewPadder(1, 0, race),
		tui.NewPadder(1, 0, age),
		tui.NewPadder(1, 0, upp),
	)
	profileBox.SetTitle("Profile")
	profileBox.SetBorder(true)
	profileBox.SetSizePolicy(tui.Preferred, tui.Minimum)

	careersBox := tui.NewVBox()
	careersBox.SetTitle("Career History")
	careersBox.SetBorder(true)
	careersBox.SetSizePolicy(tui.Preferred, tui.Expanding)

	ac := accounting.Accounting{Symbol: "Cr ", Precision: 0}
	pension := tui.NewLabel("Pension:   " + ac.FormatMoney(0))
	pension.SetStyleName("value")
	debt := tui.NewLabel("Debt:      " + ac.FormatMoney(75345332))
	debt.SetStyleName("value")
	cash := tui.NewLabel("Cash:      " + ac.FormatMoney(50000))
	cash.SetStyleName("value")
	shipCost := tui.NewLabel("Ship cost: " + ac.FormatMoney(59001))
	shipCost.SetStyleName("value")
	financesBox := tui.NewVBox(
		tui.NewPadder(1, 0, pension),
		tui.NewPadder(1, 0, debt),
		tui.NewPadder(1, 0, cash),
		tui.NewPadder(1, 0, shipCost),
	)
	financesBox.SetTitle("Finances")
	financesBox.SetBorder(true)
	financesBox.SetSizePolicy(tui.Preferred, tui.Minimum)

	leftPanel := tui.NewVBox(profileBox, careersBox, financesBox)
	leftPanel.SetSizePolicy(tui.Minimum, tui.Preferred)

	// middle panel
	mainPanel := tui.NewVBox()
	mainPanel.SetTitle("Main Panel")
	mainPanel.SetBorder(true)
	mainPanel.SetSizePolicy(tui.Expanding, tui.Preferred)

	// right panel

	skills := tui.NewList()
	skills.AddItems("Admin 2", "Advocate 0", "Stealth 0")

	skillsBox := tui.NewVBox(tui.NewPadder(1, 0, skills))
	skillsBox.SetTitle("Skills")
	skillsBox.SetBorder(true)
	rightPanel := tui.NewVBox(skillsBox)
	rightPanel.SetSizePolicy(tui.Minimum, tui.Preferred)

	grid := tui.NewHBox(leftPanel, mainPanel, rightPanel)

	return grid
}
