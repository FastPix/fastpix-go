package components

type NotFoundErrorError struct {
	Code        *int64  `json:"code,omitempty"`
	Message     *string `json:"message,omitempty"`
	Description *string `json:"description,omitempty"`
}

func (n *NotFoundErrorError) GetCode() *int64 {
	if n == nil {
		return nil
	}
	return n.Code
}

func (n *NotFoundErrorError) GetMessage() *string {
	if n == nil {
		return nil
	}
	return n.Message
}

func (n *NotFoundErrorError) GetDescription() *string {
	if n == nil {
		return nil
	}
	return n.Description
}
