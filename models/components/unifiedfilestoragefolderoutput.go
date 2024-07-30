// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UnifiedFilestorageFolderOutputFieldMappings struct {
}

type UnifiedFilestorageFolderOutputRemoteData struct {
}

type UnifiedFilestorageFolderOutputCreatedAt struct {
}

type UnifiedFilestorageFolderOutputModifiedAt struct {
}

type UnifiedFilestorageFolderOutput struct {
	// The name of the folder
	Name string `json:"name"`
	// The size of the folder
	Size string `json:"size"`
	// The url of the folder
	FolderURL string `json:"folder_url"`
	// The description of the folder
	Description string `json:"description"`
	// The UUID of the drive tied to the folder
	DriveID string `json:"drive_id"`
	// The UUID of the parent folder
	ParentFolderID string `json:"parent_folder_id"`
	// The UUID of the shared link tied to the folder
	SharedLink string `json:"shared_link"`
	// The UUID of the permission tied to the folder
	Permission    string                                      `json:"permission"`
	FieldMappings UnifiedFilestorageFolderOutputFieldMappings `json:"field_mappings"`
	// The UUID of the folder
	ID *string `json:"id,omitempty"`
	// The id of the folder in the context of the 3rd Party
	RemoteID   *string                                  `json:"remote_id,omitempty"`
	RemoteData UnifiedFilestorageFolderOutputRemoteData `json:"remote_data"`
	CreatedAt  UnifiedFilestorageFolderOutputCreatedAt  `json:"created_at"`
	ModifiedAt UnifiedFilestorageFolderOutputModifiedAt `json:"modified_at"`
}

func (o *UnifiedFilestorageFolderOutput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UnifiedFilestorageFolderOutput) GetSize() string {
	if o == nil {
		return ""
	}
	return o.Size
}

func (o *UnifiedFilestorageFolderOutput) GetFolderURL() string {
	if o == nil {
		return ""
	}
	return o.FolderURL
}

func (o *UnifiedFilestorageFolderOutput) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *UnifiedFilestorageFolderOutput) GetDriveID() string {
	if o == nil {
		return ""
	}
	return o.DriveID
}

func (o *UnifiedFilestorageFolderOutput) GetParentFolderID() string {
	if o == nil {
		return ""
	}
	return o.ParentFolderID
}

func (o *UnifiedFilestorageFolderOutput) GetSharedLink() string {
	if o == nil {
		return ""
	}
	return o.SharedLink
}

func (o *UnifiedFilestorageFolderOutput) GetPermission() string {
	if o == nil {
		return ""
	}
	return o.Permission
}

func (o *UnifiedFilestorageFolderOutput) GetFieldMappings() UnifiedFilestorageFolderOutputFieldMappings {
	if o == nil {
		return UnifiedFilestorageFolderOutputFieldMappings{}
	}
	return o.FieldMappings
}

func (o *UnifiedFilestorageFolderOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedFilestorageFolderOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedFilestorageFolderOutput) GetRemoteData() UnifiedFilestorageFolderOutputRemoteData {
	if o == nil {
		return UnifiedFilestorageFolderOutputRemoteData{}
	}
	return o.RemoteData
}

func (o *UnifiedFilestorageFolderOutput) GetCreatedAt() UnifiedFilestorageFolderOutputCreatedAt {
	if o == nil {
		return UnifiedFilestorageFolderOutputCreatedAt{}
	}
	return o.CreatedAt
}

func (o *UnifiedFilestorageFolderOutput) GetModifiedAt() UnifiedFilestorageFolderOutputModifiedAt {
	if o == nil {
		return UnifiedFilestorageFolderOutputModifiedAt{}
	}
	return o.ModifiedAt
}
