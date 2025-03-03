# QueryGuardPubkey200ResponseAddressesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The wallet address. | 
**ChainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-chains). | 

## Methods

### NewQueryGuardPubkey200ResponseAddressesInner

`func NewQueryGuardPubkey200ResponseAddressesInner(address string, chainId string, ) *QueryGuardPubkey200ResponseAddressesInner`

NewQueryGuardPubkey200ResponseAddressesInner instantiates a new QueryGuardPubkey200ResponseAddressesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryGuardPubkey200ResponseAddressesInnerWithDefaults

`func NewQueryGuardPubkey200ResponseAddressesInnerWithDefaults() *QueryGuardPubkey200ResponseAddressesInner`

NewQueryGuardPubkey200ResponseAddressesInnerWithDefaults instantiates a new QueryGuardPubkey200ResponseAddressesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *QueryGuardPubkey200ResponseAddressesInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *QueryGuardPubkey200ResponseAddressesInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *QueryGuardPubkey200ResponseAddressesInner) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChainId

`func (o *QueryGuardPubkey200ResponseAddressesInner) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *QueryGuardPubkey200ResponseAddressesInner) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *QueryGuardPubkey200ResponseAddressesInner) SetChainId(v string)`

SetChainId sets ChainId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


