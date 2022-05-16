package response

type Wrapper struct {
	ResponseCode Code        `json:"code"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

func NewResponse(code Code, data interface{}) *Wrapper {
	return &Wrapper{
		ResponseCode: code,
		Message:      code.GetMessage(),
		Data:         data,
	}
}

func (w *Wrapper) SetResponseMessage(customMessage string) {
	w.Message = customMessage
}
