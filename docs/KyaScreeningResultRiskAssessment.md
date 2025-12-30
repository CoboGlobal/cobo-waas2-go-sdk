# KyaScreeningResultRiskAssessment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | [**KyaRiskLevel**](KyaRiskLevel.md) |  | 
**Summary** | Pointer to **string** | A brief summary of the risk assessment. | [optional] 
**Details** | Pointer to [**[]KyaRiskDetail**](KyaRiskDetail.md) | Detailed risk category assessments. | [optional] 

## Methods

### NewKyaScreeningResultRiskAssessment

`func NewKyaScreeningResultRiskAssessment(level KyaRiskLevel, ) *KyaScreeningResultRiskAssessment`

NewKyaScreeningResultRiskAssessment instantiates a new KyaScreeningResultRiskAssessment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKyaScreeningResultRiskAssessmentWithDefaults

`func NewKyaScreeningResultRiskAssessmentWithDefaults() *KyaScreeningResultRiskAssessment`

NewKyaScreeningResultRiskAssessmentWithDefaults instantiates a new KyaScreeningResultRiskAssessment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *KyaScreeningResultRiskAssessment) GetLevel() KyaRiskLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *KyaScreeningResultRiskAssessment) GetLevelOk() (*KyaRiskLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *KyaScreeningResultRiskAssessment) SetLevel(v KyaRiskLevel)`

SetLevel sets Level field to given value.


### GetSummary

`func (o *KyaScreeningResultRiskAssessment) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *KyaScreeningResultRiskAssessment) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *KyaScreeningResultRiskAssessment) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *KyaScreeningResultRiskAssessment) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDetails

`func (o *KyaScreeningResultRiskAssessment) GetDetails() []KyaRiskDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *KyaScreeningResultRiskAssessment) GetDetailsOk() (*[]KyaRiskDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *KyaScreeningResultRiskAssessment) SetDetails(v []KyaRiskDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *KyaScreeningResultRiskAssessment) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


