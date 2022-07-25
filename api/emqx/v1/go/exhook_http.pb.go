// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type HookProviderHTTPServer interface {
	OnMessagePublish(context.Context, *MessagePublishRequest) (*ValuedResponse, error)
}

func RegisterHookProviderHTTPServer(s *http.Server, srv HookProviderHTTPServer) {
	r := s.Route("/")
	r.POST("/emqx/hooks/message/publish", _HookProvider_OnMessagePublish0_HTTP_Handler(srv))
}

func _HookProvider_OnMessagePublish0_HTTP_Handler(srv HookProviderHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MessagePublishRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/emqx.exhook.v2.HookProvider/OnMessagePublish")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OnMessagePublish(ctx, req.(*MessagePublishRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ValuedResponse)
		return ctx.Result(200, reply)
	}
}

type HookProviderHTTPClient interface {
	OnMessagePublish(ctx context.Context, req *MessagePublishRequest, opts ...http.CallOption) (rsp *ValuedResponse, err error)
}

type HookProviderHTTPClientImpl struct {
	cc *http.Client
}

func NewHookProviderHTTPClient(client *http.Client) HookProviderHTTPClient {
	return &HookProviderHTTPClientImpl{client}
}

func (c *HookProviderHTTPClientImpl) OnMessagePublish(ctx context.Context, in *MessagePublishRequest, opts ...http.CallOption) (*ValuedResponse, error) {
	var out ValuedResponse
	pattern := "/emqx/hooks/message/publish"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/emqx.exhook.v2.HookProvider/OnMessagePublish"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
