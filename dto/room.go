package dto

import "vuuvv.cn/youku/estate-api/entity"

type Room struct {
	entity.Floor
	BuildingName string `json:"buildingName"`
	FloorName    string `json:"floorName"`
}
