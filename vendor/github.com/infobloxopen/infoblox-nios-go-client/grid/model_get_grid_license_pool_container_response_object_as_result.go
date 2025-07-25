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

// checks if the GetGridLicensePoolContainerResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetGridLicensePoolContainerResponseObjectAsResult{}

// GetGridLicensePoolContainerResponseObjectAsResult The response format to retrieve __GridLicensePoolContainer__ objects.
type GetGridLicensePoolContainerResponseObjectAsResult struct {
	Result *GridLicensePoolContainer `json:"result,omitempty"`
}

// NewGetGridLicensePoolContainerResponseObjectAsResult instantiates a new GetGridLicensePoolContainerResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGridLicensePoolContainerResponseObjectAsResult() *GetGridLicensePoolContainerResponseObjectAsResult {
	this := GetGridLicensePoolContainerResponseObjectAsResult{}
	return &this
}

// NewGetGridLicensePoolContainerResponseObjectAsResultWithDefaults instantiates a new GetGridLicensePoolContainerResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGridLicensePoolContainerResponseObjectAsResultWithDefaults() *GetGridLicensePoolContainerResponseObjectAsResult {
	this := GetGridLicensePoolContainerResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetGridLicensePoolContainerResponseObjectAsResult) GetResult() GridLicensePoolContainer {
	if o == nil || IsNil(o.Result) {
		var ret GridLicensePoolContainer
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGridLicensePoolContainerResponseObjectAsResult) GetResultOk() (*GridLicensePoolContainer, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetGridLicensePoolContainerResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given GridLicensePoolContainer and assigns it to the Result field.
func (o *GetGridLicensePoolContainerResponseObjectAsResult) SetResult(v GridLicensePoolContainer) {
	o.Result = &v
}

func (o GetGridLicensePoolContainerResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetGridLicensePoolContainerResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetGridLicensePoolContainerResponseObjectAsResult struct {
	value *GetGridLicensePoolContainerResponseObjectAsResult
	isSet bool
}

func (v NullableGetGridLicensePoolContainerResponseObjectAsResult) Get() *GetGridLicensePoolContainerResponseObjectAsResult {
	return v.value
}

func (v *NullableGetGridLicensePoolContainerResponseObjectAsResult) Set(val *GetGridLicensePoolContainerResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGridLicensePoolContainerResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGridLicensePoolContainerResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGridLicensePoolContainerResponseObjectAsResult(val *GetGridLicensePoolContainerResponseObjectAsResult) *NullableGetGridLicensePoolContainerResponseObjectAsResult {
	return &NullableGetGridLicensePoolContainerResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetGridLicensePoolContainerResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGridLicensePoolContainerResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
