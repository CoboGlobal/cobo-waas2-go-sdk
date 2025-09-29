# \GraphQLAPI

All URIs are relative to *https://api.dev.cobo.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecuteGraphql**](GraphQLAPI.md#ExecuteGraphql) | **Post** /graphql | Execute a GraphQL query or mutation



## ExecuteGraphql

> GraphQLResponse ExecuteGraphql(ctx).GraphQLRequest(graphQLRequest).Execute()

Execute a GraphQL query or mutation



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
	graphQLRequest := *coboWaas2.NewGraphQLRequest("query {
  wallets {
    edges {
      node {
        id
        name
      }
    }
  }
}
")

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
	resp, r, err := apiClient.GraphQLAPI.ExecuteGraphql(ctx).GraphQLRequest(graphQLRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GraphQLAPI.ExecuteGraphql``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecuteGraphql`: GraphQLResponse
	fmt.Fprintf(os.Stdout, "Response from `GraphQLAPI.ExecuteGraphql`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteGraphqlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **graphQLRequest** | [**GraphQLRequest**](GraphQLRequest.md) | The request body to generate addresses within a specified wallet. | 

### Return type

[**GraphQLResponse**](GraphQLResponse.md)

### Authorization

[CoboNonce](../README.md#CoboNonce), [OAuth2](../README.md#OAuth2), [CoboAuth](../README.md#CoboAuth), [CoboSignature](../README.md#CoboSignature)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

