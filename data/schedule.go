package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"soccer/models"
)

const (
	SQLINSERTGAME = "INSERT INTO `schedule` (`id`, `leagueId`, `homeTeam`, `awayTeam`, `homeGoals`, `awayGoals`, `gameTime`, `lastUpdated`) VALUES (NULL, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP);"
)

func InsertGame(db *sql.DB, game models.Game, id int64) {

	var _, err = db.Exec(SQLINSERTGAME, id, game.HomeTeamId, game.AwayTeamId, game.HomeGoals, game.AwayGoals, game.GameTime)
	if err != nil {
		log.Print(err)
	} else {
		log.Printf("inserted game for: %s", game.GameTime)
	}
}
