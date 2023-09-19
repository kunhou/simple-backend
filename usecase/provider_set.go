package usecase

import (
	"github.com/google/wire"

	"github/kunhou/simple-backend/usecase/setting"
)

var ProviderSetUsecase = wire.NewSet(
	setting.NewSettingUsecase,
)
