// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UnifiedTicketingAccountOutputFieldMappings struct {
}

type UnifiedTicketingAccountOutputRemoteData struct {
}

type UnifiedTicketingAccountOutputCreatedAt struct {
}

type UnifiedTicketingAccountOutputModifiedAt struct {
}

type UnifiedTicketingAccountOutput struct {
	// The name of the account
	Name string `json:"name"`
	// The domains of the account
	Domains       []string                                   `json:"domains,omitempty"`
	FieldMappings UnifiedTicketingAccountOutputFieldMappings `json:"field_mappings"`
	// The UUID of the account
	ID *string `json:"id,omitempty"`
	// The id of the account in the context of the 3rd Party
	RemoteID   *string                                 `json:"remote_id,omitempty"`
	RemoteData UnifiedTicketingAccountOutputRemoteData `json:"remote_data"`
	CreatedAt  UnifiedTicketingAccountOutputCreatedAt  `json:"created_at"`
	ModifiedAt UnifiedTicketingAccountOutputModifiedAt `json:"modified_at"`
}

func (o *UnifiedTicketingAccountOutput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UnifiedTicketingAccountOutput) GetDomains() []string {
	if o == nil {
		return nil
	}
	return o.Domains
}

func (o *UnifiedTicketingAccountOutput) GetFieldMappings() UnifiedTicketingAccountOutputFieldMappings {
	if o == nil {
		return UnifiedTicketingAccountOutputFieldMappings{}
	}
	return o.FieldMappings
}

func (o *UnifiedTicketingAccountOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedTicketingAccountOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedTicketingAccountOutput) GetRemoteData() UnifiedTicketingAccountOutputRemoteData {
	if o == nil {
		return UnifiedTicketingAccountOutputRemoteData{}
	}
	return o.RemoteData
}

func (o *UnifiedTicketingAccountOutput) GetCreatedAt() UnifiedTicketingAccountOutputCreatedAt {
	if o == nil {
		return UnifiedTicketingAccountOutputCreatedAt{}
	}
	return o.CreatedAt
}

func (o *UnifiedTicketingAccountOutput) GetModifiedAt() UnifiedTicketingAccountOutputModifiedAt {
	if o == nil {
		return UnifiedTicketingAccountOutputModifiedAt{}
	}
	return o.ModifiedAt
}
