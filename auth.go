// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package gosdk

type Auth struct {
	Login *Login

	sdkConfiguration sdkConfiguration
}

func newAuth(sdkConfig sdkConfiguration) *Auth {
	return &Auth{
		sdkConfiguration: sdkConfig,
		Login:            newLogin(sdkConfig),
	}
}
