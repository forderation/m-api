package engine

import (
	"infotech.umm.ac.id/milab/structs"
	"strconv"
)

func GetBlacklistCategories(idb *InDB, offset string) (data []structs.BlackListCategory, err error) {
	o, _ := strconv.Atoi(offset)
	r := idb.DB.
		Limit(10).Offset(o).
		Find(&data)

	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}
