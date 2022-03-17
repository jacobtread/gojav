package tools

const size = 16

type bits uint16

// BitSet is a set of bits that can be set, cleared and queried.
type BitSet []bits

func (s *BitSet) Set(i uint16) {
	if len(*s) < int(i/size+1) {
		r := make([]bits, i/size+1)
		copy(r, *s)
		*s = r
	}
	(*s)[i/size] |= 1 << (i % size)
}

func (s *BitSet) Clear(i uint16) {
	if len(*s) >= int(i/size+1) {
		(*s)[i/size] &^= 1 << (i % size)
	}
}

func (s *BitSet) NextValue(v uint16) (uint16, bool) {
	if len(*s) >= int(v) {
		a := (*s)[v]
		return uint16(a), true
	} else {
		return 0, false
	}
}

func (s *BitSet) IsSet(i uint16) bool {
	return (*s)[i/size]&(1<<(i%size)) != 0
}
