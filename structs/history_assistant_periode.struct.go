package structs

import "time"

func (HistoryAssistantPeriod) TableName() string {
	return "history_assistant_periode"
}

type HistoryAssistantPeriod struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uint
	StartDate time.Time
	EndDate   time.Time
}
