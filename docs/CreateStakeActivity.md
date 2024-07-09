# CreateStakeActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The id of the wallet to stake. | 
**Address** | **string** | The staker wallet address. | 
**PoolId** | **string** | The id of the staking pool | 
**Amount** | **string** | The amount to stake | 
**Fee** | [**TransactionTransferFee**](TransactionTransferFee.md) |  | 
**Extra** | [**CreateStakeActivityExtra**](CreateStakeActivityExtra.md) |  | 

## Methods

### NewCreateStakeActivity

`func NewCreateStakeActivity(walletId string, address string, poolId string, amount string, fee TransactionTransferFee, extra CreateStakeActivityExtra, ) *CreateStakeActivity`

NewCreateStakeActivity instantiates a new CreateStakeActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStakeActivityWithDefaults

`func NewCreateStakeActivityWithDefaults() *CreateStakeActivity`

NewCreateStakeActivityWithDefaults instantiates a new CreateStakeActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *CreateStakeActivity) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *CreateStakeActivity) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *CreateStakeActivity) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *CreateStakeActivity) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateStakeActivity) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateStakeActivity) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPoolId

`func (o *CreateStakeActivity) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *CreateStakeActivity) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *CreateStakeActivity) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetAmount

`func (o *CreateStakeActivity) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateStakeActivity) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateStakeActivity) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFee

`func (o *CreateStakeActivity) GetFee() TransactionTransferFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CreateStakeActivity) GetFeeOk() (*TransactionTransferFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CreateStakeActivity) SetFee(v TransactionTransferFee)`

SetFee sets Fee field to given value.


### GetExtra

`func (o *CreateStakeActivity) GetExtra() CreateStakeActivityExtra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *CreateStakeActivity) GetExtraOk() (*CreateStakeActivityExtra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *CreateStakeActivity) SetExtra(v CreateStakeActivityExtra)`

SetExtra sets Extra field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


