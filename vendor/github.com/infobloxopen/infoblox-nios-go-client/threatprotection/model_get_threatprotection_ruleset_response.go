/*
Infoblox THREATPROTECTION API

OpenAPI specification for Infoblox NIOS WAPI THREATPROTECTION objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package threatprotection

import (
	"encoding/json"
	"fmt"
)

// GetThreatprotectionRulesetResponse - struct for GetThreatprotectionRulesetResponse
type GetThreatprotectionRulesetResponse struct {
	GetThreatprotectionRulesetResponseObjectAsResult *GetThreatprotectionRulesetResponseObjectAsResult
	ThreatprotectionRuleset                          *ThreatprotectionRuleset
}

// GetThreatprotectionRulesetResponseObjectAsResultAsGetThreatprotectionRulesetResponse is a convenience function that returns GetThreatprotectionRulesetResponseObjectAsResult wrapped in GetThreatprotectionRulesetResponse
func GetThreatprotectionRulesetResponseObjectAsResultAsGetThreatprotectionRulesetResponse(v *GetThreatprotectionRulesetResponseObjectAsResult) GetThreatprotectionRulesetResponse {
	return GetThreatprotectionRulesetResponse{
		GetThreatprotectionRulesetResponseObjectAsResult: v,
	}
}

// ThreatprotectionRulesetAsGetThreatprotectionRulesetResponse is a convenience function that returns ThreatprotectionRuleset wrapped in GetThreatprotectionRulesetResponse
func ThreatprotectionRulesetAsGetThreatprotectionRulesetResponse(v *ThreatprotectionRuleset) GetThreatprotectionRulesetResponse {
	return GetThreatprotectionRulesetResponse{
		ThreatprotectionRuleset: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetThreatprotectionRulesetResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GetThreatprotectionRulesetResponseObjectAsResult
	err = newStrictDecoder(data).Decode(&dst.GetThreatprotectionRulesetResponseObjectAsResult)
	if err == nil {
		jsonGetThreatprotectionRulesetResponseObjectAsResult, _ := json.Marshal(dst.GetThreatprotectionRulesetResponseObjectAsResult)
		if string(jsonGetThreatprotectionRulesetResponseObjectAsResult) == "{}" { // empty struct
			dst.GetThreatprotectionRulesetResponseObjectAsResult = nil
		} else {
			match++
		}
	} else {
		dst.GetThreatprotectionRulesetResponseObjectAsResult = nil
	}

	// try to unmarshal data into ThreatprotectionRuleset
	err = newStrictDecoder(data).Decode(&dst.ThreatprotectionRuleset)
	if err == nil {
		jsonThreatprotectionRuleset, _ := json.Marshal(dst.ThreatprotectionRuleset)
		if string(jsonThreatprotectionRuleset) == "{}" { // empty struct
			dst.ThreatprotectionRuleset = nil
		} else {
			match++
		}
	} else {
		dst.ThreatprotectionRuleset = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GetThreatprotectionRulesetResponseObjectAsResult = nil
		dst.ThreatprotectionRuleset = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetThreatprotectionRulesetResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetThreatprotectionRulesetResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetThreatprotectionRulesetResponse) MarshalJSON() ([]byte, error) {
	if src.GetThreatprotectionRulesetResponseObjectAsResult != nil {
		return json.Marshal(&src.GetThreatprotectionRulesetResponseObjectAsResult)
	}

	if src.ThreatprotectionRuleset != nil {
		return json.Marshal(&src.ThreatprotectionRuleset)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetThreatprotectionRulesetResponse) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GetThreatprotectionRulesetResponseObjectAsResult != nil {
		return obj.GetThreatprotectionRulesetResponseObjectAsResult
	}

	if obj.ThreatprotectionRuleset != nil {
		return obj.ThreatprotectionRuleset
	}

	// all schemas are nil
	return nil
}

type NullableGetThreatprotectionRulesetResponse struct {
	value *GetThreatprotectionRulesetResponse
	isSet bool
}

func (v NullableGetThreatprotectionRulesetResponse) Get() *GetThreatprotectionRulesetResponse {
	return v.value
}

func (v *NullableGetThreatprotectionRulesetResponse) Set(val *GetThreatprotectionRulesetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetThreatprotectionRulesetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetThreatprotectionRulesetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetThreatprotectionRulesetResponse(val *GetThreatprotectionRulesetResponse) *NullableGetThreatprotectionRulesetResponse {
	return &NullableGetThreatprotectionRulesetResponse{value: val, isSet: true}
}

func (v NullableGetThreatprotectionRulesetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetThreatprotectionRulesetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
