package data

import (
	"fmt"
	"log"
	"strings"
)

const (
	URLBASE = "http://www.primalsoccer.com/ajax/loadSchedule?origin=site&scope=program&itemType=games_events&programId=%s"
)

/*
coed B = 1
coed elite = 2
mens elite = 3
mens B = 4
*/

func getLeagueId(league string) int64 {
	var id int64

	if strings.Contains(league, "Co-Ed B") {
		id = 1

	} else if strings.Contains(league, "Co-Ed Elite") {
		id = 2
	} else if strings.Contains(league, "Men's B") {
		id = 3
	} else if strings.Contains(league, "Men's Elite") {
		id = 4
	} else {
		log.Panic("No League Id found")
	}
	return id

}

func getLeagueSchduleUrl(id string) string {
	return fmt.Sprintf(URLBASE, id)

}
