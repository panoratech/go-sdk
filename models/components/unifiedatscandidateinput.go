// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

type UnifiedAtsCandidateInputAttachmentsType string

const (
	UnifiedAtsCandidateInputAttachmentsTypeStr                        UnifiedAtsCandidateInputAttachmentsType = "str"
	UnifiedAtsCandidateInputAttachmentsTypeUnifiedAtsAttachmentOutput UnifiedAtsCandidateInputAttachmentsType = "UnifiedAtsAttachmentOutput"
)

type UnifiedAtsCandidateInputAttachments struct {
	Str                        *string
	UnifiedAtsAttachmentOutput *UnifiedAtsAttachmentOutput

	Type UnifiedAtsCandidateInputAttachmentsType
}

func CreateUnifiedAtsCandidateInputAttachmentsStr(str string) UnifiedAtsCandidateInputAttachments {
	typ := UnifiedAtsCandidateInputAttachmentsTypeStr

	return UnifiedAtsCandidateInputAttachments{
		Str:  &str,
		Type: typ,
	}
}

func CreateUnifiedAtsCandidateInputAttachmentsUnifiedAtsAttachmentOutput(unifiedAtsAttachmentOutput UnifiedAtsAttachmentOutput) UnifiedAtsCandidateInputAttachments {
	typ := UnifiedAtsCandidateInputAttachmentsTypeUnifiedAtsAttachmentOutput

	return UnifiedAtsCandidateInputAttachments{
		UnifiedAtsAttachmentOutput: &unifiedAtsAttachmentOutput,
		Type:                       typ,
	}
}

func (u *UnifiedAtsCandidateInputAttachments) UnmarshalJSON(data []byte) error {

	var unifiedAtsAttachmentOutput UnifiedAtsAttachmentOutput = UnifiedAtsAttachmentOutput{}
	if err := utils.UnmarshalJSON(data, &unifiedAtsAttachmentOutput, "", true, true); err == nil {
		u.UnifiedAtsAttachmentOutput = &unifiedAtsAttachmentOutput
		u.Type = UnifiedAtsCandidateInputAttachmentsTypeUnifiedAtsAttachmentOutput
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = UnifiedAtsCandidateInputAttachmentsTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UnifiedAtsCandidateInputAttachments", string(data))
}

func (u UnifiedAtsCandidateInputAttachments) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.UnifiedAtsAttachmentOutput != nil {
		return utils.MarshalJSON(u.UnifiedAtsAttachmentOutput, "", true)
	}

	return nil, errors.New("could not marshal union type UnifiedAtsCandidateInputAttachments: all fields are null")
}

type UnifiedAtsCandidateInputApplicationsType string

const (
	UnifiedAtsCandidateInputApplicationsTypeStr                         UnifiedAtsCandidateInputApplicationsType = "str"
	UnifiedAtsCandidateInputApplicationsTypeUnifiedAtsApplicationOutput UnifiedAtsCandidateInputApplicationsType = "UnifiedAtsApplicationOutput"
)

type UnifiedAtsCandidateInputApplications struct {
	Str                         *string
	UnifiedAtsApplicationOutput *UnifiedAtsApplicationOutput

	Type UnifiedAtsCandidateInputApplicationsType
}

func CreateUnifiedAtsCandidateInputApplicationsStr(str string) UnifiedAtsCandidateInputApplications {
	typ := UnifiedAtsCandidateInputApplicationsTypeStr

	return UnifiedAtsCandidateInputApplications{
		Str:  &str,
		Type: typ,
	}
}

func CreateUnifiedAtsCandidateInputApplicationsUnifiedAtsApplicationOutput(unifiedAtsApplicationOutput UnifiedAtsApplicationOutput) UnifiedAtsCandidateInputApplications {
	typ := UnifiedAtsCandidateInputApplicationsTypeUnifiedAtsApplicationOutput

	return UnifiedAtsCandidateInputApplications{
		UnifiedAtsApplicationOutput: &unifiedAtsApplicationOutput,
		Type:                        typ,
	}
}

func (u *UnifiedAtsCandidateInputApplications) UnmarshalJSON(data []byte) error {

	var unifiedAtsApplicationOutput UnifiedAtsApplicationOutput = UnifiedAtsApplicationOutput{}
	if err := utils.UnmarshalJSON(data, &unifiedAtsApplicationOutput, "", true, true); err == nil {
		u.UnifiedAtsApplicationOutput = &unifiedAtsApplicationOutput
		u.Type = UnifiedAtsCandidateInputApplicationsTypeUnifiedAtsApplicationOutput
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = UnifiedAtsCandidateInputApplicationsTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UnifiedAtsCandidateInputApplications", string(data))
}

func (u UnifiedAtsCandidateInputApplications) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.UnifiedAtsApplicationOutput != nil {
		return utils.MarshalJSON(u.UnifiedAtsApplicationOutput, "", true)
	}

	return nil, errors.New("could not marshal union type UnifiedAtsCandidateInputApplications: all fields are null")
}

type UnifiedAtsCandidateInputTagsType string

const (
	UnifiedAtsCandidateInputTagsTypeStr                 UnifiedAtsCandidateInputTagsType = "str"
	UnifiedAtsCandidateInputTagsTypeUnifiedAtsTagOutput UnifiedAtsCandidateInputTagsType = "UnifiedAtsTagOutput"
)

type UnifiedAtsCandidateInputTags struct {
	Str                 *string
	UnifiedAtsTagOutput *UnifiedAtsTagOutput

	Type UnifiedAtsCandidateInputTagsType
}

func CreateUnifiedAtsCandidateInputTagsStr(str string) UnifiedAtsCandidateInputTags {
	typ := UnifiedAtsCandidateInputTagsTypeStr

	return UnifiedAtsCandidateInputTags{
		Str:  &str,
		Type: typ,
	}
}

func CreateUnifiedAtsCandidateInputTagsUnifiedAtsTagOutput(unifiedAtsTagOutput UnifiedAtsTagOutput) UnifiedAtsCandidateInputTags {
	typ := UnifiedAtsCandidateInputTagsTypeUnifiedAtsTagOutput

	return UnifiedAtsCandidateInputTags{
		UnifiedAtsTagOutput: &unifiedAtsTagOutput,
		Type:                typ,
	}
}

func (u *UnifiedAtsCandidateInputTags) UnmarshalJSON(data []byte) error {

	var unifiedAtsTagOutput UnifiedAtsTagOutput = UnifiedAtsTagOutput{}
	if err := utils.UnmarshalJSON(data, &unifiedAtsTagOutput, "", true, true); err == nil {
		u.UnifiedAtsTagOutput = &unifiedAtsTagOutput
		u.Type = UnifiedAtsCandidateInputTagsTypeUnifiedAtsTagOutput
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = UnifiedAtsCandidateInputTagsTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UnifiedAtsCandidateInputTags", string(data))
}

func (u UnifiedAtsCandidateInputTags) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.UnifiedAtsTagOutput != nil {
		return utils.MarshalJSON(u.UnifiedAtsTagOutput, "", true)
	}

	return nil, errors.New("could not marshal union type UnifiedAtsCandidateInputTags: all fields are null")
}

type UnifiedAtsCandidateInput struct {
	// The first name of the candidate
	FirstName *string `json:"first_name,omitempty"`
	// The last name of the candidate
	LastName *string `json:"last_name,omitempty"`
	// The company of the candidate
	Company *string `json:"company,omitempty"`
	// The title of the candidate
	Title *string `json:"title,omitempty"`
	// The locations of the candidate
	Locations *string `json:"locations,omitempty"`
	// Whether the candidate is private
	IsPrivate *bool `json:"is_private,omitempty"`
	// Whether the candidate is reachable by email
	EmailReachable *bool `json:"email_reachable,omitempty"`
	// The remote creation date of the candidate
	RemoteCreatedAt *time.Time `json:"remote_created_at,omitempty"`
	// The remote modification date of the candidate
	RemoteModifiedAt *time.Time `json:"remote_modified_at,omitempty"`
	// The last interaction date with the candidate
	LastInteractionAt *time.Time `json:"last_interaction_at,omitempty"`
	// The attachments UUIDs of the candidate
	Attachments []UnifiedAtsCandidateInputAttachments `json:"attachments,omitempty"`
	// The applications UUIDs of the candidate
	Applications []UnifiedAtsCandidateInputApplications `json:"applications,omitempty"`
	// The tags of the candidate
	Tags []UnifiedAtsCandidateInputTags `json:"tags,omitempty"`
	// The urls of the candidate, possible values for Url type are WEBSITE, BLOG, LINKEDIN, GITHUB, or OTHER
	Urls []URL `json:"urls,omitempty"`
	// The phone numbers of the candidate
	PhoneNumbers []Phone `json:"phone_numbers,omitempty"`
	// The email addresses of the candidate
	EmailAddresses []Email `json:"email_addresses,omitempty"`
	// The custom field mappings of the object between the remote 3rd party & Panora
	FieldMappings map[string]any `json:"field_mappings,omitempty"`
}

func (u UnifiedAtsCandidateInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedAtsCandidateInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedAtsCandidateInput) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *UnifiedAtsCandidateInput) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *UnifiedAtsCandidateInput) GetCompany() *string {
	if o == nil {
		return nil
	}
	return o.Company
}

func (o *UnifiedAtsCandidateInput) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *UnifiedAtsCandidateInput) GetLocations() *string {
	if o == nil {
		return nil
	}
	return o.Locations
}

func (o *UnifiedAtsCandidateInput) GetIsPrivate() *bool {
	if o == nil {
		return nil
	}
	return o.IsPrivate
}

func (o *UnifiedAtsCandidateInput) GetEmailReachable() *bool {
	if o == nil {
		return nil
	}
	return o.EmailReachable
}

func (o *UnifiedAtsCandidateInput) GetRemoteCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteCreatedAt
}

func (o *UnifiedAtsCandidateInput) GetRemoteModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteModifiedAt
}

func (o *UnifiedAtsCandidateInput) GetLastInteractionAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastInteractionAt
}

func (o *UnifiedAtsCandidateInput) GetAttachments() []UnifiedAtsCandidateInputAttachments {
	if o == nil {
		return nil
	}
	return o.Attachments
}

func (o *UnifiedAtsCandidateInput) GetApplications() []UnifiedAtsCandidateInputApplications {
	if o == nil {
		return nil
	}
	return o.Applications
}

func (o *UnifiedAtsCandidateInput) GetTags() []UnifiedAtsCandidateInputTags {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *UnifiedAtsCandidateInput) GetUrls() []URL {
	if o == nil {
		return nil
	}
	return o.Urls
}

func (o *UnifiedAtsCandidateInput) GetPhoneNumbers() []Phone {
	if o == nil {
		return nil
	}
	return o.PhoneNumbers
}

func (o *UnifiedAtsCandidateInput) GetEmailAddresses() []Email {
	if o == nil {
		return nil
	}
	return o.EmailAddresses
}

func (o *UnifiedAtsCandidateInput) GetFieldMappings() map[string]any {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}
