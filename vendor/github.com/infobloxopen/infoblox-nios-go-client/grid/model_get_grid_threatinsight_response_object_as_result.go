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

// checks if the GetGridThreatinsightResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetGridThreatinsightResponseObjectAsResult{}

// GetGridThreatinsightResponseObjectAsResult The response format to retrieve __GridThreatinsight__ objects.
type GetGridThreatinsightResponseObjectAsResult struct {
	Result *GridThreatinsight `json:"result,omitempty"`
}

// NewGetGridThreatinsightResponseObjectAsResult instantiates a new GetGridThreatinsightResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGridThreatinsightResponseObjectAsResult() *GetGridThreatinsightResponseObjectAsResult {
	this := GetGridThreatinsightResponseObjectAsResult{}
	return &this
}

// NewGetGridThreatinsightResponseObjectAsResultWithDefaults instantiates a new GetGridThreatinsightResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGridThreatinsightResponseObjectAsResultWithDefaults() *GetGridThreatinsightResponseObjectAsResult {
	this := GetGridThreatinsightResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetGridThreatinsightResponseObjectAsResult) GetResult() GridThreatinsight {
	if o == nil || IsNil(o.Result) {
		var ret GridThreatinsight
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGridThreatinsightResponseObjectAsResult) GetResultOk() (*GridThreatinsight, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetGridThreatinsightResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given GridThreatinsight and assigns it to the Result field.
func (o *GetGridThreatinsightResponseObjectAsResult) SetResult(v GridThreatinsight) {
	o.Result = &v
}

func (o GetGridThreatinsightResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetGridThreatinsightResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetGridThreatinsightResponseObjectAsResult struct {
	value *GetGridThreatinsightResponseObjectAsResult
	isSet bool
}

func (v NullableGetGridThreatinsightResponseObjectAsResult) Get() *GetGridThreatinsightResponseObjectAsResult {
	return v.value
}

func (v *NullableGetGridThreatinsightResponseObjectAsResult) Set(val *GetGridThreatinsightResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGridThreatinsightResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGridThreatinsightResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGridThreatinsightResponseObjectAsResult(val *GetGridThreatinsightResponseObjectAsResult) *NullableGetGridThreatinsightResponseObjectAsResult {
	return &NullableGetGridThreatinsightResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetGridThreatinsightResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGridThreatinsightResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
