// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

type UsersType string

const (
	UsersTypeStr                          UsersType = "str"
	UsersTypeUnifiedFilestorageUserOutput UsersType = "UnifiedFilestorageUserOutput"
)

type Users struct {
	Str                          *string
	UnifiedFilestorageUserOutput *UnifiedFilestorageUserOutput

	Type UsersType
}

func CreateUsersStr(str string) Users {
	typ := UsersTypeStr

	return Users{
		Str:  &str,
		Type: typ,
	}
}

func CreateUsersUnifiedFilestorageUserOutput(unifiedFilestorageUserOutput UnifiedFilestorageUserOutput) Users {
	typ := UsersTypeUnifiedFilestorageUserOutput

	return Users{
		UnifiedFilestorageUserOutput: &unifiedFilestorageUserOutput,
		Type:                         typ,
	}
}

func (u *Users) UnmarshalJSON(data []byte) error {

	var unifiedFilestorageUserOutput UnifiedFilestorageUserOutput = UnifiedFilestorageUserOutput{}
	if err := utils.UnmarshalJSON(data, &unifiedFilestorageUserOutput, "", true, true); err == nil {
		u.UnifiedFilestorageUserOutput = &unifiedFilestorageUserOutput
		u.Type = UsersTypeUnifiedFilestorageUserOutput
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = UsersTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Users", string(data))
}

func (u Users) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.UnifiedFilestorageUserOutput != nil {
		return utils.MarshalJSON(u.UnifiedFilestorageUserOutput, "", true)
	}

	return nil, errors.New("could not marshal union type Users: all fields are null")
}

type UnifiedFilestorageGroupOutput struct {
	// The name of the group
	Name *string `json:"name"`
	// Uuids of users of the group
	Users []Users `json:"users"`
	// Indicates whether or not this object has been deleted in the third party platform.
	RemoteWasDeleted *bool `json:"remote_was_deleted"`
	// The custom field mappings of the object between the remote 3rd party & Panora
	FieldMappings map[string]any `json:"field_mappings,omitempty"`
	// The UUID of the group
	ID *string `json:"id,omitempty"`
	// The id of the group in the context of the 3rd Party
	RemoteID *string `json:"remote_id,omitempty"`
	// The remote data of the group in the context of the 3rd Party
	RemoteData map[string]any `json:"remote_data,omitempty"`
	// The created date of the object
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The modified date of the object
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
}

func (u UnifiedFilestorageGroupOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedFilestorageGroupOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedFilestorageGroupOutput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UnifiedFilestorageGroupOutput) GetUsers() []Users {
	if o == nil {
		return []Users{}
	}
	return o.Users
}

func (o *UnifiedFilestorageGroupOutput) GetRemoteWasDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteWasDeleted
}

func (o *UnifiedFilestorageGroupOutput) GetFieldMappings() map[string]any {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}

func (o *UnifiedFilestorageGroupOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedFilestorageGroupOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedFilestorageGroupOutput) GetRemoteData() map[string]any {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *UnifiedFilestorageGroupOutput) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UnifiedFilestorageGroupOutput) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}
