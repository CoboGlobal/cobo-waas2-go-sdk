# DestinationWalletAddressDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationId** | **string** | The destination ID. | 
**DestinationName** | **string** | The name of the destination. | 
**DestinationType** | [**DestinationType**](DestinationType.md) |  | 
**MerchantId** | Pointer to **string** | The ID of the merchant linked to the destination. | [optional] 
**WalletAddressId** | **string** | The wallet address ID. | 
**Address** | **string** | The wallet address. | 
**ChainId** | **string** | The chain ID of the cryptocurrency. | 
**UpdatedTimestamp** | **int32** | The updated time of the wallet address, represented as a UNIX timestamp in seconds. | 

## Methods

### NewDestinationWalletAddressDetail

`func NewDestinationWalletAddressDetail(destinationId string, destinationName string, destinationType DestinationType, walletAddressId string, address string, chainId string, updatedTimestamp int32, ) *DestinationWalletAddressDetail`

NewDestinationWalletAddressDetail instantiates a new DestinationWalletAddressDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationWalletAddressDetailWithDefaults

`func NewDestinationWalletAddressDetailWithDefaults() *DestinationWalletAddressDetail`

NewDestinationWalletAddressDetailWithDefaults instantiates a new DestinationWalletAddressDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationId

`func (o *DestinationWalletAddressDetail) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *DestinationWalletAddressDetail) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *DestinationWalletAddressDetail) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetDestinationName

`func (o *DestinationWalletAddressDetail) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *DestinationWalletAddressDetail) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *DestinationWalletAddressDetail) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.


### GetDestinationType

`func (o *DestinationWalletAddressDetail) GetDestinationType() DestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *DestinationWalletAddressDetail) GetDestinationTypeOk() (*DestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *DestinationWalletAddressDetail) SetDestinationType(v DestinationType)`

SetDestinationType sets DestinationType field to given value.


### GetMerchantId

`func (o *DestinationWalletAddressDetail) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *DestinationWalletAddressDetail) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *DestinationWalletAddressDetail) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *DestinationWalletAddressDetail) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetWalletAddressId

`func (o *DestinationWalletAddressDetail) GetWalletAddressId() string`

GetWalletAddressId returns the WalletAddressId field if non-nil, zero value otherwise.

### GetWalletAddressIdOk

`func (o *DestinationWalletAddressDetail) GetWalletAddressIdOk() (*string, bool)`

GetWalletAddressIdOk returns a tuple with the WalletAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddressId

`func (o *DestinationWalletAddressDetail) SetWalletAddressId(v string)`

SetWalletAddressId sets WalletAddressId field to given value.


### GetAddress

`func (o *DestinationWalletAddressDetail) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DestinationWalletAddressDetail) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DestinationWalletAddressDetail) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChainId

`func (o *DestinationWalletAddressDetail) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *DestinationWalletAddressDetail) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *DestinationWalletAddressDetail) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetUpdatedTimestamp

`func (o *DestinationWalletAddressDetail) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *DestinationWalletAddressDetail) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *DestinationWalletAddressDetail) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


