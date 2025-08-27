# SweepToAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The wallet address. | 
**ChainId** | **string** | The chain ID, which is the unique identifier of a blockchain. You can retrieve the IDs of all the chains you can use by calling [List enabled chains](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-chains). | 
**Status** | Pointer to [**SweepToAddressStatus**](SweepToAddressStatus.md) |  | [optional] 

## Methods

### NewSweepToAddress

`func NewSweepToAddress(address string, chainId string, ) *SweepToAddress`

NewSweepToAddress instantiates a new SweepToAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSweepToAddressWithDefaults

`func NewSweepToAddressWithDefaults() *SweepToAddress`

NewSweepToAddressWithDefaults instantiates a new SweepToAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SweepToAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SweepToAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SweepToAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChainId

`func (o *SweepToAddress) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *SweepToAddress) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *SweepToAddress) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetStatus

`func (o *SweepToAddress) GetStatus() SweepToAddressStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SweepToAddress) GetStatusOk() (*SweepToAddressStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SweepToAddress) SetStatus(v SweepToAddressStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SweepToAddress) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


