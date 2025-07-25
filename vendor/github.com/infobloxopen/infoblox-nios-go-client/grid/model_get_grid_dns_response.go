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

// GetGridDnsResponse - struct for GetGridDnsResponse
type GetGridDnsResponse struct {
	GetGridDnsResponseObjectAsResult *GetGridDnsResponseObjectAsResult
	GridDns                          *GridDns
}

// GetGridDnsResponseObjectAsResultAsGetGridDnsResponse is a convenience function that returns GetGridDnsResponseObjectAsResult wrapped in GetGridDnsResponse
func GetGridDnsResponseObjectAsResultAsGetGridDnsResponse(v *GetGridDnsResponseObjectAsResult) GetGridDnsResponse {
	return GetGridDnsResponse{
		GetGridDnsResponseObjectAsResult: v,
	}
}

// GridDnsAsGetGridDnsResponse is a convenience function that returns GridDns wrapped in GetGridDnsResponse
func GridDnsAsGetGridDnsResponse(v *GridDns) GetGridDnsResponse {
	return GetGridDnsResponse{
		GridDns: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetGridDnsResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetGridDnsResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetGridDnsResponseObjectAsResult)
	if err == nil {
		jsonGetGridDnsResponseObjectAsResult, _ := json.Marshal(dst.GetGridDnsResponseObjectAsResult)
		if string(jsonGetGridDnsResponseObjectAsResult) == "{}" { // empty struct
			dst.GetGridDnsResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetGridDnsResponseObjectAsResult = nil
	}

	// try to unmarshal data into GridDns
	err = newStrictDecoder(data).Decode(&dst.GridDns)
	if err == nil {
		jsonGridDns, _ := json.Marshal(dst.GridDns)
		if string(jsonGridDns) == "{}" { // empty struct
			dst.GridDns = nil
		} else {
			match++
		}
	} else {
		dst.GridDns = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetGridDnsResponseObjectAsResult = nil
		dst.GridDns = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetGridDnsResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetGridDnsResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetGridDnsResponse) MarshalJSON() ([]byte, error) {
	if src.GetGridDnsResponseObjectAsResult != nil {
		return json.Marshal(&src.GetGridDnsResponseObjectAsResult)
	}

	if src.GridDns != nil {
		return json.Marshal(&src.GridDns)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetGridDnsResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetGridDnsResponseObjectAsResult != nil {
		return obj.GetGridDnsResponseObjectAsResult
	}

	if obj.GridDns != nil {
		return obj.GridDns
	}

	// all schemas are nil
	return nil
}

type NullableGetGridDnsResponse struct {
	value *GetGridDnsResponse
	isSet bool
}

func (v NullableGetGridDnsResponse) Get() *GetGridDnsResponse {
	return v.value
}

func (v *NullableGetGridDnsResponse) Set(val *GetGridDnsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGridDnsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGridDnsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGridDnsResponse(val *GetGridDnsResponse) *NullableGetGridDnsResponse {
	return &NullableGetGridDnsResponse{value: val, isSet: true}
}

func (v NullableGetGridDnsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGridDnsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
