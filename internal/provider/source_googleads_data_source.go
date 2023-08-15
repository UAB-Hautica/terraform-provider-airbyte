// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"airbyte/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SourceGoogleAdsDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceGoogleAdsDataSource{}

func NewSourceGoogleAdsDataSource() datasource.DataSource {
	return &SourceGoogleAdsDataSource{}
}

// SourceGoogleAdsDataSource is the data source implementation.
type SourceGoogleAdsDataSource struct {
	client *sdk.SDK
}

// SourceGoogleAdsDataSourceModel describes the data model.
type SourceGoogleAdsDataSourceModel struct {
	Configuration SourceGoogleAds `tfsdk:"configuration"`
	Name          types.String    `tfsdk:"name"`
	SecretID      types.String    `tfsdk:"secret_id"`
	SourceID      types.String    `tfsdk:"source_id"`
	WorkspaceID   types.String    `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceGoogleAdsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_google_ads"
}

// Schema defines the schema for the data source.
func (r *SourceGoogleAdsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceGoogleAds DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"conversion_window_days": schema.Int64Attribute{
						Computed:    true,
						Description: `A conversion window is the period of time after an ad interaction (such as an ad click or video view) during which a conversion, such as a purchase, is recorded in Google Ads. For more information, see Google's <a href="https://support.google.com/google-ads/answer/3123169?hl=en">documentation</a>.`,
					},
					"credentials": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"access_token": schema.StringAttribute{
								Computed:    true,
								Description: `Access Token for making authenticated requests. More instruction on how to find this value in our <a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide">docs</a>`,
							},
							"client_id": schema.StringAttribute{
								Computed:    true,
								Description: `The Client ID of your Google Ads developer application. More instruction on how to find this value in our <a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide">docs</a>`,
							},
							"client_secret": schema.StringAttribute{
								Computed:    true,
								Description: `The Client Secret of your Google Ads developer application. More instruction on how to find this value in our <a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide">docs</a>`,
							},
							"developer_token": schema.StringAttribute{
								Computed:    true,
								Description: `Developer token granted by Google to use their APIs. More instruction on how to find this value in our <a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide">docs</a>`,
							},
							"refresh_token": schema.StringAttribute{
								Computed:    true,
								Description: `The token for obtaining a new access token. More instruction on how to find this value in our <a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide">docs</a>`,
							},
						},
					},
					"custom_queries": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"query": schema.StringAttribute{
									Computed:    true,
									Description: `A custom defined GAQL query for building the report. Should not contain segments.date expression because it is used by incremental streams. See Google's <a href="https://developers.google.com/google-ads/api/fields/v11/overview_query_builder">query builder</a> for more information.`,
								},
								"table_name": schema.StringAttribute{
									Computed:    true,
									Description: `The table name in your destination database for choosen query.`,
								},
							},
						},
					},
					"customer_id": schema.StringAttribute{
						Computed:    true,
						Description: `Comma separated list of (client) customer IDs. Each customer ID must be specified as a 10-digit number without dashes. More instruction on how to find this value in our <a href="https://docs.airbyte.com/integrations/sources/google-ads#setup-guide">docs</a>. Metrics streams like AdGroupAdReport cannot be requested for a manager account.`,
					},
					"end_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
						Description: `UTC date and time in the format 2017-01-25. Any data after this date will not be replicated.`,
					},
					"login_customer_id": schema.StringAttribute{
						Computed:    true,
						Description: `If your access to the customer account is through a manager account, this field is required and must be set to the customer ID of the manager account (10-digit number without dashes). More information about this field you can see <a href="https://developers.google.com/google-ads/api/docs/concepts/call-structure#cid">here</a>`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"google-ads",
							),
						},
						Description: `must be one of ["google-ads"]`,
					},
					"start_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
						Description: `UTC date and time in the format 2017-01-25. Any data before this date will not be replicated.`,
					},
				},
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"secret_id": schema.StringAttribute{
				Optional:    true,
				Description: `Optional secretID obtained through the public API OAuth redirect flow.`,
			},
			"source_id": schema.StringAttribute{
				Required: true,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *SourceGoogleAdsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SourceGoogleAdsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceGoogleAdsDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	sourceID := data.SourceID.ValueString()
	request := operations.GetSourceGoogleAdsRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceGoogleAds(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
