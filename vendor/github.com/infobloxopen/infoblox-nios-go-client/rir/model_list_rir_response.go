/*
Infoblox RIR API

OpenAPI specification for Infoblox NIOS WAPI RIR objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rir

import (
	"encoding/json"
	"fmt"
)

// ListRirResponse - struct for ListRirResponse
type ListRirResponse struct {
	ListRirResponseObject *ListRirResponseObject
	ArrayOfRir            *[]Rir
}

// ListRirResponseObjectAsListRirResponse is a convenience function that returns ListRirResponseObject wrapped in ListRirResponse
func ListRirResponseObjectAsListRirResponse(v *ListRirResponseObject) ListRirResponse {
	return ListRirResponse{
		ListRirResponseObject: v,
	}
}

// []RirAsListRirResponse is a convenience function that returns []Rir wrapped in ListRirResponse
func ArrayOfRirAsListRirResponse(v *[]Rir) ListRirResponse {
	return ListRirResponse{
		ArrayOfRir: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListRirResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListRirResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListRirResponseObject)
	if err == nil {
		jsonListRirResponseObject, _ := json.Marshal(dst.ListRirResponseObject)
		if string(jsonListRirResponseObject) == "{}" { // empty struct
			dst.ListRirResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListRirResponseObject = nil
	}

	// try to unmarshal data into ArrayOfRir
	err = newStrictDecoder(data).Decode(&dst.ArrayOfRir)
	if err == nil {
		jsonArrayOfRir, _ := json.Marshal(dst.ArrayOfRir)
		if string(jsonArrayOfRir) == "{}" { // empty struct
			dst.ArrayOfRir = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfRir = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListRirResponseObject = nil
		dst.ArrayOfRir = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListRirResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListRirResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListRirResponse) MarshalJSON() ([]byte, error) {
	if src.ListRirResponseObject != nil {
		return json.Marshal(&src.ListRirResponseObject)
	}

	if src.ArrayOfRir != nil {
		return json.Marshal(&src.ArrayOfRir)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListRirResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListRirResponseObject != nil {
		return obj.ListRirResponseObject
	}

	if obj.ArrayOfRir != nil {
		return obj.ArrayOfRir
	}

	// all schemas are nil
	return nil
}

type NullableListRirResponse struct {
	value *ListRirResponse
	isSet bool
}

func (v NullableListRirResponse) Get() *ListRirResponse {
	return v.value
}

func (v *NullableListRirResponse) Set(val *ListRirResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRirResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRirResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRirResponse(val *ListRirResponse) *NullableListRirResponse {
	return &NullableListRirResponse{value: val, isSet: true}
}

func (v NullableListRirResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRirResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
