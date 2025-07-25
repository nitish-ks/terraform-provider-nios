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

// checks if the GetRecordRpzAaaaResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRecordRpzAaaaResponseObjectAsResult{}

// GetRecordRpzAaaaResponseObjectAsResult The response format to retrieve __RecordRpzAaaa__ objects.
type GetRecordRpzAaaaResponseObjectAsResult struct {
	Result *RecordRpzAaaa `json:"result,omitempty"`
}

// NewGetRecordRpzAaaaResponseObjectAsResult instantiates a new GetRecordRpzAaaaResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecordRpzAaaaResponseObjectAsResult() *GetRecordRpzAaaaResponseObjectAsResult {
	this := GetRecordRpzAaaaResponseObjectAsResult{}
	return &this
}

// NewGetRecordRpzAaaaResponseObjectAsResultWithDefaults instantiates a new GetRecordRpzAaaaResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecordRpzAaaaResponseObjectAsResultWithDefaults() *GetRecordRpzAaaaResponseObjectAsResult {
	this := GetRecordRpzAaaaResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetRecordRpzAaaaResponseObjectAsResult) GetResult() RecordRpzAaaa {
	if o == nil || IsNil(o.Result) {
		var ret RecordRpzAaaa
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordRpzAaaaResponseObjectAsResult) GetResultOk() (*RecordRpzAaaa, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetRecordRpzAaaaResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given RecordRpzAaaa and assigns it to the Result field.
func (o *GetRecordRpzAaaaResponseObjectAsResult) SetResult(v RecordRpzAaaa) {
	o.Result = &v
}

func (o GetRecordRpzAaaaResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRecordRpzAaaaResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetRecordRpzAaaaResponseObjectAsResult struct {
	value *GetRecordRpzAaaaResponseObjectAsResult
	isSet bool
}

func (v NullableGetRecordRpzAaaaResponseObjectAsResult) Get() *GetRecordRpzAaaaResponseObjectAsResult {
	return v.value
}

func (v *NullableGetRecordRpzAaaaResponseObjectAsResult) Set(val *GetRecordRpzAaaaResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecordRpzAaaaResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecordRpzAaaaResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecordRpzAaaaResponseObjectAsResult(val *GetRecordRpzAaaaResponseObjectAsResult) *NullableGetRecordRpzAaaaResponseObjectAsResult {
	return &NullableGetRecordRpzAaaaResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetRecordRpzAaaaResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecordRpzAaaaResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
