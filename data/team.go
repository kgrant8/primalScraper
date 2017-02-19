package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	SQLGETTEAM    = "select id from `teams` where name=? AND leagueId=?;"
	SQLINSERTTEAM = "INSERT INTO `teams` (`id`, `name`, `leagueId`) VALUES (NULL, ?, ?);"
)

func GetTeamId(db *sql.DB, name string, leagueId int64) int64 {
	var id int64
	err := db.QueryRow(SQLGETTEAM, name, leagueId).Scan(&id)

	if err == sql.ErrNoRows {
		log.Printf("No team with that name.")
		id = insertTeam(db, name, leagueId)
	} else if err != nil {
		log.Fatal(err)
	}
	return id
}

func insertTeam(db *sql.DB, name string, league int64) int64 {

	var id int64
	var results, err = db.Exec(SQLINSERTTEAM, name, league)
	if err != nil {
		log.Print(err)
	}
	id, err = results.LastInsertId()
	if id > 0 {
		log.Printf("inserted: %s", name)
	} else if err != nil {
		log.Print(err)
	} else {
		log.Printf("Error inserting: %s", name)
	}
	return id

}
