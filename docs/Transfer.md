# Transfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | The request ID that is used to track a withdrawal request. The request ID is provided by you and must be unique within your organization. | 
**Source** | [**TransferSource**](TransferSource.md) |  | 
**TokenId** | **string** | The token ID, which is the unique identifier of a token. You can retrieve the IDs of all the tokens you can use by calling [List organization enabled tokens](/v2/api-references/wallets/list-organization-enabled-tokens). | 
**Destination** | [**TransferDestination**](TransferDestination.md) |  | 
**CategoryNames** | Pointer to **[]string** | The custom category for you to identify your transactions. | [optional] 
**Description** | Pointer to **string** | The description of the transfer. | [optional] 
**Fee** | Pointer to [**TransactionTransferFee**](TransactionTransferFee.md) |  | [optional] 

## Methods

### NewTransfer

`func NewTransfer(requestId string, source TransferSource, tokenId string, destination TransferDestination, ) *Transfer`

NewTransfer instantiates a new Transfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferWithDefaults

`func NewTransferWithDefaults() *Transfer`

NewTransferWithDefaults instantiates a new Transfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *Transfer) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Transfer) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Transfer) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetSource

`func (o *Transfer) GetSource() TransferSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Transfer) GetSourceOk() (*TransferSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Transfer) SetSource(v TransferSource)`

SetSource sets Source field to given value.


### GetTokenId

`func (o *Transfer) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *Transfer) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *Transfer) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetDestination

`func (o *Transfer) GetDestination() TransferDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Transfer) GetDestinationOk() (*TransferDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Transfer) SetDestination(v TransferDestination)`

SetDestination sets Destination field to given value.


### GetCategoryNames

`func (o *Transfer) GetCategoryNames() []string`

GetCategoryNames returns the CategoryNames field if non-nil, zero value otherwise.

### GetCategoryNamesOk

`func (o *Transfer) GetCategoryNamesOk() (*[]string, bool)`

GetCategoryNamesOk returns a tuple with the CategoryNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryNames

`func (o *Transfer) SetCategoryNames(v []string)`

SetCategoryNames sets CategoryNames field to given value.

### HasCategoryNames

`func (o *Transfer) HasCategoryNames() bool`

HasCategoryNames returns a boolean if a field has been set.

### GetDescription

`func (o *Transfer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Transfer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Transfer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Transfer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFee

`func (o *Transfer) GetFee() TransactionTransferFee`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *Transfer) GetFeeOk() (*TransactionTransferFee, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *Transfer) SetFee(v TransactionTransferFee)`

SetFee sets Fee field to given value.

### HasFee

`func (o *Transfer) HasFee() bool`

HasFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


