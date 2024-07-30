# TransferParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a transaction request. The request ID is provided by you and must be unique within your organization. | 
**Source** | [**TransferSource**](TransferSource.md) |  | 
**TokenId** | **string** | The token ID of the transferred token. You can retrieve the IDs of all the tokens you can use by calling [List enabled tokens](/v2/api-references/wallets/list-enabled-tokens). | 
**Destination** | [**TransferDestination**](TransferDestination.md) |  | 
**CategoryNames** | Pointer to **[]string** | The custom category for you to identify your transactions. | [optional] 
**Description** | Pointer to **string** | The description of the transfer. | [optional] 
**Fee** | Pointer to [**TransactionRequestFee**](TransactionRequestFee.md) |  | [optional] 

## Methods

### NewTransferParams

`func NewTransferParams(requestId string, source TransferSource, tokenId string, destination TransferDestination, ) *TransferParams`

NewTransferParams instantiates a new TransferParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferParamsWithDefaults

`func NewTransferParamsWithDefaults() *TransferParams`

NewTransferParamsWithDefaults instantiates a new TransferParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *TransferParams) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransferParams) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransferParams) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetSource

`func (o *TransferParams) GetSource() TransferSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TransferParams) GetSourceOk() (*TransferSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TransferParams) SetSource(v TransferSource)`

SetSource sets Source field to given value.


### GetTokenId

`func (o *TransferParams) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TransferParams) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TransferParams) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetDestination

`func (o *TransferParams) GetDestination() TransferDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *TransferParams) GetDestinationOk() (*TransferDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *TransferParams) SetDestination(v TransferDestination)`

SetDestination sets Destination field to given value.


### GetCategoryNames

`func (o *TransferParams) GetCategoryNames() []string`

GetCategoryNames returns the CategoryNames field if non-nil, zero value otherwise.

### GetCategoryNamesOk

`func (o *TransferParams) GetCategoryNamesOk() (*[]string, bool)`

GetCategoryNamesOk returns a tuple with the CategoryNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryNames

`func (o *TransferParams) SetCategoryNames(v []string)`

SetCategoryNames sets CategoryNames field to given value.

### HasCategoryNames

`func (o *TransferParams) HasCategoryNames() bool`

HasCategoryNames returns a boolean if a field has been set.

### GetDescription

`func (o *TransferParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransferParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransferParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransferParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFee

`func (o *TransferParams) GetFee() TransactionRequestFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *TransferParams) GetFeeOk() (*TransactionRequestFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *TransferParams) SetFee(v TransactionRequestFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *TransferParams) HasFee() bool`

HasFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


