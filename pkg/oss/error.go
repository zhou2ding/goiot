package oss

type NoPermissionError struct{}

func (e NoPermissionError) Error() string {
	return "no permission"
}

func NewNoPermissionError() error {
	return NoPermissionError{}
}
