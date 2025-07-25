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

// checks if the GridDashboard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GridDashboard{}

// GridDashboard struct for GridDashboard
type GridDashboard struct {
	// The reference to the object.
	Ref *string `json:"_ref,omitempty"`
	// The Grid Dashboard critical threshold for Analytics tunneling events.
	AnalyticsTunnelingEventCriticalThreshold *int64 `json:"analytics_tunneling_event_critical_threshold,omitempty"`
	// The Grid Dashboard warning threshold for Analytics tunneling events.
	AnalyticsTunnelingEventWarningThreshold *int64 `json:"analytics_tunneling_event_warning_threshold,omitempty"`
	// The Grid Dashboard critical threshold for ATP critical events.
	AtpCriticalEventCriticalThreshold *int64 `json:"atp_critical_event_critical_threshold,omitempty"`
	// The Grid Dashboard warning threshold for ATP critical events.
	AtpCriticalEventWarningThreshold *int64 `json:"atp_critical_event_warning_threshold,omitempty"`
	// The Grid Dashboard critical threshold for ATP major events.
	AtpMajorEventCriticalThreshold *int64 `json:"atp_major_event_critical_threshold,omitempty"`
	// The Grid Dashboard warning threshold for ATP major events.
	AtpMajorEventWarningThreshold *int64 `json:"atp_major_event_warning_threshold,omitempty"`
	// The Grid Dashboard critical threshold for ATP warning events.
	AtpWarningEventCriticalThreshold *int64 `json:"atp_warning_event_critical_threshold,omitempty"`
	// The Grid Dashboard warning threshold for ATP warning events.
	AtpWarningEventWarningThreshold *int64 `json:"atp_warning_event_warning_threshold,omitempty"`
	// The critical threshold value for blocked RPZ hits in the Grid dashboard.
	RpzBlockedHitCriticalThreshold *int64 `json:"rpz_blocked_hit_critical_threshold,omitempty"`
	// The warning threshold value for blocked RPZ hits in the Grid dashboard.
	RpzBlockedHitWarningThreshold *int64 `json:"rpz_blocked_hit_warning_threshold,omitempty"`
	// The Grid Dashboard critical threshold for RPZ passthru events.
	RpzPassthruEventCriticalThreshold *int64 `json:"rpz_passthru_event_critical_threshold,omitempty"`
	// The Grid Dashboard warning threshold for RPZ passthru events.
	RpzPassthruEventWarningThreshold *int64 `json:"rpz_passthru_event_warning_threshold,omitempty"`
	// The critical threshold value for substituted RPZ hits in the Grid dashboard.
	RpzSubstitutedHitCriticalThreshold *int64 `json:"rpz_substituted_hit_critical_threshold,omitempty"`
	// The warning threshold value for substituted RPZ hits in the Grid dashboard.
	RpzSubstitutedHitWarningThreshold *int64 `json:"rpz_substituted_hit_warning_threshold,omitempty"`
}

// NewGridDashboard instantiates a new GridDashboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGridDashboard() *GridDashboard {
	this := GridDashboard{}
	return &this
}

// NewGridDashboardWithDefaults instantiates a new GridDashboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridDashboardWithDefaults() *GridDashboard {
	this := GridDashboard{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *GridDashboard) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDashboard) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *GridDashboard) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *GridDashboard) SetRef(v string) {
	o.Ref = &v
}

// GetAnalyticsTunnelingEventCriticalThreshold returns the AnalyticsTunnelingEventCriticalThreshold field value if set, zero value otherwise.
func (o *GridDashboard) GetAnalyticsTunnelingEventCriticalThreshold() int64 {
	if o == nil || IsNil(o.AnalyticsTunnelingEventCriticalThreshold) {
		var ret int64
		return ret
	}
	return *o.AnalyticsTunnelingEventCriticalThreshold
}

// GetAnalyticsTunnelingEventCriticalThresholdOk returns a tuple with the AnalyticsTunnelingEventCriticalThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDashboard) GetAnalyticsTunnelingEventCriticalThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.AnalyticsTunnelingEventCriticalThreshold) {
		return nil, false
	}
	return o.AnalyticsTunnelingEventCriticalThreshold, true
}

// HasAnalyticsTunnelingEventCriticalThreshold returns a boolean if a field has been set.
func (o *GridDashboard) HasAnalyticsTunnelingEventCriticalThreshold() bool {
	if o != nil && !IsNil(o.AnalyticsTunnelingEventCriticalThreshold) {
		return true
	}

	return false
}

// SetAnalyticsTunnelingEventCriticalThreshold gets a reference to the given int64 and assigns it to the AnalyticsTunnelingEventCriticalThreshold field.
func (o *GridDashboard) SetAnalyticsTunnelingEventCriticalThreshold(v int64) {
	o.AnalyticsTunnelingEventCriticalThreshold = &v
}

// GetAnalyticsTunnelingEventWarningThreshold returns the AnalyticsTunnelingEventWarningThreshold field value if set, zero value otherwise.
func (o *GridDashboard) GetAnalyticsTunnelingEventWarningThreshold() int64 {
	if o == nil || IsNil(o.AnalyticsTunnelingEventWarningThreshold) {
		var ret int64
		return ret
	}
	return *o.AnalyticsTunnelingEventWarningThreshold
}

// GetAnalyticsTunnelingEventWarningThresholdOk returns a tuple with the AnalyticsTunnelingEventWarningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDashboard) GetAnalyticsTunnelingEventWarningThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.AnalyticsTunnelingEventWarningThreshold) {
		return nil, false
	}
	return o.AnalyticsTunnelingEventWarningThreshold, true
}

// HasAnalyticsTunnelingEventWarningThreshold returns a boolean if a field has been set.
func (o *GridDashboard) HasAnalyticsTunnelingEventWarningThreshold() bool {
	if o != nil && !IsNil(o.AnalyticsTunnelingEventWarningThreshold) {
		return true
	}

	return false
}

// SetAnalyticsTunnelingEventWarningThreshold gets a reference to the given int64 and assigns it to the AnalyticsTunnelingEventWarningThreshold field.
func (o *GridDashboard) SetAnalyticsTunnelingEventWarningThreshold(v int64) {
	o.AnalyticsTunnelingEventWarningThreshold = &v
}

// GetAtpCriticalEventCriticalThreshold returns the AtpCriticalEventCriticalThreshold field value if set, zero value otherwise.
func (o *GridDashboard) GetAtpCriticalEventCriticalThreshold() int64 {
	if o == nil || IsNil(o.AtpCriticalEventCriticalThreshold) {
		var ret int64
		return ret
	}
	return *o.AtpCriticalEventCriticalThreshold
}

// GetAtpCriticalEventCriticalThresholdOk returns a tuple with the AtpCriticalEventCriticalThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDashboard) GetAtpCriticalEventCriticalThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.AtpCriticalEventCriticalThreshold) {
		return nil, false
	}
	return o.AtpCriticalEventCriticalThreshold, true
}

// HasAtpCriticalEventCriticalThreshold returns a boolean if a field has been set.
func (o *GridDashboard) HasAtpCriticalEventCriticalThreshold() bool {
	if o != nil && !IsNil(o.AtpCriticalEventCriticalThreshold) {
		return true
	}

	return false
}

// SetAtpCriticalEventCriticalThreshold gets a reference to the given int64 and assigns it to the AtpCriticalEventCriticalThreshold field.
func (o *GridDashboard) SetAtpCriticalEventCriticalThreshold(v int64) {
	o.AtpCriticalEventCriticalThreshold = &v
}

// GetAtpCriticalEventWarningThreshold returns the AtpCriticalEventWarningThreshold field value if set, zero value otherwise.
func (o *GridDashboard) GetAtpCriticalEventWarningThreshold() int64 {
	if o == nil || IsNil(o.AtpCriticalEventWarningThreshold) {
		var ret int64
		return ret
	}
	return *o.AtpCriticalEventWarningThreshold
}

// GetAtpCriticalEventWarningThresholdOk returns a tuple with the AtpCriticalEventWarningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDashboard) GetAtpCriticalEventWarningThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.AtpCriticalEventWarningThreshold) {
		return nil, false
	}
	return o.AtpCriticalEventWarningThreshold, true
}

// HasAtpCriticalEventWarningThreshold returns a boolean if a field has been set.
func (o *GridDashboard) HasAtpCriticalEventWarningThreshold() bool {
	if o != nil && !IsNil(o.AtpCriticalEventWarningThreshold) {
		return true
	}

	return false
}

// SetAtpCriticalEventWarningThreshold gets a reference to the given int64 and assigns it to the AtpCriticalEventWarningThreshold field.
func (o *GridDashboard) SetAtpCriticalEventWarningThreshold(v int64) {
	o.AtpCriticalEventWarningThreshold = &v
}

// GetAtpMajorEventCriticalThreshold returns the AtpMajorEventCriticalThreshold field value if set, zero value otherwise.
func (o *GridDashboard) GetAtpMajorEventCriticalThreshold() int64 {
	if o == nil || IsNil(o.AtpMajorEventCriticalThreshold) {
		var ret int64
		return ret
	}
	return *o.AtpMajorEventCriticalThreshold
}

// GetAtpMajorEventCriticalThresholdOk returns a tuple with the AtpMajorEventCriticalThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDashboard) GetAtpMajorEventCriticalThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.AtpMajorEventCriticalThreshold) {
		return nil, false
	}
	return o.AtpMajorEventCriticalThreshold, true
}

// HasAtpMajorEventCriticalThreshold returns a boolean if a field has been set.
func (o *GridDashboard) HasAtpMajorEventCriticalThreshold() bool {
	if o != nil && !IsNil(o.AtpMajorEventCriticalThreshold) {
		return true
	}

	return false
}

// SetAtpMajorEventCriticalThreshold gets a reference to the given int64 and assigns it to the AtpMajorEventCriticalThreshold field.
func (o *GridDashboard) SetAtpMajorEventCriticalThreshold(v int64) {
	o.AtpMajorEventCriticalThreshold = &v
}

// GetAtpMajorEventWarningThreshold returns the AtpMajorEventWarningThreshold field value if set, zero value otherwise.
func (o *GridDashboard) GetAtpMajorEventWarningThreshold() int64 {
	if o == nil || IsNil(o.AtpMajorEventWarningThreshold) {
		var ret int64
		return ret
	}
	return *o.AtpMajorEventWarningThreshold
}

// GetAtpMajorEventWarningThresholdOk returns a tuple with the AtpMajorEventWarningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDashboard) GetAtpMajorEventWarningThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.AtpMajorEventWarningThreshold) {
		return nil, false
	}
	return o.AtpMajorEventWarningThreshold, true
}

// HasAtpMajorEventWarningThreshold returns a boolean if a field has been set.
func (o *GridDashboard) HasAtpMajorEventWarningThreshold() bool {
	if o != nil && !IsNil(o.AtpMajorEventWarningThreshold) {
		return true
	}

	return false
}

// SetAtpMajorEventWarningThreshold gets a reference to the given int64 and assigns it to the AtpMajorEventWarningThreshold field.
func (o *GridDashboard) SetAtpMajorEventWarningThreshold(v int64) {
	o.AtpMajorEventWarningThreshold = &v
}

// GetAtpWarningEventCriticalThreshold returns the AtpWarningEventCriticalThreshold field value if set, zero value otherwise.
func (o *GridDashboard) GetAtpWarningEventCriticalThreshold() int64 {
	if o == nil || IsNil(o.AtpWarningEventCriticalThreshold) {
		var ret int64
		return ret
	}
	return *o.AtpWarningEventCriticalThreshold
}

// GetAtpWarningEventCriticalThresholdOk returns a tuple with the AtpWarningEventCriticalThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDashboard) GetAtpWarningEventCriticalThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.AtpWarningEventCriticalThreshold) {
		return nil, false
	}
	return o.AtpWarningEventCriticalThreshold, true
}

// HasAtpWarningEventCriticalThreshold returns a boolean if a field has been set.
func (o *GridDashboard) HasAtpWarningEventCriticalThreshold() bool {
	if o != nil && !IsNil(o.AtpWarningEventCriticalThreshold) {
		return true
	}

	return false
}

// SetAtpWarningEventCriticalThreshold gets a reference to the given int64 and assigns it to the AtpWarningEventCriticalThreshold field.
func (o *GridDashboard) SetAtpWarningEventCriticalThreshold(v int64) {
	o.AtpWarningEventCriticalThreshold = &v
}

// GetAtpWarningEventWarningThreshold returns the AtpWarningEventWarningThreshold field value if set, zero value otherwise.
func (o *GridDashboard) GetAtpWarningEventWarningThreshold() int64 {
	if o == nil || IsNil(o.AtpWarningEventWarningThreshold) {
		var ret int64
		return ret
	}
	return *o.AtpWarningEventWarningThreshold
}

// GetAtpWarningEventWarningThresholdOk returns a tuple with the AtpWarningEventWarningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDashboard) GetAtpWarningEventWarningThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.AtpWarningEventWarningThreshold) {
		return nil, false
	}
	return o.AtpWarningEventWarningThreshold, true
}

// HasAtpWarningEventWarningThreshold returns a boolean if a field has been set.
func (o *GridDashboard) HasAtpWarningEventWarningThreshold() bool {
	if o != nil && !IsNil(o.AtpWarningEventWarningThreshold) {
		return true
	}

	return false
}

// SetAtpWarningEventWarningThreshold gets a reference to the given int64 and assigns it to the AtpWarningEventWarningThreshold field.
func (o *GridDashboard) SetAtpWarningEventWarningThreshold(v int64) {
	o.AtpWarningEventWarningThreshold = &v
}

// GetRpzBlockedHitCriticalThreshold returns the RpzBlockedHitCriticalThreshold field value if set, zero value otherwise.
func (o *GridDashboard) GetRpzBlockedHitCriticalThreshold() int64 {
	if o == nil || IsNil(o.RpzBlockedHitCriticalThreshold) {
		var ret int64
		return ret
	}
	return *o.RpzBlockedHitCriticalThreshold
}

// GetRpzBlockedHitCriticalThresholdOk returns a tuple with the RpzBlockedHitCriticalThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDashboard) GetRpzBlockedHitCriticalThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.RpzBlockedHitCriticalThreshold) {
		return nil, false
	}
	return o.RpzBlockedHitCriticalThreshold, true
}

// HasRpzBlockedHitCriticalThreshold returns a boolean if a field has been set.
func (o *GridDashboard) HasRpzBlockedHitCriticalThreshold() bool {
	if o != nil && !IsNil(o.RpzBlockedHitCriticalThreshold) {
		return true
	}

	return false
}

// SetRpzBlockedHitCriticalThreshold gets a reference to the given int64 and assigns it to the RpzBlockedHitCriticalThreshold field.
func (o *GridDashboard) SetRpzBlockedHitCriticalThreshold(v int64) {
	o.RpzBlockedHitCriticalThreshold = &v
}

// GetRpzBlockedHitWarningThreshold returns the RpzBlockedHitWarningThreshold field value if set, zero value otherwise.
func (o *GridDashboard) GetRpzBlockedHitWarningThreshold() int64 {
	if o == nil || IsNil(o.RpzBlockedHitWarningThreshold) {
		var ret int64
		return ret
	}
	return *o.RpzBlockedHitWarningThreshold
}

// GetRpzBlockedHitWarningThresholdOk returns a tuple with the RpzBlockedHitWarningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDashboard) GetRpzBlockedHitWarningThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.RpzBlockedHitWarningThreshold) {
		return nil, false
	}
	return o.RpzBlockedHitWarningThreshold, true
}

// HasRpzBlockedHitWarningThreshold returns a boolean if a field has been set.
func (o *GridDashboard) HasRpzBlockedHitWarningThreshold() bool {
	if o != nil && !IsNil(o.RpzBlockedHitWarningThreshold) {
		return true
	}

	return false
}

// SetRpzBlockedHitWarningThreshold gets a reference to the given int64 and assigns it to the RpzBlockedHitWarningThreshold field.
func (o *GridDashboard) SetRpzBlockedHitWarningThreshold(v int64) {
	o.RpzBlockedHitWarningThreshold = &v
}

// GetRpzPassthruEventCriticalThreshold returns the RpzPassthruEventCriticalThreshold field value if set, zero value otherwise.
func (o *GridDashboard) GetRpzPassthruEventCriticalThreshold() int64 {
	if o == nil || IsNil(o.RpzPassthruEventCriticalThreshold) {
		var ret int64
		return ret
	}
	return *o.RpzPassthruEventCriticalThreshold
}

// GetRpzPassthruEventCriticalThresholdOk returns a tuple with the RpzPassthruEventCriticalThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDashboard) GetRpzPassthruEventCriticalThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.RpzPassthruEventCriticalThreshold) {
		return nil, false
	}
	return o.RpzPassthruEventCriticalThreshold, true
}

// HasRpzPassthruEventCriticalThreshold returns a boolean if a field has been set.
func (o *GridDashboard) HasRpzPassthruEventCriticalThreshold() bool {
	if o != nil && !IsNil(o.RpzPassthruEventCriticalThreshold) {
		return true
	}

	return false
}

// SetRpzPassthruEventCriticalThreshold gets a reference to the given int64 and assigns it to the RpzPassthruEventCriticalThreshold field.
func (o *GridDashboard) SetRpzPassthruEventCriticalThreshold(v int64) {
	o.RpzPassthruEventCriticalThreshold = &v
}

// GetRpzPassthruEventWarningThreshold returns the RpzPassthruEventWarningThreshold field value if set, zero value otherwise.
func (o *GridDashboard) GetRpzPassthruEventWarningThreshold() int64 {
	if o == nil || IsNil(o.RpzPassthruEventWarningThreshold) {
		var ret int64
		return ret
	}
	return *o.RpzPassthruEventWarningThreshold
}

// GetRpzPassthruEventWarningThresholdOk returns a tuple with the RpzPassthruEventWarningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDashboard) GetRpzPassthruEventWarningThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.RpzPassthruEventWarningThreshold) {
		return nil, false
	}
	return o.RpzPassthruEventWarningThreshold, true
}

// HasRpzPassthruEventWarningThreshold returns a boolean if a field has been set.
func (o *GridDashboard) HasRpzPassthruEventWarningThreshold() bool {
	if o != nil && !IsNil(o.RpzPassthruEventWarningThreshold) {
		return true
	}

	return false
}

// SetRpzPassthruEventWarningThreshold gets a reference to the given int64 and assigns it to the RpzPassthruEventWarningThreshold field.
func (o *GridDashboard) SetRpzPassthruEventWarningThreshold(v int64) {
	o.RpzPassthruEventWarningThreshold = &v
}

// GetRpzSubstitutedHitCriticalThreshold returns the RpzSubstitutedHitCriticalThreshold field value if set, zero value otherwise.
func (o *GridDashboard) GetRpzSubstitutedHitCriticalThreshold() int64 {
	if o == nil || IsNil(o.RpzSubstitutedHitCriticalThreshold) {
		var ret int64
		return ret
	}
	return *o.RpzSubstitutedHitCriticalThreshold
}

// GetRpzSubstitutedHitCriticalThresholdOk returns a tuple with the RpzSubstitutedHitCriticalThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDashboard) GetRpzSubstitutedHitCriticalThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.RpzSubstitutedHitCriticalThreshold) {
		return nil, false
	}
	return o.RpzSubstitutedHitCriticalThreshold, true
}

// HasRpzSubstitutedHitCriticalThreshold returns a boolean if a field has been set.
func (o *GridDashboard) HasRpzSubstitutedHitCriticalThreshold() bool {
	if o != nil && !IsNil(o.RpzSubstitutedHitCriticalThreshold) {
		return true
	}

	return false
}

// SetRpzSubstitutedHitCriticalThreshold gets a reference to the given int64 and assigns it to the RpzSubstitutedHitCriticalThreshold field.
func (o *GridDashboard) SetRpzSubstitutedHitCriticalThreshold(v int64) {
	o.RpzSubstitutedHitCriticalThreshold = &v
}

// GetRpzSubstitutedHitWarningThreshold returns the RpzSubstitutedHitWarningThreshold field value if set, zero value otherwise.
func (o *GridDashboard) GetRpzSubstitutedHitWarningThreshold() int64 {
	if o == nil || IsNil(o.RpzSubstitutedHitWarningThreshold) {
		var ret int64
		return ret
	}
	return *o.RpzSubstitutedHitWarningThreshold
}

// GetRpzSubstitutedHitWarningThresholdOk returns a tuple with the RpzSubstitutedHitWarningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDashboard) GetRpzSubstitutedHitWarningThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.RpzSubstitutedHitWarningThreshold) {
		return nil, false
	}
	return o.RpzSubstitutedHitWarningThreshold, true
}

// HasRpzSubstitutedHitWarningThreshold returns a boolean if a field has been set.
func (o *GridDashboard) HasRpzSubstitutedHitWarningThreshold() bool {
	if o != nil && !IsNil(o.RpzSubstitutedHitWarningThreshold) {
		return true
	}

	return false
}

// SetRpzSubstitutedHitWarningThreshold gets a reference to the given int64 and assigns it to the RpzSubstitutedHitWarningThreshold field.
func (o *GridDashboard) SetRpzSubstitutedHitWarningThreshold(v int64) {
	o.RpzSubstitutedHitWarningThreshold = &v
}

func (o GridDashboard) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GridDashboard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["_ref"] = o.Ref
	}
	if !IsNil(o.AnalyticsTunnelingEventCriticalThreshold) {
		toSerialize["analytics_tunneling_event_critical_threshold"] = o.AnalyticsTunnelingEventCriticalThreshold
	}
	if !IsNil(o.AnalyticsTunnelingEventWarningThreshold) {
		toSerialize["analytics_tunneling_event_warning_threshold"] = o.AnalyticsTunnelingEventWarningThreshold
	}
	if !IsNil(o.AtpCriticalEventCriticalThreshold) {
		toSerialize["atp_critical_event_critical_threshold"] = o.AtpCriticalEventCriticalThreshold
	}
	if !IsNil(o.AtpCriticalEventWarningThreshold) {
		toSerialize["atp_critical_event_warning_threshold"] = o.AtpCriticalEventWarningThreshold
	}
	if !IsNil(o.AtpMajorEventCriticalThreshold) {
		toSerialize["atp_major_event_critical_threshold"] = o.AtpMajorEventCriticalThreshold
	}
	if !IsNil(o.AtpMajorEventWarningThreshold) {
		toSerialize["atp_major_event_warning_threshold"] = o.AtpMajorEventWarningThreshold
	}
	if !IsNil(o.AtpWarningEventCriticalThreshold) {
		toSerialize["atp_warning_event_critical_threshold"] = o.AtpWarningEventCriticalThreshold
	}
	if !IsNil(o.AtpWarningEventWarningThreshold) {
		toSerialize["atp_warning_event_warning_threshold"] = o.AtpWarningEventWarningThreshold
	}
	if !IsNil(o.RpzBlockedHitCriticalThreshold) {
		toSerialize["rpz_blocked_hit_critical_threshold"] = o.RpzBlockedHitCriticalThreshold
	}
	if !IsNil(o.RpzBlockedHitWarningThreshold) {
		toSerialize["rpz_blocked_hit_warning_threshold"] = o.RpzBlockedHitWarningThreshold
	}
	if !IsNil(o.RpzPassthruEventCriticalThreshold) {
		toSerialize["rpz_passthru_event_critical_threshold"] = o.RpzPassthruEventCriticalThreshold
	}
	if !IsNil(o.RpzPassthruEventWarningThreshold) {
		toSerialize["rpz_passthru_event_warning_threshold"] = o.RpzPassthruEventWarningThreshold
	}
	if !IsNil(o.RpzSubstitutedHitCriticalThreshold) {
		toSerialize["rpz_substituted_hit_critical_threshold"] = o.RpzSubstitutedHitCriticalThreshold
	}
	if !IsNil(o.RpzSubstitutedHitWarningThreshold) {
		toSerialize["rpz_substituted_hit_warning_threshold"] = o.RpzSubstitutedHitWarningThreshold
	}
	return toSerialize, nil
}

type NullableGridDashboard struct {
	value *GridDashboard
	isSet bool
}

func (v NullableGridDashboard) Get() *GridDashboard {
	return v.value
}

func (v *NullableGridDashboard) Set(val *GridDashboard) {
	v.value = val
	v.isSet = true
}

func (v NullableGridDashboard) IsSet() bool {
	return v.isSet
}

func (v *NullableGridDashboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridDashboard(val *GridDashboard) *NullableGridDashboard {
	return &NullableGridDashboard{value: val, isSet: true}
}

func (v NullableGridDashboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridDashboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
