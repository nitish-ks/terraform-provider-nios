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

// CreateSharedrecordAResponse - struct for CreateSharedrecordAResponse
type CreateSharedrecordAResponse struct {
	CreateSharedrecordAResponseAsObject *CreateSharedrecordAResponseAsObject
	String                              *string
}

// CreateSharedrecordAResponseAsObjectAsCreateSharedrecordAResponse is a convenience function that returns CreateSharedrecordAResponseAsObject wrapped in CreateSharedrecordAResponse
func CreateSharedrecordAResponseAsObjectAsCreateSharedrecordAResponse(v *CreateSharedrecordAResponseAsObject) CreateSharedrecordAResponse {
	return CreateSharedrecordAResponse{
		CreateSharedrecordAResponseAsObject: v,
	}
}

// stringAsCreateSharedrecordAResponse is a convenience function that returns string wrapped in CreateSharedrecordAResponse
func StringAsCreateSharedrecordAResponse(v *string) CreateSharedrecordAResponse {
	return CreateSharedrecordAResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateSharedrecordAResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateSharedrecordAResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.CreateSharedrecordAResponseAsObject)
	if err == nil {
		jsonCreateSharedrecordAResponseAsObject, _ := json.Marshal(dst.CreateSharedrecordAResponseAsObject)
		if string(jsonCreateSharedrecordAResponseAsObject) == "{}" { // empty struct
			dst.CreateSharedrecordAResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.CreateSharedrecordAResponseAsObject = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CreateSharedrecordAResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateSharedrecordAResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateSharedrecordAResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateSharedrecordAResponse) MarshalJSON() ([]byte, error) {
	if src.CreateSharedrecordAResponseAsObject != nil {
		return json.Marshal(&src.CreateSharedrecordAResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateSharedrecordAResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateSharedrecordAResponseAsObject != nil {
		return obj.CreateSharedrecordAResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateSharedrecordAResponse struct {
	value *CreateSharedrecordAResponse
	isSet bool
}

func (v NullableCreateSharedrecordAResponse) Get() *CreateSharedrecordAResponse {
	return v.value
}

func (v *NullableCreateSharedrecordAResponse) Set(val *CreateSharedrecordAResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSharedrecordAResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSharedrecordAResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSharedrecordAResponse(val *CreateSharedrecordAResponse) *NullableCreateSharedrecordAResponse {
	return &NullableCreateSharedrecordAResponse{value: val, isSet: true}
}

func (v NullableCreateSharedrecordAResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSharedrecordAResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
