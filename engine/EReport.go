package engine

import "infotech.umm.ac.id/milab/structs"

//func GetBlacklistCategories(idb *InDB) (data []structs.BlackListCategory, err error) {
//	r := idb.DB.Find(&data)
//	if r.Error != nil {
//		return data, r.Error
//	}
//	return data, err
//}

func GetBlacklistCategoryById(idb *InDB, id string) (data []structs.BlackListCategory, err error) {
	r := idb.DB.Where("id = (?)", id).Find(&data)
	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}

func CreateReportStudent(idb *InDB, blackList *structs.BlackList) (err error) {
	result := idb.DB.Create(&blackList)
	if result.Error != nil {
		return err
	}
	return
}
