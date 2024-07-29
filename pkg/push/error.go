package push

type NotVerifiedError struct{}

func (e NotVerifiedError) Error() string {
	return "sender not verified"
}

func NewNotVerifiedError() error {
	return NotVerifiedError{}
}
