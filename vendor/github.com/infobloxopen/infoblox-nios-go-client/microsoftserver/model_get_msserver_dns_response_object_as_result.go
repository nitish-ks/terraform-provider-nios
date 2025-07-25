/*
Infoblox MICROSOFTSERVER API

OpenAPI specification for Infoblox NIOS WAPI MICROSOFTSERVER objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package microsoftserver

import (
	"encoding/json"
)

// checks if the GetMsserverDnsResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMsserverDnsResponseObjectAsResult{}

// GetMsserverDnsResponseObjectAsResult The response format to retrieve __MsserverDns__ objects.
type GetMsserverDnsResponseObjectAsResult struct {
	Result *MsserverDns `json:"result,omitempty"`
}

// NewGetMsserverDnsResponseObjectAsResult instantiates a new GetMsserverDnsResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMsserverDnsResponseObjectAsResult() *GetMsserverDnsResponseObjectAsResult {
	this := GetMsserverDnsResponseObjectAsResult{}
	return &this
}

// NewGetMsserverDnsResponseObjectAsResultWithDefaults instantiates a new GetMsserverDnsResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMsserverDnsResponseObjectAsResultWithDefaults() *GetMsserverDnsResponseObjectAsResult {
	this := GetMsserverDnsResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetMsserverDnsResponseObjectAsResult) GetResult() MsserverDns {
	if o == nil || IsNil(o.Result) {
		var ret MsserverDns
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMsserverDnsResponseObjectAsResult) GetResultOk() (*MsserverDns, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetMsserverDnsResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given MsserverDns and assigns it to the Result field.
func (o *GetMsserverDnsResponseObjectAsResult) SetResult(v MsserverDns) {
	o.Result = &v
}

func (o GetMsserverDnsResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMsserverDnsResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetMsserverDnsResponseObjectAsResult struct {
	value *GetMsserverDnsResponseObjectAsResult
	isSet bool
}

func (v NullableGetMsserverDnsResponseObjectAsResult) Get() *GetMsserverDnsResponseObjectAsResult {
	return v.value
}

func (v *NullableGetMsserverDnsResponseObjectAsResult) Set(val *GetMsserverDnsResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMsserverDnsResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMsserverDnsResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMsserverDnsResponseObjectAsResult(val *GetMsserverDnsResponseObjectAsResult) *NullableGetMsserverDnsResponseObjectAsResult {
	return &NullableGetMsserverDnsResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetMsserverDnsResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMsserverDnsResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
