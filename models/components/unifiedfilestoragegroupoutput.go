// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UnifiedFilestorageGroupOutputFieldMappings struct {
}

type UnifiedFilestorageGroupOutputRemoteData struct {
}

type UnifiedFilestorageGroupOutputCreatedAt struct {
}

type UnifiedFilestorageGroupOutputModifiedAt struct {
}

type UnifiedFilestorageGroupOutput struct {
	// The name of the group
	Name string `json:"name"`
	// Uuids of users of the group
	Users []string `json:"users"`
	// Indicates whether or not this object has been deleted in the third party platform.
	RemoteWasDeleted bool                                       `json:"remote_was_deleted"`
	FieldMappings    UnifiedFilestorageGroupOutputFieldMappings `json:"field_mappings"`
	// The UUID of the group
	ID *string `json:"id,omitempty"`
	// The id of the group in the context of the 3rd Party
	RemoteID   *string                                 `json:"remote_id,omitempty"`
	RemoteData UnifiedFilestorageGroupOutputRemoteData `json:"remote_data"`
	CreatedAt  UnifiedFilestorageGroupOutputCreatedAt  `json:"created_at"`
	ModifiedAt UnifiedFilestorageGroupOutputModifiedAt `json:"modified_at"`
}

func (o *UnifiedFilestorageGroupOutput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UnifiedFilestorageGroupOutput) GetUsers() []string {
	if o == nil {
		return []string{}
	}
	return o.Users
}

func (o *UnifiedFilestorageGroupOutput) GetRemoteWasDeleted() bool {
	if o == nil {
		return false
	}
	return o.RemoteWasDeleted
}

func (o *UnifiedFilestorageGroupOutput) GetFieldMappings() UnifiedFilestorageGroupOutputFieldMappings {
	if o == nil {
		return UnifiedFilestorageGroupOutputFieldMappings{}
	}
	return o.FieldMappings
}

func (o *UnifiedFilestorageGroupOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedFilestorageGroupOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedFilestorageGroupOutput) GetRemoteData() UnifiedFilestorageGroupOutputRemoteData {
	if o == nil {
		return UnifiedFilestorageGroupOutputRemoteData{}
	}
	return o.RemoteData
}

func (o *UnifiedFilestorageGroupOutput) GetCreatedAt() UnifiedFilestorageGroupOutputCreatedAt {
	if o == nil {
		return UnifiedFilestorageGroupOutputCreatedAt{}
	}
	return o.CreatedAt
}

func (o *UnifiedFilestorageGroupOutput) GetModifiedAt() UnifiedFilestorageGroupOutputModifiedAt {
	if o == nil {
		return UnifiedFilestorageGroupOutputModifiedAt{}
	}
	return o.ModifiedAt
}