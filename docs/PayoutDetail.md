# PayoutDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayoutId** | **string** | The payout ID. | 
**Source** | Pointer to [**PayoutSource**](PayoutSource.md) |  | [optional] 
**DestinationCount** | **int32** | The destination count. | 
**TokenId** | **string** | The token ID. | 
**TotalAmount** | **string** | The total amount. | 
**Status** | [**PayoutStatus**](PayoutStatus.md) |  | 
**Description** | Pointer to **string** | The description. | [optional] 
**CreatedTimestamp** | **int32** | The created time of the payout, represented as a UNIX timestamp in seconds. | 
**UpdatedTimestamp** | Pointer to **int32** | The updated time of the payout, represented as a UNIX timestamp in seconds. | [optional] 
**Destinations** | Pointer to [**[]PayoutDestinationDetail**](PayoutDestinationDetail.md) | The destinations of the payout. | [optional] 

## Methods

### NewPayoutDetail

`func NewPayoutDetail(payoutId string, destinationCount int32, tokenId string, totalAmount string, status PayoutStatus, createdTimestamp int32, ) *PayoutDetail`

NewPayoutDetail instantiates a new PayoutDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutDetailWithDefaults

`func NewPayoutDetailWithDefaults() *PayoutDetail`

NewPayoutDetailWithDefaults instantiates a new PayoutDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayoutId

`func (o *PayoutDetail) GetPayoutId() string`

GetPayoutId returns the PayoutId field if non-nil, zero value otherwise.

### GetPayoutIdOk

`func (o *PayoutDetail) GetPayoutIdOk() (*string, bool)`

GetPayoutIdOk returns a tuple with the PayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutId

`func (o *PayoutDetail) SetPayoutId(v string)`

SetPayoutId sets PayoutId field to given value.


### GetSource

`func (o *PayoutDetail) GetSource() PayoutSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PayoutDetail) GetSourceOk() (*PayoutSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PayoutDetail) SetSource(v PayoutSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *PayoutDetail) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDestinationCount

`func (o *PayoutDetail) GetDestinationCount() int32`

GetDestinationCount returns the DestinationCount field if non-nil, zero value otherwise.

### GetDestinationCountOk

`func (o *PayoutDetail) GetDestinationCountOk() (*int32, bool)`

GetDestinationCountOk returns a tuple with the DestinationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCount

`func (o *PayoutDetail) SetDestinationCount(v int32)`

SetDestinationCount sets DestinationCount field to given value.


### GetTokenId

`func (o *PayoutDetail) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PayoutDetail) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PayoutDetail) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetTotalAmount

`func (o *PayoutDetail) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *PayoutDetail) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *PayoutDetail) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.


### GetStatus

`func (o *PayoutDetail) GetStatus() PayoutStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PayoutDetail) GetStatusOk() (*PayoutStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PayoutDetail) SetStatus(v PayoutStatus)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *PayoutDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PayoutDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PayoutDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PayoutDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *PayoutDetail) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *PayoutDetail) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *PayoutDetail) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *PayoutDetail) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *PayoutDetail) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *PayoutDetail) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *PayoutDetail) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.

### GetDestinations

`func (o *PayoutDetail) GetDestinations() []PayoutDestinationDetail`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *PayoutDetail) GetDestinationsOk() (*[]PayoutDestinationDetail, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *PayoutDetail) SetDestinations(v []PayoutDestinationDetail)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *PayoutDetail) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


