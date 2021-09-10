package structs

import "time"

type ItemLoan struct {
	Model
	ItemCode   int
	MemberId   int
	Name       string
	LoanFor    int
	LoanDate   time.Time
	DueDate    time.Time
	DueTime    string
	IsLent     bool
	IsReturn   bool
	ReturnDate time.Time
	UserId     int
}
