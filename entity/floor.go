package entity

import (
	"github.com/vuuvv/orca/orm"
)

type Floor struct {
	orm.Id
	EstateId   int64  `json:"estateId" gorm:"comment:所属园区"`
	OrgPath    string `json:"orgPath" gorm:"comment:所属机构路径"`
	BuildingId int64  `json:"buildingId" gorm:"comment:所属楼栋"`
	Name       string `json:"name" gorm:"comment:楼层名称"`
	orm.Entity

	Building *Building `json:"building" gorm:"-"`
}

func (this *Floor) TableName() string {
	return "t_floor"
}

func (this *Floor) TableTitle() string {
	return "楼层"
}
