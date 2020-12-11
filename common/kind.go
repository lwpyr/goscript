package common

type DataKind struct {
	// maybe this struct will be useful, so I just keep it here
	Kind int // int/float/message/slice/map/...
}

const ( // Kind
	UInt8 = iota
	UInt32
	UInt64
	Int32
	Int64
	Float32
	Float64
	Bool
	Bytes
	String
	Slice
	Map
	Message
	Channel
	Nil
	Enum
	Func
	ReflectType
	Any
)

var KindMap = map[int]*DataKind{
	UInt8: {
		Kind: UInt8,
	},
	Int32: {
		Kind: Int32,
	},
	Int64: {
		Kind: Int64,
	},
	UInt32: {
		Kind: UInt32,
	},
	UInt64: {
		Kind: UInt64,
	},
	Float32: {
		Kind: Float32,
	},
	Float64: {
		Kind: Float64,
	},
	Bool: {
		Kind: Bool,
	},
	String: {
		Kind: String,
	},
	Bytes: {
		Kind: Bytes,
	},
	Slice: {
		Kind: Slice,
	},
	Map: {
		Kind: Map,
	},
	Message: {
		Kind: Message,
	},
	Channel: {
		Kind: Channel,
	},
	Nil: {
		Kind: Nil,
	},
	Enum: {
		Kind: Enum,
	},
	Func: {
		Kind: Func,
	},
	ReflectType: {
		Kind: ReflectType,
	},
	Any: {
		Kind: Any,
	},
}
