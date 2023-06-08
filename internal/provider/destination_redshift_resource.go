// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"context"
	"fmt"

	"airbyte/internal/sdk/pkg/models/operations"
	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DestinationRedshiftResource{}
var _ resource.ResourceWithImportState = &DestinationRedshiftResource{}

func NewDestinationRedshiftResource() resource.Resource {
	return &DestinationRedshiftResource{}
}

// DestinationRedshiftResource defines the resource implementation.
type DestinationRedshiftResource struct {
	client *sdk.SDK
}

// DestinationRedshiftResourceModel describes the resource data model.
type DestinationRedshiftResourceModel struct {
	Configuration   DestinationRedshift `tfsdk:"configuration"`
	DestinationID   types.String        `tfsdk:"destination_id"`
	DestinationType types.String        `tfsdk:"destination_type"`
	Name            types.String        `tfsdk:"name"`
	WorkspaceID     types.String        `tfsdk:"workspace_id"`
}

func (r *DestinationRedshiftResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_redshift"
}

func (r *DestinationRedshiftResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationRedshift Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"database": schema.StringAttribute{
						Required: true,
					},
					"destination_type": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"redshift",
							),
						},
					},
					"host": schema.StringAttribute{
						Required: true,
					},
					"jdbc_url_params": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"password": schema.StringAttribute{
						Required: true,
					},
					"port": schema.Int64Attribute{
						Required: true,
					},
					"schema": schema.StringAttribute{
						Required: true,
					},
					"tunnel_method": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"destination_redshift_ssh_tunnel_method_no_tunnel": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"NO_TUNNEL",
											),
										},
										Description: `No ssh tunnel needed to connect to database`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"destination_redshift_ssh_tunnel_method_password_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_host": schema.StringAttribute{
										Required: true,
									},
									"tunnel_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_PASSWORD_AUTH",
											),
										},
										Description: `Connect through a jump server tunnel host using username and password authentication`,
									},
									"tunnel_port": schema.Int64Attribute{
										Required: true,
									},
									"tunnel_user": schema.StringAttribute{
										Required: true,
									},
									"tunnel_user_password": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"destination_redshift_ssh_tunnel_method_ssh_key_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ssh_key": schema.StringAttribute{
										Required: true,
									},
									"tunnel_host": schema.StringAttribute{
										Required: true,
									},
									"tunnel_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_KEY_AUTH",
											),
										},
										Description: `Connect through a jump server tunnel host using username and ssh key`,
									},
									"tunnel_port": schema.Int64Attribute{
										Required: true,
									},
									"tunnel_user": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"destination_redshift_update_ssh_tunnel_method_no_tunnel": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"NO_TUNNEL",
											),
										},
										Description: `No ssh tunnel needed to connect to database`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"destination_redshift_update_ssh_tunnel_method_password_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_host": schema.StringAttribute{
										Required: true,
									},
									"tunnel_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_PASSWORD_AUTH",
											),
										},
										Description: `Connect through a jump server tunnel host using username and password authentication`,
									},
									"tunnel_port": schema.Int64Attribute{
										Required: true,
									},
									"tunnel_user": schema.StringAttribute{
										Required: true,
									},
									"tunnel_user_password": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"destination_redshift_update_ssh_tunnel_method_ssh_key_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ssh_key": schema.StringAttribute{
										Required: true,
									},
									"tunnel_host": schema.StringAttribute{
										Required: true,
									},
									"tunnel_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_KEY_AUTH",
											),
										},
										Description: `Connect through a jump server tunnel host using username and ssh key`,
									},
									"tunnel_port": schema.Int64Attribute{
										Required: true,
									},
									"tunnel_user": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"uploading_method": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"destination_redshift_uploading_method_s3_staging": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_key_id": schema.StringAttribute{
										Required: true,
									},
									"encryption": schema.SingleNestedAttribute{
										Computed: true,
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"destination_redshift_uploading_method_s3_staging_encryption_aes_cbc_envelope_encryption": schema.SingleNestedAttribute{
												Computed: true,
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"encryption_type": schema.StringAttribute{
														Required: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"aes_cbc_envelope",
															),
														},
													},
													"key_encrypting_key": schema.StringAttribute{
														Computed: true,
														Optional: true,
													},
												},
												Description: `Staging data will be encrypted using AES-CBC envelope encryption.`,
											},
											"destination_redshift_uploading_method_s3_staging_encryption_no_encryption": schema.SingleNestedAttribute{
												Computed: true,
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"encryption_type": schema.StringAttribute{
														Required: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"none",
															),
														},
													},
												},
												Description: `Staging data will be stored in plaintext.`,
											},
										},
										Validators: []validator.Object{
											validators.ExactlyOneChild(),
										},
									},
									"file_buffer_count": schema.Int64Attribute{
										Computed: true,
										Optional: true,
									},
									"file_name_pattern": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"S3 Staging",
											),
										},
									},
									"purge_staging_data": schema.BoolAttribute{
										Computed: true,
										Optional: true,
									},
									"s3_bucket_name": schema.StringAttribute{
										Required: true,
									},
									"s3_bucket_path": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"s3_bucket_region": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"",
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
											),
										},
										Description: `The region of the S3 staging bucket to use if utilising a COPY strategy. See <a href="https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-bucket.html#:~:text=In-,Region,-%2C%20choose%20the%20AWS">AWS docs</a> for details.`,
									},
									"secret_access_key": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `The method how the data will be uploaded to the database.`,
							},
							"destination_redshift_uploading_method_standard": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Standard",
											),
										},
									},
								},
								Description: `The method how the data will be uploaded to the database.`,
							},
							"destination_redshift_update_uploading_method_s3_staging": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"access_key_id": schema.StringAttribute{
										Required: true,
									},
									"encryption": schema.SingleNestedAttribute{
										Computed: true,
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"destination_redshift_update_uploading_method_s3_staging_encryption_no_encryption": schema.SingleNestedAttribute{
												Computed: true,
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"encryption_type": schema.StringAttribute{
														Required: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"none",
															),
														},
													},
												},
												Description: `Staging data will be stored in plaintext.`,
											},
											"destination_redshift_update_uploading_method_s3_staging_encryption_aes_cbc_envelope_encryption": schema.SingleNestedAttribute{
												Computed: true,
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"encryption_type": schema.StringAttribute{
														Required: true,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"aes_cbc_envelope",
															),
														},
													},
													"key_encrypting_key": schema.StringAttribute{
														Computed: true,
														Optional: true,
													},
												},
												Description: `Staging data will be encrypted using AES-CBC envelope encryption.`,
											},
										},
										Validators: []validator.Object{
											validators.ExactlyOneChild(),
										},
									},
									"file_buffer_count": schema.Int64Attribute{
										Computed: true,
										Optional: true,
									},
									"file_name_pattern": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"S3 Staging",
											),
										},
									},
									"purge_staging_data": schema.BoolAttribute{
										Computed: true,
										Optional: true,
									},
									"s3_bucket_name": schema.StringAttribute{
										Required: true,
									},
									"s3_bucket_path": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"s3_bucket_region": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"",
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
											),
										},
										Description: `The region of the S3 staging bucket to use if utilising a COPY strategy. See <a href="https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-bucket.html#:~:text=In-,Region,-%2C%20choose%20the%20AWS">AWS docs</a> for details.`,
									},
									"secret_access_key": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `The method how the data will be uploaded to the database.`,
							},
							"destination_redshift_update_uploading_method_standard": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Standard",
											),
										},
									},
								},
								Description: `The method how the data will be uploaded to the database.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"username": schema.StringAttribute{
						Required: true,
					},
				},
			},
			"destination_id": schema.StringAttribute{
				Computed: true,
			},
			"destination_type": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"workspace_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *DestinationRedshiftResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *DestinationRedshiftResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DestinationRedshiftResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
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

	request := *data.ToCreateSDKType()
	res, err := r.client.Destinations.CreateDestinationRedshift(ctx, request)
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
	data.RefreshFromCreateResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationRedshiftResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DestinationRedshiftResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationRedshiftResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DestinationRedshiftResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	destinationRedshiftPutRequest := data.ToUpdateSDKType()
	destinationID := data.DestinationID.ValueString()
	request := operations.PutDestinationRedshiftRequest{
		DestinationRedshiftPutRequest: destinationRedshiftPutRequest,
		DestinationID:                 destinationID,
	}
	res, err := r.client.Destinations.PutDestinationRedshift(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationRedshiftResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DestinationRedshiftResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	request := operations.DeleteDestinationRedshiftRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.DeleteDestinationRedshift(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *DestinationRedshiftResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
