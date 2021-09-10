package structs

type Certification struct {
	Model
	Name           string
	Description    string
	Price          int
	AllUser        bool
	OnlyStudent    bool
	OnlyAssistant  bool
	OnlyInstructor bool
	Active         int
}
