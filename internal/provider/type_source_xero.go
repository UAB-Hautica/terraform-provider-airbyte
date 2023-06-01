// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceXero struct {
	Authentication SourceXeroAuthenticateViaXeroOAuth `tfsdk:"authentication"`
	SourceType     types.String                       `tfsdk:"source_type"`
	StartDate      types.String                       `tfsdk:"start_date"`
	TenantID       types.String                       `tfsdk:"tenant_id"`
}
