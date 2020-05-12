// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bot_message.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Bot struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Command              []*Command `protobuf:"bytes,6,rep,name=command,proto3" json:"command,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Bot) Reset()         { *m = Bot{} }
func (m *Bot) String() string { return proto.CompactTextString(m) }
func (*Bot) ProtoMessage()    {}
func (*Bot) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee8f4877009a2ebf, []int{0}
}

func (m *Bot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bot.Unmarshal(m, b)
}
func (m *Bot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bot.Marshal(b, m, deterministic)
}
func (m *Bot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bot.Merge(m, src)
}
func (m *Bot) XXX_Size() int {
	return xxx_messageInfo_Bot.Size(m)
}
func (m *Bot) XXX_DiscardUnknown() {
	xxx_messageInfo_Bot.DiscardUnknown(m)
}

var xxx_messageInfo_Bot proto.InternalMessageInfo

func (m *Bot) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Bot) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Bot) GetCommand() []*Command {
	if m != nil {
		return m.Command
	}
	return nil
}

type Command struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Answer               string   `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee8f4877009a2ebf, []int{1}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Command) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

type NewBotRequest struct {
	Bot                  *Bot     `protobuf:"bytes,1,opt,name=bot,proto3" json:"bot,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewBotRequest) Reset()         { *m = NewBotRequest{} }
func (m *NewBotRequest) String() string { return proto.CompactTextString(m) }
func (*NewBotRequest) ProtoMessage()    {}
func (*NewBotRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee8f4877009a2ebf, []int{2}
}

func (m *NewBotRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewBotRequest.Unmarshal(m, b)
}
func (m *NewBotRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewBotRequest.Marshal(b, m, deterministic)
}
func (m *NewBotRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewBotRequest.Merge(m, src)
}
func (m *NewBotRequest) XXX_Size() int {
	return xxx_messageInfo_NewBotRequest.Size(m)
}
func (m *NewBotRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewBotRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewBotRequest proto.InternalMessageInfo

func (m *NewBotRequest) GetBot() *Bot {
	if m != nil {
		return m.Bot
	}
	return nil
}

type NewBotResponse struct {
	BotId                string   `protobuf:"bytes,1,opt,name=bot_id,json=botId,proto3" json:"bot_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewBotResponse) Reset()         { *m = NewBotResponse{} }
func (m *NewBotResponse) String() string { return proto.CompactTextString(m) }
func (*NewBotResponse) ProtoMessage()    {}
func (*NewBotResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee8f4877009a2ebf, []int{3}
}

func (m *NewBotResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewBotResponse.Unmarshal(m, b)
}
func (m *NewBotResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewBotResponse.Marshal(b, m, deterministic)
}
func (m *NewBotResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewBotResponse.Merge(m, src)
}
func (m *NewBotResponse) XXX_Size() int {
	return xxx_messageInfo_NewBotResponse.Size(m)
}
func (m *NewBotResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewBotResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewBotResponse proto.InternalMessageInfo

func (m *NewBotResponse) GetBotId() string {
	if m != nil {
		return m.BotId
	}
	return ""
}

type AnswerRequest struct {
	BotName              string   `protobuf:"bytes,1,opt,name=bot_name,json=botName,proto3" json:"bot_name,omitempty"`
	Command              string   `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnswerRequest) Reset()         { *m = AnswerRequest{} }
func (m *AnswerRequest) String() string { return proto.CompactTextString(m) }
func (*AnswerRequest) ProtoMessage()    {}
func (*AnswerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee8f4877009a2ebf, []int{4}
}

func (m *AnswerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnswerRequest.Unmarshal(m, b)
}
func (m *AnswerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnswerRequest.Marshal(b, m, deterministic)
}
func (m *AnswerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnswerRequest.Merge(m, src)
}
func (m *AnswerRequest) XXX_Size() int {
	return xxx_messageInfo_AnswerRequest.Size(m)
}
func (m *AnswerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnswerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnswerRequest proto.InternalMessageInfo

func (m *AnswerRequest) GetBotName() string {
	if m != nil {
		return m.BotName
	}
	return ""
}

func (m *AnswerRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

type AnswerResponse struct {
	Answer               string   `protobuf:"bytes,1,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnswerResponse) Reset()         { *m = AnswerResponse{} }
func (m *AnswerResponse) String() string { return proto.CompactTextString(m) }
func (*AnswerResponse) ProtoMessage()    {}
func (*AnswerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee8f4877009a2ebf, []int{5}
}

func (m *AnswerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnswerResponse.Unmarshal(m, b)
}
func (m *AnswerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnswerResponse.Marshal(b, m, deterministic)
}
func (m *AnswerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnswerResponse.Merge(m, src)
}
func (m *AnswerResponse) XXX_Size() int {
	return xxx_messageInfo_AnswerResponse.Size(m)
}
func (m *AnswerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AnswerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AnswerResponse proto.InternalMessageInfo

func (m *AnswerResponse) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func init() {
	proto.RegisterType((*Bot)(nil), "education.code.codeedu.Bot")
	proto.RegisterType((*Command)(nil), "education.code.codeedu.Command")
	proto.RegisterType((*NewBotRequest)(nil), "education.code.codeedu.NewBotRequest")
	proto.RegisterType((*NewBotResponse)(nil), "education.code.codeedu.NewBotResponse")
	proto.RegisterType((*AnswerRequest)(nil), "education.code.codeedu.AnswerRequest")
	proto.RegisterType((*AnswerResponse)(nil), "education.code.codeedu.AnswerResponse")
}

func init() {
	proto.RegisterFile("bot_message.proto", fileDescriptor_ee8f4877009a2ebf)
}

var fileDescriptor_ee8f4877009a2ebf = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3f, 0x4b, 0xc3, 0x40,
	0x14, 0x6f, 0xda, 0x9a, 0xda, 0x57, 0x5a, 0xf0, 0xc0, 0x12, 0x75, 0xb0, 0x1c, 0xa8, 0x5d, 0xcc,
	0x50, 0x27, 0x11, 0x04, 0x53, 0x17, 0x97, 0x0e, 0x75, 0x52, 0x04, 0x49, 0x72, 0x0f, 0xc9, 0x90,
	0xbc, 0x98, 0xbb, 0xb4, 0x9f, 0xd0, 0xef, 0x25, 0x77, 0xc9, 0xd1, 0x04, 0x0c, 0x2e, 0x21, 0xf7,
	0xf2, 0x7b, 0xbf, 0x7f, 0x39, 0x38, 0x89, 0x48, 0x7d, 0xa6, 0x28, 0x65, 0xf8, 0x85, 0x7e, 0x5e,
	0x90, 0x22, 0x36, 0x47, 0x51, 0xc6, 0xa1, 0x4a, 0x28, 0xf3, 0x63, 0x12, 0x68, 0x1e, 0x28, 0x4a,
	0xbe, 0x83, 0x41, 0x40, 0x8a, 0x31, 0x18, 0x66, 0x61, 0x8a, 0x9e, 0xb3, 0x70, 0x96, 0xe3, 0xad,
	0x79, 0x67, 0x0b, 0x98, 0x08, 0x94, 0x71, 0x91, 0xe4, 0x7a, 0xcd, 0xeb, 0x9b, 0x4f, 0xcd, 0x11,
	0xbb, 0x87, 0x51, 0x4c, 0x69, 0x1a, 0x66, 0xc2, 0x73, 0x17, 0x83, 0xe5, 0x64, 0x75, 0xe9, 0xff,
	0x2d, 0xe3, 0xaf, 0x2b, 0xd8, 0xd6, 0xe2, 0xf9, 0x03, 0x8c, 0xea, 0x19, 0xf3, 0x0e, 0x2c, 0x95,
	0xbc, 0x3d, 0xb2, 0x39, 0xb8, 0x61, 0x26, 0xf7, 0x58, 0xd4, 0xe2, 0xf5, 0x89, 0x3f, 0xc2, 0x74,
	0x83, 0xfb, 0x80, 0xd4, 0x16, 0xbf, 0x4b, 0x94, 0x8a, 0xdd, 0xc2, 0x20, 0x22, 0x65, 0xd6, 0x27,
	0xab, 0x8b, 0x2e, 0x13, 0x7a, 0x41, 0xe3, 0xf8, 0x0d, 0xcc, 0xec, 0xbe, 0xcc, 0x29, 0x93, 0xc8,
	0x4e, 0xc1, 0xd5, 0x9d, 0x25, 0xd6, 0xc2, 0x51, 0x44, 0xea, 0x45, 0xf0, 0x67, 0x98, 0x3e, 0x19,
	0x49, 0x2b, 0x74, 0x06, 0xc7, 0x1a, 0xd7, 0xe8, 0x6a, 0x14, 0x91, 0xda, 0xe8, 0xba, 0x1a, 0x31,
	0xfa, 0xad, 0x18, 0x7c, 0x09, 0x33, 0xcb, 0x52, 0xcb, 0x1d, 0x82, 0x39, 0xcd, 0x60, 0xab, 0x1f,
	0x07, 0x20, 0x20, 0xf5, 0x8a, 0xc5, 0x2e, 0x89, 0x91, 0x7d, 0xc0, 0x78, 0x5d, 0x60, 0xa8, 0x50,
	0xff, 0xa2, 0xab, 0xae, 0x58, 0xad, 0x2a, 0xce, 0xaf, 0xff, 0x83, 0x55, 0x16, 0x78, 0x8f, 0xbd,
	0x81, 0x5b, 0xd9, 0xea, 0xa6, 0x6e, 0x85, 0xef, 0xa6, 0x6e, 0xa7, 0xe3, 0xbd, 0x60, 0xf8, 0xde,
	0xcf, 0xa3, 0xc8, 0x35, 0x57, 0xef, 0xee, 0x37, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xd9, 0x7a, 0xd4,
	0x8f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BotServiceClient is the client API for BotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BotServiceClient interface {
	CreateBot(ctx context.Context, in *NewBotRequest, opts ...grpc.CallOption) (*NewBotResponse, error)
	Answer(ctx context.Context, in *AnswerRequest, opts ...grpc.CallOption) (*AnswerResponse, error)
}

type botServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBotServiceClient(cc grpc.ClientConnInterface) BotServiceClient {
	return &botServiceClient{cc}
}

func (c *botServiceClient) CreateBot(ctx context.Context, in *NewBotRequest, opts ...grpc.CallOption) (*NewBotResponse, error) {
	out := new(NewBotResponse)
	err := c.cc.Invoke(ctx, "/education.code.codeedu.BotService/CreateBot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServiceClient) Answer(ctx context.Context, in *AnswerRequest, opts ...grpc.CallOption) (*AnswerResponse, error) {
	out := new(AnswerResponse)
	err := c.cc.Invoke(ctx, "/education.code.codeedu.BotService/Answer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotServiceServer is the server API for BotService service.
type BotServiceServer interface {
	CreateBot(context.Context, *NewBotRequest) (*NewBotResponse, error)
	Answer(context.Context, *AnswerRequest) (*AnswerResponse, error)
}

// UnimplementedBotServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBotServiceServer struct {
}

func (*UnimplementedBotServiceServer) CreateBot(ctx context.Context, req *NewBotRequest) (*NewBotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBot not implemented")
}
func (*UnimplementedBotServiceServer) Answer(ctx context.Context, req *AnswerRequest) (*AnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Answer not implemented")
}

func RegisterBotServiceServer(s *grpc.Server, srv BotServiceServer) {
	s.RegisterService(&_BotService_serviceDesc, srv)
}

func _BotService_CreateBot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewBotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServiceServer).CreateBot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/education.code.codeedu.BotService/CreateBot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServiceServer).CreateBot(ctx, req.(*NewBotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotService_Answer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServiceServer).Answer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/education.code.codeedu.BotService/Answer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServiceServer).Answer(ctx, req.(*AnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BotService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "education.code.codeedu.BotService",
	HandlerType: (*BotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBot",
			Handler:    _BotService_CreateBot_Handler,
		},
		{
			MethodName: "Answer",
			Handler:    _BotService_Answer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bot_message.proto",
}