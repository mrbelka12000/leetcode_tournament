package usrevent

type UsrEvent struct {
	usrEventRepo Repo
}

func New(usrEventRepo Repo) *UsrEvent {
	return &UsrEvent{
		usrEventRepo: usrEventRepo,
	}
}
