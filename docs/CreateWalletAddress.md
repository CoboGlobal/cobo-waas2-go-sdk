# CreateWalletAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The wallet address. | 
**ChainId** | **string** | The chain ID of the address. | 

## Methods

### NewCreateWalletAddress

`func NewCreateWalletAddress(address string, chainId string, ) *CreateWalletAddress`

NewCreateWalletAddress instantiates a new CreateWalletAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWalletAddressWithDefaults

`func NewCreateWalletAddressWithDefaults() *CreateWalletAddress`

NewCreateWalletAddressWithDefaults instantiates a new CreateWalletAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *CreateWalletAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateWalletAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateWalletAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChainId

`func (o *CreateWalletAddress) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *CreateWalletAddress) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *CreateWalletAddress) SetChainId(v string)`

SetChainId sets ChainId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


