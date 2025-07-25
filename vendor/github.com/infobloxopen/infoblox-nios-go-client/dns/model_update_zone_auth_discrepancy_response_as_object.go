/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the UpdateZoneAuthDiscrepancyResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateZoneAuthDiscrepancyResponseAsObject{}

// UpdateZoneAuthDiscrepancyResponseAsObject The response format to update __ZoneAuthDiscrepancy__ in object format.
type UpdateZoneAuthDiscrepancyResponseAsObject struct {
	Result               *ZoneAuthDiscrepancy `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateZoneAuthDiscrepancyResponseAsObject UpdateZoneAuthDiscrepancyResponseAsObject

// NewUpdateZoneAuthDiscrepancyResponseAsObject instantiates a new UpdateZoneAuthDiscrepancyResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateZoneAuthDiscrepancyResponseAsObject() *UpdateZoneAuthDiscrepancyResponseAsObject {
	this := UpdateZoneAuthDiscrepancyResponseAsObject{}
	return &this
}

// NewUpdateZoneAuthDiscrepancyResponseAsObjectWithDefaults instantiates a new UpdateZoneAuthDiscrepancyResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateZoneAuthDiscrepancyResponseAsObjectWithDefaults() *UpdateZoneAuthDiscrepancyResponseAsObject {
	this := UpdateZoneAuthDiscrepancyResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateZoneAuthDiscrepancyResponseAsObject) GetResult() ZoneAuthDiscrepancy {
	if o == nil || IsNil(o.Result) {
		var ret ZoneAuthDiscrepancy
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateZoneAuthDiscrepancyResponseAsObject) GetResultOk() (*ZoneAuthDiscrepancy, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateZoneAuthDiscrepancyResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ZoneAuthDiscrepancy and assigns it to the Result field.
func (o *UpdateZoneAuthDiscrepancyResponseAsObject) SetResult(v ZoneAuthDiscrepancy) {
	o.Result = &v
}

func (o UpdateZoneAuthDiscrepancyResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateZoneAuthDiscrepancyResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateZoneAuthDiscrepancyResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateZoneAuthDiscrepancyResponseAsObject := _UpdateZoneAuthDiscrepancyResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateZoneAuthDiscrepancyResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateZoneAuthDiscrepancyResponseAsObject(varUpdateZoneAuthDiscrepancyResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateZoneAuthDiscrepancyResponseAsObject struct {
	value *UpdateZoneAuthDiscrepancyResponseAsObject
	isSet bool
}

func (v NullableUpdateZoneAuthDiscrepancyResponseAsObject) Get() *UpdateZoneAuthDiscrepancyResponseAsObject {
	return v.value
}

func (v *NullableUpdateZoneAuthDiscrepancyResponseAsObject) Set(val *UpdateZoneAuthDiscrepancyResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateZoneAuthDiscrepancyResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateZoneAuthDiscrepancyResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateZoneAuthDiscrepancyResponseAsObject(val *UpdateZoneAuthDiscrepancyResponseAsObject) *NullableUpdateZoneAuthDiscrepancyResponseAsObject {
	return &NullableUpdateZoneAuthDiscrepancyResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateZoneAuthDiscrepancyResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateZoneAuthDiscrepancyResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
