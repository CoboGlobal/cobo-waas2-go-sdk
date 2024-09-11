# BookkeepingSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalTransactionCount** | **int32** | Total transaction count. | 
**TotalInflowValue** | **string** | The USD value of the inflow. | 
**TotalOutflowValue** | **string** | The USD value of the outflow. | 
**TotalFeeValue** | Pointer to **string** | The USD value of the fee. | [optional] 

## Methods

### NewBookkeepingSummary

`func NewBookkeepingSummary(totalTransactionCount int32, totalInflowValue string, totalOutflowValue string, ) *BookkeepingSummary`

NewBookkeepingSummary instantiates a new BookkeepingSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookkeepingSummaryWithDefaults

`func NewBookkeepingSummaryWithDefaults() *BookkeepingSummary`

NewBookkeepingSummaryWithDefaults instantiates a new BookkeepingSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalTransactionCount

`func (o *BookkeepingSummary) GetTotalTransactionCount() int32`

GetTotalTransactionCount returns the TotalTransactionCount field if non-nil, zero value otherwise.

### GetTotalTransactionCountOk

`func (o *BookkeepingSummary) GetTotalTransactionCountOk() (*int32, bool)`

GetTotalTransactionCountOk returns a tuple with the TotalTransactionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTransactionCount

`func (o *BookkeepingSummary) SetTotalTransactionCount(v int32)`

SetTotalTransactionCount sets TotalTransactionCount field to given value.


### GetTotalInflowValue

`func (o *BookkeepingSummary) GetTotalInflowValue() string`

GetTotalInflowValue returns the TotalInflowValue field if non-nil, zero value otherwise.

### GetTotalInflowValueOk

`func (o *BookkeepingSummary) GetTotalInflowValueOk() (*string, bool)`

GetTotalInflowValueOk returns a tuple with the TotalInflowValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInflowValue

`func (o *BookkeepingSummary) SetTotalInflowValue(v string)`

SetTotalInflowValue sets TotalInflowValue field to given value.


### GetTotalOutflowValue

`func (o *BookkeepingSummary) GetTotalOutflowValue() string`

GetTotalOutflowValue returns the TotalOutflowValue field if non-nil, zero value otherwise.

### GetTotalOutflowValueOk

`func (o *BookkeepingSummary) GetTotalOutflowValueOk() (*string, bool)`

GetTotalOutflowValueOk returns a tuple with the TotalOutflowValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOutflowValue

`func (o *BookkeepingSummary) SetTotalOutflowValue(v string)`

SetTotalOutflowValue sets TotalOutflowValue field to given value.


### GetTotalFeeValue

`func (o *BookkeepingSummary) GetTotalFeeValue() string`

GetTotalFeeValue returns the TotalFeeValue field if non-nil, zero value otherwise.

### GetTotalFeeValueOk

`func (o *BookkeepingSummary) GetTotalFeeValueOk() (*string, bool)`

GetTotalFeeValueOk returns a tuple with the TotalFeeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFeeValue

`func (o *BookkeepingSummary) SetTotalFeeValue(v string)`

SetTotalFeeValue sets TotalFeeValue field to given value.

### HasTotalFeeValue

`func (o *BookkeepingSummary) HasTotalFeeValue() bool`

HasTotalFeeValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


