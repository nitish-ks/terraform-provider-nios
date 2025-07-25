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

// checks if the DiscoveryGridproperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscoveryGridproperties{}

// DiscoveryGridproperties struct for DiscoveryGridproperties
type DiscoveryGridproperties struct {
	// The reference to the object.
	Ref                        *string                                            `json:"_ref,omitempty"`
	AdvancedPollingSettings    *DiscoveryGridpropertiesAdvancedPollingSettings    `json:"advanced_polling_settings,omitempty"`
	AdvancedSdnPollingSettings *DiscoveryGridpropertiesAdvancedSdnPollingSettings `json:"advanced_sdn_polling_settings,omitempty"`
	AdvisorSettings            *DiscoveryGridpropertiesAdvisorSettings            `json:"advisor_settings,omitempty"`
	// Automatic conversion settings.
	AutoConversionSettings  []DiscoveryGridpropertiesAutoConversionSettings `json:"auto_conversion_settings,omitempty"`
	BasicPollingSettings    *DiscoveryGridpropertiesBasicPollingSettings    `json:"basic_polling_settings,omitempty"`
	BasicSdnPollingSettings *DiscoveryGridpropertiesBasicSdnPollingSettings `json:"basic_sdn_polling_settings,omitempty"`
	// Discovery CLI credentials.
	CliCredentials []DiscoveryGridpropertiesCliCredentials `json:"cli_credentials,omitempty"`
	// Device Hints.
	DeviceHints              []DiscoveryGridpropertiesDeviceHints             `json:"device_hints,omitempty"`
	DiscoveryBlackoutSetting *DiscoveryGridpropertiesDiscoveryBlackoutSetting `json:"discovery_blackout_setting,omitempty"`
	// The type of the devices the DNS processor operates on.
	DnsLookupOption *string `json:"dns_lookup_option,omitempty"`
	// The percentage of available capacity the DNS processor operates at. Valid values are unsigned integer between 1 and 100, inclusive.
	DnsLookupThrottle *int64 `json:"dns_lookup_throttle,omitempty"`
	// Advisor application enabled/disabled.
	EnableAdvisor *bool `json:"enable_advisor,omitempty"`
	// The flag that enables automatic conversion of discovered data.
	EnableAutoConversion *bool `json:"enable_auto_conversion,omitempty"`
	// The flag that enables updating discovered data for managed objects.
	EnableAutoUpdates *bool `json:"enable_auto_updates,omitempty"`
	// The Grid name.
	GridName *string `json:"grid_name,omitempty"`
	// Determines the timeout to ignore the discovery conflict duration (in seconds).
	IgnoreConflictDuration     *int64                                             `json:"ignore_conflict_duration,omitempty"`
	PortControlBlackoutSetting *DiscoveryGridpropertiesPortControlBlackoutSetting `json:"port_control_blackout_setting,omitempty"`
	// Ports to scan.
	Ports []DiscoveryGridpropertiesPorts `json:"ports,omitempty"`
	// Determines if the same port control is used for discovery blackout.
	SamePortControlDiscoveryBlackout *bool `json:"same_port_control_discovery_blackout,omitempty"`
	// Discovery SNMP v1 and v2 credentials.
	Snmpv1v2Credentials []DiscoveryGridpropertiesSnmpv1v2Credentials `json:"snmpv1v2_credentials,omitempty"`
	// Discovery SNMP v3 credentials.
	Snmpv3Credentials []DiscoveryGridpropertiesSnmpv3Credentials `json:"snmpv3_credentials,omitempty"`
	// Limit of discovered unmanaged IP address which determines how frequently the user is notified about the new unmanaged IP address in a particular network.
	UnmanagedIpsLimit *int64 `json:"unmanaged_ips_limit,omitempty"`
	// Determines the timeout between two notifications (in seconds) about the new unmanaged IP address in a particular network. The value must be between 60 seconds and the number of seconds remaining to Jan 2038.
	UnmanagedIpsTimeout *int64 `json:"unmanaged_ips_timeout,omitempty"`
	// The policy type used to define the behavior of the VRF mapping.
	VrfMappingPolicy *string `json:"vrf_mapping_policy,omitempty"`
	// VRF mapping rules.
	VrfMappingRules []DiscoveryGridpropertiesVrfMappingRules `json:"vrf_mapping_rules,omitempty"`
}

// NewDiscoveryGridproperties instantiates a new DiscoveryGridproperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveryGridproperties() *DiscoveryGridproperties {
	this := DiscoveryGridproperties{}
	return &this
}

// NewDiscoveryGridpropertiesWithDefaults instantiates a new DiscoveryGridproperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveryGridpropertiesWithDefaults() *DiscoveryGridproperties {
	this := DiscoveryGridproperties{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *DiscoveryGridproperties) SetRef(v string) {
	o.Ref = &v
}

// GetAdvancedPollingSettings returns the AdvancedPollingSettings field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetAdvancedPollingSettings() DiscoveryGridpropertiesAdvancedPollingSettings {
	if o == nil || IsNil(o.AdvancedPollingSettings) {
		var ret DiscoveryGridpropertiesAdvancedPollingSettings
		return ret
	}
	return *o.AdvancedPollingSettings
}

// GetAdvancedPollingSettingsOk returns a tuple with the AdvancedPollingSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetAdvancedPollingSettingsOk() (*DiscoveryGridpropertiesAdvancedPollingSettings, bool) {
	if o == nil || IsNil(o.AdvancedPollingSettings) {
		return nil, false
	}
	return o.AdvancedPollingSettings, true
}

// HasAdvancedPollingSettings returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasAdvancedPollingSettings() bool {
	if o != nil && !IsNil(o.AdvancedPollingSettings) {
		return true
	}

	return false
}

// SetAdvancedPollingSettings gets a reference to the given DiscoveryGridpropertiesAdvancedPollingSettings and assigns it to the AdvancedPollingSettings field.
func (o *DiscoveryGridproperties) SetAdvancedPollingSettings(v DiscoveryGridpropertiesAdvancedPollingSettings) {
	o.AdvancedPollingSettings = &v
}

// GetAdvancedSdnPollingSettings returns the AdvancedSdnPollingSettings field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetAdvancedSdnPollingSettings() DiscoveryGridpropertiesAdvancedSdnPollingSettings {
	if o == nil || IsNil(o.AdvancedSdnPollingSettings) {
		var ret DiscoveryGridpropertiesAdvancedSdnPollingSettings
		return ret
	}
	return *o.AdvancedSdnPollingSettings
}

// GetAdvancedSdnPollingSettingsOk returns a tuple with the AdvancedSdnPollingSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetAdvancedSdnPollingSettingsOk() (*DiscoveryGridpropertiesAdvancedSdnPollingSettings, bool) {
	if o == nil || IsNil(o.AdvancedSdnPollingSettings) {
		return nil, false
	}
	return o.AdvancedSdnPollingSettings, true
}

// HasAdvancedSdnPollingSettings returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasAdvancedSdnPollingSettings() bool {
	if o != nil && !IsNil(o.AdvancedSdnPollingSettings) {
		return true
	}

	return false
}

// SetAdvancedSdnPollingSettings gets a reference to the given DiscoveryGridpropertiesAdvancedSdnPollingSettings and assigns it to the AdvancedSdnPollingSettings field.
func (o *DiscoveryGridproperties) SetAdvancedSdnPollingSettings(v DiscoveryGridpropertiesAdvancedSdnPollingSettings) {
	o.AdvancedSdnPollingSettings = &v
}

// GetAdvisorSettings returns the AdvisorSettings field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetAdvisorSettings() DiscoveryGridpropertiesAdvisorSettings {
	if o == nil || IsNil(o.AdvisorSettings) {
		var ret DiscoveryGridpropertiesAdvisorSettings
		return ret
	}
	return *o.AdvisorSettings
}

// GetAdvisorSettingsOk returns a tuple with the AdvisorSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetAdvisorSettingsOk() (*DiscoveryGridpropertiesAdvisorSettings, bool) {
	if o == nil || IsNil(o.AdvisorSettings) {
		return nil, false
	}
	return o.AdvisorSettings, true
}

// HasAdvisorSettings returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasAdvisorSettings() bool {
	if o != nil && !IsNil(o.AdvisorSettings) {
		return true
	}

	return false
}

// SetAdvisorSettings gets a reference to the given DiscoveryGridpropertiesAdvisorSettings and assigns it to the AdvisorSettings field.
func (o *DiscoveryGridproperties) SetAdvisorSettings(v DiscoveryGridpropertiesAdvisorSettings) {
	o.AdvisorSettings = &v
}

// GetAutoConversionSettings returns the AutoConversionSettings field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetAutoConversionSettings() []DiscoveryGridpropertiesAutoConversionSettings {
	if o == nil || IsNil(o.AutoConversionSettings) {
		var ret []DiscoveryGridpropertiesAutoConversionSettings
		return ret
	}
	return o.AutoConversionSettings
}

// GetAutoConversionSettingsOk returns a tuple with the AutoConversionSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetAutoConversionSettingsOk() ([]DiscoveryGridpropertiesAutoConversionSettings, bool) {
	if o == nil || IsNil(o.AutoConversionSettings) {
		return nil, false
	}
	return o.AutoConversionSettings, true
}

// HasAutoConversionSettings returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasAutoConversionSettings() bool {
	if o != nil && !IsNil(o.AutoConversionSettings) {
		return true
	}

	return false
}

// SetAutoConversionSettings gets a reference to the given []DiscoveryGridpropertiesAutoConversionSettings and assigns it to the AutoConversionSettings field.
func (o *DiscoveryGridproperties) SetAutoConversionSettings(v []DiscoveryGridpropertiesAutoConversionSettings) {
	o.AutoConversionSettings = v
}

// GetBasicPollingSettings returns the BasicPollingSettings field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetBasicPollingSettings() DiscoveryGridpropertiesBasicPollingSettings {
	if o == nil || IsNil(o.BasicPollingSettings) {
		var ret DiscoveryGridpropertiesBasicPollingSettings
		return ret
	}
	return *o.BasicPollingSettings
}

// GetBasicPollingSettingsOk returns a tuple with the BasicPollingSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetBasicPollingSettingsOk() (*DiscoveryGridpropertiesBasicPollingSettings, bool) {
	if o == nil || IsNil(o.BasicPollingSettings) {
		return nil, false
	}
	return o.BasicPollingSettings, true
}

// HasBasicPollingSettings returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasBasicPollingSettings() bool {
	if o != nil && !IsNil(o.BasicPollingSettings) {
		return true
	}

	return false
}

// SetBasicPollingSettings gets a reference to the given DiscoveryGridpropertiesBasicPollingSettings and assigns it to the BasicPollingSettings field.
func (o *DiscoveryGridproperties) SetBasicPollingSettings(v DiscoveryGridpropertiesBasicPollingSettings) {
	o.BasicPollingSettings = &v
}

// GetBasicSdnPollingSettings returns the BasicSdnPollingSettings field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetBasicSdnPollingSettings() DiscoveryGridpropertiesBasicSdnPollingSettings {
	if o == nil || IsNil(o.BasicSdnPollingSettings) {
		var ret DiscoveryGridpropertiesBasicSdnPollingSettings
		return ret
	}
	return *o.BasicSdnPollingSettings
}

// GetBasicSdnPollingSettingsOk returns a tuple with the BasicSdnPollingSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetBasicSdnPollingSettingsOk() (*DiscoveryGridpropertiesBasicSdnPollingSettings, bool) {
	if o == nil || IsNil(o.BasicSdnPollingSettings) {
		return nil, false
	}
	return o.BasicSdnPollingSettings, true
}

// HasBasicSdnPollingSettings returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasBasicSdnPollingSettings() bool {
	if o != nil && !IsNil(o.BasicSdnPollingSettings) {
		return true
	}

	return false
}

// SetBasicSdnPollingSettings gets a reference to the given DiscoveryGridpropertiesBasicSdnPollingSettings and assigns it to the BasicSdnPollingSettings field.
func (o *DiscoveryGridproperties) SetBasicSdnPollingSettings(v DiscoveryGridpropertiesBasicSdnPollingSettings) {
	o.BasicSdnPollingSettings = &v
}

// GetCliCredentials returns the CliCredentials field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetCliCredentials() []DiscoveryGridpropertiesCliCredentials {
	if o == nil || IsNil(o.CliCredentials) {
		var ret []DiscoveryGridpropertiesCliCredentials
		return ret
	}
	return o.CliCredentials
}

// GetCliCredentialsOk returns a tuple with the CliCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetCliCredentialsOk() ([]DiscoveryGridpropertiesCliCredentials, bool) {
	if o == nil || IsNil(o.CliCredentials) {
		return nil, false
	}
	return o.CliCredentials, true
}

// HasCliCredentials returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasCliCredentials() bool {
	if o != nil && !IsNil(o.CliCredentials) {
		return true
	}

	return false
}

// SetCliCredentials gets a reference to the given []DiscoveryGridpropertiesCliCredentials and assigns it to the CliCredentials field.
func (o *DiscoveryGridproperties) SetCliCredentials(v []DiscoveryGridpropertiesCliCredentials) {
	o.CliCredentials = v
}

// GetDeviceHints returns the DeviceHints field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetDeviceHints() []DiscoveryGridpropertiesDeviceHints {
	if o == nil || IsNil(o.DeviceHints) {
		var ret []DiscoveryGridpropertiesDeviceHints
		return ret
	}
	return o.DeviceHints
}

// GetDeviceHintsOk returns a tuple with the DeviceHints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetDeviceHintsOk() ([]DiscoveryGridpropertiesDeviceHints, bool) {
	if o == nil || IsNil(o.DeviceHints) {
		return nil, false
	}
	return o.DeviceHints, true
}

// HasDeviceHints returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasDeviceHints() bool {
	if o != nil && !IsNil(o.DeviceHints) {
		return true
	}

	return false
}

// SetDeviceHints gets a reference to the given []DiscoveryGridpropertiesDeviceHints and assigns it to the DeviceHints field.
func (o *DiscoveryGridproperties) SetDeviceHints(v []DiscoveryGridpropertiesDeviceHints) {
	o.DeviceHints = v
}

// GetDiscoveryBlackoutSetting returns the DiscoveryBlackoutSetting field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetDiscoveryBlackoutSetting() DiscoveryGridpropertiesDiscoveryBlackoutSetting {
	if o == nil || IsNil(o.DiscoveryBlackoutSetting) {
		var ret DiscoveryGridpropertiesDiscoveryBlackoutSetting
		return ret
	}
	return *o.DiscoveryBlackoutSetting
}

// GetDiscoveryBlackoutSettingOk returns a tuple with the DiscoveryBlackoutSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetDiscoveryBlackoutSettingOk() (*DiscoveryGridpropertiesDiscoveryBlackoutSetting, bool) {
	if o == nil || IsNil(o.DiscoveryBlackoutSetting) {
		return nil, false
	}
	return o.DiscoveryBlackoutSetting, true
}

// HasDiscoveryBlackoutSetting returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasDiscoveryBlackoutSetting() bool {
	if o != nil && !IsNil(o.DiscoveryBlackoutSetting) {
		return true
	}

	return false
}

// SetDiscoveryBlackoutSetting gets a reference to the given DiscoveryGridpropertiesDiscoveryBlackoutSetting and assigns it to the DiscoveryBlackoutSetting field.
func (o *DiscoveryGridproperties) SetDiscoveryBlackoutSetting(v DiscoveryGridpropertiesDiscoveryBlackoutSetting) {
	o.DiscoveryBlackoutSetting = &v
}

// GetDnsLookupOption returns the DnsLookupOption field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetDnsLookupOption() string {
	if o == nil || IsNil(o.DnsLookupOption) {
		var ret string
		return ret
	}
	return *o.DnsLookupOption
}

// GetDnsLookupOptionOk returns a tuple with the DnsLookupOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetDnsLookupOptionOk() (*string, bool) {
	if o == nil || IsNil(o.DnsLookupOption) {
		return nil, false
	}
	return o.DnsLookupOption, true
}

// HasDnsLookupOption returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasDnsLookupOption() bool {
	if o != nil && !IsNil(o.DnsLookupOption) {
		return true
	}

	return false
}

// SetDnsLookupOption gets a reference to the given string and assigns it to the DnsLookupOption field.
func (o *DiscoveryGridproperties) SetDnsLookupOption(v string) {
	o.DnsLookupOption = &v
}

// GetDnsLookupThrottle returns the DnsLookupThrottle field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetDnsLookupThrottle() int64 {
	if o == nil || IsNil(o.DnsLookupThrottle) {
		var ret int64
		return ret
	}
	return *o.DnsLookupThrottle
}

// GetDnsLookupThrottleOk returns a tuple with the DnsLookupThrottle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetDnsLookupThrottleOk() (*int64, bool) {
	if o == nil || IsNil(o.DnsLookupThrottle) {
		return nil, false
	}
	return o.DnsLookupThrottle, true
}

// HasDnsLookupThrottle returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasDnsLookupThrottle() bool {
	if o != nil && !IsNil(o.DnsLookupThrottle) {
		return true
	}

	return false
}

// SetDnsLookupThrottle gets a reference to the given int64 and assigns it to the DnsLookupThrottle field.
func (o *DiscoveryGridproperties) SetDnsLookupThrottle(v int64) {
	o.DnsLookupThrottle = &v
}

// GetEnableAdvisor returns the EnableAdvisor field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetEnableAdvisor() bool {
	if o == nil || IsNil(o.EnableAdvisor) {
		var ret bool
		return ret
	}
	return *o.EnableAdvisor
}

// GetEnableAdvisorOk returns a tuple with the EnableAdvisor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetEnableAdvisorOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAdvisor) {
		return nil, false
	}
	return o.EnableAdvisor, true
}

// HasEnableAdvisor returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasEnableAdvisor() bool {
	if o != nil && !IsNil(o.EnableAdvisor) {
		return true
	}

	return false
}

// SetEnableAdvisor gets a reference to the given bool and assigns it to the EnableAdvisor field.
func (o *DiscoveryGridproperties) SetEnableAdvisor(v bool) {
	o.EnableAdvisor = &v
}

// GetEnableAutoConversion returns the EnableAutoConversion field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetEnableAutoConversion() bool {
	if o == nil || IsNil(o.EnableAutoConversion) {
		var ret bool
		return ret
	}
	return *o.EnableAutoConversion
}

// GetEnableAutoConversionOk returns a tuple with the EnableAutoConversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetEnableAutoConversionOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAutoConversion) {
		return nil, false
	}
	return o.EnableAutoConversion, true
}

// HasEnableAutoConversion returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasEnableAutoConversion() bool {
	if o != nil && !IsNil(o.EnableAutoConversion) {
		return true
	}

	return false
}

// SetEnableAutoConversion gets a reference to the given bool and assigns it to the EnableAutoConversion field.
func (o *DiscoveryGridproperties) SetEnableAutoConversion(v bool) {
	o.EnableAutoConversion = &v
}

// GetEnableAutoUpdates returns the EnableAutoUpdates field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetEnableAutoUpdates() bool {
	if o == nil || IsNil(o.EnableAutoUpdates) {
		var ret bool
		return ret
	}
	return *o.EnableAutoUpdates
}

// GetEnableAutoUpdatesOk returns a tuple with the EnableAutoUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetEnableAutoUpdatesOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAutoUpdates) {
		return nil, false
	}
	return o.EnableAutoUpdates, true
}

// HasEnableAutoUpdates returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasEnableAutoUpdates() bool {
	if o != nil && !IsNil(o.EnableAutoUpdates) {
		return true
	}

	return false
}

// SetEnableAutoUpdates gets a reference to the given bool and assigns it to the EnableAutoUpdates field.
func (o *DiscoveryGridproperties) SetEnableAutoUpdates(v bool) {
	o.EnableAutoUpdates = &v
}

// GetGridName returns the GridName field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetGridName() string {
	if o == nil || IsNil(o.GridName) {
		var ret string
		return ret
	}
	return *o.GridName
}

// GetGridNameOk returns a tuple with the GridName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetGridNameOk() (*string, bool) {
	if o == nil || IsNil(o.GridName) {
		return nil, false
	}
	return o.GridName, true
}

// HasGridName returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasGridName() bool {
	if o != nil && !IsNil(o.GridName) {
		return true
	}

	return false
}

// SetGridName gets a reference to the given string and assigns it to the GridName field.
func (o *DiscoveryGridproperties) SetGridName(v string) {
	o.GridName = &v
}

// GetIgnoreConflictDuration returns the IgnoreConflictDuration field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetIgnoreConflictDuration() int64 {
	if o == nil || IsNil(o.IgnoreConflictDuration) {
		var ret int64
		return ret
	}
	return *o.IgnoreConflictDuration
}

// GetIgnoreConflictDurationOk returns a tuple with the IgnoreConflictDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetIgnoreConflictDurationOk() (*int64, bool) {
	if o == nil || IsNil(o.IgnoreConflictDuration) {
		return nil, false
	}
	return o.IgnoreConflictDuration, true
}

// HasIgnoreConflictDuration returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasIgnoreConflictDuration() bool {
	if o != nil && !IsNil(o.IgnoreConflictDuration) {
		return true
	}

	return false
}

// SetIgnoreConflictDuration gets a reference to the given int64 and assigns it to the IgnoreConflictDuration field.
func (o *DiscoveryGridproperties) SetIgnoreConflictDuration(v int64) {
	o.IgnoreConflictDuration = &v
}

// GetPortControlBlackoutSetting returns the PortControlBlackoutSetting field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetPortControlBlackoutSetting() DiscoveryGridpropertiesPortControlBlackoutSetting {
	if o == nil || IsNil(o.PortControlBlackoutSetting) {
		var ret DiscoveryGridpropertiesPortControlBlackoutSetting
		return ret
	}
	return *o.PortControlBlackoutSetting
}

// GetPortControlBlackoutSettingOk returns a tuple with the PortControlBlackoutSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetPortControlBlackoutSettingOk() (*DiscoveryGridpropertiesPortControlBlackoutSetting, bool) {
	if o == nil || IsNil(o.PortControlBlackoutSetting) {
		return nil, false
	}
	return o.PortControlBlackoutSetting, true
}

// HasPortControlBlackoutSetting returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasPortControlBlackoutSetting() bool {
	if o != nil && !IsNil(o.PortControlBlackoutSetting) {
		return true
	}

	return false
}

// SetPortControlBlackoutSetting gets a reference to the given DiscoveryGridpropertiesPortControlBlackoutSetting and assigns it to the PortControlBlackoutSetting field.
func (o *DiscoveryGridproperties) SetPortControlBlackoutSetting(v DiscoveryGridpropertiesPortControlBlackoutSetting) {
	o.PortControlBlackoutSetting = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetPorts() []DiscoveryGridpropertiesPorts {
	if o == nil || IsNil(o.Ports) {
		var ret []DiscoveryGridpropertiesPorts
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetPortsOk() ([]DiscoveryGridpropertiesPorts, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []DiscoveryGridpropertiesPorts and assigns it to the Ports field.
func (o *DiscoveryGridproperties) SetPorts(v []DiscoveryGridpropertiesPorts) {
	o.Ports = v
}

// GetSamePortControlDiscoveryBlackout returns the SamePortControlDiscoveryBlackout field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetSamePortControlDiscoveryBlackout() bool {
	if o == nil || IsNil(o.SamePortControlDiscoveryBlackout) {
		var ret bool
		return ret
	}
	return *o.SamePortControlDiscoveryBlackout
}

// GetSamePortControlDiscoveryBlackoutOk returns a tuple with the SamePortControlDiscoveryBlackout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetSamePortControlDiscoveryBlackoutOk() (*bool, bool) {
	if o == nil || IsNil(o.SamePortControlDiscoveryBlackout) {
		return nil, false
	}
	return o.SamePortControlDiscoveryBlackout, true
}

// HasSamePortControlDiscoveryBlackout returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasSamePortControlDiscoveryBlackout() bool {
	if o != nil && !IsNil(o.SamePortControlDiscoveryBlackout) {
		return true
	}

	return false
}

// SetSamePortControlDiscoveryBlackout gets a reference to the given bool and assigns it to the SamePortControlDiscoveryBlackout field.
func (o *DiscoveryGridproperties) SetSamePortControlDiscoveryBlackout(v bool) {
	o.SamePortControlDiscoveryBlackout = &v
}

// GetSnmpv1v2Credentials returns the Snmpv1v2Credentials field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetSnmpv1v2Credentials() []DiscoveryGridpropertiesSnmpv1v2Credentials {
	if o == nil || IsNil(o.Snmpv1v2Credentials) {
		var ret []DiscoveryGridpropertiesSnmpv1v2Credentials
		return ret
	}
	return o.Snmpv1v2Credentials
}

// GetSnmpv1v2CredentialsOk returns a tuple with the Snmpv1v2Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetSnmpv1v2CredentialsOk() ([]DiscoveryGridpropertiesSnmpv1v2Credentials, bool) {
	if o == nil || IsNil(o.Snmpv1v2Credentials) {
		return nil, false
	}
	return o.Snmpv1v2Credentials, true
}

// HasSnmpv1v2Credentials returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasSnmpv1v2Credentials() bool {
	if o != nil && !IsNil(o.Snmpv1v2Credentials) {
		return true
	}

	return false
}

// SetSnmpv1v2Credentials gets a reference to the given []DiscoveryGridpropertiesSnmpv1v2Credentials and assigns it to the Snmpv1v2Credentials field.
func (o *DiscoveryGridproperties) SetSnmpv1v2Credentials(v []DiscoveryGridpropertiesSnmpv1v2Credentials) {
	o.Snmpv1v2Credentials = v
}

// GetSnmpv3Credentials returns the Snmpv3Credentials field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetSnmpv3Credentials() []DiscoveryGridpropertiesSnmpv3Credentials {
	if o == nil || IsNil(o.Snmpv3Credentials) {
		var ret []DiscoveryGridpropertiesSnmpv3Credentials
		return ret
	}
	return o.Snmpv3Credentials
}

// GetSnmpv3CredentialsOk returns a tuple with the Snmpv3Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetSnmpv3CredentialsOk() ([]DiscoveryGridpropertiesSnmpv3Credentials, bool) {
	if o == nil || IsNil(o.Snmpv3Credentials) {
		return nil, false
	}
	return o.Snmpv3Credentials, true
}

// HasSnmpv3Credentials returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasSnmpv3Credentials() bool {
	if o != nil && !IsNil(o.Snmpv3Credentials) {
		return true
	}

	return false
}

// SetSnmpv3Credentials gets a reference to the given []DiscoveryGridpropertiesSnmpv3Credentials and assigns it to the Snmpv3Credentials field.
func (o *DiscoveryGridproperties) SetSnmpv3Credentials(v []DiscoveryGridpropertiesSnmpv3Credentials) {
	o.Snmpv3Credentials = v
}

// GetUnmanagedIpsLimit returns the UnmanagedIpsLimit field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetUnmanagedIpsLimit() int64 {
	if o == nil || IsNil(o.UnmanagedIpsLimit) {
		var ret int64
		return ret
	}
	return *o.UnmanagedIpsLimit
}

// GetUnmanagedIpsLimitOk returns a tuple with the UnmanagedIpsLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetUnmanagedIpsLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.UnmanagedIpsLimit) {
		return nil, false
	}
	return o.UnmanagedIpsLimit, true
}

// HasUnmanagedIpsLimit returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasUnmanagedIpsLimit() bool {
	if o != nil && !IsNil(o.UnmanagedIpsLimit) {
		return true
	}

	return false
}

// SetUnmanagedIpsLimit gets a reference to the given int64 and assigns it to the UnmanagedIpsLimit field.
func (o *DiscoveryGridproperties) SetUnmanagedIpsLimit(v int64) {
	o.UnmanagedIpsLimit = &v
}

// GetUnmanagedIpsTimeout returns the UnmanagedIpsTimeout field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetUnmanagedIpsTimeout() int64 {
	if o == nil || IsNil(o.UnmanagedIpsTimeout) {
		var ret int64
		return ret
	}
	return *o.UnmanagedIpsTimeout
}

// GetUnmanagedIpsTimeoutOk returns a tuple with the UnmanagedIpsTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetUnmanagedIpsTimeoutOk() (*int64, bool) {
	if o == nil || IsNil(o.UnmanagedIpsTimeout) {
		return nil, false
	}
	return o.UnmanagedIpsTimeout, true
}

// HasUnmanagedIpsTimeout returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasUnmanagedIpsTimeout() bool {
	if o != nil && !IsNil(o.UnmanagedIpsTimeout) {
		return true
	}

	return false
}

// SetUnmanagedIpsTimeout gets a reference to the given int64 and assigns it to the UnmanagedIpsTimeout field.
func (o *DiscoveryGridproperties) SetUnmanagedIpsTimeout(v int64) {
	o.UnmanagedIpsTimeout = &v
}

// GetVrfMappingPolicy returns the VrfMappingPolicy field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetVrfMappingPolicy() string {
	if o == nil || IsNil(o.VrfMappingPolicy) {
		var ret string
		return ret
	}
	return *o.VrfMappingPolicy
}

// GetVrfMappingPolicyOk returns a tuple with the VrfMappingPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetVrfMappingPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.VrfMappingPolicy) {
		return nil, false
	}
	return o.VrfMappingPolicy, true
}

// HasVrfMappingPolicy returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasVrfMappingPolicy() bool {
	if o != nil && !IsNil(o.VrfMappingPolicy) {
		return true
	}

	return false
}

// SetVrfMappingPolicy gets a reference to the given string and assigns it to the VrfMappingPolicy field.
func (o *DiscoveryGridproperties) SetVrfMappingPolicy(v string) {
	o.VrfMappingPolicy = &v
}

// GetVrfMappingRules returns the VrfMappingRules field value if set, zero value otherwise.
func (o *DiscoveryGridproperties) GetVrfMappingRules() []DiscoveryGridpropertiesVrfMappingRules {
	if o == nil || IsNil(o.VrfMappingRules) {
		var ret []DiscoveryGridpropertiesVrfMappingRules
		return ret
	}
	return o.VrfMappingRules
}

// GetVrfMappingRulesOk returns a tuple with the VrfMappingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryGridproperties) GetVrfMappingRulesOk() ([]DiscoveryGridpropertiesVrfMappingRules, bool) {
	if o == nil || IsNil(o.VrfMappingRules) {
		return nil, false
	}
	return o.VrfMappingRules, true
}

// HasVrfMappingRules returns a boolean if a field has been set.
func (o *DiscoveryGridproperties) HasVrfMappingRules() bool {
	if o != nil && !IsNil(o.VrfMappingRules) {
		return true
	}

	return false
}

// SetVrfMappingRules gets a reference to the given []DiscoveryGridpropertiesVrfMappingRules and assigns it to the VrfMappingRules field.
func (o *DiscoveryGridproperties) SetVrfMappingRules(v []DiscoveryGridpropertiesVrfMappingRules) {
	o.VrfMappingRules = v
}

func (o DiscoveryGridproperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveryGridproperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.AdvancedPollingSettings) {
		toSerialize["advanced_polling_settings"] = o.AdvancedPollingSettings
	}
	if !IsNil(o.AdvancedSdnPollingSettings) {
		toSerialize["advanced_sdn_polling_settings"] = o.AdvancedSdnPollingSettings
	}
	if !IsNil(o.AdvisorSettings) {
		toSerialize["advisor_settings"] = o.AdvisorSettings
	}
	if !IsNil(o.AutoConversionSettings) {
		toSerialize["auto_conversion_settings"] = o.AutoConversionSettings
	}
	if !IsNil(o.BasicPollingSettings) {
		toSerialize["basic_polling_settings"] = o.BasicPollingSettings
	}
	if !IsNil(o.BasicSdnPollingSettings) {
		toSerialize["basic_sdn_polling_settings"] = o.BasicSdnPollingSettings
	}
	if !IsNil(o.CliCredentials) {
		toSerialize["cli_credentials"] = o.CliCredentials
	}
	if !IsNil(o.DeviceHints) {
		toSerialize["device_hints"] = o.DeviceHints
	}
	if !IsNil(o.DiscoveryBlackoutSetting) {
		toSerialize["discovery_blackout_setting"] = o.DiscoveryBlackoutSetting
	}
	if !IsNil(o.DnsLookupOption) {
		toSerialize["dns_lookup_option"] = o.DnsLookupOption
	}
	if !IsNil(o.DnsLookupThrottle) {
		toSerialize["dns_lookup_throttle"] = o.DnsLookupThrottle
	}
	if !IsNil(o.EnableAdvisor) {
		toSerialize["enable_advisor"] = o.EnableAdvisor
	}
	if !IsNil(o.EnableAutoConversion) {
		toSerialize["enable_auto_conversion"] = o.EnableAutoConversion
	}
	if !IsNil(o.EnableAutoUpdates) {
		toSerialize["enable_auto_updates"] = o.EnableAutoUpdates
	}
	if !IsNil(o.GridName) {
		toSerialize["grid_name"] = o.GridName
	}
	if !IsNil(o.IgnoreConflictDuration) {
		toSerialize["ignore_conflict_duration"] = o.IgnoreConflictDuration
	}
	if !IsNil(o.PortControlBlackoutSetting) {
		toSerialize["port_control_blackout_setting"] = o.PortControlBlackoutSetting
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.SamePortControlDiscoveryBlackout) {
		toSerialize["same_port_control_discovery_blackout"] = o.SamePortControlDiscoveryBlackout
	}
	if !IsNil(o.Snmpv1v2Credentials) {
		toSerialize["snmpv1v2_credentials"] = o.Snmpv1v2Credentials
	}
	if !IsNil(o.Snmpv3Credentials) {
		toSerialize["snmpv3_credentials"] = o.Snmpv3Credentials
	}
	if !IsNil(o.UnmanagedIpsLimit) {
		toSerialize["unmanaged_ips_limit"] = o.UnmanagedIpsLimit
	}
	if !IsNil(o.UnmanagedIpsTimeout) {
		toSerialize["unmanaged_ips_timeout"] = o.UnmanagedIpsTimeout
	}
	if !IsNil(o.VrfMappingPolicy) {
		toSerialize["vrf_mapping_policy"] = o.VrfMappingPolicy
	}
	if !IsNil(o.VrfMappingRules) {
		toSerialize["vrf_mapping_rules"] = o.VrfMappingRules
	}
	return toSerialize, nil
}

type NullableDiscoveryGridproperties struct {
	value *DiscoveryGridproperties
	isSet bool
}

func (v NullableDiscoveryGridproperties) Get() *DiscoveryGridproperties {
	return v.value
}

func (v *NullableDiscoveryGridproperties) Set(val *DiscoveryGridproperties) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryGridproperties) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryGridproperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryGridproperties(val *DiscoveryGridproperties) *NullableDiscoveryGridproperties {
	return &NullableDiscoveryGridproperties{value: val, isSet: true}
}

func (v NullableDiscoveryGridproperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryGridproperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
