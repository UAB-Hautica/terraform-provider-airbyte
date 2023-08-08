// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationDatabricks1 struct {
	AcceptTerms                   types.Bool                       `tfsdk:"accept_terms"`
	DataSource                    DestinationDatabricksDataSource2 `tfsdk:"data_source"`
	Database                      types.String                     `tfsdk:"database"`
	DatabricksHTTPPath            types.String                     `tfsdk:"databricks_http_path"`
	DatabricksPersonalAccessToken types.String                     `tfsdk:"databricks_personal_access_token"`
	DatabricksPort                types.String                     `tfsdk:"databricks_port"`
	DatabricksServerHostname      types.String                     `tfsdk:"databricks_server_hostname"`
	DestinationType               types.String                     `tfsdk:"destination_type"`
	PurgeStagingData              types.Bool                       `tfsdk:"purge_staging_data"`
	Schema                        types.String                     `tfsdk:"schema"`
}
