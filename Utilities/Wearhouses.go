package utilities

import (
	dbstore "stockManagment/DbStore"
	models "stockManagment/DbStore/Models"
	dtos "stockManagment/Dtos"
	"time"

	"github.com/google/uuid"
)

func GetAllWearhouses(companyId string) []models.Warehouse {
	db, err := dbstore.Connect()

	if err != nil {
		return nil
	}
	defer dbstore.Disconnect()

	var wearehouses []models.Warehouse
	err = db.Where("company_id = ?", companyId).Find(&wearehouses).Error

	if err != nil {
		return nil
	}

	return wearehouses
}

func CreateWearhouse(name string, companyId string, userId string) bool {
	db, err := dbstore.Connect()
	if err != nil {
		return false
	}

	defer dbstore.Disconnect()

	newWarehouse := models.Warehouse{
		ID:           uuid.New().String(),
		Name:         name,
		CreatorID:    userId,
		CreationDate: time.Now(),
		CompanyID:    companyId,
	}

	result := db.Create(&newWarehouse)
	if result.Error != nil {
		return false
	}

	return true
}

func DeleteWarehouse(id string) error {
	db, err := dbstore.Connect()

	if err != nil {
		return err
	}
	defer dbstore.Disconnect()

	var warehouse models.Warehouse
	err = db.Where("id = ?", id).Find(&warehouse).Error
	if err != nil {
		return err
	}

	err = db.Delete(&warehouse).Error
	if err != nil {
		return err
	}

	return nil
}

func GetWarehouseById(id string) (models.Warehouse, error) {
	db, err := dbstore.Connect()
	var warehouse models.Warehouse

	if err != nil {
		return warehouse, err
	}
	defer dbstore.Disconnect()

	err = db.Where("id = ?", id).Find(&warehouse).Error
	if err != nil {
		return warehouse, err
	}

	return warehouse, nil
}

func EditWarehouse(warehouseInfo dtos.WarehouseInfo) error {
	db, err := dbstore.Connect()

	if err != nil {
		return err
	}
	defer dbstore.Disconnect()

	var warehouse models.Warehouse
	err = db.Where("id = ?", warehouseInfo.Id).Find(&warehouse).Error
	if err != nil {
		return err
	}

	warehouse.Name = warehouseInfo.Name

	err = db.Save(&warehouse).Error
	if err != nil {
		return err
	}

	return nil
}
