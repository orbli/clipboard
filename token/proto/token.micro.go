// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: token/proto/token.proto

package orbli_micro_token

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for TokenService service

type TokenService interface {
	Create(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
	Read(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
	Update(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
}

type tokenService struct {
	c    client.Client
	name string
}

func NewTokenService(name string, c client.Client) TokenService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "orbli.micro.token"
	}
	return &tokenService{
		c:    c,
		name: name,
	}
}

func (c *tokenService) Create(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.name, "TokenService.Create", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenService) Read(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.name, "TokenService.Read", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenService) Update(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.name, "TokenService.Update", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TokenService service

type TokenServiceHandler interface {
	Create(context.Context, *Token, *Token) error
	Read(context.Context, *Token, *Token) error
	Update(context.Context, *Token, *Token) error
}

func RegisterTokenServiceHandler(s server.Server, hdlr TokenServiceHandler, opts ...server.HandlerOption) error {
	type tokenService interface {
		Create(ctx context.Context, in *Token, out *Token) error
		Read(ctx context.Context, in *Token, out *Token) error
		Update(ctx context.Context, in *Token, out *Token) error
	}
	type TokenService struct {
		tokenService
	}
	h := &tokenServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TokenService{h}, opts...))
}

type tokenServiceHandler struct {
	TokenServiceHandler
}

func (h *tokenServiceHandler) Create(ctx context.Context, in *Token, out *Token) error {
	return h.TokenServiceHandler.Create(ctx, in, out)
}

func (h *tokenServiceHandler) Read(ctx context.Context, in *Token, out *Token) error {
	return h.TokenServiceHandler.Read(ctx, in, out)
}

func (h *tokenServiceHandler) Update(ctx context.Context, in *Token, out *Token) error {
	return h.TokenServiceHandler.Update(ctx, in, out)
}
