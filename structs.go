package gojav

const (
	ClassConstant              uint8 = 7
	FieldrefConstant                 = 9
	MethodrefConstant                = 10
	InterfaceMethodrefConstant       = 11
	StringConstant                   = 8
	IntegerConstant                  = 3
	FloatConstant                    = 4
	LongConstant                     = 5
	DoubleConstant                   = 6
	NameAndTypeConstant              = 12
	Utf8Constant                     = 1
	MethodHandleConstant             = 15
	MethodTypeConstant               = 16
	InvokeDynamicConstant            = 18
)

type ClassFile struct {
	Magic        uint32
	MinorVersion uint16
	MajorVersion uint16
	ConstantPool *ConstantPool
	AccessFlags  uint16
	ThisClass    uint16
	SuperClass   uint16
	Interfaces   *Interfaces

	FieldCount     uint16
	Fields         []*FieldInfo
	MethodCount    uint16
	Methods        []MethodInfo
	AttributeCount uint16
	Attributes     []AttributeInfo
}

type Interfaces struct {
	Size    uint16
	Indexes []uint16
	Names   []string
}

type ConstantPool struct {
	Size uint16
	Pool []*ConstantPoolEntry
}

type ConstantPoolEntry struct {
	Tag   uint8
	Value interface{}
}

type IndexConstant struct {
	Tag   uint8
	Index uint16
}

type RIndexConstant struct {
	Tag   uint8
	Value string
}

type LinkConstant struct {
	Tag              uint8
	Index            uint16
	NameAndTypeIndex uint16
}

type RLinkConstant struct {
	ClassName   string
	ElementName string
	Descriptor  string
}

type FloatInfo struct {
	Tag   uint8
	Bytes uint32
}

type DoubleInfo struct {
	Tag       uint8
	HighBytes uint32
	LowBytes  uint32
}

type LongInfo struct {
	Tag       uint8
	HighBytes uint32
	LowBytes  uint32
}

type Utf8Info struct {
	Tag    uint8
	Length uint16
	Bytes  []uint8
}

type MethodHandleInfo struct {
	Tag            uint8
	ReferenceKind  uint8
	ReferenceIndex uint16
}

type AccessFlag uint16

const (
	PublicFieldACC    AccessFlag = 0x0001
	PrivateFieldACC              = 0x0002
	ProtectedFieldACC            = 0x0004
	StaticFieldACC               = 0x0004
	FinalFieldACC                = 0x0008
	VolatileFieldACC             = 0x0010
	TransientFieldACC            = 0x0040
	SyntheticFieldACC            = 0x1000
	EnumFieldACC                 = 0x4000
)

type FieldInfo struct {
	AccessFlags     uint16 // See AccessFlag
	NameIndex       uint16
	DescriptorIndex uint16

	AttributeCount uint16
	Attributes     []AttributeInfo
}

type Attributes struct {
	Size       uint16
	Attributes []any
}

type Attribute uint8

type AttributeInfo struct {
	NameIndex uint16
	Length    uint32
	Info      []uint8
}

const (
	PublicMethodACC       AccessFlag = 0x0001
	PrivateMethodACC                 = 0x0002
	ProtectedMethodACC               = 0x0004
	StaticMethodACC                  = 0x0008
	FinalMethodACC                   = 0x0010
	SynchronizedMethodACC            = 0x0020
	BridgeMethodACC                  = 0x0040
	VarargsMethodACC                 = 0x0080
	NativeMethodACC                  = 0x0100
	AbstractMethodACC                = 0x0400
	StrictMethodACC                  = 0x0800
	SyntheticMethodACC               = 0x1000
)

type MethodInfo struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16

	AttributeCount uint16
	Attributes     []AttributeInfo
}
