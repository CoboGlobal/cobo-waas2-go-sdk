# SafeTxExtraData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** | The recipient address of the transaction | 
**Value** | **string** | Readable transaction value (e.g., 1 ETH) | 
**Data** | **string** | The transaction data | 
**DomainHash** | **string** | EIP712 structured data domain hash | 
**MessageHash** | **string** | Hash of the structured message | 
**SafeAddress** | **string** | Address of the Safe contract | 
**SafeTxHash** | **string** | Hash of the Safe transaction | 
**SafeNonce** | **int32** | Safe transaction nonce | 
**Operation** | **string** | Type of operation performed in the transaction | 
**GasTokenAddr** | Pointer to **string** | Address of the gas token | [optional] 
**SafeTxGas** | Pointer to **int32** | Gas used for the Safe transaction | [optional] 
**BaseGas** | Pointer to **int32** | Base gas for the transaction | [optional] 
**GasPrice** | Pointer to **string** | Gas price used in the transaction | [optional] 
**RefundReceiver** | Pointer to **string** | Address to receive the gas refund | [optional] 
**ToContractName** | Pointer to **string** | Name of the recipient contract (if available) | [optional] 
**DecodedData** | Pointer to [**SafeTxDecodedData**](SafeTxDecodedData.md) |  | [optional] 
**Signature** | Pointer to **string** | Signature of the transaction (if signed by Cobo Signer) | [optional] 
**Wei** | Pointer to **NullableString** | Transaction amount in Wei | [optional] 

## Methods

### NewSafeTxExtraData

`func NewSafeTxExtraData(to string, value string, data string, domainHash string, messageHash string, safeAddress string, safeTxHash string, safeNonce int32, operation string, ) *SafeTxExtraData`

NewSafeTxExtraData instantiates a new SafeTxExtraData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafeTxExtraDataWithDefaults

`func NewSafeTxExtraDataWithDefaults() *SafeTxExtraData`

NewSafeTxExtraDataWithDefaults instantiates a new SafeTxExtraData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *SafeTxExtraData) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SafeTxExtraData) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SafeTxExtraData) SetTo(v string)`

SetTo sets To field to given value.


### GetValue

`func (o *SafeTxExtraData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SafeTxExtraData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SafeTxExtraData) SetValue(v string)`

SetValue sets Value field to given value.


### GetData

`func (o *SafeTxExtraData) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SafeTxExtraData) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SafeTxExtraData) SetData(v string)`

SetData sets Data field to given value.


### GetDomainHash

`func (o *SafeTxExtraData) GetDomainHash() string`

GetDomainHash returns the DomainHash field if non-nil, zero value otherwise.

### GetDomainHashOk

`func (o *SafeTxExtraData) GetDomainHashOk() (*string, bool)`

GetDomainHashOk returns a tuple with the DomainHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainHash

`func (o *SafeTxExtraData) SetDomainHash(v string)`

SetDomainHash sets DomainHash field to given value.


### GetMessageHash

`func (o *SafeTxExtraData) GetMessageHash() string`

GetMessageHash returns the MessageHash field if non-nil, zero value otherwise.

### GetMessageHashOk

`func (o *SafeTxExtraData) GetMessageHashOk() (*string, bool)`

GetMessageHashOk returns a tuple with the MessageHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageHash

`func (o *SafeTxExtraData) SetMessageHash(v string)`

SetMessageHash sets MessageHash field to given value.


### GetSafeAddress

`func (o *SafeTxExtraData) GetSafeAddress() string`

GetSafeAddress returns the SafeAddress field if non-nil, zero value otherwise.

### GetSafeAddressOk

`func (o *SafeTxExtraData) GetSafeAddressOk() (*string, bool)`

GetSafeAddressOk returns a tuple with the SafeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeAddress

`func (o *SafeTxExtraData) SetSafeAddress(v string)`

SetSafeAddress sets SafeAddress field to given value.


### GetSafeTxHash

`func (o *SafeTxExtraData) GetSafeTxHash() string`

GetSafeTxHash returns the SafeTxHash field if non-nil, zero value otherwise.

### GetSafeTxHashOk

`func (o *SafeTxExtraData) GetSafeTxHashOk() (*string, bool)`

GetSafeTxHashOk returns a tuple with the SafeTxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeTxHash

`func (o *SafeTxExtraData) SetSafeTxHash(v string)`

SetSafeTxHash sets SafeTxHash field to given value.


### GetSafeNonce

`func (o *SafeTxExtraData) GetSafeNonce() int32`

GetSafeNonce returns the SafeNonce field if non-nil, zero value otherwise.

### GetSafeNonceOk

`func (o *SafeTxExtraData) GetSafeNonceOk() (*int32, bool)`

GetSafeNonceOk returns a tuple with the SafeNonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeNonce

`func (o *SafeTxExtraData) SetSafeNonce(v int32)`

SetSafeNonce sets SafeNonce field to given value.


### GetOperation

`func (o *SafeTxExtraData) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *SafeTxExtraData) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *SafeTxExtraData) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetGasTokenAddr

`func (o *SafeTxExtraData) GetGasTokenAddr() string`

GetGasTokenAddr returns the GasTokenAddr field if non-nil, zero value otherwise.

### GetGasTokenAddrOk

`func (o *SafeTxExtraData) GetGasTokenAddrOk() (*string, bool)`

GetGasTokenAddrOk returns a tuple with the GasTokenAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasTokenAddr

`func (o *SafeTxExtraData) SetGasTokenAddr(v string)`

SetGasTokenAddr sets GasTokenAddr field to given value.

### HasGasTokenAddr

`func (o *SafeTxExtraData) HasGasTokenAddr() bool`

HasGasTokenAddr returns a boolean if a field has been set.

### GetSafeTxGas

`func (o *SafeTxExtraData) GetSafeTxGas() int32`

GetSafeTxGas returns the SafeTxGas field if non-nil, zero value otherwise.

### GetSafeTxGasOk

`func (o *SafeTxExtraData) GetSafeTxGasOk() (*int32, bool)`

GetSafeTxGasOk returns a tuple with the SafeTxGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeTxGas

`func (o *SafeTxExtraData) SetSafeTxGas(v int32)`

SetSafeTxGas sets SafeTxGas field to given value.

### HasSafeTxGas

`func (o *SafeTxExtraData) HasSafeTxGas() bool`

HasSafeTxGas returns a boolean if a field has been set.

### GetBaseGas

`func (o *SafeTxExtraData) GetBaseGas() int32`

GetBaseGas returns the BaseGas field if non-nil, zero value otherwise.

### GetBaseGasOk

`func (o *SafeTxExtraData) GetBaseGasOk() (*int32, bool)`

GetBaseGasOk returns a tuple with the BaseGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseGas

`func (o *SafeTxExtraData) SetBaseGas(v int32)`

SetBaseGas sets BaseGas field to given value.

### HasBaseGas

`func (o *SafeTxExtraData) HasBaseGas() bool`

HasBaseGas returns a boolean if a field has been set.

### GetGasPrice

`func (o *SafeTxExtraData) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *SafeTxExtraData) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *SafeTxExtraData) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *SafeTxExtraData) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### GetRefundReceiver

`func (o *SafeTxExtraData) GetRefundReceiver() string`

GetRefundReceiver returns the RefundReceiver field if non-nil, zero value otherwise.

### GetRefundReceiverOk

`func (o *SafeTxExtraData) GetRefundReceiverOk() (*string, bool)`

GetRefundReceiverOk returns a tuple with the RefundReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundReceiver

`func (o *SafeTxExtraData) SetRefundReceiver(v string)`

SetRefundReceiver sets RefundReceiver field to given value.

### HasRefundReceiver

`func (o *SafeTxExtraData) HasRefundReceiver() bool`

HasRefundReceiver returns a boolean if a field has been set.

### GetToContractName

`func (o *SafeTxExtraData) GetToContractName() string`

GetToContractName returns the ToContractName field if non-nil, zero value otherwise.

### GetToContractNameOk

`func (o *SafeTxExtraData) GetToContractNameOk() (*string, bool)`

GetToContractNameOk returns a tuple with the ToContractName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToContractName

`func (o *SafeTxExtraData) SetToContractName(v string)`

SetToContractName sets ToContractName field to given value.

### HasToContractName

`func (o *SafeTxExtraData) HasToContractName() bool`

HasToContractName returns a boolean if a field has been set.

### GetDecodedData

`func (o *SafeTxExtraData) GetDecodedData() SafeTxDecodedData`

GetDecodedData returns the DecodedData field if non-nil, zero value otherwise.

### GetDecodedDataOk

`func (o *SafeTxExtraData) GetDecodedDataOk() (*SafeTxDecodedData, bool)`

GetDecodedDataOk returns a tuple with the DecodedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecodedData

`func (o *SafeTxExtraData) SetDecodedData(v SafeTxDecodedData)`

SetDecodedData sets DecodedData field to given value.

### HasDecodedData

`func (o *SafeTxExtraData) HasDecodedData() bool`

HasDecodedData returns a boolean if a field has been set.

### GetSignature

`func (o *SafeTxExtraData) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *SafeTxExtraData) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *SafeTxExtraData) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *SafeTxExtraData) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetWei

`func (o *SafeTxExtraData) GetWei() string`

GetWei returns the Wei field if non-nil, zero value otherwise.

### GetWeiOk

`func (o *SafeTxExtraData) GetWeiOk() (*string, bool)`

GetWeiOk returns a tuple with the Wei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWei

`func (o *SafeTxExtraData) SetWei(v string)`

SetWei sets Wei field to given value.

### HasWei

`func (o *SafeTxExtraData) HasWei() bool`

HasWei returns a boolean if a field has been set.

### SetWeiNil

`func (o *SafeTxExtraData) SetWeiNil(b bool)`

 SetWeiNil sets the value for Wei to be an explicit nil

### UnsetWei
`func (o *SafeTxExtraData) UnsetWei()`

UnsetWei ensures that no value is present for Wei, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


