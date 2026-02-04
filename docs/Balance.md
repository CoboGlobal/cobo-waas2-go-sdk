# Balance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **string** | The current amount of tokens in an address, which is retrieved directly from the network. To learn more, see [Balances and transaction amounts for MPC Wallets](https://www.cobo.com/developers/v2/guides/mpc-wallets/balance-amounts) for more details. | 
**Available** | **string** | The amount of tokens ready to be spent. To learn more, see [Balances and transaction amounts for MPC Wallets](https://www.cobo.com/developers/v2/guides/mpc-wallets/balance-amounts) for more details. | 
**Pending** | Pointer to **string** | The total amount being sent in a transaction, which is calculated as the withdrawal amount plus the transaction fee. To learn more, see [Balances and transaction amounts for MPC Wallets](https://www.cobo.com/developers/v2/guides/mpc-wallets/balance-amounts) for more details. | [optional] [default to "0"]
**Locked** | Pointer to **string** | For UTXO chains, this is the combined value of the selected UTXOs for the transaction. For other chains, it is equal to the Pending amount. To learn more, see [Balances and transaction amounts for MPC Wallets](https://www.cobo.com/developers/v2/guides/mpc-wallets/balance-amounts) for more details. | [optional] [default to "0"]
**Frozen** | Pointer to **string** | Amount frozen due to compliance inspection. To learn more, see [Balances and transaction amounts for MPC Wallets](https://www.cobo.com/developers/v2/guides/mpc-wallets/balance-amounts) for more details. | [optional] [default to "0"]

## Methods

### NewBalance

`func NewBalance(total string, available string, ) *Balance`

NewBalance instantiates a new Balance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceWithDefaults

`func NewBalanceWithDefaults() *Balance`

NewBalanceWithDefaults instantiates a new Balance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *Balance) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Balance) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Balance) SetTotal(v string)`

SetTotal sets Total field to given value.


### GetAvailable

`func (o *Balance) GetAvailable() string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *Balance) GetAvailableOk() (*string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *Balance) SetAvailable(v string)`

SetAvailable sets Available field to given value.


### GetPending

`func (o *Balance) GetPending() string`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *Balance) GetPendingOk() (*string, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *Balance) SetPending(v string)`

SetPending sets Pending field to given value.

### HasPending

`func (o *Balance) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetLocked

`func (o *Balance) GetLocked() string`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Balance) GetLockedOk() (*string, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Balance) SetLocked(v string)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Balance) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetFrozen

`func (o *Balance) GetFrozen() string`

GetFrozen returns the Frozen field if non-nil, zero value otherwise.

### GetFrozenOk

`func (o *Balance) GetFrozenOk() (*string, bool)`

GetFrozenOk returns a tuple with the Frozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozen

`func (o *Balance) SetFrozen(v string)`

SetFrozen sets Frozen field to given value.

### HasFrozen

`func (o *Balance) HasFrozen() bool`

HasFrozen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


