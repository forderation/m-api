package engine

type StandardResponse struct {
	Code string
	Message string
	Data interface{}
}

const CodeOkResponse = "00"