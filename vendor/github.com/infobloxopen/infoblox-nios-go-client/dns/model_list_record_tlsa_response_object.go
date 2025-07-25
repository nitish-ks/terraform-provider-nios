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

// checks if the ListRecordTlsaResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRecordTlsaResponseObject{}

// ListRecordTlsaResponseObject The response format to retrieve __RecordTlsa__ objects.
type ListRecordTlsaResponseObject struct {
	Result               []RecordTlsa `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListRecordTlsaResponseObject ListRecordTlsaResponseObject

// NewListRecordTlsaResponseObject instantiates a new ListRecordTlsaResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRecordTlsaResponseObject() *ListRecordTlsaResponseObject {
	this := ListRecordTlsaResponseObject{}
	return &this
}

// NewListRecordTlsaResponseObjectWithDefaults instantiates a new ListRecordTlsaResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRecordTlsaResponseObjectWithDefaults() *ListRecordTlsaResponseObject {
	this := ListRecordTlsaResponseObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListRecordTlsaResponseObject) GetResult() []RecordTlsa {
	if o == nil || IsNil(o.Result) {
		var ret []RecordTlsa
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRecordTlsaResponseObject) GetResultOk() ([]RecordTlsa, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListRecordTlsaResponseObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []RecordTlsa and assigns it to the Result field.
func (o *ListRecordTlsaResponseObject) SetResult(v []RecordTlsa) {
	o.Result = v
}

func (o ListRecordTlsaResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRecordTlsaResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListRecordTlsaResponseObject) UnmarshalJSON(data []byte) (err error) {
	varListRecordTlsaResponseObject := _ListRecordTlsaResponseObject{}

	err = json.Unmarshal(data, &varListRecordTlsaResponseObject)

	if err != nil {
		return err
	}

	*o = ListRecordTlsaResponseObject(varListRecordTlsaResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListRecordTlsaResponseObject struct {
	value *ListRecordTlsaResponseObject
	isSet bool
}

func (v NullableListRecordTlsaResponseObject) Get() *ListRecordTlsaResponseObject {
	return v.value
}

func (v *NullableListRecordTlsaResponseObject) Set(val *ListRecordTlsaResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableListRecordTlsaResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableListRecordTlsaResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRecordTlsaResponseObject(val *ListRecordTlsaResponseObject) *NullableListRecordTlsaResponseObject {
	return &NullableListRecordTlsaResponseObject{value: val, isSet: true}
}

func (v NullableListRecordTlsaResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRecordTlsaResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
