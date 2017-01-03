package j_proxy_restriction

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

// NewPostRemoteAPIJProxyRestrictionCreateParams creates a new PostRemoteAPIJProxyRestrictionCreateParams object
// with the default values initialized.
func NewPostRemoteAPIJProxyRestrictionCreateParams() *PostRemoteAPIJProxyRestrictionCreateParams {
	var ()
	return &PostRemoteAPIJProxyRestrictionCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJProxyRestrictionCreateParamsWithTimeout creates a new PostRemoteAPIJProxyRestrictionCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJProxyRestrictionCreateParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJProxyRestrictionCreateParams {
	var ()
	return &PostRemoteAPIJProxyRestrictionCreateParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJProxyRestrictionCreateParamsWithContext creates a new PostRemoteAPIJProxyRestrictionCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJProxyRestrictionCreateParamsWithContext(ctx context.Context) *PostRemoteAPIJProxyRestrictionCreateParams {
	var ()
	return &PostRemoteAPIJProxyRestrictionCreateParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJProxyRestrictionCreateParams contains all the parameters to send to the API endpoint
for the post remote API j proxy restriction create operation typically these are written to a http.Request
*/
type PostRemoteAPIJProxyRestrictionCreateParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j proxy restriction create params
func (o *PostRemoteAPIJProxyRestrictionCreateParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJProxyRestrictionCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j proxy restriction create params
func (o *PostRemoteAPIJProxyRestrictionCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j proxy restriction create params
func (o *PostRemoteAPIJProxyRestrictionCreateParams) WithContext(ctx context.Context) *PostRemoteAPIJProxyRestrictionCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j proxy restriction create params
func (o *PostRemoteAPIJProxyRestrictionCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j proxy restriction create params
func (o *PostRemoteAPIJProxyRestrictionCreateParams) WithBody(body *models.DefaultSelector) *PostRemoteAPIJProxyRestrictionCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j proxy restriction create params
func (o *PostRemoteAPIJProxyRestrictionCreateParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJProxyRestrictionCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
