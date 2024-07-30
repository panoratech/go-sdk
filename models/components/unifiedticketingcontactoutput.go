// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UnifiedTicketingContactOutputFieldMappings struct {
}

type UnifiedTicketingContactOutputRemoteData struct {
}

type UnifiedTicketingContactOutputCreatedAt struct {
}

type UnifiedTicketingContactOutputModifiedAt struct {
}

type UnifiedTicketingContactOutput struct {
	// The name of the contact
	Name string `json:"name"`
	// The email address of the contact
	EmailAddress string `json:"email_address"`
	// The phone number of the contact
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The details of the contact
	Details       *string                                    `json:"details,omitempty"`
	FieldMappings UnifiedTicketingContactOutputFieldMappings `json:"field_mappings"`
	// The UUID of the contact
	ID *string `json:"id,omitempty"`
	// The id of the contact in the context of the 3rd Party
	RemoteID   *string                                 `json:"remote_id,omitempty"`
	RemoteData UnifiedTicketingContactOutputRemoteData `json:"remote_data"`
	CreatedAt  UnifiedTicketingContactOutputCreatedAt  `json:"created_at"`
	ModifiedAt UnifiedTicketingContactOutputModifiedAt `json:"modified_at"`
}

func (o *UnifiedTicketingContactOutput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UnifiedTicketingContactOutput) GetEmailAddress() string {
	if o == nil {
		return ""
	}
	return o.EmailAddress
}

func (o *UnifiedTicketingContactOutput) GetPhoneNumber() *string {
	if o == nil {
		return nil
	}
	return o.PhoneNumber
}

func (o *UnifiedTicketingContactOutput) GetDetails() *string {
	if o == nil {
		return nil
	}
	return o.Details
}

func (o *UnifiedTicketingContactOutput) GetFieldMappings() UnifiedTicketingContactOutputFieldMappings {
	if o == nil {
		return UnifiedTicketingContactOutputFieldMappings{}
	}
	return o.FieldMappings
}

func (o *UnifiedTicketingContactOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedTicketingContactOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedTicketingContactOutput) GetRemoteData() UnifiedTicketingContactOutputRemoteData {
	if o == nil {
		return UnifiedTicketingContactOutputRemoteData{}
	}
	return o.RemoteData
}

func (o *UnifiedTicketingContactOutput) GetCreatedAt() UnifiedTicketingContactOutputCreatedAt {
	if o == nil {
		return UnifiedTicketingContactOutputCreatedAt{}
	}
	return o.CreatedAt
}

func (o *UnifiedTicketingContactOutput) GetModifiedAt() UnifiedTicketingContactOutputModifiedAt {
	if o == nil {
		return UnifiedTicketingContactOutputModifiedAt{}
	}
	return o.ModifiedAt
}
