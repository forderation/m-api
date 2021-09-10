package structs

type CertificationGrade struct {
	Model
	CertificationScheduleID    uint
	CertificationParticipantID uint
	CertificationSessionID     uint
	UserIDAdmin                uint
	Grade                      int
	Pass                       int8
}
