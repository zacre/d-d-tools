package ddtools

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func RollAbilityScores() []int {
	scores := make([]int, 6, 6)
	var rolls [4]int
	rand.Seed(time.Now().Unix())
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
				scores[i] += rolls[j]
			}
		}
	}
	return scores
}

func PrintAbilityScores(scores []int) {
	for _, score := range scores {
		fmt.Printf("%v:\t%s\n", score, modifierToString(getModifier(score)))
	}
	fmt.Println("Sum of modifiers:", sumModifiers(scores))
}

func getModifier(score int) int {
	modifier := score/2 - 5
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

func sumModifiers(scores []int) int {
	sum := 0
	for _, score := range scores {
		sum += getModifier(score)
	}
	return sum
}
