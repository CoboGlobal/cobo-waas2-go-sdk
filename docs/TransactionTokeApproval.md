# TransactionTokeApproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token ID, which is the unique identifier of a token. | 
**ChainId** | **string** | The ID of the chain on which the token operates. | 
**AssetId** | Pointer to **string** | (This concept applies to Exchange Wallets only) The asset ID. An asset ID is the unique identifier of the asset held within your linked exchange account. | [optional] 
**Symbol** | Pointer to **string** | The token symbol, which is the abbreviated name of a token. | [optional] 
**Name** | Pointer to **string** | The token name, which is the full name of a token. | [optional] 
**Decimal** | Pointer to **int32** | The token decimal. | [optional] 
**IconUrl** | Pointer to **string** | The URL of the token icon. | [optional] 
**TokenAddress** | Pointer to **string** | The token address, if applicable. | [optional] 
**FeeTokenId** | Pointer to **string** | The fee token ID. A fee token is the token with which you pay transaction fees. | [optional] 
**CanDeposit** | Pointer to **bool** | Whether deposits are enabled for this token. | [optional] 
**CanWithdraw** | Pointer to **bool** | Whether withdrawals are enabled for this token. | [optional] 
**DustThreshold** | Pointer to **string** | The minimum withdrawal amount for Custodial Wallets. If your withdrawal amount is smaller than this threshold, the withdrawal request will receive an error.  Note: [Cobo Loop](https://manuals.cobo.com/en/portal/custodial-wallets/cobo-loop) transfers do not have this limitation.  | [optional] 
**CustodialMinimumDepositThreshold** | Pointer to **string** | The minimum deposit amount for Custodial Wallets. If the amount you deposit to a Custodial Wallet is smaller than this threshold, the deposit will not show up on Cobo Portal or trigger any webhook events.  Note: [Cobo Loop](https://manuals.cobo.com/en/portal/custodial-wallets/cobo-loop)transfers do not have this limitation.  | [optional] 
**AssetModelType** | Pointer to [**TokenAssetModelType**](TokenAssetModelType.md) |  | [optional] 
**Amount** | Pointer to **float32** | Transaction value (Note that this is an absolute value. If you trade 1.5 BTC, then the value is 1.5)  | [optional] 
**Spender** | Pointer to **string** | Spender address | [optional] 

## Methods

### NewTransactionTokeApproval

`func NewTransactionTokeApproval(tokenId string, chainId string, ) *TransactionTokeApproval`

NewTransactionTokeApproval instantiates a new TransactionTokeApproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionTokeApprovalWithDefaults

`func NewTransactionTokeApprovalWithDefaults() *TransactionTokeApproval`

NewTransactionTokeApprovalWithDefaults instantiates a new TransactionTokeApproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *TransactionTokeApproval) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransactionTokeApproval) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransactionTokeApproval) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChainId

`func (o *TransactionTokeApproval) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TransactionTokeApproval) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TransactionTokeApproval) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetAssetId

`func (o *TransactionTokeApproval) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *TransactionTokeApproval) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *TransactionTokeApproval) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *TransactionTokeApproval) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetSymbol

`func (o *TransactionTokeApproval) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TransactionTokeApproval) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TransactionTokeApproval) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TransactionTokeApproval) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetName

`func (o *TransactionTokeApproval) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TransactionTokeApproval) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TransactionTokeApproval) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TransactionTokeApproval) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDecimal

`func (o *TransactionTokeApproval) GetDecimal() int32`

GetDecimal returns the Decimal field if non-nil, zero value otherwise.

### GetDecimalOk

`func (o *TransactionTokeApproval) GetDecimalOk() (*int32, bool)`

GetDecimalOk returns a tuple with the Decimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimal

`func (o *TransactionTokeApproval) SetDecimal(v int32)`

SetDecimal sets Decimal field to given value.

### HasDecimal

`func (o *TransactionTokeApproval) HasDecimal() bool`

HasDecimal returns a boolean if a field has been set.

### GetIconUrl

`func (o *TransactionTokeApproval) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *TransactionTokeApproval) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *TransactionTokeApproval) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *TransactionTokeApproval) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetTokenAddress

`func (o *TransactionTokeApproval) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TransactionTokeApproval) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TransactionTokeApproval) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.

### HasTokenAddress

`func (o *TransactionTokeApproval) HasTokenAddress() bool`

HasTokenAddress returns a boolean if a field has been set.

### GetFeeTokenId

`func (o *TransactionTokeApproval) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *TransactionTokeApproval) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *TransactionTokeApproval) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.

### HasFeeTokenId

`func (o *TransactionTokeApproval) HasFeeTokenId() bool`

HasFeeTokenId returns a boolean if a field has been set.

### GetCanDeposit

`func (o *TransactionTokeApproval) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *TransactionTokeApproval) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *TransactionTokeApproval) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *TransactionTokeApproval) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *TransactionTokeApproval) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *TransactionTokeApproval) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *TransactionTokeApproval) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *TransactionTokeApproval) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetDustThreshold

`func (o *TransactionTokeApproval) GetDustThreshold() string`

GetDustThreshold returns the DustThreshold field if non-nil, zero value otherwise.

### GetDustThresholdOk

`func (o *TransactionTokeApproval) GetDustThresholdOk() (*string, bool)`

GetDustThresholdOk returns a tuple with the DustThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDustThreshold

`func (o *TransactionTokeApproval) SetDustThreshold(v string)`

SetDustThreshold sets DustThreshold field to given value.

### HasDustThreshold

`func (o *TransactionTokeApproval) HasDustThreshold() bool`

HasDustThreshold returns a boolean if a field has been set.

### GetCustodialMinimumDepositThreshold

`func (o *TransactionTokeApproval) GetCustodialMinimumDepositThreshold() string`

GetCustodialMinimumDepositThreshold returns the CustodialMinimumDepositThreshold field if non-nil, zero value otherwise.

### GetCustodialMinimumDepositThresholdOk

`func (o *TransactionTokeApproval) GetCustodialMinimumDepositThresholdOk() (*string, bool)`

GetCustodialMinimumDepositThresholdOk returns a tuple with the CustodialMinimumDepositThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustodialMinimumDepositThreshold

`func (o *TransactionTokeApproval) SetCustodialMinimumDepositThreshold(v string)`

SetCustodialMinimumDepositThreshold sets CustodialMinimumDepositThreshold field to given value.

### HasCustodialMinimumDepositThreshold

`func (o *TransactionTokeApproval) HasCustodialMinimumDepositThreshold() bool`

HasCustodialMinimumDepositThreshold returns a boolean if a field has been set.

### GetAssetModelType

`func (o *TransactionTokeApproval) GetAssetModelType() TokenAssetModelType`

GetAssetModelType returns the AssetModelType field if non-nil, zero value otherwise.

### GetAssetModelTypeOk

`func (o *TransactionTokeApproval) GetAssetModelTypeOk() (*TokenAssetModelType, bool)`

GetAssetModelTypeOk returns a tuple with the AssetModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetModelType

`func (o *TransactionTokeApproval) SetAssetModelType(v TokenAssetModelType)`

SetAssetModelType sets AssetModelType field to given value.

### HasAssetModelType

`func (o *TransactionTokeApproval) HasAssetModelType() bool`

HasAssetModelType returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionTokeApproval) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionTokeApproval) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionTokeApproval) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionTokeApproval) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetSpender

`func (o *TransactionTokeApproval) GetSpender() string`

GetSpender returns the Spender field if non-nil, zero value otherwise.

### GetSpenderOk

`func (o *TransactionTokeApproval) GetSpenderOk() (*string, bool)`

GetSpenderOk returns a tuple with the Spender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpender

`func (o *TransactionTokeApproval) SetSpender(v string)`

SetSpender sets Spender field to given value.

### HasSpender

`func (o *TransactionTokeApproval) HasSpender() bool`

HasSpender returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


