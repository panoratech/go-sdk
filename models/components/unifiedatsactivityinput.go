// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

type UnifiedAtsActivityInputFieldMappings struct {
}

type UnifiedAtsActivityInput struct {
	// The type of activity
	ActivityType *string `json:"activity_type,omitempty"`
	// The subject of the activity
	Subject *string `json:"subject,omitempty"`
	// The body of the activity
	Body *string `json:"body,omitempty"`
	// The visibility of the activity
	Visibility *string `json:"visibility,omitempty"`
	// The UUID of the candidate
	CandidateID *string `json:"candidate_id,omitempty"`
	// The remote creation date of the activity
	RemoteCreatedAt *time.Time                           `json:"remote_created_at,omitempty"`
	FieldMappings   UnifiedAtsActivityInputFieldMappings `json:"field_mappings"`
}

func (u UnifiedAtsActivityInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedAtsActivityInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedAtsActivityInput) GetActivityType() *string {
	if o == nil {
		return nil
	}
	return o.ActivityType
}

func (o *UnifiedAtsActivityInput) GetSubject() *string {
	if o == nil {
		return nil
	}
	return o.Subject
}

func (o *UnifiedAtsActivityInput) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *UnifiedAtsActivityInput) GetVisibility() *string {
	if o == nil {
		return nil
	}
	return o.Visibility
}

func (o *UnifiedAtsActivityInput) GetCandidateID() *string {
	if o == nil {
		return nil
	}
	return o.CandidateID
}

func (o *UnifiedAtsActivityInput) GetRemoteCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteCreatedAt
}

func (o *UnifiedAtsActivityInput) GetFieldMappings() UnifiedAtsActivityInputFieldMappings {
	if o == nil {
		return UnifiedAtsActivityInputFieldMappings{}
	}
	return o.FieldMappings
}