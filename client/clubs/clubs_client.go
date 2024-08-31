// Code generated by go-swagger; DO NOT EDIT.

package clubs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new clubs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new clubs API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new clubs API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for clubs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetClubActivitiesByID(params *GetClubActivitiesByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClubActivitiesByIDOK, error)

	GetClubAdminsByID(params *GetClubAdminsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClubAdminsByIDOK, error)

	GetClubByID(params *GetClubByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClubByIDOK, error)

	GetClubMembersByID(params *GetClubMembersByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClubMembersByIDOK, error)

	GetLoggedInAthleteClubs(params *GetLoggedInAthleteClubsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLoggedInAthleteClubsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetClubActivitiesByID lists club activities

Retrieve recent activities from members of a specific club. The authenticated athlete must belong to the requested club in order to hit this endpoint. Pagination is supported. Athlete profile visibility is respected for all activities.
*/
func (a *Client) GetClubActivitiesByID(params *GetClubActivitiesByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClubActivitiesByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClubActivitiesByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getClubActivitiesById",
		Method:             "GET",
		PathPattern:        "/clubs/{id}/activities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClubActivitiesByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClubActivitiesByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClubActivitiesByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetClubAdminsByID lists club administrators

Returns a list of the administrators of a given club.
*/
func (a *Client) GetClubAdminsByID(params *GetClubAdminsByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClubAdminsByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClubAdminsByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getClubAdminsById",
		Method:             "GET",
		PathPattern:        "/clubs/{id}/admins",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClubAdminsByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClubAdminsByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClubAdminsByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetClubByID gets club

Returns a given club using its identifier.
*/
func (a *Client) GetClubByID(params *GetClubByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClubByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClubByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getClubById",
		Method:             "GET",
		PathPattern:        "/clubs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClubByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClubByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClubByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetClubMembersByID lists club members

Returns a list of the athletes who are members of a given club.
*/
func (a *Client) GetClubMembersByID(params *GetClubMembersByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClubMembersByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClubMembersByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getClubMembersById",
		Method:             "GET",
		PathPattern:        "/clubs/{id}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClubMembersByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetClubMembersByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClubMembersByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetLoggedInAthleteClubs lists athlete clubs

Returns a list of the clubs whose membership includes the authenticated athlete.
*/
func (a *Client) GetLoggedInAthleteClubs(params *GetLoggedInAthleteClubsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLoggedInAthleteClubsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoggedInAthleteClubsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLoggedInAthleteClubs",
		Method:             "GET",
		PathPattern:        "/athlete/clubs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLoggedInAthleteClubsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLoggedInAthleteClubsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLoggedInAthleteClubsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
