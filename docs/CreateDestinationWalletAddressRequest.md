# CreateDestinationWalletAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationId** | **string** | The destination ID. | 
**Address** | **string** | The wallet address. | 
**ChainId** | **string** | The chain ID of the cryptocurrency. | 

## Methods

### NewCreateDestinationWalletAddressRequest

`func NewCreateDestinationWalletAddressRequest(destinationId string, address string, chainId string, ) *CreateDestinationWalletAddressRequest`

NewCreateDestinationWalletAddressRequest instantiates a new CreateDestinationWalletAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDestinationWalletAddressRequestWithDefaults

`func NewCreateDestinationWalletAddressRequestWithDefaults() *CreateDestinationWalletAddressRequest`

NewCreateDestinationWalletAddressRequestWithDefaults instantiates a new CreateDestinationWalletAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationId

`func (o *CreateDestinationWalletAddressRequest) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *CreateDestinationWalletAddressRequest) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *CreateDestinationWalletAddressRequest) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetAddress

`func (o *CreateDestinationWalletAddressRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateDestinationWalletAddressRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateDestinationWalletAddressRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChainId

`func (o *CreateDestinationWalletAddressRequest) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *CreateDestinationWalletAddressRequest) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *CreateDestinationWalletAddressRequest) SetChainId(v string)`

SetChainId sets ChainId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


