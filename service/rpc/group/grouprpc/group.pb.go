// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/group.proto

package admin

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type GroupParams struct {
	Rid                  string   `protobuf:"bytes,3,opt,name=rid,proto3" json:"rid,omitempty"`
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	GroupId              string   `protobuf:"bytes,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupParams) Reset()         { *m = GroupParams{} }
func (m *GroupParams) String() string { return proto.CompactTextString(m) }
func (*GroupParams) ProtoMessage()    {}
func (*GroupParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_73aa59f1565abd82, []int{0}
}

func (m *GroupParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupParams.Unmarshal(m, b)
}
func (m *GroupParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupParams.Marshal(b, m, deterministic)
}
func (m *GroupParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupParams.Merge(m, src)
}
func (m *GroupParams) XXX_Size() int {
	return xxx_messageInfo_GroupParams.Size(m)
}
func (m *GroupParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupParams.DiscardUnknown(m)
}

var xxx_messageInfo_GroupParams proto.InternalMessageInfo

func (m *GroupParams) GetRid() string {
	if m != nil {
		return m.Rid
	}
	return ""
}

func (m *GroupParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GroupParams) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

type Group struct {
	Gid                  string   `protobuf:"bytes,1,opt,name=gid,proto3" json:"gid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Create               int64    `protobuf:"varint,4,opt,name=create,proto3" json:"create,omitempty"`
	Update               int64    `protobuf:"varint,5,opt,name=update,proto3" json:"update,omitempty"`
	IsValid              int64    `protobuf:"varint,6,opt,name=isValid,proto3" json:"isValid,omitempty"`
	IsDel                int64    `protobuf:"varint,7,opt,name=isDel,proto3" json:"isDel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_73aa59f1565abd82, []int{1}
}

func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetGid() string {
	if m != nil {
		return m.Gid
	}
	return ""
}

func (m *Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Group) GetCreate() int64 {
	if m != nil {
		return m.Create
	}
	return 0
}

func (m *Group) GetUpdate() int64 {
	if m != nil {
		return m.Update
	}
	return 0
}

func (m *Group) GetIsValid() int64 {
	if m != nil {
		return m.IsValid
	}
	return 0
}

func (m *Group) GetIsDel() int64 {
	if m != nil {
		return m.IsDel
	}
	return 0
}

//这里返回结果应该是一个Role类型
type GroupPn struct {
	Pn                   int64    `protobuf:"varint,1,opt,name=pn,proto3" json:"pn,omitempty"`
	Ps                   int64    `protobuf:"varint,2,opt,name=ps,proto3" json:"ps,omitempty"`
	Group                []*Group `protobuf:"bytes,3,rep,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupPn) Reset()         { *m = GroupPn{} }
func (m *GroupPn) String() string { return proto.CompactTextString(m) }
func (*GroupPn) ProtoMessage()    {}
func (*GroupPn) Descriptor() ([]byte, []int) {
	return fileDescriptor_73aa59f1565abd82, []int{2}
}

func (m *GroupPn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupPn.Unmarshal(m, b)
}
func (m *GroupPn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupPn.Marshal(b, m, deterministic)
}
func (m *GroupPn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupPn.Merge(m, src)
}
func (m *GroupPn) XXX_Size() int {
	return xxx_messageInfo_GroupPn.Size(m)
}
func (m *GroupPn) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupPn.DiscardUnknown(m)
}

var xxx_messageInfo_GroupPn proto.InternalMessageInfo

func (m *GroupPn) GetPn() int64 {
	if m != nil {
		return m.Pn
	}
	return 0
}

func (m *GroupPn) GetPs() int64 {
	if m != nil {
		return m.Ps
	}
	return 0
}

func (m *GroupPn) GetGroup() []*Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func init() {
	proto.RegisterType((*GroupParams)(nil), "admin.GroupParams")
	proto.RegisterType((*Group)(nil), "admin.Group")
	proto.RegisterType((*GroupPn)(nil), "admin.GroupPn")
}

func init() { proto.RegisterFile("proto/group.proto", fileDescriptor_73aa59f1565abd82) }

var fileDescriptor_73aa59f1565abd82 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4a, 0x3b, 0x31,
	0x10, 0xc6, 0xff, 0xbb, 0xe9, 0xb6, 0xfd, 0xcf, 0x96, 0xaa, 0x41, 0x24, 0xf4, 0x54, 0xf6, 0xd4,
	0x8b, 0x15, 0x6b, 0x5f, 0x40, 0x28, 0x14, 0x0f, 0x05, 0x59, 0xd1, 0x7b, 0x6c, 0x86, 0x12, 0xd8,
	0xcd, 0x86, 0x64, 0xd7, 0xb7, 0xf0, 0x0d, 0x7d, 0x18, 0xc9, 0x64, 0x2b, 0x15, 0xbc, 0xf4, 0x36,
	0xdf, 0x6f, 0xf2, 0x7d, 0xc9, 0x0c, 0x81, 0x2b, 0xeb, 0x9a, 0xb6, 0xb9, 0x3b, 0xb8, 0xa6, 0xb3,
	0x4b, 0xaa, 0x79, 0x26, 0x55, 0xad, 0xcd, 0xac, 0xef, 0x74, 0xad, 0xae, 0x7c, 0xec, 0x14, 0x3b,
	0xc8, 0xb7, 0xe1, 0xe0, 0xb3, 0x74, 0xb2, 0xf6, 0xfc, 0x12, 0x98, 0xd3, 0x4a, 0xb0, 0x79, 0xb2,
	0xf8, 0x5f, 0x86, 0x92, 0x73, 0x18, 0x18, 0x59, 0xa3, 0x48, 0x08, 0x51, 0xcd, 0x05, 0x8c, 0x28,
	0xfd, 0x49, 0x89, 0x94, 0xf0, 0x51, 0x16, 0x9f, 0x09, 0x64, 0x94, 0x17, 0x92, 0x0e, 0x5a, 0xf5,
	0xb6, 0x50, 0xfe, 0x24, 0xa5, 0x27, 0x49, 0x37, 0x30, 0xdc, 0x3b, 0x94, 0x2d, 0x8a, 0xc1, 0x3c,
	0x59, 0xb0, 0xb2, 0x57, 0x81, 0x77, 0x56, 0x05, 0x9e, 0x45, 0x1e, 0x55, 0xb8, 0x59, 0xfb, 0x37,
	0x59, 0x69, 0x25, 0x86, 0xd4, 0x38, 0x4a, 0x7e, 0x0d, 0x99, 0xf6, 0x1b, 0xac, 0xc4, 0x88, 0x78,
	0x14, 0xc5, 0x0e, 0x46, 0x71, 0x3c, 0xc3, 0xa7, 0x90, 0x5a, 0x43, 0xef, 0x61, 0x65, 0x6a, 0xa3,
	0xf6, 0xf4, 0x98, 0xa0, 0x3d, 0x2f, 0x20, 0xa3, 0x29, 0x04, 0x9b, 0xb3, 0x45, 0xbe, 0x9a, 0x2c,
	0x69, 0x67, 0x4b, 0xb2, 0x97, 0xb1, 0xb5, 0xfa, 0x4a, 0x60, 0x42, 0xe0, 0x05, 0xdd, 0x87, 0xde,
	0x23, 0xbf, 0x87, 0xf1, 0xa3, 0x52, 0x71, 0x62, 0x7e, 0xea, 0x88, 0xfb, 0x9c, 0x5d, 0xf4, 0xac,
	0x44, 0x6f, 0x1b, 0xe3, 0xb1, 0xf8, 0xc7, 0x6f, 0x61, 0xbc, 0xc5, 0x36, 0x5a, 0xa6, 0xbf, 0x2c,
	0xe6, 0xaf, 0xe3, 0x6b, 0xc8, 0x37, 0x58, 0x61, 0x8b, 0x67, 0x5d, 0xb2, 0x86, 0xfc, 0x95, 0x36,
	0x76, 0x8e, 0xeb, 0x7d, 0x48, 0x7f, 0xe2, 0xe1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xf8, 0xaf,
	0x94, 0x42, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for GroupService service

type GroupServiceClient interface {
	AddGroup(ctx context.Context, in *GroupParams, opts ...client.CallOption) (*Response, error)
	GetGroup(ctx context.Context, in *GroupPn, opts ...client.CallOption) (*Response, error)
	DeleteGroup(ctx context.Context, in *GroupParams, opts ...client.CallOption) (*Response, error)
	UpdateGroup(ctx context.Context, in *GroupParams, opts ...client.CallOption) (*Response, error)
}

type groupServiceClient struct {
	c           client.Client
	serviceName string
}

func NewGroupServiceClient(serviceName string, c client.Client) GroupServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "admin"
	}
	return &groupServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *groupServiceClient) AddGroup(ctx context.Context, in *GroupParams, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "GroupService.AddGroup", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) GetGroup(ctx context.Context, in *GroupPn, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "GroupService.GetGroup", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) DeleteGroup(ctx context.Context, in *GroupParams, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "GroupService.DeleteGroup", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupServiceClient) UpdateGroup(ctx context.Context, in *GroupParams, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "GroupService.UpdateGroup", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GroupService service

type GroupServiceHandler interface {
	AddGroup(context.Context, *GroupParams, *Response) error
	GetGroup(context.Context, *GroupPn, *Response) error
	DeleteGroup(context.Context, *GroupParams, *Response) error
	UpdateGroup(context.Context, *GroupParams, *Response) error
}

func RegisterGroupServiceHandler(s server.Server, hdlr GroupServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&GroupService{hdlr}, opts...))
}

type GroupService struct {
	GroupServiceHandler
}

func (h *GroupService) AddGroup(ctx context.Context, in *GroupParams, out *Response) error {
	return h.GroupServiceHandler.AddGroup(ctx, in, out)
}

func (h *GroupService) GetGroup(ctx context.Context, in *GroupPn, out *Response) error {
	return h.GroupServiceHandler.GetGroup(ctx, in, out)
}

func (h *GroupService) DeleteGroup(ctx context.Context, in *GroupParams, out *Response) error {
	return h.GroupServiceHandler.DeleteGroup(ctx, in, out)
}

func (h *GroupService) UpdateGroup(ctx context.Context, in *GroupParams, out *Response) error {
	return h.GroupServiceHandler.UpdateGroup(ctx, in, out)
}
