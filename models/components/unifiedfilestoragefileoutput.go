// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

type UnifiedFilestorageFileOutput struct {
	// The name of the file
	Name *string `json:"name"`
	// The url of the file
	FileURL *string `json:"file_url"`
	// The mime type of the file
	MimeType *string `json:"mime_type"`
	// The size of the file
	Size *string `json:"size"`
	// The UUID of the folder tied to the file
	FolderID *string `json:"folder_id"`
	// The UUID of the permission tied to the file
	Permission *string `json:"permission"`
	// The UUID of the shared link tied to the file
	SharedLink *string `json:"shared_link"`
	// The custom field mappings of the object between the remote 3rd party & Panora
	FieldMappings map[string]any `json:"field_mappings,omitempty"`
	// The UUID of the file
	ID *string `json:"id,omitempty"`
	// The id of the file in the context of the 3rd Party
	RemoteID *string `json:"remote_id,omitempty"`
	// The remote data of the file in the context of the 3rd Party
	RemoteData map[string]any `json:"remote_data,omitempty"`
	// The created date of the object
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The modified date of the object
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
}

func (u UnifiedFilestorageFileOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedFilestorageFileOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedFilestorageFileOutput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UnifiedFilestorageFileOutput) GetFileURL() *string {
	if o == nil {
		return nil
	}
	return o.FileURL
}

func (o *UnifiedFilestorageFileOutput) GetMimeType() *string {
	if o == nil {
		return nil
	}
	return o.MimeType
}

func (o *UnifiedFilestorageFileOutput) GetSize() *string {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *UnifiedFilestorageFileOutput) GetFolderID() *string {
	if o == nil {
		return nil
	}
	return o.FolderID
}

func (o *UnifiedFilestorageFileOutput) GetPermission() *string {
	if o == nil {
		return nil
	}
	return o.Permission
}

func (o *UnifiedFilestorageFileOutput) GetSharedLink() *string {
	if o == nil {
		return nil
	}
	return o.SharedLink
}

func (o *UnifiedFilestorageFileOutput) GetFieldMappings() map[string]any {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}

func (o *UnifiedFilestorageFileOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedFilestorageFileOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedFilestorageFileOutput) GetRemoteData() map[string]any {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *UnifiedFilestorageFileOutput) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UnifiedFilestorageFileOutput) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}
