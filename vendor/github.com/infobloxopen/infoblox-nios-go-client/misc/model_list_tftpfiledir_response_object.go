/*
Infoblox MISC API

OpenAPI specification for Infoblox NIOS WAPI MISC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package misc

import (
	"encoding/json"
)

// checks if the ListTftpfiledirResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTftpfiledirResponseObject{}

// ListTftpfiledirResponseObject The response format to retrieve __Tftpfiledir__ objects.
type ListTftpfiledirResponseObject struct {
	Result               []Tftpfiledir `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListTftpfiledirResponseObject ListTftpfiledirResponseObject

// NewListTftpfiledirResponseObject instantiates a new ListTftpfiledirResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTftpfiledirResponseObject() *ListTftpfiledirResponseObject {
	this := ListTftpfiledirResponseObject{}
	return &this
}

// NewListTftpfiledirResponseObjectWithDefaults instantiates a new ListTftpfiledirResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTftpfiledirResponseObjectWithDefaults() *ListTftpfiledirResponseObject {
	this := ListTftpfiledirResponseObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListTftpfiledirResponseObject) GetResult() []Tftpfiledir {
	if o == nil || IsNil(o.Result) {
		var ret []Tftpfiledir
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTftpfiledirResponseObject) GetResultOk() ([]Tftpfiledir, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListTftpfiledirResponseObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []Tftpfiledir and assigns it to the Result field.
func (o *ListTftpfiledirResponseObject) SetResult(v []Tftpfiledir) {
	o.Result = v
}

func (o ListTftpfiledirResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTftpfiledirResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListTftpfiledirResponseObject) UnmarshalJSON(data []byte) (err error) {
	varListTftpfiledirResponseObject := _ListTftpfiledirResponseObject{}

	err = json.Unmarshal(data, &varListTftpfiledirResponseObject)

	if err != nil {
		return err
	}

	*o = ListTftpfiledirResponseObject(varListTftpfiledirResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListTftpfiledirResponseObject struct {
	value *ListTftpfiledirResponseObject
	isSet bool
}

func (v NullableListTftpfiledirResponseObject) Get() *ListTftpfiledirResponseObject {
	return v.value
}

func (v *NullableListTftpfiledirResponseObject) Set(val *ListTftpfiledirResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableListTftpfiledirResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableListTftpfiledirResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTftpfiledirResponseObject(val *ListTftpfiledirResponseObject) *NullableListTftpfiledirResponseObject {
	return &NullableListTftpfiledirResponseObject{value: val, isSet: true}
}

func (v NullableListTftpfiledirResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTftpfiledirResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
