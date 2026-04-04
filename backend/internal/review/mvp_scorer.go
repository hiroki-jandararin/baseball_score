package review

func ScoreMVP(match Match, stats PlayerMatchStats) int {
	score := 0
	score += stats.Hits * 2
	score += stats.RBI * 3
	score += stats.Runs * 2
	score += stats.Walks
	score -= stats.Errors * 2
	score -= stats.Strikeouts
	return score
}

func SelectMVP(match Match) MVPResult {
	best := MVPResult{}
	bestScore := 0
	first := true

	for _, stats := range match.PlayerStats {
		score := ScoreMVP(match, stats)
		if first || score > bestScore {
			best = MVPResult{
				PlayerID:   stats.PlayerID,
				PlayerName: stats.PlayerName,
				Score:      score,
			}
			bestScore = score
			first = false
		}
	}

	return best
}
