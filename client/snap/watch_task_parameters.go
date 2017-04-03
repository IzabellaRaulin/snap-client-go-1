package snap

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

// NewWatchTaskParams creates a new WatchTaskParams object
// with the default values initialized.
func NewWatchTaskParams() *WatchTaskParams {
	var ()
	return &WatchTaskParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWatchTaskParamsWithTimeout creates a new WatchTaskParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWatchTaskParamsWithTimeout(timeout time.Duration) *WatchTaskParams {
	var ()
	return &WatchTaskParams{

		timeout: timeout,
	}
}

// NewWatchTaskParamsWithContext creates a new WatchTaskParams object
// with the default values initialized, and the ability to set a context for a request
func NewWatchTaskParamsWithContext(ctx context.Context) *WatchTaskParams {
	var ()
	return &WatchTaskParams{

		Context: ctx,
	}
}

// NewWatchTaskParamsWithHTTPClient creates a new WatchTaskParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWatchTaskParamsWithHTTPClient(client *http.Client) *WatchTaskParams {
	var ()
	return &WatchTaskParams{
		HTTPClient: client,
	}
}

/*WatchTaskParams contains all the parameters to send to the API endpoint
for the watch task operation typically these are written to a http.Request
*/
type WatchTaskParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the watch task params
func (o *WatchTaskParams) WithTimeout(timeout time.Duration) *WatchTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the watch task params
func (o *WatchTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the watch task params
func (o *WatchTaskParams) WithContext(ctx context.Context) *WatchTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the watch task params
func (o *WatchTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the watch task params
func (o *WatchTaskParams) WithHTTPClient(client *http.Client) *WatchTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the watch task params
func (o *WatchTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the watch task params
func (o *WatchTaskParams) WithID(id string) *WatchTaskParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the watch task params
func (o *WatchTaskParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *WatchTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
