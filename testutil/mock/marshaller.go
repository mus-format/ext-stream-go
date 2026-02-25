package mock

import (
	"github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

type (
	MarshalMUSFn func(w mus.Writer) (n int, err error)
	SizeMUSFn    func() (size int)
)

func NewMarshallerMUS() MarshallerMUS {
	return MarshallerMUS{mok.New("MarshallerMUS")}
}

type MarshallerMUS struct {
	*mok.Mock
}

func (m MarshallerMUS) RegisterMarshalMUS(fn MarshalMUSFn) MarshallerMUS {
	m.Register("MarshalMUS", fn)
	return m
}

func (m MarshallerMUS) RegisterSizeMUS(fn SizeMUSFn) MarshallerMUS {
	m.Register("SizeeMUS", fn)
	return m
}

func (m MarshallerMUS) MarshalMUS(w mus.Writer) (n int, err error) {
	result, err := m.Call("MarshalMUS", w)
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}

func (m MarshallerMUS) SizeMUS() (size int) {
	result, err := m.Call("SizeMUS")
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}
