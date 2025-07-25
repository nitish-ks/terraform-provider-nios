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

// checks if the MemberthreatprotectionnatrulesNatPorts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberthreatprotectionnatrulesNatPorts{}

// MemberthreatprotectionnatrulesNatPorts struct for MemberthreatprotectionnatrulesNatPorts
type MemberthreatprotectionnatrulesNatPorts struct {
	// The start port value for the NAT port configuration object.
	StartPort *int64 `json:"start_port,omitempty"`
	// The end port value for the NAT port configuration object.
	EndPort *int64 `json:"end_port,omitempty"`
	// The block size for the NAT Port configuration object.
	BlockSize            *int64 `json:"block_size,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemberthreatprotectionnatrulesNatPorts MemberthreatprotectionnatrulesNatPorts

// NewMemberthreatprotectionnatrulesNatPorts instantiates a new MemberthreatprotectionnatrulesNatPorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberthreatprotectionnatrulesNatPorts() *MemberthreatprotectionnatrulesNatPorts {
	this := MemberthreatprotectionnatrulesNatPorts{}
	return &this
}

// NewMemberthreatprotectionnatrulesNatPortsWithDefaults instantiates a new MemberthreatprotectionnatrulesNatPorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberthreatprotectionnatrulesNatPortsWithDefaults() *MemberthreatprotectionnatrulesNatPorts {
	this := MemberthreatprotectionnatrulesNatPorts{}
	return &this
}

// GetStartPort returns the StartPort field value if set, zero value otherwise.
func (o *MemberthreatprotectionnatrulesNatPorts) GetStartPort() int64 {
	if o == nil || IsNil(o.StartPort) {
		var ret int64
		return ret
	}
	return *o.StartPort
}

// GetStartPortOk returns a tuple with the StartPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberthreatprotectionnatrulesNatPorts) GetStartPortOk() (*int64, bool) {
	if o == nil || IsNil(o.StartPort) {
		return nil, false
	}
	return o.StartPort, true
}

// HasStartPort returns a boolean if a field has been set.
func (o *MemberthreatprotectionnatrulesNatPorts) HasStartPort() bool {
	if o != nil && !IsNil(o.StartPort) {
		return true
	}

	return false
}

// SetStartPort gets a reference to the given int64 and assigns it to the StartPort field.
func (o *MemberthreatprotectionnatrulesNatPorts) SetStartPort(v int64) {
	o.StartPort = &v
}

// GetEndPort returns the EndPort field value if set, zero value otherwise.
func (o *MemberthreatprotectionnatrulesNatPorts) GetEndPort() int64 {
	if o == nil || IsNil(o.EndPort) {
		var ret int64
		return ret
	}
	return *o.EndPort
}

// GetEndPortOk returns a tuple with the EndPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberthreatprotectionnatrulesNatPorts) GetEndPortOk() (*int64, bool) {
	if o == nil || IsNil(o.EndPort) {
		return nil, false
	}
	return o.EndPort, true
}

// HasEndPort returns a boolean if a field has been set.
func (o *MemberthreatprotectionnatrulesNatPorts) HasEndPort() bool {
	if o != nil && !IsNil(o.EndPort) {
		return true
	}

	return false
}

// SetEndPort gets a reference to the given int64 and assigns it to the EndPort field.
func (o *MemberthreatprotectionnatrulesNatPorts) SetEndPort(v int64) {
	o.EndPort = &v
}

// GetBlockSize returns the BlockSize field value if set, zero value otherwise.
func (o *MemberthreatprotectionnatrulesNatPorts) GetBlockSize() int64 {
	if o == nil || IsNil(o.BlockSize) {
		var ret int64
		return ret
	}
	return *o.BlockSize
}

// GetBlockSizeOk returns a tuple with the BlockSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberthreatprotectionnatrulesNatPorts) GetBlockSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.BlockSize) {
		return nil, false
	}
	return o.BlockSize, true
}

// HasBlockSize returns a boolean if a field has been set.
func (o *MemberthreatprotectionnatrulesNatPorts) HasBlockSize() bool {
	if o != nil && !IsNil(o.BlockSize) {
		return true
	}

	return false
}

// SetBlockSize gets a reference to the given int64 and assigns it to the BlockSize field.
func (o *MemberthreatprotectionnatrulesNatPorts) SetBlockSize(v int64) {
	o.BlockSize = &v
}

func (o MemberthreatprotectionnatrulesNatPorts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberthreatprotectionnatrulesNatPorts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartPort) {
		toSerialize["start_port"] = o.StartPort
	}
	if !IsNil(o.EndPort) {
		toSerialize["end_port"] = o.EndPort
	}
	if !IsNil(o.BlockSize) {
		toSerialize["block_size"] = o.BlockSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemberthreatprotectionnatrulesNatPorts) UnmarshalJSON(data []byte) (err error) {
	varMemberthreatprotectionnatrulesNatPorts := _MemberthreatprotectionnatrulesNatPorts{}

	err = json.Unmarshal(data, &varMemberthreatprotectionnatrulesNatPorts)

	if err != nil {
		return err
	}

	*o = MemberthreatprotectionnatrulesNatPorts(varMemberthreatprotectionnatrulesNatPorts)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "start_port")
		delete(additionalProperties, "end_port")
		delete(additionalProperties, "block_size")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemberthreatprotectionnatrulesNatPorts struct {
	value *MemberthreatprotectionnatrulesNatPorts
	isSet bool
}

func (v NullableMemberthreatprotectionnatrulesNatPorts) Get() *MemberthreatprotectionnatrulesNatPorts {
	return v.value
}

func (v *NullableMemberthreatprotectionnatrulesNatPorts) Set(val *MemberthreatprotectionnatrulesNatPorts) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberthreatprotectionnatrulesNatPorts) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberthreatprotectionnatrulesNatPorts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberthreatprotectionnatrulesNatPorts(val *MemberthreatprotectionnatrulesNatPorts) *NullableMemberthreatprotectionnatrulesNatPorts {
	return &NullableMemberthreatprotectionnatrulesNatPorts{value: val, isSet: true}
}

func (v NullableMemberthreatprotectionnatrulesNatPorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberthreatprotectionnatrulesNatPorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
