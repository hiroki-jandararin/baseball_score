package review

func calculateBasePlayerScore(stats PlayerMatchStats) int {
	score := 0
	score += stats.Hits * 2
	score += stats.RBI * 3
	score += stats.Runs * 2
	score += stats.Walks
	score -= stats.Errors * 2
	score -= stats.Strikeouts
	return score
}

func CalculatePlayerScore(stats PlayerMatchStats) int {
	return calculateBasePlayerScore(stats)
}