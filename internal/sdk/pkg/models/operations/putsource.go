// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceRequest struct {
	SourcePutRequest *shared.SourcePutRequest `request:"mediaType=application/json"`
	SourceID         string                   `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceRequest) GetSourcePutRequest() *shared.SourcePutRequest {
	if o == nil {
		return nil
	}
	return o.SourcePutRequest
}

func (o *PutSourceRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Update a source and fully overwrite it
	SourceResponse *shared.SourceResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceResponse) GetSourceResponse() *shared.SourceResponse {
	if o == nil {
		return nil
	}
	return o.SourceResponse
}

func (o *PutSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
