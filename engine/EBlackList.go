package engine

import (
	"infotech.umm.ac.id/milab/structs"
	"strconv"
)

func GetBlackLists(idb *InDB, offset string) (data []structs.BlackListWithAssociation, err error) {
	o, _ := strconv.Atoi(offset)
	r := idb.DB.
		Limit(10).Offset(o).Preload("User").
		Find(&data)

	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}

func CreateBlacklist(idb *InDB, newBlacklist structs.BlackList) (structs.BlackList, error) {
	r := idb.DB.
		Create(&newBlacklist)

	if r.Error != nil {
		return newBlacklist, r.Error
	}
	return newBlacklist, nil
}
