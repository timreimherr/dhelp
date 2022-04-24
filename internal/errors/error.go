package errors

type Error string

func (e Error) Error() string { return string(e) }

const (
	ErrSectionAlreadyExists = Error("section already exists")
	ErrSectionDoesNotExist  = Error("entity does not exist")
)
