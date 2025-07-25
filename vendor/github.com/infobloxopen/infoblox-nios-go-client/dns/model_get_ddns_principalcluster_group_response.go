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

// GetDdnsPrincipalclusterGroupResponse - struct for GetDdnsPrincipalclusterGroupResponse
type GetDdnsPrincipalclusterGroupResponse struct {
	DdnsPrincipalclusterGroup                          *DdnsPrincipalclusterGroup
	GetDdnsPrincipalclusterGroupResponseObjectAsResult *GetDdnsPrincipalclusterGroupResponseObjectAsResult
}

// DdnsPrincipalclusterGroupAsGetDdnsPrincipalclusterGroupResponse is a convenience function that returns DdnsPrincipalclusterGroup wrapped in GetDdnsPrincipalclusterGroupResponse
func DdnsPrincipalclusterGroupAsGetDdnsPrincipalclusterGroupResponse(v *DdnsPrincipalclusterGroup) GetDdnsPrincipalclusterGroupResponse {
	return GetDdnsPrincipalclusterGroupResponse{
		DdnsPrincipalclusterGroup: v,
	}
}

// GetDdnsPrincipalclusterGroupResponseObjectAsResultAsGetDdnsPrincipalclusterGroupResponse is a convenience function that returns GetDdnsPrincipalclusterGroupResponseObjectAsResult wrapped in GetDdnsPrincipalclusterGroupResponse
func GetDdnsPrincipalclusterGroupResponseObjectAsResultAsGetDdnsPrincipalclusterGroupResponse(v *GetDdnsPrincipalclusterGroupResponseObjectAsResult) GetDdnsPrincipalclusterGroupResponse {
	return GetDdnsPrincipalclusterGroupResponse{
		GetDdnsPrincipalclusterGroupResponseObjectAsResult: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetDdnsPrincipalclusterGroupResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DdnsPrincipalclusterGroup
	err = newStrictDecoder(data).Decode(&dst.DdnsPrincipalclusterGroup)
	if err == nil {
		jsonDdnsPrincipalclusterGroup, _ := json.Marshal(dst.DdnsPrincipalclusterGroup)
		if string(jsonDdnsPrincipalclusterGroup) == "{}" { // empty struct
			dst.DdnsPrincipalclusterGroup = nil
		} else {
			match++
		}
	} else {
		dst.DdnsPrincipalclusterGroup = nil
	}

	// try to unmarshal data into GetDdnsPrincipalclusterGroupResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetDdnsPrincipalclusterGroupResponseObjectAsResult)
	if err == nil {
		jsonGetDdnsPrincipalclusterGroupResponseObjectAsResult, _ := json.Marshal(dst.GetDdnsPrincipalclusterGroupResponseObjectAsResult)
		if string(jsonGetDdnsPrincipalclusterGroupResponseObjectAsResult) == "{}" { // empty struct
			dst.GetDdnsPrincipalclusterGroupResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetDdnsPrincipalclusterGroupResponseObjectAsResult = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DdnsPrincipalclusterGroup = nil
		dst.GetDdnsPrincipalclusterGroupResponseObjectAsResult = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetDdnsPrincipalclusterGroupResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetDdnsPrincipalclusterGroupResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetDdnsPrincipalclusterGroupResponse) MarshalJSON() ([]byte, error) {
	if src.DdnsPrincipalclusterGroup != nil {
		return json.Marshal(&src.DdnsPrincipalclusterGroup)
	}

	if src.GetDdnsPrincipalclusterGroupResponseObjectAsResult != nil {
		return json.Marshal(&src.GetDdnsPrincipalclusterGroupResponseObjectAsResult)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetDdnsPrincipalclusterGroupResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DdnsPrincipalclusterGroup != nil {
		return obj.DdnsPrincipalclusterGroup
	}

	if obj.GetDdnsPrincipalclusterGroupResponseObjectAsResult != nil {
		return obj.GetDdnsPrincipalclusterGroupResponseObjectAsResult
	}

	// all schemas are nil
	return nil
}

type NullableGetDdnsPrincipalclusterGroupResponse struct {
	value *GetDdnsPrincipalclusterGroupResponse
	isSet bool
}

func (v NullableGetDdnsPrincipalclusterGroupResponse) Get() *GetDdnsPrincipalclusterGroupResponse {
	return v.value
}

func (v *NullableGetDdnsPrincipalclusterGroupResponse) Set(val *GetDdnsPrincipalclusterGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDdnsPrincipalclusterGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDdnsPrincipalclusterGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDdnsPrincipalclusterGroupResponse(val *GetDdnsPrincipalclusterGroupResponse) *NullableGetDdnsPrincipalclusterGroupResponse {
	return &NullableGetDdnsPrincipalclusterGroupResponse{value: val, isSet: true}
}

func (v NullableGetDdnsPrincipalclusterGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDdnsPrincipalclusterGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
