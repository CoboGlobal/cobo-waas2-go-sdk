# TSSKeyShareSignEventData

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

### NewTSSKeyShareSignEventData

`func NewTSSKeyShareSignEventData(dataType TSSEventDataType, ) *TSSKeyShareSignEventData`

NewTSSKeyShareSignEventData instantiates a new TSSKeyShareSignEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSKeyShareSignEventDataWithDefaults

`func NewTSSKeyShareSignEventDataWithDefaults() *TSSKeyShareSignEventData`

NewTSSKeyShareSignEventDataWithDefaults instantiates a new TSSKeyShareSignEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *TSSKeyShareSignEventData) GetDataType() TSSEventDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *TSSKeyShareSignEventData) GetDataTypeOk() (*TSSEventDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *TSSKeyShareSignEventData) SetDataType(v TSSEventDataType)`

SetDataType sets DataType field to given value.


### GetRequestId

`func (o *TSSKeyShareSignEventData) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TSSKeyShareSignEventData) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TSSKeyShareSignEventData) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *TSSKeyShareSignEventData) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestType

`func (o *TSSKeyShareSignEventData) GetRequestType() TSSRequestTypeEenum`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *TSSKeyShareSignEventData) GetRequestTypeOk() (*TSSRequestTypeEenum, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *TSSKeyShareSignEventData) SetRequestType(v TSSRequestTypeEenum)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *TSSKeyShareSignEventData) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetRequestStatus

`func (o *TSSKeyShareSignEventData) GetRequestStatus() TSSStatus`

GetRequestStatus returns the RequestStatus field if non-nil, zero value otherwise.

### GetRequestStatusOk

`func (o *TSSKeyShareSignEventData) GetRequestStatusOk() (*TSSStatus, bool)`

GetRequestStatusOk returns a tuple with the RequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStatus

`func (o *TSSKeyShareSignEventData) SetRequestStatus(v TSSStatus)`

SetRequestStatus sets RequestStatus field to given value.

### HasRequestStatus

`func (o *TSSKeyShareSignEventData) HasRequestStatus() bool`

HasRequestStatus returns a boolean if a field has been set.

### GetExtraInfo

`func (o *TSSKeyShareSignEventData) GetExtraInfo() string`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *TSSKeyShareSignEventData) GetExtraInfoOk() (*string, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *TSSKeyShareSignEventData) SetExtraInfo(v string)`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *TSSKeyShareSignEventData) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### GetFailedReason

`func (o *TSSKeyShareSignEventData) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *TSSKeyShareSignEventData) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *TSSKeyShareSignEventData) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *TSSKeyShareSignEventData) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.

### GetRequestDetail

`func (o *TSSKeyShareSignEventData) GetRequestDetail() TSSKeyShareSignRequest`

GetRequestDetail returns the RequestDetail field if non-nil, zero value otherwise.

### GetRequestDetailOk

`func (o *TSSKeyShareSignEventData) GetRequestDetailOk() (*TSSKeyShareSignRequest, bool)`

GetRequestDetailOk returns a tuple with the RequestDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDetail

`func (o *TSSKeyShareSignEventData) SetRequestDetail(v TSSKeyShareSignRequest)`

SetRequestDetail sets RequestDetail field to given value.

### HasRequestDetail

`func (o *TSSKeyShareSignEventData) HasRequestDetail() bool`

HasRequestDetail returns a boolean if a field has been set.

### GetResult

`func (o *TSSKeyShareSignEventData) GetResult() TSSKeyShareSignSignatures`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TSSKeyShareSignEventData) GetResultOk() (*TSSKeyShareSignSignatures, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TSSKeyShareSignEventData) SetResult(v TSSKeyShareSignSignatures)`

SetResult sets Result field to given value.

### HasResult

`func (o *TSSKeyShareSignEventData) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


