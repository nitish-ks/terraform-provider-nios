/*
Infoblox DNS API

OpenAPI specification for Infoblox NIOS WAPI DNS objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the RecordAIpv4addrOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordAIpv4addrOneOf{}

// RecordAIpv4addrOneOf record:a: The IPv4 Address of the record.
type RecordAIpv4addrOneOf struct {
	ObjectFunction       *string                `json:"_object_function,omitempty"`
	Parameters           map[string]interface{} `json:"_parameters,omitempty"`
	ResultField          *string                `json:"_result_field,omitempty"`
	Object               *string                `json:"_object,omitempty"`
	ObjectParameters     map[string]interface{} `json:"_object_parameters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecordAIpv4addrOneOf RecordAIpv4addrOneOf

// NewRecordAIpv4addrOneOf instantiates a new RecordAIpv4addrOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordAIpv4addrOneOf() *RecordAIpv4addrOneOf {
	this := RecordAIpv4addrOneOf{}
	return &this
}

// NewRecordAIpv4addrOneOfWithDefaults instantiates a new RecordAIpv4addrOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordAIpv4addrOneOfWithDefaults() *RecordAIpv4addrOneOf {
	this := RecordAIpv4addrOneOf{}
	return &this
}

// GetObjectFunction returns the ObjectFunction field value if set, zero value otherwise.
func (o *RecordAIpv4addrOneOf) GetObjectFunction() string {
	if o == nil || IsNil(o.ObjectFunction) {
		var ret string
		return ret
	}
	return *o.ObjectFunction
}

// GetObjectFunctionOk returns a tuple with the ObjectFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordAIpv4addrOneOf) GetObjectFunctionOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectFunction) {
		return nil, false
	}
	return o.ObjectFunction, true
}

// HasObjectFunction returns a boolean if a field has been set.
func (o *RecordAIpv4addrOneOf) HasObjectFunction() bool {
	if o != nil && !IsNil(o.ObjectFunction) {
		return true
	}

	return false
}

// SetObjectFunction gets a reference to the given string and assigns it to the ObjectFunction field.
func (o *RecordAIpv4addrOneOf) SetObjectFunction(v string) {
	o.ObjectFunction = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *RecordAIpv4addrOneOf) GetParameters() map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordAIpv4addrOneOf) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *RecordAIpv4addrOneOf) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *RecordAIpv4addrOneOf) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

// GetResultField returns the ResultField field value if set, zero value otherwise.
func (o *RecordAIpv4addrOneOf) GetResultField() string {
	if o == nil || IsNil(o.ResultField) {
		var ret string
		return ret
	}
	return *o.ResultField
}

// GetResultFieldOk returns a tuple with the ResultField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordAIpv4addrOneOf) GetResultFieldOk() (*string, bool) {
	if o == nil || IsNil(o.ResultField) {
		return nil, false
	}
	return o.ResultField, true
}

// HasResultField returns a boolean if a field has been set.
func (o *RecordAIpv4addrOneOf) HasResultField() bool {
	if o != nil && !IsNil(o.ResultField) {
		return true
	}

	return false
}

// SetResultField gets a reference to the given string and assigns it to the ResultField field.
func (o *RecordAIpv4addrOneOf) SetResultField(v string) {
	o.ResultField = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *RecordAIpv4addrOneOf) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordAIpv4addrOneOf) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *RecordAIpv4addrOneOf) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *RecordAIpv4addrOneOf) SetObject(v string) {
	o.Object = &v
}

// GetObjectParameters returns the ObjectParameters field value if set, zero value otherwise.
func (o *RecordAIpv4addrOneOf) GetObjectParameters() map[string]interface{} {
	if o == nil || IsNil(o.ObjectParameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.ObjectParameters
}

// GetObjectParametersOk returns a tuple with the ObjectParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordAIpv4addrOneOf) GetObjectParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ObjectParameters) {
		return map[string]interface{}{}, false
	}
	return o.ObjectParameters, true
}

// HasObjectParameters returns a boolean if a field has been set.
func (o *RecordAIpv4addrOneOf) HasObjectParameters() bool {
	if o != nil && !IsNil(o.ObjectParameters) {
		return true
	}

	return false
}

// SetObjectParameters gets a reference to the given map[string]interface{} and assigns it to the ObjectParameters field.
func (o *RecordAIpv4addrOneOf) SetObjectParameters(v map[string]interface{}) {
	o.ObjectParameters = v
}

func (o RecordAIpv4addrOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordAIpv4addrOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObjectFunction) {
		toSerialize["_object_function"] = o.ObjectFunction
	}
	if !IsNil(o.Parameters) {
		toSerialize["_parameters"] = o.Parameters
	}
	if !IsNil(o.ResultField) {
		toSerialize["_result_field"] = o.ResultField
	}
	if !IsNil(o.Object) {
		toSerialize["_object"] = o.Object
	}
	if !IsNil(o.ObjectParameters) {
		toSerialize["_object_parameters"] = o.ObjectParameters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecordAIpv4addrOneOf) UnmarshalJSON(data []byte) (err error) {
	varRecordAIpv4addrOneOf := _RecordAIpv4addrOneOf{}

	err = json.Unmarshal(data, &varRecordAIpv4addrOneOf)

	if err != nil {
		return err
	}

	*o = RecordAIpv4addrOneOf(varRecordAIpv4addrOneOf)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_object_function")
		delete(additionalProperties, "_parameters")
		delete(additionalProperties, "_result_field")
		delete(additionalProperties, "_object")
		delete(additionalProperties, "_object_parameters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecordAIpv4addrOneOf struct {
	value *RecordAIpv4addrOneOf
	isSet bool
}

func (v NullableRecordAIpv4addrOneOf) Get() *RecordAIpv4addrOneOf {
	return v.value
}

func (v *NullableRecordAIpv4addrOneOf) Set(val *RecordAIpv4addrOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordAIpv4addrOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordAIpv4addrOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordAIpv4addrOneOf(val *RecordAIpv4addrOneOf) *NullableRecordAIpv4addrOneOf {
	return &NullableRecordAIpv4addrOneOf{value: val, isSet: true}
}

func (v NullableRecordAIpv4addrOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordAIpv4addrOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
