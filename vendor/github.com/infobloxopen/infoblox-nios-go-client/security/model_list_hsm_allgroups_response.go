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

// ListHsmAllgroupsResponse - struct for ListHsmAllgroupsResponse
type ListHsmAllgroupsResponse struct {
	ListHsmAllgroupsResponseObject *ListHsmAllgroupsResponseObject
	ArrayOfHsmAllgroups            *[]HsmAllgroups
}

// ListHsmAllgroupsResponseObjectAsListHsmAllgroupsResponse is a convenience function that returns ListHsmAllgroupsResponseObject wrapped in ListHsmAllgroupsResponse
func ListHsmAllgroupsResponseObjectAsListHsmAllgroupsResponse(v *ListHsmAllgroupsResponseObject) ListHsmAllgroupsResponse {
	return ListHsmAllgroupsResponse{
		ListHsmAllgroupsResponseObject: v,
	}
}

// []HsmAllgroupsAsListHsmAllgroupsResponse is a convenience function that returns []HsmAllgroups wrapped in ListHsmAllgroupsResponse
func ArrayOfHsmAllgroupsAsListHsmAllgroupsResponse(v *[]HsmAllgroups) ListHsmAllgroupsResponse {
	return ListHsmAllgroupsResponse{
		ArrayOfHsmAllgroups: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListHsmAllgroupsResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListHsmAllgroupsResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListHsmAllgroupsResponseObject)
	if err == nil {
		jsonListHsmAllgroupsResponseObject, _ := json.Marshal(dst.ListHsmAllgroupsResponseObject)
		if string(jsonListHsmAllgroupsResponseObject) == "{}" { // empty struct
			dst.ListHsmAllgroupsResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListHsmAllgroupsResponseObject = nil
	}

	// try to unmarshal data into ArrayOfHsmAllgroups
	err = newStrictDecoder(data).Decode(&dst.ArrayOfHsmAllgroups)
	if err == nil {
		jsonArrayOfHsmAllgroups, _ := json.Marshal(dst.ArrayOfHsmAllgroups)
		if string(jsonArrayOfHsmAllgroups) == "{}" { // empty struct
			dst.ArrayOfHsmAllgroups = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfHsmAllgroups = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListHsmAllgroupsResponseObject = nil
		dst.ArrayOfHsmAllgroups = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListHsmAllgroupsResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListHsmAllgroupsResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListHsmAllgroupsResponse) MarshalJSON() ([]byte, error) {
	if src.ListHsmAllgroupsResponseObject != nil {
		return json.Marshal(&src.ListHsmAllgroupsResponseObject)
	}

	if src.ArrayOfHsmAllgroups != nil {
		return json.Marshal(&src.ArrayOfHsmAllgroups)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListHsmAllgroupsResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListHsmAllgroupsResponseObject != nil {
		return obj.ListHsmAllgroupsResponseObject
	}

	if obj.ArrayOfHsmAllgroups != nil {
		return obj.ArrayOfHsmAllgroups
	}

	// all schemas are nil
	return nil
}

type NullableListHsmAllgroupsResponse struct {
	value *ListHsmAllgroupsResponse
	isSet bool
}

func (v NullableListHsmAllgroupsResponse) Get() *ListHsmAllgroupsResponse {
	return v.value
}

func (v *NullableListHsmAllgroupsResponse) Set(val *ListHsmAllgroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListHsmAllgroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListHsmAllgroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListHsmAllgroupsResponse(val *ListHsmAllgroupsResponse) *NullableListHsmAllgroupsResponse {
	return &NullableListHsmAllgroupsResponse{value: val, isSet: true}
}

func (v NullableListHsmAllgroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListHsmAllgroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
