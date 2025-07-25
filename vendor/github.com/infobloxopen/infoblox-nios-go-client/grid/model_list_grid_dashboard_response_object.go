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

// checks if the ListGridDashboardResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListGridDashboardResponseObject{}

// ListGridDashboardResponseObject The response format to retrieve __GridDashboard__ objects.
type ListGridDashboardResponseObject struct {
	Result               []GridDashboard `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListGridDashboardResponseObject ListGridDashboardResponseObject

// NewListGridDashboardResponseObject instantiates a new ListGridDashboardResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListGridDashboardResponseObject() *ListGridDashboardResponseObject {
	this := ListGridDashboardResponseObject{}
	return &this
}

// NewListGridDashboardResponseObjectWithDefaults instantiates a new ListGridDashboardResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListGridDashboardResponseObjectWithDefaults() *ListGridDashboardResponseObject {
	this := ListGridDashboardResponseObject{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ListGridDashboardResponseObject) GetResult() []GridDashboard {
	if o == nil || IsNil(o.Result) {
		var ret []GridDashboard
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGridDashboardResponseObject) GetResultOk() ([]GridDashboard, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ListGridDashboardResponseObject) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []GridDashboard and assigns it to the Result field.
func (o *ListGridDashboardResponseObject) SetResult(v []GridDashboard) {
	o.Result = v
}

func (o ListGridDashboardResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListGridDashboardResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListGridDashboardResponseObject) UnmarshalJSON(data []byte) (err error) {
	varListGridDashboardResponseObject := _ListGridDashboardResponseObject{}

	err = json.Unmarshal(data, &varListGridDashboardResponseObject)

	if err != nil {
		return err
	}

	*o = ListGridDashboardResponseObject(varListGridDashboardResponseObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListGridDashboardResponseObject struct {
	value *ListGridDashboardResponseObject
	isSet bool
}

func (v NullableListGridDashboardResponseObject) Get() *ListGridDashboardResponseObject {
	return v.value
}

func (v *NullableListGridDashboardResponseObject) Set(val *ListGridDashboardResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableListGridDashboardResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableListGridDashboardResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGridDashboardResponseObject(val *ListGridDashboardResponseObject) *NullableListGridDashboardResponseObject {
	return &NullableListGridDashboardResponseObject{value: val, isSet: true}
}

func (v NullableListGridDashboardResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListGridDashboardResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
