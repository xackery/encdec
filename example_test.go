package encdec

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"testing"
)

type exampleStruct struct {
	val1        int16
	val2        uint32
	someStr1    string
	someStr2    string
	someStrZero string
	val3        float32
	someBytes   []byte
}

func TestExample(t *testing.T) {
	r := bytes.NewReader([]byte{
		0x01, 0x00, // val1
		0x02, 0x00, 0x00, 0x00, // val2
		0x61, 0x62, 0x63, // someStr1
		0x04, 0x00, 0x00, 0x00, 0x64, 0x65, 0x66, 0x67, // someStr2
		0x68, 0x69, 0x6a, 0x6b, 0x00, // someStrZero
		0x00, 0x00, 0x80, 0x3f, // val3
		0x01, 0x02, 0x03, 0x04, // someBytes
	})
	def, err := someReadSeekerExample(r)
	if err != nil {
		t.Fatalf("Failed someReadSeekerExample: %v", err)
	}
	fmt.Printf("%+v\n", def)

	w := bytes.NewBuffer(nil)
	err = def.someWriterExample(w)
	if err != nil {
		t.Fatalf("Failed someWriterExample: %v", err)
	}

}

func someReadSeekerExample(r io.ReadSeeker) (*exampleStruct, error) {
	dec := NewDecoder(r, binary.LittleEndian)

	def := &exampleStruct{} // initialize an example struct in def

	def.val1 = dec.Int16()                     // decode int16 worth of bytes to val1,
	def.val2 = dec.Uint32()                    // decode int16 worth of bytes to val2
	def.someStr1 = dec.StringFixed(3)          // read 3 bytes and convert to a string
	def.someStr2 = dec.StringLenPrefixUint32() // read 4 bytes (uint32) to sort length of string, then read length and convert to string
	def.someStrZero = dec.StringZero()         // read until 0x00 and convert to string
	def.val3 = dec.Float32()                   // read 4 bytes and convertt float and place into val3
	def.someBytes = dec.Bytes(4)               // read 4 bytes and place into someBytes

	// all error handling can be checked at end like this:
	if dec.Error() != nil {
		return nil, fmt.Errorf("decode: %w", dec.Error())
	}
	return def, nil
}

func (def *exampleStruct) someWriterExample(w io.Writer) error {
	enc := NewEncoder(w, binary.LittleEndian)

	enc.Int16(def.val1)                     // encode val1 as int16
	enc.Uint32(def.val2)                    // encode val2 as uint32
	enc.StringZero(def.someStrZero)         // write string until 0x00
	enc.StringFixed(def.someStr1, 3)        // write 3 bytes of string
	enc.StringLenPrefixUint32(def.someStr2) // write 4 bytes (uint32) to sort length of string, then write length and convert to string
	enc.Float32(def.val3)                   // write 4 bytes and convertt float and place into val3
	if enc.Error() != nil {
		return fmt.Errorf("encode: %w", enc.Error())
	}
	return nil
}
