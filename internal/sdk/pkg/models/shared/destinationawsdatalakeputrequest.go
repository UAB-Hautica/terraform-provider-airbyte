// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationAwsDatalakePutRequest struct {
	Configuration DestinationAwsDatalakeUpdate `json:"configuration"`
	Name          string                       `json:"name"`
	WorkspaceID   string                       `json:"workspaceId"`
}

func (o *DestinationAwsDatalakePutRequest) GetConfiguration() DestinationAwsDatalakeUpdate {
	if o == nil {
		return DestinationAwsDatalakeUpdate{}
	}
	return o.Configuration
}

func (o *DestinationAwsDatalakePutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationAwsDatalakePutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
