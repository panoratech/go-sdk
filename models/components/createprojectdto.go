// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CreateProjectDto struct {
	// The name of the project
	Name string `json:"name"`
	// The organization ID
	IDOrganization *string `json:"id_organization,omitempty"`
	// The user ID
	IDUser string `json:"id_user"`
}

func (o *CreateProjectDto) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateProjectDto) GetIDOrganization() *string {
	if o == nil {
		return nil
	}
	return o.IDOrganization
}

func (o *CreateProjectDto) GetIDUser() string {
	if o == nil {
		return ""
	}
	return o.IDUser
}
