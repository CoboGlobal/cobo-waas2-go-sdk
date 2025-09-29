# TokenizationBlocklistAddressNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The allowed address. | 
**Note** | Pointer to **string** | The note for the allowed address. | [optional] 
**CreatedTimestamp** | Pointer to **int64** | The time when the blocklist address was created, in Unix timestamp format, measured in milliseconds. | [optional] 

## Methods

### NewTokenizationBlocklistAddressNote

`func NewTokenizationBlocklistAddressNote(address string, ) *TokenizationBlocklistAddressNote`

NewTokenizationBlocklistAddressNote instantiates a new TokenizationBlocklistAddressNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationBlocklistAddressNoteWithDefaults

`func NewTokenizationBlocklistAddressNoteWithDefaults() *TokenizationBlocklistAddressNote`

NewTokenizationBlocklistAddressNoteWithDefaults instantiates a new TokenizationBlocklistAddressNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TokenizationBlocklistAddressNote) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TokenizationBlocklistAddressNote) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TokenizationBlocklistAddressNote) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetNote

`func (o *TokenizationBlocklistAddressNote) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *TokenizationBlocklistAddressNote) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *TokenizationBlocklistAddressNote) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *TokenizationBlocklistAddressNote) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *TokenizationBlocklistAddressNote) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *TokenizationBlocklistAddressNote) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *TokenizationBlocklistAddressNote) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *TokenizationBlocklistAddressNote) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


