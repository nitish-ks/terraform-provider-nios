/*
Infoblox GRID API

OpenAPI specification for Infoblox NIOS WAPI GRID objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grid

import (
	"encoding/json"
)

// checks if the UpdateMemberResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMemberResponseAsObject{}

// UpdateMemberResponseAsObject The response format to update __Member__ in object format.
type UpdateMemberResponseAsObject struct {
	Result               *Member `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateMemberResponseAsObject UpdateMemberResponseAsObject

// NewUpdateMemberResponseAsObject instantiates a new UpdateMemberResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMemberResponseAsObject() *UpdateMemberResponseAsObject {
	this := UpdateMemberResponseAsObject{}
	return &this
}

// NewUpdateMemberResponseAsObjectWithDefaults instantiates a new UpdateMemberResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMemberResponseAsObjectWithDefaults() *UpdateMemberResponseAsObject {
	this := UpdateMemberResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateMemberResponseAsObject) GetResult() Member {
	if o == nil || IsNil(o.Result) {
		var ret Member
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMemberResponseAsObject) GetResultOk() (*Member, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateMemberResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Member and assigns it to the Result field.
func (o *UpdateMemberResponseAsObject) SetResult(v Member) {
	o.Result = &v
}

func (o UpdateMemberResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMemberResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateMemberResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varUpdateMemberResponseAsObject := _UpdateMemberResponseAsObject{}

	err = json.Unmarshal(data, &varUpdateMemberResponseAsObject)

	if err != nil {
		return err
	}

	*o = UpdateMemberResponseAsObject(varUpdateMemberResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateMemberResponseAsObject struct {
	value *UpdateMemberResponseAsObject
	isSet bool
}

func (v NullableUpdateMemberResponseAsObject) Get() *UpdateMemberResponseAsObject {
	return v.value
}

func (v *NullableUpdateMemberResponseAsObject) Set(val *UpdateMemberResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMemberResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMemberResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMemberResponseAsObject(val *UpdateMemberResponseAsObject) *NullableUpdateMemberResponseAsObject {
	return &NullableUpdateMemberResponseAsObject{value: val, isSet: true}
}

func (v NullableUpdateMemberResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMemberResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
