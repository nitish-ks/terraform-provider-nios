/*
Infoblox MICROSOFTSERVER API

OpenAPI specification for Infoblox NIOS WAPI MICROSOFTSERVER objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package microsoftserver

import (
	"encoding/json"
	"fmt"
)

// GetMsserverDhcpResponse - struct for GetMsserverDhcpResponse
type GetMsserverDhcpResponse struct {
	GetMsserverDhcpResponseObjectAsResult *GetMsserverDhcpResponseObjectAsResult
	MsserverDhcp                          *MsserverDhcp
}

// GetMsserverDhcpResponseObjectAsResultAsGetMsserverDhcpResponse is a convenience function that returns GetMsserverDhcpResponseObjectAsResult wrapped in GetMsserverDhcpResponse
func GetMsserverDhcpResponseObjectAsResultAsGetMsserverDhcpResponse(v *GetMsserverDhcpResponseObjectAsResult) GetMsserverDhcpResponse {
	return GetMsserverDhcpResponse{
		GetMsserverDhcpResponseObjectAsResult: v,
	}
}

// MsserverDhcpAsGetMsserverDhcpResponse is a convenience function that returns MsserverDhcp wrapped in GetMsserverDhcpResponse
func MsserverDhcpAsGetMsserverDhcpResponse(v *MsserverDhcp) GetMsserverDhcpResponse {
	return GetMsserverDhcpResponse{
		MsserverDhcp: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetMsserverDhcpResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetMsserverDhcpResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetMsserverDhcpResponseObjectAsResult)
	if err == nil {
		jsonGetMsserverDhcpResponseObjectAsResult, _ := json.Marshal(dst.GetMsserverDhcpResponseObjectAsResult)
		if string(jsonGetMsserverDhcpResponseObjectAsResult) == "{}" { // empty struct
			dst.GetMsserverDhcpResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetMsserverDhcpResponseObjectAsResult = nil
	}

	// try to unmarshal data into MsserverDhcp
	err = newStrictDecoder(data).Decode(&dst.MsserverDhcp)
	if err == nil {
		jsonMsserverDhcp, _ := json.Marshal(dst.MsserverDhcp)
		if string(jsonMsserverDhcp) == "{}" { // empty struct
			dst.MsserverDhcp = nil
		} else {
			match++
		}
	} else {
		dst.MsserverDhcp = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetMsserverDhcpResponseObjectAsResult = nil
		dst.MsserverDhcp = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetMsserverDhcpResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetMsserverDhcpResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetMsserverDhcpResponse) MarshalJSON() ([]byte, error) {
	if src.GetMsserverDhcpResponseObjectAsResult != nil {
		return json.Marshal(&src.GetMsserverDhcpResponseObjectAsResult)
	}

	if src.MsserverDhcp != nil {
		return json.Marshal(&src.MsserverDhcp)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetMsserverDhcpResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetMsserverDhcpResponseObjectAsResult != nil {
		return obj.GetMsserverDhcpResponseObjectAsResult
	}

	if obj.MsserverDhcp != nil {
		return obj.MsserverDhcp
	}

	// all schemas are nil
	return nil
}

type NullableGetMsserverDhcpResponse struct {
	value *GetMsserverDhcpResponse
	isSet bool
}

func (v NullableGetMsserverDhcpResponse) Get() *GetMsserverDhcpResponse {
	return v.value
}

func (v *NullableGetMsserverDhcpResponse) Set(val *GetMsserverDhcpResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMsserverDhcpResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMsserverDhcpResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMsserverDhcpResponse(val *GetMsserverDhcpResponse) *NullableGetMsserverDhcpResponse {
	return &NullableGetMsserverDhcpResponse{value: val, isSet: true}
}

func (v NullableGetMsserverDhcpResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMsserverDhcpResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
