# TSSKeyGenEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | [**TSSEventDataType**](TSSEventDataType.md) |  | 
**RequestId** | Pointer to **string** | The request ID. | [optional] 
**RequestType** | Pointer to [**TSSRequestTypeEenum**](TSSRequestTypeEenum.md) |  | [optional] 
**RequestStatus** | Pointer to [**TSSStatus**](TSSStatus.md) |  | [optional] 
**ExtraInfo** | Pointer to **string** | The extra info. | [optional] 
**FailedReason** | Pointer to **string** | The failed reason. | [optional] 
**RequestDetail** | Pointer to [**TSSKeyGenRequest**](TSSKeyGenRequest.md) |  | [optional] 
**Result** | Pointer to [**TSSGroup**](TSSGroup.md) |  | [optional] 

## Methods

### NewTSSKeyGenEventData

`func NewTSSKeyGenEventData(dataType TSSEventDataType, ) *TSSKeyGenEventData`

NewTSSKeyGenEventData instantiates a new TSSKeyGenEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSKeyGenEventDataWithDefaults

`func NewTSSKeyGenEventDataWithDefaults() *TSSKeyGenEventData`

NewTSSKeyGenEventDataWithDefaults instantiates a new TSSKeyGenEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *TSSKeyGenEventData) GetDataType() TSSEventDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *TSSKeyGenEventData) GetDataTypeOk() (*TSSEventDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *TSSKeyGenEventData) SetDataType(v TSSEventDataType)`

SetDataType sets DataType field to given value.


### GetRequestId

`func (o *TSSKeyGenEventData) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TSSKeyGenEventData) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TSSKeyGenEventData) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *TSSKeyGenEventData) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestType

`func (o *TSSKeyGenEventData) GetRequestType() TSSRequestTypeEenum`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *TSSKeyGenEventData) GetRequestTypeOk() (*TSSRequestTypeEenum, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *TSSKeyGenEventData) SetRequestType(v TSSRequestTypeEenum)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *TSSKeyGenEventData) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetRequestStatus

`func (o *TSSKeyGenEventData) GetRequestStatus() TSSStatus`

GetRequestStatus returns the RequestStatus field if non-nil, zero value otherwise.

### GetRequestStatusOk

`func (o *TSSKeyGenEventData) GetRequestStatusOk() (*TSSStatus, bool)`

GetRequestStatusOk returns a tuple with the RequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStatus

`func (o *TSSKeyGenEventData) SetRequestStatus(v TSSStatus)`

SetRequestStatus sets RequestStatus field to given value.

### HasRequestStatus

`func (o *TSSKeyGenEventData) HasRequestStatus() bool`

HasRequestStatus returns a boolean if a field has been set.

### GetExtraInfo

`func (o *TSSKeyGenEventData) GetExtraInfo() string`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *TSSKeyGenEventData) GetExtraInfoOk() (*string, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *TSSKeyGenEventData) SetExtraInfo(v string)`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *TSSKeyGenEventData) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### GetFailedReason

`func (o *TSSKeyGenEventData) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *TSSKeyGenEventData) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *TSSKeyGenEventData) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *TSSKeyGenEventData) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.

### GetRequestDetail

`func (o *TSSKeyGenEventData) GetRequestDetail() TSSKeyGenRequest`

GetRequestDetail returns the RequestDetail field if non-nil, zero value otherwise.

### GetRequestDetailOk

`func (o *TSSKeyGenEventData) GetRequestDetailOk() (*TSSKeyGenRequest, bool)`

GetRequestDetailOk returns a tuple with the RequestDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDetail

`func (o *TSSKeyGenEventData) SetRequestDetail(v TSSKeyGenRequest)`

SetRequestDetail sets RequestDetail field to given value.

### HasRequestDetail

`func (o *TSSKeyGenEventData) HasRequestDetail() bool`

HasRequestDetail returns a boolean if a field has been set.

### GetResult

`func (o *TSSKeyGenEventData) GetResult() TSSGroup`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TSSKeyGenEventData) GetResultOk() (*TSSGroup, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TSSKeyGenEventData) SetResult(v TSSGroup)`

SetResult sets Result field to given value.

### HasResult

`func (o *TSSKeyGenEventData) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


