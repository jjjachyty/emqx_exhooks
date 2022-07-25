package data

import (
	"context"

	"emqx_exhooks/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type msgRepo struct {
	data *Data
	log  *log.Helper
}

// NewMqttMsgRepo .
func NewMqttMsgRepo(data *Data, logger log.Logger) biz.MsgRepo {
	return &msgRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *msgRepo) Save(ctx context.Context, g *biz.MqttMsg) (*biz.MqttMsg, error) {
	err := r.data.db.Model(g).Create(g).Error
	return g, err
}

func (r *msgRepo) Update(ctx context.Context, g *biz.MqttMsg) (*biz.MqttMsg, error) {
	return g, nil
}

func (r *msgRepo) FindByID(context.Context, int64) (*biz.MqttMsg, error) {
	return nil, nil
}

func (r *msgRepo) ListByHello(context.Context, string) ([]*biz.MqttMsg, error) {
	return nil, nil
}

func (r *msgRepo) ListAll(context.Context) ([]*biz.MqttMsg, error) {
	return nil, nil
}
