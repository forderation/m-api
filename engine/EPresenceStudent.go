package engine

import (
	"infotech.umm.ac.id/milab/structs"
	"strconv"
)

func GetPresenceStudent(idb *InDB, limit string) (data []structs.PresenceStudent, err error) {
	limits, _ := strconv.Atoi(limit)
	r := idb.DB.
		Limit(limits).
		Offset(15).
		Find(&data)

	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}

func GetSinglePresenceStudentUsingID(idb *InDB, id string) (data structs.PresenceStudent, err error) {
	r := idb.DB.
		Where("id = ?", id).
		First(&data)

	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}

func UpdatePresenceStudent(idb *InDB, newPresenceStudent structs.PresenceStudent) (structs.PresenceStudent, error) {
	r := idb.DB.
		Debug().
		Save(&newPresenceStudent)

	if r.Error != nil {
		return newPresenceStudent, r.Error
	}

	return newPresenceStudent, nil
}
