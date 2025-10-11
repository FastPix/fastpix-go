

package components

// SimulcastUnavailableError - Returns the problem that has occured.
type SimulcastUnavailableError struct {
	// An error code indicating the type of the error.
	//
	Code *int64 `json:"code,omitempty"`
	// A descriptive message providing more details for the error
	//
	Message *string `json:"message,omitempty"`
	// A detailed explanation of the possible causes for the error.
	//
	Description *string `json:"description,omitempty"`
}

func (s *SimulcastUnavailableError) GetCode() *int64 {
	if s == nil {
		return nil
	}
	return s.Code
}

func (s *SimulcastUnavailableError) GetMessage() *string {
	if s == nil {
		return nil
	}
	return s.Message
}

func (s *SimulcastUnavailableError) GetDescription() *string {
	if s == nil {
		return nil
	}
	return s.Description
}
