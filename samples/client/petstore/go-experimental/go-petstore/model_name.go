/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"encoding/json"
)

// Name Model for testing model name same as property name
type Name struct {
	Name int32 `json:"name"`
	SnakeCase *int32 `json:"snake_case,omitempty"`
	Property *string `json:"property,omitempty"`
	Var123Number *int32 `json:"123Number,omitempty"`
}

// NewName instantiates a new Name object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewName(name int32, ) *Name {
	this := Name{}
	this.Name = name
	return &this
}

// NewNameWithDefaults instantiates a new Name object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameWithDefaults() *Name {
	this := Name{}
	return &this
}

// GetName returns the Name field value
func (o *Name) GetName() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Name) GetNameOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Name) SetName(v int32) {
	o.Name = v
}

// GetSnakeCase returns the SnakeCase field value if set, zero value otherwise.
func (o *Name) GetSnakeCase() int32 {
	if o == nil || o.SnakeCase == nil {
		var ret int32
		return ret
	}
	return *o.SnakeCase
}

// GetSnakeCaseOk returns a tuple with the SnakeCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Name) GetSnakeCaseOk() (*int32, bool) {
	if o == nil || o.SnakeCase == nil {
		return nil, false
	}
	return o.SnakeCase, true
}

// HasSnakeCase returns a boolean if a field has been set.
func (o *Name) HasSnakeCase() bool {
	if o != nil && o.SnakeCase != nil {
		return true
	}

	return false
}

// SetSnakeCase gets a reference to the given int32 and assigns it to the SnakeCase field.
func (o *Name) SetSnakeCase(v int32) {
	o.SnakeCase = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *Name) GetProperty() string {
	if o == nil || o.Property == nil {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Name) GetPropertyOk() (*string, bool) {
	if o == nil || o.Property == nil {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *Name) HasProperty() bool {
	if o != nil && o.Property != nil {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *Name) SetProperty(v string) {
	o.Property = &v
}

// GetVar123Number returns the Var123Number field value if set, zero value otherwise.
func (o *Name) GetVar123Number() int32 {
	if o == nil || o.Var123Number == nil {
		var ret int32
		return ret
	}
	return *o.Var123Number
}

// GetVar123NumberOk returns a tuple with the Var123Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Name) GetVar123NumberOk() (*int32, bool) {
	if o == nil || o.Var123Number == nil {
		return nil, false
	}
	return o.Var123Number, true
}

// HasVar123Number returns a boolean if a field has been set.
func (o *Name) HasVar123Number() bool {
	if o != nil && o.Var123Number != nil {
		return true
	}

	return false
}

// SetVar123Number gets a reference to the given int32 and assigns it to the Var123Number field.
func (o *Name) SetVar123Number(v int32) {
	o.Var123Number = &v
}

func (o Name) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.SnakeCase != nil {
		toSerialize["snake_case"] = o.SnakeCase
	}
	if o.Property != nil {
		toSerialize["property"] = o.Property
	}
	if o.Var123Number != nil {
		toSerialize["123Number"] = o.Var123Number
	}
	return json.Marshal(toSerialize)
}

type NullableName struct {
	value *Name
	isSet bool
}

func (v NullableName) Get() *Name {
	return v.value
}

func (v *NullableName) Set(val *Name) {
	v.value = val
	v.isSet = true
}

func (v NullableName) IsSet() bool {
	return v.isSet
}

func (v *NullableName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableName(val *Name) *NullableName {
	return &NullableName{value: val, isSet: true}
}

func (v NullableName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


