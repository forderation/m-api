package engine

import "infotech.umm.ac.id/milab/structs"

func GetSingleTaskDetail(idb *InDB, param string, paramValue string) (data structs.TaskDetail, err error) {
	r := idb.DB.
		Where(param+" = ?", paramValue).
		First(&data)

	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}

func UpdateTaskDetail(idb *InDB, newTaskDetail structs.TaskDetail) (structs.TaskDetail, error) {
	r := idb.DB.
		Save(&newTaskDetail)

	if r.Error != nil {
		return newTaskDetail, r.Error
	}
	return newTaskDetail, nil
}

func GenerateInitialTaskDetail(idb *InDB, task structs.Task, studentIds []uint) (data []structs.TaskDetail, err error) {
	for _, e := range studentIds {
		data = append(data, structs.TaskDetail{
			TaskId:         task.ID,
			UserStudentId:  e,
			Grade:          0,
			GradeEachPoint: task.GradePoints,
		})
	}

	r := idb.DB.
		Create(&data)
	if r.Error != nil {
		return data, r.Error
	}
	return data, err
}
