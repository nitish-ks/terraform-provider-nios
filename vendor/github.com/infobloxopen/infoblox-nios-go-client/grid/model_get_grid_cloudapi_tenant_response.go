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

// GetGridCloudapiTenantResponse - struct for GetGridCloudapiTenantResponse
type GetGridCloudapiTenantResponse struct {
	GetGridCloudapiTenantResponseObjectAsResult *GetGridCloudapiTenantResponseObjectAsResult
	GridCloudapiTenant                          *GridCloudapiTenant
}

// GetGridCloudapiTenantResponseObjectAsResultAsGetGridCloudapiTenantResponse is a convenience function that returns GetGridCloudapiTenantResponseObjectAsResult wrapped in GetGridCloudapiTenantResponse
func GetGridCloudapiTenantResponseObjectAsResultAsGetGridCloudapiTenantResponse(v *GetGridCloudapiTenantResponseObjectAsResult) GetGridCloudapiTenantResponse {
	return GetGridCloudapiTenantResponse{
		GetGridCloudapiTenantResponseObjectAsResult: v,
	}
}

// GridCloudapiTenantAsGetGridCloudapiTenantResponse is a convenience function that returns GridCloudapiTenant wrapped in GetGridCloudapiTenantResponse
func GridCloudapiTenantAsGetGridCloudapiTenantResponse(v *GridCloudapiTenant) GetGridCloudapiTenantResponse {
	return GetGridCloudapiTenantResponse{
		GridCloudapiTenant: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetGridCloudapiTenantResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetGridCloudapiTenantResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetGridCloudapiTenantResponseObjectAsResult)
	if err == nil {
		jsonGetGridCloudapiTenantResponseObjectAsResult, _ := json.Marshal(dst.GetGridCloudapiTenantResponseObjectAsResult)
		if string(jsonGetGridCloudapiTenantResponseObjectAsResult) == "{}" { // empty struct
			dst.GetGridCloudapiTenantResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetGridCloudapiTenantResponseObjectAsResult = nil
	}

	// try to unmarshal data into GridCloudapiTenant
	err = newStrictDecoder(data).Decode(&dst.GridCloudapiTenant)
	if err == nil {
		jsonGridCloudapiTenant, _ := json.Marshal(dst.GridCloudapiTenant)
		if string(jsonGridCloudapiTenant) == "{}" { // empty struct
			dst.GridCloudapiTenant = nil
		} else {
			match++
		}
	} else {
		dst.GridCloudapiTenant = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetGridCloudapiTenantResponseObjectAsResult = nil
		dst.GridCloudapiTenant = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetGridCloudapiTenantResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetGridCloudapiTenantResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetGridCloudapiTenantResponse) MarshalJSON() ([]byte, error) {
	if src.GetGridCloudapiTenantResponseObjectAsResult != nil {
		return json.Marshal(&src.GetGridCloudapiTenantResponseObjectAsResult)
	}

	if src.GridCloudapiTenant != nil {
		return json.Marshal(&src.GridCloudapiTenant)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetGridCloudapiTenantResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetGridCloudapiTenantResponseObjectAsResult != nil {
		return obj.GetGridCloudapiTenantResponseObjectAsResult
	}

	if obj.GridCloudapiTenant != nil {
		return obj.GridCloudapiTenant
	}

	// all schemas are nil
	return nil
}

type NullableGetGridCloudapiTenantResponse struct {
	value *GetGridCloudapiTenantResponse
	isSet bool
}

func (v NullableGetGridCloudapiTenantResponse) Get() *GetGridCloudapiTenantResponse {
	return v.value
}

func (v *NullableGetGridCloudapiTenantResponse) Set(val *GetGridCloudapiTenantResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGridCloudapiTenantResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGridCloudapiTenantResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGridCloudapiTenantResponse(val *GetGridCloudapiTenantResponse) *NullableGetGridCloudapiTenantResponse {
	return &NullableGetGridCloudapiTenantResponse{value: val, isSet: true}
}

func (v NullableGetGridCloudapiTenantResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGridCloudapiTenantResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
