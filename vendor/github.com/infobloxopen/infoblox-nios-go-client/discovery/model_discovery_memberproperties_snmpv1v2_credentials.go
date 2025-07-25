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

// checks if the DiscoveryMemberpropertiesSnmpv1v2Credentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscoveryMemberpropertiesSnmpv1v2Credentials{}

// DiscoveryMemberpropertiesSnmpv1v2Credentials struct for DiscoveryMemberpropertiesSnmpv1v2Credentials
type DiscoveryMemberpropertiesSnmpv1v2Credentials struct {
	// The public community string.
	CommunityString *string `json:"community_string,omitempty"`
	// Comments for the SNMPv1 and SNMPv2 users.
	Comment *string `json:"comment,omitempty"`
	// Group for the SNMPv1 and SNMPv2 credential.
	CredentialGroup      *string `json:"credential_group,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiscoveryMemberpropertiesSnmpv1v2Credentials DiscoveryMemberpropertiesSnmpv1v2Credentials

// NewDiscoveryMemberpropertiesSnmpv1v2Credentials instantiates a new DiscoveryMemberpropertiesSnmpv1v2Credentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveryMemberpropertiesSnmpv1v2Credentials() *DiscoveryMemberpropertiesSnmpv1v2Credentials {
	this := DiscoveryMemberpropertiesSnmpv1v2Credentials{}
	return &this
}

// NewDiscoveryMemberpropertiesSnmpv1v2CredentialsWithDefaults instantiates a new DiscoveryMemberpropertiesSnmpv1v2Credentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveryMemberpropertiesSnmpv1v2CredentialsWithDefaults() *DiscoveryMemberpropertiesSnmpv1v2Credentials {
	this := DiscoveryMemberpropertiesSnmpv1v2Credentials{}
	return &this
}

// GetCommunityString returns the CommunityString field value if set, zero value otherwise.
func (o *DiscoveryMemberpropertiesSnmpv1v2Credentials) GetCommunityString() string {
	if o == nil || IsNil(o.CommunityString) {
		var ret string
		return ret
	}
	return *o.CommunityString
}

// GetCommunityStringOk returns a tuple with the CommunityString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryMemberpropertiesSnmpv1v2Credentials) GetCommunityStringOk() (*string, bool) {
	if o == nil || IsNil(o.CommunityString) {
		return nil, false
	}
	return o.CommunityString, true
}

// HasCommunityString returns a boolean if a field has been set.
func (o *DiscoveryMemberpropertiesSnmpv1v2Credentials) HasCommunityString() bool {
	if o != nil && !IsNil(o.CommunityString) {
		return true
	}

	return false
}

// SetCommunityString gets a reference to the given string and assigns it to the CommunityString field.
func (o *DiscoveryMemberpropertiesSnmpv1v2Credentials) SetCommunityString(v string) {
	o.CommunityString = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *DiscoveryMemberpropertiesSnmpv1v2Credentials) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryMemberpropertiesSnmpv1v2Credentials) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *DiscoveryMemberpropertiesSnmpv1v2Credentials) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *DiscoveryMemberpropertiesSnmpv1v2Credentials) SetComment(v string) {
	o.Comment = &v
}

// GetCredentialGroup returns the CredentialGroup field value if set, zero value otherwise.
func (o *DiscoveryMemberpropertiesSnmpv1v2Credentials) GetCredentialGroup() string {
	if o == nil || IsNil(o.CredentialGroup) {
		var ret string
		return ret
	}
	return *o.CredentialGroup
}

// GetCredentialGroupOk returns a tuple with the CredentialGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryMemberpropertiesSnmpv1v2Credentials) GetCredentialGroupOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialGroup) {
		return nil, false
	}
	return o.CredentialGroup, true
}

// HasCredentialGroup returns a boolean if a field has been set.
func (o *DiscoveryMemberpropertiesSnmpv1v2Credentials) HasCredentialGroup() bool {
	if o != nil && !IsNil(o.CredentialGroup) {
		return true
	}

	return false
}

// SetCredentialGroup gets a reference to the given string and assigns it to the CredentialGroup field.
func (o *DiscoveryMemberpropertiesSnmpv1v2Credentials) SetCredentialGroup(v string) {
	o.CredentialGroup = &v
}

func (o DiscoveryMemberpropertiesSnmpv1v2Credentials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveryMemberpropertiesSnmpv1v2Credentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommunityString) {
		toSerialize["community_string"] = o.CommunityString
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CredentialGroup) {
		toSerialize["credential_group"] = o.CredentialGroup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DiscoveryMemberpropertiesSnmpv1v2Credentials) UnmarshalJSON(data []byte) (err error) {
	varDiscoveryMemberpropertiesSnmpv1v2Credentials := _DiscoveryMemberpropertiesSnmpv1v2Credentials{}

	err = json.Unmarshal(data, &varDiscoveryMemberpropertiesSnmpv1v2Credentials)

	if err != nil {
		return err
	}

	*o = DiscoveryMemberpropertiesSnmpv1v2Credentials(varDiscoveryMemberpropertiesSnmpv1v2Credentials)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "community_string")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "credential_group")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiscoveryMemberpropertiesSnmpv1v2Credentials struct {
	value *DiscoveryMemberpropertiesSnmpv1v2Credentials
	isSet bool
}

func (v NullableDiscoveryMemberpropertiesSnmpv1v2Credentials) Get() *DiscoveryMemberpropertiesSnmpv1v2Credentials {
	return v.value
}

func (v *NullableDiscoveryMemberpropertiesSnmpv1v2Credentials) Set(val *DiscoveryMemberpropertiesSnmpv1v2Credentials) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryMemberpropertiesSnmpv1v2Credentials) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryMemberpropertiesSnmpv1v2Credentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryMemberpropertiesSnmpv1v2Credentials(val *DiscoveryMemberpropertiesSnmpv1v2Credentials) *NullableDiscoveryMemberpropertiesSnmpv1v2Credentials {
	return &NullableDiscoveryMemberpropertiesSnmpv1v2Credentials{value: val, isSet: true}
}

func (v NullableDiscoveryMemberpropertiesSnmpv1v2Credentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryMemberpropertiesSnmpv1v2Credentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
