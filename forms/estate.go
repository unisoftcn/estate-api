package forms

import (
	"github.com/vuuvv/orca/orm"
)

type EstateData struct {
	Lnglat  []string `json:"lnglat"`
	Contact string   `json:"contact" valid:"required~请填写联系人"`
	Phone   string   `json:"phone" valid:"required~请填写联系电话"`
	Remark  string   `json:"remark"`
}

type Estate struct {
	orm.Id
	ParentId int64      `json:"parentId"`
	Name     string     `json:"name" valid:"required~请填写园区名称"`
	Address  string     `json:"address" valid:"required~请填写园区地址"`
	AreaCode string     `json:"areaCode" valid:"required~请选择区域"`
	AreaName string     `json:"-"`
	Logo     string     `json:"logo"`
	RawData  EstateData `json:"rawData"`
}
