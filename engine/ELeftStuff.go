package engine

import (
	"gorm.io/gorm/clause"
	"infotech.umm.ac.id/milab/structs"
	"strconv"
)

func GetLeftStuff(idb *InDB, offset string) (data []structs.LeftStuff, err error) {
	o, _ := strconv.Atoi(offset)
	r := idb.DB.
		Limit(10).
		Offset(o).
		Find(&data)

	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}

func GetSingleLeftStuffUsingID(idb *InDB, id string) (data []structs.LeftStuffWithAssociation, err error) {
	r := idb.DB.
		Preload(clause.Associations).
		Where("id = ?", id).
		First(&data)

	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}
