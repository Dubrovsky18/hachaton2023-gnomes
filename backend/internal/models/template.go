package models

import (
	"gorm.io/gorm"
)

type Template struct {
	*gorm.Model
	TemplateName string `json:"templateName"`
	Html         string `json:"html"`
}
