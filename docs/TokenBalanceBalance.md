# TokenBalanceBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **string** | The total balance, which is the sum of the available, pending, and locked balances. | 
**Available** | **string** | The balance free to use. | 
**Pending** | Pointer to **string** | The balance that is currently pending confirmation. | [optional] [default to "0"]
**Locked** | Pointer to **string** | The balance that is currently not accessible to transactions. | [optional] [default to "0"]

## Methods

### NewTokenBalanceBalance

`func NewTokenBalanceBalance(total string, available string, ) *TokenBalanceBalance`

NewTokenBalanceBalance instantiates a new TokenBalanceBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenBalanceBalanceWithDefaults

`func NewTokenBalanceBalanceWithDefaults() *TokenBalanceBalance`

NewTokenBalanceBalanceWithDefaults instantiates a new TokenBalanceBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *TokenBalanceBalance) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TokenBalanceBalance) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TokenBalanceBalance) SetTotal(v string)`

SetTotal sets Total field to given value.


### GetAvailable

`func (o *TokenBalanceBalance) GetAvailable() string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *TokenBalanceBalance) GetAvailableOk() (*string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *TokenBalanceBalance) SetAvailable(v string)`

SetAvailable sets Available field to given value.


### GetPending

`func (o *TokenBalanceBalance) GetPending() string`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *TokenBalanceBalance) GetPendingOk() (*string, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *TokenBalanceBalance) SetPending(v string)`

SetPending sets Pending field to given value.

### HasPending

`func (o *TokenBalanceBalance) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetLocked

`func (o *TokenBalanceBalance) GetLocked() string`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *TokenBalanceBalance) GetLockedOk() (*string, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *TokenBalanceBalance) SetLocked(v string)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *TokenBalanceBalance) HasLocked() bool`

HasLocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


