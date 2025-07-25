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

// checks if the GridNtpSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GridNtpSetting{}

// GridNtpSetting struct for GridNtpSetting
type GridNtpSetting struct {
	// Determines whether NTP is enabled on the Grid.
	EnableNtp *bool `json:"enable_ntp,omitempty"`
	// The list of NTP servers configured on a Grid.
	NtpServers []GridntpsettingNtpServers `json:"ntp_servers,omitempty"`
	// The list of NTP authentication keys used to authenticate NTP clients.
	NtpKeys []GridntpsettingNtpKeys `json:"ntp_keys,omitempty"`
	NtpAcl  *GridntpsettingNtpAcl   `json:"ntp_acl,omitempty"`
	// Determines whether the Kiss-o'-Death packets are enabled.
	NtpKod *bool `json:"ntp_kod,omitempty"`
	// Grid level GM local NTP stratum.
	GmLocalNtpStratum *int64 `json:"gm_local_ntp_stratum,omitempty"`
	// Local NTP stratum for non-GM members.
	LocalNtpStratum *int64 `json:"local_ntp_stratum,omitempty"`
	// This flag controls whether gm_local_ntp_stratum value be set to a default value
	UseDefaultStratum    *bool `json:"use_default_stratum,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GridNtpSetting GridNtpSetting

// NewGridNtpSetting instantiates a new GridNtpSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGridNtpSetting() *GridNtpSetting {
	this := GridNtpSetting{}
	return &this
}

// NewGridNtpSettingWithDefaults instantiates a new GridNtpSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridNtpSettingWithDefaults() *GridNtpSetting {
	this := GridNtpSetting{}
	return &this
}

// GetEnableNtp returns the EnableNtp field value if set, zero value otherwise.
func (o *GridNtpSetting) GetEnableNtp() bool {
	if o == nil || IsNil(o.EnableNtp) {
		var ret bool
		return ret
	}
	return *o.EnableNtp
}

// GetEnableNtpOk returns a tuple with the EnableNtp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridNtpSetting) GetEnableNtpOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableNtp) {
		return nil, false
	}
	return o.EnableNtp, true
}

// HasEnableNtp returns a boolean if a field has been set.
func (o *GridNtpSetting) HasEnableNtp() bool {
	if o != nil && !IsNil(o.EnableNtp) {
		return true
	}

	return false
}

// SetEnableNtp gets a reference to the given bool and assigns it to the EnableNtp field.
func (o *GridNtpSetting) SetEnableNtp(v bool) {
	o.EnableNtp = &v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise.
func (o *GridNtpSetting) GetNtpServers() []GridntpsettingNtpServers {
	if o == nil || IsNil(o.NtpServers) {
		var ret []GridntpsettingNtpServers
		return ret
	}
	return o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridNtpSetting) GetNtpServersOk() ([]GridntpsettingNtpServers, bool) {
	if o == nil || IsNil(o.NtpServers) {
		return nil, false
	}
	return o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *GridNtpSetting) HasNtpServers() bool {
	if o != nil && !IsNil(o.NtpServers) {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []GridntpsettingNtpServers and assigns it to the NtpServers field.
func (o *GridNtpSetting) SetNtpServers(v []GridntpsettingNtpServers) {
	o.NtpServers = v
}

// GetNtpKeys returns the NtpKeys field value if set, zero value otherwise.
func (o *GridNtpSetting) GetNtpKeys() []GridntpsettingNtpKeys {
	if o == nil || IsNil(o.NtpKeys) {
		var ret []GridntpsettingNtpKeys
		return ret
	}
	return o.NtpKeys
}

// GetNtpKeysOk returns a tuple with the NtpKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridNtpSetting) GetNtpKeysOk() ([]GridntpsettingNtpKeys, bool) {
	if o == nil || IsNil(o.NtpKeys) {
		return nil, false
	}
	return o.NtpKeys, true
}

// HasNtpKeys returns a boolean if a field has been set.
func (o *GridNtpSetting) HasNtpKeys() bool {
	if o != nil && !IsNil(o.NtpKeys) {
		return true
	}

	return false
}

// SetNtpKeys gets a reference to the given []GridntpsettingNtpKeys and assigns it to the NtpKeys field.
func (o *GridNtpSetting) SetNtpKeys(v []GridntpsettingNtpKeys) {
	o.NtpKeys = v
}

// GetNtpAcl returns the NtpAcl field value if set, zero value otherwise.
func (o *GridNtpSetting) GetNtpAcl() GridntpsettingNtpAcl {
	if o == nil || IsNil(o.NtpAcl) {
		var ret GridntpsettingNtpAcl
		return ret
	}
	return *o.NtpAcl
}

// GetNtpAclOk returns a tuple with the NtpAcl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridNtpSetting) GetNtpAclOk() (*GridntpsettingNtpAcl, bool) {
	if o == nil || IsNil(o.NtpAcl) {
		return nil, false
	}
	return o.NtpAcl, true
}

// HasNtpAcl returns a boolean if a field has been set.
func (o *GridNtpSetting) HasNtpAcl() bool {
	if o != nil && !IsNil(o.NtpAcl) {
		return true
	}

	return false
}

// SetNtpAcl gets a reference to the given GridntpsettingNtpAcl and assigns it to the NtpAcl field.
func (o *GridNtpSetting) SetNtpAcl(v GridntpsettingNtpAcl) {
	o.NtpAcl = &v
}

// GetNtpKod returns the NtpKod field value if set, zero value otherwise.
func (o *GridNtpSetting) GetNtpKod() bool {
	if o == nil || IsNil(o.NtpKod) {
		var ret bool
		return ret
	}
	return *o.NtpKod
}

// GetNtpKodOk returns a tuple with the NtpKod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridNtpSetting) GetNtpKodOk() (*bool, bool) {
	if o == nil || IsNil(o.NtpKod) {
		return nil, false
	}
	return o.NtpKod, true
}

// HasNtpKod returns a boolean if a field has been set.
func (o *GridNtpSetting) HasNtpKod() bool {
	if o != nil && !IsNil(o.NtpKod) {
		return true
	}

	return false
}

// SetNtpKod gets a reference to the given bool and assigns it to the NtpKod field.
func (o *GridNtpSetting) SetNtpKod(v bool) {
	o.NtpKod = &v
}

// GetGmLocalNtpStratum returns the GmLocalNtpStratum field value if set, zero value otherwise.
func (o *GridNtpSetting) GetGmLocalNtpStratum() int64 {
	if o == nil || IsNil(o.GmLocalNtpStratum) {
		var ret int64
		return ret
	}
	return *o.GmLocalNtpStratum
}

// GetGmLocalNtpStratumOk returns a tuple with the GmLocalNtpStratum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridNtpSetting) GetGmLocalNtpStratumOk() (*int64, bool) {
	if o == nil || IsNil(o.GmLocalNtpStratum) {
		return nil, false
	}
	return o.GmLocalNtpStratum, true
}

// HasGmLocalNtpStratum returns a boolean if a field has been set.
func (o *GridNtpSetting) HasGmLocalNtpStratum() bool {
	if o != nil && !IsNil(o.GmLocalNtpStratum) {
		return true
	}

	return false
}

// SetGmLocalNtpStratum gets a reference to the given int64 and assigns it to the GmLocalNtpStratum field.
func (o *GridNtpSetting) SetGmLocalNtpStratum(v int64) {
	o.GmLocalNtpStratum = &v
}

// GetLocalNtpStratum returns the LocalNtpStratum field value if set, zero value otherwise.
func (o *GridNtpSetting) GetLocalNtpStratum() int64 {
	if o == nil || IsNil(o.LocalNtpStratum) {
		var ret int64
		return ret
	}
	return *o.LocalNtpStratum
}

// GetLocalNtpStratumOk returns a tuple with the LocalNtpStratum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridNtpSetting) GetLocalNtpStratumOk() (*int64, bool) {
	if o == nil || IsNil(o.LocalNtpStratum) {
		return nil, false
	}
	return o.LocalNtpStratum, true
}

// HasLocalNtpStratum returns a boolean if a field has been set.
func (o *GridNtpSetting) HasLocalNtpStratum() bool {
	if o != nil && !IsNil(o.LocalNtpStratum) {
		return true
	}

	return false
}

// SetLocalNtpStratum gets a reference to the given int64 and assigns it to the LocalNtpStratum field.
func (o *GridNtpSetting) SetLocalNtpStratum(v int64) {
	o.LocalNtpStratum = &v
}

// GetUseDefaultStratum returns the UseDefaultStratum field value if set, zero value otherwise.
func (o *GridNtpSetting) GetUseDefaultStratum() bool {
	if o == nil || IsNil(o.UseDefaultStratum) {
		var ret bool
		return ret
	}
	return *o.UseDefaultStratum
}

// GetUseDefaultStratumOk returns a tuple with the UseDefaultStratum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridNtpSetting) GetUseDefaultStratumOk() (*bool, bool) {
	if o == nil || IsNil(o.UseDefaultStratum) {
		return nil, false
	}
	return o.UseDefaultStratum, true
}

// HasUseDefaultStratum returns a boolean if a field has been set.
func (o *GridNtpSetting) HasUseDefaultStratum() bool {
	if o != nil && !IsNil(o.UseDefaultStratum) {
		return true
	}

	return false
}

// SetUseDefaultStratum gets a reference to the given bool and assigns it to the UseDefaultStratum field.
func (o *GridNtpSetting) SetUseDefaultStratum(v bool) {
	o.UseDefaultStratum = &v
}

func (o GridNtpSetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GridNtpSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnableNtp) {
		toSerialize["enable_ntp"] = o.EnableNtp
	}
	if !IsNil(o.NtpServers) {
		toSerialize["ntp_servers"] = o.NtpServers
	}
	if !IsNil(o.NtpKeys) {
		toSerialize["ntp_keys"] = o.NtpKeys
	}
	if !IsNil(o.NtpAcl) {
		toSerialize["ntp_acl"] = o.NtpAcl
	}
	if !IsNil(o.NtpKod) {
		toSerialize["ntp_kod"] = o.NtpKod
	}
	if !IsNil(o.GmLocalNtpStratum) {
		toSerialize["gm_local_ntp_stratum"] = o.GmLocalNtpStratum
	}
	if !IsNil(o.LocalNtpStratum) {
		toSerialize["local_ntp_stratum"] = o.LocalNtpStratum
	}
	if !IsNil(o.UseDefaultStratum) {
		toSerialize["use_default_stratum"] = o.UseDefaultStratum
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GridNtpSetting) UnmarshalJSON(data []byte) (err error) {
	varGridNtpSetting := _GridNtpSetting{}

	err = json.Unmarshal(data, &varGridNtpSetting)

	if err != nil {
		return err
	}

	*o = GridNtpSetting(varGridNtpSetting)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enable_ntp")
		delete(additionalProperties, "ntp_servers")
		delete(additionalProperties, "ntp_keys")
		delete(additionalProperties, "ntp_acl")
		delete(additionalProperties, "ntp_kod")
		delete(additionalProperties, "gm_local_ntp_stratum")
		delete(additionalProperties, "local_ntp_stratum")
		delete(additionalProperties, "use_default_stratum")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGridNtpSetting struct {
	value *GridNtpSetting
	isSet bool
}

func (v NullableGridNtpSetting) Get() *GridNtpSetting {
	return v.value
}

func (v *NullableGridNtpSetting) Set(val *GridNtpSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableGridNtpSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableGridNtpSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridNtpSetting(val *GridNtpSetting) *NullableGridNtpSetting {
	return &NullableGridNtpSetting{value: val, isSet: true}
}

func (v NullableGridNtpSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridNtpSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
