/*
Infoblox GRID API

OpenAPI specification for Infoblox NIOS WAPI GRID objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grid

import (
	"encoding/json"
	"fmt"
)

// UpdateGridServicerestartRequestResponse - struct for UpdateGridServicerestartRequestResponse
type UpdateGridServicerestartRequestResponse struct {
	UpdateGridServicerestartRequestResponseAsObject *UpdateGridServicerestartRequestResponseAsObject
	String                                          *string
}

// UpdateGridServicerestartRequestResponseAsObjectAsUpdateGridServicerestartRequestResponse is a convenience function that returns UpdateGridServicerestartRequestResponseAsObject wrapped in UpdateGridServicerestartRequestResponse
func UpdateGridServicerestartRequestResponseAsObjectAsUpdateGridServicerestartRequestResponse(v *UpdateGridServicerestartRequestResponseAsObject) UpdateGridServicerestartRequestResponse {
	return UpdateGridServicerestartRequestResponse{
		UpdateGridServicerestartRequestResponseAsObject: v,
	}
}

// stringAsUpdateGridServicerestartRequestResponse is a convenience function that returns string wrapped in UpdateGridServicerestartRequestResponse
func StringAsUpdateGridServicerestartRequestResponse(v *string) UpdateGridServicerestartRequestResponse {
	return UpdateGridServicerestartRequestResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateGridServicerestartRequestResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateGridServicerestartRequestResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdateGridServicerestartRequestResponseAsObject)
	if err == nil {
		jsonUpdateGridServicerestartRequestResponseAsObject, _ := json.Marshal(dst.UpdateGridServicerestartRequestResponseAsObject)
		if string(jsonUpdateGridServicerestartRequestResponseAsObject) == "{}" { // empty struct
			dst.UpdateGridServicerestartRequestResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdateGridServicerestartRequestResponseAsObject = nil
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
		dst.UpdateGridServicerestartRequestResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateGridServicerestartRequestResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateGridServicerestartRequestResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateGridServicerestartRequestResponse) MarshalJSON() ([]byte, error) {
	if src.UpdateGridServicerestartRequestResponseAsObject != nil {
		return json.Marshal(&src.UpdateGridServicerestartRequestResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateGridServicerestartRequestResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateGridServicerestartRequestResponseAsObject != nil {
		return obj.UpdateGridServicerestartRequestResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdateGridServicerestartRequestResponse struct {
	value *UpdateGridServicerestartRequestResponse
	isSet bool
}

func (v NullableUpdateGridServicerestartRequestResponse) Get() *UpdateGridServicerestartRequestResponse {
	return v.value
}

func (v *NullableUpdateGridServicerestartRequestResponse) Set(val *UpdateGridServicerestartRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGridServicerestartRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGridServicerestartRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGridServicerestartRequestResponse(val *UpdateGridServicerestartRequestResponse) *NullableUpdateGridServicerestartRequestResponse {
	return &NullableUpdateGridServicerestartRequestResponse{value: val, isSet: true}
}

func (v NullableUpdateGridServicerestartRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGridServicerestartRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
