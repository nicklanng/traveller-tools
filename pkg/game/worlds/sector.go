package worlds

import "sort"

type Sector struct {
	Name               string                `json:"name"`
	ImperialCoordinate ImperialCoordinate    `json:"coords"`
	SubSectors         map[string]*SubSector `json:"subSectors"`
}

func NewSector(name string, coords ImperialCoordinate) (sector *Sector) {
	sector = new(Sector)
	sector.Name = name
	sector.ImperialCoordinate = coords
	sector.SubSectors = make(map[string]*SubSector, 16)
	for _, v := range "ABCDEFGHIJKLMNOP" {
		sector.SubSectors[string(v)] = NewSubSector("")
	}

	return
}

func (s *Sector) String() (str string) {
	str = s.Name + ", " + s.ImperialCoordinate.String() + "\n"

	var keys []string
	for k := range s.SubSectors {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		str += "  " + k + ":\n    " + s.SubSectors[k].String() + "\n"
	}
	return
}
