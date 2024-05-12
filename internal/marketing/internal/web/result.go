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

package web

import (
	"github.com/ecodeclub/ginx"
	"github.com/ecodeclub/webook/internal/marketing/internal/errs"
)

var (
	systemErrorResult = ginx.Result{
		Code: errs.SystemError.Code,
		Msg:  errs.SystemError.Msg,
	}
	redemptionCodeUsedErrResult = ginx.Result{
		Code: errs.RedemptionCodeUsedError.Code,
		Msg:  errs.RedemptionCodeUsedError.Msg,
	}
	redemptionCodeNotFoundErrResult = ginx.Result{
		Code: errs.RedemptionCodeNotFoundErr.Code,
		Msg:  errs.RedemptionCodeNotFoundErr.Msg,
	}
)
