package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type SysArea struct {
	service.Service
}

// GetPage 获取SysArea列表
func (e *SysArea) GetPage(c *dto.SysAreaReq, p *actions.DataPermission, list *[]models.SysAreaSearch, count *int64) error {
	var err error
	var data models.SysArea
	e.Log.Errorf("arrive")
	err = e.Orm.Debug().Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetByAreaCode()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("Service GetSysAreaPage error:%s", err)
		return err
	}
	return nil
}
