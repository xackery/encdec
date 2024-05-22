package encdec

import "encoding/binary"

// Coder is an interface for shared methods between Encoder and Decoder.
type Coder interface {
	SetOrder(order binary.ByteOrder)
	LastError() error
	Error() error
	Pos() int64
}
