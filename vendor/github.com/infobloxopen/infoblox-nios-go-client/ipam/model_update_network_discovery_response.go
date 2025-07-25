/*
Infoblox IPAM API

OpenAPI specification for Infoblox NIOS WAPI IPAM objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
	"fmt"
)

// UpdateNetworkDiscoveryResponse - struct for UpdateNetworkDiscoveryResponse
type UpdateNetworkDiscoveryResponse struct {
	UpdateNetworkDiscoveryResponseAsObject *UpdateNetworkDiscoveryResponseAsObject
	String                                 *string
}

// UpdateNetworkDiscoveryResponseAsObjectAsUpdateNetworkDiscoveryResponse is a convenience function that returns UpdateNetworkDiscoveryResponseAsObject wrapped in UpdateNetworkDiscoveryResponse
func UpdateNetworkDiscoveryResponseAsObjectAsUpdateNetworkDiscoveryResponse(v *UpdateNetworkDiscoveryResponseAsObject) UpdateNetworkDiscoveryResponse {
	return UpdateNetworkDiscoveryResponse{
		UpdateNetworkDiscoveryResponseAsObject: v,
	}
}

// stringAsUpdateNetworkDiscoveryResponse is a convenience function that returns string wrapped in UpdateNetworkDiscoveryResponse
func StringAsUpdateNetworkDiscoveryResponse(v *string) UpdateNetworkDiscoveryResponse {
	return UpdateNetworkDiscoveryResponse{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateNetworkDiscoveryResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into UpdateNetworkDiscoveryResponseAsObject
	err = newStrictDecoder(data).Decode(&dst.UpdateNetworkDiscoveryResponseAsObject)
	if err == nil {
		jsonUpdateNetworkDiscoveryResponseAsObject, _ := json.Marshal(dst.UpdateNetworkDiscoveryResponseAsObject)
		if string(jsonUpdateNetworkDiscoveryResponseAsObject) == "{}" { // empty struct
			dst.UpdateNetworkDiscoveryResponseAsObject = nil
		} else {
			match++
		}
	} else {
		dst.UpdateNetworkDiscoveryResponseAsObject = nil
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
		dst.UpdateNetworkDiscoveryResponseAsObject = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateNetworkDiscoveryResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateNetworkDiscoveryResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateNetworkDiscoveryResponse) MarshalJSON() ([]byte, error) {
	if src.UpdateNetworkDiscoveryResponseAsObject != nil {
		return json.Marshal(&src.UpdateNetworkDiscoveryResponseAsObject)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateNetworkDiscoveryResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.UpdateNetworkDiscoveryResponseAsObject != nil {
		return obj.UpdateNetworkDiscoveryResponseAsObject
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableUpdateNetworkDiscoveryResponse struct {
	value *UpdateNetworkDiscoveryResponse
	isSet bool
}

func (v NullableUpdateNetworkDiscoveryResponse) Get() *UpdateNetworkDiscoveryResponse {
	return v.value
}

func (v *NullableUpdateNetworkDiscoveryResponse) Set(val *UpdateNetworkDiscoveryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkDiscoveryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkDiscoveryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkDiscoveryResponse(val *UpdateNetworkDiscoveryResponse) *NullableUpdateNetworkDiscoveryResponse {
	return &NullableUpdateNetworkDiscoveryResponse{value: val, isSet: true}
}

func (v NullableUpdateNetworkDiscoveryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkDiscoveryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
