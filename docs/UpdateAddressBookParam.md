# UpdateAddressBookParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainIds** | **[]string** | A list of chain IDs. | 
**Memo** | Pointer to **string** | The memo. | [optional] 
**Label** | Pointer to **string** | The address label. | [optional] 
**Email** | Pointer to **string** | The email of the address owner. | [optional] 

## Methods

### NewUpdateAddressBookParam

`func NewUpdateAddressBookParam(chainIds []string, ) *UpdateAddressBookParam`

NewUpdateAddressBookParam instantiates a new UpdateAddressBookParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAddressBookParamWithDefaults

`func NewUpdateAddressBookParamWithDefaults() *UpdateAddressBookParam`

NewUpdateAddressBookParamWithDefaults instantiates a new UpdateAddressBookParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainIds

`func (o *UpdateAddressBookParam) GetChainIds() []string`

GetChainIds returns the ChainIds field if non-nil, zero value otherwise.

### GetChainIdsOk

`func (o *UpdateAddressBookParam) GetChainIdsOk() (*[]string, bool)`

GetChainIdsOk returns a tuple with the ChainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainIds

`func (o *UpdateAddressBookParam) SetChainIds(v []string)`

SetChainIds sets ChainIds field to given value.


### GetMemo

`func (o *UpdateAddressBookParam) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *UpdateAddressBookParam) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *UpdateAddressBookParam) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *UpdateAddressBookParam) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetLabel

`func (o *UpdateAddressBookParam) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *UpdateAddressBookParam) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *UpdateAddressBookParam) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *UpdateAddressBookParam) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateAddressBookParam) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateAddressBookParam) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateAddressBookParam) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateAddressBookParam) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


