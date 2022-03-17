package reader

import (
	"encoding/binary"
	"io"
)

type Reader struct {
	io.Reader
}

func (r *Reader) ReadU1() (uint8, error) {
	var out uint8
	err := binary.Read(r, binary.BigEndian, &out)
	if err != nil {
		return 0, err
	}
	return out, nil
}

func (r *Reader) ReadU2() (uint16, error) {
	var out uint16
	err := binary.Read(r, binary.BigEndian, &out)
	if err != nil {
		return 0, err
	}
	return out, nil
}

func (r *Reader) ReadU4() (uint16, error) {
	var out uint16
	err := binary.Read(r, binary.BigEndian, &out)
	if err != nil {
		return 0, err
	}
	return out, nil
}
