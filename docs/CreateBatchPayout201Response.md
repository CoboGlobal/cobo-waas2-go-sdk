# CreateBatchPayout201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayoutId** | **string** | The payout ID. | 
**PayoutMode** | [**PayoutMode**](PayoutMode.md) |  | 
**Status** | [**PayoutStatus**](PayoutStatus.md) |  | 
**EstimatedTotalFee** | **string** | The total fee of the payout. | 
**CreatedTimestamp** | **int32** | The created time of the payout, represented as a UNIX timestamp in seconds. | 

## Methods

### NewCreateBatchPayout201Response

`func NewCreateBatchPayout201Response(payoutId string, payoutMode PayoutMode, status PayoutStatus, estimatedTotalFee string, createdTimestamp int32, ) *CreateBatchPayout201Response`

NewCreateBatchPayout201Response instantiates a new CreateBatchPayout201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBatchPayout201ResponseWithDefaults

`func NewCreateBatchPayout201ResponseWithDefaults() *CreateBatchPayout201Response`

NewCreateBatchPayout201ResponseWithDefaults instantiates a new CreateBatchPayout201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayoutId

`func (o *CreateBatchPayout201Response) GetPayoutId() string`

GetPayoutId returns the PayoutId field if non-nil, zero value otherwise.

### GetPayoutIdOk

`func (o *CreateBatchPayout201Response) GetPayoutIdOk() (*string, bool)`

GetPayoutIdOk returns a tuple with the PayoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutId

`func (o *CreateBatchPayout201Response) SetPayoutId(v string)`

SetPayoutId sets PayoutId field to given value.


### GetPayoutMode

`func (o *CreateBatchPayout201Response) GetPayoutMode() PayoutMode`

GetPayoutMode returns the PayoutMode field if non-nil, zero value otherwise.

### GetPayoutModeOk

`func (o *CreateBatchPayout201Response) GetPayoutModeOk() (*PayoutMode, bool)`

GetPayoutModeOk returns a tuple with the PayoutMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutMode

`func (o *CreateBatchPayout201Response) SetPayoutMode(v PayoutMode)`

SetPayoutMode sets PayoutMode field to given value.


### GetStatus

`func (o *CreateBatchPayout201Response) GetStatus() PayoutStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateBatchPayout201Response) GetStatusOk() (*PayoutStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateBatchPayout201Response) SetStatus(v PayoutStatus)`

SetStatus sets Status field to given value.


### GetEstimatedTotalFee

`func (o *CreateBatchPayout201Response) GetEstimatedTotalFee() string`

GetEstimatedTotalFee returns the EstimatedTotalFee field if non-nil, zero value otherwise.

### GetEstimatedTotalFeeOk

`func (o *CreateBatchPayout201Response) GetEstimatedTotalFeeOk() (*string, bool)`

GetEstimatedTotalFeeOk returns a tuple with the EstimatedTotalFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTotalFee

`func (o *CreateBatchPayout201Response) SetEstimatedTotalFee(v string)`

SetEstimatedTotalFee sets EstimatedTotalFee field to given value.


### GetCreatedTimestamp

`func (o *CreateBatchPayout201Response) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *CreateBatchPayout201Response) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *CreateBatchPayout201Response) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


