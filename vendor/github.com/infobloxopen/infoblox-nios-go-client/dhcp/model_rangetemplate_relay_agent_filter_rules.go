/*
Infoblox DHCP API

OpenAPI specification for Infoblox NIOS WAPI DHCP objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dhcp

import (
	"encoding/json"
)

// checks if the RangetemplateRelayAgentFilterRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RangetemplateRelayAgentFilterRules{}

// RangetemplateRelayAgentFilterRules struct for RangetemplateRelayAgentFilterRules
type RangetemplateRelayAgentFilterRules struct {
	// The name of the DHCP filter.
	Filter *string `json:"filter,omitempty"`
	// The permission to be applied.
	Permission           *string `json:"permission,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RangetemplateRelayAgentFilterRules RangetemplateRelayAgentFilterRules

// NewRangetemplateRelayAgentFilterRules instantiates a new RangetemplateRelayAgentFilterRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRangetemplateRelayAgentFilterRules() *RangetemplateRelayAgentFilterRules {
	this := RangetemplateRelayAgentFilterRules{}
	return &this
}

// NewRangetemplateRelayAgentFilterRulesWithDefaults instantiates a new RangetemplateRelayAgentFilterRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangetemplateRelayAgentFilterRulesWithDefaults() *RangetemplateRelayAgentFilterRules {
	this := RangetemplateRelayAgentFilterRules{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *RangetemplateRelayAgentFilterRules) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangetemplateRelayAgentFilterRules) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *RangetemplateRelayAgentFilterRules) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *RangetemplateRelayAgentFilterRules) SetFilter(v string) {
	o.Filter = &v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *RangetemplateRelayAgentFilterRules) GetPermission() string {
	if o == nil || IsNil(o.Permission) {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangetemplateRelayAgentFilterRules) GetPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *RangetemplateRelayAgentFilterRules) HasPermission() bool {
	if o != nil && !IsNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *RangetemplateRelayAgentFilterRules) SetPermission(v string) {
	o.Permission = &v
}

func (o RangetemplateRelayAgentFilterRules) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RangetemplateRelayAgentFilterRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Permission) {
		toSerialize["permission"] = o.Permission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RangetemplateRelayAgentFilterRules) UnmarshalJSON(data []byte) (err error) {
	varRangetemplateRelayAgentFilterRules := _RangetemplateRelayAgentFilterRules{}

	err = json.Unmarshal(data, &varRangetemplateRelayAgentFilterRules)

	if err != nil {
		return err
	}

	*o = RangetemplateRelayAgentFilterRules(varRangetemplateRelayAgentFilterRules)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filter")
		delete(additionalProperties, "permission")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRangetemplateRelayAgentFilterRules struct {
	value *RangetemplateRelayAgentFilterRules
	isSet bool
}

func (v NullableRangetemplateRelayAgentFilterRules) Get() *RangetemplateRelayAgentFilterRules {
	return v.value
}

func (v *NullableRangetemplateRelayAgentFilterRules) Set(val *RangetemplateRelayAgentFilterRules) {
	v.value = val
	v.isSet = true
}

func (v NullableRangetemplateRelayAgentFilterRules) IsSet() bool {
	return v.isSet
}

func (v *NullableRangetemplateRelayAgentFilterRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRangetemplateRelayAgentFilterRules(val *RangetemplateRelayAgentFilterRules) *NullableRangetemplateRelayAgentFilterRules {
	return &NullableRangetemplateRelayAgentFilterRules{value: val, isSet: true}
}

func (v NullableRangetemplateRelayAgentFilterRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRangetemplateRelayAgentFilterRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
