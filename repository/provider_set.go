package repository

import (
	"github.com/google/wire"

	"github/kunhou/simple-backend/pkg/data"
	"github/kunhou/simple-backend/repository/setting"
)

var ProviderSetRepository = wire.NewSet(
	data.NewDB,
	setting.NewSettingRepo,
)
