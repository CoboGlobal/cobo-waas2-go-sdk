# CreateBankWithdrawalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a bank withdrawal request. The request ID is provided by you and must be unique. | 
**SourceBankAccountId** | **string** | The source bank account ID. The destination bank account must be tagged as &#x60;VA&#x60;. Cobo uses the mapped VA account to initiate the withdrawal.  | 
**TargetBankAccountId** | **string** | The target bank account ID that receives the bank withdrawal. | 
**Currency** | **string** | The fiat currency of the bank withdrawal. | 
**Amount** | **string** | The bank withdrawal amount. | 
**Remark** | Pointer to **string** | The remark for the bank withdrawal. | [optional] 

## Methods

### NewCreateBankWithdrawalRequest

`func NewCreateBankWithdrawalRequest(requestId string, sourceBankAccountId string, targetBankAccountId string, currency string, amount string, ) *CreateBankWithdrawalRequest`

NewCreateBankWithdrawalRequest instantiates a new CreateBankWithdrawalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBankWithdrawalRequestWithDefaults

`func NewCreateBankWithdrawalRequestWithDefaults() *CreateBankWithdrawalRequest`

NewCreateBankWithdrawalRequestWithDefaults instantiates a new CreateBankWithdrawalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateBankWithdrawalRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateBankWithdrawalRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateBankWithdrawalRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetSourceBankAccountId

`func (o *CreateBankWithdrawalRequest) GetSourceBankAccountId() string`

GetSourceBankAccountId returns the SourceBankAccountId field if non-nil, zero value otherwise.

### GetSourceBankAccountIdOk

`func (o *CreateBankWithdrawalRequest) GetSourceBankAccountIdOk() (*string, bool)`

GetSourceBankAccountIdOk returns a tuple with the SourceBankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBankAccountId

`func (o *CreateBankWithdrawalRequest) SetSourceBankAccountId(v string)`

SetSourceBankAccountId sets SourceBankAccountId field to given value.


### GetTargetBankAccountId

`func (o *CreateBankWithdrawalRequest) GetTargetBankAccountId() string`

GetTargetBankAccountId returns the TargetBankAccountId field if non-nil, zero value otherwise.

### GetTargetBankAccountIdOk

`func (o *CreateBankWithdrawalRequest) GetTargetBankAccountIdOk() (*string, bool)`

GetTargetBankAccountIdOk returns a tuple with the TargetBankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBankAccountId

`func (o *CreateBankWithdrawalRequest) SetTargetBankAccountId(v string)`

SetTargetBankAccountId sets TargetBankAccountId field to given value.


### GetCurrency

`func (o *CreateBankWithdrawalRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateBankWithdrawalRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateBankWithdrawalRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAmount

`func (o *CreateBankWithdrawalRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateBankWithdrawalRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateBankWithdrawalRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetRemark

`func (o *CreateBankWithdrawalRequest) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *CreateBankWithdrawalRequest) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *CreateBankWithdrawalRequest) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *CreateBankWithdrawalRequest) HasRemark() bool`

HasRemark returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


