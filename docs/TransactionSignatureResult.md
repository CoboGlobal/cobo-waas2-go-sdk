# TransactionSignatureResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultType** | Pointer to [**TransactionResultType**](TransactionResultType.md) |  | [optional] 
**Signature** | **string** | The raw data of the signature. | 

## Methods

### NewTransactionSignatureResult

`func NewTransactionSignatureResult(signature string, ) *TransactionSignatureResult`

NewTransactionSignatureResult instantiates a new TransactionSignatureResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionSignatureResultWithDefaults

`func NewTransactionSignatureResultWithDefaults() *TransactionSignatureResult`

NewTransactionSignatureResultWithDefaults instantiates a new TransactionSignatureResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultType

`func (o *TransactionSignatureResult) GetResultType() TransactionResultType`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *TransactionSignatureResult) GetResultTypeOk() (*TransactionResultType, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *TransactionSignatureResult) SetResultType(v TransactionResultType)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *TransactionSignatureResult) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetSignature

`func (o *TransactionSignatureResult) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *TransactionSignatureResult) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *TransactionSignatureResult) SetSignature(v string)`

SetSignature sets Signature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


