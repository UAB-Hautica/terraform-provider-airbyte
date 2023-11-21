// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceZendeskTalkResourceModel) ToCreateSDKType() *shared.SourceZendeskTalkCreateRequest {
	var credentials *shared.SourceZendeskTalkAuthentication
	if r.Configuration.Credentials != nil {
		var sourceZendeskTalkAPIToken *shared.SourceZendeskTalkAPIToken
		if r.Configuration.Credentials.APIToken != nil {
			var additionalProperties interface{}
			if !r.Configuration.Credentials.APIToken.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.APIToken.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.APIToken.AdditionalProperties.ValueString()), &additionalProperties)
			}
			apiToken := r.Configuration.Credentials.APIToken.APIToken.ValueString()
			email := r.Configuration.Credentials.APIToken.Email.ValueString()
			sourceZendeskTalkAPIToken = &shared.SourceZendeskTalkAPIToken{
				AdditionalProperties: additionalProperties,
				APIToken:             apiToken,
				Email:                email,
			}
		}
		if sourceZendeskTalkAPIToken != nil {
			credentials = &shared.SourceZendeskTalkAuthentication{
				SourceZendeskTalkAPIToken: sourceZendeskTalkAPIToken,
			}
		}
		var sourceZendeskTalkOAuth20 *shared.SourceZendeskTalkOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			var additionalProperties1 interface{}
			if !r.Configuration.Credentials.OAuth20.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.OAuth20.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.OAuth20.AdditionalProperties.ValueString()), &additionalProperties1)
			}
			accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			clientID := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientID.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientID.IsNull() {
				*clientID = r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientSecret.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientSecret.IsNull() {
				*clientSecret = r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			sourceZendeskTalkOAuth20 = &shared.SourceZendeskTalkOAuth20{
				AdditionalProperties: additionalProperties1,
				AccessToken:          accessToken,
				ClientID:             clientID,
				ClientSecret:         clientSecret,
			}
		}
		if sourceZendeskTalkOAuth20 != nil {
			credentials = &shared.SourceZendeskTalkAuthentication{
				SourceZendeskTalkOAuth20: sourceZendeskTalkOAuth20,
			}
		}
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	subdomain := r.Configuration.Subdomain.ValueString()
	configuration := shared.SourceZendeskTalk{
		Credentials: credentials,
		StartDate:   startDate,
		Subdomain:   subdomain,
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
	out := shared.SourceZendeskTalkCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZendeskTalkResourceModel) ToGetSDKType() *shared.SourceZendeskTalkCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZendeskTalkResourceModel) ToUpdateSDKType() *shared.SourceZendeskTalkPutRequest {
	var credentials *shared.SourceZendeskTalkUpdateAuthentication
	if r.Configuration.Credentials != nil {
		var sourceZendeskTalkUpdateAPIToken *shared.SourceZendeskTalkUpdateAPIToken
		if r.Configuration.Credentials.APIToken != nil {
			var additionalProperties interface{}
			if !r.Configuration.Credentials.APIToken.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.APIToken.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.APIToken.AdditionalProperties.ValueString()), &additionalProperties)
			}
			apiToken := r.Configuration.Credentials.APIToken.APIToken.ValueString()
			email := r.Configuration.Credentials.APIToken.Email.ValueString()
			sourceZendeskTalkUpdateAPIToken = &shared.SourceZendeskTalkUpdateAPIToken{
				AdditionalProperties: additionalProperties,
				APIToken:             apiToken,
				Email:                email,
			}
		}
		if sourceZendeskTalkUpdateAPIToken != nil {
			credentials = &shared.SourceZendeskTalkUpdateAuthentication{
				SourceZendeskTalkUpdateAPIToken: sourceZendeskTalkUpdateAPIToken,
			}
		}
		var sourceZendeskTalkUpdateOAuth20 *shared.SourceZendeskTalkUpdateOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			var additionalProperties1 interface{}
			if !r.Configuration.Credentials.OAuth20.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.OAuth20.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.OAuth20.AdditionalProperties.ValueString()), &additionalProperties1)
			}
			accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			clientID := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientID.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientID.IsNull() {
				*clientID = r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientSecret.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientSecret.IsNull() {
				*clientSecret = r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			sourceZendeskTalkUpdateOAuth20 = &shared.SourceZendeskTalkUpdateOAuth20{
				AdditionalProperties: additionalProperties1,
				AccessToken:          accessToken,
				ClientID:             clientID,
				ClientSecret:         clientSecret,
			}
		}
		if sourceZendeskTalkUpdateOAuth20 != nil {
			credentials = &shared.SourceZendeskTalkUpdateAuthentication{
				SourceZendeskTalkUpdateOAuth20: sourceZendeskTalkUpdateOAuth20,
			}
		}
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	subdomain := r.Configuration.Subdomain.ValueString()
	configuration := shared.SourceZendeskTalkUpdate{
		Credentials: credentials,
		StartDate:   startDate,
		Subdomain:   subdomain,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZendeskTalkPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZendeskTalkResourceModel) ToDeleteSDKType() *shared.SourceZendeskTalkCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZendeskTalkResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceZendeskTalkResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
