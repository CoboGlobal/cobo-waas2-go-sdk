/*
Cobo Wallet as a Service 2.0

Contact: help@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cobo_waas2

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// DevelopersAPIService DevelopersAPI service
type DevelopersAPIService service

type ApiGetApiKeyInfoRequest struct {
	ctx context.Context
	ApiService *DevelopersAPIService
}

func (r ApiGetApiKeyInfoRequest) Execute() (*GetApiKeyInfo200Response, *http.Response, error) {
	return r.ApiService.GetApiKeyInfoExecute(r)
}

/*
GetApiKeyInfo Get API key information

This operation retrieves the details of the API key that you are using.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetApiKeyInfoRequest
*/
func (a *DevelopersAPIService) GetApiKeyInfo(ctx context.Context) ApiGetApiKeyInfoRequest {
	return ApiGetApiKeyInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetApiKeyInfo200Response
func (a *DevelopersAPIService) GetApiKeyInfoExecute(r ApiGetApiKeyInfoRequest) (*GetApiKeyInfo200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetApiKeyInfo200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DevelopersAPIService.GetApiKeyInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/developers/api_key_info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListCallbackMessagesRequest struct {
	ctx context.Context
	ApiService *DevelopersAPIService
	callbackMessageIds *string
	requestIds *string
	transactionIds *string
	walletIds *string
	status *string
	limit *int32
	before *string
	after *string
}

// The IDs of the callback messages. Comma separated.
func (r ApiListCallbackMessagesRequest) CallbackMessageIds(callbackMessageIds string) ApiListCallbackMessagesRequest {
	r.callbackMessageIds = &callbackMessageIds
	return r
}

// The request_ids. Comma separated.
func (r ApiListCallbackMessagesRequest) RequestIds(requestIds string) ApiListCallbackMessagesRequest {
	r.requestIds = &requestIds
	return r
}

// The IDs of the transactions. Comma separated.
func (r ApiListCallbackMessagesRequest) TransactionIds(transactionIds string) ApiListCallbackMessagesRequest {
	r.transactionIds = &transactionIds
	return r
}

// The wallet ids of the transactions. Comma separated.
func (r ApiListCallbackMessagesRequest) WalletIds(walletIds string) ApiListCallbackMessagesRequest {
	r.walletIds = &walletIds
	return r
}

// The callback status. Possible values are &#x60;Approved&#x60;, &#x60;Denied&#x60; and &#x60;Failed&#x60;.
func (r ApiListCallbackMessagesRequest) Status(status string) ApiListCallbackMessagesRequest {
	r.status = &status
	return r
}

// The maximum number of objects to return. For most operations, the value range is [1, 50].
func (r ApiListCallbackMessagesRequest) Limit(limit int32) ApiListCallbackMessagesRequest {
	r.limit = &limit
	return r
}

// An object ID that serves as a starting point for retrieving data in reverse chronological order. For example, if you specify &#x60;before&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;, the request will retrieve a list of data objects that end before the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGmk1&#x60;. You can set this parameter to the value of &#x60;pagination.before&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned.  - If you set &#x60;before&#x60; to &#x60;infinity&#x60;, the last page of data is returned. 
func (r ApiListCallbackMessagesRequest) Before(before string) ApiListCallbackMessagesRequest {
	r.before = &before
	return r
}

// An object ID that acts as a starting point for retrieving data in chronological order. For example, if you specify &#x60;after&#x60; as &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;, the request will retrieve a list of data objects that start after the object with the object ID &#x60;RqeEoTkgKG5rpzqYzg2Hd3szmPoj2cE7w5jWwShz3C1vyGSAk&#x60;. You can set this parameter to the value of &#x60;pagination.after&#x60; in the response of the previous request.  - If you set both &#x60;after&#x60; and &#x60;before&#x60;, an error will occur.  - If you leave both &#x60;before&#x60; and &#x60;after&#x60; empty, the first page of data is returned. 
func (r ApiListCallbackMessagesRequest) After(after string) ApiListCallbackMessagesRequest {
	r.after = &after
	return r
}

func (r ApiListCallbackMessagesRequest) Execute() (*ListCallbackMessages200Response, *http.Response, error) {
	return r.ApiService.ListCallbackMessagesExecute(r)
}

/*
ListCallbackMessages List all callback messages

This operation retrieves all the callback messages.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListCallbackMessagesRequest
*/
func (a *DevelopersAPIService) ListCallbackMessages(ctx context.Context) ApiListCallbackMessagesRequest {
	return ApiListCallbackMessagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListCallbackMessages200Response
func (a *DevelopersAPIService) ListCallbackMessagesExecute(r ApiListCallbackMessagesRequest) (*ListCallbackMessages200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListCallbackMessages200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DevelopersAPIService.ListCallbackMessages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/developers/callback_messages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.callbackMessageIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "callback_message_ids", r.callbackMessageIds, "")
	}
	if r.requestIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "request_ids", r.requestIds, "")
	}
	if r.transactionIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "transaction_ids", r.transactionIds, "")
	}
	if r.walletIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "wallet_ids", r.walletIds, "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 10
		r.limit = &defaultValue
	}
	if r.before != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "before", r.before, "")
	}
	if r.after != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "after", r.after, "")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRetryCallbackMessageRequest struct {
	ctx context.Context
	ApiService *DevelopersAPIService
	messageId string
}

func (r ApiRetryCallbackMessageRequest) Execute() (*RetryWebhookEventById201Response, *http.Response, error) {
	return r.ApiService.RetryCallbackMessageExecute(r)
}

/*
RetryCallbackMessage Retry callback message

This operation retry a failed callback message.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param messageId The ID of the callback message.
 @return ApiRetryCallbackMessageRequest
*/
func (a *DevelopersAPIService) RetryCallbackMessage(ctx context.Context, messageId string) ApiRetryCallbackMessageRequest {
	return ApiRetryCallbackMessageRequest{
		ApiService: a,
		ctx: ctx,
		messageId: messageId,
	}
}

// Execute executes the request
//  @return RetryWebhookEventById201Response
func (a *DevelopersAPIService) RetryCallbackMessageExecute(r ApiRetryCallbackMessageRequest) (*RetryWebhookEventById201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RetryWebhookEventById201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DevelopersAPIService.RetryCallbackMessage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/developers/callback_message/{message_id}/retry"
	localVarPath = strings.Replace(localVarPath, "{"+"message_id"+"}", url.PathEscape(parameterValueToString(r.messageId, "messageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode >= 400 && localVarHTTPResponse.StatusCode < 500 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
