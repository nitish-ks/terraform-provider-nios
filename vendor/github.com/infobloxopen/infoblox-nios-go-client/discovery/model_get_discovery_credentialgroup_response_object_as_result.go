/*
Infoblox DISCOVERY API

OpenAPI specification for Infoblox NIOS WAPI DISCOVERY objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package discovery

import (
	"encoding/json"
)

// checks if the GetDiscoveryCredentialgroupResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDiscoveryCredentialgroupResponseObjectAsResult{}

// GetDiscoveryCredentialgroupResponseObjectAsResult The response format to retrieve __DiscoveryCredentialgroup__ objects.
type GetDiscoveryCredentialgroupResponseObjectAsResult struct {
	Result *DiscoveryCredentialgroup `json:"result,omitempty"`
}

// NewGetDiscoveryCredentialgroupResponseObjectAsResult instantiates a new GetDiscoveryCredentialgroupResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDiscoveryCredentialgroupResponseObjectAsResult() *GetDiscoveryCredentialgroupResponseObjectAsResult {
	this := GetDiscoveryCredentialgroupResponseObjectAsResult{}
	return &this
}

// NewGetDiscoveryCredentialgroupResponseObjectAsResultWithDefaults instantiates a new GetDiscoveryCredentialgroupResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDiscoveryCredentialgroupResponseObjectAsResultWithDefaults() *GetDiscoveryCredentialgroupResponseObjectAsResult {
	this := GetDiscoveryCredentialgroupResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetDiscoveryCredentialgroupResponseObjectAsResult) GetResult() DiscoveryCredentialgroup {
	if o == nil || IsNil(o.Result) {
		var ret DiscoveryCredentialgroup
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDiscoveryCredentialgroupResponseObjectAsResult) GetResultOk() (*DiscoveryCredentialgroup, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetDiscoveryCredentialgroupResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given DiscoveryCredentialgroup and assigns it to the Result field.
func (o *GetDiscoveryCredentialgroupResponseObjectAsResult) SetResult(v DiscoveryCredentialgroup) {
	o.Result = &v
}

func (o GetDiscoveryCredentialgroupResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDiscoveryCredentialgroupResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetDiscoveryCredentialgroupResponseObjectAsResult struct {
	value *GetDiscoveryCredentialgroupResponseObjectAsResult
	isSet bool
}

func (v NullableGetDiscoveryCredentialgroupResponseObjectAsResult) Get() *GetDiscoveryCredentialgroupResponseObjectAsResult {
	return v.value
}

func (v *NullableGetDiscoveryCredentialgroupResponseObjectAsResult) Set(val *GetDiscoveryCredentialgroupResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDiscoveryCredentialgroupResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDiscoveryCredentialgroupResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDiscoveryCredentialgroupResponseObjectAsResult(val *GetDiscoveryCredentialgroupResponseObjectAsResult) *NullableGetDiscoveryCredentialgroupResponseObjectAsResult {
	return &NullableGetDiscoveryCredentialgroupResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetDiscoveryCredentialgroupResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDiscoveryCredentialgroupResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
