/*
Infoblox NOTIFICATION API

OpenAPI specification for Infoblox NIOS WAPI NOTIFICATION objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"encoding/json"
)

// checks if the ListNotificationRestTemplateResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListNotificationRestTemplateResponseObject{}

// ListNotificationRestTemplateResponseObject The response format to retrieve __NotificationRestTemplate__ objects.
type ListNotificationRestTemplateResponseObject struct {
	Result               []NotificationRestTemplate `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListNotificationRestTemplateResponseObject ListNotificationRestTemplateResponseObject

// NewListNotificationRestTemplateResponseObject instantiates a new ListNotificationRestTemplateResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListNotificationRestTemplateResponseObject() *ListNotificationRestTemplateResponseObject {
	this := ListNotificationRestTemplateResponseObject{}
	return &this
}

// NewListNotificationRestTemplateResponseObjectWithDefaults instantiates a new ListNotificationRestTemplateResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListNotificationRestTemplateResponseObjectWithDefaults() *ListNotificationRestTemplateResponseObject {
	this := ListNotificationRestTemplateResponseObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListNotificationRestTemplateResponseObject) GetResult() []NotificationRestTemplate {
	if o == nil || IsNil(o.Result) {
		var ret []NotificationRestTemplate
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListNotificationRestTemplateResponseObject) GetResultOk() ([]NotificationRestTemplate, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListNotificationRestTemplateResponseObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []NotificationRestTemplate and assigns it to the Result field.
func (o *ListNotificationRestTemplateResponseObject) SetResult(v []NotificationRestTemplate) {
	o.Result = v
}

func (o ListNotificationRestTemplateResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListNotificationRestTemplateResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListNotificationRestTemplateResponseObject) UnmarshalJSON(data []byte) (err error) {
	varListNotificationRestTemplateResponseObject := _ListNotificationRestTemplateResponseObject{}

	err = json.Unmarshal(data, &varListNotificationRestTemplateResponseObject)

	if err != nil {
		return err
	}

	*o = ListNotificationRestTemplateResponseObject(varListNotificationRestTemplateResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListNotificationRestTemplateResponseObject struct {
	value *ListNotificationRestTemplateResponseObject
	isSet bool
}

func (v NullableListNotificationRestTemplateResponseObject) Get() *ListNotificationRestTemplateResponseObject {
	return v.value
}

func (v *NullableListNotificationRestTemplateResponseObject) Set(val *ListNotificationRestTemplateResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableListNotificationRestTemplateResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableListNotificationRestTemplateResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListNotificationRestTemplateResponseObject(val *ListNotificationRestTemplateResponseObject) *NullableListNotificationRestTemplateResponseObject {
	return &NullableListNotificationRestTemplateResponseObject{value: val, isSet: true}
}

func (v NullableListNotificationRestTemplateResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListNotificationRestTemplateResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
