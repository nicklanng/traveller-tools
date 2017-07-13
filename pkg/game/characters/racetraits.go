package characters

var (
	RaceTraitArmored = &RaceTrait{
		Name:        "Armored",
		Description: "The species possess thick fur, scales, a bony exoskeleton or other natural protection that gives it one point of natural armour. This works in the same way as normal armour.",
	}
	RaceTraitAquatic = &RaceTrait{
		Name:        "Aquatic",
		Description: "The species is adapted to life underwater. It can breathe underwater, or hold its breath for a long period (Endurance 10 minutes on average). If amphibious, its Dexterity is halved on land. If the species is not amphibious, then it cannot operate out of water without mechanical aid or telepresence.",
	}
	RaceTraitAtmosphericRequirements = &RaceTrait{
		Name:        "Atmospheric Requirements",
		Description: "The species requires an unusual combination of gasses to breathe, and cannot survive in most atmospheres without artificial aid. Creatures with this trait usually come from homeworlds with an Exotic atmosphere.",
	}
	RaceTraitEngineered = &RaceTrait{
		Name:        "Engineered",
		Description: "The species has been altered by some external factor to adapt to changed circumstances or a different environment. Medical treatment of Engineered species by a facility of a lower Technology Level than that at which the species was created receives a negative DM equal to the difference. Some strains of humans have been engineered to tolerate unusual atmospheres, to live on water worlds, or to cope with the stresses of space travel.",
	}
	RaceTraitFastMetabolism = &RaceTrait{
		Name:        "Fast Metabolism",
		Description: "Creatures with a fast metabolism require more food than most species, and their life support costs are doubled. In combat, fast-metabolism creatures gain a +2 initiative bonus. Fast-metabolism creatures halve their Endurance for the purposes of determining fatigue.",
	}
	RaceTraitFeral = &RaceTrait{
		Name:        "Feral",
		Description: "Feral species are uncivilised, regardless of their technological knowledge. Often, such species have acquired their technology from other races, or from Ancient ruins. Feral species are much less likely to accept the laws of more civilised societies. Feral species roll Education on 1d6 only.",
	}
	RaceTraitFlyer = &RaceTrait{
		Name:        "Flyer",
		Description: "The species can fly using wings, glider membranes, gasbags or other means. Characters of this species gain the Athletics (flying) skill at Level 0 and can travel at a speed noted in their description. Flying creatures who are aloft must spend one minor action every round on movement or stall and fall out of the air.\n• Winged flight is tiring and can only be sustained for a number of rounds equal to the creature’s Endurance before requiring a like amount of rest. Some specialised avians can increase this to minutes or even hours equal to Endurance.\n• Species with glider membranes cannot gain altitude while flying. They descend one meter every time they move forwards and cannot use more than one minor action for flying movement in a round.\n• Species that float using gasbags or some other method do not need to move to remain aloft. They are typically slower than other fliers, though.",
	}
	RaceTraitLarge = &RaceTrait{
		Name:        "Large",
		Description: "The species is considerably larger than the average for sophonts. Large creatures generally have a Strength and Endurance of 3d6 or even 4d6, and a Dexterity of 1d6. Life support requirements for Large creatures are doubled and they often have trouble operating in buildings and spacecraft designed for smaller creatures.\n• Some Large creatures are described as Huge. Attacks against Huge creatures receive a +1 DM to hit.",
	}
	RaceTraitNaturalWeapon = &RaceTrait{
		Name:        "Natural Weapon",
		Description: "The species has a natural weapon, such as claws, a strong bite or a poisonous stinger. Such weapons are usable at Personal range and deal 1 damage. The creature gains Melee (natural weapons) at level 0.",
	}
	RaceTraitNoFineManipulators = &RaceTrait{
		Name:        "No Fine Manipulators",
		Description: "The species has no fingers or other prehensile appendages, preventing them from easily picking things up, pushing small buttons, reaching into tight spaces, and so on. The species will need special equipment to function in most civilised settings.",
	}
	RaceTraitNotableStrength = &RaceTrait{
		Name:        "Notable Strength",
		Description: "Some species are notably dextrous, intelligent, tough or strong. Characters from such races have a positive Dice Modifier when rolling for that characteristic (+2 unless otherwise specified), and their racial maximum for that characteristic is increased by the same amount. Any characteristic can be Notable.",
	}
	RaceTraitNotableDexterity = &RaceTrait{
		Name:        "Notable Dexterity",
		Description: "Some species are notably dextrous, intelligent, tough or strong. Characters from such races have a positive Dice Modifier when rolling for that characteristic (+2 unless otherwise specified), and their racial maximum for that characteristic is increased by the same amount. Any characteristic can be Notable.",
	}
	RaceTraitNotableEndurance = &RaceTrait{
		Name:        "Notable Endurance",
		Description: "Some species are notably dextrous, intelligent, tough or strong. Characters from such races have a positive Dice Modifier when rolling for that characteristic (+2 unless otherwise specified), and their racial maximum for that characteristic is increased by the same amount. Any characteristic can be Notable.",
	}
	RaceTraitNotableIntelligence = &RaceTrait{
		Name:        "Notable Intelligence",
		Description: "Some species are notably dextrous, intelligent, tough or strong. Characters from such races have a positive Dice Modifier when rolling for that characteristic (+2 unless otherwise specified), and their racial maximum for that characteristic is increased by the same amount. Any characteristic can be Notable.",
	}
	RaceTraitNotableEducation = &RaceTrait{
		Name:        "Notable Education",
		Description: "Some species are notably dextrous, intelligent, tough or strong. Characters from such races have a positive Dice Modifier when rolling for that characteristic (+2 unless otherwise specified), and their racial maximum for that characteristic is increased by the same amount. Any characteristic can be Notable.",
	}
	RaceTraitNotableSocialStanding = &RaceTrait{
		Name:        "Notable Social Standing",
		Description: "Some species are notably dextrous, intelligent, tough or strong. Characters from such races have a positive Dice Modifier when rolling for that characteristic (+2 unless otherwise specified), and their racial maximum for that characteristic is increased by the same amount. Any characteristic can be Notable.",
	}
	RaceTraitPsionic = &RaceTrait{
		Name:        "Psionic",
		Description: "All members of the species are Psionic, and may determine their Psionic Strength and talents at the start of character generation.",
	}
	RaceTraitSmall = &RaceTrait{
		Name:        "Small",
		Description: "Small species generally have a Strength and Endurance of only 1d6, and a Dexterity of 3d6. The minimum size for a sophont is about half that of a human, as smaller creatures lack the cranial capacity for sophont-level intelligence. This assumes that the species has a brain structure comparable to humans. Species with a more distributed neural structure, hive intelligences, or artificial intelligences can be even smaller.\n• Some Small creatures are described as Tiny. Attacks against Tiny creatures receive a –1 DM to hit.",
	}
	RaceTraitSlowMetabolism = &RaceTrait{
		Name:        "Slow Metabolism",
		Description: "Creatures with a slow metabolism require less food than most species, and their life supportcostsarehalved. In combat, slow-metabolism creatures suffer a –2 initiative penalty.",
	}
	RaceTraitUplifted = &RaceTrait{
		Name:        "Uplifted",
		Description: "This species was originally non-sentient, but has been raised to a higher intelligence by another species. Uplifted races generally become client species of their patron. Two common uplifted animals are apes and dolphins:\n• Uplifted apes have Notable Strength and Endurance (+2) but all other characteristics are Weak (–2). They have the Uplifted trait.\n• Uplifted dolphins have Notable Strength (+4) and Notable Endurance (+2) but Weak Intelligence, Education and Social Standing (–2). They have the Uplifted, Aquatic (fully aquatic, air-breathers) and No Fine Manipulators traits",
	}
	RaceTraitWeakStrength = &RaceTrait{
		Name:        "Weak Strength",
		Description: "The opposite of Notable (Characteristic), some species are weaker, less resilient or less well educated than others. Characters from such races have a negative Dice Modifier when rolling for that characteristic (–2 unless otherwise specified), and their racial maximum for that characteristic is decreased by the same amount. Any characteristic can be Weak.",
	}
	RaceTraitWeakDexterity = &RaceTrait{
		Name:        "Weak Dexterity",
		Description: "The opposite of Notable (Characteristic), some species are weaker, less resilient or less well educated than others. Characters from such races have a negative Dice Modifier when rolling for that characteristic (–2 unless otherwise specified), and their racial maximum for that characteristic is decreased by the same amount. Any characteristic can be Weak.",
	}
	RaceTraitWeakEndurance = &RaceTrait{
		Name:        "Weak Endurance",
		Description: "The opposite of Notable (Characteristic), some species are weaker, less resilient or less well educated than others. Characters from such races have a negative Dice Modifier when rolling for that characteristic (–2 unless otherwise specified), and their racial maximum for that characteristic is decreased by the same amount. Any characteristic can be Weak.",
	}
	RaceTraitWeakIntelligence = &RaceTrait{
		Name:        "Weak Intelligence",
		Description: "The opposite of Notable (Characteristic), some species are weaker, less resilient or less well educated than others. Characters from such races have a negative Dice Modifier when rolling for that characteristic (–2 unless otherwise specified), and their racial maximum for that characteristic is decreased by the same amount. Any characteristic can be Weak.",
	}
	RaceTraitWeakEducation = &RaceTrait{
		Name:        "Weak Education",
		Description: "The opposite of Notable (Characteristic), some species are weaker, less resilient or less well educated than others. Characters from such races have a negative Dice Modifier when rolling for that characteristic (–2 unless otherwise specified), and their racial maximum for that characteristic is decreased by the same amount. Any characteristic can be Weak.",
	}
	RaceTraitWeakSocialStanding = &RaceTrait{
		Name:        "Weak Social Standing",
		Description: "The opposite of Notable (Characteristic), some species are weaker, less resilient or less well educated than others. Characters from such races have a negative Dice Modifier when rolling for that characteristic (–2 unless otherwise specified), and their racial maximum for that characteristic is decreased by the same amount. Any characteristic can be Weak.",
	}
)

type RaceTrait struct {
	Name        string
	Description string
}
