/*
Infoblox THREATINSIGHT API

OpenAPI specification for Infoblox NIOS WAPI THREATINSIGHT objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package threatinsight

import (
	"encoding/json"
)

// checks if the ThreatinsightCloudclient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreatinsightCloudclient{}

// ThreatinsightCloudclient struct for ThreatinsightCloudclient
type ThreatinsightCloudclient struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// The RPZs to which you apply newly detected domains through the Infoblox Threat Insight Cloud Client.
	BlacklistRpzList []string `json:"blacklist_rpz_list,omitempty"`
	// Determines whether the Threat Insight in Cloud Client is enabled.
	Enable *bool `json:"enable,omitempty"`
	// Force a refresh if at least one RPZ is configured.
	ForceRefresh *bool `json:"force_refresh,omitempty"`
	// The time interval (in seconds) for requesting newly detected domains by the Infoblox Threat Insight Cloud Client and applying them to the list of configured RPZs.
	Interval *int64 `json:"interval,omitempty"`
}

// NewThreatinsightCloudclient instantiates a new ThreatinsightCloudclient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreatinsightCloudclient() *ThreatinsightCloudclient {
	this := ThreatinsightCloudclient{}
	return &this
}

// NewThreatinsightCloudclientWithDefaults instantiates a new ThreatinsightCloudclient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreatinsightCloudclientWithDefaults() *ThreatinsightCloudclient {
	this := ThreatinsightCloudclient{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *ThreatinsightCloudclient) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatinsightCloudclient) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *ThreatinsightCloudclient) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *ThreatinsightCloudclient) SetRef(v string) {
	o.Ref = &v
}

// GetBlacklistRpzList returns the BlacklistRpzList field value if set, zero value otherwise.
func (o *ThreatinsightCloudclient) GetBlacklistRpzList() []string {
	if o == nil || IsNil(o.BlacklistRpzList) {
		var ret []string
		return ret
	}
	return o.BlacklistRpzList
}

// GetBlacklistRpzListOk returns a tuple with the BlacklistRpzList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatinsightCloudclient) GetBlacklistRpzListOk() ([]string, bool) {
	if o == nil || IsNil(o.BlacklistRpzList) {
		return nil, false
	}
	return o.BlacklistRpzList, true
}

// HasBlacklistRpzList returns a boolean if a field has been set.
func (o *ThreatinsightCloudclient) HasBlacklistRpzList() bool {
	if o != nil && !IsNil(o.BlacklistRpzList) {
		return true
	}

	return false
}

// SetBlacklistRpzList gets a reference to the given []string and assigns it to the BlacklistRpzList field.
func (o *ThreatinsightCloudclient) SetBlacklistRpzList(v []string) {
	o.BlacklistRpzList = v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *ThreatinsightCloudclient) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatinsightCloudclient) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *ThreatinsightCloudclient) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *ThreatinsightCloudclient) SetEnable(v bool) {
	o.Enable = &v
}

// GetForceRefresh returns the ForceRefresh field value if set, zero value otherwise.
func (o *ThreatinsightCloudclient) GetForceRefresh() bool {
	if o == nil || IsNil(o.ForceRefresh) {
		var ret bool
		return ret
	}
	return *o.ForceRefresh
}

// GetForceRefreshOk returns a tuple with the ForceRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatinsightCloudclient) GetForceRefreshOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceRefresh) {
		return nil, false
	}
	return o.ForceRefresh, true
}

// HasForceRefresh returns a boolean if a field has been set.
func (o *ThreatinsightCloudclient) HasForceRefresh() bool {
	if o != nil && !IsNil(o.ForceRefresh) {
		return true
	}

	return false
}

// SetForceRefresh gets a reference to the given bool and assigns it to the ForceRefresh field.
func (o *ThreatinsightCloudclient) SetForceRefresh(v bool) {
	o.ForceRefresh = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *ThreatinsightCloudclient) GetInterval() int64 {
	if o == nil || IsNil(o.Interval) {
		var ret int64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatinsightCloudclient) GetIntervalOk() (*int64, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *ThreatinsightCloudclient) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int64 and assigns it to the Interval field.
func (o *ThreatinsightCloudclient) SetInterval(v int64) {
	o.Interval = &v
}

func (o ThreatinsightCloudclient) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreatinsightCloudclient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.BlacklistRpzList) {
		toSerialize["blacklist_rpz_list"] = o.BlacklistRpzList
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.ForceRefresh) {
		toSerialize["force_refresh"] = o.ForceRefresh
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	return toSerialize, nil
}

type NullableThreatinsightCloudclient struct {
	value *ThreatinsightCloudclient
	isSet bool
}

func (v NullableThreatinsightCloudclient) Get() *ThreatinsightCloudclient {
	return v.value
}

func (v *NullableThreatinsightCloudclient) Set(val *ThreatinsightCloudclient) {
	v.value = val
	v.isSet = true
}

func (v NullableThreatinsightCloudclient) IsSet() bool {
	return v.isSet
}

func (v *NullableThreatinsightCloudclient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreatinsightCloudclient(val *ThreatinsightCloudclient) *NullableThreatinsightCloudclient {
	return &NullableThreatinsightCloudclient{value: val, isSet: true}
}

func (v NullableThreatinsightCloudclient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreatinsightCloudclient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
