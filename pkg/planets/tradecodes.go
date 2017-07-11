package planets

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
		TradeCodeAgricultural: TradeCode{TradeCodeAgricultural, "Ag", "Agricultural worlds are dedicated to farming and food production. Often, they are divided into vast semi-feudal estates."},
	}
)

type TradeCodeName string

type TradeCode struct {
	Name        TradeCodeName
	Code        string
	Description string
}
