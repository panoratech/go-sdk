// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

type UnifiedAtsUserOutputFieldMappings struct {
}

type UnifiedAtsUserOutputRemoteData struct {
}

type UnifiedAtsUserOutputCreatedAt struct {
}

type UnifiedAtsUserOutputModifiedAt struct {
}

type UnifiedAtsUserOutput struct {
	// The first name of the user
	FirstName *string `json:"first_name,omitempty"`
	// The last name of the user
	LastName *string `json:"last_name,omitempty"`
	// The email of the user
	Email *string `json:"email,omitempty"`
	// Whether the user is disabled
	Disabled *bool `json:"disabled,omitempty"`
	// The access role of the user
	AccessRole *string `json:"access_role,omitempty"`
	// The remote creation date of the user
	RemoteCreatedAt *time.Time `json:"remote_created_at,omitempty"`
	// The remote modification date of the user
	RemoteModifiedAt *time.Time                        `json:"remote_modified_at,omitempty"`
	FieldMappings    UnifiedAtsUserOutputFieldMappings `json:"field_mappings"`
	// The UUID of the user
	ID *string `json:"id,omitempty"`
	// The remote ID of the user in the context of the 3rd Party
	RemoteID   *string                        `json:"remote_id,omitempty"`
	RemoteData UnifiedAtsUserOutputRemoteData `json:"remote_data"`
	CreatedAt  UnifiedAtsUserOutputCreatedAt  `json:"created_at"`
	ModifiedAt UnifiedAtsUserOutputModifiedAt `json:"modified_at"`
}

func (u UnifiedAtsUserOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedAtsUserOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedAtsUserOutput) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *UnifiedAtsUserOutput) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *UnifiedAtsUserOutput) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *UnifiedAtsUserOutput) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *UnifiedAtsUserOutput) GetAccessRole() *string {
	if o == nil {
		return nil
	}
	return o.AccessRole
}

func (o *UnifiedAtsUserOutput) GetRemoteCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteCreatedAt
}

func (o *UnifiedAtsUserOutput) GetRemoteModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteModifiedAt
}

func (o *UnifiedAtsUserOutput) GetFieldMappings() UnifiedAtsUserOutputFieldMappings {
	if o == nil {
		return UnifiedAtsUserOutputFieldMappings{}
	}
	return o.FieldMappings
}

func (o *UnifiedAtsUserOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedAtsUserOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedAtsUserOutput) GetRemoteData() UnifiedAtsUserOutputRemoteData {
	if o == nil {
		return UnifiedAtsUserOutputRemoteData{}
	}
	return o.RemoteData
}

func (o *UnifiedAtsUserOutput) GetCreatedAt() UnifiedAtsUserOutputCreatedAt {
	if o == nil {
		return UnifiedAtsUserOutputCreatedAt{}
	}
	return o.CreatedAt
}

func (o *UnifiedAtsUserOutput) GetModifiedAt() UnifiedAtsUserOutputModifiedAt {
	if o == nil {
		return UnifiedAtsUserOutputModifiedAt{}
	}
	return o.ModifiedAt
}