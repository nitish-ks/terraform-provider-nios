/*
Infoblox CLOUD API

OpenAPI specification for Infoblox NIOS WAPI CLOUD objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloud

import (
	"encoding/json"
)

// checks if the CreateAzurednstaskgroupResponseAsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAzurednstaskgroupResponseAsObject{}

// CreateAzurednstaskgroupResponseAsObject The response format to create __Azurednstaskgroup__ in object format.
type CreateAzurednstaskgroupResponseAsObject struct {
	Result               *Azurednstaskgroup `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateAzurednstaskgroupResponseAsObject CreateAzurednstaskgroupResponseAsObject

// NewCreateAzurednstaskgroupResponseAsObject instantiates a new CreateAzurednstaskgroupResponseAsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAzurednstaskgroupResponseAsObject() *CreateAzurednstaskgroupResponseAsObject {
	this := CreateAzurednstaskgroupResponseAsObject{}
	return &this
}

// NewCreateAzurednstaskgroupResponseAsObjectWithDefaults instantiates a new CreateAzurednstaskgroupResponseAsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAzurednstaskgroupResponseAsObjectWithDefaults() *CreateAzurednstaskgroupResponseAsObject {
	this := CreateAzurednstaskgroupResponseAsObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CreateAzurednstaskgroupResponseAsObject) GetResult() Azurednstaskgroup {
	if o == nil || IsNil(o.Result) {
		var ret Azurednstaskgroup
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAzurednstaskgroupResponseAsObject) GetResultOk() (*Azurednstaskgroup, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CreateAzurednstaskgroupResponseAsObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Azurednstaskgroup and assigns it to the Result field.
func (o *CreateAzurednstaskgroupResponseAsObject) SetResult(v Azurednstaskgroup) {
	o.Result = &v
}

func (o CreateAzurednstaskgroupResponseAsObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAzurednstaskgroupResponseAsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateAzurednstaskgroupResponseAsObject) UnmarshalJSON(data []byte) (err error) {
	varCreateAzurednstaskgroupResponseAsObject := _CreateAzurednstaskgroupResponseAsObject{}

	err = json.Unmarshal(data, &varCreateAzurednstaskgroupResponseAsObject)

	if err != nil {
		return err
	}

	*o = CreateAzurednstaskgroupResponseAsObject(varCreateAzurednstaskgroupResponseAsObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateAzurednstaskgroupResponseAsObject struct {
	value *CreateAzurednstaskgroupResponseAsObject
	isSet bool
}

func (v NullableCreateAzurednstaskgroupResponseAsObject) Get() *CreateAzurednstaskgroupResponseAsObject {
	return v.value
}

func (v *NullableCreateAzurednstaskgroupResponseAsObject) Set(val *CreateAzurednstaskgroupResponseAsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAzurednstaskgroupResponseAsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAzurednstaskgroupResponseAsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAzurednstaskgroupResponseAsObject(val *CreateAzurednstaskgroupResponseAsObject) *NullableCreateAzurednstaskgroupResponseAsObject {
	return &NullableCreateAzurednstaskgroupResponseAsObject{value: val, isSet: true}
}

func (v NullableCreateAzurednstaskgroupResponseAsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAzurednstaskgroupResponseAsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
