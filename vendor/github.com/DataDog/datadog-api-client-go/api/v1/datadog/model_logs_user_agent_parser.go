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

// LogsUserAgentParser The User-Agent parser takes a User-Agent attribute and extracts the OS, browser, device, and other user data. It recognizes major bots like the Google Bot, Yahoo Slurp, and Bing.
type LogsUserAgentParser struct {
	// Whether or not the processor is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// Define if the source attribute is URL encoded or not.
	IsEncoded *bool `json:"is_encoded,omitempty"`
	// Name of the processor.
	Name *string `json:"name,omitempty"`
	// Array of source attributes.
	Sources []string `json:"sources"`
	// Name of the parent attribute that contains all the extracted details from the `sources`.
	Target string                  `json:"target"`
	Type   LogsUserAgentParserType `json:"type"`
}

// NewLogsUserAgentParser instantiates a new LogsUserAgentParser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsUserAgentParser(sources []string, target string, type_ LogsUserAgentParserType) *LogsUserAgentParser {
	this := LogsUserAgentParser{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	var isEncoded bool = false
	this.IsEncoded = &isEncoded
	this.Sources = sources
	this.Target = target
	this.Type = type_
	return &this
}

// NewLogsUserAgentParserWithDefaults instantiates a new LogsUserAgentParser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsUserAgentParserWithDefaults() *LogsUserAgentParser {
	this := LogsUserAgentParser{}
	var isEnabled bool = false
	this.IsEnabled = &isEnabled
	var isEncoded bool = false
	this.IsEncoded = &isEncoded
	var target string = "http.useragent_details"
	this.Target = target
	var type_ LogsUserAgentParserType = "user-agent-parser"
	this.Type = type_
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *LogsUserAgentParser) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsUserAgentParser) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *LogsUserAgentParser) HasIsEnabled() bool {
	if o != nil && o.IsEnabled != nil {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *LogsUserAgentParser) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetIsEncoded returns the IsEncoded field value if set, zero value otherwise.
func (o *LogsUserAgentParser) GetIsEncoded() bool {
	if o == nil || o.IsEncoded == nil {
		var ret bool
		return ret
	}
	return *o.IsEncoded
}

// GetIsEncodedOk returns a tuple with the IsEncoded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsUserAgentParser) GetIsEncodedOk() (*bool, bool) {
	if o == nil || o.IsEncoded == nil {
		return nil, false
	}
	return o.IsEncoded, true
}

// HasIsEncoded returns a boolean if a field has been set.
func (o *LogsUserAgentParser) HasIsEncoded() bool {
	if o != nil && o.IsEncoded != nil {
		return true
	}

	return false
}

// SetIsEncoded gets a reference to the given bool and assigns it to the IsEncoded field.
func (o *LogsUserAgentParser) SetIsEncoded(v bool) {
	o.IsEncoded = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogsUserAgentParser) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsUserAgentParser) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogsUserAgentParser) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogsUserAgentParser) SetName(v string) {
	o.Name = &v
}

// GetSources returns the Sources field value
func (o *LogsUserAgentParser) GetSources() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *LogsUserAgentParser) GetSourcesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value
func (o *LogsUserAgentParser) SetSources(v []string) {
	o.Sources = v
}

// GetTarget returns the Target field value
func (o *LogsUserAgentParser) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *LogsUserAgentParser) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *LogsUserAgentParser) SetTarget(v string) {
	o.Target = v
}

// GetType returns the Type field value
func (o *LogsUserAgentParser) GetType() LogsUserAgentParserType {
	if o == nil {
		var ret LogsUserAgentParserType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsUserAgentParser) GetTypeOk() (*LogsUserAgentParserType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LogsUserAgentParser) SetType(v LogsUserAgentParserType) {
	o.Type = v
}

func (o LogsUserAgentParser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.IsEncoded != nil {
		toSerialize["is_encoded"] = o.IsEncoded
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["sources"] = o.Sources
	}
	if true {
		toSerialize["target"] = o.Target
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableLogsUserAgentParser struct {
	value *LogsUserAgentParser
	isSet bool
}

func (v NullableLogsUserAgentParser) Get() *LogsUserAgentParser {
	return v.value
}

func (v *NullableLogsUserAgentParser) Set(val *LogsUserAgentParser) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsUserAgentParser) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsUserAgentParser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsUserAgentParser(val *LogsUserAgentParser) *NullableLogsUserAgentParser {
	return &NullableLogsUserAgentParser{value: val, isSet: true}
}

func (v NullableLogsUserAgentParser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsUserAgentParser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
