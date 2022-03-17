package parser

import (
	"encoding/binary"
	. "github.com/jacobtread/gojav/types"
	"io"
)

func ParseClassFile(r io.Reader) (*ClassFile, error) {
	out := ClassFile{}
	err := binary.Read(r, binary.BigEndian, &out.Magic)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.BigEndian, &out.MinorVersion)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.BigEndian, &out.MajorVersion)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.BigEndian, &out.ConstantPoolCount)
	if err != nil {
		return nil, err
	}

	return &out
}

func ReadConstantPool(r io.Reader, class *ClassFile) (*ConstantPoolEntry, error) {
	err := binary.Read(r, binary.BigEndian, &class.ConstantPoolCount)
	if err != nil {
		return nil, err
	}
	count := class.ConstantPoolCount

	values := make([]interface{}, count-1)

	var tag uint8
	for i := uint16(0); i < count; i++ {
		err = binary.Read(r, binary.BigEndian, &tag)
		if err != nil {
			return nil, err
		}

		switch tag {
		case Utf8Constant:
			values[i] =
		}
	}

}
