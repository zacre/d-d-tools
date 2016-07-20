package main

import (
	"d-d-tools/ddtools"
	"sort"
)

func main() {
	abilityScores := ddtools.RollAbilityScores()
	sort.Sort(sort.Reverse(sort.IntSlice(abilityScores)))
	ddtools.PrintAbilityScores(abilityScores)
}
