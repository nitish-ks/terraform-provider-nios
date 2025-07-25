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

// ListGmcscheduleResponse - struct for ListGmcscheduleResponse
type ListGmcscheduleResponse struct {
	ListGmcscheduleResponseObject *ListGmcscheduleResponseObject
	ArrayOfGmcschedule            *[]Gmcschedule
}

// ListGmcscheduleResponseObjectAsListGmcscheduleResponse is a convenience function that returns ListGmcscheduleResponseObject wrapped in ListGmcscheduleResponse
func ListGmcscheduleResponseObjectAsListGmcscheduleResponse(v *ListGmcscheduleResponseObject) ListGmcscheduleResponse {
	return ListGmcscheduleResponse{
		ListGmcscheduleResponseObject: v,
	}
}

// []GmcscheduleAsListGmcscheduleResponse is a convenience function that returns []Gmcschedule wrapped in ListGmcscheduleResponse
func ArrayOfGmcscheduleAsListGmcscheduleResponse(v *[]Gmcschedule) ListGmcscheduleResponse {
	return ListGmcscheduleResponse{
		ArrayOfGmcschedule: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListGmcscheduleResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListGmcscheduleResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListGmcscheduleResponseObject)
	if err == nil {
		jsonListGmcscheduleResponseObject, _ := json.Marshal(dst.ListGmcscheduleResponseObject)
		if string(jsonListGmcscheduleResponseObject) == "{}" { // empty struct
			dst.ListGmcscheduleResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListGmcscheduleResponseObject = nil
	}

	// try to unmarshal data into ArrayOfGmcschedule
	err = newStrictDecoder(data).Decode(&dst.ArrayOfGmcschedule)
	if err == nil {
		jsonArrayOfGmcschedule, _ := json.Marshal(dst.ArrayOfGmcschedule)
		if string(jsonArrayOfGmcschedule) == "{}" { // empty struct
			dst.ArrayOfGmcschedule = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfGmcschedule = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListGmcscheduleResponseObject = nil
		dst.ArrayOfGmcschedule = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListGmcscheduleResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListGmcscheduleResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListGmcscheduleResponse) MarshalJSON() ([]byte, error) {
	if src.ListGmcscheduleResponseObject != nil {
		return json.Marshal(&src.ListGmcscheduleResponseObject)
	}

	if src.ArrayOfGmcschedule != nil {
		return json.Marshal(&src.ArrayOfGmcschedule)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListGmcscheduleResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListGmcscheduleResponseObject != nil {
		return obj.ListGmcscheduleResponseObject
	}

	if obj.ArrayOfGmcschedule != nil {
		return obj.ArrayOfGmcschedule
	}

	// all schemas are nil
	return nil
}

type NullableListGmcscheduleResponse struct {
	value *ListGmcscheduleResponse
	isSet bool
}

func (v NullableListGmcscheduleResponse) Get() *ListGmcscheduleResponse {
	return v.value
}

func (v *NullableListGmcscheduleResponse) Set(val *ListGmcscheduleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListGmcscheduleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListGmcscheduleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGmcscheduleResponse(val *ListGmcscheduleResponse) *NullableListGmcscheduleResponse {
	return &NullableListGmcscheduleResponse{value: val, isSet: true}
}

func (v NullableListGmcscheduleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListGmcscheduleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
