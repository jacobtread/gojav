package gojav

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

func (r *Reader) ReadU1To(out *uint8) error {
	err := binary.Read(r, binary.BigEndian, &out)
	if err != nil {
		return err
	}
	return nil
}

func (r *Reader) ReadU2() (uint16, error) {
	var out uint16
	err := binary.Read(r, binary.BigEndian, &out)
	if err != nil {
		return 0, err
	}
	return out, nil
}

func (r *Reader) ReadU2To(out *uint16) error {
	err := binary.Read(r, binary.BigEndian, &out)
	if err != nil {
		return err
	}
	return nil
}

func (r *Reader) ReadU4() (uint32, error) {
	var out uint32
	err := binary.Read(r, binary.BigEndian, &out)
	if err != nil {
		return 0, err
	}
	return out, nil
}

func (r *Reader) ReadU4To(out *uint32) error {
	err := binary.Read(r, binary.BigEndian, &out)
	if err != nil {
		return err
	}
	return nil
}

func (r *Reader) ReadFloat() (float32, error) {
	var out float32
	err := binary.Read(r, binary.BigEndian, &out)
	if err != nil {
		return 0, err
	}
	return out, nil
}

func (r *Reader) ReadFloatTo(out *float32) error {
	err := binary.Read(r, binary.BigEndian, &out)
	if err != nil {
		return err
	}
	return nil
}

func (r *Reader) ReadDouble() (float64, error) {
	var out float64
	err := binary.Read(r, binary.BigEndian, &out)
	if err != nil {
		return 0, err
	}
	return out, nil
}

func (r *Reader) ReadDoubleTo(out *float64) error {
	err := binary.Read(r, binary.BigEndian, &out)
	if err != nil {
		return err
	}
	return nil
}

func (r *Reader) ReadLong() (int64, error) {
	var out int64
	err := binary.Read(r, binary.BigEndian, &out)
	if err != nil {
		return 0, err
	}
	return out, nil
}

func (r *Reader) ReadLongTo(out *int64) error {
	err := binary.Read(r, binary.BigEndian, &out)
	if err != nil {
		return err
	}
	return nil
}

func (r *Reader) ReadUTF8() (string, error) {
	length, err := r.ReadU2()
	if err != nil {
		return "", err
	}
	bytes := make([]byte, length)
	_, err = io.ReadFull(r, bytes)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (r *Reader) Discard(length int64) {
	// Skip over the unused bytes
	_, _ = io.CopyN(io.Discard, r, length)
}
