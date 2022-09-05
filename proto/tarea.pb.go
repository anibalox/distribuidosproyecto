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

type SituacionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NroEscuadra string `protobuf:"bytes,1,opt,name=nro_escuadra,json=nroEscuadra,proto3" json:"nro_escuadra,omitempty"`
}

func (x *SituacionReq) Reset() {
	*x = SituacionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarea_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SituacionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SituacionReq) ProtoMessage() {}

func (x *SituacionReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SituacionReq.ProtoReflect.Descriptor instead.
func (*SituacionReq) Descriptor() ([]byte, []int) {
	return file_tarea_proto_rawDescGZIP(), []int{0}
}

func (x *SituacionReq) GetNroEscuadra() string {
	if x != nil {
		return x.NroEscuadra
	}
	return ""
}

type Termino struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Termino string `protobuf:"bytes,1,opt,name=Termino,proto3" json:"Termino,omitempty"`
}

func (x *Termino) Reset() {
	*x = Termino{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarea_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Termino) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Termino) ProtoMessage() {}

func (x *Termino) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Termino.ProtoReflect.Descriptor instead.
func (*Termino) Descriptor() ([]byte, []int) {
	return file_tarea_proto_rawDescGZIP(), []int{1}
}

func (x *Termino) GetTermino() string {
	if x != nil {
		return x.Termino
	}
	return ""
}

type SituacionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resuelta    string `protobuf:"bytes,1,opt,name=resuelta,proto3" json:"resuelta,omitempty"`
	NroEscuadra string `protobuf:"bytes,2,opt,name=nro_escuadra,json=nroEscuadra,proto3" json:"nro_escuadra,omitempty"`
	NroLab      string `protobuf:"bytes,3,opt,name=nro_lab,json=nroLab,proto3" json:"nro_lab,omitempty"`
}

func (x *SituacionResp) Reset() {
	*x = SituacionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tarea_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SituacionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SituacionResp) ProtoMessage() {}

func (x *SituacionResp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SituacionResp.ProtoReflect.Descriptor instead.
func (*SituacionResp) Descriptor() ([]byte, []int) {
	return file_tarea_proto_rawDescGZIP(), []int{2}
}

func (x *SituacionResp) GetResuelta() string {
	if x != nil {
		return x.Resuelta
	}
	return ""
}

func (x *SituacionResp) GetNroEscuadra() string {
	if x != nil {
		return x.NroEscuadra
	}
	return ""
}

func (x *SituacionResp) GetNroLab() string {
	if x != nil {
		return x.NroLab
	}
	return ""
}

var File_tarea_proto protoreflect.FileDescriptor

var file_tarea_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67,
	0x72, 0x70, 0x63, 0x22, 0x31, 0x0a, 0x0c, 0x53, 0x69, 0x74, 0x75, 0x61, 0x63, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x72, 0x6f, 0x5f, 0x65, 0x73, 0x63, 0x75, 0x61,
	0x64, 0x72, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x72, 0x6f, 0x45, 0x73,
	0x63, 0x75, 0x61, 0x64, 0x72, 0x61, 0x22, 0x23, 0x0a, 0x07, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x6f, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x6f, 0x22, 0x67, 0x0a, 0x0d, 0x53,
	0x69, 0x74, 0x75, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x75, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x75, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x72, 0x6f, 0x5f,
	0x65, 0x73, 0x63, 0x75, 0x61, 0x64, 0x72, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x72, 0x6f, 0x45, 0x73, 0x63, 0x75, 0x61, 0x64, 0x72, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x6e,
	0x72, 0x6f, 0x5f, 0x6c, 0x61, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x72,
	0x6f, 0x4c, 0x61, 0x62, 0x32, 0x7c, 0x0a, 0x0e, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x11, 0x41, 0x62, 0x72, 0x69, 0x72, 0x43,
	0x6f, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x69, 0x74, 0x75, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x69, 0x74, 0x75, 0x61, 0x63, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x28, 0x01, 0x30, 0x01, 0x12, 0x28, 0x0a, 0x08, 0x54, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x72, 0x12, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x6f, 0x1a, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x6f, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x6e, 0x69, 0x62, 0x61, 0x6c, 0x6f, 0x78, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x69, 0x64, 0x6f, 0x73, 0x70, 0x72, 0x6f, 0x79, 0x65, 0x63, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_tarea_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tarea_proto_goTypes = []interface{}{
	(*SituacionReq)(nil),  // 0: grpc.SituacionReq
	(*Termino)(nil),       // 1: grpc.Termino
	(*SituacionResp)(nil), // 2: grpc.SituacionResp
}
var file_tarea_proto_depIdxs = []int32{
	2, // 0: grpc.CentralService.AbrirComunicacion:input_type -> grpc.SituacionResp
	1, // 1: grpc.CentralService.Terminar:input_type -> grpc.Termino
	0, // 2: grpc.CentralService.AbrirComunicacion:output_type -> grpc.SituacionReq
	1, // 3: grpc.CentralService.Terminar:output_type -> grpc.Termino
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tarea_proto_init() }
func file_tarea_proto_init() {
	if File_tarea_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tarea_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_tarea_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Termino); i {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tarea_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tarea_proto_goTypes,
		DependencyIndexes: file_tarea_proto_depIdxs,
		MessageInfos:      file_tarea_proto_msgTypes,
	}.Build()
	File_tarea_proto = out.File
	file_tarea_proto_rawDesc = nil
	file_tarea_proto_goTypes = nil
	file_tarea_proto_depIdxs = nil
}
