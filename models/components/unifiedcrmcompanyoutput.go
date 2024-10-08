// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// CreatedAt - The created date of the object
type CreatedAt struct {
}

// ModifiedAt - The modified date of the object
type ModifiedAt struct {
}

type UnifiedCrmCompanyOutput struct {
	// The name of the company
	Name *string `json:"name"`
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
	PhoneNumbers []Phone `json:"phone_numbers,omitempty"`
	// The custom field mappings of the company between the remote 3rd party & Panora
	FieldMappings map[string]any `json:"field_mappings,omitempty"`
	// The UUID of the company
	ID *string `json:"id,omitempty"`
	// The id of the company in the context of the Crm 3rd Party
	RemoteID *string `json:"remote_id,omitempty"`
	// The remote data of the company in the context of the Crm 3rd Party
	RemoteData map[string]any `json:"remote_data,omitempty"`
	// The created date of the object
	CreatedAt *CreatedAt `json:"created_at,omitempty"`
	// The modified date of the object
	ModifiedAt *ModifiedAt `json:"modified_at,omitempty"`
}

func (o *UnifiedCrmCompanyOutput) GetName() *string {
	if o == nil {
		return nil
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

func (o *UnifiedCrmCompanyOutput) GetFieldMappings() map[string]any {
	if o == nil {
		return nil
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

func (o *UnifiedCrmCompanyOutput) GetRemoteData() map[string]any {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *UnifiedCrmCompanyOutput) GetCreatedAt() *CreatedAt {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UnifiedCrmCompanyOutput) GetModifiedAt() *ModifiedAt {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}
