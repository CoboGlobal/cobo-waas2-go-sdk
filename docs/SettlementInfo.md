# SettlementInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | Pointer to **string** | The merchant ID. For developer balance, this field will be empty. | [optional] 
**TokenId** | Pointer to **string** | The ID of the cryptocurrency. | [optional] 
**AvailableAmount** | **string** | This field is no longer in use and can be ignored. | 
**AvailableCurrencyBalance** | Pointer to **string** | This field is no longer in use and can be ignored. | [optional] 
**PendingAmount** | Pointer to **string** | This field is no longer in use and can be ignored. | [optional] 
**PendingCurrencyBalance** | Pointer to **string** | This field is no longer in use and can be ignored. | [optional] 
**SettledAmount** | Pointer to **string** | The amount already settled, in the specified cryptocurrency. | [optional] 
**AvailableBalance** | Pointer to **string** | The balance available for settlement or refund, in the specified cryptocurrency. | [optional] 
**TotalBalance** | Pointer to **string** |  The total unsettled balance in the specified cryptocurrency, including: - Available balance that can be settled immediately - Amounts below the sweep threshold that require forced sweep before settlement  | [optional] 
**AcquiringType** | Pointer to [**AcquiringType**](AcquiringType.md) |  | [optional] 
**CreatedTimestamp** | Pointer to **int32** | The creation time of the settlement, represented as a UNIX timestamp in seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int32** | The last update time of the settlement, represented as a UNIX timestamp in seconds. | [optional] 

## Methods

### NewSettlementInfo

`func NewSettlementInfo(availableAmount string, ) *SettlementInfo`

NewSettlementInfo instantiates a new SettlementInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettlementInfoWithDefaults

`func NewSettlementInfoWithDefaults() *SettlementInfo`

NewSettlementInfoWithDefaults instantiates a new SettlementInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *SettlementInfo) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *SettlementInfo) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *SettlementInfo) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *SettlementInfo) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetTokenId

`func (o *SettlementInfo) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *SettlementInfo) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *SettlementInfo) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *SettlementInfo) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetAvailableAmount

`func (o *SettlementInfo) GetAvailableAmount() string`

GetAvailableAmount returns the AvailableAmount field if non-nil, zero value otherwise.

### GetAvailableAmountOk

`func (o *SettlementInfo) GetAvailableAmountOk() (*string, bool)`

GetAvailableAmountOk returns a tuple with the AvailableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAmount

`func (o *SettlementInfo) SetAvailableAmount(v string)`

SetAvailableAmount sets AvailableAmount field to given value.


### GetAvailableCurrencyBalance

`func (o *SettlementInfo) GetAvailableCurrencyBalance() string`

GetAvailableCurrencyBalance returns the AvailableCurrencyBalance field if non-nil, zero value otherwise.

### GetAvailableCurrencyBalanceOk

`func (o *SettlementInfo) GetAvailableCurrencyBalanceOk() (*string, bool)`

GetAvailableCurrencyBalanceOk returns a tuple with the AvailableCurrencyBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCurrencyBalance

`func (o *SettlementInfo) SetAvailableCurrencyBalance(v string)`

SetAvailableCurrencyBalance sets AvailableCurrencyBalance field to given value.

### HasAvailableCurrencyBalance

`func (o *SettlementInfo) HasAvailableCurrencyBalance() bool`

HasAvailableCurrencyBalance returns a boolean if a field has been set.

### GetPendingAmount

`func (o *SettlementInfo) GetPendingAmount() string`

GetPendingAmount returns the PendingAmount field if non-nil, zero value otherwise.

### GetPendingAmountOk

`func (o *SettlementInfo) GetPendingAmountOk() (*string, bool)`

GetPendingAmountOk returns a tuple with the PendingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAmount

`func (o *SettlementInfo) SetPendingAmount(v string)`

SetPendingAmount sets PendingAmount field to given value.

### HasPendingAmount

`func (o *SettlementInfo) HasPendingAmount() bool`

HasPendingAmount returns a boolean if a field has been set.

### GetPendingCurrencyBalance

`func (o *SettlementInfo) GetPendingCurrencyBalance() string`

GetPendingCurrencyBalance returns the PendingCurrencyBalance field if non-nil, zero value otherwise.

### GetPendingCurrencyBalanceOk

`func (o *SettlementInfo) GetPendingCurrencyBalanceOk() (*string, bool)`

GetPendingCurrencyBalanceOk returns a tuple with the PendingCurrencyBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCurrencyBalance

`func (o *SettlementInfo) SetPendingCurrencyBalance(v string)`

SetPendingCurrencyBalance sets PendingCurrencyBalance field to given value.

### HasPendingCurrencyBalance

`func (o *SettlementInfo) HasPendingCurrencyBalance() bool`

HasPendingCurrencyBalance returns a boolean if a field has been set.

### GetSettledAmount

`func (o *SettlementInfo) GetSettledAmount() string`

GetSettledAmount returns the SettledAmount field if non-nil, zero value otherwise.

### GetSettledAmountOk

`func (o *SettlementInfo) GetSettledAmountOk() (*string, bool)`

GetSettledAmountOk returns a tuple with the SettledAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAmount

`func (o *SettlementInfo) SetSettledAmount(v string)`

SetSettledAmount sets SettledAmount field to given value.

### HasSettledAmount

`func (o *SettlementInfo) HasSettledAmount() bool`

HasSettledAmount returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *SettlementInfo) GetAvailableBalance() string`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *SettlementInfo) GetAvailableBalanceOk() (*string, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *SettlementInfo) SetAvailableBalance(v string)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *SettlementInfo) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### GetTotalBalance

`func (o *SettlementInfo) GetTotalBalance() string`

GetTotalBalance returns the TotalBalance field if non-nil, zero value otherwise.

### GetTotalBalanceOk

`func (o *SettlementInfo) GetTotalBalanceOk() (*string, bool)`

GetTotalBalanceOk returns a tuple with the TotalBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBalance

`func (o *SettlementInfo) SetTotalBalance(v string)`

SetTotalBalance sets TotalBalance field to given value.

### HasTotalBalance

`func (o *SettlementInfo) HasTotalBalance() bool`

HasTotalBalance returns a boolean if a field has been set.

### GetAcquiringType

`func (o *SettlementInfo) GetAcquiringType() AcquiringType`

GetAcquiringType returns the AcquiringType field if non-nil, zero value otherwise.

### GetAcquiringTypeOk

`func (o *SettlementInfo) GetAcquiringTypeOk() (*AcquiringType, bool)`

GetAcquiringTypeOk returns a tuple with the AcquiringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquiringType

`func (o *SettlementInfo) SetAcquiringType(v AcquiringType)`

SetAcquiringType sets AcquiringType field to given value.

### HasAcquiringType

`func (o *SettlementInfo) HasAcquiringType() bool`

HasAcquiringType returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *SettlementInfo) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *SettlementInfo) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *SettlementInfo) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *SettlementInfo) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *SettlementInfo) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *SettlementInfo) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *SettlementInfo) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *SettlementInfo) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


