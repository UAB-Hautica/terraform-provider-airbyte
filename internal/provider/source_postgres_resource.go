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
var _ resource.Resource = &SourcePostgresResource{}
var _ resource.ResourceWithImportState = &SourcePostgresResource{}

func NewSourcePostgresResource() resource.Resource {
	return &SourcePostgresResource{}
}

// SourcePostgresResource defines the resource implementation.
type SourcePostgresResource struct {
	client *sdk.SDK
}

// SourcePostgresResourceModel describes the resource data model.
type SourcePostgresResourceModel struct {
	Configuration SourcePostgres `tfsdk:"configuration"`
	Name          types.String   `tfsdk:"name"`
	SecretID      types.String   `tfsdk:"secret_id"`
	SourceID      types.String   `tfsdk:"source_id"`
	SourceType    types.String   `tfsdk:"source_type"`
	WorkspaceID   types.String   `tfsdk:"workspace_id"`
}

func (r *SourcePostgresResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_postgres"
}

func (r *SourcePostgresResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourcePostgres Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"database": schema.StringAttribute{
						Required: true,
					},
					"host": schema.StringAttribute{
						Required: true,
					},
					"jdbc_url_params": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"password": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"port": schema.Int64Attribute{
						Required: true,
					},
					"replication_method": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"source_postgres_replication_method_logical_replication_cdc_": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"initial_waiting_seconds": schema.Int64Attribute{
										Computed: true,
										Optional: true,
									},
									"lsn_commit_behaviour": schema.StringAttribute{
										Computed: true,
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"While reading Data",
												"After loading Data in the destination",
											),
										},
										Description: `Determines when Airbtye should flush the LSN of processed WAL logs in the source database. ` + "`" + `After loading Data in the destination` + "`" + ` is default. If ` + "`" + `While reading Data` + "`" + ` is selected, in case of a downstream failure (while loading data into the destination), next sync would result in a full sync.`,
									},
									"method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"CDC",
											),
										},
									},
									"plugin": schema.StringAttribute{
										Computed: true,
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"pgoutput",
											),
										},
										Description: `A logical decoding plugin installed on the PostgreSQL server.`,
									},
									"publication": schema.StringAttribute{
										Required: true,
									},
									"queue_size": schema.Int64Attribute{
										Computed: true,
										Optional: true,
									},
									"replication_slot": schema.StringAttribute{
										Required: true,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Logical replication uses the Postgres write-ahead log (WAL) to detect inserts, updates, and deletes. This needs to be configured on the source database itself. Only available on Postgres 10 and above. Read the <a href="https://docs.airbyte.com/integrations/sources/postgres">docs</a>.`,
							},
							"source_postgres_replication_method_standard": schema.SingleNestedAttribute{
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
								Description: `Standard replication requires no setup on the DB side but will not be able to represent deletions incrementally.`,
							},
							"source_postgres_update_replication_method_logical_replication_cdc_": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"initial_waiting_seconds": schema.Int64Attribute{
										Computed: true,
										Optional: true,
									},
									"lsn_commit_behaviour": schema.StringAttribute{
										Computed: true,
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"While reading Data",
												"After loading Data in the destination",
											),
										},
										Description: `Determines when Airbtye should flush the LSN of processed WAL logs in the source database. ` + "`" + `After loading Data in the destination` + "`" + ` is default. If ` + "`" + `While reading Data` + "`" + ` is selected, in case of a downstream failure (while loading data into the destination), next sync would result in a full sync.`,
									},
									"method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"CDC",
											),
										},
									},
									"plugin": schema.StringAttribute{
										Computed: true,
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"pgoutput",
											),
										},
										Description: `A logical decoding plugin installed on the PostgreSQL server.`,
									},
									"publication": schema.StringAttribute{
										Required: true,
									},
									"queue_size": schema.Int64Attribute{
										Computed: true,
										Optional: true,
									},
									"replication_slot": schema.StringAttribute{
										Required: true,
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Logical replication uses the Postgres write-ahead log (WAL) to detect inserts, updates, and deletes. This needs to be configured on the source database itself. Only available on Postgres 10 and above. Read the <a href="https://docs.airbyte.com/integrations/sources/postgres">docs</a>.`,
							},
							"source_postgres_update_replication_method_standard": schema.SingleNestedAttribute{
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
								Description: `Standard replication requires no setup on the DB side but will not be able to represent deletions incrementally.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"schemas": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"source_type": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"postgres",
							),
						},
					},
					"ssl_mode": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"source_postgres_ssl_modes_allow": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"allow",
											),
										},
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Enables encryption only when required by the source database.`,
							},
							"source_postgres_ssl_modes_disable": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"disable",
											),
										},
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Disables encryption of communication between Airbyte and source database.`,
							},
							"source_postgres_ssl_modes_prefer": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"prefer",
											),
										},
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Allows unencrypted connection only if the source database does not support encryption.`,
							},
							"source_postgres_ssl_modes_require": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"require",
											),
										},
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Always require encryption. If the source database server does not support encryption, connection will fail.`,
							},
							"source_postgres_ssl_modes_verify_ca": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ca_certificate": schema.StringAttribute{
										Required: true,
									},
									"client_certificate": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"client_key": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"client_key_password": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"mode": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"verify-ca",
											),
										},
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Always require encryption and verifies that the source database server has a valid SSL certificate.`,
							},
							"source_postgres_ssl_modes_verify_full": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ca_certificate": schema.StringAttribute{
										Required: true,
									},
									"client_certificate": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"client_key": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"client_key_password": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"mode": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"verify-full",
											),
										},
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `This is the most secure mode. Always require encryption and verifies the identity of the source database server.`,
							},
							"source_postgres_update_ssl_modes_allow": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"allow",
											),
										},
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Enables encryption only when required by the source database.`,
							},
							"source_postgres_update_ssl_modes_disable": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"disable",
											),
										},
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Disables encryption of communication between Airbyte and source database.`,
							},
							"source_postgres_update_ssl_modes_prefer": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"prefer",
											),
										},
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Allows unencrypted connection only if the source database does not support encryption.`,
							},
							"source_postgres_update_ssl_modes_require": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"require",
											),
										},
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Always require encryption. If the source database server does not support encryption, connection will fail.`,
							},
							"source_postgres_update_ssl_modes_verify_ca": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ca_certificate": schema.StringAttribute{
										Required: true,
									},
									"client_certificate": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"client_key": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"client_key_password": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"mode": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"verify-ca",
											),
										},
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Always require encryption and verifies that the source database server has a valid SSL certificate.`,
							},
							"source_postgres_update_ssl_modes_verify_full": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ca_certificate": schema.StringAttribute{
										Required: true,
									},
									"client_certificate": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"client_key": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"client_key_password": schema.StringAttribute{
										Computed: true,
										Optional: true,
									},
									"mode": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"verify-full",
											),
										},
									},
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `This is the most secure mode. Always require encryption and verifies the identity of the source database server.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"tunnel_method": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"source_postgres_ssh_tunnel_method_no_tunnel": schema.SingleNestedAttribute{
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
							"source_postgres_ssh_tunnel_method_password_authentication": schema.SingleNestedAttribute{
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
							"source_postgres_ssh_tunnel_method_ssh_key_authentication": schema.SingleNestedAttribute{
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
							"source_postgres_update_ssh_tunnel_method_no_tunnel": schema.SingleNestedAttribute{
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
							"source_postgres_update_ssh_tunnel_method_password_authentication": schema.SingleNestedAttribute{
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
							"source_postgres_update_ssh_tunnel_method_ssh_key_authentication": schema.SingleNestedAttribute{
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
					"username": schema.StringAttribute{
						Required: true,
					},
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"secret_id": schema.StringAttribute{
				Optional: true,
			},
			"source_id": schema.StringAttribute{
				Computed: true,
			},
			"source_type": schema.StringAttribute{
				Computed: true,
			},
			"workspace_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *SourcePostgresResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SourcePostgresResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourcePostgresResourceModel
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
	res, err := r.client.Sources.CreateSourcePostgres(ctx, request)
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
	data.RefreshFromCreateResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourcePostgresResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourcePostgresResourceModel
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

func (r *SourcePostgresResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourcePostgresResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourcePostgresPutRequest := data.ToUpdateSDKType()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourcePostgresRequest{
		SourcePostgresPutRequest: sourcePostgresPutRequest,
		SourceID:                 sourceID,
	}
	res, err := r.client.Sources.PutSourcePostgres(ctx, request)
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

func (r *SourcePostgresResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourcePostgresResourceModel
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

	sourceID := data.SourceID.ValueString()
	request := operations.DeleteSourcePostgresRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourcePostgres(ctx, request)
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

func (r *SourcePostgresResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
