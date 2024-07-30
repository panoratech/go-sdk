// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type UnifiedAtsJobinterviewstageOutputFieldMappings struct {
}

type UnifiedAtsJobinterviewstageOutputRemoteData struct {
}

type UnifiedAtsJobinterviewstageOutputCreatedAt struct {
}

type UnifiedAtsJobinterviewstageOutputModifiedAt struct {
}

type UnifiedAtsJobinterviewstageOutput struct {
	// The name of the job interview stage
	Name *string `json:"name,omitempty"`
	// The order of the stage
	StageOrder *float64 `json:"stage_order,omitempty"`
	// The UUID of the job
	JobID         *string                                        `json:"job_id,omitempty"`
	FieldMappings UnifiedAtsJobinterviewstageOutputFieldMappings `json:"field_mappings"`
	// The UUID of the job interview stage
	ID *string `json:"id,omitempty"`
	// The remote ID of the job interview stage in the context of the 3rd Party
	RemoteID   *string                                     `json:"remote_id,omitempty"`
	RemoteData UnifiedAtsJobinterviewstageOutputRemoteData `json:"remote_data"`
	CreatedAt  UnifiedAtsJobinterviewstageOutputCreatedAt  `json:"created_at"`
	ModifiedAt UnifiedAtsJobinterviewstageOutputModifiedAt `json:"modified_at"`
}

func (o *UnifiedAtsJobinterviewstageOutput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UnifiedAtsJobinterviewstageOutput) GetStageOrder() *float64 {
	if o == nil {
		return nil
	}
	return o.StageOrder
}

func (o *UnifiedAtsJobinterviewstageOutput) GetJobID() *string {
	if o == nil {
		return nil
	}
	return o.JobID
}

func (o *UnifiedAtsJobinterviewstageOutput) GetFieldMappings() UnifiedAtsJobinterviewstageOutputFieldMappings {
	if o == nil {
		return UnifiedAtsJobinterviewstageOutputFieldMappings{}
	}
	return o.FieldMappings
}

func (o *UnifiedAtsJobinterviewstageOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedAtsJobinterviewstageOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedAtsJobinterviewstageOutput) GetRemoteData() UnifiedAtsJobinterviewstageOutputRemoteData {
	if o == nil {
		return UnifiedAtsJobinterviewstageOutputRemoteData{}
	}
	return o.RemoteData
}

func (o *UnifiedAtsJobinterviewstageOutput) GetCreatedAt() UnifiedAtsJobinterviewstageOutputCreatedAt {
	if o == nil {
		return UnifiedAtsJobinterviewstageOutputCreatedAt{}
	}
	return o.CreatedAt
}

func (o *UnifiedAtsJobinterviewstageOutput) GetModifiedAt() UnifiedAtsJobinterviewstageOutputModifiedAt {
	if o == nil {
		return UnifiedAtsJobinterviewstageOutputModifiedAt{}
	}
	return o.ModifiedAt
}