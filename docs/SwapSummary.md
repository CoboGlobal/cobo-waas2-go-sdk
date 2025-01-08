# SwapSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalUsdValue** | **string** | The total USD value of the swap activities, represented as a string. | 
**ActivityCount** | **int32** | The total number of swap activities. | 

## Methods

### NewSwapSummary

`func NewSwapSummary(totalUsdValue string, activityCount int32, ) *SwapSummary`

NewSwapSummary instantiates a new SwapSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwapSummaryWithDefaults

`func NewSwapSummaryWithDefaults() *SwapSummary`

NewSwapSummaryWithDefaults instantiates a new SwapSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalUsdValue

`func (o *SwapSummary) GetTotalUsdValue() string`

GetTotalUsdValue returns the TotalUsdValue field if non-nil, zero value otherwise.

### GetTotalUsdValueOk

`func (o *SwapSummary) GetTotalUsdValueOk() (*string, bool)`

GetTotalUsdValueOk returns a tuple with the TotalUsdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsdValue

`func (o *SwapSummary) SetTotalUsdValue(v string)`

SetTotalUsdValue sets TotalUsdValue field to given value.


### GetActivityCount

`func (o *SwapSummary) GetActivityCount() int32`

GetActivityCount returns the ActivityCount field if non-nil, zero value otherwise.

### GetActivityCountOk

`func (o *SwapSummary) GetActivityCountOk() (*int32, bool)`

GetActivityCountOk returns a tuple with the ActivityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityCount

`func (o *SwapSummary) SetActivityCount(v int32)`

SetActivityCount sets ActivityCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


