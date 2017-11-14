package character

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/kiwih/npc-gen/npcgen"
)

var Abilities = []string{"Str", "Dex", "Con", "Int", "Wis", "Cha"}

// RollAbilityScores rolls for six ability scores using the 4d6 drop lowest method. It returns a slice of six ability scores in the order they were generated.
// Requires rand to be seeded before use
func RollAbilityScores() []npcgen.AbilityScore {
	scores := make([]npcgen.AbilityScore, 6, 6)
	var rolls [4]int
	// get 6 ability scores
	for i := range scores {
		// roll 4d6
		for j := range rolls {
			// Intn(6) makes numbers from 0 to 5
			rolls[j] = rand.Intn(6) + 1
		}
		// get index of lowest val
		lowest := 0
		for j := range rolls {
			if rolls[j] < rolls[lowest] {
				lowest = j
			}
		}
		// sum 3 highest rolls to get score
		for j := range rolls {
			if j != lowest {
				scores[i] += npcgen.AbilityScore(rolls[j])
			}
		}
	}
	return scores
}

// SimpleAssignAbilityScores creates an AbilityScores struct from a slice of six ability score values
func SimpleAssignAbilityScores(baseScores []npcgen.AbilityScore) npcgen.AbilityScores {
	var abilityScores npcgen.AbilityScores
	abilityScores.Str = npcgen.AbilityScore(baseScores[0])
	abilityScores.Dex = npcgen.AbilityScore(baseScores[1])
	abilityScores.Con = npcgen.AbilityScore(baseScores[2])
	abilityScores.Int = npcgen.AbilityScore(baseScores[3])
	abilityScores.Wis = npcgen.AbilityScore(baseScores[4])
	abilityScores.Cha = npcgen.AbilityScore(baseScores[5])
	return abilityScores
}

// PrintRawAbilityScores prints out a slice of ability scores
func PrintRawAbilityScores(scores []npcgen.AbilityScore) {
	for _, score := range scores {
		fmt.Printf("%v:\t%s\n", score, modifierToString(GetModifier(score)))
	}
}

// PrintAbilityScores prints an AbilityScores struct
func PrintAbilityScores(scores npcgen.AbilityScores) {
	fmt.Printf("Str: %2v (%s)\n", scores.Str, modifierToString(GetModifier(scores.Str)))
	fmt.Printf("Dex: %2v (%s)\n", scores.Dex, modifierToString(GetModifier(scores.Dex)))
	fmt.Printf("Con: %2v (%s)\n", scores.Con, modifierToString(GetModifier(scores.Con)))
	fmt.Printf("Int: %2v (%s)\n", scores.Int, modifierToString(GetModifier(scores.Int)))
	fmt.Printf("Wis: %2v (%s)\n", scores.Wis, modifierToString(GetModifier(scores.Wis)))
	fmt.Printf("Cha: %2v (%s)\n", scores.Cha, modifierToString(GetModifier(scores.Cha)))
}

func GetModifier(score npcgen.AbilityScore) int {
	modifier := int(score/2 - 5)
	return modifier
}

func modifierToString(modifier int) string {
	modifierString := strconv.Itoa(modifier)
	prepend := ""
	if modifier > 0 {
		prepend = "+"
	} else if modifier == 0 {
		prepend = " "
	}
	return prepend + modifierString
}

func SumAbilityScoresRaw(scores []npcgen.AbilityScore) int {
	return SumAbilityScores(SimpleAssignAbilityScores(scores))
}

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

func SumModifiersRaw(scores []npcgen.AbilityScore) int {
	return SumModifiers(SimpleAssignAbilityScores(scores))
}

func SumModifiers(scores npcgen.AbilityScores) int {
	sum := 0
	sum += GetModifier(scores.Str)
	sum += GetModifier(scores.Dex)
	sum += GetModifier(scores.Con)
	sum += GetModifier(scores.Int)
	sum += GetModifier(scores.Wis)
	sum += GetModifier(scores.Cha)
	return sum
}
