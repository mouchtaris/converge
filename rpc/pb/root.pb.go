// Code generated by protoc-gen-go.
// source: root.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	root.proto

It has these top-level messages:
	LoadRequest
	ContentResponse
	StatusResponse
	DiffResponse
	GraphComponent
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// the stage from which this status response is being sent
type StatusResponse_Stage int32

const (
	StatusResponse_UNSPECIFIED_STAGE StatusResponse_Stage = 0
	StatusResponse_PLAN              StatusResponse_Stage = 1
	StatusResponse_APPLY             StatusResponse_Stage = 2
)

var StatusResponse_Stage_name = map[int32]string{
	0: "UNSPECIFIED_STAGE",
	1: "PLAN",
	2: "APPLY",
}
var StatusResponse_Stage_value = map[string]int32{
	"UNSPECIFIED_STAGE": 0,
	"PLAN":              1,
	"APPLY":             2,
}

func (x StatusResponse_Stage) String() string {
	return proto.EnumName(StatusResponse_Stage_name, int32(x))
}
func (StatusResponse_Stage) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

// when is this status response being sent?
type StatusResponse_Run int32

const (
	StatusResponse_UNSPECIFIED_RUN StatusResponse_Run = 0
	StatusResponse_STARTED         StatusResponse_Run = 1
	StatusResponse_FINISHED        StatusResponse_Run = 2
)

var StatusResponse_Run_name = map[int32]string{
	0: "UNSPECIFIED_RUN",
	1: "STARTED",
	2: "FINISHED",
}
var StatusResponse_Run_value = map[string]int32{
	"UNSPECIFIED_RUN": 0,
	"STARTED":         1,
	"FINISHED":        2,
}

func (x StatusResponse_Run) String() string {
	return proto.EnumName(StatusResponse_Run_name, int32(x))
}
func (StatusResponse_Run) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 1} }

type LoadRequest struct {
	Location   string            `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	Parameters map[string]string `protobuf:"bytes,2,rep,name=parameters" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Verify     bool              `protobuf:"varint,3,opt,name=verify" json:"verify,omitempty"`
}

func (m *LoadRequest) Reset()                    { *m = LoadRequest{} }
func (m *LoadRequest) String() string            { return proto.CompactTextString(m) }
func (*LoadRequest) ProtoMessage()               {}
func (*LoadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoadRequest) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *LoadRequest) GetParameters() map[string]string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *LoadRequest) GetVerify() bool {
	if m != nil {
		return m.Verify
	}
	return false
}

type ContentResponse struct {
	Content string `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
}

func (m *ContentResponse) Reset()                    { *m = ContentResponse{} }
func (m *ContentResponse) String() string            { return proto.CompactTextString(m) }
func (*ContentResponse) ProtoMessage()               {}
func (*ContentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ContentResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type StatusResponse struct {
	// TODO: preserve for compat but will stop working in 0.4.0. This has moved to
	// meta.id
	Id      string                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Stage   StatusResponse_Stage    `protobuf:"varint,2,opt,name=stage,enum=pb.StatusResponse_Stage" json:"stage,omitempty"`
	Run     StatusResponse_Run      `protobuf:"varint,3,opt,name=run,enum=pb.StatusResponse_Run" json:"run,omitempty"`
	Details *StatusResponse_Details `protobuf:"bytes,4,opt,name=details" json:"details,omitempty"`
	Meta    *StatusResponse_Meta    `protobuf:"bytes,5,opt,name=meta" json:"meta,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StatusResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StatusResponse) GetStage() StatusResponse_Stage {
	if m != nil {
		return m.Stage
	}
	return StatusResponse_UNSPECIFIED_STAGE
}

func (m *StatusResponse) GetRun() StatusResponse_Run {
	if m != nil {
		return m.Run
	}
	return StatusResponse_UNSPECIFIED_RUN
}

func (m *StatusResponse) GetDetails() *StatusResponse_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (m *StatusResponse) GetMeta() *StatusResponse_Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

// the informational message, if present
type StatusResponse_Details struct {
	Messages   []string                 `protobuf:"bytes,1,rep,name=messages" json:"messages,omitempty"`
	Changes    map[string]*DiffResponse `protobuf:"bytes,2,rep,name=changes" json:"changes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	HasChanges bool                     `protobuf:"varint,3,opt,name=hasChanges" json:"hasChanges,omitempty"`
	Error      string                   `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
	Warning    string                   `protobuf:"bytes,5,opt,name=warning" json:"warning,omitempty"`
}

func (m *StatusResponse_Details) Reset()                    { *m = StatusResponse_Details{} }
func (m *StatusResponse_Details) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse_Details) ProtoMessage()               {}
func (*StatusResponse_Details) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *StatusResponse_Details) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *StatusResponse_Details) GetChanges() map[string]*DiffResponse {
	if m != nil {
		return m.Changes
	}
	return nil
}

func (m *StatusResponse_Details) GetHasChanges() bool {
	if m != nil {
		return m.HasChanges
	}
	return false
}

func (m *StatusResponse_Details) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type StatusResponse_Meta struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *StatusResponse_Meta) Reset()                    { *m = StatusResponse_Meta{} }
func (m *StatusResponse_Meta) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse_Meta) ProtoMessage()               {}
func (*StatusResponse_Meta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 1} }

func (m *StatusResponse_Meta) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DiffResponse struct {
	Original string `protobuf:"bytes,1,opt,name=original" json:"original,omitempty"`
	Current  string `protobuf:"bytes,2,opt,name=current" json:"current,omitempty"`
	Changes  bool   `protobuf:"varint,3,opt,name=changes" json:"changes,omitempty"`
}

func (m *DiffResponse) Reset()                    { *m = DiffResponse{} }
func (m *DiffResponse) String() string            { return proto.CompactTextString(m) }
func (*DiffResponse) ProtoMessage()               {}
func (*DiffResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DiffResponse) GetOriginal() string {
	if m != nil {
		return m.Original
	}
	return ""
}

func (m *DiffResponse) GetCurrent() string {
	if m != nil {
		return m.Current
	}
	return ""
}

func (m *DiffResponse) GetChanges() bool {
	if m != nil {
		return m.Changes
	}
	return false
}

type GraphComponent struct {
	// Types that are valid to be assigned to Component:
	//	*GraphComponent_Vertex_
	//	*GraphComponent_Edge_
	Component isGraphComponent_Component `protobuf_oneof:"component"`
}

func (m *GraphComponent) Reset()                    { *m = GraphComponent{} }
func (m *GraphComponent) String() string            { return proto.CompactTextString(m) }
func (*GraphComponent) ProtoMessage()               {}
func (*GraphComponent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isGraphComponent_Component interface {
	isGraphComponent_Component()
}

type GraphComponent_Vertex_ struct {
	Vertex *GraphComponent_Vertex `protobuf:"bytes,1,opt,name=vertex,oneof"`
}
type GraphComponent_Edge_ struct {
	Edge *GraphComponent_Edge `protobuf:"bytes,2,opt,name=edge,oneof"`
}

func (*GraphComponent_Vertex_) isGraphComponent_Component() {}
func (*GraphComponent_Edge_) isGraphComponent_Component()   {}

func (m *GraphComponent) GetComponent() isGraphComponent_Component {
	if m != nil {
		return m.Component
	}
	return nil
}

func (m *GraphComponent) GetVertex() *GraphComponent_Vertex {
	if x, ok := m.GetComponent().(*GraphComponent_Vertex_); ok {
		return x.Vertex
	}
	return nil
}

func (m *GraphComponent) GetEdge() *GraphComponent_Edge {
	if x, ok := m.GetComponent().(*GraphComponent_Edge_); ok {
		return x.Edge
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GraphComponent) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GraphComponent_OneofMarshaler, _GraphComponent_OneofUnmarshaler, _GraphComponent_OneofSizer, []interface{}{
		(*GraphComponent_Vertex_)(nil),
		(*GraphComponent_Edge_)(nil),
	}
}

func _GraphComponent_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GraphComponent)
	// component
	switch x := m.Component.(type) {
	case *GraphComponent_Vertex_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Vertex); err != nil {
			return err
		}
	case *GraphComponent_Edge_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Edge); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GraphComponent.Component has unexpected type %T", x)
	}
	return nil
}

func _GraphComponent_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GraphComponent)
	switch tag {
	case 1: // component.vertex
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GraphComponent_Vertex)
		err := b.DecodeMessage(msg)
		m.Component = &GraphComponent_Vertex_{msg}
		return true, err
	case 2: // component.edge
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GraphComponent_Edge)
		err := b.DecodeMessage(msg)
		m.Component = &GraphComponent_Edge_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GraphComponent_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GraphComponent)
	// component
	switch x := m.Component.(type) {
	case *GraphComponent_Vertex_:
		s := proto.Size(x.Vertex)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GraphComponent_Edge_:
		s := proto.Size(x.Edge)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type GraphComponent_Vertex struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// the kind of node, specified as the type used to create a node of this
	// type in the Converge DSL
	Kind string `protobuf:"bytes,2,opt,name=kind" json:"kind,omitempty"`
	// detailed fields of this node, serialized as JSON
	Details []byte `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
}

func (m *GraphComponent_Vertex) Reset()                    { *m = GraphComponent_Vertex{} }
func (m *GraphComponent_Vertex) String() string            { return proto.CompactTextString(m) }
func (*GraphComponent_Vertex) ProtoMessage()               {}
func (*GraphComponent_Vertex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *GraphComponent_Vertex) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GraphComponent_Vertex) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *GraphComponent_Vertex) GetDetails() []byte {
	if m != nil {
		return m.Details
	}
	return nil
}

type GraphComponent_Edge struct {
	Source     string   `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Dest       string   `protobuf:"bytes,2,opt,name=dest" json:"dest,omitempty"`
	Attributes []string `protobuf:"bytes,3,rep,name=attributes" json:"attributes,omitempty"`
}

func (m *GraphComponent_Edge) Reset()                    { *m = GraphComponent_Edge{} }
func (m *GraphComponent_Edge) String() string            { return proto.CompactTextString(m) }
func (*GraphComponent_Edge) ProtoMessage()               {}
func (*GraphComponent_Edge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 1} }

func (m *GraphComponent_Edge) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *GraphComponent_Edge) GetDest() string {
	if m != nil {
		return m.Dest
	}
	return ""
}

func (m *GraphComponent_Edge) GetAttributes() []string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func init() {
	proto.RegisterType((*LoadRequest)(nil), "pb.LoadRequest")
	proto.RegisterType((*ContentResponse)(nil), "pb.ContentResponse")
	proto.RegisterType((*StatusResponse)(nil), "pb.StatusResponse")
	proto.RegisterType((*StatusResponse_Details)(nil), "pb.StatusResponse.Details")
	proto.RegisterType((*StatusResponse_Meta)(nil), "pb.StatusResponse.Meta")
	proto.RegisterType((*DiffResponse)(nil), "pb.DiffResponse")
	proto.RegisterType((*GraphComponent)(nil), "pb.GraphComponent")
	proto.RegisterType((*GraphComponent_Vertex)(nil), "pb.GraphComponent.Vertex")
	proto.RegisterType((*GraphComponent_Edge)(nil), "pb.GraphComponent.Edge")
	proto.RegisterEnum("pb.StatusResponse_Stage", StatusResponse_Stage_name, StatusResponse_Stage_value)
	proto.RegisterEnum("pb.StatusResponse_Run", StatusResponse_Run_name, StatusResponse_Run_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Executor service

type ExecutorClient interface {
	// Healthcheck a module given by the location
	HealthCheck(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (Executor_HealthCheckClient, error)
	// Plan out the execution of a module given by the location
	Plan(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (Executor_PlanClient, error)
	// Apply a module given by the location
	Apply(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (Executor_ApplyClient, error)
}

type executorClient struct {
	cc *grpc.ClientConn
}

func NewExecutorClient(cc *grpc.ClientConn) ExecutorClient {
	return &executorClient{cc}
}

func (c *executorClient) HealthCheck(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (Executor_HealthCheckClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Executor_serviceDesc.Streams[0], c.cc, "/pb.Executor/HealthCheck", opts...)
	if err != nil {
		return nil, err
	}
	x := &executorHealthCheckClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Executor_HealthCheckClient interface {
	Recv() (*StatusResponse, error)
	grpc.ClientStream
}

type executorHealthCheckClient struct {
	grpc.ClientStream
}

func (x *executorHealthCheckClient) Recv() (*StatusResponse, error) {
	m := new(StatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *executorClient) Plan(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (Executor_PlanClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Executor_serviceDesc.Streams[1], c.cc, "/pb.Executor/Plan", opts...)
	if err != nil {
		return nil, err
	}
	x := &executorPlanClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Executor_PlanClient interface {
	Recv() (*StatusResponse, error)
	grpc.ClientStream
}

type executorPlanClient struct {
	grpc.ClientStream
}

func (x *executorPlanClient) Recv() (*StatusResponse, error) {
	m := new(StatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *executorClient) Apply(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (Executor_ApplyClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Executor_serviceDesc.Streams[2], c.cc, "/pb.Executor/Apply", opts...)
	if err != nil {
		return nil, err
	}
	x := &executorApplyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Executor_ApplyClient interface {
	Recv() (*StatusResponse, error)
	grpc.ClientStream
}

type executorApplyClient struct {
	grpc.ClientStream
}

func (x *executorApplyClient) Recv() (*StatusResponse, error) {
	m := new(StatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Executor service

type ExecutorServer interface {
	// Healthcheck a module given by the location
	HealthCheck(*LoadRequest, Executor_HealthCheckServer) error
	// Plan out the execution of a module given by the location
	Plan(*LoadRequest, Executor_PlanServer) error
	// Apply a module given by the location
	Apply(*LoadRequest, Executor_ApplyServer) error
}

func RegisterExecutorServer(s *grpc.Server, srv ExecutorServer) {
	s.RegisterService(&_Executor_serviceDesc, srv)
}

func _Executor_HealthCheck_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LoadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExecutorServer).HealthCheck(m, &executorHealthCheckServer{stream})
}

type Executor_HealthCheckServer interface {
	Send(*StatusResponse) error
	grpc.ServerStream
}

type executorHealthCheckServer struct {
	grpc.ServerStream
}

func (x *executorHealthCheckServer) Send(m *StatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Executor_Plan_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LoadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExecutorServer).Plan(m, &executorPlanServer{stream})
}

type Executor_PlanServer interface {
	Send(*StatusResponse) error
	grpc.ServerStream
}

type executorPlanServer struct {
	grpc.ServerStream
}

func (x *executorPlanServer) Send(m *StatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Executor_Apply_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LoadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExecutorServer).Apply(m, &executorApplyServer{stream})
}

type Executor_ApplyServer interface {
	Send(*StatusResponse) error
	grpc.ServerStream
}

type executorApplyServer struct {
	grpc.ServerStream
}

func (x *executorApplyServer) Send(m *StatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Executor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Executor",
	HandlerType: (*ExecutorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HealthCheck",
			Handler:       _Executor_HealthCheck_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Plan",
			Handler:       _Executor_Plan_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Apply",
			Handler:       _Executor_Apply_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "root.proto",
}

// Client API for ResourceHost service

type ResourceHostClient interface {
	// GetBinary returns the converge binary itself
	GetBinary(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*ContentResponse, error)
	// GetModule gets the content of a module at the given path
	GetModule(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*ContentResponse, error)
}

type resourceHostClient struct {
	cc *grpc.ClientConn
}

func NewResourceHostClient(cc *grpc.ClientConn) ResourceHostClient {
	return &resourceHostClient{cc}
}

func (c *resourceHostClient) GetBinary(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*ContentResponse, error) {
	out := new(ContentResponse)
	err := grpc.Invoke(ctx, "/pb.ResourceHost/GetBinary", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceHostClient) GetModule(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*ContentResponse, error) {
	out := new(ContentResponse)
	err := grpc.Invoke(ctx, "/pb.ResourceHost/GetModule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ResourceHost service

type ResourceHostServer interface {
	// GetBinary returns the converge binary itself
	GetBinary(context.Context, *google_protobuf1.Empty) (*ContentResponse, error)
	// GetModule gets the content of a module at the given path
	GetModule(context.Context, *LoadRequest) (*ContentResponse, error)
}

func RegisterResourceHostServer(s *grpc.Server, srv ResourceHostServer) {
	s.RegisterService(&_ResourceHost_serviceDesc, srv)
}

func _ResourceHost_GetBinary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceHostServer).GetBinary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ResourceHost/GetBinary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceHostServer).GetBinary(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceHost_GetModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceHostServer).GetModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ResourceHost/GetModule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceHostServer).GetModule(ctx, req.(*LoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceHost_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ResourceHost",
	HandlerType: (*ResourceHostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBinary",
			Handler:    _ResourceHost_GetBinary_Handler,
		},
		{
			MethodName: "GetModule",
			Handler:    _ResourceHost_GetModule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "root.proto",
}

// Client API for Grapher service

type GrapherClient interface {
	Graph(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (Grapher_GraphClient, error)
}

type grapherClient struct {
	cc *grpc.ClientConn
}

func NewGrapherClient(cc *grpc.ClientConn) GrapherClient {
	return &grapherClient{cc}
}

func (c *grapherClient) Graph(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (Grapher_GraphClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Grapher_serviceDesc.Streams[0], c.cc, "/pb.Grapher/Graph", opts...)
	if err != nil {
		return nil, err
	}
	x := &grapherGraphClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Grapher_GraphClient interface {
	Recv() (*GraphComponent, error)
	grpc.ClientStream
}

type grapherGraphClient struct {
	grpc.ClientStream
}

func (x *grapherGraphClient) Recv() (*GraphComponent, error) {
	m := new(GraphComponent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Grapher service

type GrapherServer interface {
	Graph(*LoadRequest, Grapher_GraphServer) error
}

func RegisterGrapherServer(s *grpc.Server, srv GrapherServer) {
	s.RegisterService(&_Grapher_serviceDesc, srv)
}

func _Grapher_Graph_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LoadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GrapherServer).Graph(m, &grapherGraphServer{stream})
}

type Grapher_GraphServer interface {
	Send(*GraphComponent) error
	grpc.ServerStream
}

type grapherGraphServer struct {
	grpc.ServerStream
}

func (x *grapherGraphServer) Send(m *GraphComponent) error {
	return x.ServerStream.SendMsg(m)
}

var _Grapher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Grapher",
	HandlerType: (*GrapherServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Graph",
			Handler:       _Grapher_Graph_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "root.proto",
}

// Client API for Info service

type InfoClient interface {
	Ping(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type infoClient struct {
	cc *grpc.ClientConn
}

func NewInfoClient(cc *grpc.ClientConn) InfoClient {
	return &infoClient{cc}
}

func (c *infoClient) Ping(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/pb.Info/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Info service

type InfoServer interface {
	Ping(context.Context, *google_protobuf1.Empty) (*google_protobuf1.Empty, error)
}

func RegisterInfoServer(s *grpc.Server, srv InfoServer) {
	s.RegisterService(&_Info_serviceDesc, srv)
}

func _Info_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfoServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Info/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfoServer).Ping(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Info_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Info",
	HandlerType: (*InfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Info_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("root.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 957 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0xaf, 0x1d, 0xa7, 0x69, 0x5e, 0xa2, 0x34, 0xcc, 0x76, 0xbb, 0x5e, 0x2f, 0x62, 0x23, 0x1f,
	0x76, 0x4b, 0x57, 0x38, 0x90, 0x82, 0x84, 0x56, 0x5a, 0xa1, 0xb4, 0x4d, 0x9b, 0x48, 0xd9, 0x28,
	0x9a, 0x74, 0x91, 0xf8, 0x23, 0xd0, 0x24, 0x99, 0x38, 0x56, 0x9d, 0x19, 0x33, 0x1e, 0x97, 0x8d,
	0x10, 0x17, 0x0e, 0x1c, 0xb8, 0x72, 0xe6, 0xc3, 0xf0, 0x05, 0xb8, 0xf0, 0x15, 0x38, 0xf3, 0x19,
	0xd0, 0x8c, 0xed, 0x92, 0xa6, 0x89, 0xb4, 0x37, 0xff, 0x66, 0x7e, 0xef, 0xf7, 0xde, 0xfc, 0xde,
	0xf3, 0x0c, 0x80, 0xe0, 0x5c, 0x7a, 0x91, 0xe0, 0x92, 0x23, 0x33, 0x1a, 0x3b, 0xef, 0xfb, 0x9c,
	0xfb, 0x21, 0x6d, 0x92, 0x28, 0x68, 0x12, 0xc6, 0xb8, 0x24, 0x32, 0xe0, 0x2c, 0x4e, 0x19, 0xce,
	0x93, 0x6c, 0x57, 0xa3, 0x71, 0x32, 0x6b, 0xd2, 0x45, 0x24, 0x97, 0xe9, 0xa6, 0xfb, 0xa7, 0x01,
	0x95, 0x3e, 0x27, 0x53, 0x4c, 0x7f, 0x48, 0x68, 0x2c, 0x91, 0x03, 0x7b, 0x21, 0x9f, 0xe8, 0x78,
	0xdb, 0x68, 0x18, 0x47, 0x65, 0x7c, 0x8b, 0xd1, 0x17, 0x00, 0x11, 0x11, 0x64, 0x41, 0x25, 0x15,
	0xb1, 0x6d, 0x36, 0x0a, 0x47, 0x95, 0xd6, 0x53, 0x2f, 0x1a, 0x7b, 0x2b, 0x02, 0xde, 0xf0, 0x96,
	0xd1, 0x61, 0x52, 0x2c, 0xf1, 0x4a, 0x08, 0x3a, 0x84, 0xdd, 0x1b, 0x2a, 0x82, 0xd9, 0xd2, 0x2e,
	0x34, 0x8c, 0xa3, 0x3d, 0x9c, 0x21, 0xe7, 0x15, 0xec, 0xaf, 0x85, 0xa1, 0x3a, 0x14, 0xae, 0xe9,
	0x32, 0x2b, 0x41, 0x7d, 0xa2, 0x03, 0x28, 0xde, 0x90, 0x30, 0xa1, 0xb6, 0xa9, 0xd7, 0x52, 0xf0,
	0xd2, 0xfc, 0xdc, 0x70, 0x5f, 0xc0, 0xfe, 0x19, 0x67, 0x92, 0x32, 0x89, 0x69, 0x1c, 0x71, 0x16,
	0x53, 0x64, 0x43, 0x69, 0x92, 0x2e, 0x65, 0x12, 0x39, 0x74, 0xff, 0xb5, 0xa0, 0x36, 0x92, 0x44,
	0x26, 0xf1, 0x2d, 0x19, 0x81, 0x19, 0x4c, 0x53, 0xde, 0xa9, 0x69, 0x1b, 0xd8, 0x0c, 0xa6, 0xc8,
	0x83, 0x62, 0x2c, 0x89, 0x9f, 0x66, 0xab, 0xb5, 0x6c, 0x75, 0xcc, 0xbb, 0x61, 0x0a, 0xfa, 0x14,
	0xa7, 0x34, 0x74, 0x04, 0x05, 0x91, 0x30, 0x7d, 0xae, 0x5a, 0xeb, 0x70, 0x03, 0x1b, 0x27, 0x0c,
	0x2b, 0x0a, 0xfa, 0x14, 0x4a, 0x53, 0x2a, 0x49, 0x10, 0xc6, 0xb6, 0xd5, 0x30, 0x8e, 0x2a, 0x2d,
	0x67, 0x03, 0xfb, 0x3c, 0x65, 0xe0, 0x9c, 0x8a, 0x5e, 0x80, 0xb5, 0xa0, 0x92, 0xd8, 0x45, 0x1d,
	0xf2, 0x68, 0x43, 0xc8, 0x6b, 0x2a, 0x09, 0xd6, 0x24, 0xe7, 0x57, 0x13, 0x4a, 0x99, 0x82, 0x6a,
	0xe8, 0x82, 0xc6, 0x31, 0xf1, 0x69, 0x6c, 0x1b, 0x8d, 0x82, 0x6a, 0x68, 0x8e, 0x51, 0x1b, 0x4a,
	0x93, 0x39, 0x61, 0x6a, 0x2b, 0xed, 0xe6, 0xf3, 0xed, 0xa5, 0x78, 0x67, 0x29, 0x33, 0xed, 0x6a,
	0x1e, 0x87, 0x3e, 0x00, 0x98, 0x93, 0x38, 0xdb, 0xcb, 0xda, 0xba, 0xb2, 0xa2, 0xba, 0x46, 0x85,
	0xe0, 0x42, 0x9f, 0xb5, 0x8c, 0x53, 0xa0, 0xda, 0xf3, 0x23, 0x11, 0x2c, 0x60, 0xbe, 0x3e, 0x50,
	0x19, 0xe7, 0xd0, 0xe9, 0x43, 0x75, 0x35, 0xd1, 0x86, 0x39, 0x78, 0xb6, 0x3a, 0x07, 0x95, 0x56,
	0x5d, 0x95, 0x7c, 0x1e, 0xcc, 0x66, 0x79, 0xc1, 0x2b, 0x93, 0xe1, 0x1c, 0x82, 0xa5, 0x6c, 0x41,
	0xb5, 0xff, 0x3b, 0xac, 0xba, 0xeb, 0x9e, 0x40, 0x51, 0x77, 0x0f, 0x3d, 0x84, 0xf7, 0xde, 0x0c,
	0x46, 0xc3, 0xce, 0x59, 0xef, 0xa2, 0xd7, 0x39, 0xff, 0x7e, 0x74, 0xd5, 0xbe, 0xec, 0xd4, 0x77,
	0xd0, 0x1e, 0x58, 0xc3, 0x7e, 0x7b, 0x50, 0x37, 0x50, 0x19, 0x8a, 0xed, 0xe1, 0xb0, 0xff, 0x55,
	0xdd, 0x74, 0x3f, 0x83, 0x02, 0x4e, 0x18, 0x7a, 0x00, 0xfb, 0xab, 0x21, 0xf8, 0xcd, 0xa0, 0xbe,
	0x83, 0x2a, 0x50, 0x1a, 0x5d, 0xb5, 0xf1, 0x55, 0xe7, 0xbc, 0x6e, 0xa0, 0x2a, 0xec, 0x5d, 0xf4,
	0x06, 0xbd, 0x51, 0xb7, 0x73, 0x5e, 0x37, 0xdd, 0xef, 0xa0, 0xba, 0x5a, 0x9e, 0x6a, 0x08, 0x17,
	0x81, 0x1f, 0x30, 0x12, 0xe6, 0x7f, 0x58, 0x8e, 0xf5, 0xd8, 0x26, 0x42, 0xa8, 0xb1, 0x35, 0xb3,
	0xb1, 0x4d, 0xa1, 0xde, 0xb9, 0x63, 0x72, 0x0e, 0xdd, 0x3f, 0x4c, 0xa8, 0x5d, 0x0a, 0x12, 0xcd,
	0xcf, 0xf8, 0x22, 0xe2, 0x4c, 0x91, 0x4f, 0xf4, 0x7f, 0x26, 0xe9, 0x5b, 0x9d, 0xa0, 0xd2, 0x7a,
	0xac, 0x3c, 0xba, 0xcb, 0xf1, 0xbe, 0xd4, 0x84, 0xee, 0x0e, 0xce, 0xa8, 0xe8, 0x23, 0xb0, 0xe8,
	0xd4, 0xcf, 0x6d, 0x7d, 0xb4, 0x21, 0xa4, 0x33, 0xf5, 0x69, 0x77, 0x07, 0x6b, 0x9a, 0x73, 0x01,
	0xbb, 0xa9, 0xc4, 0xba, 0xb9, 0x08, 0x81, 0x75, 0x1d, 0xb0, 0x69, 0x76, 0x02, 0xfd, 0xad, 0xca,
	0xcf, 0x87, 0x5e, 0x95, 0x5f, 0xbd, 0x1d, 0x6c, 0x07, 0x83, 0xa5, 0x74, 0xd5, 0xdd, 0x10, 0xf3,
	0x44, 0x4c, 0x68, 0xa6, 0x94, 0x21, 0xa5, 0x36, 0xa5, 0x71, 0xee, 0x87, 0xfe, 0x56, 0x43, 0x47,
	0xa4, 0x14, 0xc1, 0x38, 0x91, 0xda, 0x0f, 0x35, 0xd5, 0x2b, 0x2b, 0xa7, 0x15, 0x28, 0x4f, 0xf2,
	0xaa, 0x5b, 0xbf, 0x99, 0xb0, 0xd7, 0x79, 0x4b, 0x27, 0x89, 0xe4, 0x02, 0x7d, 0x0b, 0x95, 0x2e,
	0x25, 0xa1, 0x9c, 0x9f, 0xcd, 0xe9, 0xe4, 0x1a, 0xed, 0xaf, 0xdd, 0x5e, 0x0e, 0xba, 0xff, 0x03,
	0xb8, 0xcf, 0x7e, 0xf9, 0xfb, 0x9f, 0xdf, 0xcd, 0x86, 0xfb, 0x44, 0xdf, 0xaf, 0x37, 0x9f, 0x34,
	0x17, 0x64, 0x32, 0x0f, 0x18, 0x6d, 0xce, 0xb5, 0xd2, 0x44, 0x29, 0xbd, 0x34, 0x8e, 0x3f, 0x36,
	0xd0, 0x00, 0xac, 0x61, 0x48, 0xd8, 0xbb, 0xc9, 0x3e, 0xd5, 0xb2, 0x8f, 0xdd, 0x83, 0x75, 0xd9,
	0x28, 0x24, 0x2c, 0xd5, 0x1b, 0x42, 0xb1, 0x1d, 0x45, 0xe1, 0xf2, 0xdd, 0x04, 0x1b, 0x5a, 0xd0,
	0x71, 0x1f, 0xae, 0x0b, 0x12, 0xa5, 0xa1, 0x15, 0x5b, 0x7f, 0x19, 0x50, 0xc5, 0x34, 0xb5, 0xb6,
	0xcb, 0x63, 0x89, 0xbe, 0x86, 0xf2, 0x25, 0x95, 0xa7, 0x01, 0x23, 0x62, 0x89, 0x0e, 0xbd, 0xf4,
	0xa9, 0xf0, 0xf2, 0xa7, 0xc2, 0xeb, 0xa8, 0xa7, 0xc2, 0x79, 0xa0, 0xb2, 0xad, 0x5d, 0xb1, 0x79,
	0x3a, 0x64, 0xe7, 0xe9, 0x44, 0xa6, 0x1b, 0x37, 0xc7, 0xa9, 0xdc, 0x58, 0x6b, 0xbf, 0xe6, 0xd3,
	0x24, 0xa4, 0xf7, 0x8f, 0xb0, 0x51, 0xb4, 0xa9, 0x45, 0x3f, 0x44, 0xcf, 0xef, 0x8b, 0x2e, 0xb4,
	0x4e, 0xdc, 0xfc, 0x29, 0x7f, 0x8f, 0x5e, 0x1d, 0x1f, 0xff, 0xdc, 0xfa, 0x06, 0x4a, 0x7a, 0x4a,
	0xa9, 0x50, 0x6e, 0xe9, 0xcf, 0x2d, 0x6e, 0xdd, 0x1d, 0xe6, 0xed, 0x6e, 0xf9, 0x8a, 0x97, 0xba,
	0x75, 0x05, 0x56, 0x8f, 0xcd, 0x38, 0xea, 0x83, 0x35, 0x0c, 0x98, 0xbf, 0xd5, 0x9f, 0x2d, 0xeb,
	0xee, 0x81, 0xce, 0x51, 0x43, 0xd5, 0x3c, 0x47, 0x14, 0x30, 0x7f, 0xbc, 0xab, 0x59, 0x27, 0xff,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x23, 0xe4, 0xb4, 0xc6, 0x07, 0x00, 0x00,
}
