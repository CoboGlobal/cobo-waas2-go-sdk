# KyaScreeningsEventData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScreeningId** | **string** | The unique system-generated identifier for this screening request (UUID format, fixed 36 characters). | 
**Address** | **string** | The screened blockchain address. | 
**ChainId** | **string** | The chain identifier. | 
**Status** | [**KyaScreeningStatus**](KyaScreeningStatus.md) |  | 
**CreatedTimestamp** | **int64** | The time when the screening request was created, in Unix timestamp format, measured in milliseconds. | 
**UpdatedTimestamp** | **int64** | The time when the screening status was updated, in Unix timestamp format, measured in milliseconds. | 

## Methods

### NewKyaScreeningsEventData

`func NewKyaScreeningsEventData(screeningId string, address string, chainId string, status KyaScreeningStatus, createdTimestamp int64, updatedTimestamp int64, ) *KyaScreeningsEventData`

NewKyaScreeningsEventData instantiates a new KyaScreeningsEventData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKyaScreeningsEventDataWithDefaults

`func NewKyaScreeningsEventDataWithDefaults() *KyaScreeningsEventData`

NewKyaScreeningsEventDataWithDefaults instantiates a new KyaScreeningsEventData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScreeningId

`func (o *KyaScreeningsEventData) GetScreeningId() string`

GetScreeningId returns the ScreeningId field if non-nil, zero value otherwise.

### GetScreeningIdOk

`func (o *KyaScreeningsEventData) GetScreeningIdOk() (*string, bool)`

GetScreeningIdOk returns a tuple with the ScreeningId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreeningId

`func (o *KyaScreeningsEventData) SetScreeningId(v string)`

SetScreeningId sets ScreeningId field to given value.


### GetAddress

`func (o *KyaScreeningsEventData) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *KyaScreeningsEventData) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *KyaScreeningsEventData) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChainId

`func (o *KyaScreeningsEventData) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *KyaScreeningsEventData) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *KyaScreeningsEventData) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetStatus

`func (o *KyaScreeningsEventData) GetStatus() KyaScreeningStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KyaScreeningsEventData) GetStatusOk() (*KyaScreeningStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KyaScreeningsEventData) SetStatus(v KyaScreeningStatus)`

SetStatus sets Status field to given value.


### GetCreatedTimestamp

`func (o *KyaScreeningsEventData) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *KyaScreeningsEventData) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *KyaScreeningsEventData) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetUpdatedTimestamp

`func (o *KyaScreeningsEventData) GetUpdatedTimestamp() int64`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *KyaScreeningsEventData) GetUpdatedTimestampOk() (*int64, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *KyaScreeningsEventData) SetUpdatedTimestamp(v int64)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


