package character

import "github.com/kiwih/npc-gen/npcgen"

// Character is a type to hold character data
type Character struct {
	AbilityScores npcgen.AbilityScores
}

// SetAbilityScores sets a character's ability scores to those provided
func (c *Character) SetAbilityScores(abilityScores npcgen.AbilityScores) {
	// Roll ability scores
	// TODO: choose method (4d6 drop low, 3d6, standard array)
	// TODO: Choose race
	// Note: choose method (choice, roll)
	// TODO: Decide on background
	c.AbilityScores =  abilityScores
}

// Print prints the details of a character
func (c *Character) Print() {
	PrintAbilityScores(c.AbilityScores)
}
