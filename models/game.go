package models

import (
	"time"
)

type Game struct {
	League     int64     `json:"league,omitempty"`
	Date       string    `json:"date,omitempty"`
	Time       time.Time `json:"time,omitempty"`
	GameTime   string    `json:"gameTime",omitempty"`
	AwayTeam   string    `json:"awayTeam,omitempty"`
	AwayGoals  int       `json:"awayGoals,omitempty"`
	HomeTeam   string    `json:"homeTeam,omitempty"`
	HomeGoals  int       `json:"homeGoals,omitempty"`
	HomeTeamId int64     `json:"homeTeamId,omitempty"`
	AwayTeamId int64     `json:"awayTeamId,omitempty"`
	GameType   string    `json:"gameType",omitempty"`
	Played     bool      `json:"played,omitempty"`
}
