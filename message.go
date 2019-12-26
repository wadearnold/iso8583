package iso8583

/*
An ISO8583 message is made of the following parts:
 - Message type indicator (MTI)
 - One or more bitmaps, indicating which data elements are present
 - Data elements, the actual information fields of the message
*/

// Message is a structure for an ISO  8583 message
type Message struct {
	// MTI Message type indicator (MTI)
	MTI string `json:"MTI"`

	// A bitmap is a field or subfield within a message, which indicates whether other data elements or data element subfields are present elsewhere in the message.
	// A field is considered to be present only when the corresponding bit in the bitmap is set.

	// Primary Bitmap indicates which of data elements 1 to 64 are present
	PrimaryBitmap string `json:"PrimaryBitmap"`
	// SecondaryBitmap indicates whether data elements 65 to 128 are present
	SecondaryBitmap string `json:"SecondaryBitmap,omitempty"`
	// TertiaryBitmap indicate the presence of fields 129 to 192, although these data elements are rarely used.
	TertiaryBitmap string `json:"TertiaryBitmap,omitempty"`
	// Fields are data elements of the individual fields carrying the transaction information. There are up to 128 data elements specified in the original ISO 8583:1987 standard, and up to 192 data elements in later releases
	Fields map[string]string `json:"Fields"`
}

// NewMessage constructs a Message with default values
func NewMessage() *Message {
	return &Message{
		Fields: make(map[string]string),
	}
}
