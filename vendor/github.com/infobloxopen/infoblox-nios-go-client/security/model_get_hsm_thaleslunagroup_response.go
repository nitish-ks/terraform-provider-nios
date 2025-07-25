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

// GetHsmThaleslunagroupResponse - struct for GetHsmThaleslunagroupResponse
type GetHsmThaleslunagroupResponse struct {
	GetHsmThaleslunagroupResponseObjectAsResult *GetHsmThaleslunagroupResponseObjectAsResult
	HsmThaleslunagroup                          *HsmThaleslunagroup
}

// GetHsmThaleslunagroupResponseObjectAsResultAsGetHsmThaleslunagroupResponse is a convenience function that returns GetHsmThaleslunagroupResponseObjectAsResult wrapped in GetHsmThaleslunagroupResponse
func GetHsmThaleslunagroupResponseObjectAsResultAsGetHsmThaleslunagroupResponse(v *GetHsmThaleslunagroupResponseObjectAsResult) GetHsmThaleslunagroupResponse {
	return GetHsmThaleslunagroupResponse{
		GetHsmThaleslunagroupResponseObjectAsResult: v,
	}
}

// HsmThaleslunagroupAsGetHsmThaleslunagroupResponse is a convenience function that returns HsmThaleslunagroup wrapped in GetHsmThaleslunagroupResponse
func HsmThaleslunagroupAsGetHsmThaleslunagroupResponse(v *HsmThaleslunagroup) GetHsmThaleslunagroupResponse {
	return GetHsmThaleslunagroupResponse{
		HsmThaleslunagroup: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetHsmThaleslunagroupResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetHsmThaleslunagroupResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetHsmThaleslunagroupResponseObjectAsResult)
	if err == nil {
		jsonGetHsmThaleslunagroupResponseObjectAsResult, _ := json.Marshal(dst.GetHsmThaleslunagroupResponseObjectAsResult)
		if string(jsonGetHsmThaleslunagroupResponseObjectAsResult) == "{}" { // empty struct
			dst.GetHsmThaleslunagroupResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetHsmThaleslunagroupResponseObjectAsResult = nil
	}

	// try to unmarshal data into HsmThaleslunagroup
	err = newStrictDecoder(data).Decode(&dst.HsmThaleslunagroup)
	if err == nil {
		jsonHsmThaleslunagroup, _ := json.Marshal(dst.HsmThaleslunagroup)
		if string(jsonHsmThaleslunagroup) == "{}" { // empty struct
			dst.HsmThaleslunagroup = nil
		} else {
			match++
		}
	} else {
		dst.HsmThaleslunagroup = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetHsmThaleslunagroupResponseObjectAsResult = nil
		dst.HsmThaleslunagroup = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetHsmThaleslunagroupResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetHsmThaleslunagroupResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetHsmThaleslunagroupResponse) MarshalJSON() ([]byte, error) {
	if src.GetHsmThaleslunagroupResponseObjectAsResult != nil {
		return json.Marshal(&src.GetHsmThaleslunagroupResponseObjectAsResult)
	}

	if src.HsmThaleslunagroup != nil {
		return json.Marshal(&src.HsmThaleslunagroup)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetHsmThaleslunagroupResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetHsmThaleslunagroupResponseObjectAsResult != nil {
		return obj.GetHsmThaleslunagroupResponseObjectAsResult
	}

	if obj.HsmThaleslunagroup != nil {
		return obj.HsmThaleslunagroup
	}

	// all schemas are nil
	return nil
}

type NullableGetHsmThaleslunagroupResponse struct {
	value *GetHsmThaleslunagroupResponse
	isSet bool
}

func (v NullableGetHsmThaleslunagroupResponse) Get() *GetHsmThaleslunagroupResponse {
	return v.value
}

func (v *NullableGetHsmThaleslunagroupResponse) Set(val *GetHsmThaleslunagroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetHsmThaleslunagroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetHsmThaleslunagroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetHsmThaleslunagroupResponse(val *GetHsmThaleslunagroupResponse) *NullableGetHsmThaleslunagroupResponse {
	return &NullableGetHsmThaleslunagroupResponse{value: val, isSet: true}
}

func (v NullableGetHsmThaleslunagroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetHsmThaleslunagroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
