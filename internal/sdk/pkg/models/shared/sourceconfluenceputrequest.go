// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceConfluencePutRequest struct {
	Configuration SourceConfluenceUpdate `json:"configuration"`
	Name          string                 `json:"name"`
	WorkspaceID   string                 `json:"workspaceId"`
}

func (o *SourceConfluencePutRequest) GetConfiguration() SourceConfluenceUpdate {
	if o == nil {
		return SourceConfluenceUpdate{}
	}
	return o.Configuration
}

func (o *SourceConfluencePutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceConfluencePutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
