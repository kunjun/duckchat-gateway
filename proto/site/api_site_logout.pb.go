// Code generated by protoc-gen-go. DO NOT EDIT.
// source: site/api_site_logout.proto

package site // import "github.com/duckchat/duckchat-gateway/proto/site"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// *
//
// action: api.site.logout
//
type ApiSiteLogoutRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiSiteLogoutRequest) Reset()         { *m = ApiSiteLogoutRequest{} }
func (m *ApiSiteLogoutRequest) String() string { return proto.CompactTextString(m) }
func (*ApiSiteLogoutRequest) ProtoMessage()    {}
func (*ApiSiteLogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_site_logout_93e04abc604451a9, []int{0}
}
func (m *ApiSiteLogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiSiteLogoutRequest.Unmarshal(m, b)
}
func (m *ApiSiteLogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiSiteLogoutRequest.Marshal(b, m, deterministic)
}
func (dst *ApiSiteLogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiSiteLogoutRequest.Merge(dst, src)
}
func (m *ApiSiteLogoutRequest) XXX_Size() int {
	return xxx_messageInfo_ApiSiteLogoutRequest.Size(m)
}
func (m *ApiSiteLogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiSiteLogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApiSiteLogoutRequest proto.InternalMessageInfo

type ApiSiteLogoutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiSiteLogoutResponse) Reset()         { *m = ApiSiteLogoutResponse{} }
func (m *ApiSiteLogoutResponse) String() string { return proto.CompactTextString(m) }
func (*ApiSiteLogoutResponse) ProtoMessage()    {}
func (*ApiSiteLogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_site_logout_93e04abc604451a9, []int{1}
}
func (m *ApiSiteLogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiSiteLogoutResponse.Unmarshal(m, b)
}
func (m *ApiSiteLogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiSiteLogoutResponse.Marshal(b, m, deterministic)
}
func (dst *ApiSiteLogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiSiteLogoutResponse.Merge(dst, src)
}
func (m *ApiSiteLogoutResponse) XXX_Size() int {
	return xxx_messageInfo_ApiSiteLogoutResponse.Size(m)
}
func (m *ApiSiteLogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiSiteLogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApiSiteLogoutResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ApiSiteLogoutRequest)(nil), "site.ApiSiteLogoutRequest")
	proto.RegisterType((*ApiSiteLogoutResponse)(nil), "site.ApiSiteLogoutResponse")
}

func init() {
	proto.RegisterFile("site/api_site_logout.proto", fileDescriptor_api_site_logout_93e04abc604451a9)
}

var fileDescriptor_api_site_logout_93e04abc604451a9 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0xce, 0x2c, 0x49,
	0xd5, 0x4f, 0x2c, 0xc8, 0x8c, 0x07, 0x31, 0xe2, 0x73, 0xf2, 0xd3, 0xf3, 0x4b, 0x4b, 0xf4, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x58, 0x40, 0x42, 0x4a, 0x62, 0x5c, 0x22, 0x8e, 0x05, 0x99, 0xc1,
	0x99, 0x25, 0xa9, 0x3e, 0x60, 0xc9, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x25, 0x71, 0x2e,
	0x51, 0x34, 0xf1, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xa7, 0x08, 0x2e, 0xe1, 0xe4, 0xfc, 0x5c,
	0xbd, 0xaa, 0xc4, 0x9c, 0x4a, 0x88, 0x41, 0x7a, 0x20, 0x73, 0xa2, 0xf4, 0xd3, 0x33, 0x4b, 0x32,
	0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x53, 0x4a, 0x93, 0xb3, 0x93, 0x33, 0x12, 0x4b, 0xe0,
	0x0c, 0xdd, 0xf4, 0xc4, 0x92, 0xd4, 0xf2, 0xc4, 0x4a, 0x7d, 0xb0, 0x06, 0x7d, 0x90, 0x86, 0x53,
	0x4c, 0xfc, 0x51, 0x89, 0x39, 0x95, 0x31, 0x01, 0x20, 0x91, 0x18, 0x90, 0x3d, 0x49, 0x6c, 0x60,
	0x59, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x68, 0x91, 0x8e, 0xb5, 0x00, 0x00, 0x00,
}
