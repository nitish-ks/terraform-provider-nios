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

// checks if the MemberDnsSortlist type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberDnsSortlist{}

// MemberDnsSortlist struct for MemberDnsSortlist
type MemberDnsSortlist struct {
	// The source address of a sortlist object.
	Address *string `json:"address,omitempty"`
	// The match list of a sortlist.
	MatchList            []string `json:"match_list,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemberDnsSortlist MemberDnsSortlist

// NewMemberDnsSortlist instantiates a new MemberDnsSortlist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberDnsSortlist() *MemberDnsSortlist {
	this := MemberDnsSortlist{}
	return &this
}

// NewMemberDnsSortlistWithDefaults instantiates a new MemberDnsSortlist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberDnsSortlistWithDefaults() *MemberDnsSortlist {
	this := MemberDnsSortlist{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *MemberDnsSortlist) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDnsSortlist) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *MemberDnsSortlist) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *MemberDnsSortlist) SetAddress(v string) {
	o.Address = &v
}

// GetMatchList returns the MatchList field value if set, zero value otherwise.
func (o *MemberDnsSortlist) GetMatchList() []string {
	if o == nil || IsNil(o.MatchList) {
		var ret []string
		return ret
	}
	return o.MatchList
}

// GetMatchListOk returns a tuple with the MatchList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDnsSortlist) GetMatchListOk() ([]string, bool) {
	if o == nil || IsNil(o.MatchList) {
		return nil, false
	}
	return o.MatchList, true
}

// HasMatchList returns a boolean if a field has been set.
func (o *MemberDnsSortlist) HasMatchList() bool {
	if o != nil && !IsNil(o.MatchList) {
		return true
	}

	return false
}

// SetMatchList gets a reference to the given []string and assigns it to the MatchList field.
func (o *MemberDnsSortlist) SetMatchList(v []string) {
	o.MatchList = v
}

func (o MemberDnsSortlist) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberDnsSortlist) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.MatchList) {
		toSerialize["match_list"] = o.MatchList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemberDnsSortlist) UnmarshalJSON(data []byte) (err error) {
	varMemberDnsSortlist := _MemberDnsSortlist{}

	err = json.Unmarshal(data, &varMemberDnsSortlist)

	if err != nil {
		return err
	}

	*o = MemberDnsSortlist(varMemberDnsSortlist)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "match_list")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemberDnsSortlist struct {
	value *MemberDnsSortlist
	isSet bool
}

func (v NullableMemberDnsSortlist) Get() *MemberDnsSortlist {
	return v.value
}

func (v *NullableMemberDnsSortlist) Set(val *MemberDnsSortlist) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberDnsSortlist) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberDnsSortlist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberDnsSortlist(val *MemberDnsSortlist) *NullableMemberDnsSortlist {
	return &NullableMemberDnsSortlist{value: val, isSet: true}
}

func (v NullableMemberDnsSortlist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberDnsSortlist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
