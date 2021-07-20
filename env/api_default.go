/*
 * ENV API
 *
 * Description of the ENV API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package env

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

type DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdDeleteRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	environmentId string
}


func (r DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdEnvironmentsEnvironmentIdDeleteExecute(r)
}

/*
 * OrganizationsOrgIdEnvironmentsEnvironmentIdDelete Method for OrganizationsOrgIdEnvironmentsEnvironmentIdDelete
 * Delete an environment
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The ID of the organization in GUID format
 * @param environmentId The id of an environment
 * @return DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdDeleteRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdEnvironmentsEnvironmentIdDelete(ctx _context.Context, orgId string, environmentId string) DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdDeleteRequest {
	return DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		environmentId: environmentId,
	}
}

/*
 * Execute executes the request
 */
func (a *DefaultApiService) OrganizationsOrgIdEnvironmentsEnvironmentIdDeleteExecute(r DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdEnvironmentsEnvironmentIdDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments/{environmentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	environmentId string
}


func (r DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdGetRequest) Execute() (Env, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdEnvironmentsEnvironmentIdGetExecute(r)
}

/*
 * OrganizationsOrgIdEnvironmentsEnvironmentIdGet Method for OrganizationsOrgIdEnvironmentsEnvironmentIdGet
 * Retrieves an environment by id
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The ID of the organization in GUID format
 * @param environmentId The id of an environment
 * @return DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdGetRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdEnvironmentsEnvironmentIdGet(ctx _context.Context, orgId string, environmentId string) DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdGetRequest {
	return DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdGetRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		environmentId: environmentId,
	}
}

/*
 * Execute executes the request
 * @return Env
 */
func (a *DefaultApiService) OrganizationsOrgIdEnvironmentsEnvironmentIdGetExecute(r DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdGetRequest) (Env, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Env
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdEnvironmentsEnvironmentIdGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments/{environmentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdPutRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	environmentId string
	envCore *EnvCore
}

func (r DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdPutRequest) EnvCore(envCore EnvCore) DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdPutRequest {
	r.envCore = &envCore
	return r
}

func (r DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdPutRequest) Execute() (Env, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdEnvironmentsEnvironmentIdPutExecute(r)
}

/*
 * OrganizationsOrgIdEnvironmentsEnvironmentIdPut Method for OrganizationsOrgIdEnvironmentsEnvironmentIdPut
 * Update an environment, implemented as a patch. Note that only the name is allowed to be updated, isProduction and type can not.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The ID of the organization in GUID format
 * @param environmentId The id of an environment
 * @return DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdPutRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdEnvironmentsEnvironmentIdPut(ctx _context.Context, orgId string, environmentId string) DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdPutRequest {
	return DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdPutRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		environmentId: environmentId,
	}
}

/*
 * Execute executes the request
 * @return Env
 */
func (a *DefaultApiService) OrganizationsOrgIdEnvironmentsEnvironmentIdPutExecute(r DefaultApiApiOrganizationsOrgIdEnvironmentsEnvironmentIdPutRequest) (Env, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Env
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdEnvironmentsEnvironmentIdPut")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments/{environmentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.envCore
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
			var v InlineResponse400
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

type DefaultApiApiOrganizationsOrgIdEnvironmentsGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
}


func (r DefaultApiApiOrganizationsOrgIdEnvironmentsGetRequest) Execute() (InlineResponse200, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdEnvironmentsGetExecute(r)
}

/*
 * OrganizationsOrgIdEnvironmentsGet Method for OrganizationsOrgIdEnvironmentsGet
 * Returns all matching environments
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The ID of the organization in GUID format
 * @return DefaultApiApiOrganizationsOrgIdEnvironmentsGetRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdEnvironmentsGet(ctx _context.Context, orgId string) DefaultApiApiOrganizationsOrgIdEnvironmentsGetRequest {
	return DefaultApiApiOrganizationsOrgIdEnvironmentsGetRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse200
 */
func (a *DefaultApiService) OrganizationsOrgIdEnvironmentsGetExecute(r DefaultApiApiOrganizationsOrgIdEnvironmentsGetRequest) (InlineResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdEnvironmentsGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type DefaultApiApiOrganizationsOrgIdEnvironmentsPostRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	orgId string
	envCore *EnvCore
}

func (r DefaultApiApiOrganizationsOrgIdEnvironmentsPostRequest) EnvCore(envCore EnvCore) DefaultApiApiOrganizationsOrgIdEnvironmentsPostRequest {
	r.envCore = &envCore
	return r
}

func (r DefaultApiApiOrganizationsOrgIdEnvironmentsPostRequest) Execute() (Env, *_nethttp.Response, error) {
	return r.ApiService.OrganizationsOrgIdEnvironmentsPostExecute(r)
}

/*
 * OrganizationsOrgIdEnvironmentsPost Method for OrganizationsOrgIdEnvironmentsPost
 * Creates an environment
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orgId The ID of the organization in GUID format
 * @return DefaultApiApiOrganizationsOrgIdEnvironmentsPostRequest
 */
func (a *DefaultApiService) OrganizationsOrgIdEnvironmentsPost(ctx _context.Context, orgId string) DefaultApiApiOrganizationsOrgIdEnvironmentsPostRequest {
	return DefaultApiApiOrganizationsOrgIdEnvironmentsPostRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

/*
 * Execute executes the request
 * @return Env
 */
func (a *DefaultApiService) OrganizationsOrgIdEnvironmentsPostExecute(r DefaultApiApiOrganizationsOrgIdEnvironmentsPostRequest) (Env, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Env
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.OrganizationsOrgIdEnvironmentsPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{orgId}/environments"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", _neturl.PathEscape(parameterToString(r.orgId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.envCore
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
			var v InlineResponse400
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