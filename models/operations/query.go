// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type QueryRequest struct {
	// The connection token
	XConnectionToken string               `header:"style=simple,explode=false,name=x-connection-token"`
	QueryBody        components.QueryBody `request:"mediaType=application/json"`
}

func (o *QueryRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *QueryRequest) GetQueryBody() components.QueryBody {
	if o == nil {
		return components.QueryBody{}
	}
	return o.QueryBody
}

type QueryResponse struct {
	HTTPMeta        components.HTTPMetadata `json:"-"`
	RagQueryOutputs []components.RagQueryOutput
}

func (o *QueryResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *QueryResponse) GetRagQueryOutputs() []components.RagQueryOutput {
	if o == nil {
		return nil
	}
	return o.RagQueryOutputs
}
