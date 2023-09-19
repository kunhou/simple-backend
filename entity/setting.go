package entity

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// create a struct for setting entity
type Setting struct {
	ID        uint            `json:"id" gorm:"primarykey"`
	Name      *string         `json:"name" gorm:"type:varchar(255);uniqueIndex:where:deleted_at IS NULL"`
	Value     *datatypes.JSON `json:"value" gorm:"type:jsonb;default:'{}'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
