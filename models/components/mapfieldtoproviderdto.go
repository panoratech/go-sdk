// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type MapFieldToProviderDto struct {
	// The attribute ID
	AttributeID *string `json:"attributeId"`
	// The source custom field ID
	SourceCustomFieldID *string `json:"source_custom_field_id"`
	// The source provider
	SourceProvider *string `json:"source_provider"`
	// The linked user ID
	LinkedUserID *string `json:"linked_user_id"`
}

func (o *MapFieldToProviderDto) GetAttributeID() *string {
	if o == nil {
		return nil
	}
	return o.AttributeID
}

func (o *MapFieldToProviderDto) GetSourceCustomFieldID() *string {
	if o == nil {
		return nil
	}
	return o.SourceCustomFieldID
}

func (o *MapFieldToProviderDto) GetSourceProvider() *string {
	if o == nil {
		return nil
	}
	return o.SourceProvider
}

func (o *MapFieldToProviderDto) GetLinkedUserID() *string {
	if o == nil {
		return nil
	}
	return o.LinkedUserID
}
