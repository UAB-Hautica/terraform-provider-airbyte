// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourcePaypalTransactionResourceModel) ToCreateSDKType() *shared.SourcePaypalTransactionCreateRequest {
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	isSandbox := new(bool)
	if !r.Configuration.IsSandbox.IsUnknown() && !r.Configuration.IsSandbox.IsNull() {
		*isSandbox = r.Configuration.IsSandbox.ValueBool()
	} else {
		isSandbox = nil
	}
	refreshToken := new(string)
	if !r.Configuration.RefreshToken.IsUnknown() && !r.Configuration.RefreshToken.IsNull() {
		*refreshToken = r.Configuration.RefreshToken.ValueString()
	} else {
		refreshToken = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	timeWindow := new(int64)
	if !r.Configuration.TimeWindow.IsUnknown() && !r.Configuration.TimeWindow.IsNull() {
		*timeWindow = r.Configuration.TimeWindow.ValueInt64()
	} else {
		timeWindow = nil
	}
	configuration := shared.SourcePaypalTransaction{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		IsSandbox:    isSandbox,
		RefreshToken: refreshToken,
		StartDate:    startDate,
		TimeWindow:   timeWindow,
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
	out := shared.SourcePaypalTransactionCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePaypalTransactionResourceModel) ToGetSDKType() *shared.SourcePaypalTransactionCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePaypalTransactionResourceModel) ToUpdateSDKType() *shared.SourcePaypalTransactionPutRequest {
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	isSandbox := new(bool)
	if !r.Configuration.IsSandbox.IsUnknown() && !r.Configuration.IsSandbox.IsNull() {
		*isSandbox = r.Configuration.IsSandbox.ValueBool()
	} else {
		isSandbox = nil
	}
	refreshToken := new(string)
	if !r.Configuration.RefreshToken.IsUnknown() && !r.Configuration.RefreshToken.IsNull() {
		*refreshToken = r.Configuration.RefreshToken.ValueString()
	} else {
		refreshToken = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	timeWindow := new(int64)
	if !r.Configuration.TimeWindow.IsUnknown() && !r.Configuration.TimeWindow.IsNull() {
		*timeWindow = r.Configuration.TimeWindow.ValueInt64()
	} else {
		timeWindow = nil
	}
	configuration := shared.SourcePaypalTransactionUpdate{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		IsSandbox:    isSandbox,
		RefreshToken: refreshToken,
		StartDate:    startDate,
		TimeWindow:   timeWindow,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePaypalTransactionPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePaypalTransactionResourceModel) ToDeleteSDKType() *shared.SourcePaypalTransactionCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePaypalTransactionResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourcePaypalTransactionResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
