package Types

type IError interface {
	Error() string
	Unwrap() string
}
