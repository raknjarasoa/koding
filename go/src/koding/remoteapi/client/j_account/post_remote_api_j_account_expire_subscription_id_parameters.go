package j_account

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

// NewPostRemoteAPIJAccountExpireSubscriptionIDParams creates a new PostRemoteAPIJAccountExpireSubscriptionIDParams object
// with the default values initialized.
func NewPostRemoteAPIJAccountExpireSubscriptionIDParams() *PostRemoteAPIJAccountExpireSubscriptionIDParams {
	var ()
	return &PostRemoteAPIJAccountExpireSubscriptionIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJAccountExpireSubscriptionIDParamsWithTimeout creates a new PostRemoteAPIJAccountExpireSubscriptionIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJAccountExpireSubscriptionIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJAccountExpireSubscriptionIDParams {
	var ()
	return &PostRemoteAPIJAccountExpireSubscriptionIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJAccountExpireSubscriptionIDParamsWithContext creates a new PostRemoteAPIJAccountExpireSubscriptionIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJAccountExpireSubscriptionIDParamsWithContext(ctx context.Context) *PostRemoteAPIJAccountExpireSubscriptionIDParams {
	var ()
	return &PostRemoteAPIJAccountExpireSubscriptionIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJAccountExpireSubscriptionIDParams contains all the parameters to send to the API endpoint
for the post remote API j account expire subscription ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJAccountExpireSubscriptionIDParams struct {

	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j account expire subscription ID params
func (o *PostRemoteAPIJAccountExpireSubscriptionIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJAccountExpireSubscriptionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j account expire subscription ID params
func (o *PostRemoteAPIJAccountExpireSubscriptionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j account expire subscription ID params
func (o *PostRemoteAPIJAccountExpireSubscriptionIDParams) WithContext(ctx context.Context) *PostRemoteAPIJAccountExpireSubscriptionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j account expire subscription ID params
func (o *PostRemoteAPIJAccountExpireSubscriptionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the post remote API j account expire subscription ID params
func (o *PostRemoteAPIJAccountExpireSubscriptionIDParams) WithID(id string) *PostRemoteAPIJAccountExpireSubscriptionIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j account expire subscription ID params
func (o *PostRemoteAPIJAccountExpireSubscriptionIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJAccountExpireSubscriptionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
