package dtos

type WarehouseInfo struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
