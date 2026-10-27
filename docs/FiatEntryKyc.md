# FiatEntryKyc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FiatEntryKycId** | **string** | The fiat entry KYC ID. | 
**Status** | [**FiatEntryKycStatus**](FiatEntryKycStatus.md) |  | 

## Methods

### NewFiatEntryKyc

`func NewFiatEntryKyc(fiatEntryKycId string, status FiatEntryKycStatus, ) *FiatEntryKyc`

NewFiatEntryKyc instantiates a new FiatEntryKyc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiatEntryKycWithDefaults

`func NewFiatEntryKycWithDefaults() *FiatEntryKyc`

NewFiatEntryKycWithDefaults instantiates a new FiatEntryKyc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiatEntryKycId

`func (o *FiatEntryKyc) GetFiatEntryKycId() string`

GetFiatEntryKycId returns the FiatEntryKycId field if non-nil, zero value otherwise.

### GetFiatEntryKycIdOk

`func (o *FiatEntryKyc) GetFiatEntryKycIdOk() (*string, bool)`

GetFiatEntryKycIdOk returns a tuple with the FiatEntryKycId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatEntryKycId

`func (o *FiatEntryKyc) SetFiatEntryKycId(v string)`

SetFiatEntryKycId sets FiatEntryKycId field to given value.


### GetStatus

`func (o *FiatEntryKyc) GetStatus() FiatEntryKycStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FiatEntryKyc) GetStatusOk() (*FiatEntryKycStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FiatEntryKyc) SetStatus(v FiatEntryKycStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


