# CreateStakeActivityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The id of the wallet to stake. | 
**Address** | **string** | The staker wallet address. | 
**PoolId** | **string** | The id of the staking pool | 
**Amount** | **string** | The amount to stake | 
**Fee** | [**TransactionTransferFee**](TransactionTransferFee.md) |  | 
**Extra** | [**CreateStakeActivityExtra**](CreateStakeActivityExtra.md) |  | 
**Initiator** | Pointer to **string** | The initiator of the staking activity. | [optional] 

## Methods

### NewCreateStakeActivityRequest

`func NewCreateStakeActivityRequest(walletId string, address string, poolId string, amount string, fee TransactionTransferFee, extra CreateStakeActivityExtra, ) *CreateStakeActivityRequest`

NewCreateStakeActivityRequest instantiates a new CreateStakeActivityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStakeActivityRequestWithDefaults

`func NewCreateStakeActivityRequestWithDefaults() *CreateStakeActivityRequest`

NewCreateStakeActivityRequestWithDefaults instantiates a new CreateStakeActivityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *CreateStakeActivityRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *CreateStakeActivityRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *CreateStakeActivityRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *CreateStakeActivityRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateStakeActivityRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateStakeActivityRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPoolId

`func (o *CreateStakeActivityRequest) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *CreateStakeActivityRequest) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *CreateStakeActivityRequest) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetAmount

`func (o *CreateStakeActivityRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateStakeActivityRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateStakeActivityRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFee

`func (o *CreateStakeActivityRequest) GetFee() TransactionTransferFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CreateStakeActivityRequest) GetFeeOk() (*TransactionTransferFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CreateStakeActivityRequest) SetFee(v TransactionTransferFee)`

SetFee sets Fee field to given value.


### GetExtra

`func (o *CreateStakeActivityRequest) GetExtra() CreateStakeActivityExtra`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *CreateStakeActivityRequest) GetExtraOk() (*CreateStakeActivityExtra, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *CreateStakeActivityRequest) SetExtra(v CreateStakeActivityExtra)`

SetExtra sets Extra field to given value.


### GetInitiator

`func (o *CreateStakeActivityRequest) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *CreateStakeActivityRequest) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *CreateStakeActivityRequest) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *CreateStakeActivityRequest) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


