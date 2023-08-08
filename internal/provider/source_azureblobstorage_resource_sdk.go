// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceAzureBlobStorageResourceModel) ToCreateSDKType() *shared.SourceAzureBlobStorageCreateRequest {
	azureBlobStorageAccountKey := r.Configuration.AzureBlobStorageAccountKey.ValueString()
	azureBlobStorageAccountName := r.Configuration.AzureBlobStorageAccountName.ValueString()
	azureBlobStorageBlobsPrefix := new(string)
	if !r.Configuration.AzureBlobStorageBlobsPrefix.IsUnknown() && !r.Configuration.AzureBlobStorageBlobsPrefix.IsNull() {
		*azureBlobStorageBlobsPrefix = r.Configuration.AzureBlobStorageBlobsPrefix.ValueString()
	} else {
		azureBlobStorageBlobsPrefix = nil
	}
	azureBlobStorageContainerName := r.Configuration.AzureBlobStorageContainerName.ValueString()
	azureBlobStorageEndpoint := new(string)
	if !r.Configuration.AzureBlobStorageEndpoint.IsUnknown() && !r.Configuration.AzureBlobStorageEndpoint.IsNull() {
		*azureBlobStorageEndpoint = r.Configuration.AzureBlobStorageEndpoint.ValueString()
	} else {
		azureBlobStorageEndpoint = nil
	}
	azureBlobStorageSchemaInferenceLimit := new(int64)
	if !r.Configuration.AzureBlobStorageSchemaInferenceLimit.IsUnknown() && !r.Configuration.AzureBlobStorageSchemaInferenceLimit.IsNull() {
		*azureBlobStorageSchemaInferenceLimit = r.Configuration.AzureBlobStorageSchemaInferenceLimit.ValueInt64()
	} else {
		azureBlobStorageSchemaInferenceLimit = nil
	}
	var format shared.SourceAzureBlobStorageInputFormat
	var sourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON *shared.SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON
	if r.Configuration.Format.SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON != nil {
		formatType := shared.SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatType(r.Configuration.Format.SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON.FormatType.ValueString())
		sourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON = &shared.SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON{
			FormatType: formatType,
		}
	}
	if sourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON != nil {
		format = shared.SourceAzureBlobStorageInputFormat{
			SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON: sourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON,
		}
	}
	sourceType := shared.SourceAzureBlobStorageAzureBlobStorage(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceAzureBlobStorage{
		AzureBlobStorageAccountKey:           azureBlobStorageAccountKey,
		AzureBlobStorageAccountName:          azureBlobStorageAccountName,
		AzureBlobStorageBlobsPrefix:          azureBlobStorageBlobsPrefix,
		AzureBlobStorageContainerName:        azureBlobStorageContainerName,
		AzureBlobStorageEndpoint:             azureBlobStorageEndpoint,
		AzureBlobStorageSchemaInferenceLimit: azureBlobStorageSchemaInferenceLimit,
		Format:                               format,
		SourceType:                           sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAzureBlobStorageCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAzureBlobStorageResourceModel) ToGetSDKType() *shared.SourceAzureBlobStorageCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAzureBlobStorageResourceModel) ToUpdateSDKType() *shared.SourceAzureBlobStoragePutRequest {
	azureBlobStorageAccountKey := r.Configuration.AzureBlobStorageAccountKey.ValueString()
	azureBlobStorageAccountName := r.Configuration.AzureBlobStorageAccountName.ValueString()
	azureBlobStorageBlobsPrefix := new(string)
	if !r.Configuration.AzureBlobStorageBlobsPrefix.IsUnknown() && !r.Configuration.AzureBlobStorageBlobsPrefix.IsNull() {
		*azureBlobStorageBlobsPrefix = r.Configuration.AzureBlobStorageBlobsPrefix.ValueString()
	} else {
		azureBlobStorageBlobsPrefix = nil
	}
	azureBlobStorageContainerName := r.Configuration.AzureBlobStorageContainerName.ValueString()
	azureBlobStorageEndpoint := new(string)
	if !r.Configuration.AzureBlobStorageEndpoint.IsUnknown() && !r.Configuration.AzureBlobStorageEndpoint.IsNull() {
		*azureBlobStorageEndpoint = r.Configuration.AzureBlobStorageEndpoint.ValueString()
	} else {
		azureBlobStorageEndpoint = nil
	}
	azureBlobStorageSchemaInferenceLimit := new(int64)
	if !r.Configuration.AzureBlobStorageSchemaInferenceLimit.IsUnknown() && !r.Configuration.AzureBlobStorageSchemaInferenceLimit.IsNull() {
		*azureBlobStorageSchemaInferenceLimit = r.Configuration.AzureBlobStorageSchemaInferenceLimit.ValueInt64()
	} else {
		azureBlobStorageSchemaInferenceLimit = nil
	}
	var format shared.SourceAzureBlobStorageUpdateInputFormat
	var sourceAzureBlobStorageUpdateInputFormatJSONLinesNewlineDelimitedJSON *shared.SourceAzureBlobStorageUpdateInputFormatJSONLinesNewlineDelimitedJSON
	if r.Configuration.Format.SourceAzureBlobStorageUpdateInputFormatJSONLinesNewlineDelimitedJSON != nil {
		formatType := shared.SourceAzureBlobStorageUpdateInputFormatJSONLinesNewlineDelimitedJSONFormatType(r.Configuration.Format.SourceAzureBlobStorageUpdateInputFormatJSONLinesNewlineDelimitedJSON.FormatType.ValueString())
		sourceAzureBlobStorageUpdateInputFormatJSONLinesNewlineDelimitedJSON = &shared.SourceAzureBlobStorageUpdateInputFormatJSONLinesNewlineDelimitedJSON{
			FormatType: formatType,
		}
	}
	if sourceAzureBlobStorageUpdateInputFormatJSONLinesNewlineDelimitedJSON != nil {
		format = shared.SourceAzureBlobStorageUpdateInputFormat{
			SourceAzureBlobStorageUpdateInputFormatJSONLinesNewlineDelimitedJSON: sourceAzureBlobStorageUpdateInputFormatJSONLinesNewlineDelimitedJSON,
		}
	}
	configuration := shared.SourceAzureBlobStorageUpdate{
		AzureBlobStorageAccountKey:           azureBlobStorageAccountKey,
		AzureBlobStorageAccountName:          azureBlobStorageAccountName,
		AzureBlobStorageBlobsPrefix:          azureBlobStorageBlobsPrefix,
		AzureBlobStorageContainerName:        azureBlobStorageContainerName,
		AzureBlobStorageEndpoint:             azureBlobStorageEndpoint,
		AzureBlobStorageSchemaInferenceLimit: azureBlobStorageSchemaInferenceLimit,
		Format:                               format,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAzureBlobStoragePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAzureBlobStorageResourceModel) ToDeleteSDKType() *shared.SourceAzureBlobStorageCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAzureBlobStorageResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceAzureBlobStorageResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
