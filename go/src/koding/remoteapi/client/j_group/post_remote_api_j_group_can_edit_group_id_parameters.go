package j_group

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
)

// NewPostRemoteAPIJGroupCanEditGroupIDParams creates a new PostRemoteAPIJGroupCanEditGroupIDParams object
// with the default values initialized.
func NewPostRemoteAPIJGroupCanEditGroupIDParams() *PostRemoteAPIJGroupCanEditGroupIDParams {
	var ()
	return &PostRemoteAPIJGroupCanEditGroupIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJGroupCanEditGroupIDParamsWithTimeout creates a new PostRemoteAPIJGroupCanEditGroupIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJGroupCanEditGroupIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJGroupCanEditGroupIDParams {
	var ()
	return &PostRemoteAPIJGroupCanEditGroupIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJGroupCanEditGroupIDParamsWithContext creates a new PostRemoteAPIJGroupCanEditGroupIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJGroupCanEditGroupIDParamsWithContext(ctx context.Context) *PostRemoteAPIJGroupCanEditGroupIDParams {
	var ()
	return &PostRemoteAPIJGroupCanEditGroupIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJGroupCanEditGroupIDParams contains all the parameters to send to the API endpoint
for the post remote API j group can edit group ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJGroupCanEditGroupIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j group can edit group ID params
func (o *PostRemoteAPIJGroupCanEditGroupIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJGroupCanEditGroupIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j group can edit group ID params
func (o *PostRemoteAPIJGroupCanEditGroupIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j group can edit group ID params
func (o *PostRemoteAPIJGroupCanEditGroupIDParams) WithContext(ctx context.Context) *PostRemoteAPIJGroupCanEditGroupIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j group can edit group ID params
func (o *PostRemoteAPIJGroupCanEditGroupIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j group can edit group ID params
func (o *PostRemoteAPIJGroupCanEditGroupIDParams) WithID(id string) *PostRemoteAPIJGroupCanEditGroupIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j group can edit group ID params
func (o *PostRemoteAPIJGroupCanEditGroupIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJGroupCanEditGroupIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
