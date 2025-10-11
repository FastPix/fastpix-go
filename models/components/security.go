

package components

type Security struct {
	Username *string `security:"scheme,type=http,subtype=basic,name=username,env=fastpixgo_username"`
	Password *string `security:"scheme,type=http,subtype=basic,name=password,env=fastpixgo_password"`
}

func (s *Security) GetUsername() *string {
	if s == nil {
		return nil
	}
	return s.Username
}

func (s *Security) GetPassword() *string {
	if s == nil {
		return nil
	}
	return s.Password
}
