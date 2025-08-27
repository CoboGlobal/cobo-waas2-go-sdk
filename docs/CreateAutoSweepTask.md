# CreateAutoSweepTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The wallet ID. | 
**TokenId** | **string** | The token ID of the swept token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 

## Methods

### NewCreateAutoSweepTask

`func NewCreateAutoSweepTask(walletId string, tokenId string, ) *CreateAutoSweepTask`

NewCreateAutoSweepTask instantiates a new CreateAutoSweepTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutoSweepTaskWithDefaults

`func NewCreateAutoSweepTaskWithDefaults() *CreateAutoSweepTask`

NewCreateAutoSweepTaskWithDefaults instantiates a new CreateAutoSweepTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *CreateAutoSweepTask) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *CreateAutoSweepTask) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *CreateAutoSweepTask) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetTokenId

`func (o *CreateAutoSweepTask) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *CreateAutoSweepTask) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *CreateAutoSweepTask) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


