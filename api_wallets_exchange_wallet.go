/*
Cobo Wallet as a Service 2.0

Cobo WaaS 2.0 enables you to programmatically access Cobo's full suite of crypto wallet technologies with powerful and flexible access controls.  # Wallet technologies - Custodial Wallet - MPC Wallet - Smart Contract Wallet (Based on Safe{Wallet}) - Exchange Wallet  # Risk Control technologies - Workflow - Access Control List (ACL)  # Risk Control targets - Wallet Management   - User/team and their permission management   - Risk control configurations, e.g. whitelist, blacklist, rate-limiting etc. - Blockchain Interaction   - Crypto transfer   - Smart Contract Invocation  # Important HTTPS only. RESTful, resource oriented  # Get Started Set up your APIs or get authorization  # Authentication and Authorization CoboAuth  # Request and Response application/json  # Error Handling  ### Common error codes | Error Code | Description | | -- | -- |  ### API-specific error codes For error codes that are dedicated to a specific API, see the Error codes section in each API specification, for example, /v3/wallets.  # Rate and Usage Limiting  # Idempotent Request  # Pagination # Support [Developer Hub](https://cobo.com/developers) 

API version: 1.0.0
Contact: support@cobo.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package CoboWaas2

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// WalletsExchangeWalletAPIService WalletsExchangeWalletAPI service
type WalletsExchangeWalletAPIService service

type ApiGetExchangeSupportedAssetsRequest struct {
	ctx context.Context
	ApiService *WalletsExchangeWalletAPIService
	exchangeId ExchangeId
	limit *int32
	before *string
	after *string
}

// size of page to return (pagination)
func (r ApiGetExchangeSupportedAssetsRequest) Limit(limit int32) ApiGetExchangeSupportedAssetsRequest {
	r.limit = &limit
	return r
}

// Cursor string received from previous request
func (r ApiGetExchangeSupportedAssetsRequest) Before(before string) ApiGetExchangeSupportedAssetsRequest {
	r.before = &before
	return r
}

// Cursor string received from previous request
func (r ApiGetExchangeSupportedAssetsRequest) After(after string) ApiGetExchangeSupportedAssetsRequest {
	r.after = &after
	return r
}

func (r ApiGetExchangeSupportedAssetsRequest) Execute() (*GetAssets200Response, *http.Response, error) {
	return r.ApiService.GetExchangeSupportedAssetsExecute(r)
}

/*
GetExchangeSupportedAssets List the supported assets by exchange id

Retrieve a list of supported asset.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param exchangeId Exchange ID to query
 @return ApiGetExchangeSupportedAssetsRequest
*/
func (a *WalletsExchangeWalletAPIService) GetExchangeSupportedAssets(ctx context.Context, exchangeId ExchangeId) ApiGetExchangeSupportedAssetsRequest {
	return ApiGetExchangeSupportedAssetsRequest{
		ApiService: a,
		ctx: ctx,
		exchangeId: exchangeId,
	}
}

// Execute executes the request
//  @return GetAssets200Response
func (a *WalletsExchangeWalletAPIService) GetExchangeSupportedAssetsExecute(r ApiGetExchangeSupportedAssetsRequest) (*GetAssets200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetAssets200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletsExchangeWalletAPIService.GetExchangeSupportedAssets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wallets/exchanges/{exchange_id}/supported_assets"
	localVarPath = strings.Replace(localVarPath, "{"+"exchange_id"+"}", url.PathEscape(parameterValueToString(r.exchangeId, "exchangeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 10
		r.limit = &defaultValue
	}
	if r.before != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "before", r.before, "")
	} else {
		var defaultValue string = ""
		r.before = &defaultValue
	}
	if r.after != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "after", r.after, "")
	} else {
		var defaultValue string = ""
		r.after = &defaultValue
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CoboAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["BIZ-API-KEY"] = key
			}
		}
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
		if localVarHTTPResponse.StatusCode == 400 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiGetExchangeSupportedChainsRequest struct {
	ctx context.Context
	ApiService *WalletsExchangeWalletAPIService
	exchangeId ExchangeId
	assetId *string
	limit *int32
	before *string
	after *string
}

// Unique id of the asset
func (r ApiGetExchangeSupportedChainsRequest) AssetId(assetId string) ApiGetExchangeSupportedChainsRequest {
	r.assetId = &assetId
	return r
}

// size of page to return (pagination)
func (r ApiGetExchangeSupportedChainsRequest) Limit(limit int32) ApiGetExchangeSupportedChainsRequest {
	r.limit = &limit
	return r
}

// Cursor string received from previous request
func (r ApiGetExchangeSupportedChainsRequest) Before(before string) ApiGetExchangeSupportedChainsRequest {
	r.before = &before
	return r
}

// Cursor string received from previous request
func (r ApiGetExchangeSupportedChainsRequest) After(after string) ApiGetExchangeSupportedChainsRequest {
	r.after = &after
	return r
}

func (r ApiGetExchangeSupportedChainsRequest) Execute() ([]ChainInfo, *http.Response, error) {
	return r.ApiService.GetExchangeSupportedChainsExecute(r)
}

/*
GetExchangeSupportedChains List the supported chains by exchange id and asset id

Retrieve a list of supported chains.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param exchangeId Exchange ID to query
 @return ApiGetExchangeSupportedChainsRequest
*/
func (a *WalletsExchangeWalletAPIService) GetExchangeSupportedChains(ctx context.Context, exchangeId ExchangeId) ApiGetExchangeSupportedChainsRequest {
	return ApiGetExchangeSupportedChainsRequest{
		ApiService: a,
		ctx: ctx,
		exchangeId: exchangeId,
	}
}

// Execute executes the request
//  @return []ChainInfo
func (a *WalletsExchangeWalletAPIService) GetExchangeSupportedChainsExecute(r ApiGetExchangeSupportedChainsRequest) ([]ChainInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ChainInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletsExchangeWalletAPIService.GetExchangeSupportedChains")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wallets/exchanges/{exchange_id}/assets/supported_chains"
	localVarPath = strings.Replace(localVarPath, "{"+"exchange_id"+"}", url.PathEscape(parameterValueToString(r.exchangeId, "exchangeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.assetId == nil {
		return localVarReturnValue, nil, reportError("assetId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "asset_id", r.assetId, "")
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 10
		r.limit = &defaultValue
	}
	if r.before != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "before", r.before, "")
	} else {
		var defaultValue string = ""
		r.before = &defaultValue
	}
	if r.after != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "after", r.after, "")
	} else {
		var defaultValue string = ""
		r.after = &defaultValue
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CoboAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["BIZ-API-KEY"] = key
			}
		}
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
		if localVarHTTPResponse.StatusCode == 400 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiGetExchangeWalletAssetBalancesRequest struct {
	ctx context.Context
	ApiService *WalletsExchangeWalletAPIService
	walletId string
	subWalletId *string
	assetId *string
	limit *int32
	before *string
	after *string
}

// Unique id of the wallet
func (r ApiGetExchangeWalletAssetBalancesRequest) SubWalletId(subWalletId string) ApiGetExchangeWalletAssetBalancesRequest {
	r.subWalletId = &subWalletId
	return r
}

// Unique id of the asset
func (r ApiGetExchangeWalletAssetBalancesRequest) AssetId(assetId string) ApiGetExchangeWalletAssetBalancesRequest {
	r.assetId = &assetId
	return r
}

// size of page to return (pagination)
func (r ApiGetExchangeWalletAssetBalancesRequest) Limit(limit int32) ApiGetExchangeWalletAssetBalancesRequest {
	r.limit = &limit
	return r
}

// Cursor string received from previous request
func (r ApiGetExchangeWalletAssetBalancesRequest) Before(before string) ApiGetExchangeWalletAssetBalancesRequest {
	r.before = &before
	return r
}

// Cursor string received from previous request
func (r ApiGetExchangeWalletAssetBalancesRequest) After(after string) ApiGetExchangeWalletAssetBalancesRequest {
	r.after = &after
	return r
}

func (r ApiGetExchangeWalletAssetBalancesRequest) Execute() (*GetExchangeWalletAssetBalances200Response, *http.Response, error) {
	return r.ApiService.GetExchangeWalletAssetBalancesExecute(r)
}

/*
GetExchangeWalletAssetBalances List the asset balance in exchange wallet

Retrieve a list of asset balance in exchange wallet.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param walletId Unique id of the wallet
 @return ApiGetExchangeWalletAssetBalancesRequest
*/
func (a *WalletsExchangeWalletAPIService) GetExchangeWalletAssetBalances(ctx context.Context, walletId string) ApiGetExchangeWalletAssetBalancesRequest {
	return ApiGetExchangeWalletAssetBalancesRequest{
		ApiService: a,
		ctx: ctx,
		walletId: walletId,
	}
}

// Execute executes the request
//  @return GetExchangeWalletAssetBalances200Response
func (a *WalletsExchangeWalletAPIService) GetExchangeWalletAssetBalancesExecute(r ApiGetExchangeWalletAssetBalancesRequest) (*GetExchangeWalletAssetBalances200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetExchangeWalletAssetBalances200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletsExchangeWalletAPIService.GetExchangeWalletAssetBalances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wallets/exchanges/{wallet_id}/assets"
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_id"+"}", url.PathEscape(parameterValueToString(r.walletId, "walletId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subWalletId == nil {
		return localVarReturnValue, nil, reportError("subWalletId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "sub_wallet_id", r.subWalletId, "")
	if r.assetId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "asset_id", r.assetId, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 10
		r.limit = &defaultValue
	}
	if r.before != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "before", r.before, "")
	} else {
		var defaultValue string = ""
		r.before = &defaultValue
	}
	if r.after != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "after", r.after, "")
	} else {
		var defaultValue string = ""
		r.after = &defaultValue
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CoboAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["BIZ-API-KEY"] = key
			}
		}
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
		if localVarHTTPResponse.StatusCode == 400 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiLinkSubAccountsByWalletIdRequest struct {
	ctx context.Context
	ApiService *WalletsExchangeWalletAPIService
	walletId string
	linkSubAccountsByWalletIdRequest *LinkSubAccountsByWalletIdRequest
}

// Request body for linking subaccounts
func (r ApiLinkSubAccountsByWalletIdRequest) LinkSubAccountsByWalletIdRequest(linkSubAccountsByWalletIdRequest LinkSubAccountsByWalletIdRequest) ApiLinkSubAccountsByWalletIdRequest {
	r.linkSubAccountsByWalletIdRequest = &linkSubAccountsByWalletIdRequest
	return r
}

func (r ApiLinkSubAccountsByWalletIdRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.LinkSubAccountsByWalletIdExecute(r)
}

/*
LinkSubAccountsByWalletId Link exchange sub accounts by wallet id

Link exchange sub accounts.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param walletId Unique id of the wallet
 @return ApiLinkSubAccountsByWalletIdRequest
*/
func (a *WalletsExchangeWalletAPIService) LinkSubAccountsByWalletId(ctx context.Context, walletId string) ApiLinkSubAccountsByWalletIdRequest {
	return ApiLinkSubAccountsByWalletIdRequest{
		ApiService: a,
		ctx: ctx,
		walletId: walletId,
	}
}

// Execute executes the request
//  @return []string
func (a *WalletsExchangeWalletAPIService) LinkSubAccountsByWalletIdExecute(r ApiLinkSubAccountsByWalletIdRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletsExchangeWalletAPIService.LinkSubAccountsByWalletId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wallets/{wallet_id}/exchanges/subaccounts"
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_id"+"}", url.PathEscape(parameterValueToString(r.walletId, "walletId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.linkSubAccountsByWalletIdRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CoboAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["BIZ-API-KEY"] = key
			}
		}
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
		if localVarHTTPResponse.StatusCode == 400 {
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

type ApiListExchangesRequest struct {
	ctx context.Context
	ApiService *WalletsExchangeWalletAPIService
}

func (r ApiListExchangesRequest) Execute() ([]ListExchanges200ResponseInner, *http.Response, error) {
	return r.ApiService.ListExchangesExecute(r)
}

/*
ListExchanges List of exchanges

Retrieve a list of exchanges.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListExchangesRequest
*/
func (a *WalletsExchangeWalletAPIService) ListExchanges(ctx context.Context) ApiListExchangesRequest {
	return ApiListExchangesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ListExchanges200ResponseInner
func (a *WalletsExchangeWalletAPIService) ListExchangesExecute(r ApiListExchangesRequest) ([]ListExchanges200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ListExchanges200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletsExchangeWalletAPIService.ListExchanges")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wallets/exchanges/settings"

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CoboAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["BIZ-API-KEY"] = key
			}
		}
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
		if localVarHTTPResponse.StatusCode == 401 {
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
		if localVarHTTPResponse.StatusCode == 403 {
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
		if localVarHTTPResponse.StatusCode == 429 {
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

type ApiListSubAccountsByApikeyRequest struct {
	ctx context.Context
	ApiService *WalletsExchangeWalletAPIService
	exchangeId ExchangeId
	apikey *string
	secret *string
	passphrase *string
}

// The API Key for the exchange
func (r ApiListSubAccountsByApikeyRequest) Apikey(apikey string) ApiListSubAccountsByApikeyRequest {
	r.apikey = &apikey
	return r
}

// The API Secret for the exchange.
func (r ApiListSubAccountsByApikeyRequest) Secret(secret string) ApiListSubAccountsByApikeyRequest {
	r.secret = &secret
	return r
}

// The API passphrase for the exchange wallet.
func (r ApiListSubAccountsByApikeyRequest) Passphrase(passphrase string) ApiListSubAccountsByApikeyRequest {
	r.passphrase = &passphrase
	return r
}

func (r ApiListSubAccountsByApikeyRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.ListSubAccountsByApikeyExecute(r)
}

/*
ListSubAccountsByApikey List exchange sub accounts by apikey

Retrieve a list of exchange sub accounts.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param exchangeId Exchange ID to query
 @return ApiListSubAccountsByApikeyRequest
*/
func (a *WalletsExchangeWalletAPIService) ListSubAccountsByApikey(ctx context.Context, exchangeId ExchangeId) ApiListSubAccountsByApikeyRequest {
	return ApiListSubAccountsByApikeyRequest{
		ApiService: a,
		ctx: ctx,
		exchangeId: exchangeId,
	}
}

// Execute executes the request
//  @return []string
func (a *WalletsExchangeWalletAPIService) ListSubAccountsByApikeyExecute(r ApiListSubAccountsByApikeyRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletsExchangeWalletAPIService.ListSubAccountsByApikey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wallets/exchanges/{exchange_id}/subaccounts"
	localVarPath = strings.Replace(localVarPath, "{"+"exchange_id"+"}", url.PathEscape(parameterValueToString(r.exchangeId, "exchangeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.apikey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "apikey", r.apikey, "")
	}
	if r.secret != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "secret", r.secret, "")
	}
	if r.passphrase != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "passphrase", r.passphrase, "")
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CoboAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["BIZ-API-KEY"] = key
			}
		}
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
		if localVarHTTPResponse.StatusCode == 400 {
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

type ApiListSubAccountsByWalletIdRequest struct {
	ctx context.Context
	ApiService *WalletsExchangeWalletAPIService
	walletId string
}

func (r ApiListSubAccountsByWalletIdRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.ListSubAccountsByWalletIdExecute(r)
}

/*
ListSubAccountsByWalletId List exchange sub accounts by wallet id

Retrieve a list of exchange sub accounts.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param walletId Unique id of the wallet
 @return ApiListSubAccountsByWalletIdRequest
*/
func (a *WalletsExchangeWalletAPIService) ListSubAccountsByWalletId(ctx context.Context, walletId string) ApiListSubAccountsByWalletIdRequest {
	return ApiListSubAccountsByWalletIdRequest{
		ApiService: a,
		ctx: ctx,
		walletId: walletId,
	}
}

// Execute executes the request
//  @return []string
func (a *WalletsExchangeWalletAPIService) ListSubAccountsByWalletIdExecute(r ApiListSubAccountsByWalletIdRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WalletsExchangeWalletAPIService.ListSubAccountsByWalletId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wallets/{wallet_id}/exchanges/subaccounts"
	localVarPath = strings.Replace(localVarPath, "{"+"wallet_id"+"}", url.PathEscape(parameterValueToString(r.walletId, "walletId")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CoboAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["BIZ-API-KEY"] = key
			}
		}
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
		if localVarHTTPResponse.StatusCode == 400 {
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
		if localVarHTTPResponse.StatusCode == 404 {
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
