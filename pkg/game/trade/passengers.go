package trade

import (
	"github.com/nicklanng/traveller-tools/pkg/game/dice"
	"github.com/nicklanng/traveller-tools/pkg/game/worlds"
)

var (
	passengersCurrentWorldTrade = map[worlds.TradeCodeName]int{
		worlds.TradeCodeAgricultural:    0,
		worlds.TradeCodeAsteroid:        1,
		worlds.TradeCodeBarren:          -5,
		worlds.TradeCodeDesert:          -1,
		worlds.TradeCodeFluidOceans:     0,
		worlds.TradeCodeGarden:          2,
		worlds.TradeCodeHighPopulation:  0,
		worlds.TradeCodeHighTechnology:  0,
		worlds.TradeCodeIceCapped:       1,
		worlds.TradeCodeIndustrial:      2,
		worlds.TradeCodeLowPopulation:   0,
		worlds.TradeCodeLowTechnology:   0,
		worlds.TradeCodeNonAgricultural: 0,
		worlds.TradeCodeNonIndustrial:   0,
		worlds.TradeCodePoor:            -2,
		worlds.TradeCodeRich:            -1,
		worlds.TradeCodeVacuum:          0,
		worlds.TradeCodeWaterWorld:      0,
	}
	passengersCurrentWorldTravelZone = map[worlds.TravelZone]int{
		worlds.TravelZoneUnclassified: 0,
		worlds.TravelZoneGreen:        0,
		worlds.TravelZoneAmber:        2,
		worlds.TravelZoneRed:          4,
	}
	passengersDestinationWorldTrade = map[worlds.TradeCodeName]int{
		worlds.TradeCodeAgricultural:    0,
		worlds.TradeCodeAsteroid:        -1,
		worlds.TradeCodeBarren:          -5,
		worlds.TradeCodeDesert:          -1,
		worlds.TradeCodeFluidOceans:     0,
		worlds.TradeCodeGarden:          2,
		worlds.TradeCodeHighPopulation:  4,
		worlds.TradeCodeHighTechnology:  0,
		worlds.TradeCodeIceCapped:       -1,
		worlds.TradeCodeIndustrial:      1,
		worlds.TradeCodeLowPopulation:   -4,
		worlds.TradeCodeLowTechnology:   0,
		worlds.TradeCodeNonAgricultural: 0,
		worlds.TradeCodeNonIndustrial:   -1,
		worlds.TradeCodePoor:            -1,
		worlds.TradeCodeRich:            2,
		worlds.TradeCodeVacuum:          0,
		worlds.TradeCodeWaterWorld:      0,
	}
	passengersDestinationWorldTravelZone = map[worlds.TravelZone]int{
		worlds.TravelZoneUnclassified: 0,
		worlds.TravelZoneGreen:        0,
		worlds.TravelZoneAmber:        -2,
		worlds.TravelZoneRed:          -4,
	}
)

type PassengerAvailability struct {
	LowPassages    int
	MiddlePassages int
	HighPassages   int
}

func FindAvailablePassengers(currentWorld *worlds.World, destinationWorld *worlds.World) (pa PassengerAvailability) {
	passengerTraffic := 0
	var mod int
	var ok bool

	// add current world trade code modifiers
	for _, tc := range currentWorld.TradeCodes {
		mod, ok = passengersCurrentWorldTrade[tc.Name]
		if ok {
			passengerTraffic += mod
		}
	}

	// add current world travel zone modifiers
	mod, ok = passengersCurrentWorldTravelZone[currentWorld.TravelZone]
	if ok {
		passengerTraffic += mod
	}

	// add current world population as modifier
	passengerTraffic += currentWorld.Population

	// add destination world trade code modifiers
	for _, tc := range destinationWorld.TradeCodes {
		mod, ok = passengersDestinationWorldTrade[tc.Name]
		if !ok {
			continue
		}
		passengerTraffic += mod
	}

	// add destination world travel zone modifiers
	mod, ok = passengersDestinationWorldTravelZone[destinationWorld.TravelZone]
	if ok {
		passengerTraffic += mod
	}

	if passengerTraffic > 16 {
		passengerTraffic = 16
	}

	switch passengerTraffic {
	case 1:
		pa = PassengerAvailability{dice.D6(2) - 6, dice.D6(1) - 2, 0}
	case 2:
		pa = PassengerAvailability{dice.D6(2), dice.D6(1), 0}
	case 3:
		pa = PassengerAvailability{dice.D6(2), dice.D6(2) - dice.D6(1), dice.D6(2) - dice.D6(2)}
	case 4:
		pa = PassengerAvailability{dice.D6(3) - dice.D6(1), dice.D6(2) - dice.D6(1), dice.D6(2) - dice.D6(1)}
	case 5:
		pa = PassengerAvailability{dice.D6(3) - dice.D6(1), dice.D6(3) - dice.D6(2), dice.D6(2) - dice.D6(1)}
	case 6:
		pa = PassengerAvailability{dice.D6(3), dice.D6(3) - dice.D6(2), dice.D6(3) - dice.D6(2)}
	case 7:
		pa = PassengerAvailability{dice.D6(3), dice.D6(3) - dice.D6(1), dice.D6(3) - dice.D6(2)}
	case 8:
		pa = PassengerAvailability{dice.D6(4), dice.D6(3) - dice.D6(1), dice.D6(3) - dice.D6(1)}
	case 9:
		pa = PassengerAvailability{dice.D6(4), dice.D6(3), dice.D6(3) - dice.D6(1)}
	case 10:
		pa = PassengerAvailability{dice.D6(5), dice.D6(3), dice.D6(3) - dice.D6(1)}
	case 11:
		pa = PassengerAvailability{dice.D6(5), dice.D6(4), dice.D6(3)}
	case 12:
		pa = PassengerAvailability{dice.D6(6), dice.D6(4), dice.D6(3)}
	case 13:
		pa = PassengerAvailability{dice.D6(6), dice.D6(4), dice.D6(4)}
	case 14:
		pa = PassengerAvailability{dice.D6(7), dice.D6(5), dice.D6(4)}
	case 15:
		pa = PassengerAvailability{dice.D6(8), dice.D6(5), dice.D6(4)}
	case 16:
		pa = PassengerAvailability{dice.D6(9), dice.D6(6), dice.D6(5)}
	default:
		pa = PassengerAvailability{0, 0, 0}
	}

	if pa.LowPassages < 0 {
		pa.LowPassages = 0
	}
	if pa.MiddlePassages < 0 {
		pa.MiddlePassages = 0
	}
	if pa.HighPassages < 0 {
		pa.HighPassages = 0
	}

	return
}
