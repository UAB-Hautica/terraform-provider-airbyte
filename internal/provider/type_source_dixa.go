// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceDixa struct {
	APIToken  types.String `tfsdk:"api_token"`
	BatchSize types.Int64  `tfsdk:"batch_size"`
	StartDate types.String `tfsdk:"start_date"`
}
