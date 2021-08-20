/*
Gatekeeper API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CheckDigitalCurrencyAddressResponse struct for CheckDigitalCurrencyAddressResponse
type CheckDigitalCurrencyAddressResponse struct {
	DigitalCurrencyAddress *string `json:"digital_currency_address,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewCheckDigitalCurrencyAddressResponse instantiates a new CheckDigitalCurrencyAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckDigitalCurrencyAddressResponse() *CheckDigitalCurrencyAddressResponse {
	this := CheckDigitalCurrencyAddressResponse{}
	return &this
}

// NewCheckDigitalCurrencyAddressResponseWithDefaults instantiates a new CheckDigitalCurrencyAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckDigitalCurrencyAddressResponseWithDefaults() *CheckDigitalCurrencyAddressResponse {
	this := CheckDigitalCurrencyAddressResponse{}
	return &this
}

// GetDigitalCurrencyAddress returns the DigitalCurrencyAddress field value if set, zero value otherwise.
func (o *CheckDigitalCurrencyAddressResponse) GetDigitalCurrencyAddress() string {
	if o == nil || o.DigitalCurrencyAddress == nil {
		var ret string
		return ret
	}
	return *o.DigitalCurrencyAddress
}

// GetDigitalCurrencyAddressOk returns a tuple with the DigitalCurrencyAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckDigitalCurrencyAddressResponse) GetDigitalCurrencyAddressOk() (*string, bool) {
	if o == nil || o.DigitalCurrencyAddress == nil {
		return nil, false
	}
	return o.DigitalCurrencyAddress, true
}

// HasDigitalCurrencyAddress returns a boolean if a field has been set.
func (o *CheckDigitalCurrencyAddressResponse) HasDigitalCurrencyAddress() bool {
	if o != nil && o.DigitalCurrencyAddress != nil {
		return true
	}

	return false
}

// SetDigitalCurrencyAddress gets a reference to the given string and assigns it to the DigitalCurrencyAddress field.
func (o *CheckDigitalCurrencyAddressResponse) SetDigitalCurrencyAddress(v string) {
	o.DigitalCurrencyAddress = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CheckDigitalCurrencyAddressResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckDigitalCurrencyAddressResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CheckDigitalCurrencyAddressResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CheckDigitalCurrencyAddressResponse) SetStatus(v string) {
	o.Status = &v
}

func (o CheckDigitalCurrencyAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DigitalCurrencyAddress != nil {
		toSerialize["digital_currency_address"] = o.DigitalCurrencyAddress
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableCheckDigitalCurrencyAddressResponse struct {
	value *CheckDigitalCurrencyAddressResponse
	isSet bool
}

func (v NullableCheckDigitalCurrencyAddressResponse) Get() *CheckDigitalCurrencyAddressResponse {
	return v.value
}

func (v *NullableCheckDigitalCurrencyAddressResponse) Set(val *CheckDigitalCurrencyAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckDigitalCurrencyAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckDigitalCurrencyAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckDigitalCurrencyAddressResponse(val *CheckDigitalCurrencyAddressResponse) *NullableCheckDigitalCurrencyAddressResponse {
	return &NullableCheckDigitalCurrencyAddressResponse{value: val, isSet: true}
}

func (v NullableCheckDigitalCurrencyAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckDigitalCurrencyAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


