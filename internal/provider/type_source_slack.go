// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceSlack struct {
	ChannelFilter  []types.String                      `tfsdk:"channel_filter"`
	Credentials    *SourceSlackAuthenticationMechanism `tfsdk:"credentials"`
	JoinChannels   types.Bool                          `tfsdk:"join_channels"`
	LookbackWindow types.Int64                         `tfsdk:"lookback_window"`
	StartDate      types.String                        `tfsdk:"start_date"`
}
