// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceQuickbooksResourceModel) ToCreateSDKType() *shared.SourceQuickbooksCreateRequest {
	var credentials shared.SourceQuickbooksAuthorizationMethod
	var sourceQuickbooksAuthorizationMethodOAuth20 *shared.SourceQuickbooksAuthorizationMethodOAuth20
	if r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20 != nil {
		accessToken := r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.AccessToken.ValueString()
		authType := new(shared.SourceQuickbooksAuthorizationMethodOAuth20AuthType)
		if !r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.AuthType.IsNull() {
			*authType = shared.SourceQuickbooksAuthorizationMethodOAuth20AuthType(r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.AuthType.ValueString())
		} else {
			authType = nil
		}
		clientID := r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.ClientSecret.ValueString()
		realmID := r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.RealmID.ValueString()
		refreshToken := r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.TokenExpiryDate.ValueString())
		sourceQuickbooksAuthorizationMethodOAuth20 = &shared.SourceQuickbooksAuthorizationMethodOAuth20{
			AccessToken:     accessToken,
			AuthType:        authType,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RealmID:         realmID,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceQuickbooksAuthorizationMethodOAuth20 != nil {
		credentials = shared.SourceQuickbooksAuthorizationMethod{
			SourceQuickbooksAuthorizationMethodOAuth20: sourceQuickbooksAuthorizationMethodOAuth20,
		}
	}
	sandbox := r.Configuration.Sandbox.ValueBool()
	sourceType := shared.SourceQuickbooksQuickbooks(r.Configuration.SourceType.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceQuickbooks{
		Credentials: credentials,
		Sandbox:     sandbox,
		SourceType:  sourceType,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceQuickbooksCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceQuickbooksResourceModel) ToUpdateSDKType() *shared.SourceQuickbooksPutRequest {
	var credentials shared.SourceQuickbooksUpdateAuthorizationMethod
	var sourceQuickbooksUpdateAuthorizationMethodOAuth20 *shared.SourceQuickbooksUpdateAuthorizationMethodOAuth20
	if r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20 != nil {
		accessToken := r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.AccessToken.ValueString()
		authType := new(shared.SourceQuickbooksUpdateAuthorizationMethodOAuth20AuthType)
		if !r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.AuthType.IsNull() {
			*authType = shared.SourceQuickbooksUpdateAuthorizationMethodOAuth20AuthType(r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.AuthType.ValueString())
		} else {
			authType = nil
		}
		clientID := r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.ClientSecret.ValueString()
		realmID := r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.RealmID.ValueString()
		refreshToken := r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.TokenExpiryDate.ValueString())
		sourceQuickbooksUpdateAuthorizationMethodOAuth20 = &shared.SourceQuickbooksUpdateAuthorizationMethodOAuth20{
			AccessToken:     accessToken,
			AuthType:        authType,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RealmID:         realmID,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceQuickbooksUpdateAuthorizationMethodOAuth20 != nil {
		credentials = shared.SourceQuickbooksUpdateAuthorizationMethod{
			SourceQuickbooksUpdateAuthorizationMethodOAuth20: sourceQuickbooksUpdateAuthorizationMethodOAuth20,
		}
	}
	sandbox := r.Configuration.Sandbox.ValueBool()
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceQuickbooksUpdate{
		Credentials: credentials,
		Sandbox:     sandbox,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceQuickbooksPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceQuickbooksResourceModel) ToDeleteSDKType() *shared.SourceQuickbooksCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceQuickbooksResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
