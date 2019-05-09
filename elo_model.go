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
	"math"
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
const displaySummary bool = true

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

func playerRemove(p []player, i int) []player {
	p[i] = p[len(p)-1]
	return p[:len(p)-1]
}

func stdDev(nums ...float64) (float64, float64) {
	sum := 0.00
	for _, num := range nums {
		sum += num
	}
	average := sum / float64(len(nums))
	var sD float64
	for _, num := range nums {
		sD += math.Pow((num - average), 2)
	}
	sD = math.Sqrt(sD / float64(len(nums)))
	return average, sD
}

func main() {
	// begin timing
	startTime := time.Now()
	defer func() {
		elapsedTime := time.Since(startTime)
		fmt.Println(elapsedTime)
	}()

	//create player pool
	playerPool := make([]player, totalPlayers, totalPlayers+1)
	for i := 0; i < totalPlayers; i++ {
		playerPool[i] = playerConstructor("p" + strconv.Itoa(i))
	}

	// simulate games
	for gameCounter := 0; gameCounter < totalGamesPlayed; gameCounter++ {

		// create and fill teams
		team1 := make([]player, 0, teamSize)
		for i := 0; i < teamSize; i++ {
			randomPlayerID := rand.Intn(len(playerPool))
			team1 = append(team1, playerPool[randomPlayerID])
			playerPool = playerRemove(playerPool, randomPlayerID)
		}
		team2 := make([]player, 0, teamSize)
		for i := 0; i < teamSize; i++ {
			randomPlayerID := rand.Intn(len(playerPool))
			team2 = append(team2, playerPool[randomPlayerID])
			playerPool = playerRemove(playerPool, randomPlayerID)
		}
		// calculate the team's average Elo and determine which team is more skilled
		team1AvgElo, team2AvgElo := 0.00, 0.00
		team1Skill, team2Skill := 0, 0
		for i := 0; i < teamSize; i++ {
			team1AvgElo = team1AvgElo + team1[i].elo
			team1Skill += team1[i].skill
			team2AvgElo = team2AvgElo + team2[i].elo
			team2Skill += team2[i].skill
		}
		team1AvgElo /= float64(teamSize)
		team2AvgElo /= float64(teamSize)

		// determine game winner by team's total skill
		var team1Win float64
		if team1Skill > team2Skill {
			team1Win = 1.00
		} else if team1Skill == team2Skill {
			team1Win = 0.50
		} else {
			team1Win = 0.00
		}

		// calculations
		r1 := math.Pow(10, team1AvgElo/eloSpread)
		r2 := math.Pow(10, team2AvgElo/eloSpread)

		team1WinProbability := r1 / (r1 + r2)
		team2WinProbability := r2 / (r1 + r2)

		// if team 1 wins, this creates a positive number (elo gain)
		team1EloChange := eloWeight * (team1Win - team1WinProbability)
		// if team 1 wins, this creates a negative number (elo loss)
		team2EloChange := eloWeight * (math.Abs(team1Win-1.00) - team2WinProbability)

		// update elo score and games played for each player and return to playerPool
		for i := (teamSize - 1); i >= 0; i-- {
			team1[i].elo += team1EloChange
			team1[i].gamesPlayed++
			playerPool = append(playerPool, team1[i])
			team1 = playerRemove(team1, i)
			team2[i].elo += team2EloChange
			team2[i].gamesPlayed++
			playerPool = append(playerPool, team2[i])
			team2 = playerRemove(team2, i)
		}
	}

	// organize the player pool
	sortedPlayerPool := make([]player, 0, totalPlayers)
	for j := 0; j <= maxSkillLevel; j++ {
		for k := 0; k < len(playerPool); k++ {
			if playerPool[k].skill == (maxSkillLevel - j) {
				sortedPlayerPool = append(sortedPlayerPool, playerPool[k])
				playerPool = playerRemove(playerPool, k)
				k--
			}
		}
	}

	//display the sorted player pool
	if displayAllPlayers {
		for i := 0; i < len(sortedPlayerPool); i++ {
			fmt.Println(sortedPlayerPool[i])
		}
	}

	// summarize the data - calculate relevant statistical data for each skill level
	if displaySummary {
		for j := 0; j <= maxSkillLevel; j++ {
			eloScoresThisSkillLevel := make([]float64, 0, totalPlayers)
			maxVal := 0.00
			minVal := 1000000.00
			for len(sortedPlayerPool) != 0 {
				if sortedPlayerPool[0].skill == (maxSkillLevel - j) {
					eloScoresThisSkillLevel = append(eloScoresThisSkillLevel, sortedPlayerPool[0].elo)
					maxVal = math.Max(maxVal, sortedPlayerPool[0].elo)
					minVal = math.Min(minVal, sortedPlayerPool[0].elo)
					sortedPlayerPool = sortedPlayerPool[1:]
				} else {
					avg, sD := stdDev(eloScoresThisSkillLevel...)
					fmt.Println("Skill Level:", maxSkillLevel-j, "\tNumber of Players:", len(eloScoresThisSkillLevel), "\tAverage elo:", avg, "\tStandard Deviation:", sD,
						"\tMax elo:", maxVal, "\tMinimum elo:", minVal)
					break
				}
			}
			if len(sortedPlayerPool) == 0 {
				avg, sD := stdDev(eloScoresThisSkillLevel...)
				fmt.Println("Skill Level:", maxSkillLevel-j, "\tNumber of Players:", len(eloScoresThisSkillLevel), "\tAverage elo:", avg, "\tStandard Deviation:", sD,
					"\tMax elo:", maxVal, "\tMinimum elo:", minVal)
			}
		}
	}
}
