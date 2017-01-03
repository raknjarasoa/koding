package j_provisioner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewPostRemoteAPIJProvisionerCreateParams creates a new PostRemoteAPIJProvisionerCreateParams object
// with the default values initialized.
func NewPostRemoteAPIJProvisionerCreateParams() *PostRemoteAPIJProvisionerCreateParams {
	var ()
	return &PostRemoteAPIJProvisionerCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJProvisionerCreateParamsWithTimeout creates a new PostRemoteAPIJProvisionerCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJProvisionerCreateParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJProvisionerCreateParams {
	var ()
	return &PostRemoteAPIJProvisionerCreateParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJProvisionerCreateParamsWithContext creates a new PostRemoteAPIJProvisionerCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJProvisionerCreateParamsWithContext(ctx context.Context) *PostRemoteAPIJProvisionerCreateParams {
	var ()
	return &PostRemoteAPIJProvisionerCreateParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJProvisionerCreateParams contains all the parameters to send to the API endpoint
for the post remote API j provisioner create operation typically these are written to a http.Request
*/
type PostRemoteAPIJProvisionerCreateParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j provisioner create params
func (o *PostRemoteAPIJProvisionerCreateParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJProvisionerCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j provisioner create params
func (o *PostRemoteAPIJProvisionerCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j provisioner create params
func (o *PostRemoteAPIJProvisionerCreateParams) WithContext(ctx context.Context) *PostRemoteAPIJProvisionerCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j provisioner create params
func (o *PostRemoteAPIJProvisionerCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j provisioner create params
func (o *PostRemoteAPIJProvisionerCreateParams) WithBody(body *models.DefaultSelector) *PostRemoteAPIJProvisionerCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j provisioner create params
func (o *PostRemoteAPIJProvisionerCreateParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJProvisionerCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.DefaultSelector)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
