// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

// Status - The status of the ticket. Authorized values are OPEN or CLOSED.
type Status string

const (
	StatusOpen   Status = "OPEN"
	StatusClosed Status = "CLOSED"
)

func (e Status) ToPointer() *Status {
	return &e
}
func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OPEN":
		fallthrough
	case "CLOSED":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

// Type - The type of the ticket. Authorized values are PROBLEM, QUESTION, or TASK
type Type string

const (
	TypeBug     Type = "BUG"
	TypeSubtask Type = "SUBTASK"
	TypeTask    Type = "TASK"
	TypeToDo    Type = "TO-DO"
)

func (e Type) ToPointer() *Type {
	return &e
}
func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BUG":
		fallthrough
	case "SUBTASK":
		fallthrough
	case "TASK":
		fallthrough
	case "TO-DO":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type CollectionsType string

const (
	CollectionsTypeStr                              CollectionsType = "str"
	CollectionsTypeUnifiedTicketingCollectionOutput CollectionsType = "UnifiedTicketingCollectionOutput"
)

type Collections struct {
	Str                              *string
	UnifiedTicketingCollectionOutput *UnifiedTicketingCollectionOutput

	Type CollectionsType
}

func CreateCollectionsStr(str string) Collections {
	typ := CollectionsTypeStr

	return Collections{
		Str:  &str,
		Type: typ,
	}
}

func CreateCollectionsUnifiedTicketingCollectionOutput(unifiedTicketingCollectionOutput UnifiedTicketingCollectionOutput) Collections {
	typ := CollectionsTypeUnifiedTicketingCollectionOutput

	return Collections{
		UnifiedTicketingCollectionOutput: &unifiedTicketingCollectionOutput,
		Type:                             typ,
	}
}

func (u *Collections) UnmarshalJSON(data []byte) error {

	var unifiedTicketingCollectionOutput UnifiedTicketingCollectionOutput = UnifiedTicketingCollectionOutput{}
	if err := utils.UnmarshalJSON(data, &unifiedTicketingCollectionOutput, "", true, true); err == nil {
		u.UnifiedTicketingCollectionOutput = &unifiedTicketingCollectionOutput
		u.Type = CollectionsTypeUnifiedTicketingCollectionOutput
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CollectionsTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Collections", string(data))
}

func (u Collections) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.UnifiedTicketingCollectionOutput != nil {
		return utils.MarshalJSON(u.UnifiedTicketingCollectionOutput, "", true)
	}

	return nil, errors.New("could not marshal union type Collections: all fields are null")
}

type TagsType string

const (
	TagsTypeStr                       TagsType = "str"
	TagsTypeUnifiedTicketingTagOutput TagsType = "UnifiedTicketingTagOutput"
)

type Tags struct {
	Str                       *string
	UnifiedTicketingTagOutput *UnifiedTicketingTagOutput

	Type TagsType
}

func CreateTagsStr(str string) Tags {
	typ := TagsTypeStr

	return Tags{
		Str:  &str,
		Type: typ,
	}
}

func CreateTagsUnifiedTicketingTagOutput(unifiedTicketingTagOutput UnifiedTicketingTagOutput) Tags {
	typ := TagsTypeUnifiedTicketingTagOutput

	return Tags{
		UnifiedTicketingTagOutput: &unifiedTicketingTagOutput,
		Type:                      typ,
	}
}

func (u *Tags) UnmarshalJSON(data []byte) error {

	var unifiedTicketingTagOutput UnifiedTicketingTagOutput = UnifiedTicketingTagOutput{}
	if err := utils.UnmarshalJSON(data, &unifiedTicketingTagOutput, "", true, true); err == nil {
		u.UnifiedTicketingTagOutput = &unifiedTicketingTagOutput
		u.Type = TagsTypeUnifiedTicketingTagOutput
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = TagsTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Tags", string(data))
}

func (u Tags) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.UnifiedTicketingTagOutput != nil {
		return utils.MarshalJSON(u.UnifiedTicketingTagOutput, "", true)
	}

	return nil, errors.New("could not marshal union type Tags: all fields are null")
}

// Priority - The priority of the ticket. Authorized values are HIGH, MEDIUM or LOW.
type Priority string

const (
	PriorityHigh   Priority = "HIGH"
	PriorityMedium Priority = "MEDIUM"
	PriorityLow    Priority = "LOW"
)

func (e Priority) ToPointer() *Priority {
	return &e
}
func (e *Priority) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HIGH":
		fallthrough
	case "MEDIUM":
		fallthrough
	case "LOW":
		*e = Priority(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Priority: %v", v)
	}
}

// UnifiedTicketingTicketOutputCreatorType - The creator type of the comment. Authorized values are either USER or CONTACT
type UnifiedTicketingTicketOutputCreatorType string

const (
	UnifiedTicketingTicketOutputCreatorTypeUser    UnifiedTicketingTicketOutputCreatorType = "USER"
	UnifiedTicketingTicketOutputCreatorTypeContact UnifiedTicketingTicketOutputCreatorType = "CONTACT"
)

func (e UnifiedTicketingTicketOutputCreatorType) ToPointer() *UnifiedTicketingTicketOutputCreatorType {
	return &e
}
func (e *UnifiedTicketingTicketOutputCreatorType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "USER":
		fallthrough
	case "CONTACT":
		*e = UnifiedTicketingTicketOutputCreatorType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UnifiedTicketingTicketOutputCreatorType: %v", v)
	}
}

type UnifiedTicketingTicketOutputAttachmentsType string

const (
	UnifiedTicketingTicketOutputAttachmentsTypeStr                              UnifiedTicketingTicketOutputAttachmentsType = "str"
	UnifiedTicketingTicketOutputAttachmentsTypeUnifiedTicketingAttachmentOutput UnifiedTicketingTicketOutputAttachmentsType = "UnifiedTicketingAttachmentOutput"
)

type UnifiedTicketingTicketOutputAttachments struct {
	Str                              *string
	UnifiedTicketingAttachmentOutput *UnifiedTicketingAttachmentOutput

	Type UnifiedTicketingTicketOutputAttachmentsType
}

func CreateUnifiedTicketingTicketOutputAttachmentsStr(str string) UnifiedTicketingTicketOutputAttachments {
	typ := UnifiedTicketingTicketOutputAttachmentsTypeStr

	return UnifiedTicketingTicketOutputAttachments{
		Str:  &str,
		Type: typ,
	}
}

func CreateUnifiedTicketingTicketOutputAttachmentsUnifiedTicketingAttachmentOutput(unifiedTicketingAttachmentOutput UnifiedTicketingAttachmentOutput) UnifiedTicketingTicketOutputAttachments {
	typ := UnifiedTicketingTicketOutputAttachmentsTypeUnifiedTicketingAttachmentOutput

	return UnifiedTicketingTicketOutputAttachments{
		UnifiedTicketingAttachmentOutput: &unifiedTicketingAttachmentOutput,
		Type:                             typ,
	}
}

func (u *UnifiedTicketingTicketOutputAttachments) UnmarshalJSON(data []byte) error {

	var unifiedTicketingAttachmentOutput UnifiedTicketingAttachmentOutput = UnifiedTicketingAttachmentOutput{}
	if err := utils.UnmarshalJSON(data, &unifiedTicketingAttachmentOutput, "", true, true); err == nil {
		u.UnifiedTicketingAttachmentOutput = &unifiedTicketingAttachmentOutput
		u.Type = UnifiedTicketingTicketOutputAttachmentsTypeUnifiedTicketingAttachmentOutput
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = UnifiedTicketingTicketOutputAttachmentsTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UnifiedTicketingTicketOutputAttachments", string(data))
}

func (u UnifiedTicketingTicketOutputAttachments) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.UnifiedTicketingAttachmentOutput != nil {
		return utils.MarshalJSON(u.UnifiedTicketingAttachmentOutput, "", true)
	}

	return nil, errors.New("could not marshal union type UnifiedTicketingTicketOutputAttachments: all fields are null")
}

// Comment - The comment of the ticket
type Comment struct {
	// The body of the comment
	Body *string `json:"body"`
	// The html body of the comment
	HTMLBody *string `json:"html_body,omitempty"`
	// The public status of the comment
	IsPrivate *bool `json:"is_private,omitempty"`
	// The creator type of the comment. Authorized values are either USER or CONTACT
	CreatorType *UnifiedTicketingTicketOutputCreatorType `json:"creator_type,omitempty"`
	// The UUID of the ticket the comment is tied to
	TicketID *string `json:"ticket_id,omitempty"`
	// The UUID of the contact which the comment belongs to (if no user_id specified)
	ContactID *string `json:"contact_id,omitempty"`
	// The UUID of the user which the comment belongs to (if no contact_id specified)
	UserID *string `json:"user_id,omitempty"`
	// The attachements UUIDs tied to the comment
	Attachments []UnifiedTicketingTicketOutputAttachments `json:"attachments,omitempty"`
}

func (o *Comment) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *Comment) GetHTMLBody() *string {
	if o == nil {
		return nil
	}
	return o.HTMLBody
}

func (o *Comment) GetIsPrivate() *bool {
	if o == nil {
		return nil
	}
	return o.IsPrivate
}

func (o *Comment) GetCreatorType() *UnifiedTicketingTicketOutputCreatorType {
	if o == nil {
		return nil
	}
	return o.CreatorType
}

func (o *Comment) GetTicketID() *string {
	if o == nil {
		return nil
	}
	return o.TicketID
}

func (o *Comment) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *Comment) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *Comment) GetAttachments() []UnifiedTicketingTicketOutputAttachments {
	if o == nil {
		return nil
	}
	return o.Attachments
}

type AttachmentsType string

const (
	AttachmentsTypeStr                             AttachmentsType = "str"
	AttachmentsTypeUnifiedTicketingAttachmentInput AttachmentsType = "UnifiedTicketingAttachmentInput"
)

type Attachments struct {
	Str                             *string
	UnifiedTicketingAttachmentInput *UnifiedTicketingAttachmentInput

	Type AttachmentsType
}

func CreateAttachmentsStr(str string) Attachments {
	typ := AttachmentsTypeStr

	return Attachments{
		Str:  &str,
		Type: typ,
	}
}

func CreateAttachmentsUnifiedTicketingAttachmentInput(unifiedTicketingAttachmentInput UnifiedTicketingAttachmentInput) Attachments {
	typ := AttachmentsTypeUnifiedTicketingAttachmentInput

	return Attachments{
		UnifiedTicketingAttachmentInput: &unifiedTicketingAttachmentInput,
		Type:                            typ,
	}
}

func (u *Attachments) UnmarshalJSON(data []byte) error {

	var unifiedTicketingAttachmentInput UnifiedTicketingAttachmentInput = UnifiedTicketingAttachmentInput{}
	if err := utils.UnmarshalJSON(data, &unifiedTicketingAttachmentInput, "", true, true); err == nil {
		u.UnifiedTicketingAttachmentInput = &unifiedTicketingAttachmentInput
		u.Type = AttachmentsTypeUnifiedTicketingAttachmentInput
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = AttachmentsTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Attachments", string(data))
}

func (u Attachments) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.UnifiedTicketingAttachmentInput != nil {
		return utils.MarshalJSON(u.UnifiedTicketingAttachmentInput, "", true)
	}

	return nil, errors.New("could not marshal union type Attachments: all fields are null")
}

type UnifiedTicketingTicketOutput struct {
	// The name of the ticket
	Name *string `json:"name"`
	// The status of the ticket. Authorized values are OPEN or CLOSED.
	Status *Status `json:"status,omitempty"`
	// The description of the ticket
	Description *string `json:"description"`
	// The date the ticket is due
	DueDate *time.Time `json:"due_date,omitempty"`
	// The type of the ticket. Authorized values are PROBLEM, QUESTION, or TASK
	Type *Type `json:"type,omitempty"`
	// The UUID of the parent ticket
	ParentTicket *string `json:"parent_ticket,omitempty"`
	// The collection UUIDs the ticket belongs to
	Collections []Collections `json:"collections,omitempty"`
	// The tags names of the ticket
	Tags []Tags `json:"tags,omitempty"`
	// The date the ticket has been completed
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// The priority of the ticket. Authorized values are HIGH, MEDIUM or LOW.
	Priority *Priority `json:"priority,omitempty"`
	// The users UUIDs the ticket is assigned to
	AssignedTo []string `json:"assigned_to,omitempty"`
	// The comment of the ticket
	Comment *Comment `json:"comment,omitempty"`
	// The UUID of the account which the ticket belongs to
	AccountID *string `json:"account_id,omitempty"`
	// The UUID of the contact which the ticket belongs to
	ContactID *string `json:"contact_id,omitempty"`
	// The attachements UUIDs tied to the ticket
	Attachments []Attachments `json:"attachments,omitempty"`
	// The custom field mappings of the ticket between the remote 3rd party & Panora
	FieldMappings map[string]any `json:"field_mappings,omitempty"`
	// The UUID of the ticket
	ID *string `json:"id,omitempty"`
	// The id of the ticket in the context of the 3rd Party
	RemoteID *string `json:"remote_id,omitempty"`
	// The remote data of the ticket in the context of the 3rd Party
	RemoteData map[string]any `json:"remote_data,omitempty"`
	// The created date of the object
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The modified date of the object
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
}

func (u UnifiedTicketingTicketOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedTicketingTicketOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedTicketingTicketOutput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UnifiedTicketingTicketOutput) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *UnifiedTicketingTicketOutput) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UnifiedTicketingTicketOutput) GetDueDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *UnifiedTicketingTicketOutput) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *UnifiedTicketingTicketOutput) GetParentTicket() *string {
	if o == nil {
		return nil
	}
	return o.ParentTicket
}

func (o *UnifiedTicketingTicketOutput) GetCollections() []Collections {
	if o == nil {
		return nil
	}
	return o.Collections
}

func (o *UnifiedTicketingTicketOutput) GetTags() []Tags {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *UnifiedTicketingTicketOutput) GetCompletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CompletedAt
}

func (o *UnifiedTicketingTicketOutput) GetPriority() *Priority {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *UnifiedTicketingTicketOutput) GetAssignedTo() []string {
	if o == nil {
		return nil
	}
	return o.AssignedTo
}

func (o *UnifiedTicketingTicketOutput) GetComment() *Comment {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *UnifiedTicketingTicketOutput) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *UnifiedTicketingTicketOutput) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *UnifiedTicketingTicketOutput) GetAttachments() []Attachments {
	if o == nil {
		return nil
	}
	return o.Attachments
}

func (o *UnifiedTicketingTicketOutput) GetFieldMappings() map[string]any {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}

func (o *UnifiedTicketingTicketOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedTicketingTicketOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedTicketingTicketOutput) GetRemoteData() map[string]any {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *UnifiedTicketingTicketOutput) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UnifiedTicketingTicketOutput) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}
