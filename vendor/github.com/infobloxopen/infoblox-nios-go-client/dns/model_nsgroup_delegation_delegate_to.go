/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the NsgroupDelegationDelegateTo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NsgroupDelegationDelegateTo{}

// NsgroupDelegationDelegateTo struct for NsgroupDelegationDelegateTo
type NsgroupDelegationDelegateTo struct {
	// The IPv4 Address or IPv6 Address of the server.
	Address *string `json:"address,omitempty"`
	// A resolvable domain name for the external DNS server.
	Name *string `json:"name,omitempty"`
	// This flag represents whether the name server is shared with the parent Microsoft primary zone's delegation server.
	SharedWithMsParentDelegation *bool `json:"shared_with_ms_parent_delegation,omitempty"`
	// Set this flag to hide the NS record for the primary name server from DNS queries.
	Stealth *bool `json:"stealth,omitempty"`
	// A generated TSIG key.
	TsigKey *string `json:"tsig_key,omitempty"`
	// The TSIG key algorithm.
	TsigKeyAlg *string `json:"tsig_key_alg,omitempty"`
	// The TSIG key name.
	TsigKeyName *string `json:"tsig_key_name,omitempty"`
	// Use flag for: tsig_key_name
	UseTsigKeyName       *bool `json:"use_tsig_key_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NsgroupDelegationDelegateTo NsgroupDelegationDelegateTo

// NewNsgroupDelegationDelegateTo instantiates a new NsgroupDelegationDelegateTo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsgroupDelegationDelegateTo() *NsgroupDelegationDelegateTo {
	this := NsgroupDelegationDelegateTo{}
	return &this
}

// NewNsgroupDelegationDelegateToWithDefaults instantiates a new NsgroupDelegationDelegateTo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsgroupDelegationDelegateToWithDefaults() *NsgroupDelegationDelegateTo {
	this := NsgroupDelegationDelegateTo{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *NsgroupDelegationDelegateTo) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsgroupDelegationDelegateTo) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *NsgroupDelegationDelegateTo) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *NsgroupDelegationDelegateTo) SetAddress(v string) {
	o.Address = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NsgroupDelegationDelegateTo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsgroupDelegationDelegateTo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NsgroupDelegationDelegateTo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NsgroupDelegationDelegateTo) SetName(v string) {
	o.Name = &v
}

// GetSharedWithMsParentDelegation returns the SharedWithMsParentDelegation field value if set, zero value otherwise.
func (o *NsgroupDelegationDelegateTo) GetSharedWithMsParentDelegation() bool {
	if o == nil || IsNil(o.SharedWithMsParentDelegation) {
		var ret bool
		return ret
	}
	return *o.SharedWithMsParentDelegation
}

// GetSharedWithMsParentDelegationOk returns a tuple with the SharedWithMsParentDelegation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsgroupDelegationDelegateTo) GetSharedWithMsParentDelegationOk() (*bool, bool) {
	if o == nil || IsNil(o.SharedWithMsParentDelegation) {
		return nil, false
	}
	return o.SharedWithMsParentDelegation, true
}

// HasSharedWithMsParentDelegation returns a boolean if a field has been set.
func (o *NsgroupDelegationDelegateTo) HasSharedWithMsParentDelegation() bool {
	if o != nil && !IsNil(o.SharedWithMsParentDelegation) {
		return true
	}

	return false
}

// SetSharedWithMsParentDelegation gets a reference to the given bool and assigns it to the SharedWithMsParentDelegation field.
func (o *NsgroupDelegationDelegateTo) SetSharedWithMsParentDelegation(v bool) {
	o.SharedWithMsParentDelegation = &v
}

// GetStealth returns the Stealth field value if set, zero value otherwise.
func (o *NsgroupDelegationDelegateTo) GetStealth() bool {
	if o == nil || IsNil(o.Stealth) {
		var ret bool
		return ret
	}
	return *o.Stealth
}

// GetStealthOk returns a tuple with the Stealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsgroupDelegationDelegateTo) GetStealthOk() (*bool, bool) {
	if o == nil || IsNil(o.Stealth) {
		return nil, false
	}
	return o.Stealth, true
}

// HasStealth returns a boolean if a field has been set.
func (o *NsgroupDelegationDelegateTo) HasStealth() bool {
	if o != nil && !IsNil(o.Stealth) {
		return true
	}

	return false
}

// SetStealth gets a reference to the given bool and assigns it to the Stealth field.
func (o *NsgroupDelegationDelegateTo) SetStealth(v bool) {
	o.Stealth = &v
}

// GetTsigKey returns the TsigKey field value if set, zero value otherwise.
func (o *NsgroupDelegationDelegateTo) GetTsigKey() string {
	if o == nil || IsNil(o.TsigKey) {
		var ret string
		return ret
	}
	return *o.TsigKey
}

// GetTsigKeyOk returns a tuple with the TsigKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsgroupDelegationDelegateTo) GetTsigKeyOk() (*string, bool) {
	if o == nil || IsNil(o.TsigKey) {
		return nil, false
	}
	return o.TsigKey, true
}

// HasTsigKey returns a boolean if a field has been set.
func (o *NsgroupDelegationDelegateTo) HasTsigKey() bool {
	if o != nil && !IsNil(o.TsigKey) {
		return true
	}

	return false
}

// SetTsigKey gets a reference to the given string and assigns it to the TsigKey field.
func (o *NsgroupDelegationDelegateTo) SetTsigKey(v string) {
	o.TsigKey = &v
}

// GetTsigKeyAlg returns the TsigKeyAlg field value if set, zero value otherwise.
func (o *NsgroupDelegationDelegateTo) GetTsigKeyAlg() string {
	if o == nil || IsNil(o.TsigKeyAlg) {
		var ret string
		return ret
	}
	return *o.TsigKeyAlg
}

// GetTsigKeyAlgOk returns a tuple with the TsigKeyAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsgroupDelegationDelegateTo) GetTsigKeyAlgOk() (*string, bool) {
	if o == nil || IsNil(o.TsigKeyAlg) {
		return nil, false
	}
	return o.TsigKeyAlg, true
}

// HasTsigKeyAlg returns a boolean if a field has been set.
func (o *NsgroupDelegationDelegateTo) HasTsigKeyAlg() bool {
	if o != nil && !IsNil(o.TsigKeyAlg) {
		return true
	}

	return false
}

// SetTsigKeyAlg gets a reference to the given string and assigns it to the TsigKeyAlg field.
func (o *NsgroupDelegationDelegateTo) SetTsigKeyAlg(v string) {
	o.TsigKeyAlg = &v
}

// GetTsigKeyName returns the TsigKeyName field value if set, zero value otherwise.
func (o *NsgroupDelegationDelegateTo) GetTsigKeyName() string {
	if o == nil || IsNil(o.TsigKeyName) {
		var ret string
		return ret
	}
	return *o.TsigKeyName
}

// GetTsigKeyNameOk returns a tuple with the TsigKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsgroupDelegationDelegateTo) GetTsigKeyNameOk() (*string, bool) {
	if o == nil || IsNil(o.TsigKeyName) {
		return nil, false
	}
	return o.TsigKeyName, true
}

// HasTsigKeyName returns a boolean if a field has been set.
func (o *NsgroupDelegationDelegateTo) HasTsigKeyName() bool {
	if o != nil && !IsNil(o.TsigKeyName) {
		return true
	}

	return false
}

// SetTsigKeyName gets a reference to the given string and assigns it to the TsigKeyName field.
func (o *NsgroupDelegationDelegateTo) SetTsigKeyName(v string) {
	o.TsigKeyName = &v
}

// GetUseTsigKeyName returns the UseTsigKeyName field value if set, zero value otherwise.
func (o *NsgroupDelegationDelegateTo) GetUseTsigKeyName() bool {
	if o == nil || IsNil(o.UseTsigKeyName) {
		var ret bool
		return ret
	}
	return *o.UseTsigKeyName
}

// GetUseTsigKeyNameOk returns a tuple with the UseTsigKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsgroupDelegationDelegateTo) GetUseTsigKeyNameOk() (*bool, bool) {
	if o == nil || IsNil(o.UseTsigKeyName) {
		return nil, false
	}
	return o.UseTsigKeyName, true
}

// HasUseTsigKeyName returns a boolean if a field has been set.
func (o *NsgroupDelegationDelegateTo) HasUseTsigKeyName() bool {
	if o != nil && !IsNil(o.UseTsigKeyName) {
		return true
	}

	return false
}

// SetUseTsigKeyName gets a reference to the given bool and assigns it to the UseTsigKeyName field.
func (o *NsgroupDelegationDelegateTo) SetUseTsigKeyName(v bool) {
	o.UseTsigKeyName = &v
}

func (o NsgroupDelegationDelegateTo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NsgroupDelegationDelegateTo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SharedWithMsParentDelegation) {
		toSerialize["shared_with_ms_parent_delegation"] = o.SharedWithMsParentDelegation
	}
	if !IsNil(o.Stealth) {
		toSerialize["stealth"] = o.Stealth
	}
	if !IsNil(o.TsigKey) {
		toSerialize["tsig_key"] = o.TsigKey
	}
	if !IsNil(o.TsigKeyAlg) {
		toSerialize["tsig_key_alg"] = o.TsigKeyAlg
	}
	if !IsNil(o.TsigKeyName) {
		toSerialize["tsig_key_name"] = o.TsigKeyName
	}
	if !IsNil(o.UseTsigKeyName) {
		toSerialize["use_tsig_key_name"] = o.UseTsigKeyName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NsgroupDelegationDelegateTo) UnmarshalJSON(data []byte) (err error) {
	varNsgroupDelegationDelegateTo := _NsgroupDelegationDelegateTo{}

	err = json.Unmarshal(data, &varNsgroupDelegationDelegateTo)

	if err != nil {
		return err
	}

	*o = NsgroupDelegationDelegateTo(varNsgroupDelegationDelegateTo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		delete(additionalProperties, "name")
		delete(additionalProperties, "shared_with_ms_parent_delegation")
		delete(additionalProperties, "stealth")
		delete(additionalProperties, "tsig_key")
		delete(additionalProperties, "tsig_key_alg")
		delete(additionalProperties, "tsig_key_name")
		delete(additionalProperties, "use_tsig_key_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNsgroupDelegationDelegateTo struct {
	value *NsgroupDelegationDelegateTo
	isSet bool
}

func (v NullableNsgroupDelegationDelegateTo) Get() *NsgroupDelegationDelegateTo {
	return v.value
}

func (v *NullableNsgroupDelegationDelegateTo) Set(val *NsgroupDelegationDelegateTo) {
	v.value = val
	v.isSet = true
}

func (v NullableNsgroupDelegationDelegateTo) IsSet() bool {
	return v.isSet
}

func (v *NullableNsgroupDelegationDelegateTo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsgroupDelegationDelegateTo(val *NsgroupDelegationDelegateTo) *NullableNsgroupDelegationDelegateTo {
	return &NullableNsgroupDelegationDelegateTo{value: val, isSet: true}
}

func (v NullableNsgroupDelegationDelegateTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsgroupDelegationDelegateTo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
