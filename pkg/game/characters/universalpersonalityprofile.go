package characters

import (
	"encoding/hex"
	"strings"
)

// UniversalPersonalityProfile expresses the basic characteristics of a character.
type UniversalPersonalityProfile struct {
	Strength       int
	Dexterity      int
	Endurance      int
	Intelligence   int
	Education      int
	SocialStanding int
}

func (u UniversalPersonalityProfile) String() string {
	b := []byte{
		byte(u.Strength)<<4 + byte(u.Dexterity),
		byte(u.Endurance)<<4 + byte(u.Intelligence),
		byte(u.Education)<<4 + byte(u.SocialStanding),
	}

	return strings.ToUpper(hex.EncodeToString(b))
}

// CharacteristicModifier calculates the modifier for the given score.
func CharacteristicModifier(characteristicScore int) int {
	modTable := []int{-3, -2, -2, -1, -1, -1, 0, 0, 0, 1, 1, 1, 2, 2, 2, 3}
	return modTable[characteristicScore]
}
