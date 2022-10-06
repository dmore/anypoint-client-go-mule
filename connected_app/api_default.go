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

type DefaultApiApiAccountsApiConnectedApplicationsConnAppIdDeleteRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	connAppId string
}


func (r DefaultApiApiAccountsApiConnectedApplicationsConnAppIdDeleteRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.AccountsApiConnectedApplicationsConnAppIdDeleteExecute(r)
}

/*
 * AccountsApiConnectedApplicationsConnAppIdDelete Method for AccountsApiConnectedApplicationsConnAppIdDelete
 * deletes a Connected App
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connAppId The ID of the connected app
 * @return DefaultApiApiAccountsApiConnectedApplicationsConnAppIdDeleteRequest
 */
func (a *DefaultApiService) AccountsApiConnectedApplicationsConnAppIdDelete(ctx _context.Context, connAppId string) DefaultApiApiAccountsApiConnectedApplicationsConnAppIdDeleteRequest {
	return DefaultApiApiAccountsApiConnectedApplicationsConnAppIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		connAppId: connAppId,
	}
}

/*
 * Execute executes the request
 */
func (a *DefaultApiService) AccountsApiConnectedApplicationsConnAppIdDeleteExecute(r DefaultApiApiAccountsApiConnectedApplicationsConnAppIdDeleteRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.AccountsApiConnectedApplicationsConnAppIdDelete")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/api/connectedApplications/{connAppId}"
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

type DefaultApiApiAccountsApiConnectedApplicationsConnAppIdGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	connAppId string
}


func (r DefaultApiApiAccountsApiConnectedApplicationsConnAppIdGetRequest) Execute() (ConnectedAppRespExt, *_nethttp.Response, error) {
	return r.ApiService.AccountsApiConnectedApplicationsConnAppIdGetExecute(r)
}

/*
 * AccountsApiConnectedApplicationsConnAppIdGet Method for AccountsApiConnectedApplicationsConnAppIdGet
 * Returns all connected apps
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connAppId The ID of the connected app
 * @return DefaultApiApiAccountsApiConnectedApplicationsConnAppIdGetRequest
 */
func (a *DefaultApiService) AccountsApiConnectedApplicationsConnAppIdGet(ctx _context.Context, connAppId string) DefaultApiApiAccountsApiConnectedApplicationsConnAppIdGetRequest {
	return DefaultApiApiAccountsApiConnectedApplicationsConnAppIdGetRequest{
		ApiService: a,
		ctx: ctx,
		connAppId: connAppId,
	}
}

/*
 * Execute executes the request
 * @return ConnectedAppRespExt
 */
func (a *DefaultApiService) AccountsApiConnectedApplicationsConnAppIdGetExecute(r DefaultApiApiAccountsApiConnectedApplicationsConnAppIdGetRequest) (ConnectedAppRespExt, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConnectedAppRespExt
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.AccountsApiConnectedApplicationsConnAppIdGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/api/connectedApplications/{connAppId}"
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

type DefaultApiApiAccountsApiConnectedApplicationsConnAppIdPatchRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	connAppId string
	connectedAppPatchExt *ConnectedAppPatchExt
}

func (r DefaultApiApiAccountsApiConnectedApplicationsConnAppIdPatchRequest) ConnectedAppPatchExt(connectedAppPatchExt ConnectedAppPatchExt) DefaultApiApiAccountsApiConnectedApplicationsConnAppIdPatchRequest {
	r.connectedAppPatchExt = &connectedAppPatchExt
	return r
}

func (r DefaultApiApiAccountsApiConnectedApplicationsConnAppIdPatchRequest) Execute() (ConnectedAppRespExt, *_nethttp.Response, error) {
	return r.ApiService.AccountsApiConnectedApplicationsConnAppIdPatchExecute(r)
}

/*
 * AccountsApiConnectedApplicationsConnAppIdPatch Method for AccountsApiConnectedApplicationsConnAppIdPatch
 * patches a Connected App
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connAppId The ID of the connected app
 * @return DefaultApiApiAccountsApiConnectedApplicationsConnAppIdPatchRequest
 */
func (a *DefaultApiService) AccountsApiConnectedApplicationsConnAppIdPatch(ctx _context.Context, connAppId string) DefaultApiApiAccountsApiConnectedApplicationsConnAppIdPatchRequest {
	return DefaultApiApiAccountsApiConnectedApplicationsConnAppIdPatchRequest{
		ApiService: a,
		ctx: ctx,
		connAppId: connAppId,
	}
}

/*
 * Execute executes the request
 * @return ConnectedAppRespExt
 */
func (a *DefaultApiService) AccountsApiConnectedApplicationsConnAppIdPatchExecute(r DefaultApiApiAccountsApiConnectedApplicationsConnAppIdPatchRequest) (ConnectedAppRespExt, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConnectedAppRespExt
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.AccountsApiConnectedApplicationsConnAppIdPatch")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/api/connectedApplications/{connAppId}"
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

type DefaultApiApiAccountsApiConnectedApplicationsConnAppIdScopesGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	connAppId string
}


func (r DefaultApiApiAccountsApiConnectedApplicationsConnAppIdScopesGetRequest) Execute() (InlineResponse2001, *_nethttp.Response, error) {
	return r.ApiService.AccountsApiConnectedApplicationsConnAppIdScopesGetExecute(r)
}

/*
 * AccountsApiConnectedApplicationsConnAppIdScopesGet Method for AccountsApiConnectedApplicationsConnAppIdScopesGet
 * Returns all scopes of a Connected App
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connAppId The ID of the connected app
 * @return DefaultApiApiAccountsApiConnectedApplicationsConnAppIdScopesGetRequest
 */
func (a *DefaultApiService) AccountsApiConnectedApplicationsConnAppIdScopesGet(ctx _context.Context, connAppId string) DefaultApiApiAccountsApiConnectedApplicationsConnAppIdScopesGetRequest {
	return DefaultApiApiAccountsApiConnectedApplicationsConnAppIdScopesGetRequest{
		ApiService: a,
		ctx: ctx,
		connAppId: connAppId,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse2001
 */
func (a *DefaultApiService) AccountsApiConnectedApplicationsConnAppIdScopesGetExecute(r DefaultApiApiAccountsApiConnectedApplicationsConnAppIdScopesGetRequest) (InlineResponse2001, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2001
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.AccountsApiConnectedApplicationsConnAppIdScopesGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/api/connectedApplications/{connAppId}/scopes"
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

type DefaultApiApiAccountsApiConnectedApplicationsConnAppIdScopesPutRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	connAppId string
	connectedAppScopesPutBody *ConnectedAppScopesPutBody
}

func (r DefaultApiApiAccountsApiConnectedApplicationsConnAppIdScopesPutRequest) ConnectedAppScopesPutBody(connectedAppScopesPutBody ConnectedAppScopesPutBody) DefaultApiApiAccountsApiConnectedApplicationsConnAppIdScopesPutRequest {
	r.connectedAppScopesPutBody = &connectedAppScopesPutBody
	return r
}

func (r DefaultApiApiAccountsApiConnectedApplicationsConnAppIdScopesPutRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.AccountsApiConnectedApplicationsConnAppIdScopesPutExecute(r)
}

/*
 * AccountsApiConnectedApplicationsConnAppIdScopesPut Method for AccountsApiConnectedApplicationsConnAppIdScopesPut
 * replace a Connected App scopes
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param connAppId The ID of the connected app
 * @return DefaultApiApiAccountsApiConnectedApplicationsConnAppIdScopesPutRequest
 */
func (a *DefaultApiService) AccountsApiConnectedApplicationsConnAppIdScopesPut(ctx _context.Context, connAppId string) DefaultApiApiAccountsApiConnectedApplicationsConnAppIdScopesPutRequest {
	return DefaultApiApiAccountsApiConnectedApplicationsConnAppIdScopesPutRequest{
		ApiService: a,
		ctx: ctx,
		connAppId: connAppId,
	}
}

/*
 * Execute executes the request
 */
func (a *DefaultApiService) AccountsApiConnectedApplicationsConnAppIdScopesPutExecute(r DefaultApiApiAccountsApiConnectedApplicationsConnAppIdScopesPutRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.AccountsApiConnectedApplicationsConnAppIdScopesPut")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/api/connectedApplications/{connAppId}/scopes"
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

type DefaultApiApiAccountsApiConnectedApplicationsGetRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
}


func (r DefaultApiApiAccountsApiConnectedApplicationsGetRequest) Execute() (InlineResponse200, *_nethttp.Response, error) {
	return r.ApiService.AccountsApiConnectedApplicationsGetExecute(r)
}

/*
 * AccountsApiConnectedApplicationsGet Method for AccountsApiConnectedApplicationsGet
 * Returns all connected apps
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return DefaultApiApiAccountsApiConnectedApplicationsGetRequest
 */
func (a *DefaultApiService) AccountsApiConnectedApplicationsGet(ctx _context.Context) DefaultApiApiAccountsApiConnectedApplicationsGetRequest {
	return DefaultApiApiAccountsApiConnectedApplicationsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return InlineResponse200
 */
func (a *DefaultApiService) AccountsApiConnectedApplicationsGetExecute(r DefaultApiApiAccountsApiConnectedApplicationsGetRequest) (InlineResponse200, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.AccountsApiConnectedApplicationsGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/api/connectedApplications"

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

type DefaultApiApiAccountsApiConnectedApplicationsPostRequest struct {
	ctx _context.Context
	ApiService *DefaultApiService
	connectedAppCore *ConnectedAppCore
}

func (r DefaultApiApiAccountsApiConnectedApplicationsPostRequest) ConnectedAppCore(connectedAppCore ConnectedAppCore) DefaultApiApiAccountsApiConnectedApplicationsPostRequest {
	r.connectedAppCore = &connectedAppCore
	return r
}

func (r DefaultApiApiAccountsApiConnectedApplicationsPostRequest) Execute() (ConnectedAppRespExt, *_nethttp.Response, error) {
	return r.ApiService.AccountsApiConnectedApplicationsPostExecute(r)
}

/*
 * AccountsApiConnectedApplicationsPost Method for AccountsApiConnectedApplicationsPost
 * create a Connected App
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return DefaultApiApiAccountsApiConnectedApplicationsPostRequest
 */
func (a *DefaultApiService) AccountsApiConnectedApplicationsPost(ctx _context.Context) DefaultApiApiAccountsApiConnectedApplicationsPostRequest {
	return DefaultApiApiAccountsApiConnectedApplicationsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ConnectedAppRespExt
 */
func (a *DefaultApiService) AccountsApiConnectedApplicationsPostExecute(r DefaultApiApiAccountsApiConnectedApplicationsPostRequest) (ConnectedAppRespExt, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConnectedAppRespExt
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.AccountsApiConnectedApplicationsPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/api/connectedApplications"

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
