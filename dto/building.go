package dto

import "vuuvv.cn/unisoftcn/estate-api/entity"

type Building struct {
	entity.Building
	FloorCount int `json:"floorCount"`
	RoomCount  int `json:"roomCount"`
}
