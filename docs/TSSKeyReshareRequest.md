# TSSKeyReshareRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldGroupId** | Pointer to **string** | The old TSS key share group ID. | [optional] 
**RootPubKey** | Pointer to **string** | The The old TSS key share group&#39;s root extended public key. | [optional] 
**Curve** | Pointer to [**TSSCurve**](TSSCurve.md) |  | [optional] 
**UsedNodeIds** | Pointer to **[]string** |  | [optional] 
**OldThreshold** | Pointer to **int32** | The number of key share holders required to approve each operation in the old TSS key share group. | [optional] 
**NewThreshold** | Pointer to **int32** | The number of key share holders required to approve each operation in the new TSS key share group. | [optional] 
**NewNodeIds** | Pointer to **[]string** |  | [optional] 
**TaskId** | Pointer to **string** | The task ID. | [optional] 
**BizTaskId** | Pointer to **string** | The business task ID. This field contains the TSS request ID. | [optional] 

## Methods

### NewTSSKeyReshareRequest

`func NewTSSKeyReshareRequest() *TSSKeyReshareRequest`

NewTSSKeyReshareRequest instantiates a new TSSKeyReshareRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSKeyReshareRequestWithDefaults

`func NewTSSKeyReshareRequestWithDefaults() *TSSKeyReshareRequest`

NewTSSKeyReshareRequestWithDefaults instantiates a new TSSKeyReshareRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldGroupId

`func (o *TSSKeyReshareRequest) GetOldGroupId() string`

GetOldGroupId returns the OldGroupId field if non-nil, zero value otherwise.

### GetOldGroupIdOk

`func (o *TSSKeyReshareRequest) GetOldGroupIdOk() (*string, bool)`

GetOldGroupIdOk returns a tuple with the OldGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldGroupId

`func (o *TSSKeyReshareRequest) SetOldGroupId(v string)`

SetOldGroupId sets OldGroupId field to given value.

### HasOldGroupId

`func (o *TSSKeyReshareRequest) HasOldGroupId() bool`

HasOldGroupId returns a boolean if a field has been set.

### GetRootPubKey

`func (o *TSSKeyReshareRequest) GetRootPubKey() string`

GetRootPubKey returns the RootPubKey field if non-nil, zero value otherwise.

### GetRootPubKeyOk

`func (o *TSSKeyReshareRequest) GetRootPubKeyOk() (*string, bool)`

GetRootPubKeyOk returns a tuple with the RootPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPubKey

`func (o *TSSKeyReshareRequest) SetRootPubKey(v string)`

SetRootPubKey sets RootPubKey field to given value.

### HasRootPubKey

`func (o *TSSKeyReshareRequest) HasRootPubKey() bool`

HasRootPubKey returns a boolean if a field has been set.

### GetCurve

`func (o *TSSKeyReshareRequest) GetCurve() TSSCurve`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *TSSKeyReshareRequest) GetCurveOk() (*TSSCurve, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *TSSKeyReshareRequest) SetCurve(v TSSCurve)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *TSSKeyReshareRequest) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### GetUsedNodeIds

`func (o *TSSKeyReshareRequest) GetUsedNodeIds() []string`

GetUsedNodeIds returns the UsedNodeIds field if non-nil, zero value otherwise.

### GetUsedNodeIdsOk

`func (o *TSSKeyReshareRequest) GetUsedNodeIdsOk() (*[]string, bool)`

GetUsedNodeIdsOk returns a tuple with the UsedNodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedNodeIds

`func (o *TSSKeyReshareRequest) SetUsedNodeIds(v []string)`

SetUsedNodeIds sets UsedNodeIds field to given value.

### HasUsedNodeIds

`func (o *TSSKeyReshareRequest) HasUsedNodeIds() bool`

HasUsedNodeIds returns a boolean if a field has been set.

### GetOldThreshold

`func (o *TSSKeyReshareRequest) GetOldThreshold() int32`

GetOldThreshold returns the OldThreshold field if non-nil, zero value otherwise.

### GetOldThresholdOk

`func (o *TSSKeyReshareRequest) GetOldThresholdOk() (*int32, bool)`

GetOldThresholdOk returns a tuple with the OldThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldThreshold

`func (o *TSSKeyReshareRequest) SetOldThreshold(v int32)`

SetOldThreshold sets OldThreshold field to given value.

### HasOldThreshold

`func (o *TSSKeyReshareRequest) HasOldThreshold() bool`

HasOldThreshold returns a boolean if a field has been set.

### GetNewThreshold

`func (o *TSSKeyReshareRequest) GetNewThreshold() int32`

GetNewThreshold returns the NewThreshold field if non-nil, zero value otherwise.

### GetNewThresholdOk

`func (o *TSSKeyReshareRequest) GetNewThresholdOk() (*int32, bool)`

GetNewThresholdOk returns a tuple with the NewThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewThreshold

`func (o *TSSKeyReshareRequest) SetNewThreshold(v int32)`

SetNewThreshold sets NewThreshold field to given value.

### HasNewThreshold

`func (o *TSSKeyReshareRequest) HasNewThreshold() bool`

HasNewThreshold returns a boolean if a field has been set.

### GetNewNodeIds

`func (o *TSSKeyReshareRequest) GetNewNodeIds() []string`

GetNewNodeIds returns the NewNodeIds field if non-nil, zero value otherwise.

### GetNewNodeIdsOk

`func (o *TSSKeyReshareRequest) GetNewNodeIdsOk() (*[]string, bool)`

GetNewNodeIdsOk returns a tuple with the NewNodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewNodeIds

`func (o *TSSKeyReshareRequest) SetNewNodeIds(v []string)`

SetNewNodeIds sets NewNodeIds field to given value.

### HasNewNodeIds

`func (o *TSSKeyReshareRequest) HasNewNodeIds() bool`

HasNewNodeIds returns a boolean if a field has been set.

### GetTaskId

`func (o *TSSKeyReshareRequest) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TSSKeyReshareRequest) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TSSKeyReshareRequest) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *TSSKeyReshareRequest) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetBizTaskId

`func (o *TSSKeyReshareRequest) GetBizTaskId() string`

GetBizTaskId returns the BizTaskId field if non-nil, zero value otherwise.

### GetBizTaskIdOk

`func (o *TSSKeyReshareRequest) GetBizTaskIdOk() (*string, bool)`

GetBizTaskIdOk returns a tuple with the BizTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizTaskId

`func (o *TSSKeyReshareRequest) SetBizTaskId(v string)`

SetBizTaskId sets BizTaskId field to given value.

### HasBizTaskId

`func (o *TSSKeyReshareRequest) HasBizTaskId() bool`

HasBizTaskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


