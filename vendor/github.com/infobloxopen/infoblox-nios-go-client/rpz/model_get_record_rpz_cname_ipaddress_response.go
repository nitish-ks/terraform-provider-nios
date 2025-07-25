/*
Infoblox RPZ API

OpenAPI specification for Infoblox NIOS WAPI RPZ objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rpz

import (
	"encoding/json"
	"fmt"
)

// GetRecordRpzCnameIpaddressResponse - struct for GetRecordRpzCnameIpaddressResponse
type GetRecordRpzCnameIpaddressResponse struct {
	GetRecordRpzCnameIpaddressResponseObjectAsResult *GetRecordRpzCnameIpaddressResponseObjectAsResult
	RecordRpzCnameIpaddress                          *RecordRpzCnameIpaddress
}

// GetRecordRpzCnameIpaddressResponseObjectAsResultAsGetRecordRpzCnameIpaddressResponse is a convenience function that returns GetRecordRpzCnameIpaddressResponseObjectAsResult wrapped in GetRecordRpzCnameIpaddressResponse
func GetRecordRpzCnameIpaddressResponseObjectAsResultAsGetRecordRpzCnameIpaddressResponse(v *GetRecordRpzCnameIpaddressResponseObjectAsResult) GetRecordRpzCnameIpaddressResponse {
	return GetRecordRpzCnameIpaddressResponse{
		GetRecordRpzCnameIpaddressResponseObjectAsResult: v,
	}
}

// RecordRpzCnameIpaddressAsGetRecordRpzCnameIpaddressResponse is a convenience function that returns RecordRpzCnameIpaddress wrapped in GetRecordRpzCnameIpaddressResponse
func RecordRpzCnameIpaddressAsGetRecordRpzCnameIpaddressResponse(v *RecordRpzCnameIpaddress) GetRecordRpzCnameIpaddressResponse {
	return GetRecordRpzCnameIpaddressResponse{
		RecordRpzCnameIpaddress: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetRecordRpzCnameIpaddressResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetRecordRpzCnameIpaddressResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetRecordRpzCnameIpaddressResponseObjectAsResult)
	if err == nil {
		jsonGetRecordRpzCnameIpaddressResponseObjectAsResult, _ := json.Marshal(dst.GetRecordRpzCnameIpaddressResponseObjectAsResult)
		if string(jsonGetRecordRpzCnameIpaddressResponseObjectAsResult) == "{}" { // empty struct
			dst.GetRecordRpzCnameIpaddressResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetRecordRpzCnameIpaddressResponseObjectAsResult = nil
	}

	// try to unmarshal data into RecordRpzCnameIpaddress
	err = newStrictDecoder(data).Decode(&dst.RecordRpzCnameIpaddress)
	if err == nil {
		jsonRecordRpzCnameIpaddress, _ := json.Marshal(dst.RecordRpzCnameIpaddress)
		if string(jsonRecordRpzCnameIpaddress) == "{}" { // empty struct
			dst.RecordRpzCnameIpaddress = nil
		} else {
			match++
		}
	} else {
		dst.RecordRpzCnameIpaddress = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetRecordRpzCnameIpaddressResponseObjectAsResult = nil
		dst.RecordRpzCnameIpaddress = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetRecordRpzCnameIpaddressResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetRecordRpzCnameIpaddressResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetRecordRpzCnameIpaddressResponse) MarshalJSON() ([]byte, error) {
	if src.GetRecordRpzCnameIpaddressResponseObjectAsResult != nil {
		return json.Marshal(&src.GetRecordRpzCnameIpaddressResponseObjectAsResult)
	}

	if src.RecordRpzCnameIpaddress != nil {
		return json.Marshal(&src.RecordRpzCnameIpaddress)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetRecordRpzCnameIpaddressResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetRecordRpzCnameIpaddressResponseObjectAsResult != nil {
		return obj.GetRecordRpzCnameIpaddressResponseObjectAsResult
	}

	if obj.RecordRpzCnameIpaddress != nil {
		return obj.RecordRpzCnameIpaddress
	}

	// all schemas are nil
	return nil
}

type NullableGetRecordRpzCnameIpaddressResponse struct {
	value *GetRecordRpzCnameIpaddressResponse
	isSet bool
}

func (v NullableGetRecordRpzCnameIpaddressResponse) Get() *GetRecordRpzCnameIpaddressResponse {
	return v.value
}

func (v *NullableGetRecordRpzCnameIpaddressResponse) Set(val *GetRecordRpzCnameIpaddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecordRpzCnameIpaddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecordRpzCnameIpaddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecordRpzCnameIpaddressResponse(val *GetRecordRpzCnameIpaddressResponse) *NullableGetRecordRpzCnameIpaddressResponse {
	return &NullableGetRecordRpzCnameIpaddressResponse{value: val, isSet: true}
}

func (v NullableGetRecordRpzCnameIpaddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecordRpzCnameIpaddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
