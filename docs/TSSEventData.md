# TSSEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | [**TSSEventDataType**](TSSEventDataType.md) |  | 
**RequestId** | Pointer to **string** | The request ID. | [optional] 
**RequestType** | Pointer to [**TSSRequestTypeEenum**](TSSRequestTypeEenum.md) |  | [optional] 
**RequestStatus** | Pointer to [**TSSStatus**](TSSStatus.md) |  | [optional] 
**ExtraInfo** | Pointer to **string** | The extra info. | [optional] 
**FailedReason** | Pointer to **string** | The failed reason. | [optional] 
**RequestDetail** | Pointer to [**TSSKeyShareSignRequest**](TSSKeyShareSignRequest.md) |  | [optional] 
**Result** | Pointer to [**TSSKeyShareSignSignatures**](TSSKeyShareSignSignatures.md) |  | [optional] 

## Methods

### NewTSSEventData

`func NewTSSEventData(dataType TSSEventDataType, ) *TSSEventData`

NewTSSEventData instantiates a new TSSEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSEventDataWithDefaults

`func NewTSSEventDataWithDefaults() *TSSEventData`

NewTSSEventDataWithDefaults instantiates a new TSSEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *TSSEventData) GetDataType() TSSEventDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *TSSEventData) GetDataTypeOk() (*TSSEventDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *TSSEventData) SetDataType(v TSSEventDataType)`

SetDataType sets DataType field to given value.


### GetRequestId

`func (o *TSSEventData) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TSSEventData) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TSSEventData) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *TSSEventData) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestType

`func (o *TSSEventData) GetRequestType() TSSRequestTypeEenum`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *TSSEventData) GetRequestTypeOk() (*TSSRequestTypeEenum, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *TSSEventData) SetRequestType(v TSSRequestTypeEenum)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *TSSEventData) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetRequestStatus

`func (o *TSSEventData) GetRequestStatus() TSSStatus`

GetRequestStatus returns the RequestStatus field if non-nil, zero value otherwise.

### GetRequestStatusOk

`func (o *TSSEventData) GetRequestStatusOk() (*TSSStatus, bool)`

GetRequestStatusOk returns a tuple with the RequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStatus

`func (o *TSSEventData) SetRequestStatus(v TSSStatus)`

SetRequestStatus sets RequestStatus field to given value.

### HasRequestStatus

`func (o *TSSEventData) HasRequestStatus() bool`

HasRequestStatus returns a boolean if a field has been set.

### GetExtraInfo

`func (o *TSSEventData) GetExtraInfo() string`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *TSSEventData) GetExtraInfoOk() (*string, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *TSSEventData) SetExtraInfo(v string)`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *TSSEventData) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### GetFailedReason

`func (o *TSSEventData) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *TSSEventData) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *TSSEventData) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *TSSEventData) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.

### GetRequestDetail

`func (o *TSSEventData) GetRequestDetail() TSSKeyShareSignRequest`

GetRequestDetail returns the RequestDetail field if non-nil, zero value otherwise.

### GetRequestDetailOk

`func (o *TSSEventData) GetRequestDetailOk() (*TSSKeyShareSignRequest, bool)`

GetRequestDetailOk returns a tuple with the RequestDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDetail

`func (o *TSSEventData) SetRequestDetail(v TSSKeyShareSignRequest)`

SetRequestDetail sets RequestDetail field to given value.

### HasRequestDetail

`func (o *TSSEventData) HasRequestDetail() bool`

HasRequestDetail returns a boolean if a field has been set.

### GetResult

`func (o *TSSEventData) GetResult() TSSKeyShareSignSignatures`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TSSEventData) GetResultOk() (*TSSKeyShareSignSignatures, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TSSEventData) SetResult(v TSSKeyShareSignSignatures)`

SetResult sets Result field to given value.

### HasResult

`func (o *TSSEventData) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


