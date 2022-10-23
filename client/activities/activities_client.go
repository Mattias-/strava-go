// Code generated by go-swagger; DO NOT EDIT.

package activities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new activities API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for activities API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateActivity(params *CreateActivityParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateActivityCreated, error)

	GetActivityByID(params *GetActivityByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetActivityByIDOK, error)

	GetCommentsByActivityID(params *GetCommentsByActivityIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCommentsByActivityIDOK, error)

	GetKudoersByActivityID(params *GetKudoersByActivityIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetKudoersByActivityIDOK, error)

	GetLapsByActivityID(params *GetLapsByActivityIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLapsByActivityIDOK, error)

	GetLoggedInAthleteActivities(params *GetLoggedInAthleteActivitiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLoggedInAthleteActivitiesOK, error)

	GetZonesByActivityID(params *GetZonesByActivityIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetZonesByActivityIDOK, error)

	UpdateActivityByID(params *UpdateActivityByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateActivityByIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateActivity creates an activity

Creates a manual activity for an athlete, requires activity:write scope.
*/
func (a *Client) CreateActivity(params *CreateActivityParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateActivityCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateActivityParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createActivity",
		Method:             "POST",
		PathPattern:        "/activities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateActivityReader{formats: a.formats},
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
	success, ok := result.(*CreateActivityCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateActivityDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetActivityByID gets activity

Returns the given activity that is owned by the authenticated athlete. Requires activity:read for Everyone and Followers activities. Requires activity:read_all for Only Me activities.
*/
func (a *Client) GetActivityByID(params *GetActivityByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetActivityByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActivityByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getActivityById",
		Method:             "GET",
		PathPattern:        "/activities/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetActivityByIDReader{formats: a.formats},
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
	success, ok := result.(*GetActivityByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetActivityByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCommentsByActivityID lists activity comments

Returns the comments on the given activity. Requires activity:read for Everyone and Followers activities. Requires activity:read_all for Only Me activities.
*/
func (a *Client) GetCommentsByActivityID(params *GetCommentsByActivityIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCommentsByActivityIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCommentsByActivityIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCommentsByActivityId",
		Method:             "GET",
		PathPattern:        "/activities/{id}/comments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCommentsByActivityIDReader{formats: a.formats},
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
	success, ok := result.(*GetCommentsByActivityIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCommentsByActivityIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetKudoersByActivityID lists activity kudoers

Returns the athletes who kudoed an activity identified by an identifier. Requires activity:read for Everyone and Followers activities. Requires activity:read_all for Only Me activities.
*/
func (a *Client) GetKudoersByActivityID(params *GetKudoersByActivityIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetKudoersByActivityIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKudoersByActivityIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getKudoersByActivityId",
		Method:             "GET",
		PathPattern:        "/activities/{id}/kudos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKudoersByActivityIDReader{formats: a.formats},
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
	success, ok := result.(*GetKudoersByActivityIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetKudoersByActivityIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetLapsByActivityID lists activity laps

Returns the laps of an activity identified by an identifier. Requires activity:read for Everyone and Followers activities. Requires activity:read_all for Only Me activities.
*/
func (a *Client) GetLapsByActivityID(params *GetLapsByActivityIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLapsByActivityIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLapsByActivityIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLapsByActivityId",
		Method:             "GET",
		PathPattern:        "/activities/{id}/laps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLapsByActivityIDReader{formats: a.formats},
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
	success, ok := result.(*GetLapsByActivityIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLapsByActivityIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetLoggedInAthleteActivities lists athlete activities

Returns the activities of an athlete for a specific identifier. Requires activity:read. Only Me activities will be filtered out unless requested by a token with activity:read_all.
*/
func (a *Client) GetLoggedInAthleteActivities(params *GetLoggedInAthleteActivitiesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetLoggedInAthleteActivitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoggedInAthleteActivitiesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getLoggedInAthleteActivities",
		Method:             "GET",
		PathPattern:        "/athlete/activities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLoggedInAthleteActivitiesReader{formats: a.formats},
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
	success, ok := result.(*GetLoggedInAthleteActivitiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLoggedInAthleteActivitiesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetZonesByActivityID gets activity zones

Summit Feature. Returns the zones of a given activity. Requires activity:read for Everyone and Followers activities. Requires activity:read_all for Only Me activities.
*/
func (a *Client) GetZonesByActivityID(params *GetZonesByActivityIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetZonesByActivityIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetZonesByActivityIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getZonesByActivityId",
		Method:             "GET",
		PathPattern:        "/activities/{id}/zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetZonesByActivityIDReader{formats: a.formats},
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
	success, ok := result.(*GetZonesByActivityIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetZonesByActivityIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateActivityByID updates activity

Updates the given activity that is owned by the authenticated athlete. Requires activity:write. Also requires activity:read_all in order to update Only Me activities
*/
func (a *Client) UpdateActivityByID(params *UpdateActivityByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateActivityByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateActivityByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateActivityById",
		Method:             "PUT",
		PathPattern:        "/activities/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateActivityByIDReader{formats: a.formats},
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
	success, ok := result.(*UpdateActivityByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateActivityByIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
