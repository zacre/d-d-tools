package ddtools

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func RollAbilityScores() []int {
	scores := make([]int, 6, 6)
	rolls_tmp := make([]int, 4, 4)
	rand.Seed(time.Now().Unix())
	// get 6 ability scores
	for i := 0; i < 6; i++ {
		// roll 4d6
		for j := 0; j < 4; j++ {
			// Intn(6) makes numbers from 0 to 5
			rolls_tmp[j] = rand.Intn(6) + 1
		}
		// sum 3 highest rolls to get score
		sort.Ints(rolls_tmp)
		scores[i] = rolls_tmp[1] + rolls_tmp[2] + rolls_tmp[3]
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
