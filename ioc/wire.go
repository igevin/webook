// Copyright 2023 ecodeclub
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build wireinject

package ioc

import (
	"github.com/ecodeclub/webook/internal/cases"
	"github.com/ecodeclub/webook/internal/cos"
	"github.com/ecodeclub/webook/internal/credit"
	"github.com/ecodeclub/webook/internal/feedback"
	"github.com/ecodeclub/webook/internal/interactive"
	"github.com/ecodeclub/webook/internal/label"
	"github.com/ecodeclub/webook/internal/marketing"
	"github.com/ecodeclub/webook/internal/member"
	"github.com/ecodeclub/webook/internal/order"
	"github.com/ecodeclub/webook/internal/payment"
	"github.com/ecodeclub/webook/internal/permission"
	"github.com/ecodeclub/webook/internal/pkg/middleware"
	"github.com/ecodeclub/webook/internal/product"
	"github.com/ecodeclub/webook/internal/project"
	baguwen "github.com/ecodeclub/webook/internal/question"
	"github.com/ecodeclub/webook/internal/recon"
	"github.com/ecodeclub/webook/internal/skill"
	"github.com/google/wire"
)

var BaseSet = wire.NewSet(InitDB, InitCache, InitRedis, InitMQ, InitCosConfig)

func InitApp() (*App, error) {
	wire.Build(wire.Struct(new(App), "*"),
		BaseSet,
		InitSession,
		cos.InitHandler,
		baguwen.InitModule,
		wire.FieldsOf(new(*baguwen.Module), "Hdl", "QsHdl"),
		InitUserHandler,
		label.InitHandler,
		cases.InitModule,
		wire.FieldsOf(new(*cases.Module), "Hdl"),
		skill.InitHandler,
		feedback.InitHandler,
		member.InitModule,
		wire.FieldsOf(new(*member.Module), "Svc"),
		middleware.NewCheckMembershipMiddlewareBuilder,
		product.InitModule,
		wire.FieldsOf(new(*product.Module), "Hdl"),
		order.InitModule,
		wire.FieldsOf(new(*order.Module), "Hdl", "CloseTimeoutOrdersJob"),
		payment.InitModule,
		wire.FieldsOf(new(*payment.Module), "Hdl", "SyncWechatOrderJob"),
		credit.InitModule,
		wire.FieldsOf(new(*credit.Module), "Hdl", "CloseTimeoutLockedCreditsJob"),
		project.InitModule,
		wire.FieldsOf(new(*project.Module), "AdminHdl", "Hdl"),
		recon.InitModule,
		wire.FieldsOf(new(*recon.Module), "SyncPaymentAndOrderJob"),
		marketing.InitModule,
		wire.FieldsOf(new(*marketing.Module), "AdminHdl", "Hdl"),
		interactive.InitModule,
		wire.FieldsOf(new(*interactive.Module), "Hdl"),
		permission.InitModule,
		wire.FieldsOf(new(*permission.Module), "Svc"),
		middleware.NewCheckPermissionMiddlewareBuilder,
		initLocalActiveLimiterBuilder,
		initCronJobs,
		// 这两个顺序不要换
		initGinxServer,
		InitAdminServer,
	)
	return new(App), nil
}
