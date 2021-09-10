package structs

import "time"

type Task struct {
	Model
	LockDate           *time.Time
	ClassID            uint
	UserAssistantID    uint
	ExpiredDateTime    time.Time
	Title              string
	Description        string
	FinalExam          bool
	Extra              bool
	GradePoints        string
	Percentage         int
	CodeSubmissionDate time.Time
	FileScreenshot     string
}

func (TaskWithAssociation) TableName() string {
	return "task"
}

type TaskWithAssociation struct {
	Task
	UserAssistant UserAssistant                      `gorm:"foreignkey:ID;references:UserAssistantID"`
	TaskDetails   []StudentTaskDetailWithAssociation `gorm:"foreignkey:TaskId;references:ID"`
	Class         Class                              `gorm:"foreignkey:ID;references:ClassID"`
}

func (StudentTaskWithAssociation) TableName() string {
	return "task"
}

type StudentTaskWithAssociation struct {
	Task
	TaskDetails []StudentTaskDetailWithAssociation `gorm:"foreignkey:TaskId;references:ID"`
}
