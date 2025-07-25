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

// checks if the NetworkviewAssociatedMembers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkviewAssociatedMembers{}

// NetworkviewAssociatedMembers struct for NetworkviewAssociatedMembers
type NetworkviewAssociatedMembers struct {
	// The member object associated with a network view.
	Member *string `json:"member,omitempty"`
	// The list of failover objects associated with each member.
	Failovers            []string `json:"failovers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkviewAssociatedMembers NetworkviewAssociatedMembers

// NewNetworkviewAssociatedMembers instantiates a new NetworkviewAssociatedMembers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkviewAssociatedMembers() *NetworkviewAssociatedMembers {
	this := NetworkviewAssociatedMembers{}
	return &this
}

// NewNetworkviewAssociatedMembersWithDefaults instantiates a new NetworkviewAssociatedMembers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkviewAssociatedMembersWithDefaults() *NetworkviewAssociatedMembers {
	this := NetworkviewAssociatedMembers{}
	return &this
}

// GetMember returns the Member field value if set, zero value otherwise.
func (o *NetworkviewAssociatedMembers) GetMember() string {
	if o == nil || IsNil(o.Member) {
		var ret string
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkviewAssociatedMembers) GetMemberOk() (*string, bool) {
	if o == nil || IsNil(o.Member) {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *NetworkviewAssociatedMembers) HasMember() bool {
	if o != nil && !IsNil(o.Member) {
		return true
	}

	return false
}

// SetMember gets a reference to the given string and assigns it to the Member field.
func (o *NetworkviewAssociatedMembers) SetMember(v string) {
	o.Member = &v
}

// GetFailovers returns the Failovers field value if set, zero value otherwise.
func (o *NetworkviewAssociatedMembers) GetFailovers() []string {
	if o == nil || IsNil(o.Failovers) {
		var ret []string
		return ret
	}
	return o.Failovers
}

// GetFailoversOk returns a tuple with the Failovers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkviewAssociatedMembers) GetFailoversOk() ([]string, bool) {
	if o == nil || IsNil(o.Failovers) {
		return nil, false
	}
	return o.Failovers, true
}

// HasFailovers returns a boolean if a field has been set.
func (o *NetworkviewAssociatedMembers) HasFailovers() bool {
	if o != nil && !IsNil(o.Failovers) {
		return true
	}

	return false
}

// SetFailovers gets a reference to the given []string and assigns it to the Failovers field.
func (o *NetworkviewAssociatedMembers) SetFailovers(v []string) {
	o.Failovers = v
}

func (o NetworkviewAssociatedMembers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkviewAssociatedMembers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Member) {
		toSerialize["member"] = o.Member
	}
	if !IsNil(o.Failovers) {
		toSerialize["failovers"] = o.Failovers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkviewAssociatedMembers) UnmarshalJSON(data []byte) (err error) {
	varNetworkviewAssociatedMembers := _NetworkviewAssociatedMembers{}

	err = json.Unmarshal(data, &varNetworkviewAssociatedMembers)

	if err != nil {
		return err
	}

	*o = NetworkviewAssociatedMembers(varNetworkviewAssociatedMembers)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "member")
		delete(additionalProperties, "failovers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkviewAssociatedMembers struct {
	value *NetworkviewAssociatedMembers
	isSet bool
}

func (v NullableNetworkviewAssociatedMembers) Get() *NetworkviewAssociatedMembers {
	return v.value
}

func (v *NullableNetworkviewAssociatedMembers) Set(val *NetworkviewAssociatedMembers) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkviewAssociatedMembers) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkviewAssociatedMembers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkviewAssociatedMembers(val *NetworkviewAssociatedMembers) *NullableNetworkviewAssociatedMembers {
	return &NullableNetworkviewAssociatedMembers{value: val, isSet: true}
}

func (v NullableNetworkviewAssociatedMembers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkviewAssociatedMembers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
