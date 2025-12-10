# WalletAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletAddressId** | **string** | The wallet address ID. | 
**Address** | **string** | The wallet address. | 
**ChainId** | **string** | The chain ID of the cryptocurrency. | 
**UpdatedTimestamp** | **int32** | The updated time of the wallet address, represented as a UNIX timestamp in seconds. | 

## Methods

### NewWalletAddress

`func NewWalletAddress(walletAddressId string, address string, chainId string, updatedTimestamp int32, ) *WalletAddress`

NewWalletAddress instantiates a new WalletAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletAddressWithDefaults

`func NewWalletAddressWithDefaults() *WalletAddress`

NewWalletAddressWithDefaults instantiates a new WalletAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletAddressId

`func (o *WalletAddress) GetWalletAddressId() string`

GetWalletAddressId returns the WalletAddressId field if non-nil, zero value otherwise.

### GetWalletAddressIdOk

`func (o *WalletAddress) GetWalletAddressIdOk() (*string, bool)`

GetWalletAddressIdOk returns a tuple with the WalletAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddressId

`func (o *WalletAddress) SetWalletAddressId(v string)`

SetWalletAddressId sets WalletAddressId field to given value.


### GetAddress

`func (o *WalletAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WalletAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WalletAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChainId

`func (o *WalletAddress) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *WalletAddress) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *WalletAddress) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetUpdatedTimestamp

`func (o *WalletAddress) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *WalletAddress) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *WalletAddress) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


