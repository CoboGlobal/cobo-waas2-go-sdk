# MpcStakeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | [**StakeSourceType**](StakeSourceType.md) |  | 
**WalletId** | **string** | The wallet ID. | 
**Address** | **string** | The wallet address. | 

## Methods

### NewMpcStakeSource

`func NewMpcStakeSource(sourceType StakeSourceType, walletId string, address string, ) *MpcStakeSource`

NewMpcStakeSource instantiates a new MpcStakeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMpcStakeSourceWithDefaults

`func NewMpcStakeSourceWithDefaults() *MpcStakeSource`

NewMpcStakeSourceWithDefaults instantiates a new MpcStakeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *MpcStakeSource) GetSourceType() StakeSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *MpcStakeSource) GetSourceTypeOk() (*StakeSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *MpcStakeSource) SetSourceType(v StakeSourceType)`

SetSourceType sets SourceType field to given value.


### GetWalletId

`func (o *MpcStakeSource) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *MpcStakeSource) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *MpcStakeSource) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *MpcStakeSource) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MpcStakeSource) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MpcStakeSource) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


