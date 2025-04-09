package mock

import (
	muss "github.com/mus-format/mus-stream-go"
	"github.com/ymz-ncnk/mok"
)

type MarshalTypedMUSFn[T any] func(w muss.Writer) (n T, err error)
type SizeTypedMUSFn func() (size int)

func NewMarshallerTypedMUS[T any]() MarshallerTypedMUS[T] {
	return MarshallerTypedMUS[T]{mok.New("MarshallerTypedMUS")}
}

type MarshallerTypedMUS[T any] struct {
	*mok.Mock
}

func (m MarshallerTypedMUS[T]) RegisterMarshalTypedMUS(fn MarshalTypedMUSFn[T]) MarshallerTypedMUS[T] {
	m.Register("MarshalTypedMUS", fn)
	return m
}

func (m MarshallerTypedMUS[T]) RegisterSizeTypedMUS(fn SizeTypedMUSFn) MarshallerTypedMUS[T] {
	m.Register("SizeTypedMUS", fn)
	return m
}

func (m MarshallerTypedMUS[T]) MarshalTypedMUS(w muss.Writer) (n int, err error) {
	result, err := m.Call("MarshalTypedMUS", mok.SafeVal[muss.Writer](w))
	if err != nil {
		panic(err)
	}
	n = result[0].(int)
	err, _ = result[1].(error)
	return
}

func (m MarshallerTypedMUS[T]) SizeTypedMUS() (size int) {
	result, err := m.Call("SizeTypedMUS")
	if err != nil {
		panic(err)
	}
	return result[0].(int)
}
