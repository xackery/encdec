package encdec

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
)

// Encoder is struct for encoding data.
type Encoder struct {
	order       binary.ByteOrder
	w           io.Writer
	firstError  error
	lastError   error
	lastPos     int64
	isDebugMode bool
	debugBuf    bytes.Buffer
}

// NewEncoder returns new Encoder.
func NewEncoder(w io.Writer, order binary.ByteOrder) *Encoder {
	return &Encoder{
		order: order,
		w:     w,
	}
}

// SetDebugMode enables every encode call to write to a stored buffer in the encoder to review later
func (e *Encoder) SetDebugMode(value bool) {
	e.isDebugMode = value
}

// DebugBuf returns the debug buffer
func (e *Encoder) DebugBuf() []byte {
	return e.debugBuf.Bytes()
}

// DebugString returns the debug buffer as a string
func (e *Encoder) DebugString() string {
	return string(e.debugBuf.Bytes())
}

// DebugClear clears the debug buffer
func (e *Encoder) DebugClear() {
	e.debugBuf.Reset()
}

// IsDebugMode returns if debug mode is enabled
func (e *Encoder) IsDebugMode() bool {
	return e.isDebugMode
}

// SetOrder sets byte order.
func (e *Encoder) SetOrder(order binary.ByteOrder) {
	e.order = order
}

// Pos returns current position (attemps to track if not a seeker writer).
func (e *Encoder) Pos() int64 {
	seeker, ok := e.w.(io.Seeker)
	if !ok {
		return e.lastPos
	}
	pos, err := seeker.Seek(0, io.SeekCurrent)
	if err != nil {
		pos = -1
	}
	return pos
}

// Bytes writes bytes.
func (e *Encoder) Bytes(b []byte) {
	err := binary.Write(e.w, e.order, b)
	if err != nil {
		e.lastError = err
		if e.firstError == nil {
			e.firstError = e.lastError
		}
	}
	if e.isDebugMode {
		e.debugBuf.Write(b)
	}
	e.lastPos += int64(len(b))
}

// Byte writes byte.
func (e *Encoder) Byte(b byte) {
	err := binary.Write(e.w, e.order, b)
	if err != nil {
		e.lastError = err
		if e.firstError == nil {
			e.firstError = e.lastError
		}
	}
	if e.isDebugMode {
		e.debugBuf.WriteByte(b)
	}
	e.lastPos++
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
		if e.firstError == nil {
			e.firstError = e.lastError
		}
	}
	if e.isDebugMode {
		e.debugBuf.WriteByte(v)
	}
	e.lastPos += 1
}

// Uint16 writes uint16.
func (e *Encoder) Uint16(v uint16) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError == nil {
			e.firstError = e.lastError
		}
	}
	if e.isDebugMode {
		e.debugBuf.Write([]byte{byte(v >> 8), byte(v)})
	}
	e.lastPos += 2
}

// Uint32 writes uint32.
func (e *Encoder) Uint32(v uint32) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError == nil {
			e.firstError = e.lastError
		}
	}
	if e.isDebugMode {
		e.debugBuf.Write([]byte{byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)})
	}
	e.lastPos += 4
}

// Uint64 writes uint64.
func (e *Encoder) Uint64(v uint64) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError == nil {
			e.firstError = e.lastError
		}
	}
	if e.isDebugMode {
		e.debugBuf.Write([]byte{byte(v >> 56), byte(v >> 48), byte(v >> 40), byte(v >> 32),
			byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)})
	}
	e.lastPos += 8
}

// Int8 writes int8.
func (e *Encoder) Int8(v int8) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError == nil {
			e.firstError = e.lastError
		}
	}
	if e.isDebugMode {
		e.debugBuf.WriteByte(byte(v))
	}
	e.lastPos += 1
}

// Int16 writes int16.
func (e *Encoder) Int16(v int16) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError == nil {
			e.firstError = e.lastError
		}
	}
	if e.isDebugMode {
		e.debugBuf.Write([]byte{byte(v >> 8), byte(v)})
	}
	e.lastPos += 2
}

// Int32 writes int32.
func (e *Encoder) Int32(v int32) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError == nil {
			e.firstError = e.lastError
		}
	}
	if e.isDebugMode {
		e.debugBuf.Write([]byte{byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)})
	}
	e.lastPos += 4
}

// Int64 writes int64.
func (e *Encoder) Int64(v int64) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError == nil {
			e.firstError = e.lastError
		}
	}
	if e.isDebugMode {
		e.debugBuf.Write([]byte{byte(v >> 56), byte(v >> 48), byte(v >> 40), byte(v >> 32),
			byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)})
	}
	e.lastPos += 8
}

// Float32 writes float32.
func (e *Encoder) Float32(v float32) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError == nil {
			e.firstError = e.lastError
		}
	}
	if e.isDebugMode {
		// Convert float32 to uint32 bits and write those bytes
		bits := math.Float32bits(v)
		e.debugBuf.Write([]byte{byte(bits >> 24), byte(bits >> 16), byte(bits >> 8), byte(bits)})
	}
	e.lastPos += 4
}

// Float64 writes float64.
func (e *Encoder) Float64(v float64) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError == nil {
			e.firstError = e.lastError
		}
	}
	if e.isDebugMode {
		bits := math.Float64bits(v)
		e.debugBuf.Write([]byte{byte(bits >> 56), byte(bits >> 48), byte(bits >> 40), byte(bits >> 32),
			byte(bits >> 24), byte(bits >> 16), byte(bits >> 8), byte(bits)})
	}
	e.lastPos += 8
}

// Bool writes bool.
func (e *Encoder) Bool(v bool) {
	err := binary.Write(e.w, e.order, v)
	if err != nil {
		e.lastError = err
		if e.firstError == nil {
			e.firstError = e.lastError
		}
	}
	if e.isDebugMode {
		if v {
			e.debugBuf.WriteByte(1)
		} else {
			e.debugBuf.WriteByte(0)
		}
	}
	e.lastPos += 1
}

// LastError returns last error that occurred during write.
func (e *Encoder) LastError() error {
	return e.lastError
}

// FirstError returns first error that occurred during write.
func (e *Encoder) Error() error {
	return e.firstError
}
