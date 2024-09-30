package utilities

import (
	dbstore "stockManagment/DbStore"
	models "stockManagment/DbStore/Models"
	dtos "stockManagment/Dtos"
	"time"

	"github.com/google/uuid"
)

func CreateColor(colorInfo *dtos.ColorInfo, companyId string, creatorId string) bool {
	db, err := dbstore.Connect()
	if err != nil {
		return false
	}
	defer dbstore.Disconnect()

	newColor := models.Color{
		ID:           uuid.New().String(),
		Name:         colorInfo.Name,
		Color:        colorInfo.Ref,
		CreationDate: time.Now(),
		CreatorID:    creatorId,
		CompanyID:    companyId,
	}

	result := db.Create(&newColor)
	return result.Error == nil
}
