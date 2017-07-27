package serialization

import "github.com/nicklanng/traveller-tools/pkg/game/worlds"

type SubSector struct {
	Name   string   `json:"name"`
	Worlds []string `json:"worlds"`
}

func serializeSubsector(subSector *worlds.SubSector) (s SubSector) {
	s.Name = subSector.Name

	for _, world := range subSector.Worlds {
		s.Worlds = append(s.Worlds, world.String())
	}

	return
}
