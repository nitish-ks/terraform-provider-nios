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

// checks if the ListRecordDhcidResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRecordDhcidResponseObject{}

// ListRecordDhcidResponseObject The response format to retrieve __RecordDhcid__ objects.
type ListRecordDhcidResponseObject struct {
	Result               []RecordDhcid `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListRecordDhcidResponseObject ListRecordDhcidResponseObject

// NewListRecordDhcidResponseObject instantiates a new ListRecordDhcidResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRecordDhcidResponseObject() *ListRecordDhcidResponseObject {
	this := ListRecordDhcidResponseObject{}
	return &this
}

// NewListRecordDhcidResponseObjectWithDefaults instantiates a new ListRecordDhcidResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRecordDhcidResponseObjectWithDefaults() *ListRecordDhcidResponseObject {
	this := ListRecordDhcidResponseObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListRecordDhcidResponseObject) GetResult() []RecordDhcid {
	if o == nil || IsNil(o.Result) {
		var ret []RecordDhcid
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRecordDhcidResponseObject) GetResultOk() ([]RecordDhcid, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListRecordDhcidResponseObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []RecordDhcid and assigns it to the Result field.
func (o *ListRecordDhcidResponseObject) SetResult(v []RecordDhcid) {
	o.Result = v
}

func (o ListRecordDhcidResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRecordDhcidResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListRecordDhcidResponseObject) UnmarshalJSON(data []byte) (err error) {
	varListRecordDhcidResponseObject := _ListRecordDhcidResponseObject{}

	err = json.Unmarshal(data, &varListRecordDhcidResponseObject)

	if err != nil {
		return err
	}

	*o = ListRecordDhcidResponseObject(varListRecordDhcidResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListRecordDhcidResponseObject struct {
	value *ListRecordDhcidResponseObject
	isSet bool
}

func (v NullableListRecordDhcidResponseObject) Get() *ListRecordDhcidResponseObject {
	return v.value
}

func (v *NullableListRecordDhcidResponseObject) Set(val *ListRecordDhcidResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableListRecordDhcidResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableListRecordDhcidResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRecordDhcidResponseObject(val *ListRecordDhcidResponseObject) *NullableListRecordDhcidResponseObject {
	return &NullableListRecordDhcidResponseObject{value: val, isSet: true}
}

func (v NullableListRecordDhcidResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRecordDhcidResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
