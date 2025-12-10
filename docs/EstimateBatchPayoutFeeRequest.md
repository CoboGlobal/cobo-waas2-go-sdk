# EstimateBatchPayoutFeeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The ID of the cryptocurrency used for payout.  | 
**PayoutMode** | [**PayoutMode**](PayoutMode.md) |  | 
**Source** | [**PayoutSource**](PayoutSource.md) |  | 
**Destinations** | [**[]PayoutDestination**](PayoutDestination.md) | The destinations of the payout. | 
**RbfType** | Pointer to [**PayoutRbfType**](PayoutRbfType.md) |  | [optional] 
**ReplacedPayoutId** | Pointer to **string** | The ID of the payout that this payout replaced. | [optional] 

## Methods

### NewEstimateBatchPayoutFeeRequest

`func NewEstimateBatchPayoutFeeRequest(tokenId string, payoutMode PayoutMode, source PayoutSource, destinations []PayoutDestination, ) *EstimateBatchPayoutFeeRequest`

NewEstimateBatchPayoutFeeRequest instantiates a new EstimateBatchPayoutFeeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateBatchPayoutFeeRequestWithDefaults

`func NewEstimateBatchPayoutFeeRequestWithDefaults() *EstimateBatchPayoutFeeRequest`

NewEstimateBatchPayoutFeeRequestWithDefaults instantiates a new EstimateBatchPayoutFeeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *EstimateBatchPayoutFeeRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EstimateBatchPayoutFeeRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EstimateBatchPayoutFeeRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetPayoutMode

`func (o *EstimateBatchPayoutFeeRequest) GetPayoutMode() PayoutMode`

GetPayoutMode returns the PayoutMode field if non-nil, zero value otherwise.

### GetPayoutModeOk

`func (o *EstimateBatchPayoutFeeRequest) GetPayoutModeOk() (*PayoutMode, bool)`

GetPayoutModeOk returns a tuple with the PayoutMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutMode

`func (o *EstimateBatchPayoutFeeRequest) SetPayoutMode(v PayoutMode)`

SetPayoutMode sets PayoutMode field to given value.


### GetSource

`func (o *EstimateBatchPayoutFeeRequest) GetSource() PayoutSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EstimateBatchPayoutFeeRequest) GetSourceOk() (*PayoutSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EstimateBatchPayoutFeeRequest) SetSource(v PayoutSource)`

SetSource sets Source field to given value.


### GetDestinations

`func (o *EstimateBatchPayoutFeeRequest) GetDestinations() []PayoutDestination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *EstimateBatchPayoutFeeRequest) GetDestinationsOk() (*[]PayoutDestination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *EstimateBatchPayoutFeeRequest) SetDestinations(v []PayoutDestination)`

SetDestinations sets Destinations field to given value.


### GetRbfType

`func (o *EstimateBatchPayoutFeeRequest) GetRbfType() PayoutRbfType`

GetRbfType returns the RbfType field if non-nil, zero value otherwise.

### GetRbfTypeOk

`func (o *EstimateBatchPayoutFeeRequest) GetRbfTypeOk() (*PayoutRbfType, bool)`

GetRbfTypeOk returns a tuple with the RbfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbfType

`func (o *EstimateBatchPayoutFeeRequest) SetRbfType(v PayoutRbfType)`

SetRbfType sets RbfType field to given value.

### HasRbfType

`func (o *EstimateBatchPayoutFeeRequest) HasRbfType() bool`

HasRbfType returns a boolean if a field has been set.

### GetReplacedPayoutId

`func (o *EstimateBatchPayoutFeeRequest) GetReplacedPayoutId() string`

GetReplacedPayoutId returns the ReplacedPayoutId field if non-nil, zero value otherwise.

### GetReplacedPayoutIdOk

`func (o *EstimateBatchPayoutFeeRequest) GetReplacedPayoutIdOk() (*string, bool)`

GetReplacedPayoutIdOk returns a tuple with the ReplacedPayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedPayoutId

`func (o *EstimateBatchPayoutFeeRequest) SetReplacedPayoutId(v string)`

SetReplacedPayoutId sets ReplacedPayoutId field to given value.

### HasReplacedPayoutId

`func (o *EstimateBatchPayoutFeeRequest) HasReplacedPayoutId() bool`

HasReplacedPayoutId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


