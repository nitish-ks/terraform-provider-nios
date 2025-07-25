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

// checks if the GetSnmpuserResponseObjectAsResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSnmpuserResponseObjectAsResult{}

// GetSnmpuserResponseObjectAsResult The response format to retrieve __Snmpuser__ objects.
type GetSnmpuserResponseObjectAsResult struct {
	Result *Snmpuser `json:"result,omitempty"`
}

// NewGetSnmpuserResponseObjectAsResult instantiates a new GetSnmpuserResponseObjectAsResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSnmpuserResponseObjectAsResult() *GetSnmpuserResponseObjectAsResult {
	this := GetSnmpuserResponseObjectAsResult{}
	return &this
}

// NewGetSnmpuserResponseObjectAsResultWithDefaults instantiates a new GetSnmpuserResponseObjectAsResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSnmpuserResponseObjectAsResultWithDefaults() *GetSnmpuserResponseObjectAsResult {
	this := GetSnmpuserResponseObjectAsResult{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *GetSnmpuserResponseObjectAsResult) GetResult() Snmpuser {
	if o == nil || IsNil(o.Result) {
		var ret Snmpuser
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSnmpuserResponseObjectAsResult) GetResultOk() (*Snmpuser, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *GetSnmpuserResponseObjectAsResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given Snmpuser and assigns it to the Result field.
func (o *GetSnmpuserResponseObjectAsResult) SetResult(v Snmpuser) {
	o.Result = &v
}

func (o GetSnmpuserResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSnmpuserResponseObjectAsResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableGetSnmpuserResponseObjectAsResult struct {
	value *GetSnmpuserResponseObjectAsResult
	isSet bool
}

func (v NullableGetSnmpuserResponseObjectAsResult) Get() *GetSnmpuserResponseObjectAsResult {
	return v.value
}

func (v *NullableGetSnmpuserResponseObjectAsResult) Set(val *GetSnmpuserResponseObjectAsResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSnmpuserResponseObjectAsResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSnmpuserResponseObjectAsResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSnmpuserResponseObjectAsResult(val *GetSnmpuserResponseObjectAsResult) *NullableGetSnmpuserResponseObjectAsResult {
	return &NullableGetSnmpuserResponseObjectAsResult{value: val, isSet: true}
}

func (v NullableGetSnmpuserResponseObjectAsResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSnmpuserResponseObjectAsResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
