// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RequestRequest struct {
	IntegrationID         string                           `queryParam:"style=form,explode=true,name=integrationId"`
	LinkedUserID          string                           `queryParam:"style=form,explode=true,name=linkedUserId"`
	Vertical              string                           `queryParam:"style=form,explode=true,name=vertical"`
	PassThroughRequestDto components.PassThroughRequestDto `request:"mediaType=application/json"`
}

func (o *RequestRequest) GetIntegrationID() string {
	if o == nil {
		return ""
	}
	return o.IntegrationID
}

func (o *RequestRequest) GetLinkedUserID() string {
	if o == nil {
		return ""
	}
	return o.LinkedUserID
}

func (o *RequestRequest) GetVertical() string {
	if o == nil {
		return ""
	}
	return o.Vertical
}

func (o *RequestRequest) GetPassThroughRequestDto() components.PassThroughRequestDto {
	if o == nil {
		return components.PassThroughRequestDto{}
	}
	return o.PassThroughRequestDto
}

type RequestResponse struct {
	HTTPMeta            components.HTTPMetadata `json:"-"`
	PassThroughResponse *components.PassThroughResponse
}

func (o *RequestResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RequestResponse) GetPassThroughResponse() *components.PassThroughResponse {
	if o == nil {
		return nil
	}
	return o.PassThroughResponse
}
