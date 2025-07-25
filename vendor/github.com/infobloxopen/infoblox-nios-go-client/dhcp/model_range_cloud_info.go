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

// checks if the RangeCloudInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RangeCloudInfo{}

// RangeCloudInfo struct for RangeCloudInfo
type RangeCloudInfo struct {
	DelegatedMember *RangecloudinfoDelegatedMember `json:"delegated_member,omitempty"`
	// Indicates the scope of delegation for the object. This can be one of the following: NONE (outside any delegation), ROOT (the delegation point), SUBTREE (within the scope of a delegation), RECLAIMING (within the scope of a delegation being reclaimed, either as the delegation point or in the subtree).
	DelegatedScope *string `json:"delegated_scope,omitempty"`
	// Indicates the root of the delegation if delegated_scope is SUBTREE or RECLAIMING. This is not set otherwise.
	DelegatedRoot *string `json:"delegated_root,omitempty"`
	// Determines whether the object was created by the cloud adapter or not.
	OwnedByAdaptor *bool `json:"owned_by_adaptor,omitempty"`
	// Indicates the cloud origin of the object.
	Usage *string `json:"usage,omitempty"`
	// Reference to the tenant object associated with the object, if any.
	Tenant *string `json:"tenant,omitempty"`
	// Indicates the specified cloud management platform.
	MgmtPlatform *string `json:"mgmt_platform,omitempty"`
	// Type of authority over the object.
	AuthorityType        *string `json:"authority_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RangeCloudInfo RangeCloudInfo

// NewRangeCloudInfo instantiates a new RangeCloudInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRangeCloudInfo() *RangeCloudInfo {
	this := RangeCloudInfo{}
	return &this
}

// NewRangeCloudInfoWithDefaults instantiates a new RangeCloudInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRangeCloudInfoWithDefaults() *RangeCloudInfo {
	this := RangeCloudInfo{}
	return &this
}

// GetDelegatedMember returns the DelegatedMember field value if set, zero value otherwise.
func (o *RangeCloudInfo) GetDelegatedMember() RangecloudinfoDelegatedMember {
	if o == nil || IsNil(o.DelegatedMember) {
		var ret RangecloudinfoDelegatedMember
		return ret
	}
	return *o.DelegatedMember
}

// GetDelegatedMemberOk returns a tuple with the DelegatedMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeCloudInfo) GetDelegatedMemberOk() (*RangecloudinfoDelegatedMember, bool) {
	if o == nil || IsNil(o.DelegatedMember) {
		return nil, false
	}
	return o.DelegatedMember, true
}

// HasDelegatedMember returns a boolean if a field has been set.
func (o *RangeCloudInfo) HasDelegatedMember() bool {
	if o != nil && !IsNil(o.DelegatedMember) {
		return true
	}

	return false
}

// SetDelegatedMember gets a reference to the given RangecloudinfoDelegatedMember and assigns it to the DelegatedMember field.
func (o *RangeCloudInfo) SetDelegatedMember(v RangecloudinfoDelegatedMember) {
	o.DelegatedMember = &v
}

// GetDelegatedScope returns the DelegatedScope field value if set, zero value otherwise.
func (o *RangeCloudInfo) GetDelegatedScope() string {
	if o == nil || IsNil(o.DelegatedScope) {
		var ret string
		return ret
	}
	return *o.DelegatedScope
}

// GetDelegatedScopeOk returns a tuple with the DelegatedScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeCloudInfo) GetDelegatedScopeOk() (*string, bool) {
	if o == nil || IsNil(o.DelegatedScope) {
		return nil, false
	}
	return o.DelegatedScope, true
}

// HasDelegatedScope returns a boolean if a field has been set.
func (o *RangeCloudInfo) HasDelegatedScope() bool {
	if o != nil && !IsNil(o.DelegatedScope) {
		return true
	}

	return false
}

// SetDelegatedScope gets a reference to the given string and assigns it to the DelegatedScope field.
func (o *RangeCloudInfo) SetDelegatedScope(v string) {
	o.DelegatedScope = &v
}

// GetDelegatedRoot returns the DelegatedRoot field value if set, zero value otherwise.
func (o *RangeCloudInfo) GetDelegatedRoot() string {
	if o == nil || IsNil(o.DelegatedRoot) {
		var ret string
		return ret
	}
	return *o.DelegatedRoot
}

// GetDelegatedRootOk returns a tuple with the DelegatedRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeCloudInfo) GetDelegatedRootOk() (*string, bool) {
	if o == nil || IsNil(o.DelegatedRoot) {
		return nil, false
	}
	return o.DelegatedRoot, true
}

// HasDelegatedRoot returns a boolean if a field has been set.
func (o *RangeCloudInfo) HasDelegatedRoot() bool {
	if o != nil && !IsNil(o.DelegatedRoot) {
		return true
	}

	return false
}

// SetDelegatedRoot gets a reference to the given string and assigns it to the DelegatedRoot field.
func (o *RangeCloudInfo) SetDelegatedRoot(v string) {
	o.DelegatedRoot = &v
}

// GetOwnedByAdaptor returns the OwnedByAdaptor field value if set, zero value otherwise.
func (o *RangeCloudInfo) GetOwnedByAdaptor() bool {
	if o == nil || IsNil(o.OwnedByAdaptor) {
		var ret bool
		return ret
	}
	return *o.OwnedByAdaptor
}

// GetOwnedByAdaptorOk returns a tuple with the OwnedByAdaptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeCloudInfo) GetOwnedByAdaptorOk() (*bool, bool) {
	if o == nil || IsNil(o.OwnedByAdaptor) {
		return nil, false
	}
	return o.OwnedByAdaptor, true
}

// HasOwnedByAdaptor returns a boolean if a field has been set.
func (o *RangeCloudInfo) HasOwnedByAdaptor() bool {
	if o != nil && !IsNil(o.OwnedByAdaptor) {
		return true
	}

	return false
}

// SetOwnedByAdaptor gets a reference to the given bool and assigns it to the OwnedByAdaptor field.
func (o *RangeCloudInfo) SetOwnedByAdaptor(v bool) {
	o.OwnedByAdaptor = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *RangeCloudInfo) GetUsage() string {
	if o == nil || IsNil(o.Usage) {
		var ret string
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeCloudInfo) GetUsageOk() (*string, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *RangeCloudInfo) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given string and assigns it to the Usage field.
func (o *RangeCloudInfo) SetUsage(v string) {
	o.Usage = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *RangeCloudInfo) GetTenant() string {
	if o == nil || IsNil(o.Tenant) {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeCloudInfo) GetTenantOk() (*string, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *RangeCloudInfo) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *RangeCloudInfo) SetTenant(v string) {
	o.Tenant = &v
}

// GetMgmtPlatform returns the MgmtPlatform field value if set, zero value otherwise.
func (o *RangeCloudInfo) GetMgmtPlatform() string {
	if o == nil || IsNil(o.MgmtPlatform) {
		var ret string
		return ret
	}
	return *o.MgmtPlatform
}

// GetMgmtPlatformOk returns a tuple with the MgmtPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeCloudInfo) GetMgmtPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.MgmtPlatform) {
		return nil, false
	}
	return o.MgmtPlatform, true
}

// HasMgmtPlatform returns a boolean if a field has been set.
func (o *RangeCloudInfo) HasMgmtPlatform() bool {
	if o != nil && !IsNil(o.MgmtPlatform) {
		return true
	}

	return false
}

// SetMgmtPlatform gets a reference to the given string and assigns it to the MgmtPlatform field.
func (o *RangeCloudInfo) SetMgmtPlatform(v string) {
	o.MgmtPlatform = &v
}

// GetAuthorityType returns the AuthorityType field value if set, zero value otherwise.
func (o *RangeCloudInfo) GetAuthorityType() string {
	if o == nil || IsNil(o.AuthorityType) {
		var ret string
		return ret
	}
	return *o.AuthorityType
}

// GetAuthorityTypeOk returns a tuple with the AuthorityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RangeCloudInfo) GetAuthorityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorityType) {
		return nil, false
	}
	return o.AuthorityType, true
}

// HasAuthorityType returns a boolean if a field has been set.
func (o *RangeCloudInfo) HasAuthorityType() bool {
	if o != nil && !IsNil(o.AuthorityType) {
		return true
	}

	return false
}

// SetAuthorityType gets a reference to the given string and assigns it to the AuthorityType field.
func (o *RangeCloudInfo) SetAuthorityType(v string) {
	o.AuthorityType = &v
}

func (o RangeCloudInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RangeCloudInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DelegatedMember) {
		toSerialize["delegated_member"] = o.DelegatedMember
	}
	if !IsNil(o.DelegatedScope) {
		toSerialize["delegated_scope"] = o.DelegatedScope
	}
	if !IsNil(o.DelegatedRoot) {
		toSerialize["delegated_root"] = o.DelegatedRoot
	}
	if !IsNil(o.OwnedByAdaptor) {
		toSerialize["owned_by_adaptor"] = o.OwnedByAdaptor
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	if !IsNil(o.MgmtPlatform) {
		toSerialize["mgmt_platform"] = o.MgmtPlatform
	}
	if !IsNil(o.AuthorityType) {
		toSerialize["authority_type"] = o.AuthorityType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RangeCloudInfo) UnmarshalJSON(data []byte) (err error) {
	varRangeCloudInfo := _RangeCloudInfo{}

	err = json.Unmarshal(data, &varRangeCloudInfo)

	if err != nil {
		return err
	}

	*o = RangeCloudInfo(varRangeCloudInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "delegated_member")
		delete(additionalProperties, "delegated_scope")
		delete(additionalProperties, "delegated_root")
		delete(additionalProperties, "owned_by_adaptor")
		delete(additionalProperties, "usage")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "mgmt_platform")
		delete(additionalProperties, "authority_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRangeCloudInfo struct {
	value *RangeCloudInfo
	isSet bool
}

func (v NullableRangeCloudInfo) Get() *RangeCloudInfo {
	return v.value
}

func (v *NullableRangeCloudInfo) Set(val *RangeCloudInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRangeCloudInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRangeCloudInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRangeCloudInfo(val *RangeCloudInfo) *NullableRangeCloudInfo {
	return &NullableRangeCloudInfo{value: val, isSet: true}
}

func (v NullableRangeCloudInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRangeCloudInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
