// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceGoogleAnalyticsDataAPI struct {
	Credentials         *SourceGoogleAnalyticsDataAPICredentials         `tfsdk:"credentials"`
	CustomReportsArray  []SourceGoogleAnalyticsDataAPICustomReportConfig `tfsdk:"custom_reports_array"`
	DateRangesStartDate types.String                                     `tfsdk:"date_ranges_start_date"`
	PropertyIds         []types.String                                   `tfsdk:"property_ids"`
	WindowInDays        types.Int64                                      `tfsdk:"window_in_days"`
}
