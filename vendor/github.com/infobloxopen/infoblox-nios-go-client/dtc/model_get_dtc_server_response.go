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

// GetDtcServerResponse - struct for GetDtcServerResponse
type GetDtcServerResponse struct {
	DtcServer                          *DtcServer
	GetDtcServerResponseObjectAsResult *GetDtcServerResponseObjectAsResult
}

// DtcServerAsGetDtcServerResponse is a convenience function that returns DtcServer wrapped in GetDtcServerResponse
func DtcServerAsGetDtcServerResponse(v *DtcServer) GetDtcServerResponse {
	return GetDtcServerResponse{
		DtcServer: v,
	}
}

// GetDtcServerResponseObjectAsResultAsGetDtcServerResponse is a convenience function that returns GetDtcServerResponseObjectAsResult wrapped in GetDtcServerResponse
func GetDtcServerResponseObjectAsResultAsGetDtcServerResponse(v *GetDtcServerResponseObjectAsResult) GetDtcServerResponse {
	return GetDtcServerResponse{
		GetDtcServerResponseObjectAsResult: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetDtcServerResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DtcServer
	err = newStrictDecoder(data).Decode(&dst.DtcServer)
	if err == nil {
		jsonDtcServer, _ := json.Marshal(dst.DtcServer)
		if string(jsonDtcServer) == "{}" { // empty struct
			dst.DtcServer = nil
		} else {
			match++
		}
	} else {
		dst.DtcServer = nil
	}

	// try to unmarshal data into GetDtcServerResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetDtcServerResponseObjectAsResult)
	if err == nil {
		jsonGetDtcServerResponseObjectAsResult, _ := json.Marshal(dst.GetDtcServerResponseObjectAsResult)
		if string(jsonGetDtcServerResponseObjectAsResult) == "{}" { // empty struct
			dst.GetDtcServerResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetDtcServerResponseObjectAsResult = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DtcServer = nil
		dst.GetDtcServerResponseObjectAsResult = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetDtcServerResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetDtcServerResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetDtcServerResponse) MarshalJSON() ([]byte, error) {
	if src.DtcServer != nil {
		return json.Marshal(&src.DtcServer)
	}

	if src.GetDtcServerResponseObjectAsResult != nil {
		return json.Marshal(&src.GetDtcServerResponseObjectAsResult)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetDtcServerResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DtcServer != nil {
		return obj.DtcServer
	}

	if obj.GetDtcServerResponseObjectAsResult != nil {
		return obj.GetDtcServerResponseObjectAsResult
	}

	// all schemas are nil
	return nil
}

type NullableGetDtcServerResponse struct {
	value *GetDtcServerResponse
	isSet bool
}

func (v NullableGetDtcServerResponse) Get() *GetDtcServerResponse {
	return v.value
}

func (v *NullableGetDtcServerResponse) Set(val *GetDtcServerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDtcServerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDtcServerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDtcServerResponse(val *GetDtcServerResponse) *NullableGetDtcServerResponse {
	return &NullableGetDtcServerResponse{value: val, isSet: true}
}

func (v NullableGetDtcServerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDtcServerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
