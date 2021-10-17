package models

import (
	"go-admin/common/models"
)

type SysArea struct {
	Id           int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	AreaName     string `json:"area_name" gorm:"size:255;comment:区域名称"`
	AreaCode     string `json:"area_code" gorm:"size:128;comment:区域编码"`
	AreaNameFull string `json:"area_name_full" gorm:"size:128;comment:全称"`
	ParAreaCode  string `json:"par_area_code" gorm:"size:16;comment:上级编码"`
	AreaLevel    int    `json:"area_level" gorm:"size:2;comment:区域等级"`
	models.ModelTime
	models.ControlBy
}

// SysAreaSearch 定义返回字段
type SysAreaSearch struct {
	AreaName     string `json:"area_name" gorm:"size:255;comment:区域名称"`
	AreaCode     string `json:"area_code" gorm:"size:128;comment:区域编码"`
}

func (SysArea) TableName() string {
	return "sys_area"
}

func (e *SysArea) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysArea) GetId() interface{} {
	return e.Id
}
