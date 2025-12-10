# CreateCounterpartyWalletAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CounterpartyId** | **string** | The counterparty ID. | 
**Address** | **string** | The wallet address. | 
**ChainId** | **string** | The chain ID of the cryptocurrency. | 

## Methods

### NewCreateCounterpartyWalletAddressRequest

`func NewCreateCounterpartyWalletAddressRequest(counterpartyId string, address string, chainId string, ) *CreateCounterpartyWalletAddressRequest`

NewCreateCounterpartyWalletAddressRequest instantiates a new CreateCounterpartyWalletAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCounterpartyWalletAddressRequestWithDefaults

`func NewCreateCounterpartyWalletAddressRequestWithDefaults() *CreateCounterpartyWalletAddressRequest`

NewCreateCounterpartyWalletAddressRequestWithDefaults instantiates a new CreateCounterpartyWalletAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounterpartyId

`func (o *CreateCounterpartyWalletAddressRequest) GetCounterpartyId() string`

GetCounterpartyId returns the CounterpartyId field if non-nil, zero value otherwise.

### GetCounterpartyIdOk

`func (o *CreateCounterpartyWalletAddressRequest) GetCounterpartyIdOk() (*string, bool)`

GetCounterpartyIdOk returns a tuple with the CounterpartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyId

`func (o *CreateCounterpartyWalletAddressRequest) SetCounterpartyId(v string)`

SetCounterpartyId sets CounterpartyId field to given value.


### GetAddress

`func (o *CreateCounterpartyWalletAddressRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateCounterpartyWalletAddressRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateCounterpartyWalletAddressRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChainId

`func (o *CreateCounterpartyWalletAddressRequest) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *CreateCounterpartyWalletAddressRequest) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *CreateCounterpartyWalletAddressRequest) SetChainId(v string)`

SetChainId sets ChainId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


