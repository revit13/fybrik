/*
Consolidate Services

Description of all APIs

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

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

// GPGKeyServiceApiService GPGKeyServiceApi service
type GPGKeyServiceApiService service

type ApiGPGKeyServiceCreateRequest struct {
	ctx        _context.Context
	ApiService *GPGKeyServiceApiService
	body       *V1alpha1GnuPGPublicKey
	upsert     *bool
}

// Raw key data of the GPG key(s) to create
func (r ApiGPGKeyServiceCreateRequest) Body(body V1alpha1GnuPGPublicKey) ApiGPGKeyServiceCreateRequest {
	r.body = &body
	return r
}

// Whether to upsert already existing public keys.
func (r ApiGPGKeyServiceCreateRequest) Upsert(upsert bool) ApiGPGKeyServiceCreateRequest {
	r.upsert = &upsert
	return r
}

func (r ApiGPGKeyServiceCreateRequest) Execute() (GpgkeyGnuPGPublicKeyCreateResponse, *_nethttp.Response, error) {
	return r.ApiService.GPGKeyServiceCreateExecute(r)
}

/*
GPGKeyServiceCreate Create one or more GPG public keys in the server's configuration

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGPGKeyServiceCreateRequest
*/
func (a *GPGKeyServiceApiService) GPGKeyServiceCreate(ctx _context.Context) ApiGPGKeyServiceCreateRequest {
	return ApiGPGKeyServiceCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GpgkeyGnuPGPublicKeyCreateResponse
func (a *GPGKeyServiceApiService) GPGKeyServiceCreateExecute(r ApiGPGKeyServiceCreateRequest) (GpgkeyGnuPGPublicKeyCreateResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue GpgkeyGnuPGPublicKeyCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GPGKeyServiceApiService.GPGKeyServiceCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/gpgkeys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.upsert != nil {
		localVarQueryParams.Add("upsert", parameterToString(*r.upsert, ""))
	}
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
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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
		var v RuntimeError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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

type ApiGPGKeyServiceDeleteRequest struct {
	ctx        _context.Context
	ApiService *GPGKeyServiceApiService
	keyID      *string
}

// The GPG key ID to query for.
func (r ApiGPGKeyServiceDeleteRequest) KeyID(keyID string) ApiGPGKeyServiceDeleteRequest {
	r.keyID = &keyID
	return r
}

func (r ApiGPGKeyServiceDeleteRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.GPGKeyServiceDeleteExecute(r)
}

/*
GPGKeyServiceDelete Delete specified GPG public key from the server's configuration

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGPGKeyServiceDeleteRequest
*/
func (a *GPGKeyServiceApiService) GPGKeyServiceDelete(ctx _context.Context) ApiGPGKeyServiceDeleteRequest {
	return ApiGPGKeyServiceDeleteRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return map[string]interface{}
func (a *GPGKeyServiceApiService) GPGKeyServiceDeleteExecute(r ApiGPGKeyServiceDeleteRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GPGKeyServiceApiService.GPGKeyServiceDelete")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/gpgkeys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.keyID != nil {
		localVarQueryParams.Add("keyID", parameterToString(*r.keyID, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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
		var v RuntimeError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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

type ApiGPGKeyServiceGetRequest struct {
	ctx        _context.Context
	ApiService *GPGKeyServiceApiService
	keyID      string
}

func (r ApiGPGKeyServiceGetRequest) Execute() (V1alpha1GnuPGPublicKey, *_nethttp.Response, error) {
	return r.ApiService.GPGKeyServiceGetExecute(r)
}

/*
GPGKeyServiceGet Get information about specified GPG public key from the server

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param keyID The GPG key ID to query for
	@return ApiGPGKeyServiceGetRequest
*/
func (a *GPGKeyServiceApiService) GPGKeyServiceGet(ctx _context.Context, keyID string) ApiGPGKeyServiceGetRequest {
	return ApiGPGKeyServiceGetRequest{
		ApiService: a,
		ctx:        ctx,
		keyID:      keyID,
	}
}

// Execute executes the request
//
//	@return V1alpha1GnuPGPublicKey
func (a *GPGKeyServiceApiService) GPGKeyServiceGetExecute(r ApiGPGKeyServiceGetRequest) (V1alpha1GnuPGPublicKey, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue V1alpha1GnuPGPublicKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GPGKeyServiceApiService.GPGKeyServiceGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/gpgkeys/{keyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", _neturl.PathEscape(parameterToString(r.keyID, "")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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
		var v RuntimeError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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

type ApiGPGKeyServiceListRequest struct {
	ctx        _context.Context
	ApiService *GPGKeyServiceApiService
	keyID      *string
}

// The GPG key ID to query for.
func (r ApiGPGKeyServiceListRequest) KeyID(keyID string) ApiGPGKeyServiceListRequest {
	r.keyID = &keyID
	return r
}

func (r ApiGPGKeyServiceListRequest) Execute() (V1alpha1GnuPGPublicKeyList, *_nethttp.Response, error) {
	return r.ApiService.GPGKeyServiceListExecute(r)
}

/*
GPGKeyServiceList List all available repository certificates

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGPGKeyServiceListRequest
*/
func (a *GPGKeyServiceApiService) GPGKeyServiceList(ctx _context.Context) ApiGPGKeyServiceListRequest {
	return ApiGPGKeyServiceListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return V1alpha1GnuPGPublicKeyList
func (a *GPGKeyServiceApiService) GPGKeyServiceListExecute(r ApiGPGKeyServiceListRequest) (V1alpha1GnuPGPublicKeyList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue V1alpha1GnuPGPublicKeyList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GPGKeyServiceApiService.GPGKeyServiceList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/gpgkeys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.keyID != nil {
		localVarQueryParams.Add("keyID", parameterToString(*r.keyID, ""))
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
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
		var v RuntimeError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
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
