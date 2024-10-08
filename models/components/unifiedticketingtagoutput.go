// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

type UnifiedTicketingTagOutput struct {
	// The name of the tag
	Name *string `json:"name"`
	// The custom field mappings of the tag between the remote 3rd party & Panora
	FieldMappings map[string]any `json:"field_mappings,omitempty"`
	// The UUID of the tag
	ID *string `json:"id,omitempty"`
	// The remote ID of the tag in the context of the 3rd Party
	RemoteID *string `json:"remote_id,omitempty"`
	// The remote data of the tag in the context of the 3rd Party
	RemoteData map[string]any `json:"remote_data,omitempty"`
	// The created date of the tag
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The modified date of the tag
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
}

func (u UnifiedTicketingTagOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedTicketingTagOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedTicketingTagOutput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UnifiedTicketingTagOutput) GetFieldMappings() map[string]any {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}

func (o *UnifiedTicketingTagOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedTicketingTagOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedTicketingTagOutput) GetRemoteData() map[string]any {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *UnifiedTicketingTagOutput) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UnifiedTicketingTagOutput) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}
