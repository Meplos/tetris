package leaderboard

type LeaderboardItems struct {
	Score int
	Name  string
}

type Leaderboard struct {
	Items []LeaderboardItems
}

func NewLeaderboard() *Leaderboard {
	return &Leaderboard{
		Items: make([]LeaderboardItems, 0),
	}
}
