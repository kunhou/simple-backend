package setting

// setting repo use gorm
import (
	"context"

	"gorm.io/gorm"

	"github/kunhou/simple-backend/entity"
	"github/kunhou/simple-backend/usecase/setting"
)

// create a struct for setting repository
type SettingRepo struct {
	db *gorm.DB
}

// new setting repository
func NewSettingRepo(db *gorm.DB) setting.ISettingRepository {
	return &SettingRepo{
		db: db,
	}
}

// implement ISettingRepository interface
func (r *SettingRepo) GetSettings(ctx context.Context) ([]*entity.Setting, error) {
	var settings []*entity.Setting
	if err := r.db.WithContext(ctx).Find(&settings).Error; err != nil {
		return nil, err
	}
	return settings, nil
}

func (r *SettingRepo) GetSettingByName(ctx context.Context, name string) (*entity.Setting, error) {
	var setting entity.Setting
	if err := r.db.WithContext(ctx).Where("name = ?", name).First(&setting).Error; err != nil {
		return nil, err
	}
	return &setting, nil
}

func (r *SettingRepo) CreateSetting(ctx context.Context, setting *entity.Setting) (*entity.Setting, error) {
	if err := r.db.WithContext(ctx).Create(setting).Error; err != nil {
		return nil, err
	}
	return setting, nil
}

func (r *SettingRepo) UpdateSetting(ctx context.Context, id uint, setting *entity.Setting) (*entity.Setting, error) {
	if err := r.db.WithContext(ctx).Model(&entity.Setting{}).Where("id = ?", id).Updates(setting).Error; err != nil {
		return nil, err
	}
	return setting, nil
}

func (r *SettingRepo) UpdateSettingByName(ctx context.Context, name string, setting *entity.Setting) (*entity.Setting, error) {
	res := r.db.WithContext(ctx).Model(&entity.Setting{}).Where("name = ?", name).Updates(setting)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return setting, nil
}

func (r *SettingRepo) DeleteSetting(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Setting{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *SettingRepo) DeleteSettingByName(ctx context.Context, name string) error {
	if err := r.db.WithContext(ctx).Where("name = ?", name).Delete(&entity.Setting{}).Error; err != nil {
		return err
	}
	return nil
}
