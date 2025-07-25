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

// UpdateRecordRpzNaptrResponse - struct for UpdateRecordRpzNaptrResponse
type UpdateRecordRpzNaptrResponse struct {
	UpdateRecordRpzNaptrResponseAsObject *UpdateRecordRpzNaptrResponseAsObject
	String                               *string
}

// UpdateRecordRpzNaptrResponseAsObjectAsUpdateRecordRpzNaptrResponse is a convenience function that returns UpdateRecordRpzNaptrResponseAsObject wrapped in UpdateRecordRpzNaptrResponse
func UpdateRecordRpzNaptrResponseAsObjectAsUpdateRecordRpzNaptrResponse(v *UpdateRecordRpzNaptrResponseAsObject) UpdateRecordRpzNaptrResponse {
	return UpdateRecordRpzNaptrResponse{
		UpdateRecordRpzNaptrResponseAsObject: v,
	}
}

// stringAsUpdateRecordRpzNaptrResponse is a convenience function that returns string wrapped in UpdateRecordRpzNaptrResponse
func StringAsUpdateRecordRpzNaptrResponse(v *string) UpdateRecordRpzNaptrResponse {
	return UpdateRecordRpzNaptrResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateRecordRpzNaptrResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateRecordRpzNaptrResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdateRecordRpzNaptrResponseAsObject)
	if err == nil {
		jsonUpdateRecordRpzNaptrResponseAsObject, _ := json.Marshal(dst.UpdateRecordRpzNaptrResponseAsObject)
		if string(jsonUpdateRecordRpzNaptrResponseAsObject) == "{}" { // empty struct
			dst.UpdateRecordRpzNaptrResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdateRecordRpzNaptrResponseAsObject = nil
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
		dst.UpdateRecordRpzNaptrResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateRecordRpzNaptrResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateRecordRpzNaptrResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateRecordRpzNaptrResponse) MarshalJSON() ([]byte, error) {
	if src.UpdateRecordRpzNaptrResponseAsObject != nil {
		return json.Marshal(&src.UpdateRecordRpzNaptrResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateRecordRpzNaptrResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateRecordRpzNaptrResponseAsObject != nil {
		return obj.UpdateRecordRpzNaptrResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdateRecordRpzNaptrResponse struct {
	value *UpdateRecordRpzNaptrResponse
	isSet bool
}

func (v NullableUpdateRecordRpzNaptrResponse) Get() *UpdateRecordRpzNaptrResponse {
	return v.value
}

func (v *NullableUpdateRecordRpzNaptrResponse) Set(val *UpdateRecordRpzNaptrResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRecordRpzNaptrResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRecordRpzNaptrResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRecordRpzNaptrResponse(val *UpdateRecordRpzNaptrResponse) *NullableUpdateRecordRpzNaptrResponse {
	return &NullableUpdateRecordRpzNaptrResponse{value: val, isSet: true}
}

func (v NullableUpdateRecordRpzNaptrResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRecordRpzNaptrResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
