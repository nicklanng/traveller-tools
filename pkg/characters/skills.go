package characters

var (
	SkillAdmin            SkillName = "Admin"
	SkillAdvocate         SkillName = "Advocate"
	SkillAnimals          SkillName = "Animals"
	SkillAthletics        SkillName = "Athletics"
	SkillArt              SkillName = "Art"
	SkillAstrogation      SkillName = "Astrogation"
	SkillBattleDress      SkillName = "Battle Dress"
	SkilBroker            SkillName = "Broker"
	SkillCarouse          SkillName = "Carouse"
	SkillComms            SkillName = "Comms"
	SkillComputers        SkillName = "Computers"
	SkillDeception        SkillName = "Deception"
	SkillDiplomat         SkillName = "Diplomat"
	SkillDrive            SkillName = "Drive"
	SkillEngineer         SkillName = "Engineer"
	SkillExplosives       SkillName = "Explosives"
	SkillFlyer            SkillName = "Flyer"
	SkillGambler          SkillName = "Gambler"
	SkillGunner           SkillName = "Gunner"
	SkillGunCombat        SkillName = "Gun Combat"
	SkillHeavyWeapons     SkillName = "Heavy Weapons"
	SkillInvestigate      SkillName = "Investigate"
	SkillJackOfAllTrades  SkillName = "Jack of All Trades"
	SkillLanguage         SkillName = "Language"
	SkillLeadership       SkillName = "Leadership"
	SkillLifeScience      SkillName = "Life Science"
	SkillMechanic         SkillName = "Mechanic"
	SkillMedic            SkillName = "Medic"
	SkillMelee            SkillName = "Melee"
	SkillNavigation       SkillName = "Navigation"
	SkillPersuade         SkillName = "Persuade"
	SkillPilot            SkillName = "Pilot"
	SkillPhysicalScience  SkillName = "Physical Science"
	SkillRecon            SkillName = "Recon"
	SkillRemoteOperations SkillName = "Remote Operations"
	SkillSeafarer         SkillName = "Seafarer"
	SkillSensors          SkillName = "Sensors"
	SkillSocialScience    SkillName = "Social Science"
	SkillSpaceScience     SkillName = "Space Science"
	SkillStealth          SkillName = "Stealth"
	SkillSteward          SkillName = "Steward"
	SkillStreetwise       SkillName = "Streetwise"
	SkillSurvival         SkillName = "Survival"
	SkillTactics          SkillName = "Tactics"
	SkillTrade            SkillName = "Trade"
	SkillVaccSuit         SkillName = "VaccSuit"
	SkillZeroG            SkillName = "Zero-G"

	Skills = SkillList{
		SkillAdmin:            BasicSkill{SkillAdmin, 0},
		SkillAdvocate:         BasicSkill{SkillAdvocate, 0},
		SkillAnimals:          SpecializedSkill{SkillAnimals, 0, []string{"Riding", "Veterinary", "Training", "Farming"}, ""},
		SkillAthletics:        SpecializedSkill{"Athletics", 0, []string{"Co-ordination", "Endurance", "Strength", "Flying"}, ""},
		SkillArt:              SpecializedSkill{SkillArt, 0, []string{"Acting", "Dance", "Holography", "Instrument", "Sculpting", "Writing"}, ""},
		SkillAstrogation:      BasicSkill{SkillAstrogation, 0},
		SkillBattleDress:      BasicSkill{SkillBattleDress, 0},
		SkilBroker:            BasicSkill{SkilBroker, 0},
		SkillCarouse:          BasicSkill{SkillCarouse, 0},
		SkillComms:            BasicSkill{SkillComms, 0},
		SkillComputers:        BasicSkill{SkillComputers, 0},
		SkillDeception:        BasicSkill{SkillDeception, 0},
		SkillDiplomat:         BasicSkill{SkillDiplomat, 0},
		SkillDrive:            SpecializedSkill{SkillDrive, 0, []string{"Mole", "Tracked", "Wheeled"}, ""},
		SkillEngineer:         SpecializedSkill{SkillEngineer, 0, []string{"Manoeuvre Drive", "Jump Drive", "Electronics", "Life Support", "Power"}, ""},
		SkillExplosives:       BasicSkill{SkillExplosives, 0},
		SkillFlyer:            SpecializedSkill{SkillFlyer, 0, []string{"Grav", "Rotor", "Wing"}, ""},
		SkillGambler:          BasicSkill{SkillGambler, 0},
		SkillGunner:           SpecializedSkill{SkillGunner, 0, []string{"Turrets", "Ortillery", "Screens", "Capital Weapons"}, ""},
		SkillGunCombat:        SpecializedSkill{SkillGunCombat, 0, []string{"Slug Rifle", "Slug Pistol", "Shotgun", "Energy Rifle", "Energy Pistol"}, ""},
		SkillHeavyWeapons:     SpecializedSkill{SkillHeavyWeapons, 0, []string{"Launchers", "Man Portable Artillery", "Field Artillery"}, ""},
		SkillInvestigate:      BasicSkill{SkillInvestigate, 0},
		SkillJackOfAllTrades:  BasicSkill{SkillJackOfAllTrades, 0},
		SkillLanguage:         SpecializedSkill{SkillLanguage, 0, []string{"Anglic", "Vilani", "Zdetl", "Oynprith"}, ""},
		SkillLeadership:       BasicSkill{SkillLeadership, 0},
		SkillLifeScience:      SpecializedSkill{SkillLifeScience, 0, []string{"Biology", "Cybernetics", "Genetics", "Psionicology"}, ""},
		SkillMechanic:         BasicSkill{SkillMechanic, 0},
		SkillMedic:            BasicSkill{SkillMedic, 0},
		SkillMelee:            SpecializedSkill{SkillMelee, 0, []string{"Unarmed Combat", "Blade", "Bludgeon", "Natural Weapons"}, ""},
		SkillNavigation:       BasicSkill{SkillNavigation, 0},
		SkillPersuade:         BasicSkill{SkillPersuade, 0},
		SkillPilot:            SpecializedSkill{SkillPilot, 0, []string{"Small Craft", "Spacecraft", "Capital Ships"}, ""},
		SkillPhysicalScience:  SpecializedSkill{SkillPhysicalScience, 0, []string{"Physics", "Chemistry", "Electronics"}, ""},
		SkillRecon:            BasicSkill{SkillRecon, 0},
		SkillRemoteOperations: BasicSkill{SkillRemoteOperations, 0},
		SkillSeafarer:         SpecializedSkill{SkillSeafarer, 0, []string{"Sail", "Submarine", "Ocean Ships", "Motorboats"}, ""},
		SkillSensors:          BasicSkill{SkillSensors, 0},
		SkillSocialScience:    SpecializedSkill{SkillSocialScience, 0, []string{"Archeology", "Economics", "History", "Linguistics", "Philosophy", "Psychology", "Sophontology"}, ""},
		SkillSpaceScience:     SpecializedSkill{SkillSpaceScience, 0, []string{"Planetology", "Robotics", "Xenology"}, ""},
		SkillStealth:          BasicSkill{SkillStealth, 0},
		SkillSteward:          BasicSkill{SkillSteward, 0},
		SkillStreetwise:       BasicSkill{SkillStreetwise, 0},
		SkillSurvival:         BasicSkill{SkillSurvival, 0},
		SkillTactics:          SpecializedSkill{SkillTactics, 0, []string{"Military Tactics", "Naval Tactics"}, ""},
		SkillTrade:            BasicSkill{SkillTrade, 0},
		SkillVaccSuit:         BasicSkill{SkillVaccSuit, 0},
		SkillZeroG:            BasicSkill{SkillZeroG, 0},
	}
)

type SkillName string
type SkillList map[SkillName]Skill

type Skill interface {
	Modifier(specialization string) int
}

type BasicSkill struct {
	Name     SkillName
	Training int
}

func (s BasicSkill) Modifier(specialization string) int {
	return s.Training
}

func (s BasicSkill) String() string {
	return string(s.Name) + ": " + string(s.Training)
}

type SpecializedSkill struct {
	Name                   SkillName
	Training               int
	Specializations        []string
	SelectedSpecialization string
}

func (s SpecializedSkill) Modifier(specialization string) int {
	if s.SelectedSpecialization != specialization {
		return 0
	}

	return s.Training
}

func (s SpecializedSkill) String() string {
	if s.SelectedSpecialization == "" {
		return string(s.Name) + ": " + string(s.Training)
	}

	return string(s.Name) + " (" + s.SelectedSpecialization + "): " + string(s.Training)
}
