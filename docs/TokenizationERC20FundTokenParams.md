# TokenizationERC20FundTokenParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Standard** | [**TokenizationTokenStandard**](TokenizationTokenStandard.md) |  | 
**Name** | **string** | The name of the fund token. | 
**Symbol** | **string** | The symbol of the fund token. | 
**AssetToken** | **string** | The address of the underlying asset token (e.g., XAUT, USDC). The fund will hold this token as collateral. | 
**InitialNav** | Pointer to **string** | Initial net asset value (NAV) per share (optional). Default: &#39;1.0&#39; (standard for new funds). | [optional] [default to "1.0"]
**InitialAnnualRate** | Pointer to **string** | Initial annual rate (optional). Can be updated later via NAV updater. Default: &#39;0&#39;. | [optional] [default to "0"]
**MinDeposit** | Pointer to **string** | Minimum deposit amount (optional). Default: &#39;0&#39; (no minimum, accepts any amount &gt; 0). Admin can update this later. | [optional] [default to "0"]
**MinRedemption** | Pointer to **string** | Minimum redemption amount (optional). Default: &#39;0&#39; (no minimum, accepts any amount &gt; 0). Admin can update this later. | [optional] [default to "0"]
**MaxAnnualRate** | Pointer to **string** | Maximum allowed annual rate (optional). Default: type(uint256).max (no limit). Set lower for conservative funds (e.g., &#39;0.2&#39; for 20%). | [optional] [default to "115792089237316195423570985008687907853269984665640564039457584007913129639935"]
**MaxRateChange** | Pointer to **string** | Maximum rate change per NAV update (optional). Default: type(uint256).max (no limit). Set lower to prevent volatility (e.g., &#39;0.05&#39; for 5%). | [optional] [default to "115792089237316195423570985008687907853269984665640564039457584007913129639935"]
**MinUpdateIntervalSeconds** | Pointer to **int32** | Minimum interval between NAV updates in seconds (optional). Default: 0 (no minimum). Set higher to prevent frequent updates (e.g., 86400 for 1 day). | [optional] [default to 0]
**Decimals** | Pointer to **int32** | The number of decimals for the token (0-18). | [optional] 
**Permissions** | Pointer to [**TokenizationERC20FundTokenPermissionParams**](TokenizationERC20FundTokenPermissionParams.md) |  | [optional] 

## Methods

### NewTokenizationERC20FundTokenParams

`func NewTokenizationERC20FundTokenParams(standard TokenizationTokenStandard, name string, symbol string, assetToken string, ) *TokenizationERC20FundTokenParams`

NewTokenizationERC20FundTokenParams instantiates a new TokenizationERC20FundTokenParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizationERC20FundTokenParamsWithDefaults

`func NewTokenizationERC20FundTokenParamsWithDefaults() *TokenizationERC20FundTokenParams`

NewTokenizationERC20FundTokenParamsWithDefaults instantiates a new TokenizationERC20FundTokenParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandard

`func (o *TokenizationERC20FundTokenParams) GetStandard() TokenizationTokenStandard`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *TokenizationERC20FundTokenParams) GetStandardOk() (*TokenizationTokenStandard, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *TokenizationERC20FundTokenParams) SetStandard(v TokenizationTokenStandard)`

SetStandard sets Standard field to given value.


### GetName

`func (o *TokenizationERC20FundTokenParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenizationERC20FundTokenParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenizationERC20FundTokenParams) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *TokenizationERC20FundTokenParams) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenizationERC20FundTokenParams) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenizationERC20FundTokenParams) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetAssetToken

`func (o *TokenizationERC20FundTokenParams) GetAssetToken() string`

GetAssetToken returns the AssetToken field if non-nil, zero value otherwise.

### GetAssetTokenOk

`func (o *TokenizationERC20FundTokenParams) GetAssetTokenOk() (*string, bool)`

GetAssetTokenOk returns a tuple with the AssetToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetToken

`func (o *TokenizationERC20FundTokenParams) SetAssetToken(v string)`

SetAssetToken sets AssetToken field to given value.


### GetInitialNav

`func (o *TokenizationERC20FundTokenParams) GetInitialNav() string`

GetInitialNav returns the InitialNav field if non-nil, zero value otherwise.

### GetInitialNavOk

`func (o *TokenizationERC20FundTokenParams) GetInitialNavOk() (*string, bool)`

GetInitialNavOk returns a tuple with the InitialNav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialNav

`func (o *TokenizationERC20FundTokenParams) SetInitialNav(v string)`

SetInitialNav sets InitialNav field to given value.

### HasInitialNav

`func (o *TokenizationERC20FundTokenParams) HasInitialNav() bool`

HasInitialNav returns a boolean if a field has been set.

### GetInitialAnnualRate

`func (o *TokenizationERC20FundTokenParams) GetInitialAnnualRate() string`

GetInitialAnnualRate returns the InitialAnnualRate field if non-nil, zero value otherwise.

### GetInitialAnnualRateOk

`func (o *TokenizationERC20FundTokenParams) GetInitialAnnualRateOk() (*string, bool)`

GetInitialAnnualRateOk returns a tuple with the InitialAnnualRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialAnnualRate

`func (o *TokenizationERC20FundTokenParams) SetInitialAnnualRate(v string)`

SetInitialAnnualRate sets InitialAnnualRate field to given value.

### HasInitialAnnualRate

`func (o *TokenizationERC20FundTokenParams) HasInitialAnnualRate() bool`

HasInitialAnnualRate returns a boolean if a field has been set.

### GetMinDeposit

`func (o *TokenizationERC20FundTokenParams) GetMinDeposit() string`

GetMinDeposit returns the MinDeposit field if non-nil, zero value otherwise.

### GetMinDepositOk

`func (o *TokenizationERC20FundTokenParams) GetMinDepositOk() (*string, bool)`

GetMinDepositOk returns a tuple with the MinDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDeposit

`func (o *TokenizationERC20FundTokenParams) SetMinDeposit(v string)`

SetMinDeposit sets MinDeposit field to given value.

### HasMinDeposit

`func (o *TokenizationERC20FundTokenParams) HasMinDeposit() bool`

HasMinDeposit returns a boolean if a field has been set.

### GetMinRedemption

`func (o *TokenizationERC20FundTokenParams) GetMinRedemption() string`

GetMinRedemption returns the MinRedemption field if non-nil, zero value otherwise.

### GetMinRedemptionOk

`func (o *TokenizationERC20FundTokenParams) GetMinRedemptionOk() (*string, bool)`

GetMinRedemptionOk returns a tuple with the MinRedemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRedemption

`func (o *TokenizationERC20FundTokenParams) SetMinRedemption(v string)`

SetMinRedemption sets MinRedemption field to given value.

### HasMinRedemption

`func (o *TokenizationERC20FundTokenParams) HasMinRedemption() bool`

HasMinRedemption returns a boolean if a field has been set.

### GetMaxAnnualRate

`func (o *TokenizationERC20FundTokenParams) GetMaxAnnualRate() string`

GetMaxAnnualRate returns the MaxAnnualRate field if non-nil, zero value otherwise.

### GetMaxAnnualRateOk

`func (o *TokenizationERC20FundTokenParams) GetMaxAnnualRateOk() (*string, bool)`

GetMaxAnnualRateOk returns a tuple with the MaxAnnualRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAnnualRate

`func (o *TokenizationERC20FundTokenParams) SetMaxAnnualRate(v string)`

SetMaxAnnualRate sets MaxAnnualRate field to given value.

### HasMaxAnnualRate

`func (o *TokenizationERC20FundTokenParams) HasMaxAnnualRate() bool`

HasMaxAnnualRate returns a boolean if a field has been set.

### GetMaxRateChange

`func (o *TokenizationERC20FundTokenParams) GetMaxRateChange() string`

GetMaxRateChange returns the MaxRateChange field if non-nil, zero value otherwise.

### GetMaxRateChangeOk

`func (o *TokenizationERC20FundTokenParams) GetMaxRateChangeOk() (*string, bool)`

GetMaxRateChangeOk returns a tuple with the MaxRateChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRateChange

`func (o *TokenizationERC20FundTokenParams) SetMaxRateChange(v string)`

SetMaxRateChange sets MaxRateChange field to given value.

### HasMaxRateChange

`func (o *TokenizationERC20FundTokenParams) HasMaxRateChange() bool`

HasMaxRateChange returns a boolean if a field has been set.

### GetMinUpdateIntervalSeconds

`func (o *TokenizationERC20FundTokenParams) GetMinUpdateIntervalSeconds() int32`

GetMinUpdateIntervalSeconds returns the MinUpdateIntervalSeconds field if non-nil, zero value otherwise.

### GetMinUpdateIntervalSecondsOk

`func (o *TokenizationERC20FundTokenParams) GetMinUpdateIntervalSecondsOk() (*int32, bool)`

GetMinUpdateIntervalSecondsOk returns a tuple with the MinUpdateIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUpdateIntervalSeconds

`func (o *TokenizationERC20FundTokenParams) SetMinUpdateIntervalSeconds(v int32)`

SetMinUpdateIntervalSeconds sets MinUpdateIntervalSeconds field to given value.

### HasMinUpdateIntervalSeconds

`func (o *TokenizationERC20FundTokenParams) HasMinUpdateIntervalSeconds() bool`

HasMinUpdateIntervalSeconds returns a boolean if a field has been set.

### GetDecimals

`func (o *TokenizationERC20FundTokenParams) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenizationERC20FundTokenParams) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenizationERC20FundTokenParams) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *TokenizationERC20FundTokenParams) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetPermissions

`func (o *TokenizationERC20FundTokenParams) GetPermissions() TokenizationERC20FundTokenPermissionParams`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *TokenizationERC20FundTokenParams) GetPermissionsOk() (*TokenizationERC20FundTokenPermissionParams, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *TokenizationERC20FundTokenParams) SetPermissions(v TokenizationERC20FundTokenPermissionParams)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *TokenizationERC20FundTokenParams) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


