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

// checks if the GridSnmpSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GridSnmpSetting{}

// GridSnmpSetting struct for GridSnmpSetting
type GridSnmpSetting struct {
	// The engine ID of the appliance that manages the SNMP agent.
	EngineId []string `json:"engine_id,omitempty"`
	// The community string for SNMP queries.
	QueriesCommunityString *string `json:"queries_community_string,omitempty"`
	// If set to True, SNMP queries are enabled.
	QueriesEnable *bool `json:"queries_enable,omitempty"`
	// If set to True, SNMPv3 queries are enabled.
	Snmpv3QueriesEnable *bool `json:"snmpv3_queries_enable,omitempty"`
	// A list of SNMPv3 queries users.
	Snmpv3QueriesUsers []GridsnmpsettingSnmpv3QueriesUsers `json:"snmpv3_queries_users,omitempty"`
	// If set to True, SNMPv3 traps are enabled.
	Snmpv3TrapsEnable *bool `json:"snmpv3_traps_enable,omitempty"`
	// The name of the contact person for the appliance. Second value is applicable only for HA pair. Otherwise second value is ignored.
	Syscontact []string `json:"syscontact,omitempty"`
	// Useful information about the appliance. Second value is applicable only for HA pair. Otherwise second value is ignored.
	Sysdescr []string `json:"sysdescr,omitempty"`
	// The physical location of the appliance. Second value is applicable only for HA pair. Otherwise second value is ignored.
	Syslocation []string `json:"syslocation,omitempty"`
	// The FQDN (Fully Qualified Domain Name) of the appliance. Second value is applicable only for HA pair. Otherwise second value is ignored.
	Sysname []string `json:"sysname,omitempty"`
	// A list of trap receivers.
	TrapReceivers []GridsnmpsettingTrapReceivers `json:"trap_receivers,omitempty"`
	// A string the NIOS appliance sends to the management system together with its traps. Note that this community string must match exactly what you enter in the management system.
	TrapsCommunityString *string `json:"traps_community_string,omitempty"`
	// If set to True, SNMP traps are enabled.
	TrapsEnable          *bool `json:"traps_enable,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GridSnmpSetting GridSnmpSetting

// NewGridSnmpSetting instantiates a new GridSnmpSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGridSnmpSetting() *GridSnmpSetting {
	this := GridSnmpSetting{}
	return &this
}

// NewGridSnmpSettingWithDefaults instantiates a new GridSnmpSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridSnmpSettingWithDefaults() *GridSnmpSetting {
	this := GridSnmpSetting{}
	return &this
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *GridSnmpSetting) GetEngineId() []string {
	if o == nil || IsNil(o.EngineId) {
		var ret []string
		return ret
	}
	return o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridSnmpSetting) GetEngineIdOk() ([]string, bool) {
	if o == nil || IsNil(o.EngineId) {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *GridSnmpSetting) HasEngineId() bool {
	if o != nil && !IsNil(o.EngineId) {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given []string and assigns it to the EngineId field.
func (o *GridSnmpSetting) SetEngineId(v []string) {
	o.EngineId = v
}

// GetQueriesCommunityString returns the QueriesCommunityString field value if set, zero value otherwise.
func (o *GridSnmpSetting) GetQueriesCommunityString() string {
	if o == nil || IsNil(o.QueriesCommunityString) {
		var ret string
		return ret
	}
	return *o.QueriesCommunityString
}

// GetQueriesCommunityStringOk returns a tuple with the QueriesCommunityString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridSnmpSetting) GetQueriesCommunityStringOk() (*string, bool) {
	if o == nil || IsNil(o.QueriesCommunityString) {
		return nil, false
	}
	return o.QueriesCommunityString, true
}

// HasQueriesCommunityString returns a boolean if a field has been set.
func (o *GridSnmpSetting) HasQueriesCommunityString() bool {
	if o != nil && !IsNil(o.QueriesCommunityString) {
		return true
	}

	return false
}

// SetQueriesCommunityString gets a reference to the given string and assigns it to the QueriesCommunityString field.
func (o *GridSnmpSetting) SetQueriesCommunityString(v string) {
	o.QueriesCommunityString = &v
}

// GetQueriesEnable returns the QueriesEnable field value if set, zero value otherwise.
func (o *GridSnmpSetting) GetQueriesEnable() bool {
	if o == nil || IsNil(o.QueriesEnable) {
		var ret bool
		return ret
	}
	return *o.QueriesEnable
}

// GetQueriesEnableOk returns a tuple with the QueriesEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridSnmpSetting) GetQueriesEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.QueriesEnable) {
		return nil, false
	}
	return o.QueriesEnable, true
}

// HasQueriesEnable returns a boolean if a field has been set.
func (o *GridSnmpSetting) HasQueriesEnable() bool {
	if o != nil && !IsNil(o.QueriesEnable) {
		return true
	}

	return false
}

// SetQueriesEnable gets a reference to the given bool and assigns it to the QueriesEnable field.
func (o *GridSnmpSetting) SetQueriesEnable(v bool) {
	o.QueriesEnable = &v
}

// GetSnmpv3QueriesEnable returns the Snmpv3QueriesEnable field value if set, zero value otherwise.
func (o *GridSnmpSetting) GetSnmpv3QueriesEnable() bool {
	if o == nil || IsNil(o.Snmpv3QueriesEnable) {
		var ret bool
		return ret
	}
	return *o.Snmpv3QueriesEnable
}

// GetSnmpv3QueriesEnableOk returns a tuple with the Snmpv3QueriesEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridSnmpSetting) GetSnmpv3QueriesEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Snmpv3QueriesEnable) {
		return nil, false
	}
	return o.Snmpv3QueriesEnable, true
}

// HasSnmpv3QueriesEnable returns a boolean if a field has been set.
func (o *GridSnmpSetting) HasSnmpv3QueriesEnable() bool {
	if o != nil && !IsNil(o.Snmpv3QueriesEnable) {
		return true
	}

	return false
}

// SetSnmpv3QueriesEnable gets a reference to the given bool and assigns it to the Snmpv3QueriesEnable field.
func (o *GridSnmpSetting) SetSnmpv3QueriesEnable(v bool) {
	o.Snmpv3QueriesEnable = &v
}

// GetSnmpv3QueriesUsers returns the Snmpv3QueriesUsers field value if set, zero value otherwise.
func (o *GridSnmpSetting) GetSnmpv3QueriesUsers() []GridsnmpsettingSnmpv3QueriesUsers {
	if o == nil || IsNil(o.Snmpv3QueriesUsers) {
		var ret []GridsnmpsettingSnmpv3QueriesUsers
		return ret
	}
	return o.Snmpv3QueriesUsers
}

// GetSnmpv3QueriesUsersOk returns a tuple with the Snmpv3QueriesUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridSnmpSetting) GetSnmpv3QueriesUsersOk() ([]GridsnmpsettingSnmpv3QueriesUsers, bool) {
	if o == nil || IsNil(o.Snmpv3QueriesUsers) {
		return nil, false
	}
	return o.Snmpv3QueriesUsers, true
}

// HasSnmpv3QueriesUsers returns a boolean if a field has been set.
func (o *GridSnmpSetting) HasSnmpv3QueriesUsers() bool {
	if o != nil && !IsNil(o.Snmpv3QueriesUsers) {
		return true
	}

	return false
}

// SetSnmpv3QueriesUsers gets a reference to the given []GridsnmpsettingSnmpv3QueriesUsers and assigns it to the Snmpv3QueriesUsers field.
func (o *GridSnmpSetting) SetSnmpv3QueriesUsers(v []GridsnmpsettingSnmpv3QueriesUsers) {
	o.Snmpv3QueriesUsers = v
}

// GetSnmpv3TrapsEnable returns the Snmpv3TrapsEnable field value if set, zero value otherwise.
func (o *GridSnmpSetting) GetSnmpv3TrapsEnable() bool {
	if o == nil || IsNil(o.Snmpv3TrapsEnable) {
		var ret bool
		return ret
	}
	return *o.Snmpv3TrapsEnable
}

// GetSnmpv3TrapsEnableOk returns a tuple with the Snmpv3TrapsEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridSnmpSetting) GetSnmpv3TrapsEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Snmpv3TrapsEnable) {
		return nil, false
	}
	return o.Snmpv3TrapsEnable, true
}

// HasSnmpv3TrapsEnable returns a boolean if a field has been set.
func (o *GridSnmpSetting) HasSnmpv3TrapsEnable() bool {
	if o != nil && !IsNil(o.Snmpv3TrapsEnable) {
		return true
	}

	return false
}

// SetSnmpv3TrapsEnable gets a reference to the given bool and assigns it to the Snmpv3TrapsEnable field.
func (o *GridSnmpSetting) SetSnmpv3TrapsEnable(v bool) {
	o.Snmpv3TrapsEnable = &v
}

// GetSyscontact returns the Syscontact field value if set, zero value otherwise.
func (o *GridSnmpSetting) GetSyscontact() []string {
	if o == nil || IsNil(o.Syscontact) {
		var ret []string
		return ret
	}
	return o.Syscontact
}

// GetSyscontactOk returns a tuple with the Syscontact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridSnmpSetting) GetSyscontactOk() ([]string, bool) {
	if o == nil || IsNil(o.Syscontact) {
		return nil, false
	}
	return o.Syscontact, true
}

// HasSyscontact returns a boolean if a field has been set.
func (o *GridSnmpSetting) HasSyscontact() bool {
	if o != nil && !IsNil(o.Syscontact) {
		return true
	}

	return false
}

// SetSyscontact gets a reference to the given []string and assigns it to the Syscontact field.
func (o *GridSnmpSetting) SetSyscontact(v []string) {
	o.Syscontact = v
}

// GetSysdescr returns the Sysdescr field value if set, zero value otherwise.
func (o *GridSnmpSetting) GetSysdescr() []string {
	if o == nil || IsNil(o.Sysdescr) {
		var ret []string
		return ret
	}
	return o.Sysdescr
}

// GetSysdescrOk returns a tuple with the Sysdescr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridSnmpSetting) GetSysdescrOk() ([]string, bool) {
	if o == nil || IsNil(o.Sysdescr) {
		return nil, false
	}
	return o.Sysdescr, true
}

// HasSysdescr returns a boolean if a field has been set.
func (o *GridSnmpSetting) HasSysdescr() bool {
	if o != nil && !IsNil(o.Sysdescr) {
		return true
	}

	return false
}

// SetSysdescr gets a reference to the given []string and assigns it to the Sysdescr field.
func (o *GridSnmpSetting) SetSysdescr(v []string) {
	o.Sysdescr = v
}

// GetSyslocation returns the Syslocation field value if set, zero value otherwise.
func (o *GridSnmpSetting) GetSyslocation() []string {
	if o == nil || IsNil(o.Syslocation) {
		var ret []string
		return ret
	}
	return o.Syslocation
}

// GetSyslocationOk returns a tuple with the Syslocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridSnmpSetting) GetSyslocationOk() ([]string, bool) {
	if o == nil || IsNil(o.Syslocation) {
		return nil, false
	}
	return o.Syslocation, true
}

// HasSyslocation returns a boolean if a field has been set.
func (o *GridSnmpSetting) HasSyslocation() bool {
	if o != nil && !IsNil(o.Syslocation) {
		return true
	}

	return false
}

// SetSyslocation gets a reference to the given []string and assigns it to the Syslocation field.
func (o *GridSnmpSetting) SetSyslocation(v []string) {
	o.Syslocation = v
}

// GetSysname returns the Sysname field value if set, zero value otherwise.
func (o *GridSnmpSetting) GetSysname() []string {
	if o == nil || IsNil(o.Sysname) {
		var ret []string
		return ret
	}
	return o.Sysname
}

// GetSysnameOk returns a tuple with the Sysname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridSnmpSetting) GetSysnameOk() ([]string, bool) {
	if o == nil || IsNil(o.Sysname) {
		return nil, false
	}
	return o.Sysname, true
}

// HasSysname returns a boolean if a field has been set.
func (o *GridSnmpSetting) HasSysname() bool {
	if o != nil && !IsNil(o.Sysname) {
		return true
	}

	return false
}

// SetSysname gets a reference to the given []string and assigns it to the Sysname field.
func (o *GridSnmpSetting) SetSysname(v []string) {
	o.Sysname = v
}

// GetTrapReceivers returns the TrapReceivers field value if set, zero value otherwise.
func (o *GridSnmpSetting) GetTrapReceivers() []GridsnmpsettingTrapReceivers {
	if o == nil || IsNil(o.TrapReceivers) {
		var ret []GridsnmpsettingTrapReceivers
		return ret
	}
	return o.TrapReceivers
}

// GetTrapReceiversOk returns a tuple with the TrapReceivers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridSnmpSetting) GetTrapReceiversOk() ([]GridsnmpsettingTrapReceivers, bool) {
	if o == nil || IsNil(o.TrapReceivers) {
		return nil, false
	}
	return o.TrapReceivers, true
}

// HasTrapReceivers returns a boolean if a field has been set.
func (o *GridSnmpSetting) HasTrapReceivers() bool {
	if o != nil && !IsNil(o.TrapReceivers) {
		return true
	}

	return false
}

// SetTrapReceivers gets a reference to the given []GridsnmpsettingTrapReceivers and assigns it to the TrapReceivers field.
func (o *GridSnmpSetting) SetTrapReceivers(v []GridsnmpsettingTrapReceivers) {
	o.TrapReceivers = v
}

// GetTrapsCommunityString returns the TrapsCommunityString field value if set, zero value otherwise.
func (o *GridSnmpSetting) GetTrapsCommunityString() string {
	if o == nil || IsNil(o.TrapsCommunityString) {
		var ret string
		return ret
	}
	return *o.TrapsCommunityString
}

// GetTrapsCommunityStringOk returns a tuple with the TrapsCommunityString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridSnmpSetting) GetTrapsCommunityStringOk() (*string, bool) {
	if o == nil || IsNil(o.TrapsCommunityString) {
		return nil, false
	}
	return o.TrapsCommunityString, true
}

// HasTrapsCommunityString returns a boolean if a field has been set.
func (o *GridSnmpSetting) HasTrapsCommunityString() bool {
	if o != nil && !IsNil(o.TrapsCommunityString) {
		return true
	}

	return false
}

// SetTrapsCommunityString gets a reference to the given string and assigns it to the TrapsCommunityString field.
func (o *GridSnmpSetting) SetTrapsCommunityString(v string) {
	o.TrapsCommunityString = &v
}

// GetTrapsEnable returns the TrapsEnable field value if set, zero value otherwise.
func (o *GridSnmpSetting) GetTrapsEnable() bool {
	if o == nil || IsNil(o.TrapsEnable) {
		var ret bool
		return ret
	}
	return *o.TrapsEnable
}

// GetTrapsEnableOk returns a tuple with the TrapsEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridSnmpSetting) GetTrapsEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.TrapsEnable) {
		return nil, false
	}
	return o.TrapsEnable, true
}

// HasTrapsEnable returns a boolean if a field has been set.
func (o *GridSnmpSetting) HasTrapsEnable() bool {
	if o != nil && !IsNil(o.TrapsEnable) {
		return true
	}

	return false
}

// SetTrapsEnable gets a reference to the given bool and assigns it to the TrapsEnable field.
func (o *GridSnmpSetting) SetTrapsEnable(v bool) {
	o.TrapsEnable = &v
}

func (o GridSnmpSetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GridSnmpSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EngineId) {
		toSerialize["engine_id"] = o.EngineId
	}
	if !IsNil(o.QueriesCommunityString) {
		toSerialize["queries_community_string"] = o.QueriesCommunityString
	}
	if !IsNil(o.QueriesEnable) {
		toSerialize["queries_enable"] = o.QueriesEnable
	}
	if !IsNil(o.Snmpv3QueriesEnable) {
		toSerialize["snmpv3_queries_enable"] = o.Snmpv3QueriesEnable
	}
	if !IsNil(o.Snmpv3QueriesUsers) {
		toSerialize["snmpv3_queries_users"] = o.Snmpv3QueriesUsers
	}
	if !IsNil(o.Snmpv3TrapsEnable) {
		toSerialize["snmpv3_traps_enable"] = o.Snmpv3TrapsEnable
	}
	if !IsNil(o.Syscontact) {
		toSerialize["syscontact"] = o.Syscontact
	}
	if !IsNil(o.Sysdescr) {
		toSerialize["sysdescr"] = o.Sysdescr
	}
	if !IsNil(o.Syslocation) {
		toSerialize["syslocation"] = o.Syslocation
	}
	if !IsNil(o.Sysname) {
		toSerialize["sysname"] = o.Sysname
	}
	if !IsNil(o.TrapReceivers) {
		toSerialize["trap_receivers"] = o.TrapReceivers
	}
	if !IsNil(o.TrapsCommunityString) {
		toSerialize["traps_community_string"] = o.TrapsCommunityString
	}
	if !IsNil(o.TrapsEnable) {
		toSerialize["traps_enable"] = o.TrapsEnable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GridSnmpSetting) UnmarshalJSON(data []byte) (err error) {
	varGridSnmpSetting := _GridSnmpSetting{}

	err = json.Unmarshal(data, &varGridSnmpSetting)

	if err != nil {
		return err
	}

	*o = GridSnmpSetting(varGridSnmpSetting)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "engine_id")
		delete(additionalProperties, "queries_community_string")
		delete(additionalProperties, "queries_enable")
		delete(additionalProperties, "snmpv3_queries_enable")
		delete(additionalProperties, "snmpv3_queries_users")
		delete(additionalProperties, "snmpv3_traps_enable")
		delete(additionalProperties, "syscontact")
		delete(additionalProperties, "sysdescr")
		delete(additionalProperties, "syslocation")
		delete(additionalProperties, "sysname")
		delete(additionalProperties, "trap_receivers")
		delete(additionalProperties, "traps_community_string")
		delete(additionalProperties, "traps_enable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGridSnmpSetting struct {
	value *GridSnmpSetting
	isSet bool
}

func (v NullableGridSnmpSetting) Get() *GridSnmpSetting {
	return v.value
}

func (v *NullableGridSnmpSetting) Set(val *GridSnmpSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableGridSnmpSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableGridSnmpSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridSnmpSetting(val *GridSnmpSetting) *NullableGridSnmpSetting {
	return &NullableGridSnmpSetting{value: val, isSet: true}
}

func (v NullableGridSnmpSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridSnmpSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
