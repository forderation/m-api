package engine

import (
	"gorm.io/gorm/clause"
	"infotech.umm.ac.id/milab/structs"
)

func GetSingleUser(idb *InDB, param string, paramValue string) (data structs.User, err error) {
	r := idb.DB.
		Preload(clause.Associations).
		Where(param+" = ?", paramValue).
		First(&data)

	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}

func GetSingleUserWithPreload(idb *InDB, param string, paramValue string) (data structs.UserWithAssociation, err error) {
	r := idb.DB.
		Preload(clause.Associations).
		Where(param+" = ?", paramValue).
		First(&data)

	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}

func GetUsers(idb *InDB, ids []int64) (data []structs.User, err error) {
	r := idb.DB.
		Where("id IN (?)", ids).
		Find(&data)

	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}
