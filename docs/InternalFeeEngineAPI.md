# \InternalFeeEngineAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCommissionFee**](InternalFeeEngineAPI.md#GetCommissionFee) | **Get** /internal/commission_fee | Get commission fee



## GetCommissionFee

> []CommissionFeeDetail GetCommissionFee(ctx).RequestId(requestId).BusinessTypeId(businessTypeId).Amount(amount).ChainId(chainId).TokenId(tokenId).WalletId(walletId).WalletType(walletType).RequestType(requestType).StrategyContext(strategyContext).Execute()

Get commission fee



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    coboWaas2 "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2"
    "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2/crypto"
)

func main() {
	requestId := "requestId_example"
	businessTypeId := int64(789)
	amount := int64(789)
	chainId := "chainId_example"
	tokenId := "tokenId_example"
	walletId := "walletId_example"
	walletType := int32(56)
	requestType := coboWaas2.TransactionRequestTypeParams(-999)
	strategyContext := map[string]interface{}{"key": interface{}(123)}

	configuration := coboWaas2.NewConfiguration()
	// Initialize the API client
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.Background()

    // Select the development environment. To use the production environment, replace coboWaas2.DevEnv with coboWaas2.ProdEnv
	ctx = context.WithValue(ctx, coboWaas2.ContextEnv, coboWaas2.DevEnv)
    // Replace `<YOUR_PRIVATE_KEY>` with your private key
	ctx = context.WithValue(ctx, coboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_PRIVATE_KEY>",
	})
	resp, r, err := apiClient.InternalFeeEngineAPI.GetCommissionFee(ctx).RequestId(requestId).BusinessTypeId(businessTypeId).Amount(amount).ChainId(chainId).TokenId(tokenId).WalletId(walletId).WalletType(walletType).RequestType(requestType).StrategyContext(strategyContext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalFeeEngineAPI.GetCommissionFee``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommissionFee`: []CommissionFeeDetail
	fmt.Fprintf(os.Stdout, "Response from `InternalFeeEngineAPI.GetCommissionFee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCommissionFeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestId** | **string** | Unique request identifier | 
 **businessTypeId** | **int64** |  | 
 **amount** | **int64** | Transaction amount in smallest unit | 
 **chainId** | **string** | Blockchain chain ID | 
 **tokenId** | **string** | Token ID | 
 **walletId** | **string** | Wallet ID | 
 **walletType** | **int32** | Wallet type | 
 **requestType** | [**TransactionRequestTypeParams**](TransactionRequestTypeParams.md) |  | 
 **strategyContext** | **map[string]interface{}** | Strategy context parameters.  | 

### Return type

[**[]CommissionFeeDetail**](CommissionFeeDetail.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

