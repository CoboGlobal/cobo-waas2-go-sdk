# SolContractCallAddressLookupTableAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AltAccountKey** | **string** | The on-chain public key of the Address Lookup Table (ALT) account, identifying the specific lookup table. | 
**Addresses** | **[]string** | An array of stored account addresses within the lookup table, which can be referenced in transactions by index. | 

## Methods

### NewSolContractCallAddressLookupTableAccount

`func NewSolContractCallAddressLookupTableAccount(altAccountKey string, addresses []string, ) *SolContractCallAddressLookupTableAccount`

NewSolContractCallAddressLookupTableAccount instantiates a new SolContractCallAddressLookupTableAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolContractCallAddressLookupTableAccountWithDefaults

`func NewSolContractCallAddressLookupTableAccountWithDefaults() *SolContractCallAddressLookupTableAccount`

NewSolContractCallAddressLookupTableAccountWithDefaults instantiates a new SolContractCallAddressLookupTableAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAltAccountKey

`func (o *SolContractCallAddressLookupTableAccount) GetAltAccountKey() string`

GetAltAccountKey returns the AltAccountKey field if non-nil, zero value otherwise.

### GetAltAccountKeyOk

`func (o *SolContractCallAddressLookupTableAccount) GetAltAccountKeyOk() (*string, bool)`

GetAltAccountKeyOk returns a tuple with the AltAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltAccountKey

`func (o *SolContractCallAddressLookupTableAccount) SetAltAccountKey(v string)`

SetAltAccountKey sets AltAccountKey field to given value.


### GetAddresses

`func (o *SolContractCallAddressLookupTableAccount) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *SolContractCallAddressLookupTableAccount) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *SolContractCallAddressLookupTableAccount) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


