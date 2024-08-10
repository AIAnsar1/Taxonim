package Data

type RemoteCsvError struct {
	Message      string
	WrappedError error
}
