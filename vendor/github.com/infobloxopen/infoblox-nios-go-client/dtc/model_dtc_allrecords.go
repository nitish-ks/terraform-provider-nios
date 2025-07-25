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

// checks if the DtcAllrecords type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DtcAllrecords{}

// DtcAllrecords struct for DtcAllrecords
type DtcAllrecords struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// The record comment.
	Comment *string `json:"comment,omitempty"`
	// The disable value determines if the record is disabled or not. \"False\" means the record is enabled.
	Disable *bool `json:"disable,omitempty"`
	// The name of the DTC Server object with which the record is associated.
	DtcServer *string `json:"dtc_server,omitempty"`
	// The record object, if supported by the WAPI. Otherwise, the value is \"None\".
	Record *string `json:"record,omitempty"`
	// The TTL value of the record associated with the DTC AllRecords object.
	Ttl *int64 `json:"ttl,omitempty"`
	// The record type. When searching for an unspecified record type, the search is performed for all records. On retrieval, the appliance returns \"UNSUPPORTED\" for unsupported records.
	Type *string `json:"type,omitempty"`
}

// NewDtcAllrecords instantiates a new DtcAllrecords object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDtcAllrecords() *DtcAllrecords {
	this := DtcAllrecords{}
	return &this
}

// NewDtcAllrecordsWithDefaults instantiates a new DtcAllrecords object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDtcAllrecordsWithDefaults() *DtcAllrecords {
	this := DtcAllrecords{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *DtcAllrecords) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcAllrecords) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *DtcAllrecords) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *DtcAllrecords) SetRef(v string) {
	o.Ref = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *DtcAllrecords) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcAllrecords) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *DtcAllrecords) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *DtcAllrecords) SetComment(v string) {
	o.Comment = &v
}

// GetDisable returns the Disable field value if set, zero value otherwise.
func (o *DtcAllrecords) GetDisable() bool {
	if o == nil || IsNil(o.Disable) {
		var ret bool
		return ret
	}
	return *o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcAllrecords) GetDisableOk() (*bool, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *DtcAllrecords) HasDisable() bool {
	if o != nil && !IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given bool and assigns it to the Disable field.
func (o *DtcAllrecords) SetDisable(v bool) {
	o.Disable = &v
}

// GetDtcServer returns the DtcServer field value if set, zero value otherwise.
func (o *DtcAllrecords) GetDtcServer() string {
	if o == nil || IsNil(o.DtcServer) {
		var ret string
		return ret
	}
	return *o.DtcServer
}

// GetDtcServerOk returns a tuple with the DtcServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcAllrecords) GetDtcServerOk() (*string, bool) {
	if o == nil || IsNil(o.DtcServer) {
		return nil, false
	}
	return o.DtcServer, true
}

// HasDtcServer returns a boolean if a field has been set.
func (o *DtcAllrecords) HasDtcServer() bool {
	if o != nil && !IsNil(o.DtcServer) {
		return true
	}

	return false
}

// SetDtcServer gets a reference to the given string and assigns it to the DtcServer field.
func (o *DtcAllrecords) SetDtcServer(v string) {
	o.DtcServer = &v
}

// GetRecord returns the Record field value if set, zero value otherwise.
func (o *DtcAllrecords) GetRecord() string {
	if o == nil || IsNil(o.Record) {
		var ret string
		return ret
	}
	return *o.Record
}

// GetRecordOk returns a tuple with the Record field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcAllrecords) GetRecordOk() (*string, bool) {
	if o == nil || IsNil(o.Record) {
		return nil, false
	}
	return o.Record, true
}

// HasRecord returns a boolean if a field has been set.
func (o *DtcAllrecords) HasRecord() bool {
	if o != nil && !IsNil(o.Record) {
		return true
	}

	return false
}

// SetRecord gets a reference to the given string and assigns it to the Record field.
func (o *DtcAllrecords) SetRecord(v string) {
	o.Record = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *DtcAllrecords) GetTtl() int64 {
	if o == nil || IsNil(o.Ttl) {
		var ret int64
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcAllrecords) GetTtlOk() (*int64, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *DtcAllrecords) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int64 and assigns it to the Ttl field.
func (o *DtcAllrecords) SetTtl(v int64) {
	o.Ttl = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DtcAllrecords) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DtcAllrecords) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DtcAllrecords) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DtcAllrecords) SetType(v string) {
	o.Type = &v
}

func (o DtcAllrecords) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DtcAllrecords) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Disable) {
		toSerialize["disable"] = o.Disable
	}
	if !IsNil(o.DtcServer) {
		toSerialize["dtc_server"] = o.DtcServer
	}
	if !IsNil(o.Record) {
		toSerialize["record"] = o.Record
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDtcAllrecords struct {
	value *DtcAllrecords
	isSet bool
}

func (v NullableDtcAllrecords) Get() *DtcAllrecords {
	return v.value
}

func (v *NullableDtcAllrecords) Set(val *DtcAllrecords) {
	v.value = val
	v.isSet = true
}

func (v NullableDtcAllrecords) IsSet() bool {
	return v.isSet
}

func (v *NullableDtcAllrecords) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDtcAllrecords(val *DtcAllrecords) *NullableDtcAllrecords {
	return &NullableDtcAllrecords{value: val, isSet: true}
}

func (v NullableDtcAllrecords) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDtcAllrecords) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
