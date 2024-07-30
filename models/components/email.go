// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Email struct {
	// The email address
	EmailAddress string `json:"email_address"`
	// The email address type. Authorized values are either PERSONAL or WORK.
	EmailAddressType string `json:"email_address_type"`
	// The owner type of an email
	OwnerType *string `json:"owner_type,omitempty"`
}

func (o *Email) GetEmailAddress() string {
	if o == nil {
		return ""
	}
	return o.EmailAddress
}

func (o *Email) GetEmailAddressType() string {
	if o == nil {
		return ""
	}
	return o.EmailAddressType
}

func (o *Email) GetOwnerType() *string {
	if o == nil {
		return nil
	}
	return o.OwnerType
}