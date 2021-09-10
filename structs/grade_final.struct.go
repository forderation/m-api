package structs

type GradeFinal struct {
	Model
	ClassID        uint
	Nim            string
	SchoolYearFrom int
	SchoolYearTo   int
	Semester       int
	Grade          float64
}
