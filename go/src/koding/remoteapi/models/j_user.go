package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// JUser j user
// swagger:model JUser
type JUser struct {

	// blocked reason
	BlockedReason string `json:"blockedReason,omitempty"`

	// blocked until
	BlockedUntil strfmt.Date `json:"blockedUntil,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// email frequency
	EmailFrequency interface{} `json:"emailFrequency,omitempty"`

	// last login date
	LastLoginDate strfmt.Date `json:"lastLoginDate,omitempty"`

	// old username
	OldUsername string `json:"oldUsername,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// password status
	PasswordStatus string `json:"passwordStatus,omitempty"`

	// registered at
	RegisteredAt strfmt.Date `json:"registeredAt,omitempty"`

	// registered from
	RegisteredFrom *JUserRegisteredFrom `json:"registeredFrom,omitempty"`

	// salt
	Salt string `json:"salt,omitempty"`

	// sanitized email
	SanitizedEmail string `json:"sanitizedEmail,omitempty"`

	// ssh keys
	SSHKeys interface{} `json:"sshKeys,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// twofactorkey
	Twofactorkey string `json:"twofactorkey,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this j user
func (m *JUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegisteredFrom(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JUser) validateRegisteredFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.RegisteredFrom) { // not required
		return nil
	}

	if m.RegisteredFrom != nil {

		if err := m.RegisteredFrom.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// JUserRegisteredFrom j user registered from
// swagger:model JUserRegisteredFrom
type JUserRegisteredFrom struct {

	// country
	Country string `json:"country,omitempty"`

	// ip
	IP string `json:"ip,omitempty"`

	// region
	Region string `json:"region,omitempty"`
}

// Validate validates this j user registered from
func (m *JUserRegisteredFrom) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
