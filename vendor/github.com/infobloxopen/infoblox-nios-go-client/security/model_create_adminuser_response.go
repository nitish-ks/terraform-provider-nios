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

// CreateAdminuserResponse - struct for CreateAdminuserResponse
type CreateAdminuserResponse struct {
	CreateAdminuserResponseAsObject *CreateAdminuserResponseAsObject
	String                          *string
}

// CreateAdminuserResponseAsObjectAsCreateAdminuserResponse is a convenience function that returns CreateAdminuserResponseAsObject wrapped in CreateAdminuserResponse
func CreateAdminuserResponseAsObjectAsCreateAdminuserResponse(v *CreateAdminuserResponseAsObject) CreateAdminuserResponse {
	return CreateAdminuserResponse{
		CreateAdminuserResponseAsObject: v,
	}
}

// stringAsCreateAdminuserResponse is a convenience function that returns string wrapped in CreateAdminuserResponse
func StringAsCreateAdminuserResponse(v *string) CreateAdminuserResponse {
	return CreateAdminuserResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateAdminuserResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateAdminuserResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.CreateAdminuserResponseAsObject)
	if err == nil {
		jsonCreateAdminuserResponseAsObject, _ := json.Marshal(dst.CreateAdminuserResponseAsObject)
		if string(jsonCreateAdminuserResponseAsObject) == "{}" { // empty struct
			dst.CreateAdminuserResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.CreateAdminuserResponseAsObject = nil
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
		dst.CreateAdminuserResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateAdminuserResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateAdminuserResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateAdminuserResponse) MarshalJSON() ([]byte, error) {
	if src.CreateAdminuserResponseAsObject != nil {
		return json.Marshal(&src.CreateAdminuserResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateAdminuserResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateAdminuserResponseAsObject != nil {
		return obj.CreateAdminuserResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateAdminuserResponse struct {
	value *CreateAdminuserResponse
	isSet bool
}

func (v NullableCreateAdminuserResponse) Get() *CreateAdminuserResponse {
	return v.value
}

func (v *NullableCreateAdminuserResponse) Set(val *CreateAdminuserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAdminuserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAdminuserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAdminuserResponse(val *CreateAdminuserResponse) *NullableCreateAdminuserResponse {
	return &NullableCreateAdminuserResponse{value: val, isSet: true}
}

func (v NullableCreateAdminuserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAdminuserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
