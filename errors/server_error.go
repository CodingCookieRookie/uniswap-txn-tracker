package errors

type ServerError struct {
	Msg string
}

func (s *ServerError) Error() string {
	return s.Msg
}
