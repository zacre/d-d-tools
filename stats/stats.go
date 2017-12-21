package stats

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/kiwih/npc-gen/npcgen"
)

// AbilityNames is a helper array listing the names of the abilities in modern order
var AbilityNames = []string{"Str", "Dex", "Con", "Int", "Wis", "Cha"}

// RollAbilityScores rolls for six ability scores using the 4d6 drop lowest method. It returns a slice of six ability scores in the order they were generated.
// Requires rand to be seeded before use
func RollAbilityScores(dice int, randSeeded bool) ([]npcgen.AbilityScore, error) {
	if !randSeeded {
		return nil, errors.New("Please seed rand for RollAbilityScores to work (e.g. `rand.Seed(time.Now().Unix())`")
	}
	rolls := make([]int, dice)
	scores := make([]npcgen.AbilityScore, 6, 6)
	// get 6 ability scores
	for i := range scores {
		// roll Nd6 where N == var dice
		for j := range rolls {
			// Intn(6) makes numbers from 0 to 5
			rolls[j] = rand.Intn(6) + 1
		}
		if dice <= 3 {
			// sum all dice to get score
			for j := range rolls {
				scores[i] += npcgen.AbilityScore(rolls[j])
			}
		} else {
			// sum 3 highest rolls to get score
			// NOTE: this assumption may need to change
			var highest [3]int
			for j := range rolls {
				if rolls[j] >= rolls[highest[0]] {
					// If this number is higher than the value of highest[0], stick it on top and shift everything down
					highest[2] = highest[1]
					highest[1] = highest[0]
					highest[0] = j
				} else if rolls[j] <= rolls[highest[0]] && rolls[j] >= rolls[highest[1]] {
					// If this number is lower than the value of highest[0] but higher than the value of highest[1], stick it at highest[1] and shift everything below down
					highest[2] = highest[1]
					highest[1] = j
				} else if rolls[j] <= rolls[highest[1]] && rolls[j] >= rolls[highest[2]] {
					// If this number is lower than the value of highest[0] and highest[1] but higher than the value of highest[2], stick it at highest[2] and shift everything below down
					highest[2] = j
				} else if highest[0] == highest[1] {
					// If highest[0] and highest[1] are the same index number (e.g. both initialised to 0) put this number in the array anyway (to prevent final array having duplicate references)
					// Make it both highest[1] and highest[2] so the next clause will trigger next loop
					highest[2] = j
					highest[1] = j
				} else if highest[1] == highest[2] {
					// If highest[1] and highest[2] are the same index number (e.g. both initialised to 0) put this number in the array anyway (to prevent final array having duplicate references)
					highest[2] = j
				}
			}
			scores[i] += npcgen.AbilityScore(rolls[highest[0]] + rolls[highest[1]] + rolls[highest[2]])
		}
	}
	return scores, nil
}

// SimpleAssignAbilityScores creates an AbilityScores struct from a slice of six ability score values, taking the values in order
func AbilityScores(baseScores []npcgen.AbilityScore) npcgen.AbilityScores {
	var abilityScores npcgen.AbilityScores
	abilityScores.Str = npcgen.AbilityScore(baseScores[0])
	abilityScores.Dex = npcgen.AbilityScore(baseScores[1])
	abilityScores.Con = npcgen.AbilityScore(baseScores[2])
	abilityScores.Int = npcgen.AbilityScore(baseScores[3])
	abilityScores.Wis = npcgen.AbilityScore(baseScores[4])
	abilityScores.Cha = npcgen.AbilityScore(baseScores[5])
	return abilityScores
}

func modifierToString(modifier int) string {
	modifierString := strconv.Itoa(modifier)
	prepend := ""
	if modifier > 0 {
		prepend = "+"
	} else if modifier == 0 {
		prepend = " "
	} // '-' already included from Itoa
	return prepend + modifierString
}

// AddAbilityScores adds two sets of AbilityScore structs together
func AddAbilityScores(scores npcgen.AbilityScores, scores2 npcgen.AbilityScores) npcgen.AbilityScores {
	var result npcgen.AbilityScores
	result.Str = scores.Str + scores2.Str
	result.Dex = scores.Dex + scores2.Dex
	result.Con = scores.Con + scores2.Con
	result.Int = scores.Int + scores2.Int
	result.Wis = scores.Wis + scores2.Wis
	result.Cha = scores.Cha + scores2.Cha
	return result
}

// SumAbilityScoresRaw calculates the sum of six ability scores in an array of AbilityScore
func SumAbilityScoresRaw(scores []npcgen.AbilityScore) int {
	return SumAbilityScores(AbilityScores(scores))
}

// SumAbilityScores calculates the sum of the six scores in an AbilityScores struct
func SumAbilityScores(scores npcgen.AbilityScores) int {
	sum := 0
	sum += int(scores.Str)
	sum += int(scores.Dex)
	sum += int(scores.Con)
	sum += int(scores.Int)
	sum += int(scores.Wis)
	sum += int(scores.Cha)
	return sum
}

// SumModifiersRaw calculates the sum of the six ability score modifiers calculated from an array of AbilityScore
func SumModifiersRaw(scores []npcgen.AbilityScore) int {
	return SumModifiers(AbilityScores(scores))
}

// SumModifiers calculates the sum of the six ability score modifiers calculated from an AbilityScores struct
func SumModifiers(scores npcgen.AbilityScores) int {
	sum := 0
	sum += scores.Str.Modifier()
	sum += scores.Dex.Modifier()
	sum += scores.Con.Modifier()
	sum += scores.Int.Modifier()
	sum += scores.Wis.Modifier()
	sum += scores.Cha.Modifier()
	return sum
}

// PrintRawAbilityScores prints out a slice of ability scores
func PrintRawAbilityScores(scores []npcgen.AbilityScore) {
	for _, score := range scores {
		fmt.Printf("%v:\t%s\n", score, modifierToString(score.Modifier()))
	}
}

// PrintAbilityScores prints an AbilityScores struct
func PrintAbilityScores(scores npcgen.AbilityScores) {
	fmt.Printf("Str: %2v (%s)\n", scores.Str, modifierToString(scores.Str.Modifier()))
	fmt.Printf("Dex: %2v (%s)\n", scores.Dex, modifierToString(scores.Dex.Modifier()))
	fmt.Printf("Con: %2v (%s)\n", scores.Con, modifierToString(scores.Con.Modifier()))
	fmt.Printf("Int: %2v (%s)\n", scores.Int, modifierToString(scores.Int.Modifier()))
	fmt.Printf("Wis: %2v (%s)\n", scores.Wis, modifierToString(scores.Wis.Modifier()))
	fmt.Printf("Cha: %2v (%s)\n", scores.Cha, modifierToString(scores.Cha.Modifier()))
}
