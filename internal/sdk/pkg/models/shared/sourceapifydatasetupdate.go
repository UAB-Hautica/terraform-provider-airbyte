// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceApifyDatasetUpdate struct {
	// ID of the dataset you would like to load to Airbyte. In Apify Console, you can view your datasets in the <a href="https://console.apify.com/storage/datasets">Storage section under the Datasets tab</a> after you login. See the <a href="https://docs.apify.com/platform/storage/dataset">Apify Docs</a> for more information.
	DatasetID string `json:"dataset_id"`
	// Personal API token of your Apify account. In Apify Console, you can find your API token in the <a href="https://console.apify.com/account/integrations">Settings section under the Integrations tab</a> after you login. See the <a href="https://docs.apify.com/platform/integrations/api#api-token">Apify Docs</a> for more information.
	Token string `json:"token"`
}

func (o *SourceApifyDatasetUpdate) GetDatasetID() string {
	if o == nil {
		return ""
	}
	return o.DatasetID
}

func (o *SourceApifyDatasetUpdate) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}
