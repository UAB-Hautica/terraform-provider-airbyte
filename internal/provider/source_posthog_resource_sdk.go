// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourcePosthogResourceModel) ToCreateSDKType() *shared.SourcePosthogCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	baseURL := new(string)
	if !r.Configuration.BaseURL.IsUnknown() && !r.Configuration.BaseURL.IsNull() {
		*baseURL = r.Configuration.BaseURL.ValueString()
	} else {
		baseURL = nil
	}
	eventsTimeStep := new(int64)
	if !r.Configuration.EventsTimeStep.IsUnknown() && !r.Configuration.EventsTimeStep.IsNull() {
		*eventsTimeStep = r.Configuration.EventsTimeStep.ValueInt64()
	} else {
		eventsTimeStep = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourcePosthog{
		APIKey:         apiKey,
		BaseURL:        baseURL,
		EventsTimeStep: eventsTimeStep,
		StartDate:      startDate,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePosthogCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePosthogResourceModel) ToGetSDKType() *shared.SourcePosthogCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePosthogResourceModel) ToUpdateSDKType() *shared.SourcePosthogPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	baseURL := new(string)
	if !r.Configuration.BaseURL.IsUnknown() && !r.Configuration.BaseURL.IsNull() {
		*baseURL = r.Configuration.BaseURL.ValueString()
	} else {
		baseURL = nil
	}
	eventsTimeStep := new(int64)
	if !r.Configuration.EventsTimeStep.IsUnknown() && !r.Configuration.EventsTimeStep.IsNull() {
		*eventsTimeStep = r.Configuration.EventsTimeStep.ValueInt64()
	} else {
		eventsTimeStep = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourcePosthogUpdate{
		APIKey:         apiKey,
		BaseURL:        baseURL,
		EventsTimeStep: eventsTimeStep,
		StartDate:      startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePosthogPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePosthogResourceModel) ToDeleteSDKType() *shared.SourcePosthogCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePosthogResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourcePosthogResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
