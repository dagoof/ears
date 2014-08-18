package ears

// EventualError exposes an error that may be not be immediately set. Can be
// used to return an error after some processing has taken place in a goroutine
// without the need for an explicit error channel and managing blocking states
// on a result channel and error channel.
type EventualError interface {
	Err() error
}

type SetEventualError interface {
	EventualError
	Set(error) bool
}
