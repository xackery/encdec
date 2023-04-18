# encdec
Golang encoder/decoder from io.ReadSeeker and io.Writer

Please note this is not a very rigorous way to decode a io.ReadSeeker/Writer.

I use io.ReadSeeker so I can know the position of a failure.

TODO: add notes on how to convert io.Reader to io.ReadSeeker

`go get github.com/xackery/encdec`

Example usage (can be seen as a test [here](/example_test.go))
```go
type exampleStruct struct {
	val1        int16
	val2        uint32
	someStr1    string
	someStr2    string
	someStrZero string
	val3        float32
	someBytes   []byte
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
```

The downside to this approach is that if you hit EOF or an error early, many of these fields will return default values (0, empty strings, etc) and continue your logic, so be careful to write code that ensures default values are predicted and don't panic.

Under normal circumstances, using binary.Read() is preferred, but the error check on every single read gets cumbersome for large binary decodes, and is a bit overkill.

I also considered using struct tags, but they require reflection and are a pain when you need to call a function within a struct tag to do custom condition checking on future reads based on a value of a current read.

If you've used other langauges, this approach should be familiar.