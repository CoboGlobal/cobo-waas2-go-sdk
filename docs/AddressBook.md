# AddressBook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** |  | 
**EntryId** | **string** |  | 
**Address** | **string** | address. | 
**Memo** | Pointer to **string** | memo. | [optional] 
**WalletName** | Pointer to **string** | wallet name. | [optional] 
**WalletType** | Pointer to [**WalletType**](WalletType.md) |  | [optional] 
**Label** | **string** | The label to address. | 
**Email** | Pointer to **string** | email. | [optional] 

## Methods

### NewAddressBook

`func NewAddressBook(orgId string, entryId string, address string, label string, ) *AddressBook`

NewAddressBook instantiates a new AddressBook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressBookWithDefaults

`func NewAddressBookWithDefaults() *AddressBook`

NewAddressBookWithDefaults instantiates a new AddressBook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *AddressBook) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AddressBook) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AddressBook) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetEntryId

`func (o *AddressBook) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *AddressBook) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *AddressBook) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetAddress

`func (o *AddressBook) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressBook) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressBook) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMemo

`func (o *AddressBook) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *AddressBook) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *AddressBook) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *AddressBook) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetWalletName

`func (o *AddressBook) GetWalletName() string`

GetWalletName returns the WalletName field if non-nil, zero value otherwise.

### GetWalletNameOk

`func (o *AddressBook) GetWalletNameOk() (*string, bool)`

GetWalletNameOk returns a tuple with the WalletName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletName

`func (o *AddressBook) SetWalletName(v string)`

SetWalletName sets WalletName field to given value.

### HasWalletName

`func (o *AddressBook) HasWalletName() bool`

HasWalletName returns a boolean if a field has been set.

### GetWalletType

`func (o *AddressBook) GetWalletType() WalletType`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *AddressBook) GetWalletTypeOk() (*WalletType, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *AddressBook) SetWalletType(v WalletType)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *AddressBook) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.

### GetLabel

`func (o *AddressBook) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AddressBook) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AddressBook) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetEmail

`func (o *AddressBook) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddressBook) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddressBook) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddressBook) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


