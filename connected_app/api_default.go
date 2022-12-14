/*
 * Connected App API
 *
 * Description of the Connected App API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connected_app

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

type DefaultApiApiConnectedApplicationsConnAppIdDeleteRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	connAppId string
}


func (r DefaultApiApiConnectedApplicationsConnAppIdDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ConnectedApplicationsConnAppIdDeleteExecute(r)
}

/*
 * ConnectedApplicationsConnAppIdDelete Method for ConnectedApplicationsConnAppIdDelete
 * deletes a Connected App
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connAppId The ID of the connected app
 * @return DefaultApiApiConnectedApplicationsConnAppIdDeleteRequest
 */
func (a *DefaultApiService) ConnectedApplicationsConnAppIdDelete(ctx _context.Context, connAppId string) DefaultApiApiConnectedApplicationsConnAppIdDeleteRequest {
	return DefaultApiApiConnectedApplicationsConnAppIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		connAppId: connAppId,
	}
}

/*
 * Execute executes the request
 */
func (a *DefaultApiService) ConnectedApplicationsConnAppIdDeleteExecute(r DefaultApiApiConnectedApplicationsConnAppIdDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ConnectedApplicationsConnAppIdDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connectedApplications/{connAppId}"
	localVarPath = strings.Replace(localVarPath, "{"+"connAppId"+"}", _neturl.PathEscape(parameterToString(r.connAppId, "")), -1)

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

type DefaultApiApiConnectedApplicationsConnAppIdGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	connAppId string
}


func (r DefaultApiApiConnectedApplicationsConnAppIdGetRequest) Execute() (ConnectedAppRespExt, *_nethttp.Response, error) {
	return r.ApiService.ConnectedApplicationsConnAppIdGetExecute(r)
}

/*
 * ConnectedApplicationsConnAppIdGet Method for ConnectedApplicationsConnAppIdGet
 * Returns all connected apps
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connAppId The ID of the connected app
 * @return DefaultApiApiConnectedApplicationsConnAppIdGetRequest
 */
func (a *DefaultApiService) ConnectedApplicationsConnAppIdGet(ctx _context.Context, connAppId string) DefaultApiApiConnectedApplicationsConnAppIdGetRequest {
	return DefaultApiApiConnectedApplicationsConnAppIdGetRequest{
		ApiService: a,
		ctx: ctx,
		connAppId: connAppId,
	}
}

/*
 * Execute executes the request
 * @return ConnectedAppRespExt
 */
func (a *DefaultApiService) ConnectedApplicationsConnAppIdGetExecute(r DefaultApiApiConnectedApplicationsConnAppIdGetRequest) (ConnectedAppRespExt, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConnectedAppRespExt
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ConnectedApplicationsConnAppIdGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connectedApplications/{connAppId}"
	localVarPath = strings.Replace(localVarPath, "{"+"connAppId"+"}", _neturl.PathEscape(parameterToString(r.connAppId, "")), -1)

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

type DefaultApiApiConnectedApplicationsConnAppIdPatchRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	connAppId string
	connectedAppPatchExt *ConnectedAppPatchExt
}

func (r DefaultApiApiConnectedApplicationsConnAppIdPatchRequest) ConnectedAppPatchExt(connectedAppPatchExt ConnectedAppPatchExt) DefaultApiApiConnectedApplicationsConnAppIdPatchRequest {
	r.connectedAppPatchExt = &connectedAppPatchExt
	return r
}

func (r DefaultApiApiConnectedApplicationsConnAppIdPatchRequest) Execute() (ConnectedAppRespExt, *_nethttp.Response, error) {
	return r.ApiService.ConnectedApplicationsConnAppIdPatchExecute(r)
}

/*
 * ConnectedApplicationsConnAppIdPatch Method for ConnectedApplicationsConnAppIdPatch
 * patches a Connected App
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connAppId The ID of the connected app
 * @return DefaultApiApiConnectedApplicationsConnAppIdPatchRequest
 */
func (a *DefaultApiService) ConnectedApplicationsConnAppIdPatch(ctx _context.Context, connAppId string) DefaultApiApiConnectedApplicationsConnAppIdPatchRequest {
	return DefaultApiApiConnectedApplicationsConnAppIdPatchRequest{
		ApiService: a,
		ctx: ctx,
		connAppId: connAppId,
	}
}

/*
 * Execute executes the request
 * @return ConnectedAppRespExt
 */
func (a *DefaultApiService) ConnectedApplicationsConnAppIdPatchExecute(r DefaultApiApiConnectedApplicationsConnAppIdPatchRequest) (ConnectedAppRespExt, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConnectedAppRespExt
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ConnectedApplicationsConnAppIdPatch")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connectedApplications/{connAppId}"
	localVarPath = strings.Replace(localVarPath, "{"+"connAppId"+"}", _neturl.PathEscape(parameterToString(r.connAppId, "")), -1)

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
	localVarPostBody = r.connectedAppPatchExt
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

type DefaultApiApiConnectedApplicationsConnAppIdScopesGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	connAppId string
}


func (r DefaultApiApiConnectedApplicationsConnAppIdScopesGetRequest) Execute() (InlineResponse2001, *_nethttp.Response, error) {
	return r.ApiService.ConnectedApplicationsConnAppIdScopesGetExecute(r)
}

/*
 * ConnectedApplicationsConnAppIdScopesGet Method for ConnectedApplicationsConnAppIdScopesGet
 * Returns all scopes of a Connected App
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connAppId The ID of the connected app
 * @return DefaultApiApiConnectedApplicationsConnAppIdScopesGetRequest
 */
func (a *DefaultApiService) ConnectedApplicationsConnAppIdScopesGet(ctx _context.Context, connAppId string) DefaultApiApiConnectedApplicationsConnAppIdScopesGetRequest {
	return DefaultApiApiConnectedApplicationsConnAppIdScopesGetRequest{
		ApiService: a,
		ctx: ctx,
		connAppId: connAppId,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse2001
 */
func (a *DefaultApiService) ConnectedApplicationsConnAppIdScopesGetExecute(r DefaultApiApiConnectedApplicationsConnAppIdScopesGetRequest) (InlineResponse2001, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2001
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ConnectedApplicationsConnAppIdScopesGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connectedApplications/{connAppId}/scopes"
	localVarPath = strings.Replace(localVarPath, "{"+"connAppId"+"}", _neturl.PathEscape(parameterToString(r.connAppId, "")), -1)

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

type DefaultApiApiConnectedApplicationsConnAppIdScopesPutRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	connAppId string
	connectedAppScopesPutBody *ConnectedAppScopesPutBody
}

func (r DefaultApiApiConnectedApplicationsConnAppIdScopesPutRequest) ConnectedAppScopesPutBody(connectedAppScopesPutBody ConnectedAppScopesPutBody) DefaultApiApiConnectedApplicationsConnAppIdScopesPutRequest {
	r.connectedAppScopesPutBody = &connectedAppScopesPutBody
	return r
}

func (r DefaultApiApiConnectedApplicationsConnAppIdScopesPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ConnectedApplicationsConnAppIdScopesPutExecute(r)
}

/*
 * ConnectedApplicationsConnAppIdScopesPut Method for ConnectedApplicationsConnAppIdScopesPut
 * replace a Connected App scopes
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connAppId The ID of the connected app
 * @return DefaultApiApiConnectedApplicationsConnAppIdScopesPutRequest
 */
func (a *DefaultApiService) ConnectedApplicationsConnAppIdScopesPut(ctx _context.Context, connAppId string) DefaultApiApiConnectedApplicationsConnAppIdScopesPutRequest {
	return DefaultApiApiConnectedApplicationsConnAppIdScopesPutRequest{
		ApiService: a,
		ctx: ctx,
		connAppId: connAppId,
	}
}

/*
 * Execute executes the request
 */
func (a *DefaultApiService) ConnectedApplicationsConnAppIdScopesPutExecute(r DefaultApiApiConnectedApplicationsConnAppIdScopesPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ConnectedApplicationsConnAppIdScopesPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connectedApplications/{connAppId}/scopes"
	localVarPath = strings.Replace(localVarPath, "{"+"connAppId"+"}", _neturl.PathEscape(parameterToString(r.connAppId, "")), -1)

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
	localVarPostBody = r.connectedAppScopesPutBody
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v InlineResponse400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type DefaultApiApiConnectedApplicationsGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
}


func (r DefaultApiApiConnectedApplicationsGetRequest) Execute() (InlineResponse200, *_nethttp.Response, error) {
	return r.ApiService.ConnectedApplicationsGetExecute(r)
}

/*
 * ConnectedApplicationsGet Method for ConnectedApplicationsGet
 * Returns all connected apps
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return DefaultApiApiConnectedApplicationsGetRequest
 */
func (a *DefaultApiService) ConnectedApplicationsGet(ctx _context.Context) DefaultApiApiConnectedApplicationsGetRequest {
	return DefaultApiApiConnectedApplicationsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse200
 */
func (a *DefaultApiService) ConnectedApplicationsGetExecute(r DefaultApiApiConnectedApplicationsGetRequest) (InlineResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ConnectedApplicationsGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connectedApplications"

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

type DefaultApiApiConnectedApplicationsPostRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	connectedAppCore *ConnectedAppCore
}

func (r DefaultApiApiConnectedApplicationsPostRequest) ConnectedAppCore(connectedAppCore ConnectedAppCore) DefaultApiApiConnectedApplicationsPostRequest {
	r.connectedAppCore = &connectedAppCore
	return r
}

func (r DefaultApiApiConnectedApplicationsPostRequest) Execute() (ConnectedAppRespExt, *_nethttp.Response, error) {
	return r.ApiService.ConnectedApplicationsPostExecute(r)
}

/*
 * ConnectedApplicationsPost Method for ConnectedApplicationsPost
 * create a Connected App
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return DefaultApiApiConnectedApplicationsPostRequest
 */
func (a *DefaultApiService) ConnectedApplicationsPost(ctx _context.Context) DefaultApiApiConnectedApplicationsPostRequest {
	return DefaultApiApiConnectedApplicationsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ConnectedAppRespExt
 */
func (a *DefaultApiService) ConnectedApplicationsPostExecute(r DefaultApiApiConnectedApplicationsPostRequest) (ConnectedAppRespExt, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConnectedAppRespExt
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ConnectedApplicationsPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connectedApplications"

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
	localVarPostBody = r.connectedAppCore
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
