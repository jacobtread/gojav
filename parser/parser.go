package parser

import (
	"github.com/jacobtread/gojav/reader"
	. "github.com/jacobtread/gojav/types"
)

func ParseClassFile(r reader.Reader) (*ClassFile, error) {
	out := ClassFile{}
	err := r.ReadU4To(&out.Magic)
	if err != nil {
		return nil, err
	}
	err = r.ReadU2To(&out.MinorVersion)
	if err != nil {
		return nil, err
	}
	err = r.ReadU2To(&out.MajorVersion)
	if err != nil {
		return nil, err
	}
	err = r.ReadU2To(&out.ConstantPoolCount)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func ReadConstantPool(r reader.Reader, class *ClassFile) (*ConstantPoolEntry, error) {
	count, err := r.ReadU2()
	if err != nil {
		return nil, err
	}
	class.ConstantPoolCount = count
	values := make([]*ConstantPoolEntry, count-1)
	var tag uint8
	for i := uint16(0); i < count; i++ {
		tag, err = r.ReadU1()
		if err != nil {
			return nil, err
		}
		entry := ConstantPoolEntry{Tag: tag}
		switch tag {
		case Utf8Constant:
			entry.Value, err = r.ReadUTF8()
		case IntegerConstant:
			entry.Value, err = r.ReadU4()
		case FloatConstant:
			entry.Value, err = r.ReadFloat()
		case LongConstant:
			entry.Value, err = r.ReadLong()
			i++
			values[i] = nil
		case DoubleConstant:
			entry.Value, err = r.ReadDouble()
			i++
			values[i] = nil
		}
		if err != nil {
			return nil, err
		}
		values[i] = &entry
	}

}
