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

// checks if the GetLdapAuthServiceResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLdapAuthServiceResponseObjectAsResult{}

// GetLdapAuthServiceResponseObjectAsResult The response format to retrieve __LdapAuthService__ objects.
type GetLdapAuthServiceResponseObjectAsResult struct {
	Result *LdapAuthService `json:"result,omitempty"`
}

// NewGetLdapAuthServiceResponseObjectAsResult instantiates a new GetLdapAuthServiceResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLdapAuthServiceResponseObjectAsResult() *GetLdapAuthServiceResponseObjectAsResult {
	this := GetLdapAuthServiceResponseObjectAsResult{}
	return &this
}

// NewGetLdapAuthServiceResponseObjectAsResultWithDefaults instantiates a new GetLdapAuthServiceResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLdapAuthServiceResponseObjectAsResultWithDefaults() *GetLdapAuthServiceResponseObjectAsResult {
	this := GetLdapAuthServiceResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetLdapAuthServiceResponseObjectAsResult) GetResult() LdapAuthService {
	if o == nil || IsNil(o.Result) {
		var ret LdapAuthService
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLdapAuthServiceResponseObjectAsResult) GetResultOk() (*LdapAuthService, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetLdapAuthServiceResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given LdapAuthService and assigns it to the Result field.
func (o *GetLdapAuthServiceResponseObjectAsResult) SetResult(v LdapAuthService) {
	o.Result = &v
}

func (o GetLdapAuthServiceResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLdapAuthServiceResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetLdapAuthServiceResponseObjectAsResult struct {
	value *GetLdapAuthServiceResponseObjectAsResult
	isSet bool
}

func (v NullableGetLdapAuthServiceResponseObjectAsResult) Get() *GetLdapAuthServiceResponseObjectAsResult {
	return v.value
}

func (v *NullableGetLdapAuthServiceResponseObjectAsResult) Set(val *GetLdapAuthServiceResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLdapAuthServiceResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLdapAuthServiceResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLdapAuthServiceResponseObjectAsResult(val *GetLdapAuthServiceResponseObjectAsResult) *NullableGetLdapAuthServiceResponseObjectAsResult {
	return &NullableGetLdapAuthServiceResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetLdapAuthServiceResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLdapAuthServiceResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
