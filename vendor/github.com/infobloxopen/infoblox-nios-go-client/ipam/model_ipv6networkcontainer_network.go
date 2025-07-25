/*
Infoblox IPAM API

OpenAPI specification for Infoblox NIOS WAPI IPAM objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
	"fmt"
)

// Ipv6networkcontainerNetwork - The network address in IPv6 Address/CIDR format. For regular expression searches, only the IPv6 Address portion is supported. Searches for the CIDR portion is always an exact match. For example, both network containers 16::0/28 and 26::0/24 are matched by expression '.6' and only 26::0/24 is matched by '.6/24'.
type Ipv6networkcontainerNetwork struct {
	Ipv6networkcontainerNetworkOneOf *Ipv6networkcontainerNetworkOneOf
	String                           *string
}

// Ipv6networkcontainerNetworkOneOfAsIpv6networkcontainerNetwork is a convenience function that returns Ipv6networkcontainerNetworkOneOf wrapped in Ipv6networkcontainerNetwork
func Ipv6networkcontainerNetworkOneOfAsIpv6networkcontainerNetwork(v *Ipv6networkcontainerNetworkOneOf) Ipv6networkcontainerNetwork {
	return Ipv6networkcontainerNetwork{
		Ipv6networkcontainerNetworkOneOf: v,
	}
}

// stringAsIpv6networkcontainerNetwork is a convenience function that returns string wrapped in Ipv6networkcontainerNetwork
func StringAsIpv6networkcontainerNetwork(v *string) Ipv6networkcontainerNetwork {
	return Ipv6networkcontainerNetwork{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Ipv6networkcontainerNetwork) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Ipv6networkcontainerNetworkOneOf
	err = newStrictDecoder(data).Decode(&dst.Ipv6networkcontainerNetworkOneOf)
	if err == nil {
		jsonIpv6networkcontainerNetworkOneOf, _ := json.Marshal(dst.Ipv6networkcontainerNetworkOneOf)
		if string(jsonIpv6networkcontainerNetworkOneOf) == "{}" { // empty struct
			dst.Ipv6networkcontainerNetworkOneOf = nil
		} else {
			match++
		}
	} else {
		dst.Ipv6networkcontainerNetworkOneOf = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Ipv6networkcontainerNetworkOneOf = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Ipv6networkcontainerNetwork)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Ipv6networkcontainerNetwork)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Ipv6networkcontainerNetwork) MarshalJSON() ([]byte, error) {
	if src.Ipv6networkcontainerNetworkOneOf != nil {
		return json.Marshal(&src.Ipv6networkcontainerNetworkOneOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Ipv6networkcontainerNetwork) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Ipv6networkcontainerNetworkOneOf != nil {
		return obj.Ipv6networkcontainerNetworkOneOf
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableIpv6networkcontainerNetwork struct {
	value *Ipv6networkcontainerNetwork
	isSet bool
}

func (v NullableIpv6networkcontainerNetwork) Get() *Ipv6networkcontainerNetwork {
	return v.value
}

func (v *NullableIpv6networkcontainerNetwork) Set(val *Ipv6networkcontainerNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv6networkcontainerNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv6networkcontainerNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv6networkcontainerNetwork(val *Ipv6networkcontainerNetwork) *NullableIpv6networkcontainerNetwork {
	return &NullableIpv6networkcontainerNetwork{value: val, isSet: true}
}

func (v NullableIpv6networkcontainerNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv6networkcontainerNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
