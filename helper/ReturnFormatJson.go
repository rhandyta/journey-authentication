package helper

type ReturnFormatJson struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
