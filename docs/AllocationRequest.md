# AllocationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The ID of the cryptocurrency you want to allocation. Supported values:  - USDC: &#x60;ETH_USDC&#x60;, &#x60;ARBITRUM_USDC&#x60;, &#x60;SOL_USDC&#x60;, &#x60;BASE_USDC&#x60;, &#x60;MATIC_USDC&#x60;, &#x60;BSC_USDC&#x60; - USDT: &#x60;TRON_USDT&#x60;, &#x60;ETH_USDT&#x60;, &#x60;ARBITRUM_USDT&#x60;, &#x60;SOL_USDT&#x60;, &#x60;BASE_USDT&#x60;, &#x60;MATIC_USDT&#x60;, &#x60;BSC_USDT&#x60;  | 
**Amount** | **string** | The allocation amount.  | 
**FromMerchantId** | Pointer to **string** | Only used in Merchant allocation type. The merchant ID.  | [optional] 
**ToMerchantId** | Pointer to **string** | Only used in Merchant allocation type. The merchant ID.  | [optional] 
**Description** | **string** |  | 

## Methods

### NewAllocationRequest

`func NewAllocationRequest(tokenId string, amount string, description string, ) *AllocationRequest`

NewAllocationRequest instantiates a new AllocationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocationRequestWithDefaults

`func NewAllocationRequestWithDefaults() *AllocationRequest`

NewAllocationRequestWithDefaults instantiates a new AllocationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *AllocationRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *AllocationRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *AllocationRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAmount

`func (o *AllocationRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AllocationRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AllocationRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetFromMerchantId

`func (o *AllocationRequest) GetFromMerchantId() string`

GetFromMerchantId returns the FromMerchantId field if non-nil, zero value otherwise.

### GetFromMerchantIdOk

`func (o *AllocationRequest) GetFromMerchantIdOk() (*string, bool)`

GetFromMerchantIdOk returns a tuple with the FromMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromMerchantId

`func (o *AllocationRequest) SetFromMerchantId(v string)`

SetFromMerchantId sets FromMerchantId field to given value.

### HasFromMerchantId

`func (o *AllocationRequest) HasFromMerchantId() bool`

HasFromMerchantId returns a boolean if a field has been set.

### GetToMerchantId

`func (o *AllocationRequest) GetToMerchantId() string`

GetToMerchantId returns the ToMerchantId field if non-nil, zero value otherwise.

### GetToMerchantIdOk

`func (o *AllocationRequest) GetToMerchantIdOk() (*string, bool)`

GetToMerchantIdOk returns a tuple with the ToMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToMerchantId

`func (o *AllocationRequest) SetToMerchantId(v string)`

SetToMerchantId sets ToMerchantId field to given value.

### HasToMerchantId

`func (o *AllocationRequest) HasToMerchantId() bool`

HasToMerchantId returns a boolean if a field has been set.

### GetDescription

`func (o *AllocationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AllocationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AllocationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


