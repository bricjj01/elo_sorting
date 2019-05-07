package main

/*
PURPOSE:
This program was created to demonstrate the Elo sorting method's ability
to estimate an individual player's skill level based on their performance
in a team game.

EXAMPLE:
If two teams play tug-of-war against each other, is it possible to tell
if every player on the winning team was just slightly stronger? or maybe
only one person was significantly stronger?

AUTHOR: John Brichetto
*/

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// game related constant declaration
const totalGamesPlayed int = 10000
const totalPlayers int = 1000
const teamSize int = 5

// player related constant declaration
const maxSkillLevel int = 10
const initialElo float64 = 1500.00

// Elo algorithm related constant declaration
const eloSpread float64 = 400.00
const eloWeight float64 = 32.00

// output related constants
const displayAllPlayers bool = false
const displaySummary bool = false

type player struct {
	name        string
	skill       int
	elo         float64
	gamesPlayed int
}

func playerConstructor(inputName string) player {
	p := player{
		name:        inputName,
		skill:       rand.Intn(maxSkillLevel + 1),
		elo:         initialElo,
		gamesPlayed: 0,
	}
	return p
}

func main() {
	// begin timing
	startTime := time.Now()
	defer func() {
		elapsedTime := time.Since(startTime)
		fmt.Println(elapsedTime)
	}()

	//create player pool
	playerPool := make([]player, totalPlayers, totalPlayers)
	for i := 0; i < totalPlayers; i++ {
		playerPool[i] = playerConstructor("p" + strconv.Itoa(i))
		fmt.Println(playerPool[i])
	}

	/*
		// create player pool
		playerPool := make(map[int]player, totalPlayers)
		for i := 0; i < totalPlayers; i++ {
			playerPool[i] = playerConstructor("p" + strconv.Itoa(i))
		}
	*/

	// copy delete and append from the map
	// check if key is present with _, prs := m["k2"] syntax

	// create and fill teams
	team1 := make([]player, teamSize, teamSize)
	team2 := make([]player, teamSize, teamSize)

	fmt.Println(team1, len(team1), cap(team1))
	fmt.Println(team2, len(team2), cap(team2))
}
