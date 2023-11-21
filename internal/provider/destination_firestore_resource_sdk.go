// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationFirestoreResourceModel) ToCreateSDKType() *shared.DestinationFirestoreCreateRequest {
	credentialsJSON := new(string)
	if !r.Configuration.CredentialsJSON.IsUnknown() && !r.Configuration.CredentialsJSON.IsNull() {
		*credentialsJSON = r.Configuration.CredentialsJSON.ValueString()
	} else {
		credentialsJSON = nil
	}
	projectID := r.Configuration.ProjectID.ValueString()
	configuration := shared.DestinationFirestore{
		CredentialsJSON: credentialsJSON,
		ProjectID:       projectID,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationFirestoreCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationFirestoreResourceModel) ToGetSDKType() *shared.DestinationFirestoreCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationFirestoreResourceModel) ToUpdateSDKType() *shared.DestinationFirestorePutRequest {
	credentialsJSON := new(string)
	if !r.Configuration.CredentialsJSON.IsUnknown() && !r.Configuration.CredentialsJSON.IsNull() {
		*credentialsJSON = r.Configuration.CredentialsJSON.ValueString()
	} else {
		credentialsJSON = nil
	}
	projectID := r.Configuration.ProjectID.ValueString()
	configuration := shared.DestinationFirestoreUpdate{
		CredentialsJSON: credentialsJSON,
		ProjectID:       projectID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationFirestorePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationFirestoreResourceModel) ToDeleteSDKType() *shared.DestinationFirestoreCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationFirestoreResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationFirestoreResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
