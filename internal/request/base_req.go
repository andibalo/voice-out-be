package request

type BaseRequest struct {
	// Data from Environment / Request processing
	requestedBy string
}

func (b *BaseRequest) SetRequestedBy(in string) {
	b.requestedBy = in
}

func (b *BaseRequest) GetRequestedBy() string {
	return b.requestedBy
}
