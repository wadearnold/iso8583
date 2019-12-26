package iso8583

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// Reader for the ISO 8583 Messages
type Reader struct {
	Message   *Message
	MTIEncode string

	// The io.Reader sent to be parsed.
	scanner *bufio.Scanner
	// TODO errors

	// length of the data read
	len   int
	input []byte
}

// NewReader returns a new ISO 8583 Readers that reads from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		scanner: bufio.NewScanner(r),
		Message: NewMessage(),
	}
}

func (r *Reader) Read() (*Message, error) {
	for r.scanner.Scan() {
		r.input = r.scanner.Bytes()
		r.len = len(r.input)
		// TODO error if length is 0
	}
	if err := r.scanner.Err(); err != nil {
		// TODO return an error
		fmt.Printf("Could not read input: %s\n", err)
	}

	if err := r.parseMTI(); err != nil {
		return r.Message, err
	}
	if err := r.parseBitmap(); err != nil {
		return r.Message, err
	}

	return r.Message, nil
}

func (r *Reader) parseMTI() error {
	r.Message.MTI = string(r.input[:4])
	return nil
}

func (r *Reader) parseBitmap() error {
	r.Message.PrimaryBitmap = string(r.input[4:20])
	if r.Message.PrimaryBitmap[:1] == "1" {
		r.Message.SecondaryBitmap = string(r.input[21:38])
	}

	// convert Bitmap to binary and interact for what fields are in the payload
	binaryBitmap, _ := HexToBin(r.Message.PrimaryBitmap)
	// TODO deal with the error
	for i, c := range binaryBitmap {
		if c == rune('1') {
			r.Message.Fields[string(i)] = "nil"
		}
	}

	// add the Fields map

	return nil
}

// HextToBin takes a string of hex characters and converts them to their binary equivalent
func HexToBin(hex string) (string, error) {
	ui, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		return "", err
	}
	format := fmt.Sprintf("%%0%db", len(hex)*4)
	return fmt.Sprintf(format, ui), nil
}
