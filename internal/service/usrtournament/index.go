package usrtournament

type UsrTournament struct {
	usrTourRepo Repo
}

func New(usrTourRepo Repo) *UsrTournament {
	return &UsrTournament{
		usrTourRepo: usrTourRepo,
	}
}
