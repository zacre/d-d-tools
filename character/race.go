package character

import "github.com/kiwih/npc-gen/npcgen"

// Race provides a mechanical description of racial features
// Apparently for this project races are only their base stats, so let's roll with that
type Race struct {
	Name                string
	AbilityScoreBonuses npcgen.AbilityScores
	HasSubRace          bool
}

// SubRace provides a mechanical description of subraces, linking them to races
type SubRace struct {
	Name                string
	Race                Race
	AbilityScoreBonuses npcgen.AbilityScores
}

var (
	// Human has the stat bonuses humans provide
	Human = Race{
		Name: "Human",
		AbilityScoreBonuses: npcgen.AbilityScores{
			Str: 1,
			Dex: 1,
			Con: 1,
			Int: 1,
			Wis: 1,
			Cha: 1,
		},
		HasSubRace: false,
	}

	// Elf has the base elf stat bonuses
	Elf = Race{
		Name: "Elf",
		AbilityScoreBonuses: npcgen.AbilityScores{
			Dex: 2,
		},
		HasSubRace: true,
	}
	// HighElf has the base elf bonus plus the high elf special bonus
	HighElf = SubRace{
		Name: "High Elf",
		AbilityScoreBonuses: npcgen.AbilityScores{
			Int: 1,
		},
		Race: Elf,
	}
	// WoodElf has the base elf bonus plus the wood elf special bonus
	WoodElf = SubRace{
		Name: "Wood Elf",
		AbilityScoreBonuses: npcgen.AbilityScores{
			Wis: 1,
		},
		Race: Elf,
	}
	// DarkElf has the base elf bonus plus the dark elf special bonus
	DarkElf = SubRace{
		Name: "Dark Elf",
		AbilityScoreBonuses: npcgen.AbilityScores{
			Cha: 1,
		},
		Race: Elf,
	}

	// Dwarf has the base dwarf stat bonuses
	Dwarf = Race{
		Name: "Dwarf",
		AbilityScoreBonuses: npcgen.AbilityScores{
			Con: 2,
		},
		HasSubRace: true,
	}
	// HillDwarf has the base dwarf bonus plus the hill dwarf special bonus
	HillDwarf = SubRace{
		Name: "Hill Dwarf",
		AbilityScoreBonuses: npcgen.AbilityScores{
			Wis: 1,
		},
		Race: Dwarf,
	}
	// MountainDwarf has the base dwarf bonus plus the mountain dwarf special bonus
	MountainDwarf = SubRace{
		Name: "Mountain Dwarf",
		AbilityScoreBonuses: npcgen.AbilityScores{
			Str: 2,
		},
		Race: Dwarf,
	}

	// Halfling has the base halfling stat bonuses
	Halfling = Race{
		Name: "Halfling",
		AbilityScoreBonuses: npcgen.AbilityScores{
			Dex: 2,
		},
		HasSubRace: true,
	}
	// LightfootHalfling has the base halfling bonus plus the lightfoot halfling special bonus
	LightfootHalfling = SubRace{
		Name: "Lightfoot Halfling",
		AbilityScoreBonuses: npcgen.AbilityScores{
			Cha: 1,
		},
		Race: Halfling,
	}
	// StoutHalfling has the base halfling bonus plus the stout halfling special bonus
	StoutHalfling = SubRace{
		Name: "Stout Halfling",
		AbilityScoreBonuses: npcgen.AbilityScores{
			Con: 1,
		},
		Race: Halfling,
	}
)
