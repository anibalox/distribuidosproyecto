// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: tarea.proto

package distribuidosproyecto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AyudaResp_Disponibilidad int32

const (
	AyudaResp_NO AyudaResp_Disponibilidad = 0
	AyudaResp_SI AyudaResp_Disponibilidad = 1
)

// Enum value maps for AyudaResp_Disponibilidad.
var (
	AyudaResp_Disponibilidad_name = map[int32]string{
		0: "NO",
		1: "SI",
	}
	AyudaResp_Disponibilidad_value = map[string]int32{
		"NO": 0,
		"SI": 1,
	}
)

func (x AyudaResp_Disponibilidad) Enum() *AyudaResp_Disponibilidad {
	p := new(AyudaResp_Disponibilidad)
	*p = x
	return p
}

func (x AyudaResp_Disponibilidad) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AyudaResp_Disponibilidad) Descriptor() protoreflect.EnumDescriptor {
	return file_tarea_proto_enumTypes[0].Descriptor()
}

func (AyudaResp_Disponibilidad) Type() protoreflect.EnumType {
	return &file_tarea_proto_enumTypes[0]
}

func (x AyudaResp_Disponibilidad) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AyudaResp_Disponibilidad.Descriptor instead.
func (AyudaResp_Disponibilidad) EnumDescriptor() ([]byte, []int) {
	return file_tarea_proto_rawDescGZIP(), []int{1, 0}
}

type SituacionResp_Estado int32

const (
	SituacionResp_PELIGRO  SituacionResp_Estado = 0
	SituacionResp_RESUELTO SituacionResp_Estado = 1
)

// Enum value maps for SituacionResp_Estado.
var (
	SituacionResp_Estado_name = map[int32]string{
		0: "PELIGRO",
		1: "RESUELTO",
	}
	SituacionResp_Estado_value = map[string]int32{
		"PELIGRO":  0,
		"RESUELTO": 1,
	}
)

func (x SituacionResp_Estado) Enum() *SituacionResp_Estado {
	p := new(SituacionResp_Estado)
	*p = x
	return p
}

func (x SituacionResp_Estado) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SituacionResp_Estado) Descriptor() protoreflect.EnumDescriptor {
	return file_tarea_proto_enumTypes[1].Descriptor()
}

func (SituacionResp_Estado) Type() protoreflect.EnumType {
	return &file_tarea_proto_enumTypes[1]
}

func (x SituacionResp_Estado) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SituacionResp_Estado.Descriptor instead.
func (SituacionResp_Estado) EnumDescriptor() ([]byte, []int) {
	return file_tarea_proto_rawDescGZIP(), []int{3, 0}
}

type AyudaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peticion int32 `protobuf:"varint,1,opt,name=peticion,proto3" json:"peticion,omitempty"`
}

func (x *AyudaReq) Reset() {
	*x = AyudaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarea_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AyudaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AyudaReq) ProtoMessage() {}

func (x *AyudaReq) ProtoReflect() protoreflect.Message {
	mi := &file_tarea_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AyudaReq.ProtoReflect.Descriptor instead.
func (*AyudaReq) Descriptor() ([]byte, []int) {
	return file_tarea_proto_rawDescGZIP(), []int{0}
}

func (x *AyudaReq) GetPeticion() int32 {
	if x != nil {
		return x.Peticion
	}
	return 0
}

type AyudaResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Disponibilidad AyudaResp_Disponibilidad `protobuf:"varint,1,opt,name=disponibilidad,proto3,enum=grpc.AyudaResp_Disponibilidad" json:"disponibilidad,omitempty"`
}

func (x *AyudaResp) Reset() {
	*x = AyudaResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarea_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AyudaResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AyudaResp) ProtoMessage() {}

func (x *AyudaResp) ProtoReflect() protoreflect.Message {
	mi := &file_tarea_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AyudaResp.ProtoReflect.Descriptor instead.
func (*AyudaResp) Descriptor() ([]byte, []int) {
	return file_tarea_proto_rawDescGZIP(), []int{1}
}

func (x *AyudaResp) GetDisponibilidad() AyudaResp_Disponibilidad {
	if x != nil {
		return x.Disponibilidad
	}
	return AyudaResp_NO
}

type SituacionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peticion int32 `protobuf:"varint,1,opt,name=peticion,proto3" json:"peticion,omitempty"`
}

func (x *SituacionReq) Reset() {
	*x = SituacionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarea_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SituacionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SituacionReq) ProtoMessage() {}

func (x *SituacionReq) ProtoReflect() protoreflect.Message {
	mi := &file_tarea_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SituacionReq.ProtoReflect.Descriptor instead.
func (*SituacionReq) Descriptor() ([]byte, []int) {
	return file_tarea_proto_rawDescGZIP(), []int{2}
}

func (x *SituacionReq) GetPeticion() int32 {
	if x != nil {
		return x.Peticion
	}
	return 0
}

type SituacionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Estado SituacionResp_Estado `protobuf:"varint,1,opt,name=estado,proto3,enum=grpc.SituacionResp_Estado" json:"estado,omitempty"`
}

func (x *SituacionResp) Reset() {
	*x = SituacionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarea_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SituacionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SituacionResp) ProtoMessage() {}

func (x *SituacionResp) ProtoReflect() protoreflect.Message {
	mi := &file_tarea_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SituacionResp.ProtoReflect.Descriptor instead.
func (*SituacionResp) Descriptor() ([]byte, []int) {
	return file_tarea_proto_rawDescGZIP(), []int{3}
}

func (x *SituacionResp) GetEstado() SituacionResp_Estado {
	if x != nil {
		return x.Estado
	}
	return SituacionResp_PELIGRO
}

type FinalizacionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peticion int32 `protobuf:"varint,1,opt,name=peticion,proto3" json:"peticion,omitempty"`
}

func (x *FinalizacionReq) Reset() {
	*x = FinalizacionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarea_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizacionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizacionReq) ProtoMessage() {}

func (x *FinalizacionReq) ProtoReflect() protoreflect.Message {
	mi := &file_tarea_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizacionReq.ProtoReflect.Descriptor instead.
func (*FinalizacionReq) Descriptor() ([]byte, []int) {
	return file_tarea_proto_rawDescGZIP(), []int{4}
}

func (x *FinalizacionReq) GetPeticion() int32 {
	if x != nil {
		return x.Peticion
	}
	return 0
}

type FinalizacionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Respuesta int32 `protobuf:"varint,1,opt,name=respuesta,proto3" json:"respuesta,omitempty"`
}

func (x *FinalizacionResp) Reset() {
	*x = FinalizacionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarea_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizacionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizacionResp) ProtoMessage() {}

func (x *FinalizacionResp) ProtoReflect() protoreflect.Message {
	mi := &file_tarea_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizacionResp.ProtoReflect.Descriptor instead.
func (*FinalizacionResp) Descriptor() ([]byte, []int) {
	return file_tarea_proto_rawDescGZIP(), []int{5}
}

func (x *FinalizacionResp) GetRespuesta() int32 {
	if x != nil {
		return x.Respuesta
	}
	return 0
}

var File_tarea_proto protoreflect.FileDescriptor

var file_tarea_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67,
	0x72, 0x70, 0x63, 0x22, 0x26, 0x0a, 0x08, 0x41, 0x79, 0x75, 0x64, 0x61, 0x52, 0x65, 0x71, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x65, 0x74, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x65, 0x74, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x22, 0x75, 0x0a, 0x09, 0x41,
	0x79, 0x75, 0x64, 0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x70,
	0x6f, 0x6e, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x64, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x79, 0x75, 0x64, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x6f, 0x6e, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x64, 0x61, 0x64,
	0x52, 0x0e, 0x64, 0x69, 0x73, 0x70, 0x6f, 0x6e, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x64, 0x61, 0x64,
	0x22, 0x20, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x70, 0x6f, 0x6e, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x64,
	0x61, 0x64, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x4f, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x53, 0x49,
	0x10, 0x01, 0x22, 0x2a, 0x0a, 0x0c, 0x53, 0x69, 0x74, 0x75, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x65, 0x74, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x65, 0x74, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x22, 0x68,
	0x0a, 0x0d, 0x53, 0x69, 0x74, 0x75, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x32, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x69, 0x74, 0x75, 0x61, 0x63, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x2e, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x52, 0x06, 0x65, 0x73, 0x74,
	0x61, 0x64, 0x6f, 0x22, 0x23, 0x0a, 0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x45, 0x4c, 0x49, 0x47, 0x52, 0x4f, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45,
	0x53, 0x55, 0x45, 0x4c, 0x54, 0x4f, 0x10, 0x01, 0x22, 0x2d, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x61,
	0x6c, 0x69, 0x7a, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x65, 0x74, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x65, 0x74, 0x69, 0x63, 0x69, 0x6f, 0x6e, 0x22, 0x30, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x72, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x32, 0x3a, 0x0a, 0x0e, 0x43, 0x65, 0x6e,
	0x74, 0x72, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x41,
	0x79, 0x75, 0x64, 0x61, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x79, 0x75, 0x64,
	0x61, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x79, 0x75, 0x64,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x32, 0x8b, 0x01, 0x0a, 0x12, 0x4c, 0x61, 0x62, 0x6f, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x69, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0e,
	0x50, 0x65, 0x64, 0x69, 0x72, 0x53, 0x69, 0x74, 0x75, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x69, 0x74, 0x75, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x69, 0x74, 0x75, 0x61, 0x63,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3a, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x72, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6e, 0x61,
	0x6c, 0x69, 0x7a, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6e, 0x69, 0x62, 0x61, 0x6c, 0x6f, 0x78, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x69, 0x64, 0x6f, 0x73, 0x70, 0x72, 0x6f, 0x79, 0x65, 0x63, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tarea_proto_rawDescOnce sync.Once
	file_tarea_proto_rawDescData = file_tarea_proto_rawDesc
)

func file_tarea_proto_rawDescGZIP() []byte {
	file_tarea_proto_rawDescOnce.Do(func() {
		file_tarea_proto_rawDescData = protoimpl.X.CompressGZIP(file_tarea_proto_rawDescData)
	})
	return file_tarea_proto_rawDescData
}

var file_tarea_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_tarea_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tarea_proto_goTypes = []interface{}{
	(AyudaResp_Disponibilidad)(0), // 0: grpc.AyudaResp.Disponibilidad
	(SituacionResp_Estado)(0),     // 1: grpc.SituacionResp.Estado
	(*AyudaReq)(nil),              // 2: grpc.AyudaReq
	(*AyudaResp)(nil),             // 3: grpc.AyudaResp
	(*SituacionReq)(nil),          // 4: grpc.SituacionReq
	(*SituacionResp)(nil),         // 5: grpc.SituacionResp
	(*FinalizacionReq)(nil),       // 6: grpc.FinalizacionReq
	(*FinalizacionResp)(nil),      // 7: grpc.FinalizacionResp
}
var file_tarea_proto_depIdxs = []int32{
	0, // 0: grpc.AyudaResp.disponibilidad:type_name -> grpc.AyudaResp.Disponibilidad
	1, // 1: grpc.SituacionResp.estado:type_name -> grpc.SituacionResp.Estado
	2, // 2: grpc.CentralService.Ayuda:input_type -> grpc.AyudaReq
	4, // 3: grpc.LaboratorioService.PedirSituacion:input_type -> grpc.SituacionReq
	6, // 4: grpc.LaboratorioService.Finalizar:input_type -> grpc.FinalizacionReq
	3, // 5: grpc.CentralService.Ayuda:output_type -> grpc.AyudaResp
	5, // 6: grpc.LaboratorioService.PedirSituacion:output_type -> grpc.SituacionResp
	7, // 7: grpc.LaboratorioService.Finalizar:output_type -> grpc.FinalizacionResp
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tarea_proto_init() }
func file_tarea_proto_init() {
	if File_tarea_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tarea_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AyudaReq); i {
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
		file_tarea_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AyudaResp); i {
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
		file_tarea_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SituacionReq); i {
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
		file_tarea_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SituacionResp); i {
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
		file_tarea_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalizacionReq); i {
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
		file_tarea_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalizacionResp); i {
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
			RawDescriptor: file_tarea_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_tarea_proto_goTypes,
		DependencyIndexes: file_tarea_proto_depIdxs,
		EnumInfos:         file_tarea_proto_enumTypes,
		MessageInfos:      file_tarea_proto_msgTypes,
	}.Build()
	File_tarea_proto = out.File
	file_tarea_proto_rawDesc = nil
	file_tarea_proto_goTypes = nil
	file_tarea_proto_depIdxs = nil
}
