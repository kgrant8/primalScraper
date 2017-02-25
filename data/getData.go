package data

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"soccer/models"
	"strconv"
	"strings"
	"time"
)

const (
	longForm    = "Mon, Jan 2 2006 at 3:04 PM"
	SQLDATETIME = "2006-01-02 15:04:05"
)

func GetScheduleData(siteUrl string, id int64) []models.Game {
	season := make([]models.Game, 0)

	doc, err := goquery.NewDocument(siteUrl)
	if err != nil {
		log.Fatal(err)
	}
	loc, _ := time.LoadLocation("america/los_angeles")

	doc.Find(".schedule-game").Each(func(i int, s *goquery.Selection) {
		g := models.Game{}
		g.Date = s.Find(".date").Text()
		stringTime := s.Find(".time").Text()
		g.AwayTeam = s.Find("span a").First().Text()
		g.AwayGoals, _ = strconv.Atoi(s.Find("span strong.score ").First().Text())
		g.HomeTeam = s.Find("span a").Last().Text()
		g.HomeGoals, _ = strconv.Atoi(s.Find("span strong.score").Last().Text())
		g.GameType = s.Find(".game-type.schedule-tag").Text()
		played := s.Find(".game-type.tag").Text()

		dateTime := fmt.Sprintf("%s 2017 at %s", g.Date, stringTime)
		t, err := time.ParseInLocation(longForm, dateTime, loc)
		if err != nil {
			panic(err)
		}
		g.Time = t
		//coverting time to UTC
		t = t.UTC()
		g.GameTime = t.Format(SQLDATETIME)
		g.League = id
		g.Played = hasBeenPlayed(played, t)

		season = append(season, g)
	})

	return season

}
func GetLeagueData(url string) []models.League {
	leagues := make([]models.League, 0)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("[id^=baseevent]").Each(func(i int, s *goquery.Selection) {
		l := models.League{}

		id, _ := s.Attr("id")
		l.Url = getLeagueSchduleUrl(id[10:])

		l.Name = s.Find(".meta-info h2").Text()
		l.Id = getLeagueId(l.Name)

		l.Status = s.Find(".status").Text()

		ls := s.Find(".basic.clr dd.program-list-starts").Text()
		l.StartDate = ConvertStringToTime(ls)

		leagues = append(leagues, l)

	})
	return leagues

}

func ConvertStringToTime(timeString string) time.Time {
	timeString = strings.TrimSpace(timeString)
	x := strings.Index(timeString, "-")
	year, _ := strconv.Atoi(timeString[:x])
	x = strings.LastIndex(timeString, "-")
	m, _ := (strconv.Atoi(timeString[x-2 : x]))
	month := time.Month(m)
	day, _ := strconv.Atoi(timeString[x+1 : x+3])
	date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return date

}
func hasBeenPlayed(gameType string, t time.Time) bool {
	var passed bool
	now := time.Now()
	if gameType == "Rescheduled" || gameType == "Forfeit" {
		passed = false
	}
	if ((now.Sub(t)) > 0) && (gameType != "Rescheduled" || gameType != "Forfeit") {
		passed = true
	} else {
		passed = false
	}
	return passed

}
