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

// checks if the AdmingroupDockerShowCommands type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdmingroupDockerShowCommands{}

// AdmingroupDockerShowCommands struct for AdmingroupDockerShowCommands
type AdmingroupDockerShowCommands struct {
	// If True then CLI user has permission to run the command
	ShowDockerBridge *bool `json:"show_docker_bridge,omitempty"`
	// If True then enable all fields
	EnableAll *bool `json:"enable_all,omitempty"`
	// If True then disable all fields
	DisableAll           *bool `json:"disable_all,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdmingroupDockerShowCommands AdmingroupDockerShowCommands

// NewAdmingroupDockerShowCommands instantiates a new AdmingroupDockerShowCommands object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdmingroupDockerShowCommands() *AdmingroupDockerShowCommands {
	this := AdmingroupDockerShowCommands{}
	return &this
}

// NewAdmingroupDockerShowCommandsWithDefaults instantiates a new AdmingroupDockerShowCommands object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdmingroupDockerShowCommandsWithDefaults() *AdmingroupDockerShowCommands {
	this := AdmingroupDockerShowCommands{}
	return &this
}

// GetShowDockerBridge returns the ShowDockerBridge field value if set, zero value otherwise.
func (o *AdmingroupDockerShowCommands) GetShowDockerBridge() bool {
	if o == nil || IsNil(o.ShowDockerBridge) {
		var ret bool
		return ret
	}
	return *o.ShowDockerBridge
}

// GetShowDockerBridgeOk returns a tuple with the ShowDockerBridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdmingroupDockerShowCommands) GetShowDockerBridgeOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowDockerBridge) {
		return nil, false
	}
	return o.ShowDockerBridge, true
}

// HasShowDockerBridge returns a boolean if a field has been set.
func (o *AdmingroupDockerShowCommands) HasShowDockerBridge() bool {
	if o != nil && !IsNil(o.ShowDockerBridge) {
		return true
	}

	return false
}

// SetShowDockerBridge gets a reference to the given bool and assigns it to the ShowDockerBridge field.
func (o *AdmingroupDockerShowCommands) SetShowDockerBridge(v bool) {
	o.ShowDockerBridge = &v
}

// GetEnableAll returns the EnableAll field value if set, zero value otherwise.
func (o *AdmingroupDockerShowCommands) GetEnableAll() bool {
	if o == nil || IsNil(o.EnableAll) {
		var ret bool
		return ret
	}
	return *o.EnableAll
}

// GetEnableAllOk returns a tuple with the EnableAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdmingroupDockerShowCommands) GetEnableAllOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAll) {
		return nil, false
	}
	return o.EnableAll, true
}

// HasEnableAll returns a boolean if a field has been set.
func (o *AdmingroupDockerShowCommands) HasEnableAll() bool {
	if o != nil && !IsNil(o.EnableAll) {
		return true
	}

	return false
}

// SetEnableAll gets a reference to the given bool and assigns it to the EnableAll field.
func (o *AdmingroupDockerShowCommands) SetEnableAll(v bool) {
	o.EnableAll = &v
}

// GetDisableAll returns the DisableAll field value if set, zero value otherwise.
func (o *AdmingroupDockerShowCommands) GetDisableAll() bool {
	if o == nil || IsNil(o.DisableAll) {
		var ret bool
		return ret
	}
	return *o.DisableAll
}

// GetDisableAllOk returns a tuple with the DisableAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdmingroupDockerShowCommands) GetDisableAllOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableAll) {
		return nil, false
	}
	return o.DisableAll, true
}

// HasDisableAll returns a boolean if a field has been set.
func (o *AdmingroupDockerShowCommands) HasDisableAll() bool {
	if o != nil && !IsNil(o.DisableAll) {
		return true
	}

	return false
}

// SetDisableAll gets a reference to the given bool and assigns it to the DisableAll field.
func (o *AdmingroupDockerShowCommands) SetDisableAll(v bool) {
	o.DisableAll = &v
}

func (o AdmingroupDockerShowCommands) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdmingroupDockerShowCommands) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShowDockerBridge) {
		toSerialize["show_docker_bridge"] = o.ShowDockerBridge
	}
	if !IsNil(o.EnableAll) {
		toSerialize["enable_all"] = o.EnableAll
	}
	if !IsNil(o.DisableAll) {
		toSerialize["disable_all"] = o.DisableAll
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AdmingroupDockerShowCommands) UnmarshalJSON(data []byte) (err error) {
	varAdmingroupDockerShowCommands := _AdmingroupDockerShowCommands{}

	err = json.Unmarshal(data, &varAdmingroupDockerShowCommands)

	if err != nil {
		return err
	}

	*o = AdmingroupDockerShowCommands(varAdmingroupDockerShowCommands)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "show_docker_bridge")
		delete(additionalProperties, "enable_all")
		delete(additionalProperties, "disable_all")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdmingroupDockerShowCommands struct {
	value *AdmingroupDockerShowCommands
	isSet bool
}

func (v NullableAdmingroupDockerShowCommands) Get() *AdmingroupDockerShowCommands {
	return v.value
}

func (v *NullableAdmingroupDockerShowCommands) Set(val *AdmingroupDockerShowCommands) {
	v.value = val
	v.isSet = true
}

func (v NullableAdmingroupDockerShowCommands) IsSet() bool {
	return v.isSet
}

func (v *NullableAdmingroupDockerShowCommands) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdmingroupDockerShowCommands(val *AdmingroupDockerShowCommands) *NullableAdmingroupDockerShowCommands {
	return &NullableAdmingroupDockerShowCommands{value: val, isSet: true}
}

func (v NullableAdmingroupDockerShowCommands) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdmingroupDockerShowCommands) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
