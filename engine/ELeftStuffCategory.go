package engine

import (
	"infotech.umm.ac.id/milab/structs"
	"strconv"
)

func GetLeftStuffCategory(idb *InDB, limit string) (data []structs.LeftStuffCategory, err error) {
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
