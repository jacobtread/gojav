package types

type (
	ConstantType uint16
)

const (
	ClassConstant              ConstantType = 7
	FieldrefConstant                        = 9
	MethodrefConstant                       = 10
	InterfaceMethodrefConstant              = 11
	StringConstant                          = 8
	IntegerConstant                         = 3
	FloatConstant                           = 4
	LongConstant                            = 5
	DoubleConstant                          = 6
	NameAndTypeConstant                     = 12
	Utf8Constant                            = 1
	MethodHandleConstant                    = 15
	MethodTypeConstant                      = 16
	InvokeDynamicConstant                   = 18
)

type ConstantPool struct {
	Tag  uint16
	Info []ConstantType
}

type ClassFile struct {
	Magic            uint32
	MinorVersion     uint16
	MajorVersion     uint16
	ConstantPoolSize uint16
	ConstantPool     ConstantPool
	AccessFlags      uint16
	ThisClass        uint16
	SuperClass       uint16
	InterfacesCount  uint16
}

type ClassInfo struct {
	Tag       uint8
	NameIndex uint16
}

type FieldrefInfo struct {
	Tag              uint8
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

type MethodrefInfo struct {
	Tag              uint8
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

type InterfaceMethodrefInfo struct {
	Tag              uint8
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

type StringInfo struct {
	Tag         uint8
	StringIndex uint16
}

type IntegerInfo struct {
	Tag   uint8
	Bytes uint32
}

type FloatInfo struct {
	Tag   uint8
	Bytes uint32
}

type LongInfo struct {
	Tag       uint8
	HighBytes uint32
	LowBytes  uint32
}

type DoubleInfo struct {
	Tag       uint8
	HighBytes uint32
	LowBytes  uint32
}

type NameAndTypeInfo struct {
	Tag             uint8
	NameIndex       uint16
	DescriptorIndex uint16
}

type UtfInfo struct {
	Tag    uint8
	Length uint16
	Bytes  []uint8
}

type MethodHandleInfo struct {
	Tag            uint8
	ReferenceKind  uint8
	ReferenceIndex uint16
}

type MethodTypeInfo struct {
	Tag             uint8
	DescriptorIndex uint16
}

type InvokeDynamicInfo struct {
	Tag                      uint8
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
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
	AttributesCount uint16
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
	AttributesCount uint16
	Info            []AttributeInfo
}
