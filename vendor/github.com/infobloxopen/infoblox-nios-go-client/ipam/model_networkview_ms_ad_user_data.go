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

// checks if the NetworkviewMsAdUserData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkviewMsAdUserData{}

// NetworkviewMsAdUserData struct for NetworkviewMsAdUserData
type NetworkviewMsAdUserData struct {
	// The number of active users.
	ActiveUsersCount     *int64 `json:"active_users_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkviewMsAdUserData NetworkviewMsAdUserData

// NewNetworkviewMsAdUserData instantiates a new NetworkviewMsAdUserData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkviewMsAdUserData() *NetworkviewMsAdUserData {
	this := NetworkviewMsAdUserData{}
	return &this
}

// NewNetworkviewMsAdUserDataWithDefaults instantiates a new NetworkviewMsAdUserData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkviewMsAdUserDataWithDefaults() *NetworkviewMsAdUserData {
	this := NetworkviewMsAdUserData{}
	return &this
}

// GetActiveUsersCount returns the ActiveUsersCount field value if set, zero value otherwise.
func (o *NetworkviewMsAdUserData) GetActiveUsersCount() int64 {
	if o == nil || IsNil(o.ActiveUsersCount) {
		var ret int64
		return ret
	}
	return *o.ActiveUsersCount
}

// GetActiveUsersCountOk returns a tuple with the ActiveUsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkviewMsAdUserData) GetActiveUsersCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ActiveUsersCount) {
		return nil, false
	}
	return o.ActiveUsersCount, true
}

// HasActiveUsersCount returns a boolean if a field has been set.
func (o *NetworkviewMsAdUserData) HasActiveUsersCount() bool {
	if o != nil && !IsNil(o.ActiveUsersCount) {
		return true
	}

	return false
}

// SetActiveUsersCount gets a reference to the given int64 and assigns it to the ActiveUsersCount field.
func (o *NetworkviewMsAdUserData) SetActiveUsersCount(v int64) {
	o.ActiveUsersCount = &v
}

func (o NetworkviewMsAdUserData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkviewMsAdUserData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActiveUsersCount) {
		toSerialize["active_users_count"] = o.ActiveUsersCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkviewMsAdUserData) UnmarshalJSON(data []byte) (err error) {
	varNetworkviewMsAdUserData := _NetworkviewMsAdUserData{}

	err = json.Unmarshal(data, &varNetworkviewMsAdUserData)

	if err != nil {
		return err
	}

	*o = NetworkviewMsAdUserData(varNetworkviewMsAdUserData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active_users_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkviewMsAdUserData struct {
	value *NetworkviewMsAdUserData
	isSet bool
}

func (v NullableNetworkviewMsAdUserData) Get() *NetworkviewMsAdUserData {
	return v.value
}

func (v *NullableNetworkviewMsAdUserData) Set(val *NetworkviewMsAdUserData) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkviewMsAdUserData) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkviewMsAdUserData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkviewMsAdUserData(val *NetworkviewMsAdUserData) *NullableNetworkviewMsAdUserData {
	return &NullableNetworkviewMsAdUserData{value: val, isSet: true}
}

func (v NullableNetworkviewMsAdUserData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkviewMsAdUserData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
