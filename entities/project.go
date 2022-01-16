package entities

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	ID    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name" gorm:"uniqueIndex;" validate:"required"`
	Tasks []Task `json:"tasks" form:"tasks"`
}
