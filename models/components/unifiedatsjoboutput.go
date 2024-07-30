// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

type UnifiedAtsJobOutputFieldMappings struct {
}

type UnifiedAtsJobOutputRemoteData struct {
}

type UnifiedAtsJobOutputCreatedAt struct {
}

type UnifiedAtsJobOutputModifiedAt struct {
}

type UnifiedAtsJobOutput struct {
	// The name of the job
	Name *string `json:"name,omitempty"`
	// The description of the job
	Description *string `json:"description,omitempty"`
	// The code of the job
	Code *string `json:"code,omitempty"`
	// The status of the job
	Status *string `json:"status,omitempty"`
	// The type of the job
	Type *string `json:"type,omitempty"`
	// Whether the job is confidential
	Confidential *bool `json:"confidential,omitempty"`
	// The departments UUIDs associated with the job
	Departments []string `json:"departments,omitempty"`
	// The offices UUIDs associated with the job
	Offices []string `json:"offices,omitempty"`
	// The managers UUIDs associated with the job
	Managers []string `json:"managers,omitempty"`
	// The recruiters UUIDs associated with the job
	Recruiters []string `json:"recruiters,omitempty"`
	// The remote creation date of the job
	RemoteCreatedAt *time.Time `json:"remote_created_at,omitempty"`
	// The remote modification date of the job
	RemoteUpdatedAt *time.Time                       `json:"remote_updated_at,omitempty"`
	FieldMappings   UnifiedAtsJobOutputFieldMappings `json:"field_mappings"`
	// The UUID of the job
	ID *string `json:"id,omitempty"`
	// The remote ID of the job in the context of the 3rd Party
	RemoteID   *string                       `json:"remote_id,omitempty"`
	RemoteData UnifiedAtsJobOutputRemoteData `json:"remote_data"`
	CreatedAt  UnifiedAtsJobOutputCreatedAt  `json:"created_at"`
	ModifiedAt UnifiedAtsJobOutputModifiedAt `json:"modified_at"`
}

func (u UnifiedAtsJobOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedAtsJobOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedAtsJobOutput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UnifiedAtsJobOutput) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UnifiedAtsJobOutput) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *UnifiedAtsJobOutput) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *UnifiedAtsJobOutput) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *UnifiedAtsJobOutput) GetConfidential() *bool {
	if o == nil {
		return nil
	}
	return o.Confidential
}

func (o *UnifiedAtsJobOutput) GetDepartments() []string {
	if o == nil {
		return nil
	}
	return o.Departments
}

func (o *UnifiedAtsJobOutput) GetOffices() []string {
	if o == nil {
		return nil
	}
	return o.Offices
}

func (o *UnifiedAtsJobOutput) GetManagers() []string {
	if o == nil {
		return nil
	}
	return o.Managers
}

func (o *UnifiedAtsJobOutput) GetRecruiters() []string {
	if o == nil {
		return nil
	}
	return o.Recruiters
}

func (o *UnifiedAtsJobOutput) GetRemoteCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteCreatedAt
}

func (o *UnifiedAtsJobOutput) GetRemoteUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteUpdatedAt
}

func (o *UnifiedAtsJobOutput) GetFieldMappings() UnifiedAtsJobOutputFieldMappings {
	if o == nil {
		return UnifiedAtsJobOutputFieldMappings{}
	}
	return o.FieldMappings
}

func (o *UnifiedAtsJobOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedAtsJobOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedAtsJobOutput) GetRemoteData() UnifiedAtsJobOutputRemoteData {
	if o == nil {
		return UnifiedAtsJobOutputRemoteData{}
	}
	return o.RemoteData
}

func (o *UnifiedAtsJobOutput) GetCreatedAt() UnifiedAtsJobOutputCreatedAt {
	if o == nil {
		return UnifiedAtsJobOutputCreatedAt{}
	}
	return o.CreatedAt
}

func (o *UnifiedAtsJobOutput) GetModifiedAt() UnifiedAtsJobOutputModifiedAt {
	if o == nil {
		return UnifiedAtsJobOutputModifiedAt{}
	}
	return o.ModifiedAt
}
