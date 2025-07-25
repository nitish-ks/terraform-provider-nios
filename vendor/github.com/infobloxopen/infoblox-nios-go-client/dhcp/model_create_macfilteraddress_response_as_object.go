/*
Infoblox DHCP API

OpenAPI specification for Infoblox NIOS WAPI DHCP objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dhcp

import (
	"encoding/json"
)

// checks if the CreateMacfilteraddressResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMacfilteraddressResponseAsObject{}

// CreateMacfilteraddressResponseAsObject The response format to create __Macfilteraddress__ in object format.
type CreateMacfilteraddressResponseAsObject struct {
	Result               *Macfilteraddress `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateMacfilteraddressResponseAsObject CreateMacfilteraddressResponseAsObject

// NewCreateMacfilteraddressResponseAsObject instantiates a new CreateMacfilteraddressResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMacfilteraddressResponseAsObject() *CreateMacfilteraddressResponseAsObject {
	this := CreateMacfilteraddressResponseAsObject{}
	return &this
}

// NewCreateMacfilteraddressResponseAsObjectWithDefaults instantiates a new CreateMacfilteraddressResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMacfilteraddressResponseAsObjectWithDefaults() *CreateMacfilteraddressResponseAsObject {
	this := CreateMacfilteraddressResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateMacfilteraddressResponseAsObject) GetResult() Macfilteraddress {
	if o == nil || IsNil(o.Result) {
		var ret Macfilteraddress
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMacfilteraddressResponseAsObject) GetResultOk() (*Macfilteraddress, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateMacfilteraddressResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Macfilteraddress and assigns it to the Result field.
func (o *CreateMacfilteraddressResponseAsObject) SetResult(v Macfilteraddress) {
	o.Result = &v
}

func (o CreateMacfilteraddressResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMacfilteraddressResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateMacfilteraddressResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varCreateMacfilteraddressResponseAsObject := _CreateMacfilteraddressResponseAsObject{}

	err = json.Unmarshal(data, &varCreateMacfilteraddressResponseAsObject)

	if err != nil {
		return err
	}

	*o = CreateMacfilteraddressResponseAsObject(varCreateMacfilteraddressResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateMacfilteraddressResponseAsObject struct {
	value *CreateMacfilteraddressResponseAsObject
	isSet bool
}

func (v NullableCreateMacfilteraddressResponseAsObject) Get() *CreateMacfilteraddressResponseAsObject {
	return v.value
}

func (v *NullableCreateMacfilteraddressResponseAsObject) Set(val *CreateMacfilteraddressResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMacfilteraddressResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMacfilteraddressResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMacfilteraddressResponseAsObject(val *CreateMacfilteraddressResponseAsObject) *NullableCreateMacfilteraddressResponseAsObject {
	return &NullableCreateMacfilteraddressResponseAsObject{value: val, isSet: true}
}

func (v NullableCreateMacfilteraddressResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMacfilteraddressResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
