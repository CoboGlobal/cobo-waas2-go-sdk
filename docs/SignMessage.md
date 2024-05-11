# SignMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Unique id of the request. | 
**RequestType** | **string** |  | 
**FromWalletId** | Pointer to **string** | Unique id of the wallet to sign message. | [optional] 
**FromAddressStr** | Pointer to **string** | signing address | [optional] 
**ChainId** | **string** | The blockchain on which the token operates. | 
**Message** | Pointer to **string** | Raw data to be signed, Base 64 encoded | [optional] 
**StructuredData** | Pointer to **string** | Structured data to be signed, JSON encoded | [optional] 
**MpcUsedKeyGroup** | Pointer to [**MpcSigningGroup**](MpcSigningGroup.md) |  | [optional] 

## Methods

### NewSignMessage

`func NewSignMessage(requestId string, requestType string, chainId string, ) *SignMessage`

NewSignMessage instantiates a new SignMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignMessageWithDefaults

`func NewSignMessageWithDefaults() *SignMessage`

NewSignMessageWithDefaults instantiates a new SignMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *SignMessage) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SignMessage) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SignMessage) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetRequestType

`func (o *SignMessage) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *SignMessage) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *SignMessage) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.


### GetFromWalletId

`func (o *SignMessage) GetFromWalletId() string`

GetFromWalletId returns the FromWalletId field if non-nil, zero value otherwise.

### GetFromWalletIdOk

`func (o *SignMessage) GetFromWalletIdOk() (*string, bool)`

GetFromWalletIdOk returns a tuple with the FromWalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromWalletId

`func (o *SignMessage) SetFromWalletId(v string)`

SetFromWalletId sets FromWalletId field to given value.

### HasFromWalletId

`func (o *SignMessage) HasFromWalletId() bool`

HasFromWalletId returns a boolean if a field has been set.

### GetFromAddressStr

`func (o *SignMessage) GetFromAddressStr() string`

GetFromAddressStr returns the FromAddressStr field if non-nil, zero value otherwise.

### GetFromAddressStrOk

`func (o *SignMessage) GetFromAddressStrOk() (*string, bool)`

GetFromAddressStrOk returns a tuple with the FromAddressStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddressStr

`func (o *SignMessage) SetFromAddressStr(v string)`

SetFromAddressStr sets FromAddressStr field to given value.

### HasFromAddressStr

`func (o *SignMessage) HasFromAddressStr() bool`

HasFromAddressStr returns a boolean if a field has been set.

### GetChainId

`func (o *SignMessage) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *SignMessage) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *SignMessage) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetMessage

`func (o *SignMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SignMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SignMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SignMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStructuredData

`func (o *SignMessage) GetStructuredData() string`

GetStructuredData returns the StructuredData field if non-nil, zero value otherwise.

### GetStructuredDataOk

`func (o *SignMessage) GetStructuredDataOk() (*string, bool)`

GetStructuredDataOk returns a tuple with the StructuredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredData

`func (o *SignMessage) SetStructuredData(v string)`

SetStructuredData sets StructuredData field to given value.

### HasStructuredData

`func (o *SignMessage) HasStructuredData() bool`

HasStructuredData returns a boolean if a field has been set.

### GetMpcUsedKeyGroup

`func (o *SignMessage) GetMpcUsedKeyGroup() MpcSigningGroup`

GetMpcUsedKeyGroup returns the MpcUsedKeyGroup field if non-nil, zero value otherwise.

### GetMpcUsedKeyGroupOk

`func (o *SignMessage) GetMpcUsedKeyGroupOk() (*MpcSigningGroup, bool)`

GetMpcUsedKeyGroupOk returns a tuple with the MpcUsedKeyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpcUsedKeyGroup

`func (o *SignMessage) SetMpcUsedKeyGroup(v MpcSigningGroup)`

SetMpcUsedKeyGroup sets MpcUsedKeyGroup field to given value.

### HasMpcUsedKeyGroup

`func (o *SignMessage) HasMpcUsedKeyGroup() bool`

HasMpcUsedKeyGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


