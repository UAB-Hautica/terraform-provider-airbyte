// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteDestinationQdrantRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *DeleteDestinationQdrantRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

type DeleteDestinationQdrantResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteDestinationQdrantResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteDestinationQdrantResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteDestinationQdrantResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
