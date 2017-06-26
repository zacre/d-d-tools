package character

import "github.com/kiwih/npc-gen/npcgen"

// Character is a type to hold character data
type Character struct {
	AbilityScores npcgen.AbilityScores
}

// Create creates a new character
func Create(abilityScores npcgen.AbilityScores) Character {
	// Roll ability scores
	// TODO: choose method (4d6 drop low, 3d6, standard array)
	// TODO: Choose race
	// Note: choose method (choice, roll)
	// TODO: Decide on background
	c := Character{AbilityScores: abilityScores}

	return c
}

// Print prints the details of a character
func (c *Character) Print() {
	PrintAbilityScores(c.AbilityScores)
}
