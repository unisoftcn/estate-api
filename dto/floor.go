package dto

import "vuuvv.cn/youku/estate-api/entity"

type Floor struct {
	entity.Floor
	BuildingName string `json:"buildingName"`
	RoomCount    int    `json:"roomCount"`
}
