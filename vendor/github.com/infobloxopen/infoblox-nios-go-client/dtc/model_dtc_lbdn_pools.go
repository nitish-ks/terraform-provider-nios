/*
Infoblox DTC API

OpenAPI specification for Infoblox NIOS WAPI DTC objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dtc

import (
	"encoding/json"
)

// checks if the DtcLbdnPools type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtcLbdnPools{}

// DtcLbdnPools struct for DtcLbdnPools
type DtcLbdnPools struct {
	// The pool to link with.
	Pool *string `json:"pool,omitempty"`
	// The weight of pool.
	Ratio                *int64 `json:"ratio,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DtcLbdnPools DtcLbdnPools

// NewDtcLbdnPools instantiates a new DtcLbdnPools object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtcLbdnPools() *DtcLbdnPools {
	this := DtcLbdnPools{}
	return &this
}

// NewDtcLbdnPoolsWithDefaults instantiates a new DtcLbdnPools object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtcLbdnPoolsWithDefaults() *DtcLbdnPools {
	this := DtcLbdnPools{}
	return &this
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *DtcLbdnPools) GetPool() string {
	if o == nil || IsNil(o.Pool) {
		var ret string
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcLbdnPools) GetPoolOk() (*string, bool) {
	if o == nil || IsNil(o.Pool) {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *DtcLbdnPools) HasPool() bool {
	if o != nil && !IsNil(o.Pool) {
		return true
	}

	return false
}

// SetPool gets a reference to the given string and assigns it to the Pool field.
func (o *DtcLbdnPools) SetPool(v string) {
	o.Pool = &v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *DtcLbdnPools) GetRatio() int64 {
	if o == nil || IsNil(o.Ratio) {
		var ret int64
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcLbdnPools) GetRatioOk() (*int64, bool) {
	if o == nil || IsNil(o.Ratio) {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *DtcLbdnPools) HasRatio() bool {
	if o != nil && !IsNil(o.Ratio) {
		return true
	}

	return false
}

// SetRatio gets a reference to the given int64 and assigns it to the Ratio field.
func (o *DtcLbdnPools) SetRatio(v int64) {
	o.Ratio = &v
}

func (o DtcLbdnPools) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtcLbdnPools) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pool) {
		toSerialize["pool"] = o.Pool
	}
	if !IsNil(o.Ratio) {
		toSerialize["ratio"] = o.Ratio
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DtcLbdnPools) UnmarshalJSON(data []byte) (err error) {
	varDtcLbdnPools := _DtcLbdnPools{}

	err = json.Unmarshal(data, &varDtcLbdnPools)

	if err != nil {
		return err
	}

	*o = DtcLbdnPools(varDtcLbdnPools)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pool")
		delete(additionalProperties, "ratio")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDtcLbdnPools struct {
	value *DtcLbdnPools
	isSet bool
}

func (v NullableDtcLbdnPools) Get() *DtcLbdnPools {
	return v.value
}

func (v *NullableDtcLbdnPools) Set(val *DtcLbdnPools) {
	v.value = val
	v.isSet = true
}

func (v NullableDtcLbdnPools) IsSet() bool {
	return v.isSet
}

func (v *NullableDtcLbdnPools) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtcLbdnPools(val *DtcLbdnPools) *NullableDtcLbdnPools {
	return &NullableDtcLbdnPools{value: val, isSet: true}
}

func (v NullableDtcLbdnPools) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtcLbdnPools) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
