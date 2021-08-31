package encoding

// Encoder represents encoding implementation.
type Encoder interface {
	Encode(obj interface{}) (*string, error)
	Decode(encoded string, obj interface{}) error
}
