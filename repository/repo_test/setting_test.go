package repotest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"

	"github/kunhou/simple-backend/entity"
	"github/kunhou/simple-backend/pkg/types"
	"github/kunhou/simple-backend/repository/setting"
)

func Test_CreateSetting(t *testing.T) {
	t.Log("test")
	repo := setting.NewSettingRepo(testDB)
	v := datatypes.JSON{}
	err := v.Scan([]byte(`{"test": "test"}`))
	assert.NoError(t, err)

	setting, err := repo.CreateSetting(context.TODO(), &entity.Setting{
		Name:  types.Ptr("test"),
		Value: &v,
	})
	assert.NoError(t, err)
	assert.Equal(t, uint(1), setting.ID)
	assert.Equal(t, "test", *setting.Name)
}
