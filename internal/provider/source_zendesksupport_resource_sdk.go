// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceZendeskSupportResourceModel) ToCreateSDKType() *shared.SourceZendeskSupportCreateRequest {
	var credentials *shared.SourceZendeskSupportAuthentication
	if r.Configuration.Credentials != nil {
		var sourceZendeskSupportAuthenticationOAuth20 *shared.SourceZendeskSupportAuthenticationOAuth20
		if r.Configuration.Credentials.SourceZendeskSupportAuthenticationOAuth20 != nil {
			accessToken := r.Configuration.Credentials.SourceZendeskSupportAuthenticationOAuth20.AccessToken.ValueString()
			credentials1 := new(shared.SourceZendeskSupportAuthenticationOAuth20Credentials)
			if !r.Configuration.Credentials.SourceZendeskSupportAuthenticationOAuth20.Credentials.IsUnknown() && !r.Configuration.Credentials.SourceZendeskSupportAuthenticationOAuth20.Credentials.IsNull() {
				*credentials1 = shared.SourceZendeskSupportAuthenticationOAuth20Credentials(r.Configuration.Credentials.SourceZendeskSupportAuthenticationOAuth20.Credentials.ValueString())
			} else {
				credentials1 = nil
			}
			var additionalProperties interface{}
			if !r.Configuration.Credentials.SourceZendeskSupportAuthenticationOAuth20.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.SourceZendeskSupportAuthenticationOAuth20.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.SourceZendeskSupportAuthenticationOAuth20.AdditionalProperties.ValueString()), &additionalProperties)
			}
			sourceZendeskSupportAuthenticationOAuth20 = &shared.SourceZendeskSupportAuthenticationOAuth20{
				AccessToken:          accessToken,
				Credentials:          credentials1,
				AdditionalProperties: additionalProperties,
			}
		}
		if sourceZendeskSupportAuthenticationOAuth20 != nil {
			credentials = &shared.SourceZendeskSupportAuthentication{
				SourceZendeskSupportAuthenticationOAuth20: sourceZendeskSupportAuthenticationOAuth20,
			}
		}
		var sourceZendeskSupportAuthenticationAPIToken *shared.SourceZendeskSupportAuthenticationAPIToken
		if r.Configuration.Credentials.SourceZendeskSupportAuthenticationAPIToken != nil {
			apiToken := r.Configuration.Credentials.SourceZendeskSupportAuthenticationAPIToken.APIToken.ValueString()
			credentials2 := new(shared.SourceZendeskSupportAuthenticationAPITokenCredentials)
			if !r.Configuration.Credentials.SourceZendeskSupportAuthenticationAPIToken.Credentials.IsUnknown() && !r.Configuration.Credentials.SourceZendeskSupportAuthenticationAPIToken.Credentials.IsNull() {
				*credentials2 = shared.SourceZendeskSupportAuthenticationAPITokenCredentials(r.Configuration.Credentials.SourceZendeskSupportAuthenticationAPIToken.Credentials.ValueString())
			} else {
				credentials2 = nil
			}
			email := r.Configuration.Credentials.SourceZendeskSupportAuthenticationAPIToken.Email.ValueString()
			var additionalProperties1 interface{}
			if !r.Configuration.Credentials.SourceZendeskSupportAuthenticationAPIToken.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.SourceZendeskSupportAuthenticationAPIToken.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.SourceZendeskSupportAuthenticationAPIToken.AdditionalProperties.ValueString()), &additionalProperties1)
			}
			sourceZendeskSupportAuthenticationAPIToken = &shared.SourceZendeskSupportAuthenticationAPIToken{
				APIToken:             apiToken,
				Credentials:          credentials2,
				Email:                email,
				AdditionalProperties: additionalProperties1,
			}
		}
		if sourceZendeskSupportAuthenticationAPIToken != nil {
			credentials = &shared.SourceZendeskSupportAuthentication{
				SourceZendeskSupportAuthenticationAPIToken: sourceZendeskSupportAuthenticationAPIToken,
			}
		}
	}
	ignorePagination := new(bool)
	if !r.Configuration.IgnorePagination.IsUnknown() && !r.Configuration.IgnorePagination.IsNull() {
		*ignorePagination = r.Configuration.IgnorePagination.ValueBool()
	} else {
		ignorePagination = nil
	}
	sourceType := shared.SourceZendeskSupportZendeskSupport(r.Configuration.SourceType.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	subdomain := r.Configuration.Subdomain.ValueString()
	configuration := shared.SourceZendeskSupport{
		Credentials:      credentials,
		IgnorePagination: ignorePagination,
		SourceType:       sourceType,
		StartDate:        startDate,
		Subdomain:        subdomain,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZendeskSupportCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZendeskSupportResourceModel) ToGetSDKType() *shared.SourceZendeskSupportCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZendeskSupportResourceModel) ToUpdateSDKType() *shared.SourceZendeskSupportPutRequest {
	var credentials *shared.SourceZendeskSupportUpdateAuthentication
	if r.Configuration.Credentials != nil {
		var sourceZendeskSupportUpdateAuthenticationOAuth20 *shared.SourceZendeskSupportUpdateAuthenticationOAuth20
		if r.Configuration.Credentials.SourceZendeskSupportUpdateAuthenticationOAuth20 != nil {
			accessToken := r.Configuration.Credentials.SourceZendeskSupportUpdateAuthenticationOAuth20.AccessToken.ValueString()
			credentials1 := new(shared.SourceZendeskSupportUpdateAuthenticationOAuth20Credentials)
			if !r.Configuration.Credentials.SourceZendeskSupportUpdateAuthenticationOAuth20.Credentials.IsUnknown() && !r.Configuration.Credentials.SourceZendeskSupportUpdateAuthenticationOAuth20.Credentials.IsNull() {
				*credentials1 = shared.SourceZendeskSupportUpdateAuthenticationOAuth20Credentials(r.Configuration.Credentials.SourceZendeskSupportUpdateAuthenticationOAuth20.Credentials.ValueString())
			} else {
				credentials1 = nil
			}
			var additionalProperties interface{}
			if !r.Configuration.Credentials.SourceZendeskSupportUpdateAuthenticationOAuth20.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.SourceZendeskSupportUpdateAuthenticationOAuth20.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.SourceZendeskSupportUpdateAuthenticationOAuth20.AdditionalProperties.ValueString()), &additionalProperties)
			}
			sourceZendeskSupportUpdateAuthenticationOAuth20 = &shared.SourceZendeskSupportUpdateAuthenticationOAuth20{
				AccessToken:          accessToken,
				Credentials:          credentials1,
				AdditionalProperties: additionalProperties,
			}
		}
		if sourceZendeskSupportUpdateAuthenticationOAuth20 != nil {
			credentials = &shared.SourceZendeskSupportUpdateAuthentication{
				SourceZendeskSupportUpdateAuthenticationOAuth20: sourceZendeskSupportUpdateAuthenticationOAuth20,
			}
		}
		var sourceZendeskSupportUpdateAuthenticationAPIToken *shared.SourceZendeskSupportUpdateAuthenticationAPIToken
		if r.Configuration.Credentials.SourceZendeskSupportUpdateAuthenticationAPIToken != nil {
			apiToken := r.Configuration.Credentials.SourceZendeskSupportUpdateAuthenticationAPIToken.APIToken.ValueString()
			credentials2 := new(shared.SourceZendeskSupportUpdateAuthenticationAPITokenCredentials)
			if !r.Configuration.Credentials.SourceZendeskSupportUpdateAuthenticationAPIToken.Credentials.IsUnknown() && !r.Configuration.Credentials.SourceZendeskSupportUpdateAuthenticationAPIToken.Credentials.IsNull() {
				*credentials2 = shared.SourceZendeskSupportUpdateAuthenticationAPITokenCredentials(r.Configuration.Credentials.SourceZendeskSupportUpdateAuthenticationAPIToken.Credentials.ValueString())
			} else {
				credentials2 = nil
			}
			email := r.Configuration.Credentials.SourceZendeskSupportUpdateAuthenticationAPIToken.Email.ValueString()
			var additionalProperties1 interface{}
			if !r.Configuration.Credentials.SourceZendeskSupportUpdateAuthenticationAPIToken.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.SourceZendeskSupportUpdateAuthenticationAPIToken.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.SourceZendeskSupportUpdateAuthenticationAPIToken.AdditionalProperties.ValueString()), &additionalProperties1)
			}
			sourceZendeskSupportUpdateAuthenticationAPIToken = &shared.SourceZendeskSupportUpdateAuthenticationAPIToken{
				APIToken:             apiToken,
				Credentials:          credentials2,
				Email:                email,
				AdditionalProperties: additionalProperties1,
			}
		}
		if sourceZendeskSupportUpdateAuthenticationAPIToken != nil {
			credentials = &shared.SourceZendeskSupportUpdateAuthentication{
				SourceZendeskSupportUpdateAuthenticationAPIToken: sourceZendeskSupportUpdateAuthenticationAPIToken,
			}
		}
	}
	ignorePagination := new(bool)
	if !r.Configuration.IgnorePagination.IsUnknown() && !r.Configuration.IgnorePagination.IsNull() {
		*ignorePagination = r.Configuration.IgnorePagination.ValueBool()
	} else {
		ignorePagination = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	subdomain := r.Configuration.Subdomain.ValueString()
	configuration := shared.SourceZendeskSupportUpdate{
		Credentials:      credentials,
		IgnorePagination: ignorePagination,
		StartDate:        startDate,
		Subdomain:        subdomain,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZendeskSupportPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZendeskSupportResourceModel) ToDeleteSDKType() *shared.SourceZendeskSupportCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZendeskSupportResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceZendeskSupportResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
