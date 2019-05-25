//Generated by gRPC Go plugin
//If you make any local changes, they will be lost
//source: bookmarks

package bookmarks

import "github.com/google/flatbuffers/go"

import (
  context "context"
  grpc "google.golang.org/grpc"
)

// Client API for BookmarksService service
type BookmarksServiceClient interface{
  Add(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* AddResponse, error)  
  LastAdded(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* LastAddedResponse, error)  
  All(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (BookmarksService_AllClient, error)  
  GetAll(ctx context.Context, in *flatbuffers.Builder, 
  	opts... grpc.CallOption) (* AllResponse, error)  
}

type bookmarksServiceClient struct {
  cc *grpc.ClientConn
}

func NewBookmarksServiceClient(cc *grpc.ClientConn) BookmarksServiceClient {
  return &bookmarksServiceClient{cc}
}

func (c *bookmarksServiceClient) Add(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* AddResponse, error) {
  out := new(AddResponse)
  err := grpc.Invoke(ctx, "/bookmarks.BookmarksService/Add", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

func (c *bookmarksServiceClient) LastAdded(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* LastAddedResponse, error) {
  out := new(LastAddedResponse)
  err := grpc.Invoke(ctx, "/bookmarks.BookmarksService/LastAdded", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

func (c *bookmarksServiceClient) All(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (BookmarksService_AllClient, error) {
  stream, err := grpc.NewClientStream(ctx, &_BookmarksService_serviceDesc.Streams[0], c.cc, "/bookmarks.BookmarksService/All", opts...)
  if err != nil { return nil, err }
  x := &bookmarksServiceAllClient{stream}
  if err := x.ClientStream.SendMsg(in); err != nil { return nil, err }
  if err := x.ClientStream.CloseSend(); err != nil { return nil, err }
  return x,nil
}

type BookmarksService_AllClient interface {
  Recv() (*LastAddedResponse, error)
  grpc.ClientStream
}

type bookmarksServiceAllClient struct{
  grpc.ClientStream
}

func (x *bookmarksServiceAllClient) Recv() (*LastAddedResponse, error) {
  m := new(LastAddedResponse)
  if err := x.ClientStream.RecvMsg(m); err != nil { return nil, err }
  return m, nil
}

func (c *bookmarksServiceClient) GetAll(ctx context.Context, in *flatbuffers.Builder, 
	opts... grpc.CallOption) (* AllResponse, error) {
  out := new(AllResponse)
  err := grpc.Invoke(ctx, "/bookmarks.BookmarksService/GetAll", in, out, c.cc, opts...)
  if err != nil { return nil, err }
  return out, nil
}

// Server API for BookmarksService service
type BookmarksServiceServer interface {
  Add(context.Context, *AddRequest) (*flatbuffers.Builder, error)  
  LastAdded(context.Context, *LastAddedRequest) (*flatbuffers.Builder, error)  
  All(*LastAddedRequest, BookmarksService_AllServer) error  
  GetAll(context.Context, *AllRequest) (*flatbuffers.Builder, error)  
}

func RegisterBookmarksServiceServer(s *grpc.Server, srv BookmarksServiceServer) {
  s.RegisterService(&_BookmarksService_serviceDesc, srv)
}

func _BookmarksService_Add_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(AddRequest)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(BookmarksServiceServer).Add(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/bookmarks.BookmarksService/Add",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(BookmarksServiceServer).Add(ctx, req.(* AddRequest))
  }
  return interceptor(ctx, in, info, handler)
}


func _BookmarksService_LastAdded_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(LastAddedRequest)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(BookmarksServiceServer).LastAdded(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/bookmarks.BookmarksService/LastAdded",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(BookmarksServiceServer).LastAdded(ctx, req.(* LastAddedRequest))
  }
  return interceptor(ctx, in, info, handler)
}


func _BookmarksService_All_Handler(srv interface{}, stream grpc.ServerStream) error {
  m := new(LastAddedRequest)
  if err := stream.RecvMsg(m); err != nil { return err }
  return srv.(BookmarksServiceServer).All(m, &bookmarksServiceAllServer{stream})
}

type BookmarksService_AllServer interface { 
  Send(* flatbuffers.Builder) error
  grpc.ServerStream
}

type bookmarksServiceAllServer struct {
  grpc.ServerStream
}

func (x *bookmarksServiceAllServer) Send(m *flatbuffers.Builder) error {
  return x.ServerStream.SendMsg(m)
}


func _BookmarksService_GetAll_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
  in := new(AllRequest)
  if err := dec(in); err != nil { return nil, err }
  if interceptor == nil { return srv.(BookmarksServiceServer).GetAll(ctx, in) }
  info := &grpc.UnaryServerInfo{
    Server: srv,
    FullMethod: "/bookmarks.BookmarksService/GetAll",
  }
  
  handler := func(ctx context.Context, req interface{}) (interface{}, error) {
    return srv.(BookmarksServiceServer).GetAll(ctx, req.(* AllRequest))
  }
  return interceptor(ctx, in, info, handler)
}


var _BookmarksService_serviceDesc = grpc.ServiceDesc{
  ServiceName: "bookmarks.BookmarksService",
  HandlerType: (*BookmarksServiceServer)(nil),
  Methods: []grpc.MethodDesc{
    {
      MethodName: "Add",
      Handler: _BookmarksService_Add_Handler, 
    },
    {
      MethodName: "LastAdded",
      Handler: _BookmarksService_LastAdded_Handler, 
    },
    {
      MethodName: "GetAll",
      Handler: _BookmarksService_GetAll_Handler, 
    },
  },
  Streams: []grpc.StreamDesc{
    {
      StreamName: "All",
      Handler: _BookmarksService_All_Handler, 
      ServerStreams: true,
    },
  },
}

