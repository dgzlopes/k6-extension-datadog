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

// LogsLookupProcessor Use the Lookup Processor to define a mapping between a log attribute and a human readable value saved in the processors mapping table. For example, you can use the Lookup Processor to map an internal service ID into a human readable service name. Alternatively, you could also use it to check if the MAC address that just attempted to connect to the production environment belongs to your list of stolen machines.
type LogsLookupProcessor struct {
	// Value to set the target attribute if the source value is not found in the list.
	DefaultLookup *string `json:"default_lookup,omitempty"`
	// Whether or not the processor is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Mapping table of values for the source attribute and their associated target attribute values, formatted as `[\"source_key1,target_value1\", \"source_key2,target_value2\"]`
	LookupTable []string `json:"lookup_table"`
	// Name of the processor.
	Name *string `json:"name,omitempty"`
	// Source attribute used to perform the lookup.
	Source string `json:"source"`
	// Name of the attribute that contains the corresponding value in the mapping list or the `default_lookup` if not found in the mapping list.
	Target string                  `json:"target"`
	Type   LogsLookupProcessorType `json:"type"`
}

// NewLogsLookupProcessor instantiates a new LogsLookupProcessor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsLookupProcessor(lookupTable []string, source string, target string, type_ LogsLookupProcessorType) *LogsLookupProcessor {
	this := LogsLookupProcessor{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	this.LookupTable = lookupTable
	this.Source = source
	this.Target = target
	this.Type = type_
	return &this
}

// NewLogsLookupProcessorWithDefaults instantiates a new LogsLookupProcessor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsLookupProcessorWithDefaults() *LogsLookupProcessor {
	this := LogsLookupProcessor{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	var type_ LogsLookupProcessorType = "lookup-processor"
	this.Type = type_
	return &this
}

// GetDefaultLookup returns the DefaultLookup field value if set, zero value otherwise.
func (o *LogsLookupProcessor) GetDefaultLookup() string {
	if o == nil || o.DefaultLookup == nil {
		var ret string
		return ret
	}
	return *o.DefaultLookup
}

// GetDefaultLookupOk returns a tuple with the DefaultLookup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsLookupProcessor) GetDefaultLookupOk() (*string, bool) {
	if o == nil || o.DefaultLookup == nil {
		return nil, false
	}
	return o.DefaultLookup, true
}

// HasDefaultLookup returns a boolean if a field has been set.
func (o *LogsLookupProcessor) HasDefaultLookup() bool {
	if o != nil && o.DefaultLookup != nil {
		return true
	}

	return false
}

// SetDefaultLookup gets a reference to the given string and assigns it to the DefaultLookup field.
func (o *LogsLookupProcessor) SetDefaultLookup(v string) {
	o.DefaultLookup = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *LogsLookupProcessor) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsLookupProcessor) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *LogsLookupProcessor) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *LogsLookupProcessor) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetLookupTable returns the LookupTable field value
func (o *LogsLookupProcessor) GetLookupTable() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LookupTable
}

// GetLookupTableOk returns a tuple with the LookupTable field value
// and a boolean to check if the value has been set.
func (o *LogsLookupProcessor) GetLookupTableOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LookupTable, true
}

// SetLookupTable sets field value
func (o *LogsLookupProcessor) SetLookupTable(v []string) {
	o.LookupTable = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogsLookupProcessor) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsLookupProcessor) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogsLookupProcessor) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogsLookupProcessor) SetName(v string) {
	o.Name = &v
}

// GetSource returns the Source field value
func (o *LogsLookupProcessor) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *LogsLookupProcessor) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *LogsLookupProcessor) SetSource(v string) {
	o.Source = v
}

// GetTarget returns the Target field value
func (o *LogsLookupProcessor) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *LogsLookupProcessor) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *LogsLookupProcessor) SetTarget(v string) {
	o.Target = v
}

// GetType returns the Type field value
func (o *LogsLookupProcessor) GetType() LogsLookupProcessorType {
	if o == nil {
		var ret LogsLookupProcessorType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsLookupProcessor) GetTypeOk() (*LogsLookupProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LogsLookupProcessor) SetType(v LogsLookupProcessorType) {
	o.Type = v
}

func (o LogsLookupProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultLookup != nil {
		toSerialize["default_lookup"] = o.DefaultLookup
	}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if true {
		toSerialize["lookup_table"] = o.LookupTable
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["target"] = o.Target
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableLogsLookupProcessor struct {
	value *LogsLookupProcessor
	isSet bool
}

func (v NullableLogsLookupProcessor) Get() *LogsLookupProcessor {
	return v.value
}

func (v *NullableLogsLookupProcessor) Set(val *LogsLookupProcessor) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsLookupProcessor) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsLookupProcessor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsLookupProcessor(val *LogsLookupProcessor) *NullableLogsLookupProcessor {
	return &NullableLogsLookupProcessor{value: val, isSet: true}
}

func (v NullableLogsLookupProcessor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsLookupProcessor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
