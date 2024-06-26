package controllers

import "time"

type HostTournamentRequest struct {
	Title string    `json:"title" form:"title"`
	Date  time.Time `json:"date" form:"date"`
}

type EnrollPlayerRequest struct {
	TournamentID string `query:"id"`
	UserID       string `query:"userID"`
}
