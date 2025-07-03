# \OrganizationsAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgInfo**](OrganizationsAPI.md#GetOrgInfo) | **Get** /organizations/info | Get organization information



## GetOrgInfo

> OrgInfo GetOrgInfo(ctx).Execute()

Get organization information



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
	resp, r, err := apiClient.OrganizationsAPI.GetOrgInfo(ctx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrgInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgInfo`: OrgInfo
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrgInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgInfoRequest struct via the builder pattern


### Return type

[**OrgInfo**](OrgInfo.md)

### Authorization

[CoboAuth](../README.md#CoboAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

