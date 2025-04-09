package exts

import muss "github.com/mus-format/mus-stream-go"

// MarshallerMUS interface wraps MarhsalMUS and SizeMUS methods.
type MarshallerMUS interface {
	MarshalMUS(w muss.Writer) (n int, err error)
	SizeMUS() (size int)
}
