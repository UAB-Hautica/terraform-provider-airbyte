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
var _ datasource.DataSource = &SourceE2eTestCloudDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceE2eTestCloudDataSource{}

func NewSourceE2eTestCloudDataSource() datasource.DataSource {
	return &SourceE2eTestCloudDataSource{}
}

// SourceE2eTestCloudDataSource is the data source implementation.
type SourceE2eTestCloudDataSource struct {
	client *sdk.SDK
}

// SourceE2eTestCloudDataSourceModel describes the data model.
type SourceE2eTestCloudDataSourceModel struct {
	Configuration SourceE2eTestCloud `tfsdk:"configuration"`
	Name          types.String       `tfsdk:"name"`
	SecretID      types.String       `tfsdk:"secret_id"`
	SourceID      types.String       `tfsdk:"source_id"`
	WorkspaceID   types.String       `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceE2eTestCloudDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_e2e_test_cloud"
}

// Schema defines the schema for the data source.
func (r *SourceE2eTestCloudDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceE2eTestCloud DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"max_messages": schema.Int64Attribute{
						Computed:    true,
						Description: `Number of records to emit per stream. Min 1. Max 100 billion.`,
					},
					"message_interval_ms": schema.Int64Attribute{
						Computed:    true,
						Description: `Interval between messages in ms. Min 0 ms. Max 60000 ms (1 minute).`,
					},
					"mock_catalog": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_e2e_test_cloud_mock_catalog_multi_schema": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"stream_schemas": schema.StringAttribute{
										Computed:    true,
										Description: `A Json object specifying multiple data streams and their schemas. Each key in this object is one stream name. Each value is the schema for that stream. The schema should be compatible with <a href="https://json-schema.org/draft-07/json-schema-release-notes.html">draft-07</a>. See <a href="https://cswr.github.io/JsonSchema/spec/introduction/">this doc</a> for examples.`,
									},
									"type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"MULTI_STREAM",
											),
										},
										Description: `must be one of ["MULTI_STREAM"]`,
									},
								},
								Description: `A catalog with multiple data streams, each with a different schema.`,
							},
							"source_e2e_test_cloud_mock_catalog_single_schema": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"stream_duplication": schema.Int64Attribute{
										Computed:    true,
										Description: `Duplicate the stream for easy load testing. Each stream name will have a number suffix. For example, if the stream name is "ds", the duplicated streams will be "ds_0", "ds_1", etc.`,
									},
									"stream_name": schema.StringAttribute{
										Computed:    true,
										Description: `Name of the data stream.`,
									},
									"stream_schema": schema.StringAttribute{
										Computed:    true,
										Description: `A Json schema for the stream. The schema should be compatible with <a href="https://json-schema.org/draft-07/json-schema-release-notes.html">draft-07</a>. See <a href="https://cswr.github.io/JsonSchema/spec/introduction/">this doc</a> for examples.`,
									},
									"type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SINGLE_STREAM",
											),
										},
										Description: `must be one of ["SINGLE_STREAM"]`,
									},
								},
								Description: `A catalog with one or multiple streams that share the same schema.`,
							},
							"source_e2e_test_cloud_update_mock_catalog_multi_schema": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"stream_schemas": schema.StringAttribute{
										Computed:    true,
										Description: `A Json object specifying multiple data streams and their schemas. Each key in this object is one stream name. Each value is the schema for that stream. The schema should be compatible with <a href="https://json-schema.org/draft-07/json-schema-release-notes.html">draft-07</a>. See <a href="https://cswr.github.io/JsonSchema/spec/introduction/">this doc</a> for examples.`,
									},
									"type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"MULTI_STREAM",
											),
										},
										Description: `must be one of ["MULTI_STREAM"]`,
									},
								},
								Description: `A catalog with multiple data streams, each with a different schema.`,
							},
							"source_e2e_test_cloud_update_mock_catalog_single_schema": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"stream_duplication": schema.Int64Attribute{
										Computed:    true,
										Description: `Duplicate the stream for easy load testing. Each stream name will have a number suffix. For example, if the stream name is "ds", the duplicated streams will be "ds_0", "ds_1", etc.`,
									},
									"stream_name": schema.StringAttribute{
										Computed:    true,
										Description: `Name of the data stream.`,
									},
									"stream_schema": schema.StringAttribute{
										Computed:    true,
										Description: `A Json schema for the stream. The schema should be compatible with <a href="https://json-schema.org/draft-07/json-schema-release-notes.html">draft-07</a>. See <a href="https://cswr.github.io/JsonSchema/spec/introduction/">this doc</a> for examples.`,
									},
									"type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SINGLE_STREAM",
											),
										},
										Description: `must be one of ["SINGLE_STREAM"]`,
									},
								},
								Description: `A catalog with one or multiple streams that share the same schema.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"seed": schema.Int64Attribute{
						Computed:    true,
						Description: `When the seed is unspecified, the current time millis will be used as the seed. Range: [0, 1000000].`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"e2e-test-cloud",
							),
						},
						Description: `must be one of ["e2e-test-cloud"]`,
					},
					"type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"CONTINUOUS_FEED",
							),
						},
						Description: `must be one of ["CONTINUOUS_FEED"]`,
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

func (r *SourceE2eTestCloudDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceE2eTestCloudDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceE2eTestCloudDataSourceModel
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
	request := operations.GetSourceE2eTestCloudRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceE2eTestCloud(ctx, request)
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
	if res.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
