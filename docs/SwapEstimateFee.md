# SwapEstimateFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | The ID of the wallet to pay. | 
**Address** | Pointer to **string** | The wallet address. This property is required when the wallet to pay is not a Custodial Wallet (Asset Wallet). | [optional] 
**QuoteId** | **string** | The ID of the swap quote. | 
**FeeType** | Pointer to [**FeeType**](FeeType.md) |  | [optional] [default to FEETYPE_EVM_EIP_1559]

## Methods

### NewSwapEstimateFee

`func NewSwapEstimateFee(walletId string, quoteId string, ) *SwapEstimateFee`

NewSwapEstimateFee instantiates a new SwapEstimateFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwapEstimateFeeWithDefaults

`func NewSwapEstimateFeeWithDefaults() *SwapEstimateFee`

NewSwapEstimateFeeWithDefaults instantiates a new SwapEstimateFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *SwapEstimateFee) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *SwapEstimateFee) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *SwapEstimateFee) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *SwapEstimateFee) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SwapEstimateFee) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SwapEstimateFee) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SwapEstimateFee) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetQuoteId

`func (o *SwapEstimateFee) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *SwapEstimateFee) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *SwapEstimateFee) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.


### GetFeeType

`func (o *SwapEstimateFee) GetFeeType() FeeType`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *SwapEstimateFee) GetFeeTypeOk() (*FeeType, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *SwapEstimateFee) SetFeeType(v FeeType)`

SetFeeType sets FeeType field to given value.

### HasFeeType

`func (o *SwapEstimateFee) HasFeeType() bool`

HasFeeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


