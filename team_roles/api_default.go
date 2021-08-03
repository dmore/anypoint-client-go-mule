/*
 * Team Roles API
 *
 * Description of the Team Roles API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package team_roles

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// DefaultApiService DefaultApi service
type DefaultApiService service

type DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesDeleteRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	teamId string
	requestBody *[]map[string]interface{}
}

func (r DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesDeleteRequest) RequestBody(requestBody []map[string]interface{}) DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesDeleteRequest {
	r.requestBody = &requestBody
	return r
}

func (r DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdTeamsTeamIdRolesDeleteExecute(r)
}

/*
 * OrganizationsOrgIdTeamsTeamIdRolesDelete Method for OrganizationsOrgIdTeamsTeamIdRolesDelete
 * deletes a set of roles from the given team
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The ID of the organization in GUID format
 * @param teamId The ID of the team in GUID format
 * @return DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesDeleteRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdTeamsTeamIdRolesDelete(ctx _context.Context, orgId string, teamId string) DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesDeleteRequest {
	return DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesDeleteRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		teamId: teamId,
	}
}

/*
 * Execute executes the request
 */
func (a *DefaultApiService) OrganizationsOrgIdTeamsTeamIdRolesDeleteExecute(r DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdTeamsTeamIdRolesDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/teams/{teamId}/roles"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", _neturl.PathEscape(parameterToString(r.teamId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.requestBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	teamId string
	roleId *string
	search *string
	limit *int32
	offset *int32
}

func (r DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesGetRequest) RoleId(roleId string) DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesGetRequest {
	r.roleId = &roleId
	return r
}
func (r DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesGetRequest) Search(search string) DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesGetRequest {
	r.search = &search
	return r
}
func (r DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesGetRequest) Limit(limit int32) DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesGetRequest {
	r.limit = &limit
	return r
}
func (r DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesGetRequest) Offset(offset int32) DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesGetRequest {
	r.offset = &offset
	return r
}

func (r DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesGetRequest) Execute() (TeamRoleCollection, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdTeamsTeamIdRolesGetExecute(r)
}

/*
 * OrganizationsOrgIdTeamsTeamIdRolesGet Method for OrganizationsOrgIdTeamsTeamIdRolesGet
 * retrieves team assigned roles
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The ID of the organization in GUID format
 * @param teamId The ID of the team in GUID format
 * @return DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesGetRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdTeamsTeamIdRolesGet(ctx _context.Context, orgId string, teamId string) DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesGetRequest {
	return DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesGetRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		teamId: teamId,
	}
}

/*
 * Execute executes the request
 * @return TeamRoleCollection
 */
func (a *DefaultApiService) OrganizationsOrgIdTeamsTeamIdRolesGetExecute(r DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesGetRequest) (TeamRoleCollection, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TeamRoleCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdTeamsTeamIdRolesGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/teams/{teamId}/roles"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", _neturl.PathEscape(parameterToString(r.teamId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.roleId != nil {
		localVarQueryParams.Add("role_id", parameterToString(*r.roleId, ""))
	}
	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorsResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesPostRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	teamId string
	requestBody *[]map[string]interface{}
}

func (r DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesPostRequest) RequestBody(requestBody []map[string]interface{}) DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesPostRequest {
	r.requestBody = &requestBody
	return r
}

func (r DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesPostRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdTeamsTeamIdRolesPostExecute(r)
}

/*
 * OrganizationsOrgIdTeamsTeamIdRolesPost Method for OrganizationsOrgIdTeamsTeamIdRolesPost
 * assignes a set of roles to the given team
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The ID of the organization in GUID format
 * @param teamId The ID of the team in GUID format
 * @return DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesPostRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdTeamsTeamIdRolesPost(ctx _context.Context, orgId string, teamId string) DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesPostRequest {
	return DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesPostRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		teamId: teamId,
	}
}

/*
 * Execute executes the request
 */
func (a *DefaultApiService) OrganizationsOrgIdTeamsTeamIdRolesPostExecute(r DefaultApiApiOrganizationsOrgIdTeamsTeamIdRolesPostRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdTeamsTeamIdRolesPost")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/teams/{teamId}/roles"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", _neturl.PathEscape(parameterToString(r.teamId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.requestBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
