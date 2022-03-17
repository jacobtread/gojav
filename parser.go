package gojav

import (
	"errors"
	. "github.com/jacobtread/gojav/tools"
	"log"
)

func ParseClassFile(r Reader) (*ClassFile, error) {
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

	return &out, nil
}

func ReadConstantPool(r Reader) (*ConstantPool, error) {
	count, err := r.ReadU2()
	if err != nil {
		return nil, err
	}
	nextPass := make([]BitSet, 3)

	pool := &ConstantPool{
		Size: count,
		Pool: make([]*ConstantPoolEntry, count-1),
	}

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
			pool.Pool[i] = nil
		case DoubleConstant:
			entry.Value, err = r.ReadDouble()
			i++
			pool.Pool[i] = nil
		case ClassConstant:
		case StringConstant:
		case MethodTypeConstant:
			value, err := r.ReadU2()
			if err != nil {
				return nil, err
			}
			entry.Value = &IndexConstant{Tag: tag, Index: value}
			nextPass[0].Set(i)
		case FieldrefConstant:
		case MethodrefConstant:
		case InterfaceMethodrefConstant:
		case InvokeDynamicConstant:
			a, err := r.ReadU2()
			if err != nil {
				return nil, err
			}
			b, err := r.ReadU2()
			if err != nil {
				return nil, err
			}
			entry.Value = &LinkConstant{Tag: tag, Index: a, NameAndTypeIndex: b}
			nextPass[1].Set(i)
		case MethodHandleConstant:
			a, err := r.ReadU2()
			if err != nil {
				return nil, err
			}
			b, err := r.ReadU2()
			if err != nil {
				return nil, err
			}
			entry.Value = &LinkConstant{Tag: tag, Index: a, NameAndTypeIndex: b}
			nextPass[2].Set(i)
		default:
			log.Panicf("Don't know how to handle tag %d\n", tag)
		}
		if err != nil {
			return nil, err
		}
		pool.Pool[i] = &entry
	}
	var x uint16 = 0
	next := true
	for _, set := range nextPass {
		for next {
			x, next = set.NextValue(x)
			if !next {
				continue
			}
			entry := pool.Pool[x]
			value := entry.Value
			switch value.(type) {
			case *IndexConstant:
				entry.Value, err = value.(*IndexConstant).Resolve(pool)
			case *LinkConstant:
				entry.Value, err = value.(*LinkConstant).Resolve(pool)
			}
			if err != nil {
				return nil, err
			}
			x += 1
		}
	}
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func (c *ConstantPool) GetStringEntry(index uint16) (string, error) {
	value := c.Pool[index]
	if value.Tag == StringConstant {
		return value.Value.(string), nil
	} else {
		return "", nil
	}
}

func (c *ConstantPool) GetLinkConstant(index uint16) (*RLinkConstant, error) {
	entry := c.Pool[index]
	value := entry.Value
	switch value.(type) {
	case *LinkConstant:
		return value.(*LinkConstant).Resolve(c)
	case *RLinkConstant:
		return value.(*RLinkConstant), nil
	default:
		return nil, errors.New("not a link constant")
	}
}

func (l *IndexConstant) Resolve(pool *ConstantPool) (*RIndexConstant, error) {
	switch l.Tag {
	case ClassConstant:
	case StringConstant:
	case MethodTypeConstant:
		out := RIndexConstant{Tag: l.Tag}
		var err error
		out.Value, err = pool.GetStringEntry(l.Index)
		if err != nil {
			return nil, err
		} else {
			return &out, nil
		}
	}
	return nil, errors.New("invalid tag type")
}

func (l *LinkConstant) Resolve(pool *ConstantPool) (*RLinkConstant, error) {
	var err error
	out := &RLinkConstant{}
	switch l.Tag {
	case NameAndTypeConstant:
		out.ElementName, err = pool.GetStringEntry(l.Index)
		if err != nil {
			return nil, err
		}
		out.Descriptor, err = pool.GetStringEntry(l.NameAndTypeIndex)
		if err != nil {
			return nil, err
		}
	case MethodHandleConstant:
		var value *RLinkConstant
		value, err = pool.GetLinkConstant(l.NameAndTypeIndex)
		if err != nil {
			return nil, err
		}
		out.ClassName = value.ClassName
		out.ElementName = value.ElementName
		out.Descriptor = value.Descriptor
	default:
		if l.Tag != InvokeDynamicConstant {
			out.ClassName, err = pool.GetStringEntry(l.Index)
			if err != nil {
				return nil, err
			}
		}
		var value *RLinkConstant
		value, err = pool.GetLinkConstant(l.NameAndTypeIndex)
		if err != nil {
			return nil, err
		}
		out.ElementName = value.ElementName
		out.Descriptor = value.Descriptor
	}
	return out, nil
}
