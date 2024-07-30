// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type MapFieldToProviderDto struct {
	AttributeID         string `json:"attributeId"`
	SourceCustomFieldID string `json:"source_custom_field_id"`
	SourceProvider      string `json:"source_provider"`
	LinkedUserID        string `json:"linked_user_id"`
}

func (o *MapFieldToProviderDto) GetAttributeID() string {
	if o == nil {
		return ""
	}
	return o.AttributeID
}

func (o *MapFieldToProviderDto) GetSourceCustomFieldID() string {
	if o == nil {
		return ""
	}
	return o.SourceCustomFieldID
}

func (o *MapFieldToProviderDto) GetSourceProvider() string {
	if o == nil {
		return ""
	}
	return o.SourceProvider
}

func (o *MapFieldToProviderDto) GetLinkedUserID() string {
	if o == nil {
		return ""
	}
	return o.LinkedUserID
}
