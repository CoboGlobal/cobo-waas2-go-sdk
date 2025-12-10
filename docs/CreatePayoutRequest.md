# CreatePayoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a settlement request. The request ID is provided by you and must be unique. | 
**PayoutChannel** | [**PayoutChannel**](PayoutChannel.md) |  | 
**SourceType** | [**PaymentSourceType**](PaymentSourceType.md) |  | 
**PayoutParams** | [**[]PaymentPayoutParam**](PaymentPayoutParam.md) |  | 
**BankAccountId** | Pointer to **string** | ｜ Only used in OffRamp payout channel. The ID of the bank account where the settled funds will be deposited. | [optional] 
**Currency** | Pointer to **string** | The fiat currency for the create payouts. | [optional] 
**Remark** | Pointer to **string** | The remark for the create payouts. | [optional] 

## Methods

### NewCreatePayoutRequest

`func NewCreatePayoutRequest(requestId string, payoutChannel PayoutChannel, sourceType PaymentSourceType, payoutParams []PaymentPayoutParam, ) *CreatePayoutRequest`

NewCreatePayoutRequest instantiates a new CreatePayoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePayoutRequestWithDefaults

`func NewCreatePayoutRequestWithDefaults() *CreatePayoutRequest`

NewCreatePayoutRequestWithDefaults instantiates a new CreatePayoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreatePayoutRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreatePayoutRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreatePayoutRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetPayoutChannel

`func (o *CreatePayoutRequest) GetPayoutChannel() PayoutChannel`

GetPayoutChannel returns the PayoutChannel field if non-nil, zero value otherwise.

### GetPayoutChannelOk

`func (o *CreatePayoutRequest) GetPayoutChannelOk() (*PayoutChannel, bool)`

GetPayoutChannelOk returns a tuple with the PayoutChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutChannel

`func (o *CreatePayoutRequest) SetPayoutChannel(v PayoutChannel)`

SetPayoutChannel sets PayoutChannel field to given value.


### GetSourceType

`func (o *CreatePayoutRequest) GetSourceType() PaymentSourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *CreatePayoutRequest) GetSourceTypeOk() (*PaymentSourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *CreatePayoutRequest) SetSourceType(v PaymentSourceType)`

SetSourceType sets SourceType field to given value.


### GetPayoutParams

`func (o *CreatePayoutRequest) GetPayoutParams() []PaymentPayoutParam`

GetPayoutParams returns the PayoutParams field if non-nil, zero value otherwise.

### GetPayoutParamsOk

`func (o *CreatePayoutRequest) GetPayoutParamsOk() (*[]PaymentPayoutParam, bool)`

GetPayoutParamsOk returns a tuple with the PayoutParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutParams

`func (o *CreatePayoutRequest) SetPayoutParams(v []PaymentPayoutParam)`

SetPayoutParams sets PayoutParams field to given value.


### GetBankAccountId

`func (o *CreatePayoutRequest) GetBankAccountId() string`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *CreatePayoutRequest) GetBankAccountIdOk() (*string, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *CreatePayoutRequest) SetBankAccountId(v string)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *CreatePayoutRequest) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### GetCurrency

`func (o *CreatePayoutRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreatePayoutRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreatePayoutRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreatePayoutRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetRemark

`func (o *CreatePayoutRequest) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *CreatePayoutRequest) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *CreatePayoutRequest) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *CreatePayoutRequest) HasRemark() bool`

HasRemark returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


