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

// checks if the MemberDnsAutoBlackhole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberDnsAutoBlackhole{}

// MemberDnsAutoBlackhole struct for MemberDnsAutoBlackhole
type MemberDnsAutoBlackhole struct {
	// Enables or disables the configuration of the maximum number of concurrent recursive queries the appliance sends to each upstream DNS server.
	EnableFetchesPerServer *bool `json:"enable_fetches_per_server,omitempty"`
	// Enables or disables the configuration of the maximum number of concurrent recursive queries the appliance sends to each DNS zone.
	EnableFetchesPerZone *bool `json:"enable_fetches_per_zone,omitempty"`
	// Enables or disables the holddown configuration when the appliance stops sending queries to non-responsive servers.
	EnableHolddown *bool `json:"enable_holddown,omitempty"`
	// The maximum number of concurrent recursive queries the appliance sends to a single upstream name server before blocking additional queries to that server.
	FetchesPerServer *int64 `json:"fetches_per_server,omitempty"`
	// The maximum number of concurrent recursive queries that a server sends for its domains.
	FetchesPerZone *int64 `json:"fetches_per_zone,omitempty"`
	// Determines how often (in number of recursive responses) the appliance recalculates the average timeout ratio for each DNS server.
	FpsFreq *int64 `json:"fps_freq,omitempty"`
	// The holddown duration for non-responsive servers.
	Holddown *int64 `json:"holddown,omitempty"`
	// The number of consecutive timeouts before holding down a non-responsive server.
	HolddownThreshold *int64 `json:"holddown_threshold,omitempty"`
	// The minimum time (in seconds) that needs to be passed before a timeout occurs. Note that only these timeouts are counted towards the number of consecutive timeouts.
	HolddownTimeout      *int64 `json:"holddown_timeout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemberDnsAutoBlackhole MemberDnsAutoBlackhole

// NewMemberDnsAutoBlackhole instantiates a new MemberDnsAutoBlackhole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberDnsAutoBlackhole() *MemberDnsAutoBlackhole {
	this := MemberDnsAutoBlackhole{}
	return &this
}

// NewMemberDnsAutoBlackholeWithDefaults instantiates a new MemberDnsAutoBlackhole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberDnsAutoBlackholeWithDefaults() *MemberDnsAutoBlackhole {
	this := MemberDnsAutoBlackhole{}
	return &this
}

// GetEnableFetchesPerServer returns the EnableFetchesPerServer field value if set, zero value otherwise.
func (o *MemberDnsAutoBlackhole) GetEnableFetchesPerServer() bool {
	if o == nil || IsNil(o.EnableFetchesPerServer) {
		var ret bool
		return ret
	}
	return *o.EnableFetchesPerServer
}

// GetEnableFetchesPerServerOk returns a tuple with the EnableFetchesPerServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDnsAutoBlackhole) GetEnableFetchesPerServerOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableFetchesPerServer) {
		return nil, false
	}
	return o.EnableFetchesPerServer, true
}

// HasEnableFetchesPerServer returns a boolean if a field has been set.
func (o *MemberDnsAutoBlackhole) HasEnableFetchesPerServer() bool {
	if o != nil && !IsNil(o.EnableFetchesPerServer) {
		return true
	}

	return false
}

// SetEnableFetchesPerServer gets a reference to the given bool and assigns it to the EnableFetchesPerServer field.
func (o *MemberDnsAutoBlackhole) SetEnableFetchesPerServer(v bool) {
	o.EnableFetchesPerServer = &v
}

// GetEnableFetchesPerZone returns the EnableFetchesPerZone field value if set, zero value otherwise.
func (o *MemberDnsAutoBlackhole) GetEnableFetchesPerZone() bool {
	if o == nil || IsNil(o.EnableFetchesPerZone) {
		var ret bool
		return ret
	}
	return *o.EnableFetchesPerZone
}

// GetEnableFetchesPerZoneOk returns a tuple with the EnableFetchesPerZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDnsAutoBlackhole) GetEnableFetchesPerZoneOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableFetchesPerZone) {
		return nil, false
	}
	return o.EnableFetchesPerZone, true
}

// HasEnableFetchesPerZone returns a boolean if a field has been set.
func (o *MemberDnsAutoBlackhole) HasEnableFetchesPerZone() bool {
	if o != nil && !IsNil(o.EnableFetchesPerZone) {
		return true
	}

	return false
}

// SetEnableFetchesPerZone gets a reference to the given bool and assigns it to the EnableFetchesPerZone field.
func (o *MemberDnsAutoBlackhole) SetEnableFetchesPerZone(v bool) {
	o.EnableFetchesPerZone = &v
}

// GetEnableHolddown returns the EnableHolddown field value if set, zero value otherwise.
func (o *MemberDnsAutoBlackhole) GetEnableHolddown() bool {
	if o == nil || IsNil(o.EnableHolddown) {
		var ret bool
		return ret
	}
	return *o.EnableHolddown
}

// GetEnableHolddownOk returns a tuple with the EnableHolddown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDnsAutoBlackhole) GetEnableHolddownOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableHolddown) {
		return nil, false
	}
	return o.EnableHolddown, true
}

// HasEnableHolddown returns a boolean if a field has been set.
func (o *MemberDnsAutoBlackhole) HasEnableHolddown() bool {
	if o != nil && !IsNil(o.EnableHolddown) {
		return true
	}

	return false
}

// SetEnableHolddown gets a reference to the given bool and assigns it to the EnableHolddown field.
func (o *MemberDnsAutoBlackhole) SetEnableHolddown(v bool) {
	o.EnableHolddown = &v
}

// GetFetchesPerServer returns the FetchesPerServer field value if set, zero value otherwise.
func (o *MemberDnsAutoBlackhole) GetFetchesPerServer() int64 {
	if o == nil || IsNil(o.FetchesPerServer) {
		var ret int64
		return ret
	}
	return *o.FetchesPerServer
}

// GetFetchesPerServerOk returns a tuple with the FetchesPerServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDnsAutoBlackhole) GetFetchesPerServerOk() (*int64, bool) {
	if o == nil || IsNil(o.FetchesPerServer) {
		return nil, false
	}
	return o.FetchesPerServer, true
}

// HasFetchesPerServer returns a boolean if a field has been set.
func (o *MemberDnsAutoBlackhole) HasFetchesPerServer() bool {
	if o != nil && !IsNil(o.FetchesPerServer) {
		return true
	}

	return false
}

// SetFetchesPerServer gets a reference to the given int64 and assigns it to the FetchesPerServer field.
func (o *MemberDnsAutoBlackhole) SetFetchesPerServer(v int64) {
	o.FetchesPerServer = &v
}

// GetFetchesPerZone returns the FetchesPerZone field value if set, zero value otherwise.
func (o *MemberDnsAutoBlackhole) GetFetchesPerZone() int64 {
	if o == nil || IsNil(o.FetchesPerZone) {
		var ret int64
		return ret
	}
	return *o.FetchesPerZone
}

// GetFetchesPerZoneOk returns a tuple with the FetchesPerZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDnsAutoBlackhole) GetFetchesPerZoneOk() (*int64, bool) {
	if o == nil || IsNil(o.FetchesPerZone) {
		return nil, false
	}
	return o.FetchesPerZone, true
}

// HasFetchesPerZone returns a boolean if a field has been set.
func (o *MemberDnsAutoBlackhole) HasFetchesPerZone() bool {
	if o != nil && !IsNil(o.FetchesPerZone) {
		return true
	}

	return false
}

// SetFetchesPerZone gets a reference to the given int64 and assigns it to the FetchesPerZone field.
func (o *MemberDnsAutoBlackhole) SetFetchesPerZone(v int64) {
	o.FetchesPerZone = &v
}

// GetFpsFreq returns the FpsFreq field value if set, zero value otherwise.
func (o *MemberDnsAutoBlackhole) GetFpsFreq() int64 {
	if o == nil || IsNil(o.FpsFreq) {
		var ret int64
		return ret
	}
	return *o.FpsFreq
}

// GetFpsFreqOk returns a tuple with the FpsFreq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDnsAutoBlackhole) GetFpsFreqOk() (*int64, bool) {
	if o == nil || IsNil(o.FpsFreq) {
		return nil, false
	}
	return o.FpsFreq, true
}

// HasFpsFreq returns a boolean if a field has been set.
func (o *MemberDnsAutoBlackhole) HasFpsFreq() bool {
	if o != nil && !IsNil(o.FpsFreq) {
		return true
	}

	return false
}

// SetFpsFreq gets a reference to the given int64 and assigns it to the FpsFreq field.
func (o *MemberDnsAutoBlackhole) SetFpsFreq(v int64) {
	o.FpsFreq = &v
}

// GetHolddown returns the Holddown field value if set, zero value otherwise.
func (o *MemberDnsAutoBlackhole) GetHolddown() int64 {
	if o == nil || IsNil(o.Holddown) {
		var ret int64
		return ret
	}
	return *o.Holddown
}

// GetHolddownOk returns a tuple with the Holddown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDnsAutoBlackhole) GetHolddownOk() (*int64, bool) {
	if o == nil || IsNil(o.Holddown) {
		return nil, false
	}
	return o.Holddown, true
}

// HasHolddown returns a boolean if a field has been set.
func (o *MemberDnsAutoBlackhole) HasHolddown() bool {
	if o != nil && !IsNil(o.Holddown) {
		return true
	}

	return false
}

// SetHolddown gets a reference to the given int64 and assigns it to the Holddown field.
func (o *MemberDnsAutoBlackhole) SetHolddown(v int64) {
	o.Holddown = &v
}

// GetHolddownThreshold returns the HolddownThreshold field value if set, zero value otherwise.
func (o *MemberDnsAutoBlackhole) GetHolddownThreshold() int64 {
	if o == nil || IsNil(o.HolddownThreshold) {
		var ret int64
		return ret
	}
	return *o.HolddownThreshold
}

// GetHolddownThresholdOk returns a tuple with the HolddownThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDnsAutoBlackhole) GetHolddownThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.HolddownThreshold) {
		return nil, false
	}
	return o.HolddownThreshold, true
}

// HasHolddownThreshold returns a boolean if a field has been set.
func (o *MemberDnsAutoBlackhole) HasHolddownThreshold() bool {
	if o != nil && !IsNil(o.HolddownThreshold) {
		return true
	}

	return false
}

// SetHolddownThreshold gets a reference to the given int64 and assigns it to the HolddownThreshold field.
func (o *MemberDnsAutoBlackhole) SetHolddownThreshold(v int64) {
	o.HolddownThreshold = &v
}

// GetHolddownTimeout returns the HolddownTimeout field value if set, zero value otherwise.
func (o *MemberDnsAutoBlackhole) GetHolddownTimeout() int64 {
	if o == nil || IsNil(o.HolddownTimeout) {
		var ret int64
		return ret
	}
	return *o.HolddownTimeout
}

// GetHolddownTimeoutOk returns a tuple with the HolddownTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDnsAutoBlackhole) GetHolddownTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.HolddownTimeout) {
		return nil, false
	}
	return o.HolddownTimeout, true
}

// HasHolddownTimeout returns a boolean if a field has been set.
func (o *MemberDnsAutoBlackhole) HasHolddownTimeout() bool {
	if o != nil && !IsNil(o.HolddownTimeout) {
		return true
	}

	return false
}

// SetHolddownTimeout gets a reference to the given int64 and assigns it to the HolddownTimeout field.
func (o *MemberDnsAutoBlackhole) SetHolddownTimeout(v int64) {
	o.HolddownTimeout = &v
}

func (o MemberDnsAutoBlackhole) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberDnsAutoBlackhole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableFetchesPerServer) {
		toSerialize["enable_fetches_per_server"] = o.EnableFetchesPerServer
	}
	if !IsNil(o.EnableFetchesPerZone) {
		toSerialize["enable_fetches_per_zone"] = o.EnableFetchesPerZone
	}
	if !IsNil(o.EnableHolddown) {
		toSerialize["enable_holddown"] = o.EnableHolddown
	}
	if !IsNil(o.FetchesPerServer) {
		toSerialize["fetches_per_server"] = o.FetchesPerServer
	}
	if !IsNil(o.FetchesPerZone) {
		toSerialize["fetches_per_zone"] = o.FetchesPerZone
	}
	if !IsNil(o.FpsFreq) {
		toSerialize["fps_freq"] = o.FpsFreq
	}
	if !IsNil(o.Holddown) {
		toSerialize["holddown"] = o.Holddown
	}
	if !IsNil(o.HolddownThreshold) {
		toSerialize["holddown_threshold"] = o.HolddownThreshold
	}
	if !IsNil(o.HolddownTimeout) {
		toSerialize["holddown_timeout"] = o.HolddownTimeout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemberDnsAutoBlackhole) UnmarshalJSON(data []byte) (err error) {
	varMemberDnsAutoBlackhole := _MemberDnsAutoBlackhole{}

	err = json.Unmarshal(data, &varMemberDnsAutoBlackhole)

	if err != nil {
		return err
	}

	*o = MemberDnsAutoBlackhole(varMemberDnsAutoBlackhole)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enable_fetches_per_server")
		delete(additionalProperties, "enable_fetches_per_zone")
		delete(additionalProperties, "enable_holddown")
		delete(additionalProperties, "fetches_per_server")
		delete(additionalProperties, "fetches_per_zone")
		delete(additionalProperties, "fps_freq")
		delete(additionalProperties, "holddown")
		delete(additionalProperties, "holddown_threshold")
		delete(additionalProperties, "holddown_timeout")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemberDnsAutoBlackhole struct {
	value *MemberDnsAutoBlackhole
	isSet bool
}

func (v NullableMemberDnsAutoBlackhole) Get() *MemberDnsAutoBlackhole {
	return v.value
}

func (v *NullableMemberDnsAutoBlackhole) Set(val *MemberDnsAutoBlackhole) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberDnsAutoBlackhole) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberDnsAutoBlackhole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberDnsAutoBlackhole(val *MemberDnsAutoBlackhole) *NullableMemberDnsAutoBlackhole {
	return &NullableMemberDnsAutoBlackhole{value: val, isSet: true}
}

func (v NullableMemberDnsAutoBlackhole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberDnsAutoBlackhole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
