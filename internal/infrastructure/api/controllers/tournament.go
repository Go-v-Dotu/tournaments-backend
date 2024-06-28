package controllers

import "time"

type HostTournamentRequest struct {
	TournamentInfo
}

type GetTournamentRequest struct {
	ID string `param:"id"`
}

type SelfEnrollRequest struct {
	ID string `param:"id"`
}

type GetPlayersRequest struct {
	TournamentID string `param:"id"`
}

type EnrollGuestPlayerRequest struct {
	TournamentID string `param:"id"`
	GuestUserInfo
}

type EnrollPlayerRequest struct {
	TournamentID string `param:"id"`
	UserID       string `param:"userID"`
}

type DropPlayerRequest struct {
	TournamentID string `param:"id"`
	PlayerID     string `param:"playerID"`
}

type TournamentInfo struct {
	Title string    `json:"title" form:"title"`
	Date  time.Time `json:"date" form:"date"`
}

type GuestUserInfo struct {
	Username string `json:"username" form:"username"`
}
