// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatacenterSpecVSphere DatacenterSpecVSphere describes a vSphere datacenter.
//
// swagger:model DatacenterSpecVSphere
type DatacenterSpecVSphere struct {

	// If set to true, disables the TLS certificate check against the endpoint.
	AllowInsecure bool `json:"allowInsecure,omitempty"`

	// The name of the vSphere cluster to use. Used for out-of-tree CSI Driver.
	Cluster string `json:"cluster,omitempty"`

	// The name of the datacenter to use.
	Datacenter string `json:"datacenter,omitempty"`

	// The default Datastore to be used for provisioning volumes using storage
	// classes/dynamic provisioning and for storing virtual machine files in
	// case no `Datastore` or `DatastoreCluster` is provided at Cluster level.
	DefaultDatastore string `json:"datastore,omitempty"`

	// The name of the storage policy to use for the storage class created in the user cluster.
	DefaultStoragePolicy string `json:"storagePolicy,omitempty"`

	// Endpoint URL to use, including protocol, for example "https://vcenter.example.com".
	Endpoint string `json:"endpoint,omitempty"`

	// Optional: defines if the IPv6 is enabled for the datacenter
	IPV6Enabled bool `json:"ipv6Enabled,omitempty"`

	// Optional: The root path for cluster specific VM folders. Each cluster gets its own
	// folder below the root folder. Must be the FQDN (for example
	// "/datacenter-1/vm/all-kubermatic-vms-in-here") and defaults to the root VM
	// folder: "/datacenter-1/vm"
	RootPath string `json:"rootPath,omitempty"`

	// infra management user
	InfraManagementUser *VSphereCredentials `json:"infraManagementUser,omitempty"`

	// templates
	Templates ImageList `json:"templates,omitempty"`
}

// Validate validates this datacenter spec v sphere
func (m *DatacenterSpecVSphere) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInfraManagementUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatacenterSpecVSphere) validateInfraManagementUser(formats strfmt.Registry) error {
	if swag.IsZero(m.InfraManagementUser) { // not required
		return nil
	}

	if m.InfraManagementUser != nil {
		if err := m.InfraManagementUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infraManagementUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infraManagementUser")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpecVSphere) validateTemplates(formats strfmt.Registry) error {
	if swag.IsZero(m.Templates) { // not required
		return nil
	}

	if m.Templates != nil {
		if err := m.Templates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("templates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("templates")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this datacenter spec v sphere based on the context it is used
func (m *DatacenterSpecVSphere) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInfraManagementUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTemplates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatacenterSpecVSphere) contextValidateInfraManagementUser(ctx context.Context, formats strfmt.Registry) error {

	if m.InfraManagementUser != nil {
		if err := m.InfraManagementUser.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("infraManagementUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("infraManagementUser")
			}
			return err
		}
	}

	return nil
}

func (m *DatacenterSpecVSphere) contextValidateTemplates(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Templates.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("templates")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("templates")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatacenterSpecVSphere) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatacenterSpecVSphere) UnmarshalBinary(b []byte) error {
	var res DatacenterSpecVSphere
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
