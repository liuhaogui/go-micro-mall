// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package go_micro_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Delete(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*User, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Users, error)
	UpdateInfo(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	Ping(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Delete(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Get(ctx context.Context, in *User, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.name, "UserService.Get", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Users, error) {
	req := c.c.NewRequest(c.name, "UserService.GetAll", in)
	out := new(Users)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateInfo(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.UpdateInfo", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.name, "UserService.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Ping(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserService.Ping", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *Response) error
	Delete(context.Context, *User, *Response) error
	Get(context.Context, *User, *User) error
	GetAll(context.Context, *Request, *Users) error
	UpdateInfo(context.Context, *User, *Response) error
	Auth(context.Context, *User, *Token) error
	Ping(context.Context, *Request, *Response) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Create(ctx context.Context, in *User, out *Response) error
		Delete(ctx context.Context, in *User, out *Response) error
		Get(ctx context.Context, in *User, out *User) error
		GetAll(ctx context.Context, in *Request, out *Users) error
		UpdateInfo(ctx context.Context, in *User, out *Response) error
		Auth(ctx context.Context, in *User, out *Token) error
		Ping(ctx context.Context, in *Request, out *Response) error
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

func (h *userServiceHandler) Create(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *userServiceHandler) Delete(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Delete(ctx, in, out)
}

func (h *userServiceHandler) Get(ctx context.Context, in *User, out *User) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *userServiceHandler) GetAll(ctx context.Context, in *Request, out *Users) error {
	return h.UserServiceHandler.GetAll(ctx, in, out)
}

func (h *userServiceHandler) UpdateInfo(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.UpdateInfo(ctx, in, out)
}

func (h *userServiceHandler) Auth(ctx context.Context, in *User, out *Token) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *userServiceHandler) Ping(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.Ping(ctx, in, out)
}
