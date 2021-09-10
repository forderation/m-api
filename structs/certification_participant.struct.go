package structs

type CertificationParticipant struct {
	Model
	CertificationID   uint
	UserIDParticipant uint
	CertificationTake bool
	ReceiptUploaded   bool
	ValidationPayment bool
	FinalGrade        int
	Done              bool
	GradeStatus       int
}
