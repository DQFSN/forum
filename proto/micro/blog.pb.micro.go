// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: blog.proto

package blog

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Auth service

func NewAuthEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Auth service

type AuthService interface {
	LogIn(ctx context.Context, in *LogInRequest, opts ...client.CallOption) (*LogInReply, error)
	SignUp(ctx context.Context, in *SignUpRequest, opts ...client.CallOption) (*SignUpReply, error)
	ModifyUser(ctx context.Context, in *ModifyUserRequest, opts ...client.CallOption) (*ModifyUserReply, error)
	DelUser(ctx context.Context, in *DelUserRequest, opts ...client.CallOption) (*DelUserReply, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) LogIn(ctx context.Context, in *LogInRequest, opts ...client.CallOption) (*LogInReply, error) {
	req := c.c.NewRequest(c.name, "Auth.LogIn", in)
	out := new(LogInReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) SignUp(ctx context.Context, in *SignUpRequest, opts ...client.CallOption) (*SignUpReply, error) {
	req := c.c.NewRequest(c.name, "Auth.SignUp", in)
	out := new(SignUpReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ModifyUser(ctx context.Context, in *ModifyUserRequest, opts ...client.CallOption) (*ModifyUserReply, error) {
	req := c.c.NewRequest(c.name, "Auth.ModifyUser", in)
	out := new(ModifyUserReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) DelUser(ctx context.Context, in *DelUserRequest, opts ...client.CallOption) (*DelUserReply, error) {
	req := c.c.NewRequest(c.name, "Auth.DelUser", in)
	out := new(DelUserReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	LogIn(context.Context, *LogInRequest, *LogInReply) error
	SignUp(context.Context, *SignUpRequest, *SignUpReply) error
	ModifyUser(context.Context, *ModifyUserRequest, *ModifyUserReply) error
	DelUser(context.Context, *DelUserRequest, *DelUserReply) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		LogIn(ctx context.Context, in *LogInRequest, out *LogInReply) error
		SignUp(ctx context.Context, in *SignUpRequest, out *SignUpReply) error
		ModifyUser(ctx context.Context, in *ModifyUserRequest, out *ModifyUserReply) error
		DelUser(ctx context.Context, in *DelUserRequest, out *DelUserReply) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) LogIn(ctx context.Context, in *LogInRequest, out *LogInReply) error {
	return h.AuthHandler.LogIn(ctx, in, out)
}

func (h *authHandler) SignUp(ctx context.Context, in *SignUpRequest, out *SignUpReply) error {
	return h.AuthHandler.SignUp(ctx, in, out)
}

func (h *authHandler) ModifyUser(ctx context.Context, in *ModifyUserRequest, out *ModifyUserReply) error {
	return h.AuthHandler.ModifyUser(ctx, in, out)
}

func (h *authHandler) DelUser(ctx context.Context, in *DelUserRequest, out *DelUserReply) error {
	return h.AuthHandler.DelUser(ctx, in, out)
}

// Api Endpoints for Publish service

func NewPublishEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Publish service

type PublishService interface {
	PublishBlog(ctx context.Context, in *PublishRequest, opts ...client.CallOption) (*PublishReply, error)
	GetBlogs(ctx context.Context, in *BlogsRequest, opts ...client.CallOption) (*BlogsReply, error)
	ModifyBlog(ctx context.Context, in *ModifyBlogRequest, opts ...client.CallOption) (*ModifyBlogReply, error)
	DelBlog(ctx context.Context, in *DelBlogRequest, opts ...client.CallOption) (*DelBlogReply, error)
}

type publishService struct {
	c    client.Client
	name string
}

func NewPublishService(name string, c client.Client) PublishService {
	return &publishService{
		c:    c,
		name: name,
	}
}

func (c *publishService) PublishBlog(ctx context.Context, in *PublishRequest, opts ...client.CallOption) (*PublishReply, error) {
	req := c.c.NewRequest(c.name, "Publish.PublishBlog", in)
	out := new(PublishReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishService) GetBlogs(ctx context.Context, in *BlogsRequest, opts ...client.CallOption) (*BlogsReply, error) {
	req := c.c.NewRequest(c.name, "Publish.GetBlogs", in)
	out := new(BlogsReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishService) ModifyBlog(ctx context.Context, in *ModifyBlogRequest, opts ...client.CallOption) (*ModifyBlogReply, error) {
	req := c.c.NewRequest(c.name, "Publish.ModifyBlog", in)
	out := new(ModifyBlogReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishService) DelBlog(ctx context.Context, in *DelBlogRequest, opts ...client.CallOption) (*DelBlogReply, error) {
	req := c.c.NewRequest(c.name, "Publish.DelBlog", in)
	out := new(DelBlogReply)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Publish service

type PublishHandler interface {
	PublishBlog(context.Context, *PublishRequest, *PublishReply) error
	GetBlogs(context.Context, *BlogsRequest, *BlogsReply) error
	ModifyBlog(context.Context, *ModifyBlogRequest, *ModifyBlogReply) error
	DelBlog(context.Context, *DelBlogRequest, *DelBlogReply) error
}

func RegisterPublishHandler(s server.Server, hdlr PublishHandler, opts ...server.HandlerOption) error {
	type publish interface {
		PublishBlog(ctx context.Context, in *PublishRequest, out *PublishReply) error
		GetBlogs(ctx context.Context, in *BlogsRequest, out *BlogsReply) error
		ModifyBlog(ctx context.Context, in *ModifyBlogRequest, out *ModifyBlogReply) error
		DelBlog(ctx context.Context, in *DelBlogRequest, out *DelBlogReply) error
	}
	type Publish struct {
		publish
	}
	h := &publishHandler{hdlr}
	return s.Handle(s.NewHandler(&Publish{h}, opts...))
}

type publishHandler struct {
	PublishHandler
}

func (h *publishHandler) PublishBlog(ctx context.Context, in *PublishRequest, out *PublishReply) error {
	return h.PublishHandler.PublishBlog(ctx, in, out)
}

func (h *publishHandler) GetBlogs(ctx context.Context, in *BlogsRequest, out *BlogsReply) error {
	return h.PublishHandler.GetBlogs(ctx, in, out)
}

func (h *publishHandler) ModifyBlog(ctx context.Context, in *ModifyBlogRequest, out *ModifyBlogReply) error {
	return h.PublishHandler.ModifyBlog(ctx, in, out)
}

func (h *publishHandler) DelBlog(ctx context.Context, in *DelBlogRequest, out *DelBlogReply) error {
	return h.PublishHandler.DelBlog(ctx, in, out)
}
