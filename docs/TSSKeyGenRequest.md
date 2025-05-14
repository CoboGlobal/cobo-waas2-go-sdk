# TSSKeyGenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Threshold** | Pointer to **int32** | The number of key share holders required to approve each operation in TSS key share group. | [optional] 
**NodeIds** | Pointer to **[]string** |  | [optional] 
**Curve** | Pointer to [**TSSCurve**](TSSCurve.md) |  | [optional] 
**TaskId** | Pointer to **string** | The task ID. | [optional] 
**BizTaskId** | Pointer to **string** | The business task ID. This field contains the TSS request ID. | [optional] 

## Methods

### NewTSSKeyGenRequest

`func NewTSSKeyGenRequest() *TSSKeyGenRequest`

NewTSSKeyGenRequest instantiates a new TSSKeyGenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSKeyGenRequestWithDefaults

`func NewTSSKeyGenRequestWithDefaults() *TSSKeyGenRequest`

NewTSSKeyGenRequestWithDefaults instantiates a new TSSKeyGenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreshold

`func (o *TSSKeyGenRequest) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *TSSKeyGenRequest) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *TSSKeyGenRequest) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *TSSKeyGenRequest) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetNodeIds

`func (o *TSSKeyGenRequest) GetNodeIds() []string`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *TSSKeyGenRequest) GetNodeIdsOk() (*[]string, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *TSSKeyGenRequest) SetNodeIds(v []string)`

SetNodeIds sets NodeIds field to given value.

### HasNodeIds

`func (o *TSSKeyGenRequest) HasNodeIds() bool`

HasNodeIds returns a boolean if a field has been set.

### GetCurve

`func (o *TSSKeyGenRequest) GetCurve() TSSCurve`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *TSSKeyGenRequest) GetCurveOk() (*TSSCurve, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *TSSKeyGenRequest) SetCurve(v TSSCurve)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *TSSKeyGenRequest) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### GetTaskId

`func (o *TSSKeyGenRequest) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TSSKeyGenRequest) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TSSKeyGenRequest) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *TSSKeyGenRequest) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetBizTaskId

`func (o *TSSKeyGenRequest) GetBizTaskId() string`

GetBizTaskId returns the BizTaskId field if non-nil, zero value otherwise.

### GetBizTaskIdOk

`func (o *TSSKeyGenRequest) GetBizTaskIdOk() (*string, bool)`

GetBizTaskIdOk returns a tuple with the BizTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizTaskId

`func (o *TSSKeyGenRequest) SetBizTaskId(v string)`

SetBizTaskId sets BizTaskId field to given value.

### HasBizTaskId

`func (o *TSSKeyGenRequest) HasBizTaskId() bool`

HasBizTaskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


