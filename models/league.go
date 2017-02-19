package models

import (
	"time"
)

type League struct {
	Id        int64
	Name      string
	Url       string
	StartDate time.Time
	Status    string
}
