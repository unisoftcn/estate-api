package dto

import "github.com/unisoftcn/estate-api/entity"

type Room struct {
	entity.Floor
	BuildingName string `json:"buildingName"`
	FloorName    string `json:"floorName"`
}
