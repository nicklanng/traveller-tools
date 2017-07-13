package main

import (
	"fmt"

	"github.com/nicklanng/traveller-tools/pkg/app/screens"
	"github.com/nicklanng/traveller-tools/pkg/game/characters"
	"github.com/nicklanng/traveller-tools/pkg/game/dice"
)

func main() {

	upp := characters.UniversalPersonalityProfile{
		Strength:       dice.D6() + dice.D6(),
		Dexterity:      dice.D6() + dice.D6(),
		Endurance:      dice.D6() + dice.D6(),
		Intelligence:   dice.D6() + dice.D6(),
		Education:      dice.D6() + dice.D6(),
		SocialStanding: dice.D6() + dice.D6(),
	}

	c := characters.Character{
		Name: characters.HumanName{"Bob", "Thomas", "DangerBob"},
		Age:  18,
		Race: characters.RaceHuman,
		UPP:  upp,
	}

	fmt.Println(c.String())

	mm := screens.NewMainMenu()
	mm.Render()
}
