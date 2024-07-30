// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Phone struct {
	// The phone number starting with a plus (+) followed by the country code (e.g +336676778890 for France)
	PhoneNumber string `json:"phone_number"`
	// The phone type. Authorized values are either MOBILE or WORK
	PhoneType string `json:"phone_type"`
	// The owner type of a phone number
	OwnerType *string `json:"owner_type,omitempty"`
}

func (o *Phone) GetPhoneNumber() string {
	if o == nil {
		return ""
	}
	return o.PhoneNumber
}

func (o *Phone) GetPhoneType() string {
	if o == nil {
		return ""
	}
	return o.PhoneType
}

func (o *Phone) GetOwnerType() *string {
	if o == nil {
		return nil
	}
	return o.OwnerType
}
