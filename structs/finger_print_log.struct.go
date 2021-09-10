package structs

import "time"

type FingerPrintLog struct {
	ID            uint
	DeletedAt     time.Time
	FingerPrintId int
	ScanDateTime  time.Time
	Verified      int
	Status        int
}
