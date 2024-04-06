// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package member

import (
	"context"
	"sync"
	"time"

	"github.com/ecodeclub/mq-api"
	"github.com/ecodeclub/webook/internal/member/internal/domain"
	"github.com/ecodeclub/webook/internal/member/internal/event"
	"github.com/ecodeclub/webook/internal/member/internal/repository"
	"github.com/ecodeclub/webook/internal/member/internal/repository/dao"
	"github.com/ecodeclub/webook/internal/member/internal/service"
	"github.com/ego-component/egorm"
	"github.com/gotomicro/ego/core/elog"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitModule(db *gorm.DB, q mq.MQ) (*Module, error) {
	service := InitService(db, q)
	v, err := initRegistrationConsumer(service, q)
	if err != nil {
		return nil, err
	}
	module := &Module{
		Svc:                        service,
		registrationEventConsumers: v,
	}
	return module, nil
}

// wire.go:

type Member = domain.Member

type Service = service.Service

var (
	once = &sync.Once{}
	svc  service.Service
)

func InitService(db *egorm.Component, q mq.MQ) Service {
	once.Do(func() {
		_ = dao.InitTables(db)
		d := dao.NewMemberGORMDAO(db)
		r := repository.NewMemberRepository(d)
		svc = service.NewMemberService(r)
	})
	return svc
}

func initRegistrationConsumer(svc2 service.Service, q mq.MQ) ([]*event.RegistrationEventConsumer, error) {
	startAtFunc := func() int64 {
		return time.Now().UTC().UnixMilli()
	}
	endAtFunc := func() int64 {
		return time.Date(2024, 6, 30, 23, 59, 59, 0, time.UTC).UnixMilli()
	}

	partitions := 3
	consumers := make([]*event.RegistrationEventConsumer, 0, partitions)
	for i := 0; i < partitions; i++ {
		topic := event.RegistrationEvent{}.Topic()
		groupID := topic
		c, err := q.Consumer(topic, groupID)
		if err != nil {
			return nil, err
		}
		consumer := event.NewRegistrationEventConsumer(svc2, c, startAtFunc, endAtFunc)
		consumers = append(consumers, consumer)
		go func() {
			for {
				er := consumer.Consume(context.Background())
				if er != nil {
					elog.DefaultLogger.
						Error("消费注册事件失败", elog.FieldErr(er))
				}
			}
		}()
	}
	return consumers, nil
}