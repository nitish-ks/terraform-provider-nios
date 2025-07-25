/*
Infoblox DTC API

OpenAPI specification for Infoblox NIOS WAPI DTC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dtc

import (
	"encoding/json"
	"fmt"
)

// GetDtcTopologyResponse - struct for GetDtcTopologyResponse
type GetDtcTopologyResponse struct {
	DtcTopology                          *DtcTopology
	GetDtcTopologyResponseObjectAsResult *GetDtcTopologyResponseObjectAsResult
}

// DtcTopologyAsGetDtcTopologyResponse is a convenience function that returns DtcTopology wrapped in GetDtcTopologyResponse
func DtcTopologyAsGetDtcTopologyResponse(v *DtcTopology) GetDtcTopologyResponse {
	return GetDtcTopologyResponse{
		DtcTopology: v,
	}
}

// GetDtcTopologyResponseObjectAsResultAsGetDtcTopologyResponse is a convenience function that returns GetDtcTopologyResponseObjectAsResult wrapped in GetDtcTopologyResponse
func GetDtcTopologyResponseObjectAsResultAsGetDtcTopologyResponse(v *GetDtcTopologyResponseObjectAsResult) GetDtcTopologyResponse {
	return GetDtcTopologyResponse{
		GetDtcTopologyResponseObjectAsResult: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetDtcTopologyResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DtcTopology
	err = newStrictDecoder(data).Decode(&dst.DtcTopology)
	if err == nil {
		jsonDtcTopology, _ := json.Marshal(dst.DtcTopology)
		if string(jsonDtcTopology) == "{}" { // empty struct
			dst.DtcTopology = nil
		} else {
			match++
		}
	} else {
		dst.DtcTopology = nil
	}

	// try to unmarshal data into GetDtcTopologyResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetDtcTopologyResponseObjectAsResult)
	if err == nil {
		jsonGetDtcTopologyResponseObjectAsResult, _ := json.Marshal(dst.GetDtcTopologyResponseObjectAsResult)
		if string(jsonGetDtcTopologyResponseObjectAsResult) == "{}" { // empty struct
			dst.GetDtcTopologyResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetDtcTopologyResponseObjectAsResult = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DtcTopology = nil
		dst.GetDtcTopologyResponseObjectAsResult = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetDtcTopologyResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetDtcTopologyResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetDtcTopologyResponse) MarshalJSON() ([]byte, error) {
	if src.DtcTopology != nil {
		return json.Marshal(&src.DtcTopology)
	}

	if src.GetDtcTopologyResponseObjectAsResult != nil {
		return json.Marshal(&src.GetDtcTopologyResponseObjectAsResult)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetDtcTopologyResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DtcTopology != nil {
		return obj.DtcTopology
	}

	if obj.GetDtcTopologyResponseObjectAsResult != nil {
		return obj.GetDtcTopologyResponseObjectAsResult
	}

	// all schemas are nil
	return nil
}

type NullableGetDtcTopologyResponse struct {
	value *GetDtcTopologyResponse
	isSet bool
}

func (v NullableGetDtcTopologyResponse) Get() *GetDtcTopologyResponse {
	return v.value
}

func (v *NullableGetDtcTopologyResponse) Set(val *GetDtcTopologyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDtcTopologyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDtcTopologyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDtcTopologyResponse(val *GetDtcTopologyResponse) *NullableGetDtcTopologyResponse {
	return &NullableGetDtcTopologyResponse{value: val, isSet: true}
}

func (v NullableGetDtcTopologyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDtcTopologyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
