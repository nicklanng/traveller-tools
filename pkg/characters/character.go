package characters

import "strconv"

// Character defines everything that makes up a character.
type Character struct {
	Name Name
	Age  int
	Race *Race
	UPP  UniversalPersonalityProfile
}

func (c *Character) String() string {
	return "Name: " + c.Name.String() + "\n" +
		"Age: " + strconv.Itoa(c.Age) + "\n" +
		"Race: " + c.Race.String() + "\n" +
		"UPP: " + c.UPP.String() + "\n"
}
