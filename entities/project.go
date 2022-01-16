package entities

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	ID    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name" gorm:"unique;"`
	Tasks []Task `json:"tasks" form:"tasks"`
}
