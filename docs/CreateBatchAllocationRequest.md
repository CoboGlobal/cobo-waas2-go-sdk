# CreateBatchAllocationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a batch allocation request. The request ID is provided by you and must be unique. | 
**AllocationParams** | [**[]AllocationParam**](AllocationParam.md) |  | 

## Methods

### NewCreateBatchAllocationRequest

`func NewCreateBatchAllocationRequest(requestId string, allocationParams []AllocationParam, ) *CreateBatchAllocationRequest`

NewCreateBatchAllocationRequest instantiates a new CreateBatchAllocationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBatchAllocationRequestWithDefaults

`func NewCreateBatchAllocationRequestWithDefaults() *CreateBatchAllocationRequest`

NewCreateBatchAllocationRequestWithDefaults instantiates a new CreateBatchAllocationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *CreateBatchAllocationRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateBatchAllocationRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateBatchAllocationRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetAllocationParams

`func (o *CreateBatchAllocationRequest) GetAllocationParams() []AllocationParam`

GetAllocationParams returns the AllocationParams field if non-nil, zero value otherwise.

### GetAllocationParamsOk

`func (o *CreateBatchAllocationRequest) GetAllocationParamsOk() (*[]AllocationParam, bool)`

GetAllocationParamsOk returns a tuple with the AllocationParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationParams

`func (o *CreateBatchAllocationRequest) SetAllocationParams(v []AllocationParam)`

SetAllocationParams sets AllocationParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


