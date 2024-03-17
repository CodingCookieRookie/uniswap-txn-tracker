package errors

type UserError struct {
	Msg string
}

func (u *UserError) Error() string {
	return u.Msg
}
