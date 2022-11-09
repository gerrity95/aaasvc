// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SignRequest sign request
//
// swagger:model SignRequest
type SignRequest struct {

	// base64 encoded Choria protocol.Request message to sign
	// Format: byte
	Request strfmt.Base64 `json:"request,omitempty"`

	// A signature, hex encoded, made using the private key matching the public key in the token
	Signature string `json:"signature,omitempty"`

	// The JWT token identifying the user, obtained from /login
	Token string `json:"token,omitempty"`
}

// Validate validates this sign request
func (m *SignRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sign request based on context it is used
func (m *SignRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SignRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SignRequest) UnmarshalBinary(b []byte) error {
	var res SignRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
