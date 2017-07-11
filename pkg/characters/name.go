package characters

type Name interface {
	IndexedName() string
	String() string
}

// Name represents a classic Human name.
type HumanName struct {
	Forename string
	Surname  string
	Alias    string
}

func (n HumanName) IndexedName() (s string) {
	return n.Surname + ", " + n.Forename
}

func (n HumanName) String() (s string) {
	s = n.Forename + " "

	if n.Alias != "" {
		s += "\"" + n.Alias + "\" "
	}

	s += n.Surname

	return s
}
