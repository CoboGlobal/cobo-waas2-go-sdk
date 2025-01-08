# TravelRuleWithdrawRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The transaction ID. | 
**TravelRuleInfo** | [**TravelRuleWithdrawRequestTravelRuleInfo**](TravelRuleWithdrawRequestTravelRuleInfo.md) |  | 

## Methods

### NewTravelRuleWithdrawRequest

`func NewTravelRuleWithdrawRequest(transactionId string, travelRuleInfo TravelRuleWithdrawRequestTravelRuleInfo, ) *TravelRuleWithdrawRequest`

NewTravelRuleWithdrawRequest instantiates a new TravelRuleWithdrawRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTravelRuleWithdrawRequestWithDefaults

`func NewTravelRuleWithdrawRequestWithDefaults() *TravelRuleWithdrawRequest`

NewTravelRuleWithdrawRequestWithDefaults instantiates a new TravelRuleWithdrawRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *TravelRuleWithdrawRequest) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TravelRuleWithdrawRequest) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TravelRuleWithdrawRequest) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetTravelRuleInfo

`func (o *TravelRuleWithdrawRequest) GetTravelRuleInfo() TravelRuleWithdrawRequestTravelRuleInfo`

GetTravelRuleInfo returns the TravelRuleInfo field if non-nil, zero value otherwise.

### GetTravelRuleInfoOk

`func (o *TravelRuleWithdrawRequest) GetTravelRuleInfoOk() (*TravelRuleWithdrawRequestTravelRuleInfo, bool)`

GetTravelRuleInfoOk returns a tuple with the TravelRuleInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelRuleInfo

`func (o *TravelRuleWithdrawRequest) SetTravelRuleInfo(v TravelRuleWithdrawRequestTravelRuleInfo)`

SetTravelRuleInfo sets TravelRuleInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


