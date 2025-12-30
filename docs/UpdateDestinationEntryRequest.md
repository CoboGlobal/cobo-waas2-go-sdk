# UpdateDestinationEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryType** | Pointer to [**EntryType**](EntryType.md) |  | [optional] 
**DestinationId** | **string** | The destination ID. | 
**BankAccount** | Pointer to [**UpdateDestinationBankAccount**](UpdateDestinationBankAccount.md) |  | [optional] 

## Methods

### NewUpdateDestinationEntryRequest

`func NewUpdateDestinationEntryRequest(destinationId string, ) *UpdateDestinationEntryRequest`

NewUpdateDestinationEntryRequest instantiates a new UpdateDestinationEntryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDestinationEntryRequestWithDefaults

`func NewUpdateDestinationEntryRequestWithDefaults() *UpdateDestinationEntryRequest`

NewUpdateDestinationEntryRequestWithDefaults instantiates a new UpdateDestinationEntryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryType

`func (o *UpdateDestinationEntryRequest) GetEntryType() EntryType`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *UpdateDestinationEntryRequest) GetEntryTypeOk() (*EntryType, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *UpdateDestinationEntryRequest) SetEntryType(v EntryType)`

SetEntryType sets EntryType field to given value.

### HasEntryType

`func (o *UpdateDestinationEntryRequest) HasEntryType() bool`

HasEntryType returns a boolean if a field has been set.

### GetDestinationId

`func (o *UpdateDestinationEntryRequest) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *UpdateDestinationEntryRequest) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *UpdateDestinationEntryRequest) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetBankAccount

`func (o *UpdateDestinationEntryRequest) GetBankAccount() UpdateDestinationBankAccount`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *UpdateDestinationEntryRequest) GetBankAccountOk() (*UpdateDestinationBankAccount, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *UpdateDestinationEntryRequest) SetBankAccount(v UpdateDestinationBankAccount)`

SetBankAccount sets BankAccount field to given value.

### HasBankAccount

`func (o *UpdateDestinationEntryRequest) HasBankAccount() bool`

HasBankAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


