package data

import (
	"database/sql"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"soccer/models"
)

const (
	SQLINSTERTSTANDINGS      = "INSERT INTO `standings` (`teamId`, `leagueId`, `gamesPlayed`, `wins`, `loses`, `ties`,`goalDiff`) VALUES (?,?,?,?,?,?,?);"
	SQLUPDATEWINNERSTANDINGS = "UPDATE standings SET wins = wins + 1, gamesPlayed = gamesPlayed + 1, goalDiff = goalDiff + ? WHERE teamId=? AND leagueId=?"
	SQLUPDATELOSERSTANDINGS  = "UPDATE standings SET loses = loses + 1, gamesPlayed = gamesPlayed + 1, goalDiff = goalDiff + ? WHERE teamId=? AND leagueId=?"
	SQLUPDATETIE             = "UPDATE standings SET ties = ties + 1, gamesPlayed = gamesPlayed + 1 WHERE teamId=? AND leagueId=?"
	SQLGETTEAMSTANDINGS      = "SELECT EXISTS (SELECT teamId FROM `standings` WHERE teamId=?)"
	SQLUPDATEGAMESCORE       = "UPDATE `schedule` SET `awayGoals` = ?, homeGoals = ? WHERE gameTime = ?"
)

//teamId	leagueId	gamesPlayed	wins	loses	ties	goalDiff
func insertTeamStandings(db *sql.DB, teamId int64, leagueId int64) {
	var _, err = db.Exec(SQLINSTERTSTANDINGS, teamId, leagueId, 0, 0, 0, 0, 0)
	if err != nil {
		log.Panic(err)
	} else {
		log.Printf("inserted: %d into Standings", teamId)
	}

}

func UpdateStandings(db *sql.DB, game models.Game) {
	var homeGoalDiff = game.HomeGoals - game.AwayGoals
	var homeExist bool
	var awayExist bool

	//check if team exists in standings
	home := db.QueryRow(SQLGETTEAMSTANDINGS, game.HomeTeamId).Scan(&homeExist)
	away := db.QueryRow(SQLGETTEAMSTANDINGS, game.AwayTeamId).Scan(&awayExist)

	if game.AwayTeamId == 0 || game.HomeTeamId == 0 {
		log.Panic(game.Played)
	}

	if !homeExist || home == sql.ErrNoRows {
		log.Print("NO TEAM FOUND!")
		insertTeamStandings(db, game.HomeTeamId, game.League)
	}
	if !awayExist || away == sql.ErrNoRows {
		log.Print("no away team found")
		insertTeamStandings(db, game.AwayTeamId, game.League)
	}

	if game.HomeGoals != 0 && game.AwayGoals != 0 && game.Played {

		switch {
		case homeGoalDiff == 0:
			log.Printf("Tie Game On: %d ", game.Date)
			updateTeamStandings(db, game.League, game.HomeTeamId, false, true, homeGoalDiff)
			updateTeamStandings(db, game.League, game.AwayTeamId, false, true, homeGoalDiff)
		case homeGoalDiff > 0:
			log.Print("Home Team wins!")
			updateTeamStandings(db, game.League, game.HomeTeamId, true, false, homeGoalDiff)
			updateTeamStandings(db, game.League, game.AwayTeamId, false, false, (game.AwayGoals - game.HomeGoals))
		default:
			log.Print("Away Team wins!")
			updateTeamStandings(db, game.League, game.HomeTeamId, false, false, homeGoalDiff)
			updateTeamStandings(db, game.League, game.AwayTeamId, true, false, (game.AwayGoals - game.HomeGoals))

		}
	}

}

func updateTeamStandings(db *sql.DB, leagueId int64, teamId int64, win bool, tie bool, goals int) {
	var result sql.Result
	var err error

	if tie {
		result, err = db.Exec(SQLUPDATETIE, teamId, leagueId)
	} else if win {
		result, err = db.Exec(SQLUPDATEWINNERSTANDINGS, goals, teamId, leagueId)
	} else {
		result, err = db.Exec(SQLUPDATELOSERSTANDINGS, goals, teamId, leagueId)
	}

	//check error
	if err != nil {
		log.Panic(err)
	}
	rowCnt, err := result.RowsAffected()

	if rowCnt > 0 {
		log.Printf("updated Standings for team: %d", teamId)
	}

}

func UpdateGameScore(db *sql.DB, game models.Game) {
	var _, err = db.Exec(SQLUPDATEGAMESCORE, game.AwayGoals, game.HomeGoals, game.GameTime)
	if err != nil {
		log.Panic(err)
	} else {
		log.Printf("Upated game: %s %s - %d vs %s - %d", game.GameTime, game.AwayTeam, game.AwayGoals, game.HomeTeam, game.HomeGoals)
	}

}
