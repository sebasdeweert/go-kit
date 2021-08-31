package redis

// Errors.
const (
	ErrNilResponse     = "redis: nil"
	ErrNotPong         = "unexpected return value"
	ErrNoExpirationSet = "no expiration set"
)

// Other.
const (
	PongMessage = "PONG"
)
