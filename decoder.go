package encdec

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

// Decoder is struct for decoding data.
type Decoder struct {
	order       binary.ByteOrder
	r           io.ReadSeeker
	firstError  error
	lastError   error
	isDebugMode bool
	debugBuf    bytes.Buffer
}

// NewDecoder returns new Decoder.
func NewDecoder(r io.ReadSeeker, order binary.ByteOrder) *Decoder {
	return &Decoder{
		order: order,
		r:     r,
	}
}

// SetDebugMode enables every decode call to write to a stored buffer in the decoder to review later
func (d *Decoder) SetDebugMode(value bool) {
	d.isDebugMode = value
}

// DebugBuf returns the debug buffer
func (d *Decoder) DebugBuf() []byte {
	return d.debugBuf.Bytes()
}

// DebugString returns the debug buffer as a string
func (d *Decoder) DebugString() string {
	return string(d.debugBuf.Bytes())
}

// DebugClear clears the debug buffer
func (d *Decoder) DebugClear() {
	d.debugBuf.Reset()
}

// IsDebugMode returns if debug mode is enabled
func (d *Decoder) IsDebugMode() bool {
	return d.isDebugMode
}

// SetOrder sets byte order.
func (d *Decoder) SetOrder(order binary.ByteOrder) {
	d.order = order
}

// LastError returns last error that occurred during read.
func (d *Decoder) LastError() error {
	return d.lastError
}

// Error returns first error that occurred during read.
func (d *Decoder) Error() error {
	return d.firstError
}

// Pos returns current position.
func (d *Decoder) Pos() int64 {
	pos, err := d.r.Seek(0, io.SeekCurrent)
	if err != nil {
		pos = -1
	}
	return pos
}

// Bytes returns bytes.
func (d *Decoder) Bytes(n int) []byte {
	b := make([]byte, n)
	_, err := io.ReadFull(d.r, b)
	if err != nil {
		pos, err := d.r.Seek(0, io.SeekCurrent)
		if err != nil {
			pos = -1
		}
		d.lastError = fmt.Errorf("pos %d: %w", pos, err)
		if d.firstError != nil {
			d.firstError = d.lastError
		}
	}
	return b
}

// Byte returns byte.
func (d *Decoder) Byte() byte {
	if d.isDebugMode {
		value := d.Bytes(1)[0]
		d.debugBuf.Write([]byte{value})
		return value
	}
	return d.Bytes(1)[0]
}

// StringFixed returns fixed string.
func (d *Decoder) StringFixed(n int) string {
	if d.isDebugMode {
		value := string(d.Bytes(n))
		d.debugBuf.Write([]byte(value))
		return value
	}
	return string(d.Bytes(n))
}

// StringLenPrefixUint32 returns string with length prefix assumed to be prior
func (d *Decoder) StringLenPrefixUint32() string {
	n := d.Uint32()
	if d.isDebugMode {
		value := d.StringFixed(int(n))
		d.debugBuf.Write([]byte(value))
		return value
	}
	return d.StringFixed(int(n))
}

// StringLenPrefixUint16 returns string with length prefix assumed to be prior
func (d *Decoder) StringLenPrefixUint16() string {
	n := d.Uint16()
	if d.isDebugMode {
		value := d.StringFixed(int(n))
		d.debugBuf.Write([]byte(value))
		return value
	}
	return d.StringFixed(int(n))
}

// StringLenPrefixUint8 returns string with length prefix assumed to be prior
func (d *Decoder) StringLenPrefixUint8() string {
	n := d.Uint8()
	if d.isDebugMode {
		value := d.StringFixed(int(n))
		d.debugBuf.Write([]byte(value))
		return value
	}

	return d.StringFixed(int(n))
}

// StringZero reads the read stream until a zero terminator is found.
func (d *Decoder) StringZero() string {
	var s string
	var buf [1]byte
	var err error
	for {
		_, err = io.ReadFull(d.r, buf[:])
		if err != nil {
			pos, err := d.r.Seek(0, io.SeekCurrent)
			if err != nil {
				pos = -1
			}
			d.lastError = fmt.Errorf("pos %d: %w", pos, err)
			if d.firstError != nil {
				d.firstError = d.lastError
			}
			break
		}
		if buf[0] == 0 {
			break
		}
		s += string(buf[:])
	}
	if d.isDebugMode {
		d.debugBuf.Write([]byte(s))
	}

	return s
}

// Bool returns bool.
func (d *Decoder) Bool() bool {
	if d.isDebugMode {
		value := d.Byte() != 0
		d.debugBuf.Write([]byte{0})
		return value
	}

	return d.Byte() != 0
}

// Uint8 returns uint8.
func (d *Decoder) Uint8() uint8 {
	var v uint8
	err := binary.Read(d.r, d.order, &v)
	if err != nil {
		pos, err := d.r.Seek(0, io.SeekCurrent)
		if err != nil {
			pos = -1
		}
		d.lastError = fmt.Errorf("pos %d: %w", pos, err)
		if d.firstError != nil {
			d.firstError = d.lastError
		}
	}
	if d.isDebugMode {
		d.debugBuf.Write([]byte{v})
	}
	return v
}

// Uint16 returns uint16.
func (d *Decoder) Uint16() uint16 {
	var v uint16
	err := binary.Read(d.r, d.order, &v)
	if err != nil {
		pos, err := d.r.Seek(0, io.SeekCurrent)
		if err != nil {
			pos = -1
		}
		d.lastError = fmt.Errorf("pos %d: %w", pos, err)
		if d.firstError != nil {
			d.firstError = d.lastError
		}
	}

	if d.isDebugMode {
		d.debugBuf.Write([]byte{byte(v >> 8), byte(v)})
	}
	return v
}

// Uint32 returns uint32.
func (d *Decoder) Uint32() uint32 {
	var v uint32
	err := binary.Read(d.r, d.order, &v)
	if err != nil {
		pos, err := d.r.Seek(0, io.SeekCurrent)
		if err != nil {
			pos = -1
		}

		d.lastError = fmt.Errorf("pos %d: %w", pos, err)
		if d.firstError != nil {
			d.firstError = d.lastError
		}
	}

	if d.isDebugMode {
		d.debugBuf.Write([]byte{byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)})
	}

	return v
}

// Uint64 returns uint64.
func (d *Decoder) Uint64() uint64 {
	var v uint64
	err := binary.Read(d.r, d.order, &v)
	if err != nil {
		pos, err := d.r.Seek(0, io.SeekCurrent)
		if err != nil {
			pos = -1
		}

		d.lastError = fmt.Errorf("pos %d: %w", pos, err)
		if d.firstError != nil {
			d.firstError = d.lastError
		}
	}

	if d.isDebugMode {
		d.debugBuf.Write([]byte{byte(v >> 56), byte(v >> 48), byte(v >> 40), byte(v >> 32),
			byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)})
	}
	return v
}

// Int8 returns int8.
func (d *Decoder) Int8() int8 {
	var v int8
	err := binary.Read(d.r, d.order, &v)
	if err != nil {
		pos, err := d.r.Seek(0, io.SeekCurrent)
		if err != nil {
			pos = -1
		}
		d.lastError = fmt.Errorf("pos %d: %w", pos, err)
		if d.firstError != nil {
			d.firstError = d.lastError
		}
	}

	if d.isDebugMode {
		d.debugBuf.Write([]byte{byte(v)})
	}

	return v
}

// Int16 returns int16.
func (d *Decoder) Int16() int16 {
	var v int16
	err := binary.Read(d.r, d.order, &v)
	if err != nil {
		pos, err := d.r.Seek(0, io.SeekCurrent)
		if err != nil {
			pos = -1
		}
		d.lastError = fmt.Errorf("pos %d: %w", pos, err)
		if d.firstError != nil {
			d.firstError = d.lastError
		}
	}

	if d.isDebugMode {
		d.debugBuf.Write([]byte{byte(v >> 8), byte(v)})
	}

	return v
}

// Int32 returns int32.
func (d *Decoder) Int32() int32 {
	var v int32
	err := binary.Read(d.r, d.order, &v)
	if err != nil {
		pos, err := d.r.Seek(0, io.SeekCurrent)
		if err != nil {
			pos = -1
		}

		d.lastError = fmt.Errorf("pos %d: %w", pos, err)
		if d.firstError != nil {
			d.firstError = d.lastError
		}
	}

	if d.isDebugMode {
		d.debugBuf.Write([]byte{byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)})
	}

	return v
}

// Int64 returns int64.
func (d *Decoder) Int64() int64 {
	var v int64
	err := binary.Read(d.r, d.order, &v)
	if err != nil {
		pos, err := d.r.Seek(0, io.SeekCurrent)
		if err != nil {
			pos = -1
		}

		d.lastError = fmt.Errorf("pos %d: %w", pos, err)
		if d.firstError != nil {
			d.firstError = d.lastError
		}
	}

	if d.isDebugMode {
		d.debugBuf.Write([]byte{byte(v >> 56), byte(v >> 48), byte(v >> 40), byte(v >> 32),
			byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)})
	}

	return v
}

// Float32 returns float32.
func (d *Decoder) Float32() float32 {
	var v float32
	err := binary.Read(d.r, d.order, &v)
	if err != nil {
		pos, err := d.r.Seek(0, io.SeekCurrent)
		if err != nil {
			pos = -1
		}

		d.lastError = fmt.Errorf("pos %d: %w", pos, err)
		if d.firstError != nil {
			d.firstError = d.lastError
		}
	}

	if d.isDebugMode {
		// Convert float32 to uint32 bits and write those bytes
		bits := math.Float32bits(v)
		d.debugBuf.Write([]byte{byte(bits >> 24), byte(bits >> 16), byte(bits >> 8), byte(bits)})
	}

	return v
}

// Float64 returns float64.
func (d *Decoder) Float64() float64 {
	var v float64
	err := binary.Read(d.r, d.order, &v)
	if err != nil {
		pos, err := d.r.Seek(0, io.SeekCurrent)
		if err != nil {
			pos = -1
		}

		d.lastError = fmt.Errorf("pos %d: %w", pos, err)
		if d.firstError != nil {
			d.firstError = d.lastError
		}
	}

	if d.isDebugMode {
		// Convert float64 to uint64 bits and write those bytes
		bits := math.Float64bits(v)
		d.debugBuf.Write([]byte{byte(bits >> 56), byte(bits >> 48), byte(bits >> 40), byte(bits >> 32),
			byte(bits >> 24), byte(bits >> 16), byte(bits >> 8), byte(bits)})
	}

	return v
}
