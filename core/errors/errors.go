package errors

type RecordNotFoundErr struct {
	Message string
}

func (e RecordNotFoundErr) Error() string {
	return e.Message
}
