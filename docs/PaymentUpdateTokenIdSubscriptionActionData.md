# PaymentUpdateTokenIdSubscriptionActionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | [**PaymentSubscriptionActionType**](PaymentSubscriptionActionType.md) |  | 
**SubscriptionId** | **string** | The subscription id in cobo. | 
**TokenId** | **string** | The ID of the cryptocurrency you want to subscription. Supported values:  - USDC: &#x60;ETH_USDC&#x60;, &#x60;ARBITRUM_USDCOIN&#x60;, &#x60;BASE_USDC&#x60;, &#x60;MATIC_USDC2&#x60;, &#x60;BSC_USDC&#x60; - USDT: &#x60;ETH_USDT&#x60;, &#x60;ARBITRUM_USDT&#x60;, &#x60;BASE_USDT&#x60;, &#x60;MATIC_USDT&#x60;, &#x60;BSC_USDT&#x60;  | 
**Signature** | **string** | The signature for transaction. charge action is not required. | 
**Deadline** | **int32** | The signature deadline for transaction. charge action is not required. | 

## Methods

### NewPaymentUpdateTokenIdSubscriptionActionData

`func NewPaymentUpdateTokenIdSubscriptionActionData(actionType PaymentSubscriptionActionType, subscriptionId string, tokenId string, signature string, deadline int32, ) *PaymentUpdateTokenIdSubscriptionActionData`

NewPaymentUpdateTokenIdSubscriptionActionData instantiates a new PaymentUpdateTokenIdSubscriptionActionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentUpdateTokenIdSubscriptionActionDataWithDefaults

`func NewPaymentUpdateTokenIdSubscriptionActionDataWithDefaults() *PaymentUpdateTokenIdSubscriptionActionData`

NewPaymentUpdateTokenIdSubscriptionActionDataWithDefaults instantiates a new PaymentUpdateTokenIdSubscriptionActionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *PaymentUpdateTokenIdSubscriptionActionData) GetActionType() PaymentSubscriptionActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *PaymentUpdateTokenIdSubscriptionActionData) GetActionTypeOk() (*PaymentSubscriptionActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *PaymentUpdateTokenIdSubscriptionActionData) SetActionType(v PaymentSubscriptionActionType)`

SetActionType sets ActionType field to given value.


### GetSubscriptionId

`func (o *PaymentUpdateTokenIdSubscriptionActionData) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PaymentUpdateTokenIdSubscriptionActionData) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PaymentUpdateTokenIdSubscriptionActionData) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTokenId

`func (o *PaymentUpdateTokenIdSubscriptionActionData) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PaymentUpdateTokenIdSubscriptionActionData) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PaymentUpdateTokenIdSubscriptionActionData) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetSignature

`func (o *PaymentUpdateTokenIdSubscriptionActionData) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *PaymentUpdateTokenIdSubscriptionActionData) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *PaymentUpdateTokenIdSubscriptionActionData) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetDeadline

`func (o *PaymentUpdateTokenIdSubscriptionActionData) GetDeadline() int32`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *PaymentUpdateTokenIdSubscriptionActionData) GetDeadlineOk() (*int32, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *PaymentUpdateTokenIdSubscriptionActionData) SetDeadline(v int32)`

SetDeadline sets Deadline field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


