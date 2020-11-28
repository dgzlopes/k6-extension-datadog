/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// ProcessQueryDefinition The process query to use in the widget.
type ProcessQueryDefinition struct {
	// List of processes.
	FilterBy *[]string `json:"filter_by,omitempty"`
	// Max number of items in the filter list.
	Limit *int64 `json:"limit,omitempty"`
	// Your chosen metric.
	Metric string `json:"metric"`
	// Your chosen search term.
	SearchBy *string `json:"search_by,omitempty"`
}

// NewProcessQueryDefinition instantiates a new ProcessQueryDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessQueryDefinition(metric string) *ProcessQueryDefinition {
	this := ProcessQueryDefinition{}
	this.Metric = metric
	return &this
}

// NewProcessQueryDefinitionWithDefaults instantiates a new ProcessQueryDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessQueryDefinitionWithDefaults() *ProcessQueryDefinition {
	this := ProcessQueryDefinition{}
	return &this
}

// GetFilterBy returns the FilterBy field value if set, zero value otherwise.
func (o *ProcessQueryDefinition) GetFilterBy() []string {
	if o == nil || o.FilterBy == nil {
		var ret []string
		return ret
	}
	return *o.FilterBy
}

// GetFilterByOk returns a tuple with the FilterBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessQueryDefinition) GetFilterByOk() (*[]string, bool) {
	if o == nil || o.FilterBy == nil {
		return nil, false
	}
	return o.FilterBy, true
}

// HasFilterBy returns a boolean if a field has been set.
func (o *ProcessQueryDefinition) HasFilterBy() bool {
	if o != nil && o.FilterBy != nil {
		return true
	}

	return false
}

// SetFilterBy gets a reference to the given []string and assigns it to the FilterBy field.
func (o *ProcessQueryDefinition) SetFilterBy(v []string) {
	o.FilterBy = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ProcessQueryDefinition) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessQueryDefinition) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ProcessQueryDefinition) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ProcessQueryDefinition) SetLimit(v int64) {
	o.Limit = &v
}

// GetMetric returns the Metric field value
func (o *ProcessQueryDefinition) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *ProcessQueryDefinition) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *ProcessQueryDefinition) SetMetric(v string) {
	o.Metric = v
}

// GetSearchBy returns the SearchBy field value if set, zero value otherwise.
func (o *ProcessQueryDefinition) GetSearchBy() string {
	if o == nil || o.SearchBy == nil {
		var ret string
		return ret
	}
	return *o.SearchBy
}

// GetSearchByOk returns a tuple with the SearchBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessQueryDefinition) GetSearchByOk() (*string, bool) {
	if o == nil || o.SearchBy == nil {
		return nil, false
	}
	return o.SearchBy, true
}

// HasSearchBy returns a boolean if a field has been set.
func (o *ProcessQueryDefinition) HasSearchBy() bool {
	if o != nil && o.SearchBy != nil {
		return true
	}

	return false
}

// SetSearchBy gets a reference to the given string and assigns it to the SearchBy field.
func (o *ProcessQueryDefinition) SetSearchBy(v string) {
	o.SearchBy = &v
}

func (o ProcessQueryDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FilterBy != nil {
		toSerialize["filter_by"] = o.FilterBy
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["metric"] = o.Metric
	}
	if o.SearchBy != nil {
		toSerialize["search_by"] = o.SearchBy
	}
	return json.Marshal(toSerialize)
}

type NullableProcessQueryDefinition struct {
	value *ProcessQueryDefinition
	isSet bool
}

func (v NullableProcessQueryDefinition) Get() *ProcessQueryDefinition {
	return v.value
}

func (v *NullableProcessQueryDefinition) Set(val *ProcessQueryDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessQueryDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessQueryDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessQueryDefinition(val *ProcessQueryDefinition) *NullableProcessQueryDefinition {
	return &NullableProcessQueryDefinition{value: val, isSet: true}
}

func (v NullableProcessQueryDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessQueryDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
