package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"

	"github.com/kiwih/npc-gen/npcgen"
	"github.com/zacre/d-d-tools/character"
)

// TODO: make an interactive character generator!
/*
First greet the user
Then roll some stats, and ask which stats the user wants to put them in
Then let the user choose a race, or choose a random one (TODO: suggest a recommended race based on evening up ability scores or improving the highest)
*/
func main() {
	fmt.Printf("\nWelcome to the D&D character generator! This program will generate a new character for you.\n\n")
	rawAbilityScores := rollAbilityScores()
	c := assignAbilityScores(rawAbilityScores)
	printCharacter(c)
}

func rollAbilityScores() []npcgen.AbilityScore {
	// Seed rand and roll six ability scores
	rand.Seed(time.Now().Unix())
	rawAbilityScores := character.RollAbilityScores()
	return rawAbilityScores
}

func assignAbilityScores(rawAbilityScores []npcgen.AbilityScore) character.Character {
	// Decide which method of assigning ability scores should be used
	return assignAbilityScoresByChoice(rawAbilityScores)
}

// Ability scores are assigned in order
func assignAbilityScoresStraightDown(rawAbilityScores []npcgen.AbilityScore) character.Character {
	fmt.Printf("The following numbers were rolled for your base ability scores:\n")
	fmt.Printf("%v, %v, %v, %v, %v, %v.\n", rawAbilityScores[0], rawAbilityScores[1], rawAbilityScores[2], rawAbilityScores[3], rawAbilityScores[4], rawAbilityScores[5])

	fmt.Printf("The rolled numbers will be assigned to your ability score statistics in the order they were rolled.\n")
	c := character.Character{AbilityScores: character.SimpleAssignAbilityScores(rawAbilityScores)}
	return c
}

func assignAbilityScoresByChoice(rawAbilityScores []npcgen.AbilityScore) character.Character {
	// Sort raw ability scores from highest to lowest (note reverse '>' operator for less function)
	sort.Slice(rawAbilityScores, func(i, j int) bool { return rawAbilityScores[i] > rawAbilityScores[j] })
	fmt.Printf("The following numbers were rolled for your base ability scores:\n")
	fmt.Printf("%v, %v, %v, %v, %v, %v.\n", rawAbilityScores[0], rawAbilityScores[1], rawAbilityScores[2], rawAbilityScores[3], rawAbilityScores[4], rawAbilityScores[5])

	c := character.Character{}
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
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > 6 {
			fmt.Printf("Invalid input, expecting a number between 1 and 6\n")
			continue
		}
		// If chosen stat is already assigned, tell the user to choose a different stat
		if isAssigned[choice-1] {
			fmt.Printf("You already assigned a value to %s. Choose a different stat (ability scores still to assign: ", character.Abilities[statVal-1])
			notYetAssigned := make([]string, 0, 6)
			for i := range isAssigned {
				if !isAssigned[i] {
					notYetAssigned = append(notYetAssigned, character.Abilities[i])
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
			c.AbilityScores.Str = rawAbilityScores[statVal]
		case 2:
			c.AbilityScores.Dex = rawAbilityScores[statVal]
		case 3:
			c.AbilityScores.Con = rawAbilityScores[statVal]
		case 4:
			c.AbilityScores.Int = rawAbilityScores[statVal]
		case 5:
			c.AbilityScores.Wis = rawAbilityScores[statVal]
		case 6:
			c.AbilityScores.Cha = rawAbilityScores[statVal]
		default:
			fmt.Printf("WARNING: Choice is outside the range [1-6]. This should never happen, trying again.\n")
			continue
		}
		fmt.Printf("Assigned the score %v to %v\n", rawAbilityScores[statVal], character.Abilities[choice-1])
		isAssigned[choice-1] = true
		statVal++
	}
	fmt.Println()
	return c
}

func printCharacter(c character.Character) {
	fmt.Println("Your character looks like this:")
	c.Print()
	fmt.Println("The sum of the ability score modifiers for these stats is", character.SumModifiers(c.AbilityScores))
}
