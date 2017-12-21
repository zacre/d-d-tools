package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/kiwih/npc-gen/npcgen"
	"github.com/zacre/d-d-tools/character"
)

// TODO: random races
// TODO: suggest a recommended race based on evening up ability scores or improving the highest
// TODO: proper background and class support
func main() {
	createCharacter()
}

func createCharacter() {
	fmt.Printf("\nWelcome to the D&D character generator! This program will generate a new character for you.\n\n")
	fmt.Printf("First, what method of creating a character would you like to use?\n")
	fmt.Print(`1. I already have a concept I would like to develop
2. I want to determine my character from what stats I get
`)
	var choice int
	err := errors.New("null")
	for err != nil {
		input := ""
		_, err = fmt.Scanf("%s\n", &input)
		if err != nil {
			fmt.Printf("Invalid input, expecting a number between 1 and 2\n")
			continue
		}
		choice, err = strconv.Atoi(string([]rune(input)[0]))
		if err != nil || choice < 1 || choice > 2 {
			fmt.Printf("Invalid input, expecting a number between 1 and 2\n")
			continue
		}
	}
	fmt.Println()

	var race character.Race
	var subrace character.SubRace
	var class character.Class
	var background character.Background
	var abilityScores npcgen.AbilityScores
	switch choice {
	case 1:
		// Character creation as shown in PhB (TODO) race -> class -> ability scores -> background
		race, subrace = getCharacterRace()
		class = getCharacterClass()
		abilityScores = getCharacterAbilityScores()
		background = getCharacterBackground()
	case 2:
		// Character creation starting from stats -> race -> background -> class
		abilityScores = getCharacterAbilityScores()
		race, subrace = getCharacterRace()
		background = getCharacterBackground()
		class = getCharacterClass()
	}
	c := character.Character{
		Race:          race,
		SubRace:       subrace,
		Class:         class,
		AbilityScores: abilityScores,
		Background:    background,
	}

	fmt.Println()
	printCharacter(c)
}

func getCharacterAbilityScores() npcgen.AbilityScores {
	var rawAbilityScores []npcgen.AbilityScore
	fmt.Println("What method would you like to use to determine your ability scores?")
	fmt.Print(
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
	// Choose method (4d6 drop low, 3d6, standard array, etc.)
	var choice int
	err := errors.New("null")
	for err != nil {
		input := ""
		_, err = fmt.Scanf("%s\n", &input)
		if err != nil {
			fmt.Printf("Invalid input, expecting a number between 1 and 8\n")
			continue
		}
		choice, err = strconv.Atoi(string([]rune(input)[0]))
		if err != nil || choice < 1 || choice > 8 {
			fmt.Printf("Invalid input, expecting a number between 1 and 8\n")
			continue
		}
	}
	fmt.Println()

	// This determines whether ability scores are assigned "straight" (i.e. in the order they were rolled) or by user choice
	assignStraight := false

	// Determine base ability scores (from rolling, point buy, etc.)
	switch choice {
	case 1:
		rawAbilityScores = rollAbilityScores(4)
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
		rawAbilityScores = rollAbilityScores(4)
		assignStraight = true
	case 8:
		rawAbilityScores = rollAbilityScores(3)
		assignStraight = true
	}

	var abilityScores npcgen.AbilityScores
	if assignStraight {
		abilityScores = assignAbilityScoresStraightDown(rawAbilityScores)
	} else {
		abilityScores = assignAbilityScoresByChoice(rawAbilityScores)
	}
	return abilityScores
}

// Input determines how many dice are rolled. With 0 arguments, defaults to 4, but with 1 argument, uses the input number instead
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
		threeAbilityScores[i] = rollAbilityScores(4)
		sum := character.SumAbilityScoresRaw(threeAbilityScores[i])
		if sum > largestSum {
			largestSum = sum
			largestIndex = i
		}
	}
	return threeAbilityScores[largestIndex]
}

// Ability scores are assigned in order
func assignAbilityScoresStraightDown(rawAbilityScores []npcgen.AbilityScore) npcgen.AbilityScores {
	fmt.Printf("The following numbers were obtained for your base ability scores:\n")
	fmt.Printf("%v, %v, %v, %v, %v, %v.", rawAbilityScores[0], rawAbilityScores[1], rawAbilityScores[2], rawAbilityScores[3], rawAbilityScores[4], rawAbilityScores[5])
	fmt.Printf(" (Sum: %v)", rawAbilityScores[0]+rawAbilityScores[1]+rawAbilityScores[2]+rawAbilityScores[3]+rawAbilityScores[4]+rawAbilityScores[5])
	fmt.Println()

	fmt.Printf("The rolled numbers will be assigned to your ability score statistics in the order they were rolled.\n")
	return character.SimpleAssignAbilityScores(rawAbilityScores)
}

func assignAbilityScoresByChoice(rawAbilityScores []npcgen.AbilityScore) npcgen.AbilityScores {
	// Sort raw ability scores from highest to lowest (note reverse '>' operator for less function)
	sort.Slice(rawAbilityScores, func(i, j int) bool { return rawAbilityScores[i] > rawAbilityScores[j] })
	fmt.Printf("The following numbers were obtained for your base ability scores:\n")
	fmt.Printf("%v, %v, %v, %v, %v, %v.", rawAbilityScores[0], rawAbilityScores[1], rawAbilityScores[2], rawAbilityScores[3], rawAbilityScores[4], rawAbilityScores[5])
	fmt.Printf(" (Sum: %v)", rawAbilityScores[0]+rawAbilityScores[1]+rawAbilityScores[2]+rawAbilityScores[3]+rawAbilityScores[4]+rawAbilityScores[5])
	fmt.Println()

	var as npcgen.AbilityScores
	fmt.Printf("You must now choose which ability score each value should be assigned to.\n")
	// fmt.Println("Enter 1 for Str, 2 for Dex, 3 for Con, 4 for Int, 5 for Wis, or 6 for Cha")
	var isAssigned [6]bool
	for statVal := 0; statVal < 6; statVal++ {
		fmt.Printf("Which ability score would you like to be %v? (1. Str  2. Dex  3. Con  4. Int  5. Wis  6. Cha): ", rawAbilityScores[statVal])
	abilityscorechoice:
		// Get user input
		var choice int
		err := errors.New("null")
		for err != nil {
			input := ""
			_, err = fmt.Scanf("%s\n", &input)
			if err != nil {
				fmt.Printf("Invalid input, expecting a number between 1 and 6\n")
				continue
			}
			choice, err = strconv.Atoi(string([]rune(input)[0]))
			if err != nil || choice < 1 || choice > 6 {
				fmt.Printf("Invalid input, expecting a number between 1 and 6\n")
				continue
			}
		}
		// If chosen stat is already assigned, tell the user to choose a different stat
		if isAssigned[choice-1] {
			fmt.Printf("You already assigned a value to %s. Choose a different stat (ability scores still to assign: ", character.AbilityNames[choice-1])
			notYetAssigned := make([]string, 0, 6)
			notYetAssignedNum := make([]int, 0, 6)
			for i := range isAssigned {
				if !isAssigned[i] {
					notYetAssigned = append(notYetAssigned, character.AbilityNames[i])
					notYetAssignedNum = append(notYetAssignedNum, i)
				}
			}
			for i, statName := range notYetAssigned {
				fmt.Printf("%d: %s", notYetAssignedNum[i]+1, statName)
				if i < len(notYetAssigned)-1 {
					fmt.Printf(", ")
				}
			}
			fmt.Print("): ")
			goto abilityscorechoice
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
	}
	fmt.Printf("Your base ability scores look like: Str %d, Dex %d, Con %d, Int %d, Wis %d, Cha %d\n", as.Str, as.Dex, as.Con, as.Int, as.Wis, as.Cha)
	fmt.Println()
	return as
}

func getCharacterRace() (character.Race, character.SubRace) {
	fmt.Println("Choose a race:")
	fmt.Print(`1. Dwarf (+2 Con)
2. Elf (+2 Dex)
3. Halfling (+2 Dex, small)
4. Human (+1 to all stats)
`)
	var choice int
	err := errors.New("null")
	for err != nil {
		input := ""
		_, err = fmt.Scanf("%s\n", &input)
		if err != nil {
			fmt.Printf("Invalid input, expecting a number between 1 and 4\n")
			continue
		}
		choice, err = strconv.Atoi(string([]rune(input)[0]))
		if err != nil || choice < 1 || choice > 4 {
			fmt.Printf("Invalid input, expecting a number between 1 and 4\n")
			continue
		}
	}

	var race character.Race
	var subrace character.SubRace
	switch choice {
	case 1:
		race = character.Dwarf
		fmt.Println(race.Name, "has subraces! Please pick a subrace as well:")
		fmt.Print(`1. Hill Dwarf (+1 Wis)
2. Mountain Dwarf (+2 Str)
`)
		err := errors.New("null")
		for err != nil {
			input := ""
			_, err = fmt.Scanf("%s\n", &input)
			if err != nil {
				fmt.Printf("Invalid input, expecting a number between 1 and 2\n")
				continue
			}
			choice, err = strconv.Atoi(string([]rune(input)[0]))
			if err != nil || choice < 1 || choice > 2 {
				fmt.Printf("Invalid input, expecting a number between 1 and 2\n")
				continue
			}
		}

		switch choice {
		case 1:
			subrace = character.HillDwarf
		case 2:
			subrace = character.MountainDwarf
		}
	case 2:
		race = character.Elf
		fmt.Println(race.Name, "has subraces! Please pick a subrace as well:")
		fmt.Print(`1. High Elf (+1 Int)
2. Wood Elf (+1 Wis)
3. Dark Elf (+1 Cha, sunlight sensitivity)
`)
		err := errors.New("null")
		for err != nil {
			input := ""
			_, err = fmt.Scanf("%s\n", &input)
			if err != nil {
				fmt.Printf("Invalid input, expecting a number between 1 and 3\n")
				continue
			}
			choice, err = strconv.Atoi(string([]rune(input)[0]))
			if err != nil || choice < 1 || choice > 3 {
				fmt.Printf("Invalid input, expecting a number between 1 and 3\n")
				continue
			}
		}

		switch choice {
		case 1:
			subrace = character.HighElf
		case 2:
			subrace = character.WoodElf
		case 3:
			subrace = character.DarkElf
		}
	case 3:
		race = character.Halfling
		fmt.Println(race.Name, "has subraces! Please pick a subrace as well:")
		fmt.Print(`1. Lightfoot Halfling (+1 Cha)
2. Stout Halfling (+1 Con)
`)
		err := errors.New("null")
		for err != nil {
			input := ""
			_, err = fmt.Scanf("%s\n", &input)
			if err != nil {
				fmt.Printf("Invalid input, expecting a number between 1 and 2\n")
				continue
			}
			choice, err = strconv.Atoi(string([]rune(input)[0]))
			if err != nil || choice < 1 || choice > 2 {
				fmt.Printf("Invalid input, expecting a number between 1 and 2\n")
				continue
			}
		}

		switch choice {
		case 1:
			subrace = character.LightfootHalfling
		case 2:
			subrace = character.StoutHalfling
		}
	case 4:
		race = character.Human
	}

	if subrace != (character.SubRace{}) {
		fmt.Println("You chose a", subrace.Name)
	} else {
		fmt.Println("You chose a", race.Name)
	}
	fmt.Println()
	return race, subrace
}

func getCharacterClass() character.Class {
	fmt.Println("Enter your class:")
	input := ""
	err := errors.New("null")
	for err != nil {
		_, err = fmt.Scanf("%s\n", &input)
		if err != nil {
			fmt.Printf("Scanf error\n")
		}
	}
	fmt.Println()
	var class character.Class
	class.Name = input
	return class
}

func getCharacterBackground() character.Background {
	fmt.Println("Enter your background:")
	input := ""
	err := errors.New("null")
	for err != nil {
		_, err = fmt.Scanf("%s\n", &input)
		if err != nil {
			fmt.Printf("Scanf error\n")
		}
	}
	fmt.Println()
	var background character.Background
	background.Name = input
	return background
}

func printCharacter(c character.Character) {
	fmt.Println("Your character looks like this:")
	c.Print()
	tmp := character.SumModifiers(c.GetTotalAbilityScores())
	fmt.Println("The sum of the ability score modifiers for these stats is", tmp)
	decision := ""
	switch {
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
