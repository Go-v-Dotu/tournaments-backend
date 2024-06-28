package domain

type Player struct {
	ID       string
	UserID   string
	Username string
}

func CreateGuestPlayer(id string, username string) *Player {
	return NewPlayer(id, "", username)
}

func NewPlayer(id string, userID string, username string) *Player {
	return &Player{
		ID:       id,
		UserID:   userID,
		Username: username,
	}
}
