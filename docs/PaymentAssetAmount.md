# PaymentAssetAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | The amount of the asset, represented as a decimal string. | 
**TokenId** | **string** | The token ID of the asset. | 
**ChainId** | **string** | The chain ID of the asset. | 

## Methods

### NewPaymentAssetAmount

`func NewPaymentAssetAmount(amount string, tokenId string, chainId string, ) *PaymentAssetAmount`

NewPaymentAssetAmount instantiates a new PaymentAssetAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAssetAmountWithDefaults

`func NewPaymentAssetAmountWithDefaults() *PaymentAssetAmount`

NewPaymentAssetAmountWithDefaults instantiates a new PaymentAssetAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PaymentAssetAmount) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentAssetAmount) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentAssetAmount) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetTokenId

`func (o *PaymentAssetAmount) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentAssetAmount) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentAssetAmount) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChainId

`func (o *PaymentAssetAmount) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *PaymentAssetAmount) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *PaymentAssetAmount) SetChainId(v string)`

SetChainId sets ChainId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


