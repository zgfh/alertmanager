// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package alert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new alert API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for alert API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAlerts Get a list of alerts
*/
func (a *Client) GetAlerts(params *GetAlertsParams) (*GetAlertsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAlertsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAlerts",
		Method:             "GET",
		PathPattern:        "/alerts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAlertsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAlertsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAlerts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OpenapiPostAlerts Create new Alerts
*/
func (a *Client) OpenapiPostAlerts(params *OpenapiPostAlertsParams) (*OpenapiPostAlertsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOpenapiPostAlertsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "openapiPostAlerts",
		Method:             "POST",
		PathPattern:        "/openapi/{tenantId}/alerts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OpenapiPostAlertsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OpenapiPostAlertsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for openapiPostAlerts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostAlerts Create new Alerts
*/
func (a *Client) PostAlerts(params *PostAlertsParams) (*PostAlertsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAlertsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postAlerts",
		Method:             "POST",
		PathPattern:        "/alerts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAlertsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostAlertsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postAlerts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
