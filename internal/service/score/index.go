package score

type Score struct {
	scoreRepo Repo
}

func New(scoreRepo Repo) *Score {
	return &Score{
		scoreRepo: scoreRepo,
	}
}
