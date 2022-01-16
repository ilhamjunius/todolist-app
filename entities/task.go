package entities

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID        uint   `json:"id" form:"id"`
	Name      string `json:"name" form:"name" gorm:"uniqueIndex;not null"`
	Desc      string `json:"desc" form:"desc" gorm:"not null"`
	Status    string `json:"status" form:"status" gorm:"not null"`
	UserID    uint   `json:"user_id" form:"user_id" gorm:"not null"`
	ProjectID uint   `json:"project_id" form:"project_id" gorm:"default:NULL"`
}
