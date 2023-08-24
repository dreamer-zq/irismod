// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package farm

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_GenesisState_2_list)(nil)

type _GenesisState_2_list struct {
	list *[]*FarmPool
}

func (x *_GenesisState_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*FarmPool)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*FarmPool)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_2_list) AppendMutable() protoreflect.Value {
	v := new(FarmPool)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_2_list) NewElement() protoreflect.Value {
	v := new(FarmPool)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_2_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_3_list)(nil)

type _GenesisState_3_list struct {
	list *[]*FarmInfo
}

func (x *_GenesisState_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*FarmInfo)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*FarmInfo)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_3_list) AppendMutable() protoreflect.Value {
	v := new(FarmInfo)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_3_list) NewElement() protoreflect.Value {
	v := new(FarmInfo)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_3_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_5_list)(nil)

type _GenesisState_5_list struct {
	list *[]*EscrowInfo
}

func (x *_GenesisState_5_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_5_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_5_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*EscrowInfo)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_5_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*EscrowInfo)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_5_list) AppendMutable() protoreflect.Value {
	v := new(EscrowInfo)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_5_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_5_list) NewElement() protoreflect.Value {
	v := new(EscrowInfo)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_5_list) IsValid() bool {
	return x.list != nil
}

var (
	md_GenesisState            protoreflect.MessageDescriptor
	fd_GenesisState_params     protoreflect.FieldDescriptor
	fd_GenesisState_pools      protoreflect.FieldDescriptor
	fd_GenesisState_farm_infos protoreflect.FieldDescriptor
	fd_GenesisState_sequence   protoreflect.FieldDescriptor
	fd_GenesisState_escrow     protoreflect.FieldDescriptor
)

func init() {
	file_irismod_farm_genesis_proto_init()
	md_GenesisState = File_irismod_farm_genesis_proto.Messages().ByName("GenesisState")
	fd_GenesisState_params = md_GenesisState.Fields().ByName("params")
	fd_GenesisState_pools = md_GenesisState.Fields().ByName("pools")
	fd_GenesisState_farm_infos = md_GenesisState.Fields().ByName("farm_infos")
	fd_GenesisState_sequence = md_GenesisState.Fields().ByName("sequence")
	fd_GenesisState_escrow = md_GenesisState.Fields().ByName("escrow")
}

var _ protoreflect.Message = (*fastReflection_GenesisState)(nil)

type fastReflection_GenesisState GenesisState

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GenesisState)(x)
}

func (x *GenesisState) slowProtoReflect() protoreflect.Message {
	mi := &file_irismod_farm_genesis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GenesisState_messageType fastReflection_GenesisState_messageType
var _ protoreflect.MessageType = fastReflection_GenesisState_messageType{}

type fastReflection_GenesisState_messageType struct{}

func (x fastReflection_GenesisState_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GenesisState)(nil)
}
func (x fastReflection_GenesisState_messageType) New() protoreflect.Message {
	return new(fastReflection_GenesisState)
}
func (x fastReflection_GenesisState_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GenesisState
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GenesisState) Descriptor() protoreflect.MessageDescriptor {
	return md_GenesisState
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GenesisState) Type() protoreflect.MessageType {
	return _fastReflection_GenesisState_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GenesisState) New() protoreflect.Message {
	return new(fastReflection_GenesisState)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GenesisState) Interface() protoreflect.ProtoMessage {
	return (*GenesisState)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GenesisState) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Params != nil {
		value := protoreflect.ValueOfMessage(x.Params.ProtoReflect())
		if !f(fd_GenesisState_params, value) {
			return
		}
	}
	if len(x.Pools) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_2_list{list: &x.Pools})
		if !f(fd_GenesisState_pools, value) {
			return
		}
	}
	if len(x.FarmInfos) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_3_list{list: &x.FarmInfos})
		if !f(fd_GenesisState_farm_infos, value) {
			return
		}
	}
	if x.Sequence != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Sequence)
		if !f(fd_GenesisState_sequence, value) {
			return
		}
	}
	if len(x.Escrow) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_5_list{list: &x.Escrow})
		if !f(fd_GenesisState_escrow, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_GenesisState) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "irismod.farm.GenesisState.params":
		return x.Params != nil
	case "irismod.farm.GenesisState.pools":
		return len(x.Pools) != 0
	case "irismod.farm.GenesisState.farm_infos":
		return len(x.FarmInfos) != 0
	case "irismod.farm.GenesisState.sequence":
		return x.Sequence != uint64(0)
	case "irismod.farm.GenesisState.escrow":
		return len(x.Escrow) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: irismod.farm.GenesisState"))
		}
		panic(fmt.Errorf("message irismod.farm.GenesisState does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "irismod.farm.GenesisState.params":
		x.Params = nil
	case "irismod.farm.GenesisState.pools":
		x.Pools = nil
	case "irismod.farm.GenesisState.farm_infos":
		x.FarmInfos = nil
	case "irismod.farm.GenesisState.sequence":
		x.Sequence = uint64(0)
	case "irismod.farm.GenesisState.escrow":
		x.Escrow = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: irismod.farm.GenesisState"))
		}
		panic(fmt.Errorf("message irismod.farm.GenesisState does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GenesisState) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "irismod.farm.GenesisState.params":
		value := x.Params
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "irismod.farm.GenesisState.pools":
		if len(x.Pools) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_2_list{})
		}
		listValue := &_GenesisState_2_list{list: &x.Pools}
		return protoreflect.ValueOfList(listValue)
	case "irismod.farm.GenesisState.farm_infos":
		if len(x.FarmInfos) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_3_list{})
		}
		listValue := &_GenesisState_3_list{list: &x.FarmInfos}
		return protoreflect.ValueOfList(listValue)
	case "irismod.farm.GenesisState.sequence":
		value := x.Sequence
		return protoreflect.ValueOfUint64(value)
	case "irismod.farm.GenesisState.escrow":
		if len(x.Escrow) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_5_list{})
		}
		listValue := &_GenesisState_5_list{list: &x.Escrow}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: irismod.farm.GenesisState"))
		}
		panic(fmt.Errorf("message irismod.farm.GenesisState does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "irismod.farm.GenesisState.params":
		x.Params = value.Message().Interface().(*Params)
	case "irismod.farm.GenesisState.pools":
		lv := value.List()
		clv := lv.(*_GenesisState_2_list)
		x.Pools = *clv.list
	case "irismod.farm.GenesisState.farm_infos":
		lv := value.List()
		clv := lv.(*_GenesisState_3_list)
		x.FarmInfos = *clv.list
	case "irismod.farm.GenesisState.sequence":
		x.Sequence = value.Uint()
	case "irismod.farm.GenesisState.escrow":
		lv := value.List()
		clv := lv.(*_GenesisState_5_list)
		x.Escrow = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: irismod.farm.GenesisState"))
		}
		panic(fmt.Errorf("message irismod.farm.GenesisState does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "irismod.farm.GenesisState.params":
		if x.Params == nil {
			x.Params = new(Params)
		}
		return protoreflect.ValueOfMessage(x.Params.ProtoReflect())
	case "irismod.farm.GenesisState.pools":
		if x.Pools == nil {
			x.Pools = []*FarmPool{}
		}
		value := &_GenesisState_2_list{list: &x.Pools}
		return protoreflect.ValueOfList(value)
	case "irismod.farm.GenesisState.farm_infos":
		if x.FarmInfos == nil {
			x.FarmInfos = []*FarmInfo{}
		}
		value := &_GenesisState_3_list{list: &x.FarmInfos}
		return protoreflect.ValueOfList(value)
	case "irismod.farm.GenesisState.escrow":
		if x.Escrow == nil {
			x.Escrow = []*EscrowInfo{}
		}
		value := &_GenesisState_5_list{list: &x.Escrow}
		return protoreflect.ValueOfList(value)
	case "irismod.farm.GenesisState.sequence":
		panic(fmt.Errorf("field sequence of message irismod.farm.GenesisState is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: irismod.farm.GenesisState"))
		}
		panic(fmt.Errorf("message irismod.farm.GenesisState does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GenesisState) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "irismod.farm.GenesisState.params":
		m := new(Params)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "irismod.farm.GenesisState.pools":
		list := []*FarmPool{}
		return protoreflect.ValueOfList(&_GenesisState_2_list{list: &list})
	case "irismod.farm.GenesisState.farm_infos":
		list := []*FarmInfo{}
		return protoreflect.ValueOfList(&_GenesisState_3_list{list: &list})
	case "irismod.farm.GenesisState.sequence":
		return protoreflect.ValueOfUint64(uint64(0))
	case "irismod.farm.GenesisState.escrow":
		list := []*EscrowInfo{}
		return protoreflect.ValueOfList(&_GenesisState_5_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: irismod.farm.GenesisState"))
		}
		panic(fmt.Errorf("message irismod.farm.GenesisState does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GenesisState) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in irismod.farm.GenesisState", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GenesisState) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_GenesisState) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GenesisState) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GenesisState)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Params != nil {
			l = options.Size(x.Params)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Pools) > 0 {
			for _, e := range x.Pools {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.FarmInfos) > 0 {
			for _, e := range x.FarmInfos {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Sequence != 0 {
			n += 1 + runtime.Sov(uint64(x.Sequence))
		}
		if len(x.Escrow) > 0 {
			for _, e := range x.Escrow {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*GenesisState)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Escrow) > 0 {
			for iNdEx := len(x.Escrow) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Escrow[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x2a
			}
		}
		if x.Sequence != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Sequence))
			i--
			dAtA[i] = 0x20
		}
		if len(x.FarmInfos) > 0 {
			for iNdEx := len(x.FarmInfos) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.FarmInfos[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x1a
			}
		}
		if len(x.Pools) > 0 {
			for iNdEx := len(x.Pools) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Pools[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x12
			}
		}
		if x.Params != nil {
			encoded, err := options.Marshal(x.Params)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*GenesisState)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Params == nil {
					x.Params = &Params{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Params); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pools", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Pools = append(x.Pools, &FarmPool{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Pools[len(x.Pools)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field FarmInfos", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.FarmInfos = append(x.FarmInfos, &FarmInfo{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.FarmInfos[len(x.FarmInfos)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
				}
				x.Sequence = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Sequence |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Escrow", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Escrow = append(x.Escrow, &EscrowInfo{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Escrow[len(x.Escrow)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: irismod/farm/genesis.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GenesisState defines the genesis information exported by the farm module
type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params    *Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	Pools     []*FarmPool   `protobuf:"bytes,2,rep,name=pools,proto3" json:"pools,omitempty"`
	FarmInfos []*FarmInfo   `protobuf:"bytes,3,rep,name=farm_infos,json=farmInfos,proto3" json:"farm_infos,omitempty"`
	Sequence  uint64        `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Escrow    []*EscrowInfo `protobuf:"bytes,5,rep,name=escrow,proto3" json:"escrow,omitempty"`
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_irismod_farm_genesis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

// Deprecated: Use GenesisState.ProtoReflect.Descriptor instead.
func (*GenesisState) Descriptor() ([]byte, []int) {
	return file_irismod_farm_genesis_proto_rawDescGZIP(), []int{0}
}

func (x *GenesisState) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *GenesisState) GetPools() []*FarmPool {
	if x != nil {
		return x.Pools
	}
	return nil
}

func (x *GenesisState) GetFarmInfos() []*FarmInfo {
	if x != nil {
		return x.FarmInfos
	}
	return nil
}

func (x *GenesisState) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *GenesisState) GetEscrow() []*EscrowInfo {
	if x != nil {
		return x.Escrow
	}
	return nil
}

var File_irismod_farm_genesis_proto protoreflect.FileDescriptor

var file_irismod_farm_genesis_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x72, 0x69, 0x73, 0x6d, 0x6f, 0x64, 0x2f, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x67,
	0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x69, 0x72,
	0x69, 0x73, 0x6d, 0x6f, 0x64, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x69, 0x72, 0x69, 0x73, 0x6d, 0x6f, 0x64, 0x2f, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x66,
	0x61, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x02, 0x0a, 0x0c, 0x47, 0x65,
	0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x69, 0x72, 0x69,
	0x73, 0x6d, 0x6f, 0x64, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x32,
	0x0a, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x69, 0x72, 0x69, 0x73, 0x6d, 0x6f, 0x64, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x46, 0x61, 0x72,
	0x6d, 0x50, 0x6f, 0x6f, 0x6c, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x05, 0x70, 0x6f, 0x6f,
	0x6c, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x66, 0x61, 0x72, 0x6d, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x72, 0x69, 0x73, 0x6d, 0x6f, 0x64,
	0x2e, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x46, 0x61, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x04,
	0xc8, 0xde, 0x1f, 0x00, 0x52, 0x09, 0x66, 0x61, 0x72, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x65,
	0x73, 0x63, 0x72, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x72,
	0x69, 0x73, 0x6d, 0x6f, 0x64, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x45, 0x73, 0x63, 0x72, 0x6f,
	0x77, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x65, 0x73, 0x63,
	0x72, 0x6f, 0x77, 0x42, 0x9e, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x72, 0x69, 0x73,
	0x6d, 0x6f, 0x64, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x42, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x72, 0x69, 0x73, 0x6e, 0x65, 0x74, 0x2f, 0x69, 0x72, 0x69,
	0x73, 0x6d, 0x6f, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x72, 0x69, 0x73, 0x6d, 0x6f, 0x64,
	0x2f, 0x66, 0x61, 0x72, 0x6d, 0xa2, 0x02, 0x03, 0x49, 0x46, 0x58, 0xaa, 0x02, 0x0c, 0x49, 0x72,
	0x69, 0x73, 0x6d, 0x6f, 0x64, 0x2e, 0x46, 0x61, 0x72, 0x6d, 0xca, 0x02, 0x0c, 0x49, 0x72, 0x69,
	0x73, 0x6d, 0x6f, 0x64, 0x5c, 0x46, 0x61, 0x72, 0x6d, 0xe2, 0x02, 0x18, 0x49, 0x72, 0x69, 0x73,
	0x6d, 0x6f, 0x64, 0x5c, 0x46, 0x61, 0x72, 0x6d, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x49, 0x72, 0x69, 0x73, 0x6d, 0x6f, 0x64, 0x3a, 0x3a,
	0x46, 0x61, 0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_irismod_farm_genesis_proto_rawDescOnce sync.Once
	file_irismod_farm_genesis_proto_rawDescData = file_irismod_farm_genesis_proto_rawDesc
)

func file_irismod_farm_genesis_proto_rawDescGZIP() []byte {
	file_irismod_farm_genesis_proto_rawDescOnce.Do(func() {
		file_irismod_farm_genesis_proto_rawDescData = protoimpl.X.CompressGZIP(file_irismod_farm_genesis_proto_rawDescData)
	})
	return file_irismod_farm_genesis_proto_rawDescData
}

var file_irismod_farm_genesis_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_irismod_farm_genesis_proto_goTypes = []interface{}{
	(*GenesisState)(nil), // 0: irismod.farm.GenesisState
	(*Params)(nil),       // 1: irismod.farm.Params
	(*FarmPool)(nil),     // 2: irismod.farm.FarmPool
	(*FarmInfo)(nil),     // 3: irismod.farm.FarmInfo
	(*EscrowInfo)(nil),   // 4: irismod.farm.EscrowInfo
}
var file_irismod_farm_genesis_proto_depIdxs = []int32{
	1, // 0: irismod.farm.GenesisState.params:type_name -> irismod.farm.Params
	2, // 1: irismod.farm.GenesisState.pools:type_name -> irismod.farm.FarmPool
	3, // 2: irismod.farm.GenesisState.farm_infos:type_name -> irismod.farm.FarmInfo
	4, // 3: irismod.farm.GenesisState.escrow:type_name -> irismod.farm.EscrowInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_irismod_farm_genesis_proto_init() }
func file_irismod_farm_genesis_proto_init() {
	if File_irismod_farm_genesis_proto != nil {
		return
	}
	file_irismod_farm_farm_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_irismod_farm_genesis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_irismod_farm_genesis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_irismod_farm_genesis_proto_goTypes,
		DependencyIndexes: file_irismod_farm_genesis_proto_depIdxs,
		MessageInfos:      file_irismod_farm_genesis_proto_msgTypes,
	}.Build()
	File_irismod_farm_genesis_proto = out.File
	file_irismod_farm_genesis_proto_rawDesc = nil
	file_irismod_farm_genesis_proto_goTypes = nil
	file_irismod_farm_genesis_proto_depIdxs = nil
}