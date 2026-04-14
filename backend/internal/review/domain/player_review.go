package review

type PlayerReview struct {
	PlayerID      int
	PlayerName    string
	Stats         PlayerMatchStats
	Title         string
	Comment       string
	IsMVP         bool
}