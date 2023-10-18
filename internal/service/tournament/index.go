package tournament

type Tournament struct {
	tournamentRepo Repo
}

func New(tournamentRepo Repo) *Tournament {
	return &Tournament{
		tournamentRepo: tournamentRepo,
	}
}
