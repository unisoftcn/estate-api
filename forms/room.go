package forms

import "vuuvv.cn/unisoftcn/orca/orm"

type Room struct {
	orm.Id
	FloorId int64  `json:"floorId" valid:"required~请指定楼层"`
	Name    string `json:"name" valid:"required~请填写房间名称"`
}
