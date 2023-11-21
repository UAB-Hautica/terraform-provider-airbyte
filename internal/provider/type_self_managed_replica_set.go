// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SelfManagedReplicaSet struct {
	AdditionalProperties types.String `tfsdk:"additional_properties"`
	AuthSource           types.String `tfsdk:"auth_source"`
	ConnectionString     types.String `tfsdk:"connection_string"`
	Database             types.String `tfsdk:"database"`
	Password             types.String `tfsdk:"password"`
	Username             types.String `tfsdk:"username"`
}
