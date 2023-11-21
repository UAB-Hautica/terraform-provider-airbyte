// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceGreenhouseResourceModel) ToCreateSDKType() *shared.SourceGreenhouseCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	configuration := shared.SourceGreenhouse{
		APIKey: apiKey,
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
	out := shared.SourceGreenhouseCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGreenhouseResourceModel) ToGetSDKType() *shared.SourceGreenhouseCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGreenhouseResourceModel) ToUpdateSDKType() *shared.SourceGreenhousePutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	configuration := shared.SourceGreenhouseUpdate{
		APIKey: apiKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGreenhousePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGreenhouseResourceModel) ToDeleteSDKType() *shared.SourceGreenhouseCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGreenhouseResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceGreenhouseResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
