package main

import "github.com/zacre/d-d-tools/ddtools"

func main() {
	/*
		// Roll six ability scores
		abilityScores := ddtools.RollAbilityScores()
		// Sort from highest to lowest
		sort.Sort(sort.Reverse(sort.IntSlice(abilityScores)))
		// Print ability scores and modifiers
		ddtools.PrintAbilityScores(abilityScores)
	*/
	ddtools.RollCharacter()
}
