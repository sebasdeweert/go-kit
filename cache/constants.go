package cache

import "errors"

// Direction represents a direction.
type Direction int

// Direction values.
const (
	Ascending Direction = iota
	Descending
)

// Errors messages.
const (
	ErrInvalidDirection = "the direction must be either Ascending or Descending"
)

// Errors.
var (
	ErrCacheMiss       = errors.New("cache miss")
	ErrDestNotPtrSlice = errors.New("destination is not a slice type")
)
