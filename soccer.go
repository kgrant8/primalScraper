package main

import (
	"database/sql"
	"log"
	"soccer/data"
	"soccer/models"
	"time"
)

const (
	SITE_URL = "http://www.primalsoccer.com/leagues?state=LIVE&locationId=&seasonId=&days=&levelId="
)

func main() {
	allSeasons := make([]models.Game, 0)
	allLeagues := make([]models.League, 0)
	currentTime := time.Now()

	db, err := sql.Open("mysql", "root:keith8@tcp(127.0.0.1:3306)/primal")
	if err != nil {
		log.Fatal(err)
	}
	allLeagues = data.GetLeagueData(SITE_URL)
	for _, league := range allLeagues {
		newSeason := data.GetScheduleData(league.Url, league.Id)

		for _, game := range newSeason {
			if len(game.AwayTeam) != 0 {
				game.AwayTeamId = data.GetTeamId(db, game.AwayTeam, league.Id)
			}
			if len(game.HomeTeam) != 0 {
				game.HomeTeamId = data.GetTeamId(db, game.HomeTeam, league.Id)
			}
			//getting number of days since game as been played
			diff := currentTime.YearDay() - game.Time.YearDay()

			if game.Played && diff == 1 && game.AwayGoals != 0 && game.HomeGoals != 0 {
				data.UpdateStandings(db, game)
				data.UpdateGameScore(db, game)
			}
		}
		allSeasons = append(allSeasons, newSeason...)
	}

	defer db.Close()

}
