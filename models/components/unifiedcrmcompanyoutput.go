// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UnifiedCrmCompanyOutputFieldMappings struct {
}

type UnifiedCrmCompanyOutputRemoteData struct {
}

type UnifiedCrmCompanyOutputCreatedAt struct {
}

type UnifiedCrmCompanyOutputModifiedAt struct {
}

type UnifiedCrmCompanyOutput struct {
	// The name of the company
	Name string `json:"name"`
	// The industry of the company. Authorized values can be found in the Industry enum.
	Industry *string `json:"industry,omitempty"`
	// The number of employees of the company
	NumberOfEmployees *float64 `json:"number_of_employees,omitempty"`
	// The UUID of the user who owns the company
	UserID *string `json:"user_id,omitempty"`
	// The email addresses of the company
	EmailAddresses []Email `json:"email_addresses,omitempty"`
	// The addresses of the company
	Addresses []Address `json:"addresses,omitempty"`
	// The phone numbers of the company
	PhoneNumbers  []Phone                              `json:"phone_numbers,omitempty"`
	FieldMappings UnifiedCrmCompanyOutputFieldMappings `json:"field_mappings"`
	// The UUID of the company
	ID *string `json:"id,omitempty"`
	// The id of the company in the context of the Crm 3rd Party
	RemoteID   *string                           `json:"remote_id,omitempty"`
	RemoteData UnifiedCrmCompanyOutputRemoteData `json:"remote_data"`
	CreatedAt  UnifiedCrmCompanyOutputCreatedAt  `json:"created_at"`
	ModifiedAt UnifiedCrmCompanyOutputModifiedAt `json:"modified_at"`
}

func (o *UnifiedCrmCompanyOutput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UnifiedCrmCompanyOutput) GetIndustry() *string {
	if o == nil {
		return nil
	}
	return o.Industry
}

func (o *UnifiedCrmCompanyOutput) GetNumberOfEmployees() *float64 {
	if o == nil {
		return nil
	}
	return o.NumberOfEmployees
}

func (o *UnifiedCrmCompanyOutput) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *UnifiedCrmCompanyOutput) GetEmailAddresses() []Email {
	if o == nil {
		return nil
	}
	return o.EmailAddresses
}

func (o *UnifiedCrmCompanyOutput) GetAddresses() []Address {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *UnifiedCrmCompanyOutput) GetPhoneNumbers() []Phone {
	if o == nil {
		return nil
	}
	return o.PhoneNumbers
}

func (o *UnifiedCrmCompanyOutput) GetFieldMappings() UnifiedCrmCompanyOutputFieldMappings {
	if o == nil {
		return UnifiedCrmCompanyOutputFieldMappings{}
	}
	return o.FieldMappings
}

func (o *UnifiedCrmCompanyOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedCrmCompanyOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedCrmCompanyOutput) GetRemoteData() UnifiedCrmCompanyOutputRemoteData {
	if o == nil {
		return UnifiedCrmCompanyOutputRemoteData{}
	}
	return o.RemoteData
}

func (o *UnifiedCrmCompanyOutput) GetCreatedAt() UnifiedCrmCompanyOutputCreatedAt {
	if o == nil {
		return UnifiedCrmCompanyOutputCreatedAt{}
	}
	return o.CreatedAt
}

func (o *UnifiedCrmCompanyOutput) GetModifiedAt() UnifiedCrmCompanyOutputModifiedAt {
	if o == nil {
		return UnifiedCrmCompanyOutputModifiedAt{}
	}
	return o.ModifiedAt
}
