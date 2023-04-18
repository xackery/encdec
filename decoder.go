package encdec

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Decoder is struct for decoding data.
type Decoder struct {
	order      binary.ByteOrder
	r          io.ReadSeeker
	firstError error
	lastError  error
}

// NewDecoder returns new Decoder.
func NewDecoder(r io.ReadSeeker, order binary.ByteOrder) *Decoder {
	return &Decoder{
		order: order,
		r:     r,
	}
}

// SetOrder sets byte order.
func (d *Decoder) SetOrder(order binary.ByteOrder) {
	d.order = order
}

// LastError returns last error that occured during read.
func (d *Decoder) LastError() error {
	return d.lastError
}

// Error returns first error that occured during read.
func (d *Decoder) Error() error {
	return d.firstError
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

func (d *Decoder) Byte() byte {
	return d.Bytes(1)[0]
}

// StringFixed returns fixed string.
func (d *Decoder) StringFixed(n int) string {
	return string(d.Bytes(n))
}

// StringLenPrefix returns string with length prefix assumed to be prior
func (d *Decoder) StringLenPrefixUint32() string {
	n := d.Uint32()
	return d.StringFixed(int(n))
}

// StringLenPrefix returns string with length prefix assumed to be prior
func (d *Decoder) StringLenPrefixUint16() string {
	n := d.Uint16()
	return d.StringFixed(int(n))
}

// StringLenPrefix returns string with length prefix assumed to be prior
func (d *Decoder) StringLenPrefixUint8() string {
	n := d.Uint8()
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
	return s
}

// Bool returns bool.
func (d *Decoder) Bool() bool {
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
	return v
}
