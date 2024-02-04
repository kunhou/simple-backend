//go:generate mockgen -source ./setting_usecase.go -destination=../mocks/setting_usecase.go -package=mocks
package setting

import (
	"context"

	"gorm.io/gorm"

	"github/kunhou/simple-backend/entity"
	"github/kunhou/simple-backend/pkg/errcraft"
	"github/kunhou/simple-backend/pkg/gormhelper"
	"github/kunhou/simple-backend/pkg/reason"
)

type ISettingRepository interface {
	// crud with context
	GetSettings(ctx context.Context) ([]*entity.Setting, error)
	GetSettingByName(ctx context.Context, name string) (*entity.Setting, error)
	CreateSetting(ctx context.Context, setting *entity.Setting) (*entity.Setting, error)
	UpdateSetting(ctx context.Context, id uint, setting *entity.Setting) (*entity.Setting, error)
	UpdateSettingByName(ctx context.Context, name string, setting *entity.Setting) (*entity.Setting, error)
	DeleteSetting(ctx context.Context, id uint) error
	DeleteSettingByName(ctx context.Context, name string) error
}

type SettingUsecase struct {
	repo ISettingRepository
}

// new setting usecase
func NewSettingUsecase(repo ISettingRepository) *SettingUsecase {
	return &SettingUsecase{
		repo: repo,
	}
}

// implement SettingUsecase interface
func (u *SettingUsecase) GetSettingByName(ctx context.Context, name string) (*entity.Setting, error) {
	setting, err := u.repo.GetSettingByName(ctx, name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errcraft.NotFound(reason.NotFound).SetError(err)
		}
		return nil, errcraft.InternalServer(reason.InternalServer).SetError(err)
	}
	return setting, nil
}

func (u *SettingUsecase) CreateSetting(ctx context.Context, setting *entity.Setting) (*entity.Setting, error) {
	setting, err := u.repo.CreateSetting(ctx, setting)
	if err != nil {
		if gormhelper.ParseDBError(err) == gormhelper.ErrorDuplicateValues {
			return nil, errcraft.DuplicateKey(reason.Duplicate).SetError(err)
		}
		return nil, errcraft.InternalServer(reason.InternalServer).SetError(err)
	}
	return setting, nil
}

func (u *SettingUsecase) UpdateSetting(ctx context.Context, id uint, setting *entity.Setting) (*entity.Setting, error) {
	setting, err := u.repo.UpdateSetting(ctx, id, setting)
	if err != nil {
		if gormhelper.ParseDBError(err) == gormhelper.ErrorDuplicateValues {
			return nil, errcraft.DuplicateKey(reason.Duplicate).SetError(err)
		}
		return nil, errcraft.InternalServer(reason.InternalServer).SetError(err)
	}
	return setting, nil
}

func (u *SettingUsecase) UpdateSettingByName(ctx context.Context, name string, setting *entity.Setting) (*entity.Setting, error) {
	setting, err := u.repo.UpdateSettingByName(ctx, name, setting)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errcraft.NotFound(reason.NotFound).SetError(err)
		}
		return nil, errcraft.InternalServer(reason.InternalServer).SetError(err)
	}
	return setting, nil
}

func (u *SettingUsecase) DeleteSetting(ctx context.Context, id uint) error {
	return u.repo.DeleteSetting(ctx, id)
}

func (u *SettingUsecase) DeleteSettingByName(ctx context.Context, name string) error {
	return u.repo.DeleteSettingByName(ctx, name)
}
