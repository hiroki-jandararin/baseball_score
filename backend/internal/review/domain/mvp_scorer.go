package review

func ScoreMVP(match Match, stats PlayerMatchStats) int {
	return calculateBasePlayerScore(stats)
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
