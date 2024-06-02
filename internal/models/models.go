package models

import (
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

var DB string

type RestBody struct {
	Group    string `json:"group"`
	LastName string `json:"lastName"`
}

type RestResult struct {
	LastName string `json:"lastName"`
	Stats    string `json:"stats"`
}

var Tools tools

type tools struct {
	Connection *sqlx.DB
}

type ScheduleRequest struct {
	Group string `json:"group"`
	Date  string `json:"date"`
}

type ScheduleResponse struct {
	Lessons []*Lesson
}

type Lesson struct {
	ID        uuid.UUID `json:"id"`
	Date      time.Time `json:"date"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Professor string    `json:"professor"`
	Class     string    `json:"class"`
}
