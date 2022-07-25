package data

import (
	"emqx_exhooks/internal/conf"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMqttMsgRepo)

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Helper
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (data *Data, fn func(), err error) {
	data = &Data{}
	data.log = log.NewHelper(logger)
	cleanup := func() {
		data.log.Info("closing the data resources")
	}
	data.db, err = newGormDb(c.Database, data.log)
	return data, cleanup, err
}

func newGormDb(cfgs []*conf.Data_Database, log *log.Helper) (*gorm.DB, error) {
	if len(cfgs) == 0 {
		log.Error("cfg Master is not valid", cfgs)
		return nil, errors.New("cfg  Master is not valid")
	}

	cf := cfgs[0]
	db, err := gorm.Open(mysql.Open(cf.Source), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatalf("数据库连接出错，err=%v", err)
	}

	replicas := make([]gorm.Dialector, 0)
	for _, v := range cfgs[1:] {
		replicas = append(replicas, mysql.Open(v.Source))
	}

	dbResolverCfg := dbresolver.Config{
		Sources:  replicas,
		Replicas: replicas,
		Policy:   dbresolver.RandomPolicy{}}
	readWritePlugin := dbresolver.Register(dbResolverCfg)
	db.Use(readWritePlugin)
	log.Info("初始化DB成功")
	return db, nil
}
