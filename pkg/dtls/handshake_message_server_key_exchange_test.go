package dtls

import (
	"reflect"
	"testing"
)

func TestHandshakeMessageServerKeyExchange(t *testing.T) {
	rawServerKeyExchange := []byte{
		0x03, 0x00, 0x17, 0x41, 0x04, 0x7f, 0x3a, 0xe0, 0xf8, 0x9c, 0xb2, 0xa4, 0x5e, 0x0e, 0x7c,
		0xb1, 0x2d, 0x7b, 0x3a, 0xe1, 0xbd, 0x39, 0x6c, 0x67, 0xb1, 0xea, 0xa2, 0xe1, 0xf3, 0x25,
		0x9c, 0xb0, 0x84, 0xb3, 0xc4, 0xbd, 0xe0, 0x21, 0xf9, 0x21, 0xef, 0xc3, 0xf3, 0xde, 0x19,
		0xec, 0xa8, 0x77, 0x80, 0xaa, 0x24, 0xcc, 0x83, 0x10, 0x6f, 0x5a, 0x55, 0x8a, 0x04, 0x41,
		0x9a, 0x47, 0x50, 0x1a, 0xd2, 0x9e, 0x29, 0xc1, 0xce, 0x02, 0x03, 0x00, 0x46, 0x30, 0x44,
		0x02, 0x20, 0x01, 0x71, 0xf7, 0x0c, 0xbb, 0xd0, 0x73, 0xbc, 0xeb, 0x80, 0x16, 0xe8, 0x36,
		0xe9, 0xe0, 0xa9, 0x69, 0x14, 0xf8, 0xc7, 0xae, 0x10, 0x1b, 0x1e, 0x86, 0x56, 0x65, 0xbc,
		0xf2, 0xb2, 0x4e, 0xde, 0x02, 0x20, 0x18, 0xfd, 0x33, 0x21, 0xc7, 0x3f, 0xf6, 0x5f, 0x82,
		0x72, 0x80, 0xc4, 0xd5, 0xd4, 0x83, 0xa0, 0xd5, 0x67, 0x62, 0x8f, 0x0e, 0xde, 0x77, 0xff,
		0x02, 0xa4, 0x04, 0x6a, 0x9b, 0x05, 0x89, 0x3a,
	}
	parsedServerKeyExchange := &handshakeMessageServerKeyExchange{}

	c := &handshakeMessageServerKeyExchange{}
	if err := c.unmarshal(rawServerKeyExchange); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(c, parsedServerKeyExchange) {
		t.Errorf("handshakeMessageServerKeyExchange unmarshal: got %#v, want %#v", c, parsedServerKeyExchange)
	}

	raw, err := c.marshal()
	if err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(raw, rawServerKeyExchange) {
		t.Errorf("handshakeMessageServerKeyExchange marshal: got %#v, want %#v", raw, rawServerKeyExchange)
	}
}