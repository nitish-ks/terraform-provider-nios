/*
Infoblox IPAM API

OpenAPI specification for Infoblox NIOS WAPI IPAM objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the GetNetworkviewResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkviewResponseObjectAsResult{}

// GetNetworkviewResponseObjectAsResult The response format to retrieve __Networkview__ objects.
type GetNetworkviewResponseObjectAsResult struct {
	Result *Networkview `json:"result,omitempty"`
}

// NewGetNetworkviewResponseObjectAsResult instantiates a new GetNetworkviewResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkviewResponseObjectAsResult() *GetNetworkviewResponseObjectAsResult {
	this := GetNetworkviewResponseObjectAsResult{}
	return &this
}

// NewGetNetworkviewResponseObjectAsResultWithDefaults instantiates a new GetNetworkviewResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkviewResponseObjectAsResultWithDefaults() *GetNetworkviewResponseObjectAsResult {
	this := GetNetworkviewResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetNetworkviewResponseObjectAsResult) GetResult() Networkview {
	if o == nil || IsNil(o.Result) {
		var ret Networkview
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkviewResponseObjectAsResult) GetResultOk() (*Networkview, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetNetworkviewResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Networkview and assigns it to the Result field.
func (o *GetNetworkviewResponseObjectAsResult) SetResult(v Networkview) {
	o.Result = &v
}

func (o GetNetworkviewResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkviewResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetNetworkviewResponseObjectAsResult struct {
	value *GetNetworkviewResponseObjectAsResult
	isSet bool
}

func (v NullableGetNetworkviewResponseObjectAsResult) Get() *GetNetworkviewResponseObjectAsResult {
	return v.value
}

func (v *NullableGetNetworkviewResponseObjectAsResult) Set(val *GetNetworkviewResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkviewResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkviewResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkviewResponseObjectAsResult(val *GetNetworkviewResponseObjectAsResult) *NullableGetNetworkviewResponseObjectAsResult {
	return &NullableGetNetworkviewResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetNetworkviewResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkviewResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
