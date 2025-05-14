# TSSKeyShareSignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **string** | The node ID of the key share holder. | [optional] 
**TaskId** | Pointer to **string** | The task ID. | [optional] 
**Details** | Pointer to [**[]TSSKeyShareSignDetail**](TSSKeyShareSignDetail.md) |  | [optional] 
**BizTaskId** | Pointer to **string** | The business task ID. This field contains the key share sign request ID. | [optional] 

## Methods

### NewTSSKeyShareSignRequest

`func NewTSSKeyShareSignRequest() *TSSKeyShareSignRequest`

NewTSSKeyShareSignRequest instantiates a new TSSKeyShareSignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSKeyShareSignRequestWithDefaults

`func NewTSSKeyShareSignRequestWithDefaults() *TSSKeyShareSignRequest`

NewTSSKeyShareSignRequestWithDefaults instantiates a new TSSKeyShareSignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *TSSKeyShareSignRequest) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TSSKeyShareSignRequest) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TSSKeyShareSignRequest) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *TSSKeyShareSignRequest) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetTaskId

`func (o *TSSKeyShareSignRequest) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TSSKeyShareSignRequest) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TSSKeyShareSignRequest) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *TSSKeyShareSignRequest) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetDetails

`func (o *TSSKeyShareSignRequest) GetDetails() []TSSKeyShareSignDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *TSSKeyShareSignRequest) GetDetailsOk() (*[]TSSKeyShareSignDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *TSSKeyShareSignRequest) SetDetails(v []TSSKeyShareSignDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *TSSKeyShareSignRequest) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetBizTaskId

`func (o *TSSKeyShareSignRequest) GetBizTaskId() string`

GetBizTaskId returns the BizTaskId field if non-nil, zero value otherwise.

### GetBizTaskIdOk

`func (o *TSSKeyShareSignRequest) GetBizTaskIdOk() (*string, bool)`

GetBizTaskIdOk returns a tuple with the BizTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizTaskId

`func (o *TSSKeyShareSignRequest) SetBizTaskId(v string)`

SetBizTaskId sets BizTaskId field to given value.

### HasBizTaskId

`func (o *TSSKeyShareSignRequest) HasBizTaskId() bool`

HasBizTaskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


