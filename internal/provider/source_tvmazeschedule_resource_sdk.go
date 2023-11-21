// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceTvmazeScheduleResourceModel) ToCreateSDKType() *shared.SourceTvmazeScheduleCreateRequest {
	domesticScheduleCountryCode := r.Configuration.DomesticScheduleCountryCode.ValueString()
	endDate := new(string)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate = r.Configuration.EndDate.ValueString()
	} else {
		endDate = nil
	}
	startDate := r.Configuration.StartDate.ValueString()
	webScheduleCountryCode := new(string)
	if !r.Configuration.WebScheduleCountryCode.IsUnknown() && !r.Configuration.WebScheduleCountryCode.IsNull() {
		*webScheduleCountryCode = r.Configuration.WebScheduleCountryCode.ValueString()
	} else {
		webScheduleCountryCode = nil
	}
	configuration := shared.SourceTvmazeSchedule{
		DomesticScheduleCountryCode: domesticScheduleCountryCode,
		EndDate:                     endDate,
		StartDate:                   startDate,
		WebScheduleCountryCode:      webScheduleCountryCode,
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
	out := shared.SourceTvmazeScheduleCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTvmazeScheduleResourceModel) ToGetSDKType() *shared.SourceTvmazeScheduleCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceTvmazeScheduleResourceModel) ToUpdateSDKType() *shared.SourceTvmazeSchedulePutRequest {
	domesticScheduleCountryCode := r.Configuration.DomesticScheduleCountryCode.ValueString()
	endDate := new(string)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate = r.Configuration.EndDate.ValueString()
	} else {
		endDate = nil
	}
	startDate := r.Configuration.StartDate.ValueString()
	webScheduleCountryCode := new(string)
	if !r.Configuration.WebScheduleCountryCode.IsUnknown() && !r.Configuration.WebScheduleCountryCode.IsNull() {
		*webScheduleCountryCode = r.Configuration.WebScheduleCountryCode.ValueString()
	} else {
		webScheduleCountryCode = nil
	}
	configuration := shared.SourceTvmazeScheduleUpdate{
		DomesticScheduleCountryCode: domesticScheduleCountryCode,
		EndDate:                     endDate,
		StartDate:                   startDate,
		WebScheduleCountryCode:      webScheduleCountryCode,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceTvmazeSchedulePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTvmazeScheduleResourceModel) ToDeleteSDKType() *shared.SourceTvmazeScheduleCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceTvmazeScheduleResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceTvmazeScheduleResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
