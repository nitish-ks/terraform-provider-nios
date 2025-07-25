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

// checks if the MemberIpv6StaticRoutes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberIpv6StaticRoutes{}

// MemberIpv6StaticRoutes struct for MemberIpv6StaticRoutes
type MemberIpv6StaticRoutes struct {
	// IPv6 address.
	Address *string `json:"address,omitempty"`
	// IPv6 CIDR
	Cidr *int64 `json:"cidr,omitempty"`
	// Gateway address.
	Gateway              *string `json:"gateway,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemberIpv6StaticRoutes MemberIpv6StaticRoutes

// NewMemberIpv6StaticRoutes instantiates a new MemberIpv6StaticRoutes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberIpv6StaticRoutes() *MemberIpv6StaticRoutes {
	this := MemberIpv6StaticRoutes{}
	return &this
}

// NewMemberIpv6StaticRoutesWithDefaults instantiates a new MemberIpv6StaticRoutes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberIpv6StaticRoutesWithDefaults() *MemberIpv6StaticRoutes {
	this := MemberIpv6StaticRoutes{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *MemberIpv6StaticRoutes) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberIpv6StaticRoutes) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *MemberIpv6StaticRoutes) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *MemberIpv6StaticRoutes) SetAddress(v string) {
	o.Address = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *MemberIpv6StaticRoutes) GetCidr() int64 {
	if o == nil || IsNil(o.Cidr) {
		var ret int64
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberIpv6StaticRoutes) GetCidrOk() (*int64, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *MemberIpv6StaticRoutes) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given int64 and assigns it to the Cidr field.
func (o *MemberIpv6StaticRoutes) SetCidr(v int64) {
	o.Cidr = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *MemberIpv6StaticRoutes) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberIpv6StaticRoutes) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *MemberIpv6StaticRoutes) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *MemberIpv6StaticRoutes) SetGateway(v string) {
	o.Gateway = &v
}

func (o MemberIpv6StaticRoutes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberIpv6StaticRoutes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemberIpv6StaticRoutes) UnmarshalJSON(data []byte) (err error) {
	varMemberIpv6StaticRoutes := _MemberIpv6StaticRoutes{}

	err = json.Unmarshal(data, &varMemberIpv6StaticRoutes)

	if err != nil {
		return err
	}

	*o = MemberIpv6StaticRoutes(varMemberIpv6StaticRoutes)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "cidr")
		delete(additionalProperties, "gateway")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemberIpv6StaticRoutes struct {
	value *MemberIpv6StaticRoutes
	isSet bool
}

func (v NullableMemberIpv6StaticRoutes) Get() *MemberIpv6StaticRoutes {
	return v.value
}

func (v *NullableMemberIpv6StaticRoutes) Set(val *MemberIpv6StaticRoutes) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberIpv6StaticRoutes) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberIpv6StaticRoutes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberIpv6StaticRoutes(val *MemberIpv6StaticRoutes) *NullableMemberIpv6StaticRoutes {
	return &NullableMemberIpv6StaticRoutes{value: val, isSet: true}
}

func (v NullableMemberIpv6StaticRoutes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberIpv6StaticRoutes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
