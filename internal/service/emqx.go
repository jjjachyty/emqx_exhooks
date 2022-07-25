package service

import (
	"context"
	v1 "emqx_exhooks/api/emqx/v1/go"
	"emqx_exhooks/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// EmqxHookProviderService is a greeter service.
type EmqxHookProviderService struct {
	v1.UnimplementedHookProviderServer

	uc  *biz.MsgUsecase
	log *log.Helper
}

// NewEmqxHookProviderService new a greeter service.
func NewEmqxHookProviderService(uc *biz.MsgUsecase, logger log.Logger) *EmqxHookProviderService {
	return &EmqxHookProviderService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer.
func (s *EmqxHookProviderService) OnMessagePublish(ctx context.Context, in *v1.MessagePublishRequest) (*v1.ValuedResponse, error) {
	s.log.Infof("OnMessagePublish,in:%v\n", in)
	_, err := s.uc.CreateMsg(ctx, &biz.MqttMsg{Msgid: in.Message.Id, Topic: in.Message.Topic, Sender: in.Message.From, Node: in.Message.Node, Qos: int64(in.Message.Qos), Payload: string(in.Message.Payload)})
	if err != nil {
		s.log.Errorf("OnMessagePublish,err:%v", err)
		return &v1.ValuedResponse{Type: v1.ValuedResponse_CONTINUE, Value: &v1.ValuedResponse_Message{Message: in.Message}}, err
	}
	return &v1.ValuedResponse{Type: v1.ValuedResponse_CONTINUE, Value: &v1.ValuedResponse_Message{Message: in.Message}}, nil
}
