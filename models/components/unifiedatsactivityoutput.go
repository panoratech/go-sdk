// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

type UnifiedAtsActivityOutputFieldMappings struct {
}

type UnifiedAtsActivityOutputRemoteData struct {
}

type UnifiedAtsActivityOutputCreatedAt struct {
}

type UnifiedAtsActivityOutputModifiedAt struct {
}

type UnifiedAtsActivityOutput struct {
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
	RemoteCreatedAt *time.Time                            `json:"remote_created_at,omitempty"`
	FieldMappings   UnifiedAtsActivityOutputFieldMappings `json:"field_mappings"`
	// The UUID of the activity
	ID *string `json:"id,omitempty"`
	// The remote ID of the activity in the context of the 3rd Party
	RemoteID   *string                            `json:"remote_id,omitempty"`
	RemoteData UnifiedAtsActivityOutputRemoteData `json:"remote_data"`
	CreatedAt  UnifiedAtsActivityOutputCreatedAt  `json:"created_at"`
	ModifiedAt UnifiedAtsActivityOutputModifiedAt `json:"modified_at"`
}

func (u UnifiedAtsActivityOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedAtsActivityOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedAtsActivityOutput) GetActivityType() *string {
	if o == nil {
		return nil
	}
	return o.ActivityType
}

func (o *UnifiedAtsActivityOutput) GetSubject() *string {
	if o == nil {
		return nil
	}
	return o.Subject
}

func (o *UnifiedAtsActivityOutput) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *UnifiedAtsActivityOutput) GetVisibility() *string {
	if o == nil {
		return nil
	}
	return o.Visibility
}

func (o *UnifiedAtsActivityOutput) GetCandidateID() *string {
	if o == nil {
		return nil
	}
	return o.CandidateID
}

func (o *UnifiedAtsActivityOutput) GetRemoteCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteCreatedAt
}

func (o *UnifiedAtsActivityOutput) GetFieldMappings() UnifiedAtsActivityOutputFieldMappings {
	if o == nil {
		return UnifiedAtsActivityOutputFieldMappings{}
	}
	return o.FieldMappings
}

func (o *UnifiedAtsActivityOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedAtsActivityOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedAtsActivityOutput) GetRemoteData() UnifiedAtsActivityOutputRemoteData {
	if o == nil {
		return UnifiedAtsActivityOutputRemoteData{}
	}
	return o.RemoteData
}

func (o *UnifiedAtsActivityOutput) GetCreatedAt() UnifiedAtsActivityOutputCreatedAt {
	if o == nil {
		return UnifiedAtsActivityOutputCreatedAt{}
	}
	return o.CreatedAt
}

func (o *UnifiedAtsActivityOutput) GetModifiedAt() UnifiedAtsActivityOutputModifiedAt {
	if o == nil {
		return UnifiedAtsActivityOutputModifiedAt{}
	}
	return o.ModifiedAt
}
