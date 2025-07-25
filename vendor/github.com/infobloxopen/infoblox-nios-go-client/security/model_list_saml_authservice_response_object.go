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

// checks if the ListSamlAuthserviceResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSamlAuthserviceResponseObject{}

// ListSamlAuthserviceResponseObject The response format to retrieve __SamlAuthservice__ objects.
type ListSamlAuthserviceResponseObject struct {
	Result               []SamlAuthservice `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListSamlAuthserviceResponseObject ListSamlAuthserviceResponseObject

// NewListSamlAuthserviceResponseObject instantiates a new ListSamlAuthserviceResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSamlAuthserviceResponseObject() *ListSamlAuthserviceResponseObject {
	this := ListSamlAuthserviceResponseObject{}
	return &this
}

// NewListSamlAuthserviceResponseObjectWithDefaults instantiates a new ListSamlAuthserviceResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSamlAuthserviceResponseObjectWithDefaults() *ListSamlAuthserviceResponseObject {
	this := ListSamlAuthserviceResponseObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListSamlAuthserviceResponseObject) GetResult() []SamlAuthservice {
	if o == nil || IsNil(o.Result) {
		var ret []SamlAuthservice
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSamlAuthserviceResponseObject) GetResultOk() ([]SamlAuthservice, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListSamlAuthserviceResponseObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []SamlAuthservice and assigns it to the Result field.
func (o *ListSamlAuthserviceResponseObject) SetResult(v []SamlAuthservice) {
	o.Result = v
}

func (o ListSamlAuthserviceResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSamlAuthserviceResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListSamlAuthserviceResponseObject) UnmarshalJSON(data []byte) (err error) {
	varListSamlAuthserviceResponseObject := _ListSamlAuthserviceResponseObject{}

	err = json.Unmarshal(data, &varListSamlAuthserviceResponseObject)

	if err != nil {
		return err
	}

	*o = ListSamlAuthserviceResponseObject(varListSamlAuthserviceResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListSamlAuthserviceResponseObject struct {
	value *ListSamlAuthserviceResponseObject
	isSet bool
}

func (v NullableListSamlAuthserviceResponseObject) Get() *ListSamlAuthserviceResponseObject {
	return v.value
}

func (v *NullableListSamlAuthserviceResponseObject) Set(val *ListSamlAuthserviceResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableListSamlAuthserviceResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableListSamlAuthserviceResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSamlAuthserviceResponseObject(val *ListSamlAuthserviceResponseObject) *NullableListSamlAuthserviceResponseObject {
	return &NullableListSamlAuthserviceResponseObject{value: val, isSet: true}
}

func (v NullableListSamlAuthserviceResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSamlAuthserviceResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
