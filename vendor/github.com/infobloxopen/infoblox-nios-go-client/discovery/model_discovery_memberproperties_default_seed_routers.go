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

// checks if the DiscoveryMemberpropertiesDefaultSeedRouters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscoveryMemberpropertiesDefaultSeedRouters{}

// DiscoveryMemberpropertiesDefaultSeedRouters struct for DiscoveryMemberpropertiesDefaultSeedRouters
type DiscoveryMemberpropertiesDefaultSeedRouters struct {
	// Address of the seed router.
	Address *string `json:"address,omitempty"`
	// The network view name.
	NetworkView *string `json:"network_view,omitempty"`
	// Description of the seed router.
	Comment              *string `json:"comment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiscoveryMemberpropertiesDefaultSeedRouters DiscoveryMemberpropertiesDefaultSeedRouters

// NewDiscoveryMemberpropertiesDefaultSeedRouters instantiates a new DiscoveryMemberpropertiesDefaultSeedRouters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveryMemberpropertiesDefaultSeedRouters() *DiscoveryMemberpropertiesDefaultSeedRouters {
	this := DiscoveryMemberpropertiesDefaultSeedRouters{}
	return &this
}

// NewDiscoveryMemberpropertiesDefaultSeedRoutersWithDefaults instantiates a new DiscoveryMemberpropertiesDefaultSeedRouters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveryMemberpropertiesDefaultSeedRoutersWithDefaults() *DiscoveryMemberpropertiesDefaultSeedRouters {
	this := DiscoveryMemberpropertiesDefaultSeedRouters{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *DiscoveryMemberpropertiesDefaultSeedRouters) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryMemberpropertiesDefaultSeedRouters) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *DiscoveryMemberpropertiesDefaultSeedRouters) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *DiscoveryMemberpropertiesDefaultSeedRouters) SetAddress(v string) {
	o.Address = &v
}

// GetNetworkView returns the NetworkView field value if set, zero value otherwise.
func (o *DiscoveryMemberpropertiesDefaultSeedRouters) GetNetworkView() string {
	if o == nil || IsNil(o.NetworkView) {
		var ret string
		return ret
	}
	return *o.NetworkView
}

// GetNetworkViewOk returns a tuple with the NetworkView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryMemberpropertiesDefaultSeedRouters) GetNetworkViewOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkView) {
		return nil, false
	}
	return o.NetworkView, true
}

// HasNetworkView returns a boolean if a field has been set.
func (o *DiscoveryMemberpropertiesDefaultSeedRouters) HasNetworkView() bool {
	if o != nil && !IsNil(o.NetworkView) {
		return true
	}

	return false
}

// SetNetworkView gets a reference to the given string and assigns it to the NetworkView field.
func (o *DiscoveryMemberpropertiesDefaultSeedRouters) SetNetworkView(v string) {
	o.NetworkView = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *DiscoveryMemberpropertiesDefaultSeedRouters) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryMemberpropertiesDefaultSeedRouters) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *DiscoveryMemberpropertiesDefaultSeedRouters) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *DiscoveryMemberpropertiesDefaultSeedRouters) SetComment(v string) {
	o.Comment = &v
}

func (o DiscoveryMemberpropertiesDefaultSeedRouters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveryMemberpropertiesDefaultSeedRouters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.NetworkView) {
		toSerialize["network_view"] = o.NetworkView
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DiscoveryMemberpropertiesDefaultSeedRouters) UnmarshalJSON(data []byte) (err error) {
	varDiscoveryMemberpropertiesDefaultSeedRouters := _DiscoveryMemberpropertiesDefaultSeedRouters{}

	err = json.Unmarshal(data, &varDiscoveryMemberpropertiesDefaultSeedRouters)

	if err != nil {
		return err
	}

	*o = DiscoveryMemberpropertiesDefaultSeedRouters(varDiscoveryMemberpropertiesDefaultSeedRouters)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "network_view")
		delete(additionalProperties, "comment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiscoveryMemberpropertiesDefaultSeedRouters struct {
	value *DiscoveryMemberpropertiesDefaultSeedRouters
	isSet bool
}

func (v NullableDiscoveryMemberpropertiesDefaultSeedRouters) Get() *DiscoveryMemberpropertiesDefaultSeedRouters {
	return v.value
}

func (v *NullableDiscoveryMemberpropertiesDefaultSeedRouters) Set(val *DiscoveryMemberpropertiesDefaultSeedRouters) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryMemberpropertiesDefaultSeedRouters) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryMemberpropertiesDefaultSeedRouters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryMemberpropertiesDefaultSeedRouters(val *DiscoveryMemberpropertiesDefaultSeedRouters) *NullableDiscoveryMemberpropertiesDefaultSeedRouters {
	return &NullableDiscoveryMemberpropertiesDefaultSeedRouters{value: val, isSet: true}
}

func (v NullableDiscoveryMemberpropertiesDefaultSeedRouters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryMemberpropertiesDefaultSeedRouters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
