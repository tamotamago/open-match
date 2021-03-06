// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/backend.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Configuration for the Match Function to be triggered by Open Match to
// generate proposals.
type MatchFunctionConfig struct {
	// A developer-chosen human-readable name for this Match Function.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Properties for the type of this function.
	//
	// Types that are valid to be assigned to Function:
	//	*MatchFunctionConfig_GrpcFunction
	//	*MatchFunctionConfig_RestFunction
	Function             isMatchFunctionConfig_Function `protobuf_oneof:"function"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *MatchFunctionConfig) Reset()         { *m = MatchFunctionConfig{} }
func (m *MatchFunctionConfig) String() string { return proto.CompactTextString(m) }
func (*MatchFunctionConfig) ProtoMessage()    {}
func (*MatchFunctionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{0}
}

func (m *MatchFunctionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchFunctionConfig.Unmarshal(m, b)
}
func (m *MatchFunctionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchFunctionConfig.Marshal(b, m, deterministic)
}
func (m *MatchFunctionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchFunctionConfig.Merge(m, src)
}
func (m *MatchFunctionConfig) XXX_Size() int {
	return xxx_messageInfo_MatchFunctionConfig.Size(m)
}
func (m *MatchFunctionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchFunctionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MatchFunctionConfig proto.InternalMessageInfo

func (m *MatchFunctionConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isMatchFunctionConfig_Function interface {
	isMatchFunctionConfig_Function()
}

type MatchFunctionConfig_GrpcFunction struct {
	GrpcFunction *MatchFunctionConfig_GrpcFunctionProperties `protobuf:"bytes,10001,opt,name=grpc_function,json=grpcFunction,proto3,oneof"`
}

type MatchFunctionConfig_RestFunction struct {
	RestFunction *MatchFunctionConfig_RestFunctionProperties `protobuf:"bytes,10002,opt,name=rest_function,json=restFunction,proto3,oneof"`
}

func (*MatchFunctionConfig_GrpcFunction) isMatchFunctionConfig_Function() {}

func (*MatchFunctionConfig_RestFunction) isMatchFunctionConfig_Function() {}

func (m *MatchFunctionConfig) GetFunction() isMatchFunctionConfig_Function {
	if m != nil {
		return m.Function
	}
	return nil
}

func (m *MatchFunctionConfig) GetGrpcFunction() *MatchFunctionConfig_GrpcFunctionProperties {
	if x, ok := m.GetFunction().(*MatchFunctionConfig_GrpcFunction); ok {
		return x.GrpcFunction
	}
	return nil
}

func (m *MatchFunctionConfig) GetRestFunction() *MatchFunctionConfig_RestFunctionProperties {
	if x, ok := m.GetFunction().(*MatchFunctionConfig_RestFunction); ok {
		return x.RestFunction
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MatchFunctionConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MatchFunctionConfig_GrpcFunction)(nil),
		(*MatchFunctionConfig_RestFunction)(nil),
	}
}

type MatchFunctionConfig_GrpcFunctionProperties struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MatchFunctionConfig_GrpcFunctionProperties) Reset() {
	*m = MatchFunctionConfig_GrpcFunctionProperties{}
}
func (m *MatchFunctionConfig_GrpcFunctionProperties) String() string {
	return proto.CompactTextString(m)
}
func (*MatchFunctionConfig_GrpcFunctionProperties) ProtoMessage() {}
func (*MatchFunctionConfig_GrpcFunctionProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{0, 0}
}

func (m *MatchFunctionConfig_GrpcFunctionProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchFunctionConfig_GrpcFunctionProperties.Unmarshal(m, b)
}
func (m *MatchFunctionConfig_GrpcFunctionProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchFunctionConfig_GrpcFunctionProperties.Marshal(b, m, deterministic)
}
func (m *MatchFunctionConfig_GrpcFunctionProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchFunctionConfig_GrpcFunctionProperties.Merge(m, src)
}
func (m *MatchFunctionConfig_GrpcFunctionProperties) XXX_Size() int {
	return xxx_messageInfo_MatchFunctionConfig_GrpcFunctionProperties.Size(m)
}
func (m *MatchFunctionConfig_GrpcFunctionProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchFunctionConfig_GrpcFunctionProperties.DiscardUnknown(m)
}

var xxx_messageInfo_MatchFunctionConfig_GrpcFunctionProperties proto.InternalMessageInfo

func (m *MatchFunctionConfig_GrpcFunctionProperties) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *MatchFunctionConfig_GrpcFunctionProperties) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type MatchFunctionConfig_RestFunctionProperties struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MatchFunctionConfig_RestFunctionProperties) Reset() {
	*m = MatchFunctionConfig_RestFunctionProperties{}
}
func (m *MatchFunctionConfig_RestFunctionProperties) String() string {
	return proto.CompactTextString(m)
}
func (*MatchFunctionConfig_RestFunctionProperties) ProtoMessage() {}
func (*MatchFunctionConfig_RestFunctionProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{0, 1}
}

func (m *MatchFunctionConfig_RestFunctionProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchFunctionConfig_RestFunctionProperties.Unmarshal(m, b)
}
func (m *MatchFunctionConfig_RestFunctionProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchFunctionConfig_RestFunctionProperties.Marshal(b, m, deterministic)
}
func (m *MatchFunctionConfig_RestFunctionProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchFunctionConfig_RestFunctionProperties.Merge(m, src)
}
func (m *MatchFunctionConfig_RestFunctionProperties) XXX_Size() int {
	return xxx_messageInfo_MatchFunctionConfig_RestFunctionProperties.Size(m)
}
func (m *MatchFunctionConfig_RestFunctionProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchFunctionConfig_RestFunctionProperties.DiscardUnknown(m)
}

var xxx_messageInfo_MatchFunctionConfig_RestFunctionProperties proto.InternalMessageInfo

func (m *MatchFunctionConfig_RestFunctionProperties) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *MatchFunctionConfig_RestFunctionProperties) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type GetMatchesRequest struct {
	// MatchFunction to be executed for the given list of MatchProfiles
	Function *MatchFunctionConfig `protobuf:"bytes,1,opt,name=function,proto3" json:"function,omitempty"`
	// MatchProfiles for which this MatchFunction should be executed.
	Profile              []*MatchProfile `protobuf:"bytes,2,rep,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetMatchesRequest) Reset()         { *m = GetMatchesRequest{} }
func (m *GetMatchesRequest) String() string { return proto.CompactTextString(m) }
func (*GetMatchesRequest) ProtoMessage()    {}
func (*GetMatchesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{1}
}

func (m *GetMatchesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMatchesRequest.Unmarshal(m, b)
}
func (m *GetMatchesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMatchesRequest.Marshal(b, m, deterministic)
}
func (m *GetMatchesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMatchesRequest.Merge(m, src)
}
func (m *GetMatchesRequest) XXX_Size() int {
	return xxx_messageInfo_GetMatchesRequest.Size(m)
}
func (m *GetMatchesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMatchesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMatchesRequest proto.InternalMessageInfo

func (m *GetMatchesRequest) GetFunction() *MatchFunctionConfig {
	if m != nil {
		return m.Function
	}
	return nil
}

func (m *GetMatchesRequest) GetProfile() []*MatchProfile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type GetMatchesResponse struct {
	// Result Match for the requested MatchProfile.
	Match                *Match   `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMatchesResponse) Reset()         { *m = GetMatchesResponse{} }
func (m *GetMatchesResponse) String() string { return proto.CompactTextString(m) }
func (*GetMatchesResponse) ProtoMessage()    {}
func (*GetMatchesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{2}
}

func (m *GetMatchesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMatchesResponse.Unmarshal(m, b)
}
func (m *GetMatchesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMatchesResponse.Marshal(b, m, deterministic)
}
func (m *GetMatchesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMatchesResponse.Merge(m, src)
}
func (m *GetMatchesResponse) XXX_Size() int {
	return xxx_messageInfo_GetMatchesResponse.Size(m)
}
func (m *GetMatchesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMatchesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMatchesResponse proto.InternalMessageInfo

func (m *GetMatchesResponse) GetMatch() *Match {
	if m != nil {
		return m.Match
	}
	return nil
}

type AssignTicketsRequest struct {
	// List of Ticket IDs for which the Assignment is to be made.
	TicketId []string `protobuf:"bytes,1,rep,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	// Assignment to be associated with the Ticket IDs.
	Assignment           *Assignment `protobuf:"bytes,2,opt,name=assignment,proto3" json:"assignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AssignTicketsRequest) Reset()         { *m = AssignTicketsRequest{} }
func (m *AssignTicketsRequest) String() string { return proto.CompactTextString(m) }
func (*AssignTicketsRequest) ProtoMessage()    {}
func (*AssignTicketsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{3}
}

func (m *AssignTicketsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignTicketsRequest.Unmarshal(m, b)
}
func (m *AssignTicketsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignTicketsRequest.Marshal(b, m, deterministic)
}
func (m *AssignTicketsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignTicketsRequest.Merge(m, src)
}
func (m *AssignTicketsRequest) XXX_Size() int {
	return xxx_messageInfo_AssignTicketsRequest.Size(m)
}
func (m *AssignTicketsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignTicketsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AssignTicketsRequest proto.InternalMessageInfo

func (m *AssignTicketsRequest) GetTicketId() []string {
	if m != nil {
		return m.TicketId
	}
	return nil
}

func (m *AssignTicketsRequest) GetAssignment() *Assignment {
	if m != nil {
		return m.Assignment
	}
	return nil
}

type AssignTicketsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssignTicketsResponse) Reset()         { *m = AssignTicketsResponse{} }
func (m *AssignTicketsResponse) String() string { return proto.CompactTextString(m) }
func (*AssignTicketsResponse) ProtoMessage()    {}
func (*AssignTicketsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dab762378f455cd, []int{4}
}

func (m *AssignTicketsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignTicketsResponse.Unmarshal(m, b)
}
func (m *AssignTicketsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignTicketsResponse.Marshal(b, m, deterministic)
}
func (m *AssignTicketsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignTicketsResponse.Merge(m, src)
}
func (m *AssignTicketsResponse) XXX_Size() int {
	return xxx_messageInfo_AssignTicketsResponse.Size(m)
}
func (m *AssignTicketsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignTicketsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AssignTicketsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MatchFunctionConfig)(nil), "api.MatchFunctionConfig")
	proto.RegisterType((*MatchFunctionConfig_GrpcFunctionProperties)(nil), "api.MatchFunctionConfig.GrpcFunctionProperties")
	proto.RegisterType((*MatchFunctionConfig_RestFunctionProperties)(nil), "api.MatchFunctionConfig.RestFunctionProperties")
	proto.RegisterType((*GetMatchesRequest)(nil), "api.GetMatchesRequest")
	proto.RegisterType((*GetMatchesResponse)(nil), "api.GetMatchesResponse")
	proto.RegisterType((*AssignTicketsRequest)(nil), "api.AssignTicketsRequest")
	proto.RegisterType((*AssignTicketsResponse)(nil), "api.AssignTicketsResponse")
}

func init() { proto.RegisterFile("api/backend.proto", fileDescriptor_8dab762378f455cd) }

var fileDescriptor_8dab762378f455cd = []byte{
	// 728 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x41, 0x8f, 0xe3, 0x34,
	0x18, 0x25, 0xe9, 0xee, 0xce, 0x8e, 0xcb, 0x0a, 0xc6, 0xc0, 0x6c, 0x09, 0x48, 0x98, 0x20, 0x44,
	0x55, 0x68, 0xdc, 0x0d, 0x15, 0x87, 0x22, 0xa4, 0xed, 0x0c, 0xc3, 0x30, 0xd2, 0xb2, 0x8c, 0x02,
	0x12, 0x12, 0x1c, 0x56, 0xae, 0xf3, 0x35, 0x35, 0x9b, 0xd8, 0xc6, 0x76, 0x66, 0x39, 0xf3, 0x0f,
	0x18, 0x6e, 0xfc, 0x10, 0xfe, 0x08, 0x67, 0x6e, 0x1c, 0x90, 0x38, 0x73, 0x47, 0x71, 0xa6, 0x93,
	0x0e, 0x2d, 0x82, 0x53, 0xdd, 0xf7, 0x9e, 0xdf, 0xf7, 0xbd, 0xcf, 0xb1, 0xd1, 0x01, 0xd3, 0x82,
	0x2e, 0x18, 0x7f, 0x0a, 0x32, 0x4f, 0xb4, 0x51, 0x4e, 0xe1, 0x1e, 0xd3, 0x22, 0xc2, 0x0d, 0x5e,
	0x81, 0xb5, 0xac, 0x00, 0xdb, 0x12, 0xd1, 0xeb, 0x85, 0x52, 0x45, 0x09, 0xb4, 0xa1, 0x98, 0x94,
	0xca, 0x31, 0x27, 0x94, 0x5c, 0xb3, 0xef, 0xf9, 0x1f, 0x3e, 0x2e, 0x40, 0x8e, 0xed, 0x33, 0x56,
	0x14, 0x60, 0xa8, 0xd2, 0x5e, 0xb1, 0xad, 0x8e, 0xff, 0x0c, 0xd1, 0x4b, 0x9f, 0x31, 0xc7, 0x57,
	0x9f, 0xd4, 0x92, 0x37, 0xc4, 0xb1, 0x92, 0x4b, 0x51, 0x60, 0x8c, 0x6e, 0x49, 0x56, 0xc1, 0x20,
	0x20, 0xc1, 0x70, 0x3f, 0xf3, 0x6b, 0xfc, 0x15, 0xba, 0x57, 0x18, 0xcd, 0x9f, 0x2c, 0xaf, 0xa4,
	0x83, 0x1f, 0x1f, 0x93, 0x60, 0xd8, 0x4f, 0x69, 0xc2, 0xb4, 0x48, 0x76, 0xb8, 0x24, 0xa7, 0x46,
	0xf3, 0x35, 0x74, 0x6e, 0x94, 0x06, 0xe3, 0x04, 0xd8, 0x4f, 0x9f, 0xcb, 0x9e, 0x2f, 0x36, 0x98,
	0xc6, 0xd8, 0x80, 0x75, 0x9d, 0xf1, 0xe5, 0x7f, 0x19, 0x67, 0x60, 0xdd, 0x6e, 0x63, 0xb3, 0xc1,
	0x44, 0x0f, 0xd1, 0xe1, 0xee, 0x16, 0x9a, 0x7c, 0x2b, 0x65, 0xdd, 0x3a, 0x5f, 0xb3, 0x6e, 0x30,
	0xad, 0x8c, 0x1b, 0x84, 0x24, 0x18, 0xde, 0xce, 0xfc, 0xba, 0x71, 0xd8, 0x5d, 0xeb, 0xff, 0x3a,
	0x1c, 0x21, 0x74, 0x77, 0x9d, 0x2b, 0xbe, 0x40, 0x07, 0xa7, 0xe0, 0x7c, 0x20, 0xb0, 0x19, 0x7c,
	0x57, 0x83, 0x75, 0x78, 0xda, 0x09, 0xbc, 0x59, 0x3f, 0x1d, 0xfc, 0x5b, 0xee, 0xec, 0x5a, 0x89,
	0xdf, 0x45, 0x7b, 0xda, 0xa8, 0xa5, 0x28, 0x61, 0x10, 0x92, 0xde, 0xb0, 0x9f, 0x1e, 0x74, 0x9b,
	0xce, 0x5b, 0x22, 0x5b, 0x2b, 0xe2, 0x0f, 0x10, 0xde, 0xac, 0x6b, 0xb5, 0x92, 0x16, 0x30, 0x41,
	0xb7, 0xab, 0x06, 0xba, 0xaa, 0x8a, 0x3a, 0x83, 0xac, 0x25, 0xe2, 0x1c, 0xbd, 0x3c, 0xb7, 0x56,
	0x14, 0xf2, 0x4b, 0xc1, 0x9f, 0x82, 0xbb, 0x6e, 0xf9, 0x35, 0xb4, 0xef, 0x3c, 0xf2, 0x44, 0xe4,
	0x83, 0x80, 0xf4, 0x86, 0xfb, 0xd9, 0xdd, 0x16, 0x38, 0xcb, 0x31, 0x45, 0x88, 0xf9, 0x4d, 0x15,
	0xc8, 0x76, 0x14, 0xfd, 0xf4, 0x05, 0xef, 0x3d, 0xbf, 0x86, 0xb3, 0x0d, 0x49, 0x7c, 0x1f, 0xbd,
	0xf2, 0x8f, 0x2a, 0x6d, 0x83, 0xe9, 0x1f, 0x01, 0xda, 0x3b, 0x6a, 0xef, 0x04, 0x96, 0x08, 0x75,
	0x11, 0xf0, 0xa1, 0xf7, 0xdb, 0x9a, 0x65, 0x74, 0x7f, 0x0b, 0x6f, 0xad, 0xe2, 0xf1, 0x0f, 0xbf,
	0xfe, 0xfe, 0x53, 0xf8, 0x4e, 0x1c, 0xd3, 0x8b, 0x07, 0xeb, 0x7b, 0x46, 0xab, 0x56, 0x44, 0xb9,
	0x01, 0xe6, 0xa0, 0x39, 0x5b, 0x65, 0x59, 0x39, 0x0b, 0x46, 0x93, 0x00, 0x57, 0xe8, 0xde, 0x8d,
	0xa6, 0xf0, 0xab, 0x1b, 0x11, 0x6e, 0x8e, 0x23, 0x8a, 0x76, 0x51, 0x57, 0x85, 0xdf, 0xf6, 0x85,
	0xdf, 0x88, 0xa3, 0xcd, 0xc2, 0xed, 0xac, 0x2c, 0x6d, 0x87, 0x30, 0x0b, 0x46, 0x47, 0x7f, 0x85,
	0x97, 0xf3, 0xdf, 0x42, 0xfc, 0x4b, 0x97, 0x38, 0x3e, 0x43, 0xe8, 0x73, 0x0d, 0x92, 0xf8, 0x28,
	0xf8, 0x70, 0xe5, 0x9c, 0xb6, 0x33, 0x4a, 0x95, 0x06, 0x39, 0xf6, 0x9d, 0x27, 0x39, 0x5c, 0x44,
	0x6f, 0x75, 0xff, 0xc7, 0xb9, 0xb0, 0xbc, 0xb6, 0xf6, 0x61, 0xfb, 0x3e, 0x14, 0x46, 0xd5, 0xda,
	0x26, 0x5c, 0x55, 0xa3, 0x6f, 0x10, 0x9e, 0x6b, 0xc6, 0x57, 0x40, 0xd2, 0x64, 0x42, 0x1e, 0x09,
	0x0e, 0xcd, 0xf9, 0x9f, 0xac, 0x2d, 0x0b, 0xe1, 0x56, 0xf5, 0xa2, 0x51, 0xd2, 0x53, 0xbf, 0xf5,
	0xb8, 0x54, 0x75, 0x7e, 0x5e, 0x32, 0xb7, 0x54, 0xa6, 0xda, 0xa8, 0x48, 0x17, 0xa5, 0x5a, 0xd0,
	0x8a, 0x59, 0x07, 0x86, 0x3e, 0x3a, 0x3b, 0x3e, 0x79, 0xfc, 0xc5, 0x49, 0xda, 0x7b, 0x90, 0x4c,
	0x46, 0x61, 0x10, 0xa6, 0x2f, 0x32, 0xad, 0x4b, 0xc1, 0xfd, 0xfb, 0x42, 0xbf, 0xb5, 0x4a, 0xce,
	0xb6, 0x90, 0xec, 0x43, 0xd4, 0x9b, 0x4e, 0xa6, 0x78, 0x8a, 0x46, 0x19, 0xb8, 0xda, 0x48, 0xc8,
	0xc9, 0xb3, 0x15, 0x48, 0xe2, 0x56, 0x40, 0x0c, 0x58, 0x55, 0x1b, 0x0e, 0x24, 0x57, 0x60, 0x89,
	0x54, 0x8e, 0xc0, 0xf7, 0xc2, 0xba, 0x04, 0xdf, 0x41, 0xb7, 0x7e, 0x0e, 0x83, 0x3d, 0xf3, 0x11,
	0x1a, 0x74, 0x13, 0x21, 0x1f, 0x2b, 0x5e, 0x37, 0xdf, 0x8f, 0x77, 0xc7, 0x6f, 0xee, 0x9e, 0x0f,
	0xb5, 0xc2, 0x01, 0xcd, 0x15, 0xb7, 0xf4, 0x6b, 0x2c, 0xa4, 0x03, 0x23, 0x59, 0x49, 0x97, 0xb5,
	0xab, 0x0d, 0x50, 0xbd, 0x58, 0xdc, 0xf1, 0xcf, 0xe0, 0xfb, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xe6, 0xb3, 0xfe, 0xce, 0x80, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BackendClient is the client API for Backend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BackendClient interface {
	// GetMatches triggers execution of the specfied MatchFunction for each of the
	// specified MatchProfiles. Each MatchFunction execution returns a set of
	// proposals which are then evaluated to generate results. GetMatches method
	// streams these results back to the caller.
	// TODO: Should this be renamed to createProposal? It's not a "Get" if it's
	// executing a MatchFunction.
	GetMatches(ctx context.Context, in *GetMatchesRequest, opts ...grpc.CallOption) (Backend_GetMatchesClient, error)
	// AssignTickets sets the specified Assignment on the Tickets for the Ticket
	// IDs passed.
	AssignTickets(ctx context.Context, in *AssignTicketsRequest, opts ...grpc.CallOption) (*AssignTicketsResponse, error)
}

type backendClient struct {
	cc *grpc.ClientConn
}

func NewBackendClient(cc *grpc.ClientConn) BackendClient {
	return &backendClient{cc}
}

func (c *backendClient) GetMatches(ctx context.Context, in *GetMatchesRequest, opts ...grpc.CallOption) (Backend_GetMatchesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Backend_serviceDesc.Streams[0], "/api.Backend/GetMatches", opts...)
	if err != nil {
		return nil, err
	}
	x := &backendGetMatchesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Backend_GetMatchesClient interface {
	Recv() (*GetMatchesResponse, error)
	grpc.ClientStream
}

type backendGetMatchesClient struct {
	grpc.ClientStream
}

func (x *backendGetMatchesClient) Recv() (*GetMatchesResponse, error) {
	m := new(GetMatchesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *backendClient) AssignTickets(ctx context.Context, in *AssignTicketsRequest, opts ...grpc.CallOption) (*AssignTicketsResponse, error) {
	out := new(AssignTicketsResponse)
	err := c.cc.Invoke(ctx, "/api.Backend/AssignTickets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackendServer is the server API for Backend service.
type BackendServer interface {
	// GetMatches triggers execution of the specfied MatchFunction for each of the
	// specified MatchProfiles. Each MatchFunction execution returns a set of
	// proposals which are then evaluated to generate results. GetMatches method
	// streams these results back to the caller.
	// TODO: Should this be renamed to createProposal? It's not a "Get" if it's
	// executing a MatchFunction.
	GetMatches(*GetMatchesRequest, Backend_GetMatchesServer) error
	// AssignTickets sets the specified Assignment on the Tickets for the Ticket
	// IDs passed.
	AssignTickets(context.Context, *AssignTicketsRequest) (*AssignTicketsResponse, error)
}

func RegisterBackendServer(s *grpc.Server, srv BackendServer) {
	s.RegisterService(&_Backend_serviceDesc, srv)
}

func _Backend_GetMatches_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetMatchesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BackendServer).GetMatches(m, &backendGetMatchesServer{stream})
}

type Backend_GetMatchesServer interface {
	Send(*GetMatchesResponse) error
	grpc.ServerStream
}

type backendGetMatchesServer struct {
	grpc.ServerStream
}

func (x *backendGetMatchesServer) Send(m *GetMatchesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Backend_AssignTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).AssignTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Backend/AssignTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).AssignTickets(ctx, req.(*AssignTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Backend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Backend",
	HandlerType: (*BackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignTickets",
			Handler:    _Backend_AssignTickets_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMatches",
			Handler:       _Backend_GetMatches_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/backend.proto",
}
