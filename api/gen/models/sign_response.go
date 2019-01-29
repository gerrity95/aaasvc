// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SignResponse sign response
// swagger:model SignResponse
type SignResponse struct {

	// An error message indicating why signing failed
	Error string `json:"error,omitempty"`

	// base64 encoded choria:secure:request:1 signed message
	// Format: byte
	SecureRequest strfmt.Base64 `json:"secure_request,omitempty"`
}

// Validate validates this sign response
func (m *SignResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecureRequest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SignResponse) validateSecureRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.SecureRequest) { // not required
		return nil
	}

	// Format "byte" (base64 string) is already validated when unmarshalled

	return nil
}

// MarshalBinary interface implementation
func (m *SignResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SignResponse) UnmarshalBinary(b []byte) error {
	var res SignResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
