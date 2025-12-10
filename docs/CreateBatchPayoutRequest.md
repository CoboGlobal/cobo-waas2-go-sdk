# CreateBatchPayoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the batch payout. | 
**TokenId** | **string** | The ID of the cryptocurrency used for payout.  | 
**PayoutMode** | [**PayoutMode**](PayoutMode.md) |  | 
**UnlimitedTokenApproval** | Pointer to **bool** | Whether to apply unlimited token allowance. Only applicable when: - &#x60;payout_mode&#x60; is &#x60;SmartContract&#x60;  | [optional] [default to false]
**Source** | [**PayoutSource**](PayoutSource.md) |  | 
**Destinations** | [**[]PayoutDestination**](PayoutDestination.md) | The destinations of the payout. | 

## Methods

### NewCreateBatchPayoutRequest

`func NewCreateBatchPayoutRequest(description string, tokenId string, payoutMode PayoutMode, source PayoutSource, destinations []PayoutDestination, ) *CreateBatchPayoutRequest`

NewCreateBatchPayoutRequest instantiates a new CreateBatchPayoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBatchPayoutRequestWithDefaults

`func NewCreateBatchPayoutRequestWithDefaults() *CreateBatchPayoutRequest`

NewCreateBatchPayoutRequestWithDefaults instantiates a new CreateBatchPayoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateBatchPayoutRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateBatchPayoutRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateBatchPayoutRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTokenId

`func (o *CreateBatchPayoutRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *CreateBatchPayoutRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *CreateBatchPayoutRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetPayoutMode

`func (o *CreateBatchPayoutRequest) GetPayoutMode() PayoutMode`

GetPayoutMode returns the PayoutMode field if non-nil, zero value otherwise.

### GetPayoutModeOk

`func (o *CreateBatchPayoutRequest) GetPayoutModeOk() (*PayoutMode, bool)`

GetPayoutModeOk returns a tuple with the PayoutMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutMode

`func (o *CreateBatchPayoutRequest) SetPayoutMode(v PayoutMode)`

SetPayoutMode sets PayoutMode field to given value.


### GetUnlimitedTokenApproval

`func (o *CreateBatchPayoutRequest) GetUnlimitedTokenApproval() bool`

GetUnlimitedTokenApproval returns the UnlimitedTokenApproval field if non-nil, zero value otherwise.

### GetUnlimitedTokenApprovalOk

`func (o *CreateBatchPayoutRequest) GetUnlimitedTokenApprovalOk() (*bool, bool)`

GetUnlimitedTokenApprovalOk returns a tuple with the UnlimitedTokenApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlimitedTokenApproval

`func (o *CreateBatchPayoutRequest) SetUnlimitedTokenApproval(v bool)`

SetUnlimitedTokenApproval sets UnlimitedTokenApproval field to given value.

### HasUnlimitedTokenApproval

`func (o *CreateBatchPayoutRequest) HasUnlimitedTokenApproval() bool`

HasUnlimitedTokenApproval returns a boolean if a field has been set.

### GetSource

`func (o *CreateBatchPayoutRequest) GetSource() PayoutSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateBatchPayoutRequest) GetSourceOk() (*PayoutSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateBatchPayoutRequest) SetSource(v PayoutSource)`

SetSource sets Source field to given value.


### GetDestinations

`func (o *CreateBatchPayoutRequest) GetDestinations() []PayoutDestination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *CreateBatchPayoutRequest) GetDestinationsOk() (*[]PayoutDestination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *CreateBatchPayoutRequest) SetDestinations(v []PayoutDestination)`

SetDestinations sets Destinations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


