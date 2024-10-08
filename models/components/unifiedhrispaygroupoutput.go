// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

// UnifiedHrisPaygroupOutputFieldMappings - The custom field mappings of the object between the remote 3rd party & Panora
type UnifiedHrisPaygroupOutputFieldMappings struct {
}

// UnifiedHrisPaygroupOutputRemoteData - The remote data of the pay group in the context of the 3rd Party
type UnifiedHrisPaygroupOutputRemoteData struct {
}

type UnifiedHrisPaygroupOutput struct {
	// The name of the pay group
	PayGroupName *string `json:"pay_group_name,omitempty"`
	// The custom field mappings of the object between the remote 3rd party & Panora
	FieldMappings *UnifiedHrisPaygroupOutputFieldMappings `json:"field_mappings,omitempty"`
	// The UUID of the pay group record
	ID *string `json:"id,omitempty"`
	// The remote ID of the pay group in the context of the 3rd Party
	RemoteID *string `json:"remote_id,omitempty"`
	// The remote data of the pay group in the context of the 3rd Party
	RemoteData *UnifiedHrisPaygroupOutputRemoteData `json:"remote_data,omitempty"`
	// The date when the pay group was created in the 3rd party system
	RemoteCreatedAt *time.Time `json:"remote_created_at,omitempty"`
	// The created date of the pay group record
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The last modified date of the pay group record
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// Indicates if the pay group was deleted in the remote system
	RemoteWasDeleted *bool `json:"remote_was_deleted,omitempty"`
}

func (u UnifiedHrisPaygroupOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedHrisPaygroupOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedHrisPaygroupOutput) GetPayGroupName() *string {
	if o == nil {
		return nil
	}
	return o.PayGroupName
}

func (o *UnifiedHrisPaygroupOutput) GetFieldMappings() *UnifiedHrisPaygroupOutputFieldMappings {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}

func (o *UnifiedHrisPaygroupOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedHrisPaygroupOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedHrisPaygroupOutput) GetRemoteData() *UnifiedHrisPaygroupOutputRemoteData {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *UnifiedHrisPaygroupOutput) GetRemoteCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteCreatedAt
}

func (o *UnifiedHrisPaygroupOutput) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UnifiedHrisPaygroupOutput) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}

func (o *UnifiedHrisPaygroupOutput) GetRemoteWasDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteWasDeleted
}
