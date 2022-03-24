package entity

import (
	"github.com/vuuvv/orca/orm"
)

type Room struct {
	orm.Id
	EstateId   int64  `json:"estateId" gorm:"comment:所属园区"`
	OrgPath    string `json:"orgPath" gorm:"comment:所属机构路径"`
	BuildingId int64  `json:"buildingId" gorm:"comment:所属楼栋"`
	FloorId    int64  `json:"floorId" gorm:"comment:所属楼层"`
	Name       string `json:"name" gorm:"comment:楼层名称"`
	orm.Entity

	Floor *Floor `json:"floor" gorm:"-"`
}

func (this *Room) TableName() string {
	return "t_room"
}

func (this *Room) TableTitle() string {
	return "房间"
}
