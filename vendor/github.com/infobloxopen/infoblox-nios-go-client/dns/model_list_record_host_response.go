/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
	"fmt"
)

// ListRecordHostResponse - struct for ListRecordHostResponse
type ListRecordHostResponse struct {
	ListRecordHostResponseObject *ListRecordHostResponseObject
	ArrayOfRecordHost            *[]RecordHost
}

// ListRecordHostResponseObjectAsListRecordHostResponse is a convenience function that returns ListRecordHostResponseObject wrapped in ListRecordHostResponse
func ListRecordHostResponseObjectAsListRecordHostResponse(v *ListRecordHostResponseObject) ListRecordHostResponse {
	return ListRecordHostResponse{
		ListRecordHostResponseObject: v,
	}
}

// []RecordHostAsListRecordHostResponse is a convenience function that returns []RecordHost wrapped in ListRecordHostResponse
func ArrayOfRecordHostAsListRecordHostResponse(v *[]RecordHost) ListRecordHostResponse {
	return ListRecordHostResponse{
		ArrayOfRecordHost: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListRecordHostResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ListRecordHostResponseObject
	err = newStrictDecoder(data).Decode(&dst.ListRecordHostResponseObject)
	if err == nil {
		jsonListRecordHostResponseObject, _ := json.Marshal(dst.ListRecordHostResponseObject)
		if string(jsonListRecordHostResponseObject) == "{}" { // empty struct
			dst.ListRecordHostResponseObject = nil
		} else {
			match++
		}
	} else {
		dst.ListRecordHostResponseObject = nil
	}

	// try to unmarshal data into ArrayOfRecordHost
	err = newStrictDecoder(data).Decode(&dst.ArrayOfRecordHost)
	if err == nil {
		jsonArrayOfRecordHost, _ := json.Marshal(dst.ArrayOfRecordHost)
		if string(jsonArrayOfRecordHost) == "{}" { // empty struct
			dst.ArrayOfRecordHost = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfRecordHost = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ListRecordHostResponseObject = nil
		dst.ArrayOfRecordHost = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListRecordHostResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListRecordHostResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListRecordHostResponse) MarshalJSON() ([]byte, error) {
	if src.ListRecordHostResponseObject != nil {
		return json.Marshal(&src.ListRecordHostResponseObject)
	}

	if src.ArrayOfRecordHost != nil {
		return json.Marshal(&src.ArrayOfRecordHost)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListRecordHostResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ListRecordHostResponseObject != nil {
		return obj.ListRecordHostResponseObject
	}

	if obj.ArrayOfRecordHost != nil {
		return obj.ArrayOfRecordHost
	}

	// all schemas are nil
	return nil
}

type NullableListRecordHostResponse struct {
	value *ListRecordHostResponse
	isSet bool
}

func (v NullableListRecordHostResponse) Get() *ListRecordHostResponse {
	return v.value
}

func (v *NullableListRecordHostResponse) Set(val *ListRecordHostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRecordHostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRecordHostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRecordHostResponse(val *ListRecordHostResponse) *NullableListRecordHostResponse {
	return &NullableListRecordHostResponse{value: val, isSet: true}
}

func (v NullableListRecordHostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRecordHostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
