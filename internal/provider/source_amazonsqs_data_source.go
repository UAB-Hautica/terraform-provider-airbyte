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
var _ datasource.DataSource = &SourceAmazonSqsDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceAmazonSqsDataSource{}

func NewSourceAmazonSqsDataSource() datasource.DataSource {
	return &SourceAmazonSqsDataSource{}
}

// SourceAmazonSqsDataSource is the data source implementation.
type SourceAmazonSqsDataSource struct {
	client *sdk.SDK
}

// SourceAmazonSqsDataSourceModel describes the data model.
type SourceAmazonSqsDataSourceModel struct {
	Configuration SourceAmazonSqs `tfsdk:"configuration"`
	Name          types.String    `tfsdk:"name"`
	SecretID      types.String    `tfsdk:"secret_id"`
	SourceID      types.String    `tfsdk:"source_id"`
	WorkspaceID   types.String    `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceAmazonSqsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_amazon_sqs"
}

// Schema defines the schema for the data source.
func (r *SourceAmazonSqsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceAmazonSqs DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"access_key": schema.StringAttribute{
						Computed:    true,
						Description: `The Access Key ID of the AWS IAM Role to use for pulling messages`,
					},
					"attributes_to_return": schema.StringAttribute{
						Computed:    true,
						Description: `Comma separated list of Mesage Attribute names to return`,
					},
					"delete_messages": schema.BoolAttribute{
						Computed:    true,
						Description: `If Enabled, messages will be deleted from the SQS Queue after being read. If Disabled, messages are left in the queue and can be read more than once. WARNING: Enabling this option can result in data loss in cases of failure, use with caution, see documentation for more detail. `,
					},
					"max_batch_size": schema.Int64Attribute{
						Computed:    true,
						Description: `Max amount of messages to get in one batch (10 max)`,
					},
					"max_wait_time": schema.Int64Attribute{
						Computed:    true,
						Description: `Max amount of time in seconds to wait for messages in a single poll (20 max)`,
					},
					"queue_url": schema.StringAttribute{
						Computed:    true,
						Description: `URL of the SQS Queue`,
					},
					"region": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"us-east-1",
								"us-east-2",
								"us-west-1",
								"us-west-2",
								"af-south-1",
								"ap-east-1",
								"ap-south-1",
								"ap-northeast-1",
								"ap-northeast-2",
								"ap-northeast-3",
								"ap-southeast-1",
								"ap-southeast-2",
								"ca-central-1",
								"cn-north-1",
								"cn-northwest-1",
								"eu-central-1",
								"eu-north-1",
								"eu-south-1",
								"eu-west-1",
								"eu-west-2",
								"eu-west-3",
								"sa-east-1",
								"me-south-1",
								"us-gov-east-1",
								"us-gov-west-1",
							),
						},
						MarkdownDescription: `must be one of ["us-east-1", "us-east-2", "us-west-1", "us-west-2", "af-south-1", "ap-east-1", "ap-south-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-southeast-1", "ap-southeast-2", "ca-central-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-north-1", "eu-south-1", "eu-west-1", "eu-west-2", "eu-west-3", "sa-east-1", "me-south-1", "us-gov-east-1", "us-gov-west-1"]` + "\n" +
							`AWS Region of the SQS Queue`,
					},
					"secret_key": schema.StringAttribute{
						Computed:    true,
						Description: `The Secret Key of the AWS IAM Role to use for pulling messages`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"amazon-sqs",
							),
						},
						Description: `must be one of ["amazon-sqs"]`,
					},
					"visibility_timeout": schema.Int64Attribute{
						Computed:    true,
						Description: `Modify the Visibility Timeout of the individual message from the Queue's default (seconds).`,
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

func (r *SourceAmazonSqsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceAmazonSqsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceAmazonSqsDataSourceModel
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
	request := operations.GetSourceAmazonSqsRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceAmazonSqs(ctx, request)
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
