// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type DestinationWeaviate struct {
	Embedding  DestinationWeaviateEmbedding           `tfsdk:"embedding"`
	Indexing   DestinationWeaviateIndexing            `tfsdk:"indexing"`
	Processing DestinationMilvusProcessingConfigModel `tfsdk:"processing"`
}
