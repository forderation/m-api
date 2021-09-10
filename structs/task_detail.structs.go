package structs

type TaskDetail struct {
	Model
	TaskId          uint
	UserAssistantId *uint
	UserStudentId   uint
	Grade           int
	GradeEachPoint  string
}

func (TaskDetailWithAssociation) TableName() string {
	return "task_detail"
}

type TaskDetailWithAssociation struct {
	TaskDetail
	UserStudent   UserStudent   `gorm:"foreignkey:ID;references:UserStudentId"`
	Task          Task          `gorm:"foreignkey:ID;references:TaskId"`
	UserAssistant UserAssistant `gorm:"foreignkey:ID;references:UserAssistantId"`
}

func (StudentTaskDetailWithAssociation) TableName() string {
	return "task_detail"
}

type StudentTaskDetailWithAssociation struct {
	TaskDetail
	UserAssistant UserAssistantWithAssociationDetail `gorm:"foreignkey:ID;references:UserAssistantId"`
}
