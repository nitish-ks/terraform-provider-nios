/*
Infoblox SECURITY API

OpenAPI specification for Infoblox NIOS WAPI SECURITY objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package security

import (
	"encoding/json"
	"fmt"
)

// ListAdminuserResponse - struct for ListAdminuserResponse
type ListAdminuserResponse struct {
	ListAdminuserResponseObject *ListAdminuserResponseObject
	ArrayOfAdminuser            *[]Adminuser
}

// ListAdminuserResponseObjectAsListAdminuserResponse is a convenience function that returns ListAdminuserResponseObject wrapped in ListAdminuserResponse
func ListAdminuserResponseObjectAsListAdminuserResponse(v *ListAdminuserResponseObject) ListAdminuserResponse {
	return ListAdminuserResponse{
		ListAdminuserResponseObject: v,
	}
}

// []AdminuserAsListAdminuserResponse is a convenience function that returns []Adminuser wrapped in ListAdminuserResponse
func ArrayOfAdminuserAsListAdminuserResponse(v *[]Adminuser) ListAdminuserResponse {
	return ListAdminuserResponse{
		ArrayOfAdminuser: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListAdminuserResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListAdminuserResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListAdminuserResponseObject)
	if err == nil {
		jsonListAdminuserResponseObject, _ := json.Marshal(dst.ListAdminuserResponseObject)
		if string(jsonListAdminuserResponseObject) == "{}" { // empty struct
			dst.ListAdminuserResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListAdminuserResponseObject = nil
	}

	// try to unmarshal data into ArrayOfAdminuser
	err = newStrictDecoder(data).Decode(&dst.ArrayOfAdminuser)
	if err == nil {
		jsonArrayOfAdminuser, _ := json.Marshal(dst.ArrayOfAdminuser)
		if string(jsonArrayOfAdminuser) == "{}" { // empty struct
			dst.ArrayOfAdminuser = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfAdminuser = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListAdminuserResponseObject = nil
		dst.ArrayOfAdminuser = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListAdminuserResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListAdminuserResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListAdminuserResponse) MarshalJSON() ([]byte, error) {
	if src.ListAdminuserResponseObject != nil {
		return json.Marshal(&src.ListAdminuserResponseObject)
	}

	if src.ArrayOfAdminuser != nil {
		return json.Marshal(&src.ArrayOfAdminuser)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListAdminuserResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListAdminuserResponseObject != nil {
		return obj.ListAdminuserResponseObject
	}

	if obj.ArrayOfAdminuser != nil {
		return obj.ArrayOfAdminuser
	}

	// all schemas are nil
	return nil
}

type NullableListAdminuserResponse struct {
	value *ListAdminuserResponse
	isSet bool
}

func (v NullableListAdminuserResponse) Get() *ListAdminuserResponse {
	return v.value
}

func (v *NullableListAdminuserResponse) Set(val *ListAdminuserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAdminuserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAdminuserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAdminuserResponse(val *ListAdminuserResponse) *NullableListAdminuserResponse {
	return &NullableListAdminuserResponse{value: val, isSet: true}
}

func (v NullableListAdminuserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAdminuserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
