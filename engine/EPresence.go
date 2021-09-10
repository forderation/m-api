package engine

import (
	"gorm.io/gorm/clause"
	"infotech.umm.ac.id/milab/structs"
)

func GetSingleAssociatedPresenceUsingID(idb *InDB, id string) (data structs.PresenceWithAssociation, err error) {
	r := idb.DB.
		Preload(clause.Associations).
		Preload("PresenceStudents.UserStudent.UserDetail").
		Where("id = ?", id).
		First(&data)

	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}

func CreatePresence(idb *InDB, presence structs.Presence) (structs.Presence, error) {
	r := idb.DB.
		Create(&presence)
	if r.Error != nil {
		return presence, r.Error
	}

	var UserStudentIds []uint
	r = idb.DB.
		Table("class c").
		Select("usc.user_student_id").
		Joins("RIGHT JOIN `user_student_class` usc ON c.id = usc.class_id").
		Where("c.id = ?", presence.ClassId).
		Scan(&UserStudentIds)

	_, err := GenerateInitialPresence(idb, presence, UserStudentIds)
	if err != nil {
		return presence, err
	}

	return presence, nil

}

func GenerateInitialPresence(idb *InDB, presence structs.Presence, studentIds []uint) (data []structs.PresenceStudent, err error) {
	for _, e := range studentIds {
		data = append(data, structs.PresenceStudent{
			PresenceId:      presence.ID,
			UserStudentId:   e,
			UserAssistantId: presence.UserAssistantID,
			Type:            1,
			Late:            "00:00:00",
		})
	}

	r := idb.DB.
		Create(&data)
	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}
