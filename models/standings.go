package models

type Standings struct {
	TeamId      int64 `json:"teamId,omitempty"`
	LeagueId    int64 `json:"leagueId,omitempty"`
	GamesPlated int64 `json:"gamesPlayed,omitempty"`
	wins        int64 `json:"wins,omitempty"`
	loses       int64 `json:"loses,omitempty"`
	ties        int64 `json:"ties,omitempty"`
	goalDiff    int64 `json:"goalDi,omitempty"`
}
