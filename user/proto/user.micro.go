// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user/proto/user.proto

package orbli_micro_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
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

// Client API for UserService service

type UserService interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*User, error)
	Read(ctx context.Context, in *User, opts ...client.CallOption) (*User, error)
	Update(ctx context.Context, in *User, opts ...client.CallOption) (*User, error)
	Delete(ctx context.Context, in *User, opts ...client.CallOption) (*User, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "orbli.micro.user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Create(ctx context.Context, in *User, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserService.Create", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Read(ctx context.Context, in *User, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserService.Read", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Update(ctx context.Context, in *User, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserService.Update", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Delete(ctx context.Context, in *User, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserService.Delete", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *User) error
	Read(context.Context, *User, *User) error
	Update(context.Context, *User, *User) error
	Delete(context.Context, *User, *User) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Create(ctx context.Context, in *User, out *User) error
		Read(ctx context.Context, in *User, out *User) error
		Update(ctx context.Context, in *User, out *User) error
		Delete(ctx context.Context, in *User, out *User) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) Create(ctx context.Context, in *User, out *User) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *userServiceHandler) Read(ctx context.Context, in *User, out *User) error {
	return h.UserServiceHandler.Read(ctx, in, out)
}

func (h *userServiceHandler) Update(ctx context.Context, in *User, out *User) error {
	return h.UserServiceHandler.Update(ctx, in, out)
}

func (h *userServiceHandler) Delete(ctx context.Context, in *User, out *User) error {
	return h.UserServiceHandler.Delete(ctx, in, out)
}
