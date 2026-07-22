# MerchantKycInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KycSubmissionId** | **string** | The KYC submission ID. | 
**MerchantId** | **string** | The merchant ID. | 
**Status** | [**MerchantKycStatus**](MerchantKycStatus.md) |  | 

## Methods

### NewMerchantKycInfo

`func NewMerchantKycInfo(kycSubmissionId string, merchantId string, status MerchantKycStatus, ) *MerchantKycInfo`

NewMerchantKycInfo instantiates a new MerchantKycInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantKycInfoWithDefaults

`func NewMerchantKycInfoWithDefaults() *MerchantKycInfo`

NewMerchantKycInfoWithDefaults instantiates a new MerchantKycInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKycSubmissionId

`func (o *MerchantKycInfo) GetKycSubmissionId() string`

GetKycSubmissionId returns the KycSubmissionId field if non-nil, zero value otherwise.

### GetKycSubmissionIdOk

`func (o *MerchantKycInfo) GetKycSubmissionIdOk() (*string, bool)`

GetKycSubmissionIdOk returns a tuple with the KycSubmissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKycSubmissionId

`func (o *MerchantKycInfo) SetKycSubmissionId(v string)`

SetKycSubmissionId sets KycSubmissionId field to given value.


### GetMerchantId

`func (o *MerchantKycInfo) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *MerchantKycInfo) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *MerchantKycInfo) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetStatus

`func (o *MerchantKycInfo) GetStatus() MerchantKycStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MerchantKycInfo) GetStatusOk() (*MerchantKycStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MerchantKycInfo) SetStatus(v MerchantKycStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


