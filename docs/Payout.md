# Payout

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

## Methods

### NewPayout

`func NewPayout(payoutId string, destinationCount int32, tokenId string, totalAmount string, status PayoutStatus, createdTimestamp int32, ) *Payout`

NewPayout instantiates a new Payout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutWithDefaults

`func NewPayoutWithDefaults() *Payout`

NewPayoutWithDefaults instantiates a new Payout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayoutId

`func (o *Payout) GetPayoutId() string`

GetPayoutId returns the PayoutId field if non-nil, zero value otherwise.

### GetPayoutIdOk

`func (o *Payout) GetPayoutIdOk() (*string, bool)`

GetPayoutIdOk returns a tuple with the PayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutId

`func (o *Payout) SetPayoutId(v string)`

SetPayoutId sets PayoutId field to given value.


### GetSource

`func (o *Payout) GetSource() PayoutSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Payout) GetSourceOk() (*PayoutSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Payout) SetSource(v PayoutSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *Payout) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDestinationCount

`func (o *Payout) GetDestinationCount() int32`

GetDestinationCount returns the DestinationCount field if non-nil, zero value otherwise.

### GetDestinationCountOk

`func (o *Payout) GetDestinationCountOk() (*int32, bool)`

GetDestinationCountOk returns a tuple with the DestinationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCount

`func (o *Payout) SetDestinationCount(v int32)`

SetDestinationCount sets DestinationCount field to given value.


### GetTokenId

`func (o *Payout) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *Payout) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *Payout) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetTotalAmount

`func (o *Payout) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *Payout) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *Payout) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.


### GetStatus

`func (o *Payout) GetStatus() PayoutStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Payout) GetStatusOk() (*PayoutStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Payout) SetStatus(v PayoutStatus)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *Payout) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Payout) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Payout) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Payout) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *Payout) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *Payout) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *Payout) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


