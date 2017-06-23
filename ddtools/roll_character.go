package ddtools

import "fmt"

func RollCharacter() {
	// Roll ability scores
	// TODO: choose method (4d6 drop low, 3d6, standard array)
	abilityScores := RollAbilityScores()
	// TODO: Choose race
	// Note: choose method (choice, roll)
	// TODO: Decide on background
	PrintCharacter(abilityScores)
}

func PrintCharacter(abilityScores []int) {
	for i, score := range abilityScores {
		switch i {
		case 0:
			fmt.Print("Str: ")
		case 1:
			fmt.Print("Dex: ")
		case 2:
			fmt.Print("Con: ")
		case 3:
			fmt.Print("Int: ")
		case 4:
			fmt.Print("Wis: ")
		case 5:
			fmt.Print("Cha: ")
		}
		fmt.Printf("%2v (%s)\n", score, modifierToString(getModifier(score)))
	}
	fmt.Println("Sum of modifiers:", sumModifiers(abilityScores))
}
