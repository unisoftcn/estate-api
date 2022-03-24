package forms

import "github.com/vuuvv/orca/orm"

type Building struct {
	orm.Id
	EstatePath string `json:"estatePath" valid:"required~请指定园区"`
	Name       string `json:"name" valid:"required~请填写楼栋名称"`
}
