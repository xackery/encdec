package encdec

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Decoder is struct for decoding data.
type Decoder struct {
	order       binary.ByteOrder
	r           io.ReadSeeker
	recentError error
	lastError   error
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
		if d.recentError != nil {
			d.recentError = d.lastError
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
		if d.recentError != nil {
			d.recentError = d.lastError
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
		if d.recentError != nil {
			d.recentError = d.lastError
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
		if d.recentError != nil {
			d.recentError = d.lastError
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
		if d.recentError != nil {
			d.recentError = d.lastError
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
		if d.recentError != nil {
			d.recentError = d.lastError
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
		if d.recentError != nil {
			d.recentError = d.lastError
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
		if d.recentError != nil {
			d.recentError = d.lastError
		}
	}
	return v
}
