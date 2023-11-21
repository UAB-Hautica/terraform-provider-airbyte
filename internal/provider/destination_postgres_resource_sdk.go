// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationPostgresResourceModel) ToCreateSDKType() *shared.DestinationPostgresCreateRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	schema := new(string)
	if !r.Configuration.Schema.IsUnknown() && !r.Configuration.Schema.IsNull() {
		*schema = r.Configuration.Schema.ValueString()
	} else {
		schema = nil
	}
	var sslMode *shared.DestinationPostgresSSLModes
	if r.Configuration.SslMode != nil {
		var destinationPostgresDisable *shared.DestinationPostgresDisable
		if r.Configuration.SslMode.Disable != nil {
			destinationPostgresDisable = &shared.DestinationPostgresDisable{}
		}
		if destinationPostgresDisable != nil {
			sslMode = &shared.DestinationPostgresSSLModes{
				DestinationPostgresDisable: destinationPostgresDisable,
			}
		}
		var destinationPostgresAllow *shared.DestinationPostgresAllow
		if r.Configuration.SslMode.Allow != nil {
			destinationPostgresAllow = &shared.DestinationPostgresAllow{}
		}
		if destinationPostgresAllow != nil {
			sslMode = &shared.DestinationPostgresSSLModes{
				DestinationPostgresAllow: destinationPostgresAllow,
			}
		}
		var destinationPostgresPrefer *shared.DestinationPostgresPrefer
		if r.Configuration.SslMode.Prefer != nil {
			destinationPostgresPrefer = &shared.DestinationPostgresPrefer{}
		}
		if destinationPostgresPrefer != nil {
			sslMode = &shared.DestinationPostgresSSLModes{
				DestinationPostgresPrefer: destinationPostgresPrefer,
			}
		}
		var destinationPostgresRequire *shared.DestinationPostgresRequire
		if r.Configuration.SslMode.Require != nil {
			destinationPostgresRequire = &shared.DestinationPostgresRequire{}
		}
		if destinationPostgresRequire != nil {
			sslMode = &shared.DestinationPostgresSSLModes{
				DestinationPostgresRequire: destinationPostgresRequire,
			}
		}
		var destinationPostgresVerifyCa *shared.DestinationPostgresVerifyCa
		if r.Configuration.SslMode.VerifyCa != nil {
			caCertificate := r.Configuration.SslMode.VerifyCa.CaCertificate.ValueString()
			clientKeyPassword := new(string)
			if !r.Configuration.SslMode.VerifyCa.ClientKeyPassword.IsUnknown() && !r.Configuration.SslMode.VerifyCa.ClientKeyPassword.IsNull() {
				*clientKeyPassword = r.Configuration.SslMode.VerifyCa.ClientKeyPassword.ValueString()
			} else {
				clientKeyPassword = nil
			}
			destinationPostgresVerifyCa = &shared.DestinationPostgresVerifyCa{
				CaCertificate:     caCertificate,
				ClientKeyPassword: clientKeyPassword,
			}
		}
		if destinationPostgresVerifyCa != nil {
			sslMode = &shared.DestinationPostgresSSLModes{
				DestinationPostgresVerifyCa: destinationPostgresVerifyCa,
			}
		}
		var destinationPostgresVerifyFull *shared.DestinationPostgresVerifyFull
		if r.Configuration.SslMode.VerifyFull != nil {
			caCertificate1 := r.Configuration.SslMode.VerifyFull.CaCertificate.ValueString()
			clientCertificate := r.Configuration.SslMode.VerifyFull.ClientCertificate.ValueString()
			clientKey := r.Configuration.SslMode.VerifyFull.ClientKey.ValueString()
			clientKeyPassword1 := new(string)
			if !r.Configuration.SslMode.VerifyFull.ClientKeyPassword.IsUnknown() && !r.Configuration.SslMode.VerifyFull.ClientKeyPassword.IsNull() {
				*clientKeyPassword1 = r.Configuration.SslMode.VerifyFull.ClientKeyPassword.ValueString()
			} else {
				clientKeyPassword1 = nil
			}
			destinationPostgresVerifyFull = &shared.DestinationPostgresVerifyFull{
				CaCertificate:     caCertificate1,
				ClientCertificate: clientCertificate,
				ClientKey:         clientKey,
				ClientKeyPassword: clientKeyPassword1,
			}
		}
		if destinationPostgresVerifyFull != nil {
			sslMode = &shared.DestinationPostgresSSLModes{
				DestinationPostgresVerifyFull: destinationPostgresVerifyFull,
			}
		}
	}
	var tunnelMethod *shared.DestinationPostgresSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var destinationPostgresNoTunnel *shared.DestinationPostgresNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			destinationPostgresNoTunnel = &shared.DestinationPostgresNoTunnel{}
		}
		if destinationPostgresNoTunnel != nil {
			tunnelMethod = &shared.DestinationPostgresSSHTunnelMethod{
				DestinationPostgresNoTunnel: destinationPostgresNoTunnel,
			}
		}
		var destinationPostgresSSHKeyAuthentication *shared.DestinationPostgresSSHKeyAuthentication
		if r.Configuration.TunnelMethod.SSHKeyAuthentication != nil {
			sshKey := r.Configuration.TunnelMethod.SSHKeyAuthentication.SSHKey.ValueString()
			tunnelHost := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelHost.ValueString()
			tunnelPort := new(int64)
			if !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsNull() {
				*tunnelPort = r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort = nil
			}
			tunnelUser := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelUser.ValueString()
			destinationPostgresSSHKeyAuthentication = &shared.DestinationPostgresSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if destinationPostgresSSHKeyAuthentication != nil {
			tunnelMethod = &shared.DestinationPostgresSSHTunnelMethod{
				DestinationPostgresSSHKeyAuthentication: destinationPostgresSSHKeyAuthentication,
			}
		}
		var destinationPostgresPasswordAuthentication *shared.DestinationPostgresPasswordAuthentication
		if r.Configuration.TunnelMethod.PasswordAuthentication != nil {
			tunnelHost1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelHost.ValueString()
			tunnelPort1 := new(int64)
			if !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsNull() {
				*tunnelPort1 = r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort1 = nil
			}
			tunnelUser1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUser.ValueString()
			tunnelUserPassword := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUserPassword.ValueString()
			destinationPostgresPasswordAuthentication = &shared.DestinationPostgresPasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if destinationPostgresPasswordAuthentication != nil {
			tunnelMethod = &shared.DestinationPostgresSSHTunnelMethod{
				DestinationPostgresPasswordAuthentication: destinationPostgresPasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationPostgres{
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Password:      password,
		Port:          port,
		Schema:        schema,
		SslMode:       sslMode,
		TunnelMethod:  tunnelMethod,
		Username:      username,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationPostgresCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationPostgresResourceModel) ToGetSDKType() *shared.DestinationPostgresCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationPostgresResourceModel) ToUpdateSDKType() *shared.DestinationPostgresPutRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	schema := new(string)
	if !r.Configuration.Schema.IsUnknown() && !r.Configuration.Schema.IsNull() {
		*schema = r.Configuration.Schema.ValueString()
	} else {
		schema = nil
	}
	var sslMode *shared.SSLModes
	if r.Configuration.SslMode != nil {
		var disable *shared.Disable
		if r.Configuration.SslMode.Disable != nil {
			disable = &shared.Disable{}
		}
		if disable != nil {
			sslMode = &shared.SSLModes{
				Disable: disable,
			}
		}
		var allow *shared.Allow
		if r.Configuration.SslMode.Allow != nil {
			allow = &shared.Allow{}
		}
		if allow != nil {
			sslMode = &shared.SSLModes{
				Allow: allow,
			}
		}
		var prefer *shared.Prefer
		if r.Configuration.SslMode.Prefer != nil {
			prefer = &shared.Prefer{}
		}
		if prefer != nil {
			sslMode = &shared.SSLModes{
				Prefer: prefer,
			}
		}
		var require *shared.Require
		if r.Configuration.SslMode.Require != nil {
			require = &shared.Require{}
		}
		if require != nil {
			sslMode = &shared.SSLModes{
				Require: require,
			}
		}
		var verifyCa *shared.VerifyCa
		if r.Configuration.SslMode.VerifyCa != nil {
			caCertificate := r.Configuration.SslMode.VerifyCa.CaCertificate.ValueString()
			clientKeyPassword := new(string)
			if !r.Configuration.SslMode.VerifyCa.ClientKeyPassword.IsUnknown() && !r.Configuration.SslMode.VerifyCa.ClientKeyPassword.IsNull() {
				*clientKeyPassword = r.Configuration.SslMode.VerifyCa.ClientKeyPassword.ValueString()
			} else {
				clientKeyPassword = nil
			}
			verifyCa = &shared.VerifyCa{
				CaCertificate:     caCertificate,
				ClientKeyPassword: clientKeyPassword,
			}
		}
		if verifyCa != nil {
			sslMode = &shared.SSLModes{
				VerifyCa: verifyCa,
			}
		}
		var verifyFull *shared.VerifyFull
		if r.Configuration.SslMode.VerifyFull != nil {
			caCertificate1 := r.Configuration.SslMode.VerifyFull.CaCertificate.ValueString()
			clientCertificate := r.Configuration.SslMode.VerifyFull.ClientCertificate.ValueString()
			clientKey := r.Configuration.SslMode.VerifyFull.ClientKey.ValueString()
			clientKeyPassword1 := new(string)
			if !r.Configuration.SslMode.VerifyFull.ClientKeyPassword.IsUnknown() && !r.Configuration.SslMode.VerifyFull.ClientKeyPassword.IsNull() {
				*clientKeyPassword1 = r.Configuration.SslMode.VerifyFull.ClientKeyPassword.ValueString()
			} else {
				clientKeyPassword1 = nil
			}
			verifyFull = &shared.VerifyFull{
				CaCertificate:     caCertificate1,
				ClientCertificate: clientCertificate,
				ClientKey:         clientKey,
				ClientKeyPassword: clientKeyPassword1,
			}
		}
		if verifyFull != nil {
			sslMode = &shared.SSLModes{
				VerifyFull: verifyFull,
			}
		}
	}
	var tunnelMethod *shared.DestinationPostgresUpdateSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var destinationPostgresUpdateNoTunnel *shared.DestinationPostgresUpdateNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			destinationPostgresUpdateNoTunnel = &shared.DestinationPostgresUpdateNoTunnel{}
		}
		if destinationPostgresUpdateNoTunnel != nil {
			tunnelMethod = &shared.DestinationPostgresUpdateSSHTunnelMethod{
				DestinationPostgresUpdateNoTunnel: destinationPostgresUpdateNoTunnel,
			}
		}
		var destinationPostgresUpdateSSHKeyAuthentication *shared.DestinationPostgresUpdateSSHKeyAuthentication
		if r.Configuration.TunnelMethod.SSHKeyAuthentication != nil {
			sshKey := r.Configuration.TunnelMethod.SSHKeyAuthentication.SSHKey.ValueString()
			tunnelHost := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelHost.ValueString()
			tunnelPort := new(int64)
			if !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsNull() {
				*tunnelPort = r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort = nil
			}
			tunnelUser := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelUser.ValueString()
			destinationPostgresUpdateSSHKeyAuthentication = &shared.DestinationPostgresUpdateSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if destinationPostgresUpdateSSHKeyAuthentication != nil {
			tunnelMethod = &shared.DestinationPostgresUpdateSSHTunnelMethod{
				DestinationPostgresUpdateSSHKeyAuthentication: destinationPostgresUpdateSSHKeyAuthentication,
			}
		}
		var destinationPostgresUpdatePasswordAuthentication *shared.DestinationPostgresUpdatePasswordAuthentication
		if r.Configuration.TunnelMethod.PasswordAuthentication != nil {
			tunnelHost1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelHost.ValueString()
			tunnelPort1 := new(int64)
			if !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsNull() {
				*tunnelPort1 = r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort1 = nil
			}
			tunnelUser1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUser.ValueString()
			tunnelUserPassword := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUserPassword.ValueString()
			destinationPostgresUpdatePasswordAuthentication = &shared.DestinationPostgresUpdatePasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if destinationPostgresUpdatePasswordAuthentication != nil {
			tunnelMethod = &shared.DestinationPostgresUpdateSSHTunnelMethod{
				DestinationPostgresUpdatePasswordAuthentication: destinationPostgresUpdatePasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationPostgresUpdate{
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Password:      password,
		Port:          port,
		Schema:        schema,
		SslMode:       sslMode,
		TunnelMethod:  tunnelMethod,
		Username:      username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationPostgresPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationPostgresResourceModel) ToDeleteSDKType() *shared.DestinationPostgresCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationPostgresResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationPostgresResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
