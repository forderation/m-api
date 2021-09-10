package engine

import (
	"crypto/md5"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"infotech.umm.ac.id/milab/structs"
)

func GetAllClassActive(idb *InDB, id *string) (data []structs.ClassWithAssociation, err error) {
	var (
		dataClassCategory []structs.ClassCategoryWithAssociation
		idsClassCategory  []uint
		r                 *gorm.DB
	)

	if id != nil {
		r = idb.DB.Where("id = (?)", id).Where("parent != 0").Where("is_teori_class != 1").
			Where("active = 1").Find(&dataClassCategory)
	} else {
		r = idb.DB.Where("parent != 0").Where("is_teori_class != 1").
			Where("active = 1").Find(&dataClassCategory)
	}

	if r.Error != nil {
		return data, r.Error
	}

	for _, cC := range dataClassCategory {
		idsClassCategory = append(idsClassCategory, cC.ID)
	}

	data, err = GetAllClassByCategoryIds(idb, idsClassCategory)
	if err != nil {
		return data, err
	}

	return data, err
}

func GetAllClassByCategoryIds(idb *InDB, ids []uint) (data []structs.ClassWithAssociation, err error) {
	r := idb.DB.
		Where("class_category_id IN (?)", ids).
		Find(&data)
	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}

func GetSingleClassUsingID(idb *InDB, id string) (data structs.ClassWithAssociation, err error) {
	r := idb.DB.
		Preload(clause.Associations).
		Preload("UserStudents.UserStudentDetail.UserDetail").
		Where("id = ?", id).
		First(&data)
	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}

func GetSingleStudentClassUsingID(idb *InDB, id string, studentID uint) (data structs.StudentClassWithAssociation, err error) {
	r := idb.DB.
		Preload(clause.Associations).
		Preload("Tasks.TaskDetails", "user_student_id = (?)", studentID).
		Preload("Tasks.TaskDetails.UserAssistant.Details").
		Where("id = ?", id).
		First(&data)
	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}

func DownloadPDF(task_id string) string {
	concat := fmt.Sprintf("%sabc", task_id)
	hash := md5.Sum([]byte(concat))
	folderName := fmt.Sprintf("%x%s", hash, task_id)
	hashFileName := md5.Sum([]byte(fmt.Sprintf("%smodule_upload", task_id)))
	concatFileName := fmt.Sprintf("%x%s", hashFileName, task_id)
	fmt.Printf(concatFileName)
	path := fmt.Sprintf("https://infotech.umm.ac.id/uploads/file_task/%s//%smodule.pdf", folderName, concatFileName)
	return path
}
