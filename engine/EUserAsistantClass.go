package engine

import (
	"gorm.io/gorm/clause"
	"infotech.umm.ac.id/milab/structs"
)

func GetUserAssistantClass(idb *InDB, ids []int) (data []structs.UserAssistantClassWithAssociation, err error) {

	r := idb.DB.Where("user_assistant_id IN (?)", ids).Find(&data)

	if r.Error != nil {
		return data, r.Error
	}

	return data, err
}

func GetUserAssistant(idb *InDB, ids string) (data []structs.UserAssistantWithAssociation, err error) {
	r := idb.DB.
		Where("user_id IN (?)", ids).
		Find(&data)

	if r.Error != nil {
		return data, r.Error
	}

	return data, err
}

func GetClassByUserAssistant(idb *InDB, ids []int) (data []structs.ClassWithAssociation, err error) {
	r := idb.DB.Preload("ClassCategory").Find(&data, ids)

	if r.Error != nil {
		return data, r.Error
	}

	return data, err
}

func GetUserAssistantClasses(idb *InDB, userAssistantID string) (data []structs.UserAssistantClassWithAssociation, err error) {
	r := idb.DB.
		Preload(clause.Associations).
		Preload("Class.ClassCategory").
		Preload("Class.UserAssistant").
		Where("user_assistant_id = ?", userAssistantID).
		Find(&data)

	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}
