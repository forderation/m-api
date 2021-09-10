package engine

import (
	"gorm.io/gorm/clause"
	"infotech.umm.ac.id/milab/structs"
)

func CreateUserStudentClass(idb *InDB, userStudentClass structs.UserStudentClass) (data structs.UserStudentClass, err error) {
	r := idb.DB.
		Create(&userStudentClass)

	if r.Error != nil {
		return data, r.Error
	}

	return data, err
}

func GetUserStudentClasses(idb *InDB, userStudentID string) (data []structs.UserStudentClassWithAssociation, err error) {
	r := idb.DB.
		Preload(clause.Associations).
		Preload("Class.ClassCategory").
		Preload("Class.UserAssistant").
		Preload("Class.UserInstructor").
		Where("user_student_id = ?", userStudentID).
		Find(&data)

	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}
