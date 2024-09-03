package main

import "fmt"

func main() {
	fmt.Println("Hellow World")
}

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.
	currentBestTeam := ""
	scoreMap := make(map[string]int)
	for index, team := range competitions {
		result := results[index]
		homeTeam, awayTeam := team[0], team[1]

		winningTeam := awayTeam
		if result == HOME_TEAM_WON {
			winningTeam = homeTeam
		}
		scoreMap[winningTeam] += 3

		if scoreMap[winningTeam] > scoreMap[currentBestTeam] {
			currentBestTeam = winningTeam
		}
	}
	return currentBestTeam
}
