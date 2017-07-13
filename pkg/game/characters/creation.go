package characters

import "github.com/nicklanng/traveller-tools/pkg/game/planets"

var (
	// BackgroundSkills is the set of skills a character can learn on their homeworld.
	BackgroundSkills = map[planets.TradeCodeName]Skill{
		planets.TradeCodeAgricultural:   Skills[SkillAnimals],
		planets.TradeCodeAsteroid:       Skills[SkillZeroG],
		planets.TradeCodeDesert:         Skills[SkillSurvival],
		planets.TradeCodeFluidOceans:    Skills[SkillSeafarer],
		planets.TradeCodeGarden:         Skills[SkillAnimals],
		planets.TradeCodeHighTechnology: Skills[SkillComputers],
		planets.TradeCodeHighPopulation: Skills[SkillStreetwise],
		planets.TradeCodeIceCapped:      Skills[SkillVaccSuit],
		planets.TradeCodeIndustrial:     Skills[SkillTrade],
		planets.TradeCodeLowTechnology:  Skills[SkillSurvival],
		planets.TradeCodePoor:           Skills[SkillAnimals],
		planets.TradeCodeRich:           Skills[SkillCarouse],
		planets.TradeCodeWaterWorld:     Skills[SkillSeafarer],
		planets.TradeCodeVacuum:         Skills[SkillVaccSuit],
	}

	// Education is the set of skills that can be chosen as education at character creation.
	Education = []Skill{
		Skills[SkillAdmin],
		Skills[SkillAdvocate],
		Skills[SkillArt],
		Skills[SkillCarouse],
		Skills[SkillComms],
		Skills[SkillComputers],
		Skills[SkillDrive],
		Skills[SkillEngineer],
		Skills[SkillLanguage],
		Skills[SkillMedic],
		Skills[SkillPhysicalScience],
		Skills[SkillLifeScience],
		Skills[SkillSocialScience],
		Skills[SkillSpaceScience],
		Skills[SkillTrade],
	}
)
