package entities

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID        uint   `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	Desc      string `json:"desc" form:"desc"`
	Status    string `json:"status" form:"status"`
	UserID    uint   `json:"user_id" form:"user_id"`
	ProjectID uint   `json:"project_id" form:"project_id" gorm:"default:NULL"`
}
