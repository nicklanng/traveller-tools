package worlds

type SubSector struct {
	Name   string   `json:"name"`
	Worlds []*World `json:"worlds"`
}

func NewSubSector(name string) (subSector *SubSector) {
	subSector = new(SubSector)
	subSector.Name = name
	return
}

func (s SubSector) String() (str string) {
	str = s.Name + "\n"
	for _, v := range s.Worlds {
		str += v.String() + "\n"
	}
	return
}
