# GetMaxTransferableValueWithFeeModelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | **string** | The token ID of the transferred token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](https://www.cobo.com/developers/v2/api-references/wallets/list-enabled-tokens). For transfers from Exchange Wallets, this property value represents the asset ID. | 
**Fee** | [**TransactionRequestFee**](TransactionRequestFee.md) |  | 
**ToAddress** | **string** | The recipient&#39;s address. | 
**FromAddress** | Pointer to **string** | The sender&#39;s address. This property is required when using an EVM address in an MPC Wallet. | [optional] 

## Methods

### NewGetMaxTransferableValueWithFeeModelRequest

`func NewGetMaxTransferableValueWithFeeModelRequest(tokenId string, fee TransactionRequestFee, toAddress string, ) *GetMaxTransferableValueWithFeeModelRequest`

NewGetMaxTransferableValueWithFeeModelRequest instantiates a new GetMaxTransferableValueWithFeeModelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMaxTransferableValueWithFeeModelRequestWithDefaults

`func NewGetMaxTransferableValueWithFeeModelRequestWithDefaults() *GetMaxTransferableValueWithFeeModelRequest`

NewGetMaxTransferableValueWithFeeModelRequestWithDefaults instantiates a new GetMaxTransferableValueWithFeeModelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *GetMaxTransferableValueWithFeeModelRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetMaxTransferableValueWithFeeModelRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetMaxTransferableValueWithFeeModelRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetFee

`func (o *GetMaxTransferableValueWithFeeModelRequest) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetMaxTransferableValueWithFeeModelRequest) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetMaxTransferableValueWithFeeModelRequest) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.


### GetToAddress

`func (o *GetMaxTransferableValueWithFeeModelRequest) GetToAddress() string`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *GetMaxTransferableValueWithFeeModelRequest) GetToAddressOk() (*string, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *GetMaxTransferableValueWithFeeModelRequest) SetToAddress(v string)`

SetToAddress sets ToAddress field to given value.


### GetFromAddress

`func (o *GetMaxTransferableValueWithFeeModelRequest) GetFromAddress() string`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *GetMaxTransferableValueWithFeeModelRequest) GetFromAddressOk() (*string, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *GetMaxTransferableValueWithFeeModelRequest) SetFromAddress(v string)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *GetMaxTransferableValueWithFeeModelRequest) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


