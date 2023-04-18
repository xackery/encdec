package encdec

import (
	"encoding/binary"
	"io"
)

// Encoder is struct for encoding data.
type Encoder struct {
	order      binary.ByteOrder
	w          io.Writer
	firstError error
	lastError  error
}

// NewEncoder returns new Encoder.
func NewEncoder(w io.Writer, order binary.ByteOrder) *Encoder {
	return &Encoder{
		order: order,
		w:     w,
	}
}

// SetOrder sets byte order.
func (e *Encoder) SetOrder(order binary.ByteOrder) {
	e.order = order
}

// Bytes writes bytes.
func (e *Encoder) Bytes(b []byte) {
	err := binary.Write(e.w, e.order, b)
	if err != nil {
		e.lastError = err
		if e.firstError != nil {
			e.firstError = e.lastError
		}
	}
}

// Byte writes byte.
func (e *Encoder) Byte(b byte) {
	err := binary.Write(e.w, e.order, b)
	if err != nil {
		e.lastError = err
		if e.firstError != nil {
			e.firstError = e.lastError
		}
	}
}

// String writes string.
func (e *Encoder) String(s string) {
	e.Bytes([]byte(s))
}

// StringZero writes string with zero terminator.
func (e *Encoder) StringZero(s string) {
	e.Bytes([]byte(s))
	e.Bytes([]byte{0})
}

// StringFixed writes fixed string.
func (e *Encoder) StringFixed(s string, n int) {
	if len(s) > n {
		s = s[:n]
	}
	if len(s) < n {
		s += string(make([]byte, n-len(s)))
	}
	e.Bytes([]byte(s))
}

// StringLenPrefixUint8 writes string with uint8 length prefix.
func (e *Encoder) StringLenPrefixUint8(s string) {
	e.Uint8(uint8(len(s)))
	e.String(s)
}

// StringLenPrefixUint16 writes string with uint16 length prefix.
func (e *Encoder) StringLenPrefixUint16(s string) {
	e.Uint16(uint16(len(s)))
	e.String(s)
}

// StringLenPrefixUint32 writes string with uint32 length prefix.
func (e *Encoder) StringLenPrefixUint32(s string) {
	e.Uint32(uint32(len(s)))
	e.String(s)
}

// Uint8 writes uint8.
func (e *Encoder) Uint8(v uint8) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError != nil {
			e.firstError = e.lastError
		}
	}
}

// Uint16 writes uint16.
func (e *Encoder) Uint16(v uint16) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError != nil {
			e.firstError = e.lastError
		}
	}
}

// Uint32 writes uint32.
func (e *Encoder) Uint32(v uint32) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError != nil {
			e.firstError = e.lastError
		}
	}
}

// Uint64 writes uint64.
func (e *Encoder) Uint64(v uint64) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError != nil {
			e.firstError = e.lastError
		}
	}
}

// Int8 writes int8.
func (e *Encoder) Int8(v int8) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError != nil {
			e.firstError = e.lastError
		}
	}
}

// Int16 writes int16.
func (e *Encoder) Int16(v int16) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError != nil {
			e.firstError = e.lastError
		}
	}
}

// Int32 writes int32.
func (e *Encoder) Int32(v int32) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError != nil {
			e.firstError = e.lastError
		}
	}
}

// Int64 writes int64.
func (e *Encoder) Int64(v int64) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError != nil {
			e.firstError = e.lastError
		}
	}
}

// Float32 writes float32.
func (e *Encoder) Float32(v float32) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError != nil {
			e.firstError = e.lastError
		}
	}
}

// Float64 writes float64.
func (e *Encoder) Float64(v float64) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError != nil {
			e.firstError = e.lastError
		}
	}
}

// Bool writes bool.
func (e *Encoder) Bool(v bool) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError != nil {
			e.firstError = e.lastError
		}
	}
}

// LastError returns last error that occured during write.
func (e *Encoder) LastError() error {
	return e.lastError
}

// FirstError returns first error that occured during write.
func (e *Encoder) Error() error {
	return e.firstError
}
