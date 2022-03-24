package entity

import (
	"github.com/unisoftcn/user-api/entity"
	"github.com/vuuvv/orca/orm"
)

type Building struct {
	orm.Id
	EstateId int64 `json:"estateId" gorm:"comment:所属园区"`
	OrgPath string `json:"orgPath" gorm:"comment:所属机构路径"`
	Name     string `json:"name" gorm:"comment:楼栋名称"`
	orm.Entity

	Estate *entity.Org `json:"estate" gorm:"-"`
}

func (this *Building) TableName() string {
	return "t_building"
}

func (this *Building) TableTitle() string {
	return "楼栋"
}
