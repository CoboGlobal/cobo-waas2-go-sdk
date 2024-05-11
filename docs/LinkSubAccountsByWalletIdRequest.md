# LinkSubAccountsByWalletIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | Pointer to **string** | The ID of main account wallet. | [optional] 
**SubAccountIds** | Pointer to **[]string** | The ID list of sub accounts. | [optional] 

## Methods

### NewLinkSubAccountsByWalletIdRequest

`func NewLinkSubAccountsByWalletIdRequest() *LinkSubAccountsByWalletIdRequest`

NewLinkSubAccountsByWalletIdRequest instantiates a new LinkSubAccountsByWalletIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkSubAccountsByWalletIdRequestWithDefaults

`func NewLinkSubAccountsByWalletIdRequestWithDefaults() *LinkSubAccountsByWalletIdRequest`

NewLinkSubAccountsByWalletIdRequestWithDefaults instantiates a new LinkSubAccountsByWalletIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *LinkSubAccountsByWalletIdRequest) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *LinkSubAccountsByWalletIdRequest) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *LinkSubAccountsByWalletIdRequest) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.

### HasWalletId

`func (o *LinkSubAccountsByWalletIdRequest) HasWalletId() bool`

HasWalletId returns a boolean if a field has been set.

### GetSubAccountIds

`func (o *LinkSubAccountsByWalletIdRequest) GetSubAccountIds() []string`

GetSubAccountIds returns the SubAccountIds field if non-nil, zero value otherwise.

### GetSubAccountIdsOk

`func (o *LinkSubAccountsByWalletIdRequest) GetSubAccountIdsOk() (*[]string, bool)`

GetSubAccountIdsOk returns a tuple with the SubAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccountIds

`func (o *LinkSubAccountsByWalletIdRequest) SetSubAccountIds(v []string)`

SetSubAccountIds sets SubAccountIds field to given value.

### HasSubAccountIds

`func (o *LinkSubAccountsByWalletIdRequest) HasSubAccountIds() bool`

HasSubAccountIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


