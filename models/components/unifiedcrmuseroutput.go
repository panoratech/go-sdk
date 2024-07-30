// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UnifiedCrmUserOutputFieldMappings struct {
}

type UnifiedCrmUserOutputRemoteData struct {
}

type UnifiedCrmUserOutputCreatedAt struct {
}

type UnifiedCrmUserOutputModifiedAt struct {
}

type UnifiedCrmUserOutput struct {
	// The name of the user
	Name string `json:"name"`
	// The email of the user
	Email         string                            `json:"email"`
	FieldMappings UnifiedCrmUserOutputFieldMappings `json:"field_mappings"`
	// The UUID of the user
	ID *string `json:"id,omitempty"`
	// The id of the user in the context of the Crm 3rd Party
	RemoteID   *string                        `json:"remote_id,omitempty"`
	RemoteData UnifiedCrmUserOutputRemoteData `json:"remote_data"`
	CreatedAt  UnifiedCrmUserOutputCreatedAt  `json:"created_at"`
	ModifiedAt UnifiedCrmUserOutputModifiedAt `json:"modified_at"`
}

func (o *UnifiedCrmUserOutput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UnifiedCrmUserOutput) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *UnifiedCrmUserOutput) GetFieldMappings() UnifiedCrmUserOutputFieldMappings {
	if o == nil {
		return UnifiedCrmUserOutputFieldMappings{}
	}
	return o.FieldMappings
}

func (o *UnifiedCrmUserOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedCrmUserOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedCrmUserOutput) GetRemoteData() UnifiedCrmUserOutputRemoteData {
	if o == nil {
		return UnifiedCrmUserOutputRemoteData{}
	}
	return o.RemoteData
}

func (o *UnifiedCrmUserOutput) GetCreatedAt() UnifiedCrmUserOutputCreatedAt {
	if o == nil {
		return UnifiedCrmUserOutputCreatedAt{}
	}
	return o.CreatedAt
}

func (o *UnifiedCrmUserOutput) GetModifiedAt() UnifiedCrmUserOutputModifiedAt {
	if o == nil {
		return UnifiedCrmUserOutputModifiedAt{}
	}
	return o.ModifiedAt
}