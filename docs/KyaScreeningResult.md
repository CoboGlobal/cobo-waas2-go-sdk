# KyaScreeningResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The idempotency identifier from the request, unique within your organization, used for tracking and troubleshooting. Only present in create response. | 
**ScreeningId** | **string** | The unique system-generated identifier for this screening request (UUID format, fixed 36 characters). | 
**Address** | **string** | The screened blockchain address. | 
**ChainId** | **string** | The chain identifier. | 
**Note** | Pointer to **string** | Optional note for this address screening. | [optional] 
**CreatedTimestamp** | **int64** | The time when the screening request was created, in Unix timestamp format, measured in milliseconds. | 
**RequestedBy** | **string** | The identifier of the user or application that created this screening request. | 
**Status** | [**KyaScreeningStatus**](KyaScreeningStatus.md) |  | 
**RiskAssessment** | Pointer to [**NullableKyaScreeningResultRiskAssessment**](KyaScreeningResultRiskAssessment.md) |  | [optional] 

## Methods

### NewKyaScreeningResult

`func NewKyaScreeningResult(requestId string, screeningId string, address string, chainId string, createdTimestamp int64, requestedBy string, status KyaScreeningStatus, ) *KyaScreeningResult`

NewKyaScreeningResult instantiates a new KyaScreeningResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKyaScreeningResultWithDefaults

`func NewKyaScreeningResultWithDefaults() *KyaScreeningResult`

NewKyaScreeningResultWithDefaults instantiates a new KyaScreeningResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *KyaScreeningResult) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *KyaScreeningResult) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *KyaScreeningResult) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetScreeningId

`func (o *KyaScreeningResult) GetScreeningId() string`

GetScreeningId returns the ScreeningId field if non-nil, zero value otherwise.

### GetScreeningIdOk

`func (o *KyaScreeningResult) GetScreeningIdOk() (*string, bool)`

GetScreeningIdOk returns a tuple with the ScreeningId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreeningId

`func (o *KyaScreeningResult) SetScreeningId(v string)`

SetScreeningId sets ScreeningId field to given value.


### GetAddress

`func (o *KyaScreeningResult) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *KyaScreeningResult) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *KyaScreeningResult) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChainId

`func (o *KyaScreeningResult) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *KyaScreeningResult) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *KyaScreeningResult) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetNote

`func (o *KyaScreeningResult) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *KyaScreeningResult) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *KyaScreeningResult) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *KyaScreeningResult) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *KyaScreeningResult) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *KyaScreeningResult) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *KyaScreeningResult) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.


### GetRequestedBy

`func (o *KyaScreeningResult) GetRequestedBy() string`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *KyaScreeningResult) GetRequestedByOk() (*string, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *KyaScreeningResult) SetRequestedBy(v string)`

SetRequestedBy sets RequestedBy field to given value.


### GetStatus

`func (o *KyaScreeningResult) GetStatus() KyaScreeningStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KyaScreeningResult) GetStatusOk() (*KyaScreeningStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KyaScreeningResult) SetStatus(v KyaScreeningStatus)`

SetStatus sets Status field to given value.


### GetRiskAssessment

`func (o *KyaScreeningResult) GetRiskAssessment() KyaScreeningResultRiskAssessment`

GetRiskAssessment returns the RiskAssessment field if non-nil, zero value otherwise.

### GetRiskAssessmentOk

`func (o *KyaScreeningResult) GetRiskAssessmentOk() (*KyaScreeningResultRiskAssessment, bool)`

GetRiskAssessmentOk returns a tuple with the RiskAssessment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAssessment

`func (o *KyaScreeningResult) SetRiskAssessment(v KyaScreeningResultRiskAssessment)`

SetRiskAssessment sets RiskAssessment field to given value.

### HasRiskAssessment

`func (o *KyaScreeningResult) HasRiskAssessment() bool`

HasRiskAssessment returns a boolean if a field has been set.

### SetRiskAssessmentNil

`func (o *KyaScreeningResult) SetRiskAssessmentNil(b bool)`

 SetRiskAssessmentNil sets the value for RiskAssessment to be an explicit nil

### UnsetRiskAssessment
`func (o *KyaScreeningResult) UnsetRiskAssessment()`

UnsetRiskAssessment ensures that no value is present for RiskAssessment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


