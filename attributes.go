package gojav

const (
	CodeAttr          = "Code"
	InnerClassesAttr  = "InnerClasses"
	ConstantValueAttr = "ConstantValue"
	SignatureAttr     = "Signature"
)

func (r Reader) ReadAttributes(pool *ConstantPool) (*Attributes, error) {
	count, err := r.ReadU2()
	if err != nil {
		return nil, err
	}
	out := &Attributes{
		Size:       count,
		Attributes: make([]any, count),
	}

	var nameIndex uint16
	var length uint32
	var name string
	for i := uint16(0); i < count; i++ {
		nameIndex, err = r.ReadU2()
		if err != nil {
			return nil, err
		}
		length, err = r.ReadU4()
		if err != nil {
			return nil, err
		}
		name, err = pool.GetStringEntry(nameIndex)
		var value any = nil
		switch name {
		case CodeAttr:
			value, err = r.ReadCodeAttribute(pool)
		case InnerClassesAttr:
			value, err = r.ReadInnerClassesAttribute(pool)
		case ConstantValueAttr:
			value, err = r.ReadConstantValueAttribute()
		case SignatureAttr:
			value, err = r.ReadSignatureAttribute(pool)
		default:
			r.Discard(int64(length))
		}
		if err != nil {
			return nil, err
		}
		out.Attributes[i] = value
	}
	return out, nil
}

type CodeAttribute struct {
	LocalVariables uint16
	CodeLength     uint32
	CodeFullLength uint16
	Attributes     *Attributes
}

func (r Reader) ReadCodeAttribute(pool *ConstantPool) (*CodeAttribute, error) {
	r.Discard(2)
	out := CodeAttribute{}
	err := r.ReadU2To(&out.LocalVariables)
	if err != nil {
		return nil, err
	}
	err = r.ReadU4To(&out.CodeLength)
	if err != nil {
		return nil, err
	}
	r.Discard(int64(out.CodeLength))
	execLength, err := r.ReadU2()
	if err != nil {
		return nil, err
	}
	execLength *= 8
	r.Discard(int64(execLength))
	out.CodeFullLength = out.CodeFullLength + execLength + 2
	out.Attributes, err = r.ReadAttributes(pool)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type InnerClassesAttribute struct {
	Size    uint16
	Entries []InnerClassesAttributeEntry
}

type InnerClassesAttributeEntry struct {
	InnerNameIndex  uint16
	OuterNameIndex  uint16
	SimpleNameIndex uint16
	AccessFlags     uint16

	InnerName  string
	OuterName  string
	SimpleName string
}

func (r Reader) ReadInnerClassesAttribute(pool *ConstantPool) (*InnerClassesAttribute, error) {
	count, err := r.ReadU2()
	if err != nil {
		return nil, err
	}
	out := InnerClassesAttribute{
		Size:    count,
		Entries: make([]InnerClassesAttributeEntry, count),
	}

	for i := uint16(0); i < count; i++ {
		entry := InnerClassesAttributeEntry{}
		err = r.ReadU2To(&entry.InnerNameIndex)
		if err != nil {
			return nil, err
		}
		err = r.ReadU2To(&entry.OuterNameIndex)
		if err != nil {
			return nil, err
		}
		err = r.ReadU2To(&entry.SimpleNameIndex)
		if err != nil {
			return nil, err
		}
		err = r.ReadU2To(&entry.AccessFlags)
		if err != nil {
			return nil, err
		}

		entry.SimpleName, err = pool.GetStringEntry(entry.InnerNameIndex)
		if err != nil {
			return nil, err
		}
		if entry.OuterNameIndex != 0 {
			entry.OuterName, err = pool.GetStringEntry(entry.OuterNameIndex)
			if err != nil {
				return nil, err
			}
		}
		if entry.SimpleNameIndex != 0 {
			entry.SimpleName, err = pool.GetStringEntry(entry.SimpleNameIndex)
			if err != nil {
				return nil, err
			}
		}
		out.Entries[i] = entry
	}
	return &out, nil
}

type ConstantValueAttribute struct {
	Index uint16
}

func (r Reader) ReadConstantValueAttribute() (*ConstantValueAttribute, error) {
	value, err := r.ReadU2()
	if err != nil {
		return nil, err
	}
	return &ConstantValueAttribute{Index: value}, nil
}

type SignatureAttribute struct {
	Signature string
}

func (r Reader) ReadSignatureAttribute(pool *ConstantPool) (*SignatureAttribute, error) {
	index, err := r.ReadU2()
	if err != nil {
		return nil, err
	}
	signature, err := pool.GetStringEntry(index)
	if err != nil {
		return nil, err
	}
	return &SignatureAttribute{
		Signature: signature,
	}, nil
}
