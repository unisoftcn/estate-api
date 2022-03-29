package forms

import "vuuvv.cn/unisoftcn/orca/orm"

type Floor struct {
	orm.Id
	BuildingId int64  `json:"buildingId" valid:"required~请指定楼栋名称"`
	Name       string `json:"name" valid:"required~请填写楼层名称"`
}
