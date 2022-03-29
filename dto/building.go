package dto

import "vuuvv.cn/youku/estate-api/entity"

type Building struct {
	entity.Building
	FloorCount int `json:"floorCount"`
	RoomCount  int `json:"roomCount"`
}
