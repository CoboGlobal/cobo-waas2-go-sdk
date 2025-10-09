# FeeStationCheckFeeStationUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token used to pay the gas fee for this specific transaction. You can retrieve the IDs of all supported tokens by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | 
**GasStationType** | [**FeeStationGasStationType**](FeeStationGasStationType.md) |  | 
**IsFeeStationApplicable** | **bool** | Indicates whether the fee station is applied for this transfer request. | 
**IsBalanceSufficient** | **bool** | If the fee station is used, indicates whether its balance is sufficient to cover the required gas fee. | 
**Balance** | **string** | The current token balance available in the fee station. | 
**TotalFeeAmount** | **string** | The total gas amount required for this transfer request. | 
**IsSponsorApplicable** | **bool** | Indicates whether USDT (U) sponsorship is applied when the fee station balance is insufficient. | 
**SponsoredFeeAmount** | **string** | The amount of gas fee sponsored by USDT (U) when applicable. | 

## Methods

### NewFeeStationCheckFeeStationUsageResponse

`func NewFeeStationCheckFeeStationUsageResponse(tokenId string, gasStationType FeeStationGasStationType, isFeeStationApplicable bool, isBalanceSufficient bool, balance string, totalFeeAmount string, isSponsorApplicable bool, sponsoredFeeAmount string, ) *FeeStationCheckFeeStationUsageResponse`

NewFeeStationCheckFeeStationUsageResponse instantiates a new FeeStationCheckFeeStationUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeStationCheckFeeStationUsageResponseWithDefaults

`func NewFeeStationCheckFeeStationUsageResponseWithDefaults() *FeeStationCheckFeeStationUsageResponse`

NewFeeStationCheckFeeStationUsageResponseWithDefaults instantiates a new FeeStationCheckFeeStationUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *FeeStationCheckFeeStationUsageResponse) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *FeeStationCheckFeeStationUsageResponse) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *FeeStationCheckFeeStationUsageResponse) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetGasStationType

`func (o *FeeStationCheckFeeStationUsageResponse) GetGasStationType() FeeStationGasStationType`

GetGasStationType returns the GasStationType field if non-nil, zero value otherwise.

### GetGasStationTypeOk

`func (o *FeeStationCheckFeeStationUsageResponse) GetGasStationTypeOk() (*FeeStationGasStationType, bool)`

GetGasStationTypeOk returns a tuple with the GasStationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasStationType

`func (o *FeeStationCheckFeeStationUsageResponse) SetGasStationType(v FeeStationGasStationType)`

SetGasStationType sets GasStationType field to given value.


### GetIsFeeStationApplicable

`func (o *FeeStationCheckFeeStationUsageResponse) GetIsFeeStationApplicable() bool`

GetIsFeeStationApplicable returns the IsFeeStationApplicable field if non-nil, zero value otherwise.

### GetIsFeeStationApplicableOk

`func (o *FeeStationCheckFeeStationUsageResponse) GetIsFeeStationApplicableOk() (*bool, bool)`

GetIsFeeStationApplicableOk returns a tuple with the IsFeeStationApplicable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFeeStationApplicable

`func (o *FeeStationCheckFeeStationUsageResponse) SetIsFeeStationApplicable(v bool)`

SetIsFeeStationApplicable sets IsFeeStationApplicable field to given value.


### GetIsBalanceSufficient

`func (o *FeeStationCheckFeeStationUsageResponse) GetIsBalanceSufficient() bool`

GetIsBalanceSufficient returns the IsBalanceSufficient field if non-nil, zero value otherwise.

### GetIsBalanceSufficientOk

`func (o *FeeStationCheckFeeStationUsageResponse) GetIsBalanceSufficientOk() (*bool, bool)`

GetIsBalanceSufficientOk returns a tuple with the IsBalanceSufficient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBalanceSufficient

`func (o *FeeStationCheckFeeStationUsageResponse) SetIsBalanceSufficient(v bool)`

SetIsBalanceSufficient sets IsBalanceSufficient field to given value.


### GetBalance

`func (o *FeeStationCheckFeeStationUsageResponse) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *FeeStationCheckFeeStationUsageResponse) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *FeeStationCheckFeeStationUsageResponse) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetTotalFeeAmount

`func (o *FeeStationCheckFeeStationUsageResponse) GetTotalFeeAmount() string`

GetTotalFeeAmount returns the TotalFeeAmount field if non-nil, zero value otherwise.

### GetTotalFeeAmountOk

`func (o *FeeStationCheckFeeStationUsageResponse) GetTotalFeeAmountOk() (*string, bool)`

GetTotalFeeAmountOk returns a tuple with the TotalFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFeeAmount

`func (o *FeeStationCheckFeeStationUsageResponse) SetTotalFeeAmount(v string)`

SetTotalFeeAmount sets TotalFeeAmount field to given value.


### GetIsSponsorApplicable

`func (o *FeeStationCheckFeeStationUsageResponse) GetIsSponsorApplicable() bool`

GetIsSponsorApplicable returns the IsSponsorApplicable field if non-nil, zero value otherwise.

### GetIsSponsorApplicableOk

`func (o *FeeStationCheckFeeStationUsageResponse) GetIsSponsorApplicableOk() (*bool, bool)`

GetIsSponsorApplicableOk returns a tuple with the IsSponsorApplicable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSponsorApplicable

`func (o *FeeStationCheckFeeStationUsageResponse) SetIsSponsorApplicable(v bool)`

SetIsSponsorApplicable sets IsSponsorApplicable field to given value.


### GetSponsoredFeeAmount

`func (o *FeeStationCheckFeeStationUsageResponse) GetSponsoredFeeAmount() string`

GetSponsoredFeeAmount returns the SponsoredFeeAmount field if non-nil, zero value otherwise.

### GetSponsoredFeeAmountOk

`func (o *FeeStationCheckFeeStationUsageResponse) GetSponsoredFeeAmountOk() (*string, bool)`

GetSponsoredFeeAmountOk returns a tuple with the SponsoredFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsoredFeeAmount

`func (o *FeeStationCheckFeeStationUsageResponse) SetSponsoredFeeAmount(v string)`

SetSponsoredFeeAmount sets SponsoredFeeAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


