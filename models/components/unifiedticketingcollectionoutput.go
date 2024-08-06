// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

// CollectionType - The type of the collection. Authorized values are either PROJECT or LIST
type CollectionType string

const (
	CollectionTypeProject CollectionType = "PROJECT"
	CollectionTypeList    CollectionType = "LIST"
)

func (e CollectionType) ToPointer() *CollectionType {
	return &e
}
func (e *CollectionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PROJECT":
		fallthrough
	case "LIST":
		*e = CollectionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CollectionType: %v", v)
	}
}

type UnifiedTicketingCollectionOutput struct {
	// The name of the collection
	Name *string `json:"name"`
	// The description of the collection
	Description *string `json:"description,omitempty"`
	// The type of the collection. Authorized values are either PROJECT or LIST
	CollectionType *CollectionType `json:"collection_type,omitempty"`
	// The UUID of the collection
	ID *string `json:"id,omitempty"`
	// The id of the collection in the context of the 3rd Party
	RemoteID *string `json:"remote_id,omitempty"`
	// The remote data of the collection in the context of the 3rd Party
	RemoteData map[string]any `json:"remote_data,omitempty"`
	// The created date of the object
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The modified date of the object
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
}

func (u UnifiedTicketingCollectionOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedTicketingCollectionOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedTicketingCollectionOutput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UnifiedTicketingCollectionOutput) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UnifiedTicketingCollectionOutput) GetCollectionType() *CollectionType {
	if o == nil {
		return nil
	}
	return o.CollectionType
}

func (o *UnifiedTicketingCollectionOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedTicketingCollectionOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedTicketingCollectionOutput) GetRemoteData() map[string]any {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *UnifiedTicketingCollectionOutput) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UnifiedTicketingCollectionOutput) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}
