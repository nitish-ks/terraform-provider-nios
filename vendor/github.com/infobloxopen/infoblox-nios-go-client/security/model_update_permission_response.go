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

// UpdatePermissionResponse - struct for UpdatePermissionResponse
type UpdatePermissionResponse struct {
	UpdatePermissionResponseAsObject *UpdatePermissionResponseAsObject
	String                           *string
}

// UpdatePermissionResponseAsObjectAsUpdatePermissionResponse is a convenience function that returns UpdatePermissionResponseAsObject wrapped in UpdatePermissionResponse
func UpdatePermissionResponseAsObjectAsUpdatePermissionResponse(v *UpdatePermissionResponseAsObject) UpdatePermissionResponse {
	return UpdatePermissionResponse{
		UpdatePermissionResponseAsObject: v,
	}
}

// stringAsUpdatePermissionResponse is a convenience function that returns string wrapped in UpdatePermissionResponse
func StringAsUpdatePermissionResponse(v *string) UpdatePermissionResponse {
	return UpdatePermissionResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdatePermissionResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdatePermissionResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdatePermissionResponseAsObject)
	if err == nil {
		jsonUpdatePermissionResponseAsObject, _ := json.Marshal(dst.UpdatePermissionResponseAsObject)
		if string(jsonUpdatePermissionResponseAsObject) == "{}" { // empty struct
			dst.UpdatePermissionResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdatePermissionResponseAsObject = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.UpdatePermissionResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdatePermissionResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdatePermissionResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdatePermissionResponse) MarshalJSON() ([]byte, error) {
	if src.UpdatePermissionResponseAsObject != nil {
		return json.Marshal(&src.UpdatePermissionResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdatePermissionResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdatePermissionResponseAsObject != nil {
		return obj.UpdatePermissionResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdatePermissionResponse struct {
	value *UpdatePermissionResponse
	isSet bool
}

func (v NullableUpdatePermissionResponse) Get() *UpdatePermissionResponse {
	return v.value
}

func (v *NullableUpdatePermissionResponse) Set(val *UpdatePermissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePermissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePermissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePermissionResponse(val *UpdatePermissionResponse) *NullableUpdatePermissionResponse {
	return &NullableUpdatePermissionResponse{value: val, isSet: true}
}

func (v NullableUpdatePermissionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePermissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
