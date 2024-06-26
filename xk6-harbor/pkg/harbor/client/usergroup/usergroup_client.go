// Code generated by go-swagger; DO NOT EDIT.

package usergroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery --name API --keeptree --with-expecter --case underscore

// API is the interface of the usergroup client
type API interface {
	/*
	   CreateUserGroup creates user group

	   Create user group information*/
	CreateUserGroup(ctx context.Context, params *CreateUserGroupParams) (*CreateUserGroupCreated, error)
	/*
	   DeleteUserGroup deletes user group

	   Delete user group*/
	DeleteUserGroup(ctx context.Context, params *DeleteUserGroupParams) (*DeleteUserGroupOK, error)
	/*
	   GetUserGroup gets user group information

	   Get user group information*/
	GetUserGroup(ctx context.Context, params *GetUserGroupParams) (*GetUserGroupOK, error)
	/*
	   ListUserGroups gets all user groups information

	   Get all user groups information, it is open for system admin*/
	ListUserGroups(ctx context.Context, params *ListUserGroupsParams) (*ListUserGroupsOK, error)
	/*
	   SearchUserGroups searches groups by groupname

	   This endpoint is to search groups by group name.  It's open for all authenticated requests.
	*/
	SearchUserGroups(ctx context.Context, params *SearchUserGroupsParams) (*SearchUserGroupsOK, error)
	/*
	   UpdateUserGroup updates group information

	   Update user group information*/
	UpdateUserGroup(ctx context.Context, params *UpdateUserGroupParams) (*UpdateUserGroupOK, error)
}

// New creates a new usergroup API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for usergroup API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
CreateUserGroup creates user group

Create user group information
*/
func (a *Client) CreateUserGroup(ctx context.Context, params *CreateUserGroupParams) (*CreateUserGroupCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createUserGroup",
		Method:             "POST",
		PathPattern:        "/usergroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateUserGroupReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *CreateUserGroupCreated:
		return value, nil
	case *CreateUserGroupBadRequest:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *CreateUserGroupUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *CreateUserGroupForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *CreateUserGroupConflict:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *CreateUserGroupInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createUserGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteUserGroup deletes user group

Delete user group
*/
func (a *Client) DeleteUserGroup(ctx context.Context, params *DeleteUserGroupParams) (*DeleteUserGroupOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUserGroup",
		Method:             "DELETE",
		PathPattern:        "/usergroups/{group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteUserGroupReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *DeleteUserGroupOK:
		return value, nil
	case *DeleteUserGroupBadRequest:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *DeleteUserGroupUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *DeleteUserGroupForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *DeleteUserGroupInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteUserGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUserGroup gets user group information

Get user group information
*/
func (a *Client) GetUserGroup(ctx context.Context, params *GetUserGroupParams) (*GetUserGroupOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserGroup",
		Method:             "GET",
		PathPattern:        "/usergroups/{group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetUserGroupReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetUserGroupOK:
		return value, nil
	case *GetUserGroupBadRequest:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetUserGroupUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetUserGroupForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetUserGroupNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *GetUserGroupInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListUserGroups gets all user groups information

Get all user groups information, it is open for system admin
*/
func (a *Client) ListUserGroups(ctx context.Context, params *ListUserGroupsParams) (*ListUserGroupsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listUserGroups",
		Method:             "GET",
		PathPattern:        "/usergroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListUserGroupsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *ListUserGroupsOK:
		return value, nil
	case *ListUserGroupsUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ListUserGroupsForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *ListUserGroupsInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listUserGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SearchUserGroups searches groups by groupname

This endpoint is to search groups by group name.  It's open for all authenticated requests.
*/
func (a *Client) SearchUserGroups(ctx context.Context, params *SearchUserGroupsParams) (*SearchUserGroupsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "searchUserGroups",
		Method:             "GET",
		PathPattern:        "/usergroups/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchUserGroupsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *SearchUserGroupsOK:
		return value, nil
	case *SearchUserGroupsUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *SearchUserGroupsInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchUserGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateUserGroup updates group information

Update user group information
*/
func (a *Client) UpdateUserGroup(ctx context.Context, params *UpdateUserGroupParams) (*UpdateUserGroupOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateUserGroup",
		Method:             "PUT",
		PathPattern:        "/usergroups/{group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateUserGroupReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *UpdateUserGroupOK:
		return value, nil
	case *UpdateUserGroupBadRequest:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *UpdateUserGroupUnauthorized:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *UpdateUserGroupForbidden:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *UpdateUserGroupNotFound:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	case *UpdateUserGroupInternalServerError:
		return nil, runtime.NewAPIError("unsuccessful response", value, value.Code())
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUserGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}
