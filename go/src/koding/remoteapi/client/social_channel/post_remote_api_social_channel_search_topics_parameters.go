package social_channel

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

// NewPostRemoteAPISocialChannelSearchTopicsParams creates a new PostRemoteAPISocialChannelSearchTopicsParams object
// with the default values initialized.
func NewPostRemoteAPISocialChannelSearchTopicsParams() *PostRemoteAPISocialChannelSearchTopicsParams {
	var ()
	return &PostRemoteAPISocialChannelSearchTopicsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPISocialChannelSearchTopicsParamsWithTimeout creates a new PostRemoteAPISocialChannelSearchTopicsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPISocialChannelSearchTopicsParamsWithTimeout(timeout time.Duration) *PostRemoteAPISocialChannelSearchTopicsParams {
	var ()
	return &PostRemoteAPISocialChannelSearchTopicsParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPISocialChannelSearchTopicsParamsWithContext creates a new PostRemoteAPISocialChannelSearchTopicsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPISocialChannelSearchTopicsParamsWithContext(ctx context.Context) *PostRemoteAPISocialChannelSearchTopicsParams {
	var ()
	return &PostRemoteAPISocialChannelSearchTopicsParams{

		Context: ctx,
	}
}

/*PostRemoteAPISocialChannelSearchTopicsParams contains all the parameters to send to the API endpoint
for the post remote API social channel search topics operation typically these are written to a http.Request
*/
type PostRemoteAPISocialChannelSearchTopicsParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API social channel search topics params
func (o *PostRemoteAPISocialChannelSearchTopicsParams) WithTimeout(timeout time.Duration) *PostRemoteAPISocialChannelSearchTopicsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API social channel search topics params
func (o *PostRemoteAPISocialChannelSearchTopicsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API social channel search topics params
func (o *PostRemoteAPISocialChannelSearchTopicsParams) WithContext(ctx context.Context) *PostRemoteAPISocialChannelSearchTopicsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API social channel search topics params
func (o *PostRemoteAPISocialChannelSearchTopicsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API social channel search topics params
func (o *PostRemoteAPISocialChannelSearchTopicsParams) WithBody(body *models.DefaultSelector) *PostRemoteAPISocialChannelSearchTopicsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API social channel search topics params
func (o *PostRemoteAPISocialChannelSearchTopicsParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPISocialChannelSearchTopicsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
