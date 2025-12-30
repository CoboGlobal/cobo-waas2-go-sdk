# KyaRiskDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | The risk category (e.g., sanctions, fraud, theft, etc.). | 
**Exposure** | **string** | The exposure level for this risk category (e.g., high, medium, low, none, indirect). | 

## Methods

### NewKyaRiskDetail

`func NewKyaRiskDetail(category string, exposure string, ) *KyaRiskDetail`

NewKyaRiskDetail instantiates a new KyaRiskDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKyaRiskDetailWithDefaults

`func NewKyaRiskDetailWithDefaults() *KyaRiskDetail`

NewKyaRiskDetailWithDefaults instantiates a new KyaRiskDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *KyaRiskDetail) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *KyaRiskDetail) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *KyaRiskDetail) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetExposure

`func (o *KyaRiskDetail) GetExposure() string`

GetExposure returns the Exposure field if non-nil, zero value otherwise.

### GetExposureOk

`func (o *KyaRiskDetail) GetExposureOk() (*string, bool)`

GetExposureOk returns a tuple with the Exposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposure

`func (o *KyaRiskDetail) SetExposure(v string)`

SetExposure sets Exposure field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


