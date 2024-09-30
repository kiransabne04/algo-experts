package main

import "fmt"

func main() {
	competitions := [][]string{{"HTML", "C#"}, {"C#", "Python"}, {"Python", "HTML"}}
	results := []int{0,0,1}
	
	val := TournamentWinner(competitions, results)
	fmt.Println("Hellow World", competitions, results, val)
}

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.
	bestTeam := ""
	scoreMap := make(map[string]int)

	for i, team := range competitions {
		homeTeam, awayTeam := team[0], team[1]

		winningTeam := awayTeam
		if results[i] == 1 {
			winningTeam = homeTeam
		}
		scoreMap[winningTeam] +=3
	}
	maxPoints := 0
	for team, val := range scoreMap{
		if val > maxPoints {
			maxPoints = val
			bestTeam = team
		}
	}

	fmt.Println(scoreMap)
	return bestTeam
}
