// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// CreateClusterTemplateReader is a Reader for the CreateClusterTemplate structure.
type CreateClusterTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClusterTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateClusterTemplateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateClusterTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateClusterTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateClusterTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateClusterTemplateCreated creates a CreateClusterTemplateCreated with default headers values
func NewCreateClusterTemplateCreated() *CreateClusterTemplateCreated {
	return &CreateClusterTemplateCreated{}
}

/* CreateClusterTemplateCreated describes a response with status code 201, with default header values.

ClusterTemplate
*/
type CreateClusterTemplateCreated struct {
	Payload *models.ClusterTemplate
}

func (o *CreateClusterTemplateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates][%d] createClusterTemplateCreated  %+v", 201, o.Payload)
}
func (o *CreateClusterTemplateCreated) GetPayload() *models.ClusterTemplate {
	return o.Payload
}

func (o *CreateClusterTemplateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterTemplateUnauthorized creates a CreateClusterTemplateUnauthorized with default headers values
func NewCreateClusterTemplateUnauthorized() *CreateClusterTemplateUnauthorized {
	return &CreateClusterTemplateUnauthorized{}
}

/* CreateClusterTemplateUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateClusterTemplateUnauthorized struct {
}

func (o *CreateClusterTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates][%d] createClusterTemplateUnauthorized ", 401)
}

func (o *CreateClusterTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateClusterTemplateForbidden creates a CreateClusterTemplateForbidden with default headers values
func NewCreateClusterTemplateForbidden() *CreateClusterTemplateForbidden {
	return &CreateClusterTemplateForbidden{}
}

/* CreateClusterTemplateForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateClusterTemplateForbidden struct {
}

func (o *CreateClusterTemplateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates][%d] createClusterTemplateForbidden ", 403)
}

func (o *CreateClusterTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateClusterTemplateDefault creates a CreateClusterTemplateDefault with default headers values
func NewCreateClusterTemplateDefault(code int) *CreateClusterTemplateDefault {
	return &CreateClusterTemplateDefault{
		_statusCode: code,
	}
}

/* CreateClusterTemplateDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateClusterTemplateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create cluster template default response
func (o *CreateClusterTemplateDefault) Code() int {
	return o._statusCode
}

func (o *CreateClusterTemplateDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clustertemplates][%d] createClusterTemplate default  %+v", o._statusCode, o.Payload)
}
func (o *CreateClusterTemplateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateClusterTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateClusterTemplateBody create cluster template body
swagger:model CreateClusterTemplateBody
*/
type CreateClusterTemplateBody struct {

	// name
	Name string `json:"name,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`

	// user SSH keys
	UserSSHKeys []string `json:"userSshKeys"`

	// cluster
	Cluster *models.Cluster `json:"cluster,omitempty"`

	// node deployment
	NodeDeployment *models.NodeDeployment `json:"nodeDeployment,omitempty"`
}

// Validate validates this create cluster template body
func (o *CreateClusterTemplateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNodeDeployment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateClusterTemplateBody) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(o.Cluster) { // not required
		return nil
	}

	if o.Cluster != nil {
		if err := o.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "cluster")
			}
			return err
		}
	}

	return nil
}

func (o *CreateClusterTemplateBody) validateNodeDeployment(formats strfmt.Registry) error {
	if swag.IsZero(o.NodeDeployment) { // not required
		return nil
	}

	if o.NodeDeployment != nil {
		if err := o.NodeDeployment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "nodeDeployment")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create cluster template body based on the context it is used
func (o *CreateClusterTemplateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateNodeDeployment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateClusterTemplateBody) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if o.Cluster != nil {
		if err := o.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "cluster")
			}
			return err
		}
	}

	return nil
}

func (o *CreateClusterTemplateBody) contextValidateNodeDeployment(ctx context.Context, formats strfmt.Registry) error {

	if o.NodeDeployment != nil {
		if err := o.NodeDeployment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body" + "." + "nodeDeployment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateClusterTemplateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateClusterTemplateBody) UnmarshalBinary(b []byte) error {
	var res CreateClusterTemplateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
