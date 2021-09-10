package engine

import (
	"fmt"
	"gorm.io/gorm/clause"
	"infotech.umm.ac.id/milab/structs"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func GetClassTaskByClassId(idb *InDB, id string) (data []structs.TaskWithAssociation, err error) {

	r := idb.DB.
		Where("class_id = (?)", id).
		Find(&data)

	if r.Error != nil {
		return data, r.Error
	}

	return data, err
}

func GetSingleTaskUsingID(idb *InDB, taskId string) (data structs.Task, err error) {
	r := idb.DB.
		Where("id = (?)", taskId).
		First(&data)

	if r.Error != nil {
		return data, r.Error
	}

	return data, err
}

func GetSingleAssociatedTaskUsingID(idb *InDB, taskId string) (data structs.TaskWithAssociation, err error) {
	r := idb.DB.
		Preload(clause.Associations).
		Preload("TaskDetails.UserAssistant.Details").
		Where("id = ?", taskId).
		First(&data)

	if r.Error != nil {
		return data, r.Error
	}

	return data, err
}

func GetSingleStudentAssociatedTaskUsingID(idb *InDB, taskId string) (data structs.StudentTaskWithAssociation, err error) {
	r := idb.DB.
		Preload(clause.Associations).
		Where("id = ?", taskId).
		First(&data)

	if r.Error != nil {
		return data, r.Error
	}

	return data, err
}

func CreateTask(idb *InDB, newTask structs.Task) (structs.Task, error) {
	newTask.ExpiredDateTime = time.Now().Add(time.Hour * 24 * 7)
	newTask.CodeSubmissionDate = time.Now().Add(time.Hour * 24 * 7)

	r := idb.DB.
		Create(&newTask)
	if r.Error != nil {
		return newTask, r.Error
	}

	var UserStudentIds []uint
	r = idb.DB.
		Table("class c").
		Select("usc.user_student_id").
		Joins("RIGHT JOIN `user_student_class` usc ON c.id = usc.class_id").
		Where("c.id = ?", newTask.ClassID).
		Scan(&UserStudentIds)

	_, err := GenerateInitialTaskDetail(idb, newTask, UserStudentIds)
	if err != nil {
		return newTask, r.Error
	}

	return newTask, nil
}

func UpdateTask(idb *InDB, newTask structs.Task) (structs.Task, error) {
	r := idb.DB.
		Save(&newTask)

	if r.Error != nil {
		return newTask, r.Error
	}

	return newTask, nil
}

func UpdateAssociatedTask(idb *InDB, newTask structs.StudentTaskWithAssociation) (structs.StudentTaskWithAssociation, error) {

	for _, v := range newTask.TaskDetails {
		t := structs.TaskDetail{
			Model:           v.Model,
			TaskId:          v.TaskId,
			UserAssistantId: v.UserAssistantId,
			UserStudentId:   v.UserStudentId,
			Grade:           v.Grade,
			GradeEachPoint:  v.GradeEachPoint,
		}
		_, err := UpdateTaskDetail(idb, t)
		if err != nil {
			return newTask, err
		}
	}

	r := idb.DB.
		Preload(clause.Associations).
		Save(&newTask)

	if r.Error != nil {
		return newTask, r.Error
	}

	return newTask, nil
}

func ScrambleAssistant(idb *InDB, taskDetails []structs.StudentTaskDetailWithAssociation, classId uint) (result []structs.StudentTaskDetailWithAssociation, err error) {
	classData, err := GetSingleClassUsingID(idb, strconv.Itoa(int(classId)))
	if err != nil {
		return nil, err
	}

	var assistantIDs []uint
	for _, v := range classData.UserAssistantClasses {
		assistantIDs = append(assistantIDs, v.UserAssistantID)
	}

	if len(assistantIDs) <= 0 {
		return nil, fmt.Errorf("no assistant in class")
	}

	singleLen := int64(math.RoundToEven(float64(len(taskDetails) / len(assistantIDs))))

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(taskDetails), func(i, j int) { taskDetails[i], taskDetails[j] = taskDetails[j], taskDetails[i] })

	ct := 0
	aid := 0
	for i, _ := range taskDetails {
		taskDetails[i].UserAssistantId = &assistantIDs[aid]
		if int64(ct) == singleLen {
			aid++
			ct = 0
		} else {
			ct++
		}

	}

	return taskDetails, nil

}
