package gojav

func (r Reader) ReadAttributes(pool *ConstantPool) (*Attributes, error) {
	count, err := r.ReadU2()
	if err != nil {
		return nil, err
	}
	out := &Attributes{
		Size:       count,
		Attributes: make([]*interface{}, count),
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
		case "Code":
			value, err = r.ReadCodeAttribute(pool)
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
