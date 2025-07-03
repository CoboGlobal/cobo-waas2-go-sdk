# TopUpAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The dedicated top-up address assigned to a specific payer under a merchant on a specified chain. | 
**PayerId** | **string** | A unique identifier assigned by Cobo to track and identify individual payers. | 
**CustomPayerId** | **string** | A unique identifier assigned by the developer to track and identify individual payers in their system. | 
**MerchantId** | **string** | The merchant ID. | 
**TokenId** | **string** | The token ID, which is a unique identifier that specifies both the blockchain network and cryptocurrency token in the format &#x60;{CHAIN}_{TOKEN}&#x60;. | 
**Chain** | Pointer to **string** | The chain ID. | [optional] 
**DeveloperFeeRate** | Pointer to **string** | The developer fee rate applied to top-up transactions made to this address. Expressed as a decimal string where \&quot;0.1\&quot; represents 10%. | [optional] 
**MinAmount** | **string** | The minimum top-up amount allowed for this address. Top-ups below this threshold will not be credited to merchant funds, but to developer funds instead. | 
**CreatedTimestamp** | Pointer to **int32** | The creation time of the top-up address, represented as a UNIX timestamp in seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int32** | The last update time of the top-up address, represented as a UNIX timestamp in seconds. | [optional] 

## Methods

### NewTopUpAddress

`func NewTopUpAddress(address string, payerId string, customPayerId string, merchantId string, tokenId string, minAmount string, ) *TopUpAddress`

NewTopUpAddress instantiates a new TopUpAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopUpAddressWithDefaults

`func NewTopUpAddressWithDefaults() *TopUpAddress`

NewTopUpAddressWithDefaults instantiates a new TopUpAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TopUpAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TopUpAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TopUpAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPayerId

`func (o *TopUpAddress) GetPayerId() string`

GetPayerId returns the PayerId field if non-nil, zero value otherwise.

### GetPayerIdOk

`func (o *TopUpAddress) GetPayerIdOk() (*string, bool)`

GetPayerIdOk returns a tuple with the PayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerId

`func (o *TopUpAddress) SetPayerId(v string)`

SetPayerId sets PayerId field to given value.


### GetCustomPayerId

`func (o *TopUpAddress) GetCustomPayerId() string`

GetCustomPayerId returns the CustomPayerId field if non-nil, zero value otherwise.

### GetCustomPayerIdOk

`func (o *TopUpAddress) GetCustomPayerIdOk() (*string, bool)`

GetCustomPayerIdOk returns a tuple with the CustomPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayerId

`func (o *TopUpAddress) SetCustomPayerId(v string)`

SetCustomPayerId sets CustomPayerId field to given value.


### GetMerchantId

`func (o *TopUpAddress) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *TopUpAddress) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *TopUpAddress) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetTokenId

`func (o *TopUpAddress) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TopUpAddress) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TopUpAddress) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetChain

`func (o *TopUpAddress) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *TopUpAddress) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *TopUpAddress) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *TopUpAddress) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetDeveloperFeeRate

`func (o *TopUpAddress) GetDeveloperFeeRate() string`

GetDeveloperFeeRate returns the DeveloperFeeRate field if non-nil, zero value otherwise.

### GetDeveloperFeeRateOk

`func (o *TopUpAddress) GetDeveloperFeeRateOk() (*string, bool)`

GetDeveloperFeeRateOk returns a tuple with the DeveloperFeeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperFeeRate

`func (o *TopUpAddress) SetDeveloperFeeRate(v string)`

SetDeveloperFeeRate sets DeveloperFeeRate field to given value.

### HasDeveloperFeeRate

`func (o *TopUpAddress) HasDeveloperFeeRate() bool`

HasDeveloperFeeRate returns a boolean if a field has been set.

### GetMinAmount

`func (o *TopUpAddress) GetMinAmount() string`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *TopUpAddress) GetMinAmountOk() (*string, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *TopUpAddress) SetMinAmount(v string)`

SetMinAmount sets MinAmount field to given value.


### GetCreatedTimestamp

`func (o *TopUpAddress) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *TopUpAddress) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *TopUpAddress) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *TopUpAddress) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *TopUpAddress) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *TopUpAddress) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *TopUpAddress) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *TopUpAddress) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


