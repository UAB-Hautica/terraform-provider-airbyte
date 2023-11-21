// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationWeaviatePutRequest struct {
	Configuration DestinationWeaviateUpdate `json:"configuration"`
	Name          string                    `json:"name"`
	WorkspaceID   string                    `json:"workspaceId"`
}

func (o *DestinationWeaviatePutRequest) GetConfiguration() DestinationWeaviateUpdate {
	if o == nil {
		return DestinationWeaviateUpdate{}
	}
	return o.Configuration
}

func (o *DestinationWeaviatePutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationWeaviatePutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
