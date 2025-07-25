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

// checks if the Ipv6dhcpoptionspace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ipv6dhcpoptionspace{}

// Ipv6dhcpoptionspace struct for Ipv6dhcpoptionspace
type Ipv6dhcpoptionspace struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// A descriptive comment of a DHCP IPv6 option space object.
	Comment *string `json:"comment,omitempty"`
	// The enterprise number of a DHCP IPv6 option space object.
	EnterpriseNumber *int64 `json:"enterprise_number,omitempty"`
	// The name of a DHCP IPv6 option space object.
	Name *string `json:"name,omitempty"`
	// The list of DHCP IPv6 option definition objects.
	OptionDefinitions []string `json:"option_definitions,omitempty"`
}

// NewIpv6dhcpoptionspace instantiates a new Ipv6dhcpoptionspace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv6dhcpoptionspace() *Ipv6dhcpoptionspace {
	this := Ipv6dhcpoptionspace{}
	return &this
}

// NewIpv6dhcpoptionspaceWithDefaults instantiates a new Ipv6dhcpoptionspace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv6dhcpoptionspaceWithDefaults() *Ipv6dhcpoptionspace {
	this := Ipv6dhcpoptionspace{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *Ipv6dhcpoptionspace) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6dhcpoptionspace) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *Ipv6dhcpoptionspace) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *Ipv6dhcpoptionspace) SetRef(v string) {
	o.Ref = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *Ipv6dhcpoptionspace) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6dhcpoptionspace) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *Ipv6dhcpoptionspace) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *Ipv6dhcpoptionspace) SetComment(v string) {
	o.Comment = &v
}

// GetEnterpriseNumber returns the EnterpriseNumber field value if set, zero value otherwise.
func (o *Ipv6dhcpoptionspace) GetEnterpriseNumber() int64 {
	if o == nil || IsNil(o.EnterpriseNumber) {
		var ret int64
		return ret
	}
	return *o.EnterpriseNumber
}

// GetEnterpriseNumberOk returns a tuple with the EnterpriseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6dhcpoptionspace) GetEnterpriseNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.EnterpriseNumber) {
		return nil, false
	}
	return o.EnterpriseNumber, true
}

// HasEnterpriseNumber returns a boolean if a field has been set.
func (o *Ipv6dhcpoptionspace) HasEnterpriseNumber() bool {
	if o != nil && !IsNil(o.EnterpriseNumber) {
		return true
	}

	return false
}

// SetEnterpriseNumber gets a reference to the given int64 and assigns it to the EnterpriseNumber field.
func (o *Ipv6dhcpoptionspace) SetEnterpriseNumber(v int64) {
	o.EnterpriseNumber = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Ipv6dhcpoptionspace) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6dhcpoptionspace) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ipv6dhcpoptionspace) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ipv6dhcpoptionspace) SetName(v string) {
	o.Name = &v
}

// GetOptionDefinitions returns the OptionDefinitions field value if set, zero value otherwise.
func (o *Ipv6dhcpoptionspace) GetOptionDefinitions() []string {
	if o == nil || IsNil(o.OptionDefinitions) {
		var ret []string
		return ret
	}
	return o.OptionDefinitions
}

// GetOptionDefinitionsOk returns a tuple with the OptionDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv6dhcpoptionspace) GetOptionDefinitionsOk() ([]string, bool) {
	if o == nil || IsNil(o.OptionDefinitions) {
		return nil, false
	}
	return o.OptionDefinitions, true
}

// HasOptionDefinitions returns a boolean if a field has been set.
func (o *Ipv6dhcpoptionspace) HasOptionDefinitions() bool {
	if o != nil && !IsNil(o.OptionDefinitions) {
		return true
	}

	return false
}

// SetOptionDefinitions gets a reference to the given []string and assigns it to the OptionDefinitions field.
func (o *Ipv6dhcpoptionspace) SetOptionDefinitions(v []string) {
	o.OptionDefinitions = v
}

func (o Ipv6dhcpoptionspace) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ipv6dhcpoptionspace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.EnterpriseNumber) {
		toSerialize["enterprise_number"] = o.EnterpriseNumber
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OptionDefinitions) {
		toSerialize["option_definitions"] = o.OptionDefinitions
	}
	return toSerialize, nil
}

type NullableIpv6dhcpoptionspace struct {
	value *Ipv6dhcpoptionspace
	isSet bool
}

func (v NullableIpv6dhcpoptionspace) Get() *Ipv6dhcpoptionspace {
	return v.value
}

func (v *NullableIpv6dhcpoptionspace) Set(val *Ipv6dhcpoptionspace) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv6dhcpoptionspace) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv6dhcpoptionspace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv6dhcpoptionspace(val *Ipv6dhcpoptionspace) *NullableIpv6dhcpoptionspace {
	return &NullableIpv6dhcpoptionspace{value: val, isSet: true}
}

func (v NullableIpv6dhcpoptionspace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv6dhcpoptionspace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
