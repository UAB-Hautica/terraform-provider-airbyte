// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceLokaliseResourceModel) ToCreateSDKType() *shared.SourceLokaliseCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	projectID := r.Configuration.ProjectID.ValueString()
	configuration := shared.SourceLokalise{
		APIKey:    apiKey,
		ProjectID: projectID,
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
	out := shared.SourceLokaliseCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceLokaliseResourceModel) ToGetSDKType() *shared.SourceLokaliseCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceLokaliseResourceModel) ToUpdateSDKType() *shared.SourceLokalisePutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	projectID := r.Configuration.ProjectID.ValueString()
	configuration := shared.SourceLokaliseUpdate{
		APIKey:    apiKey,
		ProjectID: projectID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceLokalisePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceLokaliseResourceModel) ToDeleteSDKType() *shared.SourceLokaliseCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceLokaliseResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceLokaliseResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
