package response

const (
	Success           Code = "VOT0000"
	ServerError       Code = "VOT0001"
	BadRequest        Code = "VOT0002"
	InvalidRequest    Code = "VOT0004"
	Failed            Code = "VOT0073"
	Pending           Code = "VOT0050"
	InvalidInputParam Code = "VOT0032"

	Unauthorized   Code = "VOT0502"
	Forbidden      Code = "VOT0503"
	GatewayTimeout Code = "VOT0048"
)

type Code string

var codeMap = map[Code]string{
	Success:           "success",
	Failed:            "failed",
	Pending:           "pending",
	BadRequest:        "bad or invalid request",
	Unauthorized:      "Unauthorized Token",
	GatewayTimeout:    "Gateway Timeout",
	ServerError:       "Internal Server Error",
	InvalidInputParam: "Other invalid argument",
}

func (c Code) GetStatus() string {
	switch c {
	case Success:
		return "SUCCESS"

	default:
		return "FAILED"
	}
}

func (c Code) GetMessage() string {
	return codeMap[c]
}

func (c Code) GetVersion() string {
	return "1"
}
