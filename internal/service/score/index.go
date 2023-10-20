package score

type Score struct {
	scoreRepo     Repo
	leetCodeStats LeetCodeStats
}

func New(scoreRepo Repo, ls LeetCodeStats) *Score {
	return &Score{
		scoreRepo:     scoreRepo,
		leetCodeStats: ls,
	}
}
