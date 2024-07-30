// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

type UnifiedTicketingTicketInputFieldMappings struct {
}

type UnifiedTicketingTicketInput struct {
	// The name of the ticket
	Name string `json:"name"`
	// The status of the ticket. Authorized values are OPEN or CLOSED.
	Status *string `json:"status,omitempty"`
	// The description of the ticket
	Description string `json:"description"`
	// The date the ticket is due
	DueDate *time.Time `json:"due_date,omitempty"`
	// The type of the ticket. Authorized values are PROBLEM, QUESTION, or TASK
	Type *string `json:"type,omitempty"`
	// The UUID of the parent ticket
	ParentTicket *string `json:"parent_ticket,omitempty"`
	// The collection UUIDs the ticket belongs to
	Collections *string `json:"collections,omitempty"`
	// The tags names of the ticket
	Tags []string `json:"tags,omitempty"`
	// The date the ticket has been completed
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// The priority of the ticket. Authorized values are HIGH, MEDIUM or LOW.
	Priority *string `json:"priority,omitempty"`
	// The users UUIDs the ticket is assigned to
	AssignedTo []string `json:"assigned_to,omitempty"`
	// The comment of the ticket
	Comment *UnifiedTicketingCommentInput `json:"comment,omitempty"`
	// The UUID of the account which the ticket belongs to
	AccountID *string `json:"account_id,omitempty"`
	// The UUID of the contact which the ticket belongs to
	ContactID *string `json:"contact_id,omitempty"`
	// The attachements UUIDs tied to the ticket
	Attachments   []string                                 `json:"attachments,omitempty"`
	FieldMappings UnifiedTicketingTicketInputFieldMappings `json:"field_mappings"`
}

func (u UnifiedTicketingTicketInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedTicketingTicketInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedTicketingTicketInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UnifiedTicketingTicketInput) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *UnifiedTicketingTicketInput) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *UnifiedTicketingTicketInput) GetDueDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *UnifiedTicketingTicketInput) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *UnifiedTicketingTicketInput) GetParentTicket() *string {
	if o == nil {
		return nil
	}
	return o.ParentTicket
}

func (o *UnifiedTicketingTicketInput) GetCollections() *string {
	if o == nil {
		return nil
	}
	return o.Collections
}

func (o *UnifiedTicketingTicketInput) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *UnifiedTicketingTicketInput) GetCompletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CompletedAt
}

func (o *UnifiedTicketingTicketInput) GetPriority() *string {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *UnifiedTicketingTicketInput) GetAssignedTo() []string {
	if o == nil {
		return nil
	}
	return o.AssignedTo
}

func (o *UnifiedTicketingTicketInput) GetComment() *UnifiedTicketingCommentInput {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *UnifiedTicketingTicketInput) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *UnifiedTicketingTicketInput) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *UnifiedTicketingTicketInput) GetAttachments() []string {
	if o == nil {
		return nil
	}
	return o.Attachments
}

func (o *UnifiedTicketingTicketInput) GetFieldMappings() UnifiedTicketingTicketInputFieldMappings {
	if o == nil {
		return UnifiedTicketingTicketInputFieldMappings{}
	}
	return o.FieldMappings
}