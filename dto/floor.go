package dto

import "vuuvv.cn/unisoftcn/estate-api/entity"

type Floor struct {
	entity.Floor
	BuildingName string `json:"buildingName"`
	RoomCount    int    `json:"roomCount"`
}
