# TokenizationIssueTokenParamsTokenParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Standard** | [**TokenizationTokenStandard**](TokenizationTokenStandard.md) |  | 
**Name** | **string** | The name of the fund token. | 
**Symbol** | **string** | The symbol of the fund token. | 
**Decimals** | **int32** | The number of decimals for the token (0-18). | 
**TokenAccessActivated** | Pointer to **bool** | Whether the allowlist feature is activated for the token. When activated, only addresses in the allowlist can perform token operations. | [optional] [default to false]
**Permissions** | Pointer to [**TokenizationERC20FundTokenPermissionParams**](TokenizationERC20FundTokenPermissionParams.md) |  | [optional] 
**UnderlyingToken** | **string** | The address of the underlying token that this tokenized asset represents. | 
**AssetToken** | **string** | The address of the underlying asset token (e.g., XAUT, USDC). The fund will hold this token as collateral. | 
**InitialNav** | Pointer to **string** | Initial net asset value (NAV) per share (optional). Default: &#39;1.0&#39; (standard for new funds). | [optional] [default to "1.0"]
**InitialAnnualRate** | Pointer to **string** | Initial annual rate (optional). Can be updated later via NAV updater. Default: &#39;0&#39;. | [optional] [default to "0"]
**MinDeposit** | Pointer to **string** | Minimum deposit amount (optional). Default: &#39;0&#39; (no minimum, accepts any amount &gt; 0). Admin can update this later. | [optional] [default to "0"]
**MinRedemption** | Pointer to **string** | Minimum redemption amount (optional). Default: &#39;0&#39; (no minimum, accepts any amount &gt; 0). Admin can update this later. | [optional] [default to "0"]
**MaxAnnualRate** | Pointer to **string** | Maximum allowed annual rate (optional). Default: type(uint256).max (no limit). Set lower for conservative funds (e.g., &#39;0.2&#39; for 20%). | [optional] [default to "115792089237316195423570985008687907853269984665640564039457584007913129639935"]
**MaxRateChange** | Pointer to **string** | Maximum rate change per NAV update (optional). Default: type(uint256).max (no limit). Set lower to prevent volatility (e.g., &#39;0.05&#39; for 5%). | [optional] [default to "115792089237316195423570985008687907853269984665640564039457584007913129639935"]
**MinUpdateIntervalSeconds** | Pointer to **int32** | Minimum interval between NAV updates in seconds (optional). Default: 0 (no minimum). Set higher to prevent frequent updates (e.g., 86400 for 1 day). | [optional] [default to 0]

## Methods

### NewTokenizationIssueTokenParamsTokenParams

`func NewTokenizationIssueTokenParamsTokenParams(standard TokenizationTokenStandard, name string, symbol string, decimals int32, underlyingToken string, assetToken string, ) *TokenizationIssueTokenParamsTokenParams`

NewTokenizationIssueTokenParamsTokenParams instantiates a new TokenizationIssueTokenParamsTokenParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationIssueTokenParamsTokenParamsWithDefaults

`func NewTokenizationIssueTokenParamsTokenParamsWithDefaults() *TokenizationIssueTokenParamsTokenParams`

NewTokenizationIssueTokenParamsTokenParamsWithDefaults instantiates a new TokenizationIssueTokenParamsTokenParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandard

`func (o *TokenizationIssueTokenParamsTokenParams) GetStandard() TokenizationTokenStandard`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetStandardOk() (*TokenizationTokenStandard, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *TokenizationIssueTokenParamsTokenParams) SetStandard(v TokenizationTokenStandard)`

SetStandard sets Standard field to given value.


### GetName

`func (o *TokenizationIssueTokenParamsTokenParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenizationIssueTokenParamsTokenParams) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *TokenizationIssueTokenParamsTokenParams) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenizationIssueTokenParamsTokenParams) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDecimals

`func (o *TokenizationIssueTokenParamsTokenParams) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenizationIssueTokenParamsTokenParams) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.


### GetTokenAccessActivated

`func (o *TokenizationIssueTokenParamsTokenParams) GetTokenAccessActivated() bool`

GetTokenAccessActivated returns the TokenAccessActivated field if non-nil, zero value otherwise.

### GetTokenAccessActivatedOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetTokenAccessActivatedOk() (*bool, bool)`

GetTokenAccessActivatedOk returns a tuple with the TokenAccessActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAccessActivated

`func (o *TokenizationIssueTokenParamsTokenParams) SetTokenAccessActivated(v bool)`

SetTokenAccessActivated sets TokenAccessActivated field to given value.

### HasTokenAccessActivated

`func (o *TokenizationIssueTokenParamsTokenParams) HasTokenAccessActivated() bool`

HasTokenAccessActivated returns a boolean if a field has been set.

### GetPermissions

`func (o *TokenizationIssueTokenParamsTokenParams) GetPermissions() TokenizationERC20FundTokenPermissionParams`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetPermissionsOk() (*TokenizationERC20FundTokenPermissionParams, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *TokenizationIssueTokenParamsTokenParams) SetPermissions(v TokenizationERC20FundTokenPermissionParams)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *TokenizationIssueTokenParamsTokenParams) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetUnderlyingToken

`func (o *TokenizationIssueTokenParamsTokenParams) GetUnderlyingToken() string`

GetUnderlyingToken returns the UnderlyingToken field if non-nil, zero value otherwise.

### GetUnderlyingTokenOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetUnderlyingTokenOk() (*string, bool)`

GetUnderlyingTokenOk returns a tuple with the UnderlyingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingToken

`func (o *TokenizationIssueTokenParamsTokenParams) SetUnderlyingToken(v string)`

SetUnderlyingToken sets UnderlyingToken field to given value.


### GetAssetToken

`func (o *TokenizationIssueTokenParamsTokenParams) GetAssetToken() string`

GetAssetToken returns the AssetToken field if non-nil, zero value otherwise.

### GetAssetTokenOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetAssetTokenOk() (*string, bool)`

GetAssetTokenOk returns a tuple with the AssetToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetToken

`func (o *TokenizationIssueTokenParamsTokenParams) SetAssetToken(v string)`

SetAssetToken sets AssetToken field to given value.


### GetInitialNav

`func (o *TokenizationIssueTokenParamsTokenParams) GetInitialNav() string`

GetInitialNav returns the InitialNav field if non-nil, zero value otherwise.

### GetInitialNavOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetInitialNavOk() (*string, bool)`

GetInitialNavOk returns a tuple with the InitialNav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialNav

`func (o *TokenizationIssueTokenParamsTokenParams) SetInitialNav(v string)`

SetInitialNav sets InitialNav field to given value.

### HasInitialNav

`func (o *TokenizationIssueTokenParamsTokenParams) HasInitialNav() bool`

HasInitialNav returns a boolean if a field has been set.

### GetInitialAnnualRate

`func (o *TokenizationIssueTokenParamsTokenParams) GetInitialAnnualRate() string`

GetInitialAnnualRate returns the InitialAnnualRate field if non-nil, zero value otherwise.

### GetInitialAnnualRateOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetInitialAnnualRateOk() (*string, bool)`

GetInitialAnnualRateOk returns a tuple with the InitialAnnualRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialAnnualRate

`func (o *TokenizationIssueTokenParamsTokenParams) SetInitialAnnualRate(v string)`

SetInitialAnnualRate sets InitialAnnualRate field to given value.

### HasInitialAnnualRate

`func (o *TokenizationIssueTokenParamsTokenParams) HasInitialAnnualRate() bool`

HasInitialAnnualRate returns a boolean if a field has been set.

### GetMinDeposit

`func (o *TokenizationIssueTokenParamsTokenParams) GetMinDeposit() string`

GetMinDeposit returns the MinDeposit field if non-nil, zero value otherwise.

### GetMinDepositOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetMinDepositOk() (*string, bool)`

GetMinDepositOk returns a tuple with the MinDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDeposit

`func (o *TokenizationIssueTokenParamsTokenParams) SetMinDeposit(v string)`

SetMinDeposit sets MinDeposit field to given value.

### HasMinDeposit

`func (o *TokenizationIssueTokenParamsTokenParams) HasMinDeposit() bool`

HasMinDeposit returns a boolean if a field has been set.

### GetMinRedemption

`func (o *TokenizationIssueTokenParamsTokenParams) GetMinRedemption() string`

GetMinRedemption returns the MinRedemption field if non-nil, zero value otherwise.

### GetMinRedemptionOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetMinRedemptionOk() (*string, bool)`

GetMinRedemptionOk returns a tuple with the MinRedemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRedemption

`func (o *TokenizationIssueTokenParamsTokenParams) SetMinRedemption(v string)`

SetMinRedemption sets MinRedemption field to given value.

### HasMinRedemption

`func (o *TokenizationIssueTokenParamsTokenParams) HasMinRedemption() bool`

HasMinRedemption returns a boolean if a field has been set.

### GetMaxAnnualRate

`func (o *TokenizationIssueTokenParamsTokenParams) GetMaxAnnualRate() string`

GetMaxAnnualRate returns the MaxAnnualRate field if non-nil, zero value otherwise.

### GetMaxAnnualRateOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetMaxAnnualRateOk() (*string, bool)`

GetMaxAnnualRateOk returns a tuple with the MaxAnnualRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAnnualRate

`func (o *TokenizationIssueTokenParamsTokenParams) SetMaxAnnualRate(v string)`

SetMaxAnnualRate sets MaxAnnualRate field to given value.

### HasMaxAnnualRate

`func (o *TokenizationIssueTokenParamsTokenParams) HasMaxAnnualRate() bool`

HasMaxAnnualRate returns a boolean if a field has been set.

### GetMaxRateChange

`func (o *TokenizationIssueTokenParamsTokenParams) GetMaxRateChange() string`

GetMaxRateChange returns the MaxRateChange field if non-nil, zero value otherwise.

### GetMaxRateChangeOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetMaxRateChangeOk() (*string, bool)`

GetMaxRateChangeOk returns a tuple with the MaxRateChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRateChange

`func (o *TokenizationIssueTokenParamsTokenParams) SetMaxRateChange(v string)`

SetMaxRateChange sets MaxRateChange field to given value.

### HasMaxRateChange

`func (o *TokenizationIssueTokenParamsTokenParams) HasMaxRateChange() bool`

HasMaxRateChange returns a boolean if a field has been set.

### GetMinUpdateIntervalSeconds

`func (o *TokenizationIssueTokenParamsTokenParams) GetMinUpdateIntervalSeconds() int32`

GetMinUpdateIntervalSeconds returns the MinUpdateIntervalSeconds field if non-nil, zero value otherwise.

### GetMinUpdateIntervalSecondsOk

`func (o *TokenizationIssueTokenParamsTokenParams) GetMinUpdateIntervalSecondsOk() (*int32, bool)`

GetMinUpdateIntervalSecondsOk returns a tuple with the MinUpdateIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUpdateIntervalSeconds

`func (o *TokenizationIssueTokenParamsTokenParams) SetMinUpdateIntervalSeconds(v int32)`

SetMinUpdateIntervalSeconds sets MinUpdateIntervalSeconds field to given value.

### HasMinUpdateIntervalSeconds

`func (o *TokenizationIssueTokenParamsTokenParams) HasMinUpdateIntervalSeconds() bool`

HasMinUpdateIntervalSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


