package sidebar

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new sidebar API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sidebar API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostRemoteAPISidebarFetchEnvironment post remote API sidebar fetch environment API
*/
func (a *Client) PostRemoteAPISidebarFetchEnvironment(params *PostRemoteAPISidebarFetchEnvironmentParams) (*PostRemoteAPISidebarFetchEnvironmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPISidebarFetchEnvironmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPISidebarFetchEnvironment",
		Method:             "POST",
		PathPattern:        "/remote.api/Sidebar.fetchEnvironment",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPISidebarFetchEnvironmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPISidebarFetchEnvironmentOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
