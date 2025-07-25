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

// CreateUpgradegroupResponse - struct for CreateUpgradegroupResponse
type CreateUpgradegroupResponse struct {
	CreateUpgradegroupResponseAsObject *CreateUpgradegroupResponseAsObject
	String                             *string
}

// CreateUpgradegroupResponseAsObjectAsCreateUpgradegroupResponse is a convenience function that returns CreateUpgradegroupResponseAsObject wrapped in CreateUpgradegroupResponse
func CreateUpgradegroupResponseAsObjectAsCreateUpgradegroupResponse(v *CreateUpgradegroupResponseAsObject) CreateUpgradegroupResponse {
	return CreateUpgradegroupResponse{
		CreateUpgradegroupResponseAsObject: v,
	}
}

// stringAsCreateUpgradegroupResponse is a convenience function that returns string wrapped in CreateUpgradegroupResponse
func StringAsCreateUpgradegroupResponse(v *string) CreateUpgradegroupResponse {
	return CreateUpgradegroupResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateUpgradegroupResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateUpgradegroupResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.CreateUpgradegroupResponseAsObject)
	if err == nil {
		jsonCreateUpgradegroupResponseAsObject, _ := json.Marshal(dst.CreateUpgradegroupResponseAsObject)
		if string(jsonCreateUpgradegroupResponseAsObject) == "{}" { // empty struct
			dst.CreateUpgradegroupResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.CreateUpgradegroupResponseAsObject = nil
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
		dst.CreateUpgradegroupResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateUpgradegroupResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateUpgradegroupResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateUpgradegroupResponse) MarshalJSON() ([]byte, error) {
	if src.CreateUpgradegroupResponseAsObject != nil {
		return json.Marshal(&src.CreateUpgradegroupResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateUpgradegroupResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateUpgradegroupResponseAsObject != nil {
		return obj.CreateUpgradegroupResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateUpgradegroupResponse struct {
	value *CreateUpgradegroupResponse
	isSet bool
}

func (v NullableCreateUpgradegroupResponse) Get() *CreateUpgradegroupResponse {
	return v.value
}

func (v *NullableCreateUpgradegroupResponse) Set(val *CreateUpgradegroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUpgradegroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUpgradegroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUpgradegroupResponse(val *CreateUpgradegroupResponse) *NullableCreateUpgradegroupResponse {
	return &NullableCreateUpgradegroupResponse{value: val, isSet: true}
}

func (v NullableCreateUpgradegroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUpgradegroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
