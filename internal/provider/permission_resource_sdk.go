// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *PermissionResourceModel) ToSharedPermissionCreateRequest() *shared.PermissionCreateRequest {
	organizationID := new(string)
	if !r.OrganizationID.IsUnknown() && !r.OrganizationID.IsNull() {
		*organizationID = r.OrganizationID.ValueString()
	} else {
		organizationID = nil
	}
	permissionType := shared.PermissionType(r.PermissionType.ValueString())
	userID := r.UserID.ValueString()
	workspaceID := new(string)
	if !r.WorkspaceID.IsUnknown() && !r.WorkspaceID.IsNull() {
		*workspaceID = r.WorkspaceID.ValueString()
	} else {
		workspaceID = nil
	}
	out := shared.PermissionCreateRequest{
		OrganizationID: organizationID,
		PermissionType: permissionType,
		UserID:         userID,
		WorkspaceID:    workspaceID,
	}
	return &out
}

func (r *PermissionResourceModel) RefreshFromSharedPermissionResponse(resp *shared.PermissionResponse) {
	if resp != nil {
		r.OrganizationID = types.StringPointerValue(resp.OrganizationID)
		r.PermissionID = types.StringValue(resp.PermissionID)
		r.PermissionType = types.StringValue(string(resp.PermissionType))
		r.UserID = types.StringValue(resp.UserID)
		r.WorkspaceID = types.StringPointerValue(resp.WorkspaceID)
	}
}

func (r *PermissionResourceModel) ToSharedPermissionUpdateRequest() *shared.PermissionUpdateRequest {
	permissionType := shared.PermissionType(r.PermissionType.ValueString())
	out := shared.PermissionUpdateRequest{
		PermissionType: permissionType,
	}
	return &out
}
