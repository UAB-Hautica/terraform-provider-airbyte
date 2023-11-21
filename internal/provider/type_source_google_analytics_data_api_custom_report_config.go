// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceGoogleAnalyticsDataAPICustomReportConfig struct {
	DimensionFilter *SourceGoogleAnalyticsDataAPIDimensionsFilter `tfsdk:"dimension_filter"`
	Dimensions      []types.String                                `tfsdk:"dimensions"`
	MetricFilter    *SourceGoogleAnalyticsDataAPIDimensionsFilter `tfsdk:"metric_filter"`
	Metrics         []types.String                                `tfsdk:"metrics"`
	Name            types.String                                  `tfsdk:"name"`
}
