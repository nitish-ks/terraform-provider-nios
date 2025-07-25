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

// checks if the MemberSyslogProxySetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberSyslogProxySetting{}

// MemberSyslogProxySetting struct for MemberSyslogProxySetting
type MemberSyslogProxySetting struct {
	// If set to True, the member receives syslog messages from specified devices, such as syslog servers and routers, and then forwards these messages to an external syslog server.
	Enable *bool `json:"enable,omitempty"`
	// If set to True, the appliance can receive messages from other devices via TCP.
	TcpEnable *bool `json:"tcp_enable,omitempty"`
	// The TCP port the appliance must listen on.
	TcpPort *int64 `json:"tcp_port,omitempty"`
	// If set to True, the appliance can receive messages from other devices via UDP.
	UdpEnable *bool `json:"udp_enable,omitempty"`
	// The UDP port the appliance must listen on.
	UdpPort *int64 `json:"udp_port,omitempty"`
	// This list controls the IP addresses and networks that are allowed to access the syslog proxy.
	ClientAcls           []MembersyslogproxysettingClientAcls `json:"client_acls,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemberSyslogProxySetting MemberSyslogProxySetting

// NewMemberSyslogProxySetting instantiates a new MemberSyslogProxySetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberSyslogProxySetting() *MemberSyslogProxySetting {
	this := MemberSyslogProxySetting{}
	return &this
}

// NewMemberSyslogProxySettingWithDefaults instantiates a new MemberSyslogProxySetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberSyslogProxySettingWithDefaults() *MemberSyslogProxySetting {
	this := MemberSyslogProxySetting{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *MemberSyslogProxySetting) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberSyslogProxySetting) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *MemberSyslogProxySetting) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *MemberSyslogProxySetting) SetEnable(v bool) {
	o.Enable = &v
}

// GetTcpEnable returns the TcpEnable field value if set, zero value otherwise.
func (o *MemberSyslogProxySetting) GetTcpEnable() bool {
	if o == nil || IsNil(o.TcpEnable) {
		var ret bool
		return ret
	}
	return *o.TcpEnable
}

// GetTcpEnableOk returns a tuple with the TcpEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberSyslogProxySetting) GetTcpEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.TcpEnable) {
		return nil, false
	}
	return o.TcpEnable, true
}

// HasTcpEnable returns a boolean if a field has been set.
func (o *MemberSyslogProxySetting) HasTcpEnable() bool {
	if o != nil && !IsNil(o.TcpEnable) {
		return true
	}

	return false
}

// SetTcpEnable gets a reference to the given bool and assigns it to the TcpEnable field.
func (o *MemberSyslogProxySetting) SetTcpEnable(v bool) {
	o.TcpEnable = &v
}

// GetTcpPort returns the TcpPort field value if set, zero value otherwise.
func (o *MemberSyslogProxySetting) GetTcpPort() int64 {
	if o == nil || IsNil(o.TcpPort) {
		var ret int64
		return ret
	}
	return *o.TcpPort
}

// GetTcpPortOk returns a tuple with the TcpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberSyslogProxySetting) GetTcpPortOk() (*int64, bool) {
	if o == nil || IsNil(o.TcpPort) {
		return nil, false
	}
	return o.TcpPort, true
}

// HasTcpPort returns a boolean if a field has been set.
func (o *MemberSyslogProxySetting) HasTcpPort() bool {
	if o != nil && !IsNil(o.TcpPort) {
		return true
	}

	return false
}

// SetTcpPort gets a reference to the given int64 and assigns it to the TcpPort field.
func (o *MemberSyslogProxySetting) SetTcpPort(v int64) {
	o.TcpPort = &v
}

// GetUdpEnable returns the UdpEnable field value if set, zero value otherwise.
func (o *MemberSyslogProxySetting) GetUdpEnable() bool {
	if o == nil || IsNil(o.UdpEnable) {
		var ret bool
		return ret
	}
	return *o.UdpEnable
}

// GetUdpEnableOk returns a tuple with the UdpEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberSyslogProxySetting) GetUdpEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.UdpEnable) {
		return nil, false
	}
	return o.UdpEnable, true
}

// HasUdpEnable returns a boolean if a field has been set.
func (o *MemberSyslogProxySetting) HasUdpEnable() bool {
	if o != nil && !IsNil(o.UdpEnable) {
		return true
	}

	return false
}

// SetUdpEnable gets a reference to the given bool and assigns it to the UdpEnable field.
func (o *MemberSyslogProxySetting) SetUdpEnable(v bool) {
	o.UdpEnable = &v
}

// GetUdpPort returns the UdpPort field value if set, zero value otherwise.
func (o *MemberSyslogProxySetting) GetUdpPort() int64 {
	if o == nil || IsNil(o.UdpPort) {
		var ret int64
		return ret
	}
	return *o.UdpPort
}

// GetUdpPortOk returns a tuple with the UdpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberSyslogProxySetting) GetUdpPortOk() (*int64, bool) {
	if o == nil || IsNil(o.UdpPort) {
		return nil, false
	}
	return o.UdpPort, true
}

// HasUdpPort returns a boolean if a field has been set.
func (o *MemberSyslogProxySetting) HasUdpPort() bool {
	if o != nil && !IsNil(o.UdpPort) {
		return true
	}

	return false
}

// SetUdpPort gets a reference to the given int64 and assigns it to the UdpPort field.
func (o *MemberSyslogProxySetting) SetUdpPort(v int64) {
	o.UdpPort = &v
}

// GetClientAcls returns the ClientAcls field value if set, zero value otherwise.
func (o *MemberSyslogProxySetting) GetClientAcls() []MembersyslogproxysettingClientAcls {
	if o == nil || IsNil(o.ClientAcls) {
		var ret []MembersyslogproxysettingClientAcls
		return ret
	}
	return o.ClientAcls
}

// GetClientAclsOk returns a tuple with the ClientAcls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberSyslogProxySetting) GetClientAclsOk() ([]MembersyslogproxysettingClientAcls, bool) {
	if o == nil || IsNil(o.ClientAcls) {
		return nil, false
	}
	return o.ClientAcls, true
}

// HasClientAcls returns a boolean if a field has been set.
func (o *MemberSyslogProxySetting) HasClientAcls() bool {
	if o != nil && !IsNil(o.ClientAcls) {
		return true
	}

	return false
}

// SetClientAcls gets a reference to the given []MembersyslogproxysettingClientAcls and assigns it to the ClientAcls field.
func (o *MemberSyslogProxySetting) SetClientAcls(v []MembersyslogproxysettingClientAcls) {
	o.ClientAcls = v
}

func (o MemberSyslogProxySetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberSyslogProxySetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.TcpEnable) {
		toSerialize["tcp_enable"] = o.TcpEnable
	}
	if !IsNil(o.TcpPort) {
		toSerialize["tcp_port"] = o.TcpPort
	}
	if !IsNil(o.UdpEnable) {
		toSerialize["udp_enable"] = o.UdpEnable
	}
	if !IsNil(o.UdpPort) {
		toSerialize["udp_port"] = o.UdpPort
	}
	if !IsNil(o.ClientAcls) {
		toSerialize["client_acls"] = o.ClientAcls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MemberSyslogProxySetting) UnmarshalJSON(data []byte) (err error) {
	varMemberSyslogProxySetting := _MemberSyslogProxySetting{}

	err = json.Unmarshal(data, &varMemberSyslogProxySetting)

	if err != nil {
		return err
	}

	*o = MemberSyslogProxySetting(varMemberSyslogProxySetting)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enable")
		delete(additionalProperties, "tcp_enable")
		delete(additionalProperties, "tcp_port")
		delete(additionalProperties, "udp_enable")
		delete(additionalProperties, "udp_port")
		delete(additionalProperties, "client_acls")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMemberSyslogProxySetting struct {
	value *MemberSyslogProxySetting
	isSet bool
}

func (v NullableMemberSyslogProxySetting) Get() *MemberSyslogProxySetting {
	return v.value
}

func (v *NullableMemberSyslogProxySetting) Set(val *MemberSyslogProxySetting) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberSyslogProxySetting) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberSyslogProxySetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberSyslogProxySetting(val *MemberSyslogProxySetting) *NullableMemberSyslogProxySetting {
	return &NullableMemberSyslogProxySetting{value: val, isSet: true}
}

func (v NullableMemberSyslogProxySetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberSyslogProxySetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
