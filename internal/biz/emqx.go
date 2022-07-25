package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// MqttMsg  MQTT 消息
type MqttMsg struct {
	ID      int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	Msgid   string    `gorm:"column:msgid" db:"msgid" json:"msgid" form:"msgid"`
	Topic   string    `gorm:"column:topic" db:"topic" json:"topic" form:"topic"`
	Sender  string    `gorm:"column:sender" db:"sender" json:"sender" form:"sender"`
	Node    string    `gorm:"column:node" db:"node" json:"node" form:"node"`
	Qos     int64     `gorm:"column:qos" db:"qos" json:"qos" form:"qos"`
	Retain  int64     `gorm:"column:retain" db:"retain" json:"retain" form:"retain"`
	Payload string    `gorm:"column:payload" db:"payload" json:"payload" form:"payload"`
	Arrived time.Time `gorm:"column:arrived" db:"arrived" json:"arrived" form:"arrived"`
}

func (MqttMsg) TabName() string {
	return "mqtt_msg"
}

// MsgRepo is a Greater repo.
type MsgRepo interface {
	Save(context.Context, *MqttMsg) (*MqttMsg, error)
	Update(context.Context, *MqttMsg) (*MqttMsg, error)
	FindByID(context.Context, int64) (*MqttMsg, error)
	ListByHello(context.Context, string) ([]*MqttMsg, error)
	ListAll(context.Context) ([]*MqttMsg, error)
}

// MsgUsecase is a Msg usecase.
type MsgUsecase struct {
	repo MsgRepo
	log  *log.Helper
}

// NewMsgUsecase new a Msg usecase.
func NewMsgUsecase(repo MsgRepo, logger log.Logger) *MsgUsecase {
	return &MsgUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateMsg creates a Msg, and returns the new Msg.
func (uc *MsgUsecase) CreateMsg(ctx context.Context, g *MqttMsg) (*MqttMsg, error) {
	uc.log.WithContext(ctx).Infof("CreateMsg: %v", g)
	return uc.repo.Save(ctx, g)
}
