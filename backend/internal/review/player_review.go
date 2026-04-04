package review

type PlayerReview struct {
	PlayerID      string
	PlayerName    string
	Stats         PlayerMatchStats
	Title         string
	Comment       string
	IsMVP         bool
}
