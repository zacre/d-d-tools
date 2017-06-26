package character

import "github.com/kiwih/npc-gen/npcgen"

// Race provides a mechanical description of racial features
// Apparently for this project races are only their base stats, so let's roll with that
type Race struct {
	AbilityScoreBonuses npcgen.AbilityScores
}

var (
	// Human has the stat bonuses humans provide
	Human = Race{
		AbilityScoreBonuses: npcgen.AbilityScores{
			Str: 1,
			Dex: 1,
			Con: 1,
			Int: 1,
			Wis: 1,
			Cha: 1,
		},
	}
)
