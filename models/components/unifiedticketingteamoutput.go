// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UnifiedTicketingTeamOutputFieldMappings struct {
}

type UnifiedTicketingTeamOutputRemoteData struct {
}

type UnifiedTicketingTeamOutputCreatedAt struct {
}

type UnifiedTicketingTeamOutputModifiedAt struct {
}

type UnifiedTicketingTeamOutput struct {
	// The name of the team
	Name string `json:"name"`
	// The description of the team
	Description   *string                                 `json:"description,omitempty"`
	FieldMappings UnifiedTicketingTeamOutputFieldMappings `json:"field_mappings"`
	// The UUID of the team
	ID *string `json:"id,omitempty"`
	// The id of the team in the context of the 3rd Party
	RemoteID   *string                              `json:"remote_id,omitempty"`
	RemoteData UnifiedTicketingTeamOutputRemoteData `json:"remote_data"`
	CreatedAt  UnifiedTicketingTeamOutputCreatedAt  `json:"created_at"`
	ModifiedAt UnifiedTicketingTeamOutputModifiedAt `json:"modified_at"`
}

func (o *UnifiedTicketingTeamOutput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UnifiedTicketingTeamOutput) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UnifiedTicketingTeamOutput) GetFieldMappings() UnifiedTicketingTeamOutputFieldMappings {
	if o == nil {
		return UnifiedTicketingTeamOutputFieldMappings{}
	}
	return o.FieldMappings
}

func (o *UnifiedTicketingTeamOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedTicketingTeamOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedTicketingTeamOutput) GetRemoteData() UnifiedTicketingTeamOutputRemoteData {
	if o == nil {
		return UnifiedTicketingTeamOutputRemoteData{}
	}
	return o.RemoteData
}

func (o *UnifiedTicketingTeamOutput) GetCreatedAt() UnifiedTicketingTeamOutputCreatedAt {
	if o == nil {
		return UnifiedTicketingTeamOutputCreatedAt{}
	}
	return o.CreatedAt
}

func (o *UnifiedTicketingTeamOutput) GetModifiedAt() UnifiedTicketingTeamOutputModifiedAt {
	if o == nil {
		return UnifiedTicketingTeamOutputModifiedAt{}
	}
	return o.ModifiedAt
}
