# TokenInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](/v2/api-references/wallets/list-enabled-tokens). | 
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

## Methods

### NewTokenInfo

`func NewTokenInfo(tokenId string, chainId string, ) *TokenInfo`

NewTokenInfo instantiates a new TokenInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenInfoWithDefaults

`func NewTokenInfoWithDefaults() *TokenInfo`

NewTokenInfoWithDefaults instantiates a new TokenInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *TokenInfo) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenInfo) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenInfo) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChainId

`func (o *TokenInfo) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TokenInfo) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TokenInfo) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetAssetId

`func (o *TokenInfo) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *TokenInfo) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *TokenInfo) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *TokenInfo) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetSymbol

`func (o *TokenInfo) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenInfo) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenInfo) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TokenInfo) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetName

`func (o *TokenInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TokenInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDecimal

`func (o *TokenInfo) GetDecimal() int32`

GetDecimal returns the Decimal field if non-nil, zero value otherwise.

### GetDecimalOk

`func (o *TokenInfo) GetDecimalOk() (*int32, bool)`

GetDecimalOk returns a tuple with the Decimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimal

`func (o *TokenInfo) SetDecimal(v int32)`

SetDecimal sets Decimal field to given value.

### HasDecimal

`func (o *TokenInfo) HasDecimal() bool`

HasDecimal returns a boolean if a field has been set.

### GetIconUrl

`func (o *TokenInfo) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *TokenInfo) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *TokenInfo) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *TokenInfo) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetTokenAddress

`func (o *TokenInfo) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TokenInfo) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TokenInfo) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.

### HasTokenAddress

`func (o *TokenInfo) HasTokenAddress() bool`

HasTokenAddress returns a boolean if a field has been set.

### GetFeeTokenId

`func (o *TokenInfo) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *TokenInfo) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *TokenInfo) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.

### HasFeeTokenId

`func (o *TokenInfo) HasFeeTokenId() bool`

HasFeeTokenId returns a boolean if a field has been set.

### GetCanDeposit

`func (o *TokenInfo) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *TokenInfo) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *TokenInfo) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *TokenInfo) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *TokenInfo) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *TokenInfo) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *TokenInfo) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *TokenInfo) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetDustThreshold

`func (o *TokenInfo) GetDustThreshold() string`

GetDustThreshold returns the DustThreshold field if non-nil, zero value otherwise.

### GetDustThresholdOk

`func (o *TokenInfo) GetDustThresholdOk() (*string, bool)`

GetDustThresholdOk returns a tuple with the DustThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDustThreshold

`func (o *TokenInfo) SetDustThreshold(v string)`

SetDustThreshold sets DustThreshold field to given value.

### HasDustThreshold

`func (o *TokenInfo) HasDustThreshold() bool`

HasDustThreshold returns a boolean if a field has been set.

### GetCustodialMinimumDepositThreshold

`func (o *TokenInfo) GetCustodialMinimumDepositThreshold() string`

GetCustodialMinimumDepositThreshold returns the CustodialMinimumDepositThreshold field if non-nil, zero value otherwise.

### GetCustodialMinimumDepositThresholdOk

`func (o *TokenInfo) GetCustodialMinimumDepositThresholdOk() (*string, bool)`

GetCustodialMinimumDepositThresholdOk returns a tuple with the CustodialMinimumDepositThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustodialMinimumDepositThreshold

`func (o *TokenInfo) SetCustodialMinimumDepositThreshold(v string)`

SetCustodialMinimumDepositThreshold sets CustodialMinimumDepositThreshold field to given value.

### HasCustodialMinimumDepositThreshold

`func (o *TokenInfo) HasCustodialMinimumDepositThreshold() bool`

HasCustodialMinimumDepositThreshold returns a boolean if a field has been set.

### GetAssetModelType

`func (o *TokenInfo) GetAssetModelType() TokenAssetModelType`

GetAssetModelType returns the AssetModelType field if non-nil, zero value otherwise.

### GetAssetModelTypeOk

`func (o *TokenInfo) GetAssetModelTypeOk() (*TokenAssetModelType, bool)`

GetAssetModelTypeOk returns a tuple with the AssetModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetModelType

`func (o *TokenInfo) SetAssetModelType(v TokenAssetModelType)`

SetAssetModelType sets AssetModelType field to given value.

### HasAssetModelType

`func (o *TokenInfo) HasAssetModelType() bool`

HasAssetModelType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


