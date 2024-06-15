package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decode(io.Reader, *Message) error
}

type GOBDecoder struct {
}

func (dec GOBDecoder) Decode(reader io.Reader, val any) error {
	return gob.NewDecoder(reader).Decode(val)
}

type NOPDecoder struct {
}

func (dec NOPDecoder) Decode(reader io.Reader, msg *Message) error {
	buf := make([]byte, 1028)

	n, err := reader.Read(buf)

	if err != nil {
		return err
	}

	msg.Payload = buf[:n]

	return nil
}
