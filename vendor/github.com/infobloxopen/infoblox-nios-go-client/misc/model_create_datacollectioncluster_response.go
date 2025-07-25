/*
Infoblox MISC API

OpenAPI specification for Infoblox NIOS WAPI MISC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package misc

import (
	"encoding/json"
	"fmt"
)

// CreateDatacollectionclusterResponse - struct for CreateDatacollectionclusterResponse
type CreateDatacollectionclusterResponse struct {
	CreateDatacollectionclusterResponseAsObject *CreateDatacollectionclusterResponseAsObject
	String                                      *string
}

// CreateDatacollectionclusterResponseAsObjectAsCreateDatacollectionclusterResponse is a convenience function that returns CreateDatacollectionclusterResponseAsObject wrapped in CreateDatacollectionclusterResponse
func CreateDatacollectionclusterResponseAsObjectAsCreateDatacollectionclusterResponse(v *CreateDatacollectionclusterResponseAsObject) CreateDatacollectionclusterResponse {
	return CreateDatacollectionclusterResponse{
		CreateDatacollectionclusterResponseAsObject: v,
	}
}

// stringAsCreateDatacollectionclusterResponse is a convenience function that returns string wrapped in CreateDatacollectionclusterResponse
func StringAsCreateDatacollectionclusterResponse(v *string) CreateDatacollectionclusterResponse {
	return CreateDatacollectionclusterResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateDatacollectionclusterResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateDatacollectionclusterResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.CreateDatacollectionclusterResponseAsObject)
	if err == nil {
		jsonCreateDatacollectionclusterResponseAsObject, _ := json.Marshal(dst.CreateDatacollectionclusterResponseAsObject)
		if string(jsonCreateDatacollectionclusterResponseAsObject) == "{}" { // empty struct
			dst.CreateDatacollectionclusterResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.CreateDatacollectionclusterResponseAsObject = nil
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
		dst.CreateDatacollectionclusterResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateDatacollectionclusterResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateDatacollectionclusterResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateDatacollectionclusterResponse) MarshalJSON() ([]byte, error) {
	if src.CreateDatacollectionclusterResponseAsObject != nil {
		return json.Marshal(&src.CreateDatacollectionclusterResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateDatacollectionclusterResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateDatacollectionclusterResponseAsObject != nil {
		return obj.CreateDatacollectionclusterResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableCreateDatacollectionclusterResponse struct {
	value *CreateDatacollectionclusterResponse
	isSet bool
}

func (v NullableCreateDatacollectionclusterResponse) Get() *CreateDatacollectionclusterResponse {
	return v.value
}

func (v *NullableCreateDatacollectionclusterResponse) Set(val *CreateDatacollectionclusterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDatacollectionclusterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDatacollectionclusterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDatacollectionclusterResponse(val *CreateDatacollectionclusterResponse) *NullableCreateDatacollectionclusterResponse {
	return &NullableCreateDatacollectionclusterResponse{value: val, isSet: true}
}

func (v NullableCreateDatacollectionclusterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDatacollectionclusterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
