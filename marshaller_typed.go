package exts

import muss "github.com/mus-format/mus-stream-go"

// MarshallerTypedMUS interface wraps MarhsalMUS and SizeMUS methods. It is
// intended for use with DTS.
type MarshallerTypedMUS interface {
	MarshalTypedMUS(w muss.Writer) (n int, err error)
	SizeTypedMUS() (size int)
}
