package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/kiwih/npc-gen/npcgen"
	"github.com/zacre/d-d-tools/character"
)

/*
First greet the user
Then roll some stats, and ask which stats the user wants to put them in
Then let the user choose a race, or choose a random one (TODO: suggest a recommended race based on evening up ability scores or improving the highest)
*/
func main() {
	fmt.Printf("\nWelcome to the D&D character generator! This program will generate a new character for you.\n\n")
	var rawAbilityScores []npcgen.AbilityScore
	fmt.Printf("What method would you like to use to determine your ability scores?\n")
	fmt.Printf(
		`1. Rolling 4d6 drop the lowest (Sum Avg: 73, SD: 7)
2. Standard array (15, 14, 13, 12, 10, 8. Sum: 72)
3. Best of 3: Rolling 4d6 drop the lowest (Sum Avg: 79, SD: 5)
4. Heroic array (16, 15, 14, 12, 11, 10. Sum: 78)
5. Rolling 5d6 drop the two lowest (Sum Avg: 81, SD: 6)
6. Rolling 3d6 (Sum Avg: 63, SD: 7)
7. Rolling 4d6 drop the lowest straight down
8. Rolling 3d6 straight down
`)
	fmt.Printf("Enter your choice: ")
	input := ""
	_, err := fmt.Scanf("%s\n", &input)
	if err != nil {
		fmt.Printf("Invalid input, expecting a number between 1 and 8\n")
		return
	}
	choice, err := strconv.Atoi(string([]rune(input)[0]))
	if err != nil || choice < 1 || choice > 8 {
		fmt.Printf("Invalid input, expecting a number between 1 and 8\n")
		return
	}
	assignStraight := false
	switch choice {
	case 1:
		rawAbilityScores = rollAbilityScores()
	case 2:
		rawAbilityScores = []npcgen.AbilityScore{15, 14, 13, 12, 10, 8}
	case 3:
		rawAbilityScores = rollAbilityScoresBestOf3()
	case 4:
		rawAbilityScores = []npcgen.AbilityScore{16, 15, 14, 12, 11, 10}
	case 5:
		rawAbilityScores = rollAbilityScores(5)
	case 6:
		rawAbilityScores = rollAbilityScores(3)
	case 7:
		rawAbilityScores = rollAbilityScores()
		assignStraight = true
	case 8:
		rawAbilityScores = rollAbilityScores(3)
		assignStraight = true
	default:
		rawAbilityScores = nil
	}
	if rawAbilityScores == nil {
		fmt.Printf("An error has occurred somehow; your choice of determining ability scores was %d, but this program doesn't recognise that\n", choice)
		return
	}
	fmt.Println()
	c := createCharacter(rawAbilityScores, assignStraight)
	printCharacter(c)
}

func rollAbilityScores(dice ...int) []npcgen.AbilityScore {
	// Seed rand and roll six ability scores
	rand.Seed(time.Now().Unix())
	randSeeded := true
	diceToRoll := 4
	for i, die := range dice {
		// Only first option is investigated
		if i == 0 {
			diceToRoll = die
		} else {
			break
		}
	}
	rawAbilityScores, err := character.RollAbilityScores(diceToRoll, randSeeded)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	return rawAbilityScores
}

func rollAbilityScoresBestOf3() []npcgen.AbilityScore {
	var threeAbilityScores [3][]npcgen.AbilityScore
	largestSum := 0
	largestIndex := 0
	for i := 0; i < 3; i++ {
		threeAbilityScores[i] = rollAbilityScores()
		sum := character.SumAbilityScoresRaw(threeAbilityScores[i])
		if sum > largestSum {
			largestSum = sum
			largestIndex = i
		}
	}
	return threeAbilityScores[largestIndex]
}

// TODO: Currently only creates a set of ability scores, not a character
func createCharacter(rawAbilityScores []npcgen.AbilityScore, assignStraight bool) character.Character {
	c := character.Character{}
	if assignStraight {
		assignAbilityScoresStraightDown(&c, rawAbilityScores)
	} else {
		assignAbilityScoresByChoice(&c, rawAbilityScores)
	}
	return c
}

// Ability scores are assigned in order
func assignAbilityScoresStraightDown(c *character.Character, rawAbilityScores []npcgen.AbilityScore) {
	fmt.Printf("The following numbers were obtained for your base ability scores:\n")
	fmt.Printf("%v, %v, %v, %v, %v, %v.", rawAbilityScores[0], rawAbilityScores[1], rawAbilityScores[2], rawAbilityScores[3], rawAbilityScores[4], rawAbilityScores[5])
	fmt.Printf(" (Sum: %v)", rawAbilityScores[0]+rawAbilityScores[1]+rawAbilityScores[2]+rawAbilityScores[3]+rawAbilityScores[4]+rawAbilityScores[5])
	fmt.Println()

	fmt.Printf("The rolled numbers will be assigned to your ability score statistics in the order they were rolled.\n")
	c.SetAbilityScores(character.SimpleAssignAbilityScores(rawAbilityScores))
}

func assignAbilityScoresByChoice(c *character.Character, rawAbilityScores []npcgen.AbilityScore) {
	// Sort raw ability scores from highest to lowest (note reverse '>' operator for less function)
	sort.Slice(rawAbilityScores, func(i, j int) bool { return rawAbilityScores[i] > rawAbilityScores[j] })
	fmt.Printf("The following numbers were obtained for your base ability scores:\n")
	fmt.Printf("%v, %v, %v, %v, %v, %v.", rawAbilityScores[0], rawAbilityScores[1], rawAbilityScores[2], rawAbilityScores[3], rawAbilityScores[4], rawAbilityScores[5])
	fmt.Printf(" (Sum: %v)", rawAbilityScores[0]+rawAbilityScores[1]+rawAbilityScores[2]+rawAbilityScores[3]+rawAbilityScores[4]+rawAbilityScores[5])
	fmt.Println()

	as := npcgen.AbilityScores{}
	fmt.Printf("You must now choose which ability score each value should be assigned to.\n")
	// fmt.Println("Enter 1 for Str, 2 for Dex, 3 for Con, 4 for Int, 5 for Wis, or 6 for Cha")
	var isAssigned [6]bool
	statVal := 0
	for statVal < 6 {
		fmt.Printf("Which ability score would you like to be %v? (1. Str  2. Dex  3. Con  4. Int  5. Wis  6. Cha): ", rawAbilityScores[statVal])
		// Get user input
		input := ""
		_, err := fmt.Scanf("%s\n", &input)
		if err != nil {
			fmt.Printf("Invalid input, expecting a number between 1 and 6\n")
			continue
		}
		// Convert input to integer
		choice, err := strconv.Atoi(string([]rune(input)[0]))
		if err != nil || choice < 1 || choice > 6 {
			fmt.Printf("Invalid input, expecting a number between 1 and 6\n")
			continue
		}
		// If chosen stat is already assigned, tell the user to choose a different stat
		if isAssigned[choice-1] {
			fmt.Printf("You already assigned a value to %s. Choose a different stat (ability scores still to assign: ", character.AbilityNames[choice-1])
			notYetAssigned := make([]string, 0, 6)
			for i := range isAssigned {
				if !isAssigned[i] {
					notYetAssigned = append(notYetAssigned, character.AbilityNames[i])
				}
			}
			for i, statName := range notYetAssigned {
				fmt.Printf("%s", statName)
				if i < len(notYetAssigned)-1 {
					fmt.Printf(", ")
				}
			}
			fmt.Printf(")\n")
			continue
		}
		// Assign the correct ability score
		switch choice {
		case 1:
			as.Str = rawAbilityScores[statVal]
		case 2:
			as.Dex = rawAbilityScores[statVal]
		case 3:
			as.Con = rawAbilityScores[statVal]
		case 4:
			as.Int = rawAbilityScores[statVal]
		case 5:
			as.Wis = rawAbilityScores[statVal]
		case 6:
			as.Cha = rawAbilityScores[statVal]
		default:
			fmt.Printf("WARNING: Choice is outside the range [1-6]. This should never happen, trying again.\n")
			continue
		}
		fmt.Printf("Assigned the score %v to %v\n", rawAbilityScores[statVal], character.AbilityNames[choice-1])
		isAssigned[choice-1] = true
		statVal++
	}
	fmt.Println()
	c.SetAbilityScores(as)
}

func printCharacter(c character.Character) {
	fmt.Println("Your character looks like this:")
	c.Print()
	fmt.Println("The sum of the ability score modifiers for these stats is", character.SumModifiers(c.AbilityScores))
	decision := ""
	switch tmp := character.SumModifiers(c.AbilityScores); {
	case tmp <= 3:
		decision = "Oh, those stats are a bit below average."
	case tmp >= 4 && tmp <= 6:
		decision = "Those stats are quite balanced overall."
	case tmp >= 7 && tmp <= 9:
		decision = "Nice job, those stats are pretty high!"
	case tmp >= 10:
		decision = "Congratulations, those are phenomenal stats!"
	}
	fmt.Println(decision)
}
