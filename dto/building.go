package dto

import "github.com/unisoftcn/estate-api/entity"

type Building struct {
	entity.Building
	FloorCount int `json:"floorCount"`
	RoomCount  int `json:"roomCount"`
}
