# PreCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipChecks** | Pointer to [**[]SkipCheckType**](SkipCheckType.md) | Selection to skip verification. | [optional] 

## Methods

### NewPreCheck

`func NewPreCheck() *PreCheck`

NewPreCheck instantiates a new PreCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreCheckWithDefaults

`func NewPreCheckWithDefaults() *PreCheck`

NewPreCheckWithDefaults instantiates a new PreCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipChecks

`func (o *PreCheck) GetSkipChecks() []SkipCheckType`

GetSkipChecks returns the SkipChecks field if non-nil, zero value otherwise.

### GetSkipChecksOk

`func (o *PreCheck) GetSkipChecksOk() (*[]SkipCheckType, bool)`

GetSkipChecksOk returns a tuple with the SkipChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipChecks

`func (o *PreCheck) SetSkipChecks(v []SkipCheckType)`

SetSkipChecks sets SkipChecks field to given value.

### HasSkipChecks

`func (o *PreCheck) HasSkipChecks() bool`

HasSkipChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


