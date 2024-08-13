// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourcePunkAPIUpdate struct {
	// To extract specific data with Unique ID
	BrewedAfter string `json:"brewed_after"`
	// To extract specific data with Unique ID
	BrewedBefore string `json:"brewed_before"`
	// To extract specific data with Unique ID
	ID *string `json:"id,omitempty"`
}

func (o *SourcePunkAPIUpdate) GetBrewedAfter() string {
	if o == nil {
		return ""
	}
	return o.BrewedAfter
}

func (o *SourcePunkAPIUpdate) GetBrewedBefore() string {
	if o == nil {
		return ""
	}
	return o.BrewedBefore
}

func (o *SourcePunkAPIUpdate) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
