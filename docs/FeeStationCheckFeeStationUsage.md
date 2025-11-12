# FeeStationCheckFeeStationUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. | 
**Amount** | Pointer to **string** | Applicable to transfer requests only. The amount of tokens to be transferred in this request. | [optional] 
**TokenId** | Pointer to **string** | Applicable to transfer requests only. The token ID of the asset to be transferred.   You can retrieve available token IDs by calling   [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens).  | [optional] 
**FeeTokenId** | Pointer to **string** | The token ID used to pay the gas fee for the main transaction. You can retrieve available token IDs by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). | [optional] 
**EstimatedFeeAmount** | **string** | The estimated transaction fee required for this transfer, before applying any Fee Station rules. | 
**FromAddress** | **string** | The blockchain address that initiates the transfer. | 
**FromWalletId** | **string** | The wallet ID. | 
**AutoFuel** | Pointer to [**AutoFuelType**](AutoFuelType.md) |  | [optional] 

## Methods

### NewFeeStationCheckFeeStationUsage

`func NewFeeStationCheckFeeStationUsage(requestId string, estimatedFeeAmount string, fromAddress string, fromWalletId string, ) *FeeStationCheckFeeStationUsage`

NewFeeStationCheckFeeStationUsage instantiates a new FeeStationCheckFeeStationUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeStationCheckFeeStationUsageWithDefaults

`func NewFeeStationCheckFeeStationUsageWithDefaults() *FeeStationCheckFeeStationUsage`

NewFeeStationCheckFeeStationUsageWithDefaults instantiates a new FeeStationCheckFeeStationUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *FeeStationCheckFeeStationUsage) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *FeeStationCheckFeeStationUsage) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *FeeStationCheckFeeStationUsage) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetAmount

`func (o *FeeStationCheckFeeStationUsage) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FeeStationCheckFeeStationUsage) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FeeStationCheckFeeStationUsage) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *FeeStationCheckFeeStationUsage) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTokenId

`func (o *FeeStationCheckFeeStationUsage) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *FeeStationCheckFeeStationUsage) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *FeeStationCheckFeeStationUsage) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *FeeStationCheckFeeStationUsage) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetFeeTokenId

`func (o *FeeStationCheckFeeStationUsage) GetFeeTokenId() string`

GetFeeTokenId returns the FeeTokenId field if non-nil, zero value otherwise.

### GetFeeTokenIdOk

`func (o *FeeStationCheckFeeStationUsage) GetFeeTokenIdOk() (*string, bool)`

GetFeeTokenIdOk returns a tuple with the FeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeTokenId

`func (o *FeeStationCheckFeeStationUsage) SetFeeTokenId(v string)`

SetFeeTokenId sets FeeTokenId field to given value.

### HasFeeTokenId

`func (o *FeeStationCheckFeeStationUsage) HasFeeTokenId() bool`

HasFeeTokenId returns a boolean if a field has been set.

### GetEstimatedFeeAmount

`func (o *FeeStationCheckFeeStationUsage) GetEstimatedFeeAmount() string`

GetEstimatedFeeAmount returns the EstimatedFeeAmount field if non-nil, zero value otherwise.

### GetEstimatedFeeAmountOk

`func (o *FeeStationCheckFeeStationUsage) GetEstimatedFeeAmountOk() (*string, bool)`

GetEstimatedFeeAmountOk returns a tuple with the EstimatedFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedFeeAmount

`func (o *FeeStationCheckFeeStationUsage) SetEstimatedFeeAmount(v string)`

SetEstimatedFeeAmount sets EstimatedFeeAmount field to given value.


### GetFromAddress

`func (o *FeeStationCheckFeeStationUsage) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *FeeStationCheckFeeStationUsage) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *FeeStationCheckFeeStationUsage) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.


### GetFromWalletId

`func (o *FeeStationCheckFeeStationUsage) GetFromWalletId() string`

GetFromWalletId returns the FromWalletId field if non-nil, zero value otherwise.

### GetFromWalletIdOk

`func (o *FeeStationCheckFeeStationUsage) GetFromWalletIdOk() (*string, bool)`

GetFromWalletIdOk returns a tuple with the FromWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWalletId

`func (o *FeeStationCheckFeeStationUsage) SetFromWalletId(v string)`

SetFromWalletId sets FromWalletId field to given value.


### GetAutoFuel

`func (o *FeeStationCheckFeeStationUsage) GetAutoFuel() AutoFuelType`

GetAutoFuel returns the AutoFuel field if non-nil, zero value otherwise.

### GetAutoFuelOk

`func (o *FeeStationCheckFeeStationUsage) GetAutoFuelOk() (*AutoFuelType, bool)`

GetAutoFuelOk returns a tuple with the AutoFuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoFuel

`func (o *FeeStationCheckFeeStationUsage) SetAutoFuel(v AutoFuelType)`

SetAutoFuel sets AutoFuel field to given value.

### HasAutoFuel

`func (o *FeeStationCheckFeeStationUsage) HasAutoFuel() bool`

HasAutoFuel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


