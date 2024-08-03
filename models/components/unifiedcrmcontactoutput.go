// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

type UnifiedCrmContactOutput struct {
	// The first name of the contact
	FirstName *string `json:"first_name"`
	// The last name of the contact
	LastName *string `json:"last_name"`
	// The email addresses of the contact
	EmailAddresses []Email `json:"email_addresses,omitempty"`
	// The phone numbers of the contact
	PhoneNumbers []Phone `json:"phone_numbers,omitempty"`
	// The addresses of the contact
	Addresses []Address `json:"addresses,omitempty"`
	// The UUID of the user who owns the contact
	UserID *string `json:"user_id,omitempty"`
	// The custom field mappings of the contact between the remote 3rd party & Panora
	FieldMappings map[string]any `json:"field_mappings,omitempty"`
	// The UUID of the contact
	ID *string `json:"id,omitempty"`
	// The id of the contact in the context of the Crm 3rd Party
	RemoteID *string `json:"remote_id,omitempty"`
	// The remote data of the contact in the context of the Crm 3rd Party
	RemoteData map[string]any `json:"remote_data,omitempty"`
	// The created date of the object
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The modified date of the object
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
}

func (u UnifiedCrmContactOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedCrmContactOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedCrmContactOutput) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *UnifiedCrmContactOutput) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *UnifiedCrmContactOutput) GetEmailAddresses() []Email {
	if o == nil {
		return nil
	}
	return o.EmailAddresses
}

func (o *UnifiedCrmContactOutput) GetPhoneNumbers() []Phone {
	if o == nil {
		return nil
	}
	return o.PhoneNumbers
}

func (o *UnifiedCrmContactOutput) GetAddresses() []Address {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *UnifiedCrmContactOutput) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *UnifiedCrmContactOutput) GetFieldMappings() map[string]any {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}

func (o *UnifiedCrmContactOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedCrmContactOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedCrmContactOutput) GetRemoteData() map[string]any {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *UnifiedCrmContactOutput) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UnifiedCrmContactOutput) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}
