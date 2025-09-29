# CryptoAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token identifier (e.g., ETH_USDT, TRON_USDT) that this address is associated with. | 
**CryptoAddressId** | **string** | Unique identifier for the pre-approved crypto address, used to reference the address securely in requests. This ID is returned by the system and should be used instead of the raw blockchain address in API calls. | 
**Address** | **string** | The actual blockchain address to which funds will be transferred. This is for display purposes only; external clients should always use address_id to refer to the address in secure operations. | 
**Label** | Pointer to **string** | A human-readable label or alias for the crypto address, set by the merchant or platform operator. This field is optional and intended to help distinguish addresses by usage or purpose (e.g., \&quot;Main Payout Wallet\&quot;, \&quot;Cold Wallet\&quot;). | [optional] 
**CreatedTimestamp** | Pointer to **int32** | The created time of the crypto address, represented as a UNIX timestamp in seconds. | [optional] 
**UpdatedTimestamp** | Pointer to **int32** | The updated time of the crypto address, represented as a UNIX timestamp in seconds. | [optional] 

## Methods

### NewCryptoAddress

`func NewCryptoAddress(tokenId string, cryptoAddressId string, address string, ) *CryptoAddress`

NewCryptoAddress instantiates a new CryptoAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoAddressWithDefaults

`func NewCryptoAddressWithDefaults() *CryptoAddress`

NewCryptoAddressWithDefaults instantiates a new CryptoAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *CryptoAddress) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *CryptoAddress) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *CryptoAddress) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetCryptoAddressId

`func (o *CryptoAddress) GetCryptoAddressId() string`

GetCryptoAddressId returns the CryptoAddressId field if non-nil, zero value otherwise.

### GetCryptoAddressIdOk

`func (o *CryptoAddress) GetCryptoAddressIdOk() (*string, bool)`

GetCryptoAddressIdOk returns a tuple with the CryptoAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAddressId

`func (o *CryptoAddress) SetCryptoAddressId(v string)`

SetCryptoAddressId sets CryptoAddressId field to given value.


### GetAddress

`func (o *CryptoAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CryptoAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CryptoAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetLabel

`func (o *CryptoAddress) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CryptoAddress) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CryptoAddress) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CryptoAddress) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetCreatedTimestamp

`func (o *CryptoAddress) GetCreatedTimestamp() int32`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *CryptoAddress) GetCreatedTimestampOk() (*int32, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *CryptoAddress) SetCreatedTimestamp(v int32)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *CryptoAddress) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetUpdatedTimestamp

`func (o *CryptoAddress) GetUpdatedTimestamp() int32`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *CryptoAddress) GetUpdatedTimestampOk() (*int32, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *CryptoAddress) SetUpdatedTimestamp(v int32)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.

### HasUpdatedTimestamp

`func (o *CryptoAddress) HasUpdatedTimestamp() bool`

HasUpdatedTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


