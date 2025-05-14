# TSSBaseRequestEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataType** | [**TSSEventDataType**](TSSEventDataType.md) |  | 
**RequestId** | Pointer to **string** | The request ID. | [optional] 
**RequestType** | Pointer to [**TSSRequestTypeEenum**](TSSRequestTypeEenum.md) |  | [optional] 
**RequestStatus** | Pointer to [**TSSStatus**](TSSStatus.md) |  | [optional] 
**ExtraInfo** | Pointer to **string** | The extra info. | [optional] 
**FailedReason** | Pointer to **string** | The failed reason. | [optional] 

## Methods

### NewTSSBaseRequestEventData

`func NewTSSBaseRequestEventData(dataType TSSEventDataType, ) *TSSBaseRequestEventData`

NewTSSBaseRequestEventData instantiates a new TSSBaseRequestEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSSBaseRequestEventDataWithDefaults

`func NewTSSBaseRequestEventDataWithDefaults() *TSSBaseRequestEventData`

NewTSSBaseRequestEventDataWithDefaults instantiates a new TSSBaseRequestEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataType

`func (o *TSSBaseRequestEventData) GetDataType() TSSEventDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *TSSBaseRequestEventData) GetDataTypeOk() (*TSSEventDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *TSSBaseRequestEventData) SetDataType(v TSSEventDataType)`

SetDataType sets DataType field to given value.


### GetRequestId

`func (o *TSSBaseRequestEventData) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TSSBaseRequestEventData) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TSSBaseRequestEventData) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *TSSBaseRequestEventData) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestType

`func (o *TSSBaseRequestEventData) GetRequestType() TSSRequestTypeEenum`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *TSSBaseRequestEventData) GetRequestTypeOk() (*TSSRequestTypeEenum, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *TSSBaseRequestEventData) SetRequestType(v TSSRequestTypeEenum)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *TSSBaseRequestEventData) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetRequestStatus

`func (o *TSSBaseRequestEventData) GetRequestStatus() TSSStatus`

GetRequestStatus returns the RequestStatus field if non-nil, zero value otherwise.

### GetRequestStatusOk

`func (o *TSSBaseRequestEventData) GetRequestStatusOk() (*TSSStatus, bool)`

GetRequestStatusOk returns a tuple with the RequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStatus

`func (o *TSSBaseRequestEventData) SetRequestStatus(v TSSStatus)`

SetRequestStatus sets RequestStatus field to given value.

### HasRequestStatus

`func (o *TSSBaseRequestEventData) HasRequestStatus() bool`

HasRequestStatus returns a boolean if a field has been set.

### GetExtraInfo

`func (o *TSSBaseRequestEventData) GetExtraInfo() string`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *TSSBaseRequestEventData) GetExtraInfoOk() (*string, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *TSSBaseRequestEventData) SetExtraInfo(v string)`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *TSSBaseRequestEventData) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### GetFailedReason

`func (o *TSSBaseRequestEventData) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *TSSBaseRequestEventData) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *TSSBaseRequestEventData) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *TSSBaseRequestEventData) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


