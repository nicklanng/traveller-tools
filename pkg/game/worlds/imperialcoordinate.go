package worlds

import "strconv"

type ImperialCoordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (ic ImperialCoordinate) String() string {
	return strconv.Itoa(ic.X) + "," + strconv.Itoa(ic.Y)
}
