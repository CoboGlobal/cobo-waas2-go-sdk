# CreatePayoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a payout request. The request ID is provided by you and must be unique. | 
**SourceAccount** | **string** | The source account from which the payout will be made. - If the source account is a merchant account, provide the merchant&#39;s ID (e.g., \&quot;M1001\&quot;). - If the source account is the developer account, use the string &#x60;\&quot;developer\&quot;&#x60;.  | 
**PayoutChannel** | [**PayoutChannel**](PayoutChannel.md) |  | 
**PayoutParams** | [**[]PaymentPayoutParam**](PaymentPayoutParam.md) |  | 
**RecipientInfo** | [**PaymentPayoutRecipientInfo**](PaymentPayoutRecipientInfo.md) |  | 
**Remark** | Pointer to **string** | An optional note or comment about the payout for your internal reference. | [optional] 

## Methods

### NewCreatePayoutRequest

`func NewCreatePayoutRequest(requestId string, sourceAccount string, payoutChannel PayoutChannel, payoutParams []PaymentPayoutParam, recipientInfo PaymentPayoutRecipientInfo, ) *CreatePayoutRequest`

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


### GetSourceAccount

`func (o *CreatePayoutRequest) GetSourceAccount() string`

GetSourceAccount returns the SourceAccount field if non-nil, zero value otherwise.

### GetSourceAccountOk

`func (o *CreatePayoutRequest) GetSourceAccountOk() (*string, bool)`

GetSourceAccountOk returns a tuple with the SourceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAccount

`func (o *CreatePayoutRequest) SetSourceAccount(v string)`

SetSourceAccount sets SourceAccount field to given value.


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


### GetRecipientInfo

`func (o *CreatePayoutRequest) GetRecipientInfo() PaymentPayoutRecipientInfo`

GetRecipientInfo returns the RecipientInfo field if non-nil, zero value otherwise.

### GetRecipientInfoOk

`func (o *CreatePayoutRequest) GetRecipientInfoOk() (*PaymentPayoutRecipientInfo, bool)`

GetRecipientInfoOk returns a tuple with the RecipientInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientInfo

`func (o *CreatePayoutRequest) SetRecipientInfo(v PaymentPayoutRecipientInfo)`

SetRecipientInfo sets RecipientInfo field to given value.


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


