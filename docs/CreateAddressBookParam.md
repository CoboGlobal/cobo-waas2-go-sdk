# CreateAddressBookParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainIds** | **[]string** | A list of chain IDs. | 
**Address** | **string** | The wallet address. | 
**Memo** | Pointer to **string** | The memo. | [optional] 
**Label** | Pointer to **string** | The address label. | [optional] 
**Email** | Pointer to **string** | The email of the address owner. | [optional] 

## Methods

### NewCreateAddressBookParam

`func NewCreateAddressBookParam(chainIds []string, address string, ) *CreateAddressBookParam`

NewCreateAddressBookParam instantiates a new CreateAddressBookParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAddressBookParamWithDefaults

`func NewCreateAddressBookParamWithDefaults() *CreateAddressBookParam`

NewCreateAddressBookParamWithDefaults instantiates a new CreateAddressBookParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainIds

`func (o *CreateAddressBookParam) GetChainIds() []string`

GetChainIds returns the ChainIds field if non-nil, zero value otherwise.

### GetChainIdsOk

`func (o *CreateAddressBookParam) GetChainIdsOk() (*[]string, bool)`

GetChainIdsOk returns a tuple with the ChainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainIds

`func (o *CreateAddressBookParam) SetChainIds(v []string)`

SetChainIds sets ChainIds field to given value.


### GetAddress

`func (o *CreateAddressBookParam) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CreateAddressBookParam) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CreateAddressBookParam) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMemo

`func (o *CreateAddressBookParam) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *CreateAddressBookParam) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *CreateAddressBookParam) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *CreateAddressBookParam) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetLabel

`func (o *CreateAddressBookParam) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateAddressBookParam) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateAddressBookParam) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateAddressBookParam) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetEmail

`func (o *CreateAddressBookParam) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateAddressBookParam) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateAddressBookParam) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateAddressBookParam) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


