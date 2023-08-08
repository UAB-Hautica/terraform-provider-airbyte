// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"airbyte/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &DestinationPubsubDataSource{}
var _ datasource.DataSourceWithConfigure = &DestinationPubsubDataSource{}

func NewDestinationPubsubDataSource() datasource.DataSource {
	return &DestinationPubsubDataSource{}
}

// DestinationPubsubDataSource is the data source implementation.
type DestinationPubsubDataSource struct {
	client *sdk.SDK
}

// DestinationPubsubDataSourceModel describes the data model.
type DestinationPubsubDataSourceModel struct {
	Configuration DestinationPubsub `tfsdk:"configuration"`
	DestinationID types.String      `tfsdk:"destination_id"`
	Name          types.String      `tfsdk:"name"`
	WorkspaceID   types.String      `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *DestinationPubsubDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_pubsub"
}

// Schema defines the schema for the data source.
func (r *DestinationPubsubDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationPubsub DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"batching_delay_threshold": schema.Int64Attribute{
						Computed:    true,
						Description: `Number of ms before the buffer is flushed`,
					},
					"batching_element_count_threshold": schema.Int64Attribute{
						Computed:    true,
						Description: `Number of messages before the buffer is flushed`,
					},
					"batching_enabled": schema.BoolAttribute{
						Computed:    true,
						Description: `If TRUE messages will be buffered instead of sending them one by one`,
					},
					"batching_request_bytes_threshold": schema.Int64Attribute{
						Computed:    true,
						Description: `Number of bytes before the buffer is flushed`,
					},
					"credentials_json": schema.StringAttribute{
						Computed:    true,
						Description: `The contents of the JSON service account key. Check out the <a href="https://docs.airbyte.com/integrations/destinations/pubsub">docs</a> if you need help generating this key.`,
					},
					"destination_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"pubsub",
							),
						},
						Description: `must be one of ["pubsub"]`,
					},
					"ordering_enabled": schema.BoolAttribute{
						Computed:    true,
						Description: `If TRUE PubSub publisher will have <a href="https://cloud.google.com/pubsub/docs/ordering">message ordering</a> enabled. Every message will have an ordering key of stream`,
					},
					"project_id": schema.StringAttribute{
						Computed:    true,
						Description: `The GCP project ID for the project containing the target PubSub.`,
					},
					"topic_id": schema.StringAttribute{
						Computed:    true,
						Description: `The PubSub topic ID in the given GCP project ID.`,
					},
				},
			},
			"destination_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *DestinationPubsubDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *DestinationPubsubDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *DestinationPubsubDataSourceModel
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

	destinationID := data.DestinationID.ValueString()
	request := operations.GetDestinationPubsubRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationPubsub(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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
	if res.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
