package setting

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"gorm.io/datatypes"

	"github/kunhou/simple-backend/entity"
	"github/kunhou/simple-backend/pkg/types"
	"github/kunhou/simple-backend/usecase/mocks"
)

type settingTestSuite struct {
	suite.Suite
	ctrl            *gomock.Controller
	mockSettingRepo *mocks.MockISettingRepository

	settingUsecase *SettingUsecase
}

func (suite *settingTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockSettingRepo = mocks.NewMockISettingRepository(suite.ctrl)

	suite.settingUsecase = NewSettingUsecase(suite.mockSettingRepo)
}

func (suite *settingTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func TestSettingTestSuite(t *testing.T) {
	suite.Run(t, new(settingTestSuite))
}

func (suite *settingTestSuite) TestGetSettingByName() {
	suite.mockSettingRepo.EXPECT().GetSettingByName(gomock.Any(), gomock.Any()).Return(&entity.Setting{
		ID:    1,
		Name:  types.Ptr("test"),
		Value: types.Ptr[datatypes.JSON]([]byte(`{"test": 123}`)),
	}, nil)
	setting, err := suite.settingUsecase.GetSettingByName(context.Background(), "test")
	suite.NoError(err)
	suite.Equal(uint(1), setting.ID)
	suite.Equal("test", *setting.Name)

	v := datatypes.JSON{}
	err = v.Scan([]byte(`{"test": 123}`))
	suite.NoError(err)
	suite.Equal(v, *setting.Value)
}

func (suite *settingTestSuite) TestCreateSetting() {
	suite.mockSettingRepo.EXPECT().CreateSetting(gomock.Any(), gomock.Any()).Return(&entity.Setting{
		ID:    1,
		Name:  types.Ptr("test"),
		Value: types.Ptr[datatypes.JSON]([]byte(`{"test": 123}`)),
	}, nil)

	setting, err := suite.settingUsecase.CreateSetting(context.Background(), &entity.Setting{
		Name:  types.Ptr("test"),
		Value: types.Ptr[datatypes.JSON]([]byte(`{"test": 123}`)),
	})
	suite.NoError(err)
	suite.Equal(uint(1), setting.ID)
	suite.Equal("test", *setting.Name)

	v := datatypes.JSON{}
	err = v.Scan([]byte(`{"test": 123}`))
	suite.NoError(err)
}

func (suite *settingTestSuite) TestUpdateSetting() {
	suite.mockSettingRepo.EXPECT().UpdateSetting(gomock.Any(), uint(1), &entity.Setting{
		Name:  types.Ptr("test2"),
		Value: types.Ptr[datatypes.JSON]([]byte(`{"test": 123}`)),
	}).Return(&entity.Setting{
		ID:    1,
		Name:  types.Ptr("test2"),
		Value: types.Ptr[datatypes.JSON]([]byte(`{"test": 123}`)),
	}, nil)

	setting, err := suite.settingUsecase.UpdateSetting(context.Background(), uint(1), &entity.Setting{
		Name:  types.Ptr("test2"),
		Value: types.Ptr[datatypes.JSON]([]byte(`{"test": 123}`)),
	})
	suite.NoError(err)
	suite.Equal(uint(1), setting.ID)
	suite.Equal("test2", *setting.Name)

	v := datatypes.JSON{}
	err = v.Scan([]byte(`{"test": 123}`))
	suite.NoError(err)
	suite.Equal(v, *setting.Value)
}

func (suite *settingTestSuite) TestUpdateSettingByName() {
	suite.mockSettingRepo.EXPECT().UpdateSettingByName(gomock.Any(), "test", &entity.Setting{
		Name:  types.Ptr("test2"),
		Value: types.Ptr[datatypes.JSON]([]byte(`{"test": 321}`)),
	}).Return(&entity.Setting{
		ID:    1,
		Name:  types.Ptr("test2"),
		Value: types.Ptr[datatypes.JSON]([]byte(`{"test": 321}`)),
	}, nil)

	setting, err := suite.settingUsecase.UpdateSettingByName(context.Background(), "test", &entity.Setting{
		Name:  types.Ptr("test2"),
		Value: types.Ptr[datatypes.JSON]([]byte(`{"test": 321}`)),
	})
	suite.NoError(err)
	suite.Equal(uint(1), setting.ID)
	suite.Equal("test2", *setting.Name)

	v := datatypes.JSON{}
	err = v.Scan([]byte(`{"test": 321}`))
	suite.NoError(err)
	suite.Equal(v, *setting.Value)
}

func (suite *settingTestSuite) TestDeleteSetting() {
	suite.mockSettingRepo.EXPECT().DeleteSetting(gomock.Any(), uint(1)).Return(nil)

	err := suite.settingUsecase.DeleteSetting(context.Background(), uint(1))
	suite.NoError(err)
}

func (suite *settingTestSuite) TestDeleteSettingByName() {
	suite.mockSettingRepo.EXPECT().DeleteSettingByName(gomock.Any(), "test").Return(nil)

	err := suite.settingUsecase.DeleteSettingByName(context.Background(), "test")
	suite.NoError(err)
}
