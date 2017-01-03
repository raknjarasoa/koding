package j_workspace

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

// NewPostRemoteAPIJWorkspaceDeleteByIDParams creates a new PostRemoteAPIJWorkspaceDeleteByIDParams object
// with the default values initialized.
func NewPostRemoteAPIJWorkspaceDeleteByIDParams() *PostRemoteAPIJWorkspaceDeleteByIDParams {
	var ()
	return &PostRemoteAPIJWorkspaceDeleteByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJWorkspaceDeleteByIDParamsWithTimeout creates a new PostRemoteAPIJWorkspaceDeleteByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJWorkspaceDeleteByIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJWorkspaceDeleteByIDParams {
	var ()
	return &PostRemoteAPIJWorkspaceDeleteByIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJWorkspaceDeleteByIDParamsWithContext creates a new PostRemoteAPIJWorkspaceDeleteByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJWorkspaceDeleteByIDParamsWithContext(ctx context.Context) *PostRemoteAPIJWorkspaceDeleteByIDParams {
	var ()
	return &PostRemoteAPIJWorkspaceDeleteByIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJWorkspaceDeleteByIDParams contains all the parameters to send to the API endpoint
for the post remote API j workspace delete by ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJWorkspaceDeleteByIDParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j workspace delete by ID params
func (o *PostRemoteAPIJWorkspaceDeleteByIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJWorkspaceDeleteByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j workspace delete by ID params
func (o *PostRemoteAPIJWorkspaceDeleteByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j workspace delete by ID params
func (o *PostRemoteAPIJWorkspaceDeleteByIDParams) WithContext(ctx context.Context) *PostRemoteAPIJWorkspaceDeleteByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j workspace delete by ID params
func (o *PostRemoteAPIJWorkspaceDeleteByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j workspace delete by ID params
func (o *PostRemoteAPIJWorkspaceDeleteByIDParams) WithBody(body *models.DefaultSelector) *PostRemoteAPIJWorkspaceDeleteByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j workspace delete by ID params
func (o *PostRemoteAPIJWorkspaceDeleteByIDParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJWorkspaceDeleteByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
