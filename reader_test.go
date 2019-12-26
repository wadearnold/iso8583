package iso8583

import (
	"bytes"
	"testing"
)

var presentment = []byte("020042000400000000021612345678901234560609173030123456789ABC1000123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789")

func TestRead(t *testing.T) {
	r := NewReader(bytes.NewReader(presentment))
	msg, _ := r.Read()
	if msg.MTI != "0200" {
		t.Errorf("Expected MTI 0200 got: %s", msg.MTI)
	}
	if msg.PrimaryBitmap != "4200040000000002" {
		t.Errorf("Expected MTI 4200040000000002 got: %s", msg.PrimaryBitmap)
	}
	if len(msg.Fields) != 4 {
		t.Errorf("Expected 4 fields got: %v \n", len(msg.Fields))
	}
}

// TODO test for SecondaryBitmap, TertiaryBitmap

func TestHexToBin(t *testing.T) {
	primaryBitmap := "4200040000000002"
	binaryResult := "0100001000000000000001000000000000000000000000000000000000000010"

	result, _ := HexToBin(primaryBitmap)
	if binaryResult != result {
		t.Errorf("Expected HexToBinary failed got: %s", result)
	}
}
