/*
Infoblox RPZ API

OpenAPI specification for Infoblox NIOS WAPI RPZ objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rpz

import (
	"encoding/json"
)

// checks if the UpdateRecordRpzTxtResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRecordRpzTxtResponseAsObject{}

// UpdateRecordRpzTxtResponseAsObject The response format to update __RecordRpzTxt__ in object format.
type UpdateRecordRpzTxtResponseAsObject struct {
	Result               *RecordRpzTxt `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateRecordRpzTxtResponseAsObject UpdateRecordRpzTxtResponseAsObject

// NewUpdateRecordRpzTxtResponseAsObject instantiates a new UpdateRecordRpzTxtResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRecordRpzTxtResponseAsObject() *UpdateRecordRpzTxtResponseAsObject {
	this := UpdateRecordRpzTxtResponseAsObject{}
	return &this
}

// NewUpdateRecordRpzTxtResponseAsObjectWithDefaults instantiates a new UpdateRecordRpzTxtResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRecordRpzTxtResponseAsObjectWithDefaults() *UpdateRecordRpzTxtResponseAsObject {
	this := UpdateRecordRpzTxtResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateRecordRpzTxtResponseAsObject) GetResult() RecordRpzTxt {
	if o == nil || IsNil(o.Result) {
		var ret RecordRpzTxt
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRecordRpzTxtResponseAsObject) GetResultOk() (*RecordRpzTxt, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateRecordRpzTxtResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given RecordRpzTxt and assigns it to the Result field.
func (o *UpdateRecordRpzTxtResponseAsObject) SetResult(v RecordRpzTxt) {
	o.Result = &v
}

func (o UpdateRecordRpzTxtResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRecordRpzTxtResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateRecordRpzTxtResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateRecordRpzTxtResponseAsObject := _UpdateRecordRpzTxtResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateRecordRpzTxtResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateRecordRpzTxtResponseAsObject(varUpdateRecordRpzTxtResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateRecordRpzTxtResponseAsObject struct {
	value *UpdateRecordRpzTxtResponseAsObject
	isSet bool
}

func (v NullableUpdateRecordRpzTxtResponseAsObject) Get() *UpdateRecordRpzTxtResponseAsObject {
	return v.value
}

func (v *NullableUpdateRecordRpzTxtResponseAsObject) Set(val *UpdateRecordRpzTxtResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRecordRpzTxtResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRecordRpzTxtResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRecordRpzTxtResponseAsObject(val *UpdateRecordRpzTxtResponseAsObject) *NullableUpdateRecordRpzTxtResponseAsObject {
	return &NullableUpdateRecordRpzTxtResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateRecordRpzTxtResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRecordRpzTxtResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
