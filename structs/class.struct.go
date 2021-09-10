package structs

type Class struct {
	Model
	ClassCategoryID     uint
	FullName            string
	ShortName           string
	Active              bool
	UserInstructorID    uint
	UserAssistantID     uint
	UserAdminID         uint
	ScheduleLocation    int
	ScheduleWeek        int
	ScheduleWeeks       int
	ScheduleTimeFrom    int
	ScheduleTimeUntil   int
	PercentagePresence  int
	PercentageTask      int
	PercentageFinalExam int
	PercentageExtra     int
}

func (ClassWithAssociation) TableName() string { return "class" }

type ClassWithAssociation struct {
	Class
	ClassCategory        ClassCategory                `gorm:"foreignKey:ClassCategoryID;references:ID"`
	UserAdmin            UserAdmin                    `gorm:"foreignKey:ID;references:UserAdminID"`
	Presences            []Presence                   `gorm:"foreignKey:ClassId;references:ID"`
	UserInstructor       UserInstructor               `gorm:"foreignKey:ID;references:UserInstructorID"`
	UserStudents         []UserStudentClassWithDetail `gorm:"foreignKey:ClassId;references:ID"`
	UserAssistant        UserAssistant                `gorm:"foreignKey:ID;references:UserAssistantID"`
	UserAssistantClasses []UserAssistantClass         `gorm:"foreignKey:ClassID;references:ID"`
	Tasks                []Task                       `gorm:"foreignKey:ClassID;references:ID"`
}

func (StudentClassWithAssociation) TableName() string { return "class" }

type StudentClassWithAssociation struct {
	Class
	ClassCategory        ClassCategory                `gorm:"foreignKey:ClassCategoryID;references:ID"`
	UserAdmin            UserAdmin                    `gorm:"foreignKey:ID;references:UserAdminID"`
	Presences            []Presence                   `gorm:"foreignKey:ClassId;references:ID"`
	UserInstructor       UserInstructor               `gorm:"foreignKey:ID;references:UserInstructorID"`
	UserAssistant        UserAssistant                `gorm:"foreignKey:ID;references:UserAssistantID"`
	UserAssistantClasses []UserAssistantClass         `gorm:"foreignKey:ClassID;references:ID"`
	Tasks                []StudentTaskWithAssociation `gorm:"foreignKey:ClassID;references:ID"`
}
