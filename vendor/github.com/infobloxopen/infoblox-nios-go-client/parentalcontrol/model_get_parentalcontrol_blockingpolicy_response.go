/*
Infoblox PARENTALCONTROL API

OpenAPI specification for Infoblox NIOS WAPI PARENTALCONTROL objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package parentalcontrol

import (
	"encoding/json"
	"fmt"
)

// GetParentalcontrolBlockingpolicyResponse - struct for GetParentalcontrolBlockingpolicyResponse
type GetParentalcontrolBlockingpolicyResponse struct {
	GetParentalcontrolBlockingpolicyResponseObjectAsResult *GetParentalcontrolBlockingpolicyResponseObjectAsResult
	ParentalcontrolBlockingpolicy                          *ParentalcontrolBlockingpolicy
}

// GetParentalcontrolBlockingpolicyResponseObjectAsResultAsGetParentalcontrolBlockingpolicyResponse is a convenience function that returns GetParentalcontrolBlockingpolicyResponseObjectAsResult wrapped in GetParentalcontrolBlockingpolicyResponse
func GetParentalcontrolBlockingpolicyResponseObjectAsResultAsGetParentalcontrolBlockingpolicyResponse(v *GetParentalcontrolBlockingpolicyResponseObjectAsResult) GetParentalcontrolBlockingpolicyResponse {
	return GetParentalcontrolBlockingpolicyResponse{
		GetParentalcontrolBlockingpolicyResponseObjectAsResult: v,
	}
}

// ParentalcontrolBlockingpolicyAsGetParentalcontrolBlockingpolicyResponse is a convenience function that returns ParentalcontrolBlockingpolicy wrapped in GetParentalcontrolBlockingpolicyResponse
func ParentalcontrolBlockingpolicyAsGetParentalcontrolBlockingpolicyResponse(v *ParentalcontrolBlockingpolicy) GetParentalcontrolBlockingpolicyResponse {
	return GetParentalcontrolBlockingpolicyResponse{
		ParentalcontrolBlockingpolicy: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetParentalcontrolBlockingpolicyResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetParentalcontrolBlockingpolicyResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetParentalcontrolBlockingpolicyResponseObjectAsResult)
	if err == nil {
		jsonGetParentalcontrolBlockingpolicyResponseObjectAsResult, _ := json.Marshal(dst.GetParentalcontrolBlockingpolicyResponseObjectAsResult)
		if string(jsonGetParentalcontrolBlockingpolicyResponseObjectAsResult) == "{}" { // empty struct
			dst.GetParentalcontrolBlockingpolicyResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetParentalcontrolBlockingpolicyResponseObjectAsResult = nil
	}

	// try to unmarshal data into ParentalcontrolBlockingpolicy
	err = newStrictDecoder(data).Decode(&dst.ParentalcontrolBlockingpolicy)
	if err == nil {
		jsonParentalcontrolBlockingpolicy, _ := json.Marshal(dst.ParentalcontrolBlockingpolicy)
		if string(jsonParentalcontrolBlockingpolicy) == "{}" { // empty struct
			dst.ParentalcontrolBlockingpolicy = nil
		} else {
			match++
		}
	} else {
		dst.ParentalcontrolBlockingpolicy = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetParentalcontrolBlockingpolicyResponseObjectAsResult = nil
		dst.ParentalcontrolBlockingpolicy = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetParentalcontrolBlockingpolicyResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetParentalcontrolBlockingpolicyResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetParentalcontrolBlockingpolicyResponse) MarshalJSON() ([]byte, error) {
	if src.GetParentalcontrolBlockingpolicyResponseObjectAsResult != nil {
		return json.Marshal(&src.GetParentalcontrolBlockingpolicyResponseObjectAsResult)
	}

	if src.ParentalcontrolBlockingpolicy != nil {
		return json.Marshal(&src.ParentalcontrolBlockingpolicy)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetParentalcontrolBlockingpolicyResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetParentalcontrolBlockingpolicyResponseObjectAsResult != nil {
		return obj.GetParentalcontrolBlockingpolicyResponseObjectAsResult
	}

	if obj.ParentalcontrolBlockingpolicy != nil {
		return obj.ParentalcontrolBlockingpolicy
	}

	// all schemas are nil
	return nil
}

type NullableGetParentalcontrolBlockingpolicyResponse struct {
	value *GetParentalcontrolBlockingpolicyResponse
	isSet bool
}

func (v NullableGetParentalcontrolBlockingpolicyResponse) Get() *GetParentalcontrolBlockingpolicyResponse {
	return v.value
}

func (v *NullableGetParentalcontrolBlockingpolicyResponse) Set(val *GetParentalcontrolBlockingpolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetParentalcontrolBlockingpolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetParentalcontrolBlockingpolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetParentalcontrolBlockingpolicyResponse(val *GetParentalcontrolBlockingpolicyResponse) *NullableGetParentalcontrolBlockingpolicyResponse {
	return &NullableGetParentalcontrolBlockingpolicyResponse{value: val, isSet: true}
}

func (v NullableGetParentalcontrolBlockingpolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetParentalcontrolBlockingpolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
