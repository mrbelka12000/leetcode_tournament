package errs

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrEmailNotFound     = Error("email_not_found")
	ErrInvalidEmail      = Error("invalid_email")
	ErrEmailExists       = Error("email_exists")
	ErrUsernameNotFound  = Error("username_not_found")
	ErrInvalidUsername   = Error("invalid_username")
	ErrUsernameExists    = Error("username_exists")
	ErrNameNotFound      = Error("name_not_found")
	ErrInvalidName       = Error("invalid_name")
	ErrPasswordNotFound  = Error("password_not_found")
	ErrInvalidPassword   = Error("invalid_password")
	ErrPasswordDontMatch = Error("password_dont_match")
	ErrUsrStatusNotFound = Error("status_not_found")
	ErrInvalidUsrStatus  = Error("invalid_status")
	ErrUsrTypeNotFound   = Error("type_not_found")
	ErrInvalidUsrType    = Error("invalid_type")

	ErrUsrNotFound = Error("usr_not_found")
)
