package ext

import "github.com/mus-format/mus-stream-go"

// MarshallerMUS interface wraps MarhsalMUS and SizeMUS methods.
type MarshallerMUS interface {
	MarshalMUS(w mus.Writer) (n int, err error)
	SizeMUS() (size int)
}

// MarshallerTypedMUS interface wraps MarhsalTypedMUS and SizeTypedMUS methods.
// It is intended for use with DTS.
type MarshallerTypedMUS interface {
	MarshalTypedMUS(w mus.Writer) (n int, err error)
	SizeTypedMUS() (size int)
}
