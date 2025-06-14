// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Security struct {
	Username *string `security:"scheme,type=http,subtype=basic,name=username,env=fastpixsdk_username"`
	Password *string `security:"scheme,type=http,subtype=basic,name=password,env=fastpixsdk_password"`
}

func (o *Security) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *Security) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}
