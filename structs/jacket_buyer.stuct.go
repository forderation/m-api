package structs

import "time"

type JacketBuyer struct {
	Model
	IdJacketSession  uint
	IdUserBuyer      uint
	JacketBuyerSize  string
	FittingDate      time.Time
	IsCostume        int
	ArmLength        int
	StomachCircle    int
	ChestCircle      int
	WideBack         int
	PayDates         time.Time
	PayMethod        string
	DatetimeTake     time.Time
	Notes            string
	UserAdminIdGiver uint
}
