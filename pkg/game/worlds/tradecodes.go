package worlds

var (
	TradeCodeAgricultural    TradeCodeName = "Agricultural"
	TradeCodeAsteroid        TradeCodeName = "Asteroid"
	TradeCodeBarren          TradeCodeName = "Barren"
	TradeCodeDesert          TradeCodeName = "Desert"
	TradeCodeFluidOceans     TradeCodeName = "Fluid Oceans"
	TradeCodeGarden          TradeCodeName = "Garden"
	TradeCodeHighPopulation  TradeCodeName = "High Population"
	TradeCodeHighTechnology  TradeCodeName = "High Technology"
	TradeCodeIceCapped       TradeCodeName = "Ice-Capped"
	TradeCodeIndustrial      TradeCodeName = "Industrial"
	TradeCodeLowPopulation   TradeCodeName = "Low Population"
	TradeCodeLowTechnology   TradeCodeName = "Low Technology"
	TradeCodeNonAgricultural TradeCodeName = "Non-Agricultural"
	TradeCodeNonIndustrial   TradeCodeName = "Non-Industrial"
	TradeCodePoor            TradeCodeName = "Poor"
	TradeCodeRich            TradeCodeName = "Rich"
	TradeCodeVacuum          TradeCodeName = "Vacuum"
	TradeCodeWaterWorld      TradeCodeName = "Water World"

	TradeCodes = map[TradeCodeName]TradeCode{
		TradeCodeAgricultural:    TradeCode{TradeCodeAgricultural, "Ag", "Agricultural worlds are dedicated to farming and food production. Often, they are divided into vast semi-feudal estates."},
		TradeCodeAsteroid:        TradeCode{TradeCodeAsteroid, "As", "Asteroids are usually mining colonies, but can also be orbital factories or colonies."},
		TradeCodeBarren:          TradeCode{TradeCodeBarren, "Ba", "Barren worlds are uncolonised and empty."},
		TradeCodeDesert:          TradeCode{TradeCodeDesert, "De", "Desert worlds are dry and barely habitable."},
		TradeCodeFluidOceans:     TradeCode{TradeCodeFluidOceans, "Fl", "Fluid Oceans are worlds where the surface liquid is something other than water, and so are incompatible with Earth-derived life."},
		TradeCodeGarden:          TradeCode{TradeCodeGarden, "Ga", "Garden worlds are Earth-like."},
		TradeCodeHighPopulation:  TradeCode{TradeCodeHighPopulation, "Hi", "High Population worlds have a population in the billions."},
		TradeCodeHighTechnology:  TradeCode{TradeCodeHighTechnology, "Ht", "High Technology worlds are among the most technologically advanced in the Imperium."},
		TradeCodeIceCapped:       TradeCode{TradeCodeIceCapped, "IC", "Ice-Capped worlds have most of their surface liquid frozen in polar ice caps, and are cold and dry."},
		TradeCodeIndustrial:      TradeCode{TradeCodeIndustrial, "In", "Industrial worlds are dominated by factories and cities."},
		TradeCodeLowPopulation:   TradeCode{TradeCodeLowPopulation, "Lo", "Low Population worlds have a population of only a few thousand or less."},
		TradeCodeLowTechnology:   TradeCode{TradeCodeLowTechnology, "Lt", "Low Technology worlds are pre- industrial and cannot produce advanced goods."},
		TradeCodeNonAgricultural: TradeCode{TradeCodeNonAgricultural, "Na", "Non-Agricultural worlds are too dry or barren to support their populations using conventional food production."},
		TradeCodeNonIndustrial:   TradeCode{TradeCodeNonIndustrial, "NI", "Non-Industrial worlds are too low-population to maintain an industrial base."},
		TradeCodePoor:            TradeCode{TradeCodePoor, "Po", "Poor worlds lack resources, viable land or sufficient population to be anything other than marginal colonies."},
		TradeCodeRich:            TradeCode{TradeCodeRich, "Ri", "Rich worlds are blessed with a stable government and viable biosphere, making them economic powerhouses."},
		TradeCodeVacuum:          TradeCode{TradeCodeVacuum, "Va", "Vacuum worlds have no atmosphere."},
		TradeCodeWaterWorld:      TradeCode{TradeCodeWaterWorld, "Wa", "Water Worlds are nearly entirely water-ocean."},
	}
)

type TradeCodeName string

type TradeCode struct {
	Name        TradeCodeName
	Code        string
	Description string
}

func (t TradeCode) String() string {
	return t.Code
}
