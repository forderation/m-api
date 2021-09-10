package structs

import "time"

type CertificationSchedule struct {
	Model
	CertificationSessionID uint
	Location               int
	ScheduleFrom           time.Time
	ScheduleUntil          time.Time
	Limit                  int
	Active                 int
}
