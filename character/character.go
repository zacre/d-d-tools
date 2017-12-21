package character

import (
	"fmt"

	"github.com/kiwih/npc-gen/npcgen"
	"github.com/zacre/d-d-tools/stats"
)

// Character is a type to hold character data
type Character struct {
	Race          Race
	SubRace       SubRace
	Class         Class
	AbilityScores npcgen.AbilityScores
	Background    Background
}

// GetTotalAbilityScores adds racial bonuses to a character's base ability scores, getting the total ability scores
func (c *Character) GetTotalAbilityScores() npcgen.AbilityScores {
	totalBonuses := stats.AddAbilityScores(c.Race.AbilityScoreBonuses, c.SubRace.AbilityScoreBonuses)
	totalAbilityScores := stats.AddAbilityScores(c.AbilityScores, totalBonuses)
	return totalAbilityScores
}

// Print prints the details of a character
func (c *Character) Print() {
	c.PrintTitle()
	c.PrintAbilityScores()
}

// PrintTitle prints the 'title' of a character, in format Race Class (Background) - e.g. Human Fighter (Soldier)
func (c *Character) PrintTitle() {
	if c.SubRace != (SubRace{}) {
		fmt.Printf("%s %s (%s)\n", c.SubRace.Name, c.Class.Name, c.Background.Name)
	} else {
		fmt.Printf("%s %s (%s)\n", c.Race.Name, c.Class.Name, c.Background.Name)
	}
}

// PrintAbilityScores prints the total ability scores of a character, including racial bonuses
func (c *Character) PrintAbilityScores() {
	stats.PrintAbilityScores(c.GetTotalAbilityScores())
}
