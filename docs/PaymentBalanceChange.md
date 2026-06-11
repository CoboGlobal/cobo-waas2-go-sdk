# PaymentBalanceChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** | The source ID of the balance change. | 
**SourceType** | [**PaymentBalanceChangeSourceType**](PaymentBalanceChangeSourceType.md) |  | 
**TokenId** | **string** | The token ID of the balance change. | 
**Amount** | **string** | The balance change amount, truncated to two decimal places and represented as a numeric string. | 
**AmountRaw** | **string** | The balance change amount in the token&#39;s decimal precision, represented as a numeric string. | 
**BalanceBefore** | **string** | The account balance before the balance change, truncated to two decimal places and represented as a numeric string. | 
**BalanceBeforeRaw** | **string** | The account balance before the balance change in the token&#39;s decimal precision, represented as a numeric string. | 
**BalanceAfter** | **string** | The account balance after the balance change, truncated to two decimal places and represented as a numeric string. | 
**BalanceAfterRaw** | **string** | The account balance after the balance change in the token&#39;s decimal precision, represented as a numeric string. | 
**FlowDirection** | [**PaymentBalanceFlowDirection**](PaymentBalanceFlowDirection.md) |  | 
**CreatedTimestamp** | **int64** | The time when the balance change was created, represented as a UNIX timestamp in seconds. | 

## Methods

### NewPaymentBalanceChange

`func NewPaymentBalanceChange(sourceId string, sourceType PaymentBalanceChangeSourceType, tokenId string, amount string, amountRaw string, balanceBefore string, balanceBeforeRaw string, balanceAfter string, balanceAfterRaw string, flowDirection PaymentBalanceFlowDirection, createdTimestamp int64, ) *PaymentBalanceChange`

NewPaymentBalanceChange instantiates a new PaymentBalanceChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentBalanceChangeWithDefaults

`func NewPaymentBalanceChangeWithDefaults() *PaymentBalanceChange`

NewPaymentBalanceChangeWithDefaults instantiates a new PaymentBalanceChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *PaymentBalanceChange) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *PaymentBalanceChange) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *PaymentBalanceChange) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceType

`func (o *PaymentBalanceChange) GetSourceType() PaymentBalanceChangeSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *PaymentBalanceChange) GetSourceTypeOk() (*PaymentBalanceChangeSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *PaymentBalanceChange) SetSourceType(v PaymentBalanceChangeSourceType)`

SetSourceType sets SourceType field to given value.


### GetTokenId

`func (o *PaymentBalanceChange) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentBalanceChange) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentBalanceChange) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetAmount

`func (o *PaymentBalanceChange) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentBalanceChange) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentBalanceChange) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAmountRaw

`func (o *PaymentBalanceChange) GetAmountRaw() string`

GetAmountRaw returns the AmountRaw field if non-nil, zero value otherwise.

### GetAmountRawOk

`func (o *PaymentBalanceChange) GetAmountRawOk() (*string, bool)`

GetAmountRawOk returns a tuple with the AmountRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRaw

`func (o *PaymentBalanceChange) SetAmountRaw(v string)`

SetAmountRaw sets AmountRaw field to given value.


### GetBalanceBefore

`func (o *PaymentBalanceChange) GetBalanceBefore() string`

GetBalanceBefore returns the BalanceBefore field if non-nil, zero value otherwise.

### GetBalanceBeforeOk

`func (o *PaymentBalanceChange) GetBalanceBeforeOk() (*string, bool)`

GetBalanceBeforeOk returns a tuple with the BalanceBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceBefore

`func (o *PaymentBalanceChange) SetBalanceBefore(v string)`

SetBalanceBefore sets BalanceBefore field to given value.


### GetBalanceBeforeRaw

`func (o *PaymentBalanceChange) GetBalanceBeforeRaw() string`

GetBalanceBeforeRaw returns the BalanceBeforeRaw field if non-nil, zero value otherwise.

### GetBalanceBeforeRawOk

`func (o *PaymentBalanceChange) GetBalanceBeforeRawOk() (*string, bool)`

GetBalanceBeforeRawOk returns a tuple with the BalanceBeforeRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceBeforeRaw

`func (o *PaymentBalanceChange) SetBalanceBeforeRaw(v string)`

SetBalanceBeforeRaw sets BalanceBeforeRaw field to given value.


### GetBalanceAfter

`func (o *PaymentBalanceChange) GetBalanceAfter() string`

GetBalanceAfter returns the BalanceAfter field if non-nil, zero value otherwise.

### GetBalanceAfterOk

`func (o *PaymentBalanceChange) GetBalanceAfterOk() (*string, bool)`

GetBalanceAfterOk returns a tuple with the BalanceAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAfter

`func (o *PaymentBalanceChange) SetBalanceAfter(v string)`

SetBalanceAfter sets BalanceAfter field to given value.


### GetBalanceAfterRaw

`func (o *PaymentBalanceChange) GetBalanceAfterRaw() string`

GetBalanceAfterRaw returns the BalanceAfterRaw field if non-nil, zero value otherwise.

### GetBalanceAfterRawOk

`func (o *PaymentBalanceChange) GetBalanceAfterRawOk() (*string, bool)`

GetBalanceAfterRawOk returns a tuple with the BalanceAfterRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAfterRaw

`func (o *PaymentBalanceChange) SetBalanceAfterRaw(v string)`

SetBalanceAfterRaw sets BalanceAfterRaw field to given value.


### GetFlowDirection

`func (o *PaymentBalanceChange) GetFlowDirection() PaymentBalanceFlowDirection`

GetFlowDirection returns the FlowDirection field if non-nil, zero value otherwise.

### GetFlowDirectionOk

`func (o *PaymentBalanceChange) GetFlowDirectionOk() (*PaymentBalanceFlowDirection, bool)`

GetFlowDirectionOk returns a tuple with the FlowDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDirection

`func (o *PaymentBalanceChange) SetFlowDirection(v PaymentBalanceFlowDirection)`

SetFlowDirection sets FlowDirection field to given value.


### GetCreatedTimestamp

`func (o *PaymentBalanceChange) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PaymentBalanceChange) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PaymentBalanceChange) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


