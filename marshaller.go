package exts

import muss "github.com/mus-format/mus-stream-go"

// MarshallerMUS interface wraps MarhsalMUS and SizeMUS methods.
type MarshallerMUS interface {
	MarshalMUS(w muss.Writer) (n int, err error)
	SizeMUS() (size int)
}

// MarshallerTypedMUS interface wraps MarhsalTypedMUS and SizeTypedMUS methods.
// It is intended for use with DTS.
type MarshallerTypedMUS interface {
	MarshalTypedMUS(w muss.Writer) (n int, err error)
	SizeTypedMUS() (size int)
}
