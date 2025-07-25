/*
Infoblox SECURITY API

OpenAPI specification for Infoblox NIOS WAPI SECURITY objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package security

import (
	"encoding/json"
)

// checks if the UpdateLocaluserAuthserviceResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateLocaluserAuthserviceResponseAsObject{}

// UpdateLocaluserAuthserviceResponseAsObject The response format to update __LocaluserAuthservice__ in object format.
type UpdateLocaluserAuthserviceResponseAsObject struct {
	Result               *LocaluserAuthservice `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateLocaluserAuthserviceResponseAsObject UpdateLocaluserAuthserviceResponseAsObject

// NewUpdateLocaluserAuthserviceResponseAsObject instantiates a new UpdateLocaluserAuthserviceResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLocaluserAuthserviceResponseAsObject() *UpdateLocaluserAuthserviceResponseAsObject {
	this := UpdateLocaluserAuthserviceResponseAsObject{}
	return &this
}

// NewUpdateLocaluserAuthserviceResponseAsObjectWithDefaults instantiates a new UpdateLocaluserAuthserviceResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLocaluserAuthserviceResponseAsObjectWithDefaults() *UpdateLocaluserAuthserviceResponseAsObject {
	this := UpdateLocaluserAuthserviceResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateLocaluserAuthserviceResponseAsObject) GetResult() LocaluserAuthservice {
	if o == nil || IsNil(o.Result) {
		var ret LocaluserAuthservice
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLocaluserAuthserviceResponseAsObject) GetResultOk() (*LocaluserAuthservice, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateLocaluserAuthserviceResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given LocaluserAuthservice and assigns it to the Result field.
func (o *UpdateLocaluserAuthserviceResponseAsObject) SetResult(v LocaluserAuthservice) {
	o.Result = &v
}

func (o UpdateLocaluserAuthserviceResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateLocaluserAuthserviceResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateLocaluserAuthserviceResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateLocaluserAuthserviceResponseAsObject := _UpdateLocaluserAuthserviceResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateLocaluserAuthserviceResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateLocaluserAuthserviceResponseAsObject(varUpdateLocaluserAuthserviceResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateLocaluserAuthserviceResponseAsObject struct {
	value *UpdateLocaluserAuthserviceResponseAsObject
	isSet bool
}

func (v NullableUpdateLocaluserAuthserviceResponseAsObject) Get() *UpdateLocaluserAuthserviceResponseAsObject {
	return v.value
}

func (v *NullableUpdateLocaluserAuthserviceResponseAsObject) Set(val *UpdateLocaluserAuthserviceResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLocaluserAuthserviceResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLocaluserAuthserviceResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLocaluserAuthserviceResponseAsObject(val *UpdateLocaluserAuthserviceResponseAsObject) *NullableUpdateLocaluserAuthserviceResponseAsObject {
	return &NullableUpdateLocaluserAuthserviceResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateLocaluserAuthserviceResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLocaluserAuthserviceResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
