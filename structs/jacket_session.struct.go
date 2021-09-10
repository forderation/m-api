package structs

import "time"

type JacketSession struct {
	Model
	Name        string
	Description string
	Price       int
	FinishDate  time.Time
}
