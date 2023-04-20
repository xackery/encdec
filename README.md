# encdec
[![GoDoc](https://godoc.org/github.com/xackery/encdec?status.svg)](https://godoc.org/github.com/xackery/encdec) [![Go Report Card](https://goreportcard.com/badge/github.com/xackery/encdec)](https://goreportcard.com/report/github.com/xackery/encdec)

Golang encoder/decoder from io.ReadSeeker and io.Writer

Please note this is not a very rigorous way to decode a io.ReadSeeker/Writer.

I use io.ReadSeeker over io.Reader so the package can know the position of a failure.

To obtain, just run  `go get github.com/xackery/encdec`

- Perk: Github copilot works very smoothly with this approach. Define a struct, get decoder initialized, and watch as copilot fills all the decoding fields one by one, even the subStruct example below was filled with copilot. Same flow for encoder.
- Perk: Easy to read and modify later. Most fields are a single line, making it easy to identify and fix later.
- Perk: Don't need to expose properties in a struct. Public (uppercase) is optional
- Perk: No reflection used, no struct tags needed
- Perk: Easy to lace in conditional values for variable binary streams
- Con: Not always super intuitive where a failure occured, since no context of which property failed like with binary.Read/Write
- Con: Always sanitize default value cases, or a panic may occur with returned values

Example usage (can be seen as a test [here](/example_test.go))
```go

import (
    "encoding/binary"
    "github.com/xackery/encdec"
    "fmt"
)


type exampleStruct struct {
	val1           int16
	val2           uint32
	someStr1       string
	someStr2       string
	someStrZero    string
	val3           float32
	someBytes      []byte
	someSubStructs []exampleSubStruct
}

type exampleSubStruct struct {
	val1 bool
	val2 float64
	val3 exampleVector3
}

type exampleVector3 struct {
	x float32
	y float32
	z float32
}


func someReadSeekerExample(r io.ReadSeeker) (*exampleStruct, error) {
	dec := NewDecoder(r, binary.LittleEndian)

	def := &exampleStruct{} // initialize an example struct in def

	def.val1 = dec.Int16()                     // decode int16 worth of bytes to val1
	def.val2 = dec.Uint32()                    // decode uint16 worth of bytes to val2
	def.someStr1 = dec.StringFixed(3)          // read 3 bytes and convert to a string
	def.someStr2 = dec.StringLenPrefixUint32() // read 4 bytes (uint32) to sort length of string, then read length and convert to string
	def.someStrZero = dec.StringZero()         // read until 0x00 and convert to string
	def.val3 = dec.Float32()                   // read 4 bytes and convertt float and place into val3
	def.someBytes = dec.Bytes(4)               // read 4 bytes and place into someBytes

	subStructLen := dec.Uint32() // read 4 bytes (uint32) to sort length of sub struct
	for i := 0; i < int(subStructLen); i++ {
		subStruct := exampleSubStruct{}
		subStruct.val1 = dec.Bool()      // read 1 byte and convert to bool
		subStruct.val2 = dec.Float64()   // read 8 bytes and convert to float64
		subStruct.val3.x = dec.Float32() // read 4 bytes and convert to float32
		subStruct.val3.y = dec.Float32() // read 4 bytes and convert to float32
		subStruct.val3.z = dec.Float32() // read 4 bytes and convert to float32

		def.someSubStructs = append(def.someSubStructs, subStruct) // append sub struct to someSubStructs
	}

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
	enc.Bytes(def.someBytes)                // write 4 bytes and place into someBytes

	enc.Uint32(uint32(len(def.someSubStructs))) // write 4 bytes (uint32) to sort length of sub struct
	for _, subStruct := range def.someSubStructs {
		enc.Bool(subStruct.val1)      // write 1 byte and convert to bool
		enc.Float64(subStruct.val2)   // write 8 bytes and convert to float64
		enc.Float32(subStruct.val3.x) // write 4 bytes and convert to float32
		enc.Float32(subStruct.val3.y) // write 4 bytes and convert to float32
		enc.Float32(subStruct.val3.z) // write 4 bytes and convert to float32
	}

	if enc.Error() != nil {
		return fmt.Errorf("encode: %w", enc.Error())
	}
	return nil
}
```