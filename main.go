package main

import (
  "sort"

	"github.com/zacre/d-d-tools/ddtools"
)

func main() {
	abilityScores := ddtools.RollAbilityScores()
	sort.Sort(sort.Reverse(sort.IntSlice(abilityScores)))
	ddtools.PrintAbilityScores(abilityScores)
}
