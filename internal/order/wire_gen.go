// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package order

import (
	"sync"

	"github.com/ecodeclub/ecache"
	"github.com/ecodeclub/webook/internal/credit"
	"github.com/ecodeclub/webook/internal/order/internal/repository"
	"github.com/ecodeclub/webook/internal/order/internal/repository/dao"
	service4 "github.com/ecodeclub/webook/internal/order/internal/service"
	"github.com/ecodeclub/webook/internal/order/internal/web"
	"github.com/ecodeclub/webook/internal/payment"
	"github.com/ecodeclub/webook/internal/pkg/sequencenumber"
	"github.com/ecodeclub/webook/internal/product"
	"github.com/ego-component/egorm"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitHandler(db *gorm.DB, paymentSvc payment.Service, productSvc product.Service, creditSvc credit.Service, cache ecache.Cache) *web.Handler {
	orderDAO := InitTablesOnce(db)
	orderRepository := repository.NewRepository(orderDAO)
	serviceService := service4.NewService(orderRepository)
	generator := sequencenumber.NewGenerator()
	handler := web.NewHandler(serviceService, paymentSvc, productSvc, creditSvc, generator, cache)
	return handler
}

// wire.go:

var HandlerSet = wire.NewSet(
	InitTablesOnce, repository.NewRepository, service4.NewService, sequencenumber.NewGenerator, web.NewHandler,
)

var once = &sync.Once{}

func InitTablesOnce(db *egorm.Component) dao.OrderDAO {
	once.Do(func() {
		_ = dao.InitTables(db)
	})
	return dao.NewOrderGORMDAO(db)
}

type Handler = web.Handler
