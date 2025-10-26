# ext-stream-go

**ext-stream-go** provides an extension for the [mus-stream-go](https://github.com/mus-format/mus-stream-go)
serializer, enabling additional functionality for the MUS format.

This package includes:

- `MarshallerMUS` — an interface for types that can marshal themselves into the
  MUS format.
- `MarshallerTypedMUS` — an interface for types that support typed MUS
  serialization (designed for use with [DTS](https://github.com/mus-format/dts-go)).

