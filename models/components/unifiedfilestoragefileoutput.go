// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UnifiedFilestorageFileOutputFieldMappings struct {
}

type UnifiedFilestorageFileOutputRemoteData struct {
}

type UnifiedFilestorageFileOutputCreatedAt struct {
}

type UnifiedFilestorageFileOutputModifiedAt struct {
}

type UnifiedFilestorageFileOutput struct {
	// The name of the file
	Name string `json:"name"`
	// The url of the file
	FileURL string `json:"file_url"`
	// The mime type of the file
	MimeType string `json:"mime_type"`
	// The size of the file
	Size string `json:"size"`
	// The UUID of the folder tied to the file
	FolderID string `json:"folder_id"`
	// The UUID of the permission tied to the file
	Permission string `json:"permission"`
	// The UUID of the shared link tied to the file
	SharedLink    string                                    `json:"shared_link"`
	FieldMappings UnifiedFilestorageFileOutputFieldMappings `json:"field_mappings"`
	// The UUID of the file
	ID *string `json:"id,omitempty"`
	// The id of the file in the context of the 3rd Party
	RemoteID   *string                                `json:"remote_id,omitempty"`
	RemoteData UnifiedFilestorageFileOutputRemoteData `json:"remote_data"`
	CreatedAt  UnifiedFilestorageFileOutputCreatedAt  `json:"created_at"`
	ModifiedAt UnifiedFilestorageFileOutputModifiedAt `json:"modified_at"`
}

func (o *UnifiedFilestorageFileOutput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UnifiedFilestorageFileOutput) GetFileURL() string {
	if o == nil {
		return ""
	}
	return o.FileURL
}

func (o *UnifiedFilestorageFileOutput) GetMimeType() string {
	if o == nil {
		return ""
	}
	return o.MimeType
}

func (o *UnifiedFilestorageFileOutput) GetSize() string {
	if o == nil {
		return ""
	}
	return o.Size
}

func (o *UnifiedFilestorageFileOutput) GetFolderID() string {
	if o == nil {
		return ""
	}
	return o.FolderID
}

func (o *UnifiedFilestorageFileOutput) GetPermission() string {
	if o == nil {
		return ""
	}
	return o.Permission
}

func (o *UnifiedFilestorageFileOutput) GetSharedLink() string {
	if o == nil {
		return ""
	}
	return o.SharedLink
}

func (o *UnifiedFilestorageFileOutput) GetFieldMappings() UnifiedFilestorageFileOutputFieldMappings {
	if o == nil {
		return UnifiedFilestorageFileOutputFieldMappings{}
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

func (o *UnifiedFilestorageFileOutput) GetRemoteData() UnifiedFilestorageFileOutputRemoteData {
	if o == nil {
		return UnifiedFilestorageFileOutputRemoteData{}
	}
	return o.RemoteData
}

func (o *UnifiedFilestorageFileOutput) GetCreatedAt() UnifiedFilestorageFileOutputCreatedAt {
	if o == nil {
		return UnifiedFilestorageFileOutputCreatedAt{}
	}
	return o.CreatedAt
}

func (o *UnifiedFilestorageFileOutput) GetModifiedAt() UnifiedFilestorageFileOutputModifiedAt {
	if o == nil {
		return UnifiedFilestorageFileOutputModifiedAt{}
	}
	return o.ModifiedAt
}
