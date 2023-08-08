// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceMicrosoftTeamsResourceModel) ToCreateSDKType() *shared.SourceMicrosoftTeamsCreateRequest {
	var credentials *shared.SourceMicrosoftTeamsAuthenticationMechanism
	if r.Configuration.Credentials != nil {
		var sourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20 *shared.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20
		if r.Configuration.Credentials.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20 != nil {
			authType := new(shared.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20AuthType)
			if !r.Configuration.Credentials.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20.AuthType.IsNull() {
				*authType = shared.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20AuthType(r.Configuration.Credentials.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20.AuthType.ValueString())
			} else {
				authType = nil
			}
			clientID := r.Configuration.Credentials.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20.RefreshToken.ValueString()
			tenantID := r.Configuration.Credentials.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20.TenantID.ValueString()
			sourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20 = &shared.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20{
				AuthType:     authType,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
				TenantID:     tenantID,
			}
		}
		if sourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20 != nil {
			credentials = &shared.SourceMicrosoftTeamsAuthenticationMechanism{
				SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20: sourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftOAuth20,
			}
		}
		var sourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft *shared.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft
		if r.Configuration.Credentials.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft != nil {
			authType1 := new(shared.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftAuthType)
			if !r.Configuration.Credentials.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft.AuthType.IsNull() {
				*authType1 = shared.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoftAuthType(r.Configuration.Credentials.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft.AuthType.ValueString())
			} else {
				authType1 = nil
			}
			clientId1 := r.Configuration.Credentials.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft.ClientID.ValueString()
			clientSecret1 := r.Configuration.Credentials.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft.ClientSecret.ValueString()
			tenantId1 := r.Configuration.Credentials.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft.TenantID.ValueString()
			sourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft = &shared.SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft{
				AuthType:     authType1,
				ClientID:     clientId1,
				ClientSecret: clientSecret1,
				TenantID:     tenantId1,
			}
		}
		if sourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft != nil {
			credentials = &shared.SourceMicrosoftTeamsAuthenticationMechanism{
				SourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft: sourceMicrosoftTeamsAuthenticationMechanismAuthenticateViaMicrosoft,
			}
		}
	}
	period := r.Configuration.Period.ValueString()
	sourceType := shared.SourceMicrosoftTeamsMicrosoftTeams(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceMicrosoftTeams{
		Credentials: credentials,
		Period:      period,
		SourceType:  sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMicrosoftTeamsCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMicrosoftTeamsResourceModel) ToGetSDKType() *shared.SourceMicrosoftTeamsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMicrosoftTeamsResourceModel) ToUpdateSDKType() *shared.SourceMicrosoftTeamsPutRequest {
	var credentials *shared.SourceMicrosoftTeamsUpdateAuthenticationMechanism
	if r.Configuration.Credentials != nil {
		var sourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20 *shared.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20
		if r.Configuration.Credentials.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20 != nil {
			authType := new(shared.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20AuthType)
			if !r.Configuration.Credentials.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20.AuthType.IsNull() {
				*authType = shared.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20AuthType(r.Configuration.Credentials.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20.AuthType.ValueString())
			} else {
				authType = nil
			}
			clientID := r.Configuration.Credentials.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20.RefreshToken.ValueString()
			tenantID := r.Configuration.Credentials.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20.TenantID.ValueString()
			sourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20 = &shared.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20{
				AuthType:     authType,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
				TenantID:     tenantID,
			}
		}
		if sourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20 != nil {
			credentials = &shared.SourceMicrosoftTeamsUpdateAuthenticationMechanism{
				SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20: sourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20,
			}
		}
		var sourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoft *shared.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoft
		if r.Configuration.Credentials.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoft != nil {
			authType1 := new(shared.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftAuthType)
			if !r.Configuration.Credentials.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoft.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoft.AuthType.IsNull() {
				*authType1 = shared.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftAuthType(r.Configuration.Credentials.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoft.AuthType.ValueString())
			} else {
				authType1 = nil
			}
			clientId1 := r.Configuration.Credentials.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoft.ClientID.ValueString()
			clientSecret1 := r.Configuration.Credentials.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoft.ClientSecret.ValueString()
			tenantId1 := r.Configuration.Credentials.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoft.TenantID.ValueString()
			sourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoft = &shared.SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoft{
				AuthType:     authType1,
				ClientID:     clientId1,
				ClientSecret: clientSecret1,
				TenantID:     tenantId1,
			}
		}
		if sourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoft != nil {
			credentials = &shared.SourceMicrosoftTeamsUpdateAuthenticationMechanism{
				SourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoft: sourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoft,
			}
		}
	}
	period := r.Configuration.Period.ValueString()
	configuration := shared.SourceMicrosoftTeamsUpdate{
		Credentials: credentials,
		Period:      period,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMicrosoftTeamsPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMicrosoftTeamsResourceModel) ToDeleteSDKType() *shared.SourceMicrosoftTeamsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMicrosoftTeamsResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceMicrosoftTeamsResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
