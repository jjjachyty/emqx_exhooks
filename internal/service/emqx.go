package service

import (
	"context"
	v1 "emqx_exhooks/api/emqx/v1/go"
	"emqx_exhooks/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// EmqxHookProviderService is a greeter service.
type EmqxHookProviderService struct {
	uc  *biz.MsgUsecase
	log *log.Helper
	v1.UnimplementedHookProviderServer
}

// var _ EmqxHookProviderService = (v1.UnimplementedHookProviderServer)(&EmqxHookProviderService{})
var _ = (v1.HookProviderServer)(&EmqxHookProviderService{})

// NewEmqxHookProviderService new a greeter service.
func NewEmqxHookProviderService(uc *biz.MsgUsecase, logger log.Logger) *EmqxHookProviderService {
	return &EmqxHookProviderService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer.
func (s *EmqxHookProviderService) OnMessagePublish(ctx context.Context, in *v1.MessagePublishRequest) (*v1.ValuedResponse, error) {
	s.log.Infof("OnMessagePublish,in:%v\n", in)
	_, err := s.uc.CreateMsg(ctx, &biz.MqttMsg{Msgid: in.Message.Id, Topic: in.Message.Topic, Sender: in.Message.From, Node: in.Message.Node, Qos: int64(in.Message.Qos), Payload: string(in.Message.Payload), Timestamp: in.Message.Timestamp})
	if err != nil {
		s.log.Errorf("OnMessagePublish,err:%v", err)
		return &v1.ValuedResponse{Type: v1.ValuedResponse_CONTINUE, Value: &v1.ValuedResponse_Message{Message: in.Message}}, err
	}
	return &v1.ValuedResponse{Type: v1.ValuedResponse_CONTINUE, Value: &v1.ValuedResponse_Message{Message: in.Message}}, nil
}

func (s *EmqxHookProviderService) OnProviderLoaded(ctx context.Context, in *v1.ProviderLoadedRequest) (*v1.LoadedResponse, error) {
	s.log.Infof("OnProviderLoaded,in:%v\n", in)
	hooks := []*v1.HookSpec{
		{Name: "message.publish"},
	}
	return &v1.LoadedResponse{Hooks: hooks}, nil
}

func (s *EmqxHookProviderService) OnProviderUnloaded(ctx context.Context, in *v1.ProviderUnloadedRequest) (*v1.EmptySuccess, error) {
	s.log.Infof("OnProviderUnloaded,in:%v\n", in)
	return &v1.EmptySuccess{}, nil
}
func (s *EmqxHookProviderService) OnClientConnect(ctx context.Context, in *v1.ClientConnectRequest) (*v1.EmptySuccess, error) {
	return &v1.EmptySuccess{}, nil
}
func (s *EmqxHookProviderService) OnClientConnack(ctx context.Context, in *v1.ClientConnackRequest) (*v1.EmptySuccess, error) {
	return &v1.EmptySuccess{}, nil
}
func (s *EmqxHookProviderService) OnClientConnected(ctx context.Context, in *v1.ClientConnectedRequest) (*v1.EmptySuccess, error) {
	return &v1.EmptySuccess{}, nil
}
func (s *EmqxHookProviderService) OnClientDisconnected(ctx context.Context, in *v1.ClientDisconnectedRequest) (*v1.EmptySuccess, error) {
	return &v1.EmptySuccess{}, nil
}
func (s *EmqxHookProviderService) OnClientAuthenticate(ctx context.Context, in *v1.ClientAuthenticateRequest) (*v1.ValuedResponse, error) {
	reply := &v1.ValuedResponse{}
	reply.Type = v1.ValuedResponse_STOP_AND_RETURN
	reply.Value = &v1.ValuedResponse_BoolResult{BoolResult: true}
	return reply, nil
}
func (s *EmqxHookProviderService) OnClientAuthorize(ctx context.Context, in *v1.ClientAuthorizeRequest) (*v1.ValuedResponse, error) {
	reply := &v1.ValuedResponse{}
	reply.Type = v1.ValuedResponse_STOP_AND_RETURN
	reply.Value = &v1.ValuedResponse_BoolResult{BoolResult: true}
	return reply, nil
}
func (s *EmqxHookProviderService) OnClientSubscribe(ctx context.Context, in *v1.ClientSubscribeRequest) (*v1.EmptySuccess, error) {
	return &v1.EmptySuccess{}, nil
}
func (s *EmqxHookProviderService) OnClientUnsubscribe(ctx context.Context, in *v1.ClientUnsubscribeRequest) (*v1.EmptySuccess, error) {
	return &v1.EmptySuccess{}, nil
}
func (s *EmqxHookProviderService) OnSessionCreated(ctx context.Context, in *v1.SessionCreatedRequest) (*v1.EmptySuccess, error) {
	return &v1.EmptySuccess{}, nil
}
func (s *EmqxHookProviderService) OnSessionSubscribed(ctx context.Context, in *v1.SessionSubscribedRequest) (*v1.EmptySuccess, error) {
	return &v1.EmptySuccess{}, nil
}
func (s *EmqxHookProviderService) OnSessionUnsubscribed(ctx context.Context, in *v1.SessionUnsubscribedRequest) (*v1.EmptySuccess, error) {
	return &v1.EmptySuccess{}, nil
}
func (s *EmqxHookProviderService) OnSessionResumed(ctx context.Context, in *v1.SessionResumedRequest) (*v1.EmptySuccess, error) {
	return &v1.EmptySuccess{}, nil
}
func (s *EmqxHookProviderService) OnSessionDiscarded(ctx context.Context, in *v1.SessionDiscardedRequest) (*v1.EmptySuccess, error) {
	return &v1.EmptySuccess{}, nil
}
func (s *EmqxHookProviderService) OnSessionTakenover(ctx context.Context, in *v1.SessionTakenoverRequest) (*v1.EmptySuccess, error) {
	return &v1.EmptySuccess{}, nil
}
func (s *EmqxHookProviderService) OnSessionTerminated(ctx context.Context, in *v1.SessionTerminatedRequest) (*v1.EmptySuccess, error) {
	return &v1.EmptySuccess{}, nil
}
func (s *EmqxHookProviderService) OnMessageDelivered(ctx context.Context, in *v1.MessageDeliveredRequest) (*v1.EmptySuccess, error) {
	return &v1.EmptySuccess{}, nil
}
func (s *EmqxHookProviderService) OnMessageDropped(ctx context.Context, in *v1.MessageDroppedRequest) (*v1.EmptySuccess, error) {
	return &v1.EmptySuccess{}, nil
}
func (s *EmqxHookProviderService) OnMessageAcked(ctx context.Context, in *v1.MessageAckedRequest) (*v1.EmptySuccess, error) {
	return &v1.EmptySuccess{}, nil
}
