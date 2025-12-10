# CounterpartyWalletAddressDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CounterpartyId** | **string** | The counterparty ID. | 
**CounterpartyName** | **string** | The name of the counterparty. | 
**CounterpartyType** | [**CounterpartyType**](CounterpartyType.md) |  | 
**WalletAddressId** | **string** | The wallet address ID. | 
**Address** | **string** | The wallet address. | 
**ChainId** | **string** | The chain ID of the cryptocurrency. | 
**UpdatedTimestamp** | **int32** | The updated time of the wallet address, represented as a UNIX timestamp in seconds. | 

## Methods

### NewCounterpartyWalletAddressDetail

`func NewCounterpartyWalletAddressDetail(counterpartyId string, counterpartyName string, counterpartyType CounterpartyType, walletAddressId string, address string, chainId string, updatedTimestamp int32, ) *CounterpartyWalletAddressDetail`

NewCounterpartyWalletAddressDetail instantiates a new CounterpartyWalletAddressDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCounterpartyWalletAddressDetailWithDefaults

`func NewCounterpartyWalletAddressDetailWithDefaults() *CounterpartyWalletAddressDetail`

NewCounterpartyWalletAddressDetailWithDefaults instantiates a new CounterpartyWalletAddressDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounterpartyId

`func (o *CounterpartyWalletAddressDetail) GetCounterpartyId() string`

GetCounterpartyId returns the CounterpartyId field if non-nil, zero value otherwise.

### GetCounterpartyIdOk

`func (o *CounterpartyWalletAddressDetail) GetCounterpartyIdOk() (*string, bool)`

GetCounterpartyIdOk returns a tuple with the CounterpartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyId

`func (o *CounterpartyWalletAddressDetail) SetCounterpartyId(v string)`

SetCounterpartyId sets CounterpartyId field to given value.


### GetCounterpartyName

`func (o *CounterpartyWalletAddressDetail) GetCounterpartyName() string`

GetCounterpartyName returns the CounterpartyName field if non-nil, zero value otherwise.

### GetCounterpartyNameOk

`func (o *CounterpartyWalletAddressDetail) GetCounterpartyNameOk() (*string, bool)`

GetCounterpartyNameOk returns a tuple with the CounterpartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyName

`func (o *CounterpartyWalletAddressDetail) SetCounterpartyName(v string)`

SetCounterpartyName sets CounterpartyName field to given value.


### GetCounterpartyType

`func (o *CounterpartyWalletAddressDetail) GetCounterpartyType() CounterpartyType`

GetCounterpartyType returns the CounterpartyType field if non-nil, zero value otherwise.

### GetCounterpartyTypeOk

`func (o *CounterpartyWalletAddressDetail) GetCounterpartyTypeOk() (*CounterpartyType, bool)`

GetCounterpartyTypeOk returns a tuple with the CounterpartyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyType

`func (o *CounterpartyWalletAddressDetail) SetCounterpartyType(v CounterpartyType)`

SetCounterpartyType sets CounterpartyType field to given value.


### GetWalletAddressId

`func (o *CounterpartyWalletAddressDetail) GetWalletAddressId() string`

GetWalletAddressId returns the WalletAddressId field if non-nil, zero value otherwise.

### GetWalletAddressIdOk

`func (o *CounterpartyWalletAddressDetail) GetWalletAddressIdOk() (*string, bool)`

GetWalletAddressIdOk returns a tuple with the WalletAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddressId

`func (o *CounterpartyWalletAddressDetail) SetWalletAddressId(v string)`

SetWalletAddressId sets WalletAddressId field to given value.


### GetAddress

`func (o *CounterpartyWalletAddressDetail) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CounterpartyWalletAddressDetail) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CounterpartyWalletAddressDetail) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChainId

`func (o *CounterpartyWalletAddressDetail) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *CounterpartyWalletAddressDetail) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *CounterpartyWalletAddressDetail) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetUpdatedTimestamp

`func (o *CounterpartyWalletAddressDetail) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *CounterpartyWalletAddressDetail) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *CounterpartyWalletAddressDetail) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


