// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model.proto

package pb

import (
	fmt "fmt"
	proto1 "github.com/bifrostcloud/protoc-gen-httpclient/proto"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Package struct {
	PackageName          string   `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	PackagePath          string   `protobuf:"bytes,2,opt,name=package_path,json=packagePath,proto3" json:"package_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Package) Reset()         { *m = Package{} }
func (m *Package) String() string { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()    {}
func (*Package) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}
func (m *Package) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Package.Unmarshal(m, b)
}
func (m *Package) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Package.Marshal(b, m, deterministic)
}
func (m *Package) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Package.Merge(m, src)
}
func (m *Package) XXX_Size() int {
	return xxx_messageInfo_Package.Size(m)
}
func (m *Package) XXX_DiscardUnknown() {
	xxx_messageInfo_Package.DiscardUnknown(m)
}

var xxx_messageInfo_Package proto.InternalMessageInfo

func (m *Package) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *Package) GetPackagePath() string {
	if m != nil {
		return m.PackagePath
	}
	return ""
}

type Input struct {
	FieldName            string   `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	FieldValue           string   `protobuf:"bytes,2,opt,name=field_value,json=fieldValue,proto3" json:"field_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{1}
}
func (m *Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Input.Unmarshal(m, b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Input.Marshal(b, m, deterministic)
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return xxx_messageInfo_Input.Size(m)
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

func (m *Input) GetFieldValue() string {
	if m != nil {
		return m.FieldValue
	}
	return ""
}

type Base struct {
	Package              string    `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
	Imports              []Package `protobuf:"bytes,2,rep,name=imports,proto3" json:"imports"`
	Services             []Service `protobuf:"bytes,3,rep,name=services,proto3" json:"services"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Base) Reset()         { *m = Base{} }
func (m *Base) String() string { return proto.CompactTextString(m) }
func (*Base) ProtoMessage()    {}
func (*Base) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{2}
}
func (m *Base) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Base.Unmarshal(m, b)
}
func (m *Base) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Base.Marshal(b, m, deterministic)
}
func (m *Base) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Base.Merge(m, src)
}
func (m *Base) XXX_Size() int {
	return xxx_messageInfo_Base.Size(m)
}
func (m *Base) XXX_DiscardUnknown() {
	xxx_messageInfo_Base.DiscardUnknown(m)
}

var xxx_messageInfo_Base proto.InternalMessageInfo

func (m *Base) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

func (m *Base) GetImports() []Package {
	if m != nil {
		return m.Imports
	}
	return nil
}

func (m *Base) GetServices() []Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type Service struct {
	UpperCamelCaseServiceName string   `protobuf:"bytes,1,opt,name=upper_camel_case_service_name,json=upperCamelCaseServiceName,proto3" json:"upper_camel_case_service_name,omitempty"`
	LowerCamelCaseServiceName string   `protobuf:"bytes,2,opt,name=lower_camel_case_service_name,json=lowerCamelCaseServiceName,proto3" json:"lower_camel_case_service_name,omitempty"`
	Auth                      string   `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
	Methods                   []Method `protobuf:"bytes,4,rep,name=methods,proto3" json:"methods"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{3}
}
func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetUpperCamelCaseServiceName() string {
	if m != nil {
		return m.UpperCamelCaseServiceName
	}
	return ""
}

func (m *Service) GetLowerCamelCaseServiceName() string {
	if m != nil {
		return m.LowerCamelCaseServiceName
	}
	return ""
}

func (m *Service) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

func (m *Service) GetMethods() []Method {
	if m != nil {
		return m.Methods
	}
	return nil
}

type Method struct {
	UpperCamelCaseServiceName string                `protobuf:"bytes,1,opt,name=upper_camel_case_service_name,json=upperCamelCaseServiceName,proto3" json:"upper_camel_case_service_name,omitempty"`
	UpperCamelCaseMethodName  string                `protobuf:"bytes,2,opt,name=upper_camel_case_method_name,json=upperCamelCaseMethodName,proto3" json:"upper_camel_case_method_name,omitempty"`
	InputType                 string                `protobuf:"bytes,3,opt,name=input_type,json=inputType,proto3" json:"input_type,omitempty"`
	InputFields               Fields                `protobuf:"bytes,4,opt,name=input_fields,json=inputFields,proto3" json:"input_fields"`
	OutputType                string                `protobuf:"bytes,5,opt,name=output_type,json=outputType,proto3" json:"output_type,omitempty"`
	Auth                      string                `protobuf:"bytes,6,opt,name=auth,proto3" json:"auth,omitempty"`
	RequestOptions            proto1.RequestOptions `protobuf:"bytes,7,opt,name=request_options,json=requestOptions,proto3" json:"request_options"`
	XXX_NoUnkeyedLiteral      struct{}              `json:"-"`
	XXX_unrecognized          []byte                `json:"-"`
	XXX_sizecache             int32                 `json:"-"`
}

func (m *Method) Reset()         { *m = Method{} }
func (m *Method) String() string { return proto.CompactTextString(m) }
func (*Method) ProtoMessage()    {}
func (*Method) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{4}
}
func (m *Method) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Method.Unmarshal(m, b)
}
func (m *Method) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Method.Marshal(b, m, deterministic)
}
func (m *Method) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Method.Merge(m, src)
}
func (m *Method) XXX_Size() int {
	return xxx_messageInfo_Method.Size(m)
}
func (m *Method) XXX_DiscardUnknown() {
	xxx_messageInfo_Method.DiscardUnknown(m)
}

var xxx_messageInfo_Method proto.InternalMessageInfo

func (m *Method) GetUpperCamelCaseServiceName() string {
	if m != nil {
		return m.UpperCamelCaseServiceName
	}
	return ""
}

func (m *Method) GetUpperCamelCaseMethodName() string {
	if m != nil {
		return m.UpperCamelCaseMethodName
	}
	return ""
}

func (m *Method) GetInputType() string {
	if m != nil {
		return m.InputType
	}
	return ""
}

func (m *Method) GetInputFields() Fields {
	if m != nil {
		return m.InputFields
	}
	return Fields{}
}

func (m *Method) GetOutputType() string {
	if m != nil {
		return m.OutputType
	}
	return ""
}

func (m *Method) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

func (m *Method) GetRequestOptions() proto1.RequestOptions {
	if m != nil {
		return m.RequestOptions
	}
	return proto1.RequestOptions{}
}

type Fields struct {
	FieldImport          []FieldImport `protobuf:"bytes,1,rep,name=field_import,json=fieldImport,proto3" json:"field_import"`
	Type                 string        `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	UpperCamelCase       string        `protobuf:"bytes,3,opt,name=upper_camel_case,json=upperCamelCase,proto3" json:"upper_camel_case,omitempty"`
	Base                 []string      `protobuf:"bytes,4,rep,name=base,proto3" json:"base,omitempty"`
	Lowercase            []string      `protobuf:"bytes,5,rep,name=lowercase,proto3" json:"lowercase,omitempty"`
	DotNotation          []string      `protobuf:"bytes,6,rep,name=dot_notation,json=dotNotation,proto3" json:"dot_notation,omitempty"`
	ParamCase            []string      `protobuf:"bytes,7,rep,name=param_case,json=paramCase,proto3" json:"param_case,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Fields) Reset()         { *m = Fields{} }
func (m *Fields) String() string { return proto.CompactTextString(m) }
func (*Fields) ProtoMessage()    {}
func (*Fields) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{5}
}
func (m *Fields) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fields.Unmarshal(m, b)
}
func (m *Fields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fields.Marshal(b, m, deterministic)
}
func (m *Fields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fields.Merge(m, src)
}
func (m *Fields) XXX_Size() int {
	return xxx_messageInfo_Fields.Size(m)
}
func (m *Fields) XXX_DiscardUnknown() {
	xxx_messageInfo_Fields.DiscardUnknown(m)
}

var xxx_messageInfo_Fields proto.InternalMessageInfo

func (m *Fields) GetFieldImport() []FieldImport {
	if m != nil {
		return m.FieldImport
	}
	return nil
}

func (m *Fields) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Fields) GetUpperCamelCase() string {
	if m != nil {
		return m.UpperCamelCase
	}
	return ""
}

func (m *Fields) GetBase() []string {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *Fields) GetLowercase() []string {
	if m != nil {
		return m.Lowercase
	}
	return nil
}

func (m *Fields) GetDotNotation() []string {
	if m != nil {
		return m.DotNotation
	}
	return nil
}

func (m *Fields) GetParamCase() []string {
	if m != nil {
		return m.ParamCase
	}
	return nil
}

type FieldImport struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tag                  string   `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldImport) Reset()         { *m = FieldImport{} }
func (m *FieldImport) String() string { return proto.CompactTextString(m) }
func (*FieldImport) ProtoMessage()    {}
func (*FieldImport) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{6}
}
func (m *FieldImport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldImport.Unmarshal(m, b)
}
func (m *FieldImport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldImport.Marshal(b, m, deterministic)
}
func (m *FieldImport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldImport.Merge(m, src)
}
func (m *FieldImport) XXX_Size() int {
	return xxx_messageInfo_FieldImport.Size(m)
}
func (m *FieldImport) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldImport.DiscardUnknown(m)
}

var xxx_messageInfo_FieldImport proto.InternalMessageInfo

func (m *FieldImport) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FieldImport) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func init() {
	proto.RegisterType((*Package)(nil), "Package")
	proto.RegisterType((*Input)(nil), "Input")
	proto.RegisterType((*Base)(nil), "Base")
	proto.RegisterType((*Service)(nil), "Service")
	proto.RegisterType((*Method)(nil), "Method")
	proto.RegisterType((*Fields)(nil), "Fields")
	proto.RegisterType((*FieldImport)(nil), "FieldImport")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x4f, 0xdb, 0x4e,
	0x10, 0xcd, 0x3f, 0x62, 0x32, 0x8e, 0x00, 0xed, 0xc9, 0xfc, 0x44, 0x80, 0x9f, 0x2f, 0x8d, 0x2a,
	0x11, 0xaa, 0xa2, 0x5e, 0x29, 0x02, 0x89, 0x8a, 0x43, 0x01, 0xa5, 0x55, 0x0f, 0xbd, 0x58, 0x1b,
	0x7b, 0x63, 0x5b, 0xb5, 0xbd, 0x5b, 0x7b, 0x4d, 0xc5, 0x37, 0x44, 0x3d, 0xf1, 0x09, 0xaa, 0x8a,
	0x6f, 0xd0, 0x6f, 0x50, 0xed, 0xcc, 0x06, 0xec, 0xb6, 0xea, 0xa9, 0xb7, 0xd9, 0xf7, 0xde, 0xbe,
	0xd9, 0x37, 0x9e, 0x04, 0xdc, 0x5c, 0x46, 0x22, 0x9b, 0xa9, 0x52, 0x6a, 0xf9, 0xdf, 0xeb, 0x38,
	0xd5, 0x49, 0xbd, 0x98, 0x85, 0x32, 0x3f, 0x5c, 0xa4, 0xcb, 0x52, 0x56, 0x3a, 0xcc, 0x64, 0x1d,
	0x1d, 0x22, 0x1d, 0x1e, 0xc4, 0xa2, 0x38, 0x48, 0xb4, 0x56, 0x61, 0x96, 0x8a, 0x42, 0x13, 0x7a,
	0xa8, 0x6f, 0x95, 0xa8, 0xac, 0xc1, 0x41, 0xc3, 0x20, 0x96, 0xb1, 0x24, 0xc9, 0xa2, 0x5e, 0xe2,
	0x89, 0xf4, 0xa6, 0x22, 0xb9, 0x7f, 0x05, 0xce, 0x35, 0x0f, 0x3f, 0xf1, 0x58, 0xb0, 0xff, 0x61,
	0xac, 0xa8, 0x0c, 0x0a, 0x9e, 0x0b, 0xaf, 0xbb, 0xdf, 0x9d, 0x8e, 0xe6, 0xae, 0xc5, 0x2e, 0x79,
	0xde, 0x92, 0x28, 0xae, 0x13, 0xaf, 0xd7, 0x92, 0x5c, 0x73, 0x9d, 0xf8, 0x6f, 0x60, 0xed, 0xa2,
	0x50, 0xb5, 0x66, 0x13, 0x80, 0x65, 0x2a, 0xb2, 0xa8, 0x69, 0x36, 0x42, 0x04, 0xad, 0xf6, 0xc0,
	0x25, 0xfa, 0x86, 0x67, 0xb5, 0xb0, 0x4e, 0x74, 0xe3, 0x83, 0x41, 0xfc, 0x1b, 0x18, 0x9c, 0xf2,
	0x4a, 0x30, 0x0f, 0x1c, 0xeb, 0x6f, 0x4d, 0x56, 0x47, 0x36, 0x05, 0x27, 0xcd, 0x95, 0x2c, 0x75,
	0xe5, 0xf5, 0xf6, 0xfb, 0x53, 0xf7, 0xe5, 0xfa, 0xcc, 0x66, 0x39, 0x1d, 0xdc, 0x7d, 0xdb, 0xeb,
	0xcc, 0x57, 0x34, 0x7b, 0x0e, 0xeb, 0x95, 0x28, 0x6f, 0xd2, 0x50, 0x54, 0x5e, 0xdf, 0x4a, 0xdf,
	0x11, 0x60, 0xa5, 0x8f, 0xbc, 0xff, 0xb5, 0x0b, 0x8e, 0xe5, 0xd8, 0x09, 0x4c, 0x6a, 0xa5, 0x44,
	0x19, 0x84, 0x3c, 0x17, 0x59, 0x10, 0xf2, 0x4a, 0x04, 0x56, 0xd8, 0x8c, 0xb5, 0x8d, 0xa2, 0x33,
	0xa3, 0x39, 0xe3, 0x95, 0xb0, 0xb7, 0x31, 0xe6, 0x09, 0x4c, 0x32, 0xf9, 0xe5, 0x2f, 0x0e, 0x14,
	0x7c, 0x1b, 0x45, 0x7f, 0x74, 0x60, 0x30, 0xe0, 0xb5, 0x4e, 0xbc, 0x3e, 0x0a, 0xb1, 0x66, 0xcf,
	0xc0, 0xc9, 0x85, 0x4e, 0x64, 0x54, 0x79, 0x03, 0x8c, 0xe3, 0xcc, 0xde, 0xe2, 0x79, 0x15, 0xdc,
	0xb2, 0xfe, 0x7d, 0x0f, 0x86, 0xc4, 0xfc, 0x83, 0x2c, 0xc7, 0xb0, 0xf3, 0x9b, 0x03, 0x35, 0x6a,
	0x46, 0xf1, 0xda, 0x06, 0xd4, 0x1d, 0xef, 0x4f, 0x00, 0x52, 0xb3, 0x1a, 0x81, 0xd9, 0x57, 0x9b,
	0x67, 0x84, 0xc8, 0xfb, 0x5b, 0x25, 0xd8, 0x0b, 0x18, 0x13, 0x8d, 0x4b, 0x60, 0x92, 0x75, 0x31,
	0xd9, 0x39, 0x1e, 0x6d, 0x32, 0x17, 0x25, 0x04, 0x99, 0x1d, 0x92, 0xb5, 0x7e, 0x74, 0x5c, 0xa3,
	0x1d, 0x22, 0x08, 0x2d, 0x57, 0xb3, 0x1b, 0x36, 0x66, 0x77, 0x0c, 0x9b, 0xa5, 0xf8, 0x5c, 0x8b,
	0x4a, 0x07, 0x52, 0xe9, 0x54, 0x16, 0x95, 0xe7, 0x60, 0xa7, 0xcd, 0xd9, 0x9c, 0xf0, 0x2b, 0x82,
	0x6d, 0xc7, 0x8d, 0xb2, 0x85, 0xfa, 0x3f, 0xba, 0x30, 0xb4, 0xfd, 0x5f, 0xc1, 0x98, 0x76, 0x98,
	0xf6, 0xcc, 0xeb, 0xe2, 0xb7, 0x18, 0xd3, 0x8b, 0x2f, 0x10, 0x5b, 0x3d, 0x7b, 0xf9, 0x04, 0x99,
	0x57, 0xe1, 0x7b, 0x69, 0x5e, 0x58, 0xb3, 0x29, 0x6c, 0xfd, 0x3a, 0x5b, 0x3b, 0xa1, 0x8d, 0xf6,
	0x3c, 0xcd, 0xed, 0x85, 0x61, 0xcd, 0x87, 0x1f, 0xcd, 0xb1, 0x66, 0x3b, 0x30, 0xc2, 0x05, 0xc2,
	0x6b, 0x6b, 0x48, 0x3c, 0x01, 0xe6, 0x57, 0x1b, 0x49, 0x1d, 0x14, 0x52, 0x73, 0x13, 0xc1, 0x1b,
	0xa2, 0xc0, 0x8d, 0xa4, 0xbe, 0xb4, 0x90, 0xf9, 0x34, 0x8a, 0x97, 0x3c, 0xa7, 0xc6, 0x0e, 0x39,
	0x20, 0x62, 0x7a, 0xfa, 0x47, 0xe0, 0x9e, 0xb7, 0x03, 0x34, 0x36, 0x06, 0x6b, 0xb6, 0x05, 0x7d,
	0xcd, 0x63, 0x9b, 0xc9, 0x94, 0xa7, 0x5b, 0x77, 0x0f, 0xbb, 0x9d, 0xfb, 0x87, 0xdd, 0xce, 0xf7,
	0x87, 0xdd, 0xce, 0xc7, 0x9e, 0x5a, 0x2c, 0x86, 0xf8, 0x9f, 0x73, 0xf4, 0x33, 0x00, 0x00, 0xff,
	0xff, 0x01, 0xfb, 0x63, 0x63, 0xf2, 0x04, 0x00, 0x00,
}
